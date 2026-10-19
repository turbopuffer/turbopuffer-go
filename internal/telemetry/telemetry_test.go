package telemetry

import (
	"bytes"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/turbopuffer/turbopuffer-go/v2/internal/observer"
)

const (
	queryURL   = "/v2/namespaces/my-ns/query"
	queryRoute = "POST /v2/namespaces/{ns}/query"
)

func addRecord(b *buffer, status int) {
	b.add(record{Route: queryRoute, Status: status})
}

type payload struct {
	Version int              `json:"version"`
	Records []map[string]any `json:"records"`
	Dropped *int             `json:"dropped"`
}

func decodeHeader(t *testing.T, header string) payload {
	t.Helper()
	var p payload
	if err := json.Unmarshal([]byte(header), &p); err != nil {
		t.Fatalf("invalid header %q: %v", header, err)
	}
	if p.Version != 1 {
		t.Fatalf("expected version 1, got %d", p.Version)
	}
	return p
}

func takePayload(t *testing.T, b *buffer) payload {
	t.Helper()
	bt := b.take()
	if bt == nil {
		t.Fatal("expected a batch")
	}
	return decodeHeader(t, bt.header)
}

func statuses(records []map[string]any) []int {
	var out []int
	for _, r := range records {
		out = append(out, int(r["status"].(float64)))
	}
	return out
}

func rateHeader(value string) http.Header {
	h := http.Header{}
	h.Set(SampleRateHeader, value)
	return h
}

func setRandom(t *testing.T, value float64) {
	t.Helper()
	prev := random
	random = func() float64 { return value }
	t.Cleanup(func() { random = prev })
}

func TestEnabledByEnv(t *testing.T) {
	for value, want := range map[string]bool{
		"1": true, "true": true, " ON ": true, "yes": true,
		"0": false, "false": false, "off": false, "": false,
	} {
		t.Setenv(EnvVar, value)
		if got := enabledByEnv(); got != want {
			t.Errorf("%s=%q: expected %t, got %t", EnvVar, value, want, got)
		}
		if got := len(EnvOptions()) > 0; got != want {
			t.Errorf("%s=%q: expected options %t, got %t", EnvVar, value, want, got)
		}
	}
}

func TestEnabledByInvalidEnvWarns(t *testing.T) {
	var logs bytes.Buffer
	prev := slog.Default()
	slog.SetDefault(slog.New(slog.NewTextHandler(&logs, nil)))
	t.Cleanup(func() { slog.SetDefault(prev) })

	t.Setenv(EnvVar, "sometimes")
	if enabledByEnv() {
		t.Error("expected an invalid value to disable telemetry")
	}
	if !strings.Contains(logs.String(), EnvVar) {
		t.Errorf("expected a warning naming %s, got %q", EnvVar, logs.String())
	}
}

func TestRouteTemplate(t *testing.T) {
	for _, tc := range []struct{ method, url, want string }{
		{"post", queryURL, queryRoute},
		{"post", "/v2/namespaces/my-ns/query?stainless_overload=multiQuery", "POST /v2/namespaces/{ns}/query?stainless_overload=multiQuery"},
		{"post", "/v2/namespaces/my%2Fns", "POST /v2/namespaces/{ns}"},
		{"get", "/v1/namespaces", "GET /v1/namespaces"},
		{"get", "/v1/namespaces/my-ns/operations/abc123?stainless_overload=pollCopyFrom", "GET /v1/namespaces/{ns}/operations/{token}?stainless_overload=pollCopyFrom"},
		{"get", "https://gcp-us-central1.turbopuffer.com/v1/namespaces/my-ns/metadata", "GET /v1/namespaces/{ns}/metadata"},
		{"get", "http://localhost:5000/v1/namespaces/my-ns/operations/op-abc", "GET /v1/namespaces/{ns}/operations/{token}"},
	} {
		if got := routeTemplate(tc.method, tc.url); got != tc.want {
			t.Errorf("routeTemplate(%q, %q): expected %q, got %q", tc.method, tc.url, tc.want, got)
		}
	}
}

func TestHeaderEncoding(t *testing.T) {
	reused := false
	b := newBuffer()
	b.entries = []entry{{
		recordedAt: time.Now().Add(-2 * time.Second),
		record: record{
			Route:         queryRoute,
			Status:        0,
			Error:         "timeout",
			Retries:       2,
			SampleRate:    0.25,
			TotalUs:       micros(3 * time.Millisecond),
			ConnectUs:     micros(0),
			ServerTotalUs: micros(time.Millisecond),
			Reused:        &reused,
		},
	}}
	b.dropped = 3

	header := b.take().header
	var compact bytes.Buffer
	if err := json.Compact(&compact, []byte(header)); err != nil || compact.String() != header {
		t.Errorf("expected compact JSON, got %q", header)
	}
	p := decodeHeader(t, header)
	if p.Dropped == nil || *p.Dropped != 3 {
		t.Errorf("expected 3 dropped, got %v", p.Dropped)
	}
	rec := p.Records[0]
	age := rec["age_ms"].(float64)
	if age < 2000 || age > 60_000 {
		t.Errorf("expected an age of about 2000ms, got %v", age)
	}
	delete(rec, "age_ms")
	want := map[string]any{
		"route":           queryRoute,
		"status":          0.0,
		"error":           "timeout",
		"retries":         2.0,
		"sample_rate":     0.25,
		"total_us":        3000.0,
		"connect_us":      0.0,
		"server_total_us": 1000.0,
		"reused":          false,
	}
	if !reflect.DeepEqual(rec, want) {
		t.Errorf("expected record %v, got %v", want, rec)
	}

	// Zero and unset values are left out.
	addRecord(b, 200)
	header = regexp.MustCompile(`"age_ms":\d+`).ReplaceAllString(b.take().header, `"age_ms":0`)
	if want := `{"version":1,"records":[{"route":"POST /v2/namespaces/{ns}/query","status":200,"age_ms":0}]}`; header != want {
		t.Errorf("expected %s, got %s", want, header)
	}
}

func TestFullBufferDropsOldestRecords(t *testing.T) {
	b := newBuffer()
	for status := range maxBufferedRecords + 5 {
		addRecord(b, status)
	}

	first := takePayload(t, b)
	if first.Dropped == nil || *first.Dropped != 5 {
		t.Errorf("expected 5 dropped, got %v", first.Dropped)
	}
	if len(first.Records) != maxRecordsPerHeader || statuses(first.Records)[0] != 5 {
		t.Errorf("expected %d records from status 5, got %v", maxRecordsPerHeader, statuses(first.Records))
	}

	second := takePayload(t, b)
	if second.Dropped != nil {
		t.Errorf("expected dropped to be omitted, got %d", *second.Dropped)
	}
	if got := statuses(second.Records)[0]; got != 5+maxRecordsPerHeader {
		t.Errorf("expected records from status %d, got %d", 5+maxRecordsPerHeader, got)
	}
}

func TestRestorePutsRecordsBackInOrder(t *testing.T) {
	b := newBuffer()
	for status := range 3 {
		addRecord(b, status)
	}
	bt := b.take()

	addRecord(b, 99)
	b.restore(bt)

	if got := statuses(takePayload(t, b).Records); !reflect.DeepEqual(got, []int{0, 1, 2, 99}) {
		t.Errorf("expected statuses [0 1 2 99], got %v", got)
	}
	if b.take() != nil {
		t.Error("expected the buffer to be empty")
	}
}

func TestRestoreIntoFullBufferDropsOldest(t *testing.T) {
	b := newBuffer()
	for status := range maxBufferedRecords {
		addRecord(b, status)
	}
	bt := b.take()
	b.dropped = 2
	bt.dropped = 1
	for status := range 15 {
		addRecord(b, 1000+status)
	}

	// Only the newest 5 records of the batch fit.
	b.restore(bt)

	p := takePayload(t, b)
	if p.Dropped == nil || *p.Dropped != 2+1+15 {
		t.Errorf("expected 18 dropped, got %v", p.Dropped)
	}
	if got := statuses(p.Records)[:6]; !reflect.DeepEqual(got, []int{15, 16, 17, 18, 19, 20}) {
		t.Errorf("expected statuses from 15, got %v", got)
	}
}

func TestSampleRateFromResponse(t *testing.T) {
	b := newBuffer()
	if b.rate() != 1 {
		t.Errorf("expected a default rate of 1, got %v", b.rate())
	}
	for _, value := range []string{"abc", "1.5", "-0.1", "NaN", "Inf", ""} {
		b.observeResponse(rateHeader("0.5"))
		b.observeResponse(rateHeader(value))
		if b.rate() != 0.5 {
			t.Errorf("expected %q to be ignored, got rate %v", value, b.rate())
		}
	}
	b.observeResponse(http.Header{})
	if b.rate() != 0.5 {
		t.Errorf("expected a missing header to be ignored, got rate %v", b.rate())
	}
	for value, rate := range map[string]float64{"0": 0, "0.001": 0.001, "1": 1} {
		b.observeResponse(rateHeader(value))
		if b.rate() != rate {
			t.Errorf("expected rate %v from %q, got %v", rate, value, b.rate())
		}
	}
}

func TestSampling(t *testing.T) {
	result := &observer.RequestResult{Timings: &observer.Timings{AttemptBegin: time.Now()}}
	finish := func(b *buffer) {
		(&requestScope{buffer: b, info: observer.RequestInfo{Method: "POST", URL: queryURL}}).Finish(result)
	}

	b := newBuffer()
	b.observeResponse(rateHeader("0.25"))
	setRandom(t, 0.2)
	finish(b)
	setRandom(t, 0.3)
	finish(b)

	records := takePayload(t, b).Records
	if len(records) != 1 || records[0]["sample_rate"] != 0.25 {
		t.Errorf("expected one record sampled at 0.25, got %v", records)
	}

	b.observeResponse(rateHeader("0"))
	setRandom(t, 0)
	finish(b)
	if b.take() != nil {
		t.Error("expected nothing to be recorded at rate 0")
	}
}

func TestDurations(t *testing.T) {
	at := func(us int) time.Time { return time.Unix(0, 0).Add(time.Duration(us) * time.Microsecond) }
	durations := func(t *observer.Timings, end int) map[string]any {
		r := record{}
		r.setDurations(t, at(end))
		var out map[string]any
		b, _ := json.Marshal(r)
		_ = json.Unmarshal(b, &out)
		for _, key := range []string{"route", "status", "age_ms"} {
			delete(out, key)
		}
		return out
	}

	newConnection := &observer.Timings{
		AttemptBegin:     at(0),
		SerializeBegin:   at(0),
		SerializeEnd:     at(20),
		RequestReady:     at(50),
		ConnectBegin:     at(150),
		ConnectEnd:       at(400),
		SendBegin:        at(410),
		SendEnd:          at(500),
		ResponseHeaders:  at(900),
		BodyEnd:          at(950),
		DeserializeBegin: at(960),
		DeserializeEnd:   at(990),
		Extra:            map[string]time.Time{"compress_begin": at(60), "compress_end": at(100)},
	}
	want := map[string]any{
		"total_us": 1000.0, "serialize_us": 20.0, "compress_us": 40.0, "pool_us": 50.0, "connect_us": 250.0,
		"send_us": 90.0, "wait_us": 400.0, "body_read_us": 50.0, "deserialize_us": 30.0,
	}
	if got := durations(newConnection, 1000); !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v, got %v", want, got)
	}

	reused := &observer.Timings{AttemptBegin: at(0), RequestReady: at(100), SendBegin: at(130), SendEnd: at(160), ResponseHeaders: at(300)}
	got := durations(reused, 400)
	if _, ok := got["connect_us"]; ok || got["pool_us"] != 30.0 {
		t.Errorf("expected a pool wait of 30us without connecting, got %v", got)
	}

	// Sending can finish after the response headers arrive.
	reused.SendEnd = at(350)
	if got := durations(reused, 400); got["wait_us"] != 0.0 {
		t.Errorf("expected no wait, got %v", got["wait_us"])
	}

	// Equal compress marks mean the body wasn't compressed.
	reused.Extra = map[string]time.Time{"compress_begin": at(110), "compress_end": at(110)}
	if got := durations(reused, 400); got["compress_us"] != nil || got["pool_us"] != 20.0 {
		t.Errorf("expected no compression and a pool wait of 20us, got %v", got)
	}

	// Durations are floored to whole microseconds.
	partial := &observer.Timings{AttemptBegin: at(0), DeserializeBegin: at(0), DeserializeEnd: at(0).Add(999 * time.Nanosecond)}
	if got := durations(partial, 1); !reflect.DeepEqual(got, map[string]any{"total_us": 1.0, "deserialize_us": 0.0}) {
		t.Errorf("expected whole microseconds, got %v", got)
	}

	// Untraced requests, e.g. through a custom HTTP client.
	if got := durations(&observer.Timings{AttemptBegin: at(0)}, 400); !reflect.DeepEqual(got, map[string]any{"total_us": 400.0}) {
		t.Errorf("expected only a total, got %v", got)
	}
}

type fakePerformance struct{ raw string }

func (p fakePerformance) RawJSON() string { return p.raw }

type fakeResponse struct {
	Performance fakePerformance
}

func TestServerTotalMs(t *testing.T) {
	res := &fakeResponse{Performance: fakePerformance{raw: `{"query_execution_ms":5,"server_total_ms":7}`}}
	if ms, ok := serverTotalMs(&res); !ok || ms != 7 {
		t.Errorf("expected 7, got %v %t", ms, ok)
	}

	var nilRes *fakeResponse
	body := []byte("{}")
	for _, parsed := range []any{
		nil,
		&nilRes,
		&body,
		&map[string]any{"performance": map[string]any{"server_total_ms": 7}},
		&fakeResponse{},
		&fakeResponse{Performance: fakePerformance{raw: `{"server_total_ms":null}`}},
		&struct{ Performance int }{Performance: 7},
	} {
		if ms, ok := serverTotalMs(parsed); ok {
			t.Errorf("expected no server total for %#v, got %v", parsed, ms)
		}
	}
}

func TestNewRecord(t *testing.T) {
	reused := true
	start := time.Now().Add(-time.Second)
	res := &fakeResponse{Performance: fakePerformance{raw: `{"server_total_ms":7.4}`}}
	rec := newRecord(observer.RequestInfo{Method: "post", URL: queryURL}, &observer.RequestResult{
		Retries:  2,
		Response: &http.Response{StatusCode: 503},
		Parsed:   &res,
		Timings: &observer.Timings{
			AttemptBegin:     start,
			DeserializeBegin: start.Add(time.Millisecond),
			DeserializeEnd:   start.Add(1234567 * time.Nanosecond),
			Reused:           &reused,
		},
	}, 1)

	if rec.Route != queryRoute || rec.Status != 503 || rec.Retries != 2 || rec.Error != "" || rec.SampleRate != 0 {
		t.Errorf("unexpected record %+v", rec)
	}
	// The end is when deserializing finished, not now.
	if *rec.TotalUs != 1234 || *rec.DeserializeUs != 234 || *rec.ServerTotalUs != 7400 || !*rec.Reused {
		t.Errorf("unexpected record %+v", rec)
	}
}

func TestErrorName(t *testing.T) {
	for kind, want := range map[observer.ErrorKind]string{
		"":                               "",
		observer.ErrorKindCanceled:       "canceled",
		observer.ErrorKindTimeout:        "timeout",
		observer.ErrorKindPoolTimeout:    "timeout",
		observer.ErrorKindConnectTimeout: "timeout",
		observer.ErrorKindReadTimeout:    "timeout",
		observer.ErrorKindWriteTimeout:   "timeout",
		observer.ErrorKindConnectError:   "connect",
		observer.ErrorKindReadError:      "connect",
		observer.ErrorKindOther:          "other",
	} {
		if got := errorName(kind); got != want {
			t.Errorf("errorName(%q): expected %q, got %q", kind, want, got)
		}
	}
}

func TestScopeBatchLifecycle(t *testing.T) {
	no, yes := false, true
	for _, tc := range []struct {
		name          string
		reachedServer *bool
		resent        bool
		// Records buffered after the request, including its own.
		buffered int
	}{
		{"never reached the server", &no, true, 2},
		{"reached the server", &yes, false, 1},
		{"may have reached the server", nil, false, 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			b := newBuffer()
			addRecord(b, 200)
			scope := newScope(b, observer.RequestInfo{Method: "POST", URL: queryURL})
			if got := len(b.entries); got != 1 {
				t.Fatalf("expected records to stay buffered until the first attempt, got %d", got)
			}

			first, _ := http.NewRequest(http.MethodPost, queryURL, nil)
			scope.AttemptBegin(0, first)
			header := first.Header.Get(Header)
			if header == "" {
				t.Fatal("expected the first attempt to carry the batch")
			}
			scope.AttemptEnd(&observer.AttemptResult{Error: errors.New("failed"), ReachedServer: tc.reachedServer})

			second, _ := http.NewRequest(http.MethodPost, queryURL, nil)
			scope.AttemptBegin(1, second)
			want := ""
			if tc.resent {
				want = header
			}
			if got := second.Header.Get(Header); got != want {
				t.Errorf("expected header %q on the retry, got %q", want, got)
			}

			// A batch still held when the request fails goes back to the buffer.
			scope.AttemptEnd(&observer.AttemptResult{Error: errors.New("failed"), ReachedServer: tc.reachedServer})
			setRandom(t, 0)
			scope.Finish(&observer.RequestResult{Timings: &observer.Timings{}})
			if got := len(b.entries); got != tc.buffered {
				t.Errorf("expected %d buffered records, got %d", tc.buffered, got)
			}
		})
	}
}

func TestSharedBuffer(t *testing.T) {
	b := sharedBuffer("tpuf_shared", "https://a.example.com/")
	if sharedBuffer("tpuf_shared", "https://a.example.com/") != b {
		t.Error("expected the same key and base URL to share a buffer")
	}
	if sharedBuffer("tpuf_other", "https://a.example.com/") == b || sharedBuffer("tpuf_shared", "https://b.example.com/") == b {
		t.Error("expected another key or base URL to get its own buffer")
	}
}
