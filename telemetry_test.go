package turbopuffer_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"slices"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/turbopuffer/turbopuffer-go/v2"
	"github.com/turbopuffer/turbopuffer-go/v2/internal/respondasync"
	"github.com/turbopuffer/turbopuffer-go/v2/internal/telemetry"
	"github.com/turbopuffer/turbopuffer-go/v2/option"
)

const (
	queryRoute = "POST /v2/namespaces/{ns}/query"
	writeRoute = "POST /v2/namespaces/{ns}"
	pollRoute  = "GET /v1/namespaces/{ns}/operations/{token}"
)

var queryOK = `{
	"rows": [],
	"billing": {"billable_logical_bytes_queried": 0, "billable_logical_bytes_returned": 0},
	"performance": {
		"approx_namespace_size": 1,
		"cache_hit_ratio": 1.0,
		"cache_temperature": "hot",
		"exhaustive_search_count": 0,
		"query_execution_ms": 5,
		"server_total_ms": 7
	}
}`

type telemetryRecord = map[string]any

var requestIDPattern = regexp.MustCompile(`^tpreq_[2-7a-z]{29}$`)

// telemetryTransport responds with respond and keeps the telemetry header and request ID of each
// request it sees, or "" for a request without one.
type telemetryTransport struct {
	mu      sync.Mutex
	headers []string
	ids     []string
	respond func(n int, req *http.Request) (*http.Response, error)
}

func (tr *telemetryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	tr.mu.Lock()
	n := len(tr.headers)
	tr.headers = append(tr.headers, req.Header.Get(telemetry.Header))
	tr.ids = append(tr.ids, req.Header.Get(telemetry.RequestIDHeader))
	tr.mu.Unlock()
	return tr.respond(n, req)
}

// requestIDs returns the request ID each request was sent with, or "" for a request without one.
func (tr *telemetryTransport) requestIDs() []string {
	tr.mu.Lock()
	defer tr.mu.Unlock()
	return slices.Clone(tr.ids)
}

// records returns the records the nth request carried, or nil if it carried none.
func (tr *telemetryTransport) records(t *testing.T, n int) []telemetryRecord {
	t.Helper()
	tr.mu.Lock()
	defer tr.mu.Unlock()
	if n >= len(tr.headers) {
		t.Fatalf("expected at least %d requests, got %d", n+1, len(tr.headers))
	}
	return decodeRecords(t, tr.headers[n])
}

// allRecords returns the records carried by every request, in order.
func (tr *telemetryTransport) allRecords(t *testing.T) []telemetryRecord {
	t.Helper()
	var records []telemetryRecord
	for n := range len(tr.headers) {
		records = append(records, tr.records(t, n)...)
	}
	return records
}

func decodeRecords(t *testing.T, header string) []telemetryRecord {
	t.Helper()
	if header == "" {
		return nil
	}
	var payload struct {
		Version int               `json:"version"`
		Records []telemetryRecord `json:"records"`
	}
	if err := json.Unmarshal([]byte(header), &payload); err != nil || payload.Version != 1 {
		t.Fatalf("invalid telemetry header %q: %v", header, err)
	}
	return payload.Records
}

func respondWith(status int, body string, headers http.Header) func(int, *http.Request) (*http.Response, error) {
	return func(int, *http.Request) (*http.Response, error) {
		return makeResponse(status, body, headers), nil
	}
}

var apiKeys atomic.Int64

// uniqueAPIKey gives a test a telemetry buffer of its own, including when it runs again.
func uniqueAPIKey(t *testing.T) string {
	return fmt.Sprintf("tpuf_%s_%d", t.Name(), apiKeys.Add(1))
}

func newTelemetryClient(t *testing.T, tr *telemetryTransport, opts ...option.RequestOption) turbopuffer.Client {
	t.Helper()
	opts = append([]option.RequestOption{option.WithAPIKey(uniqueAPIKey(t))}, opts...)
	return newClient(tr, opts...)
}

func namespaceOf(client turbopuffer.Client, name string) turbopuffer.Namespace {
	return client.Namespace(name)
}

func query(t *testing.T, ns turbopuffer.Namespace, opts ...option.RequestOption) error {
	t.Helper()
	_, err := ns.Query(context.Background(), turbopuffer.NamespaceQueryParams{TopK: turbopuffer.Int(1)}, opts...)
	return err
}

func mustQuery(t *testing.T, ns turbopuffer.Namespace, opts ...option.RequestOption) {
	t.Helper()
	if err := query(t, ns, opts...); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func checkFailures(t *testing.T, records []telemetryRecord, want ...[2]any) {
	t.Helper()
	if len(records) != len(want) {
		t.Fatalf("expected %d records, got %v", len(want), records)
	}
	for i, rec := range records {
		if rec["status"] != want[i][0] || rec["error"] != want[i][1] {
			t.Errorf("record #%d: expected status %v and error %v, got %v", i, want[i][0], want[i][1], rec)
		}
	}
}

func TestTelemetryRecordsRideOnNextRequest(t *testing.T) {
	tr := &telemetryTransport{respond: respondWith(http.StatusOK, queryOK, nil)}
	ns := namespaceOf(newTelemetryClient(t, tr, option.WithTelemetry(true)), "my-ns")

	mustQuery(t, ns)
	mustQuery(t, ns)

	if records := tr.records(t, 0); records != nil {
		t.Errorf("expected the first request to carry nothing, got %v", records)
	}
	records := tr.records(t, 1)
	if len(records) != 1 {
		t.Fatalf("expected one record, got %v", records)
	}
	rec := records[0]
	if rec["route"] != queryRoute || rec["status"] != 200.0 || rec["server_total_us"] != 7000.0 {
		t.Errorf("unexpected record %v", rec)
	}
	// Custom transports aren't traced, so network phases are only covered by TestTelemetryConnectionPhases.
	for _, key := range []string{"total_us", "serialize_us", "deserialize_us", "age_ms"} {
		if v, ok := rec[key].(float64); !ok || v < 0 || v != float64(int64(v)) {
			t.Errorf("expected %s to be a non-negative integer, got %v", key, rec[key])
		}
	}
	if rec["total_us"].(float64) < rec["serialize_us"].(float64)+rec["deserialize_us"].(float64) {
		t.Errorf("expected the total to cover serializing and deserializing, got %v", rec)
	}
	for _, key := range []string{"compress_us", "response_us", "retries", "error", "sample_rate", "reused"} {
		if _, ok := rec[key]; ok {
			t.Errorf("expected no %s, got %v", key, rec)
		}
	}
}

func TestTelemetryEnabling(t *testing.T) {
	for _, tc := range []struct {
		name    string
		env     string
		opts    []option.RequestOption
		enabled bool
	}{
		{"disabled by default", "", nil, false},
		{"enabled by env", "1", nil, true},
		{"invalid env", "sometimes", nil, false},
		{"option overrides env to disable", "1", []option.RequestOption{option.WithTelemetry(false)}, false},
		{"option overrides env to enable", "0", []option.RequestOption{option.WithTelemetry(true)}, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv(telemetry.EnvVar, tc.env)
			tr := &telemetryTransport{respond: respondWith(http.StatusOK, queryOK, nil)}
			client := newTelemetryClient(t, tr, tc.opts...)

			// Copies of the client keep the setting it was made with.
			os.Unsetenv(telemetry.EnvVar)
			mustQuery(t, client.Namespace("a"))
			mustQuery(t, client.Namespace("b"))

			if got := tr.records(t, 1) != nil; got != tc.enabled {
				t.Errorf("expected telemetry enabled %t, got %t", tc.enabled, got)
			}
			if got := tr.requestIDs()[0] != ""; got != tc.enabled {
				t.Errorf("expected a request ID sent %t, got %t", tc.enabled, got)
			}
		})
	}
}

func TestTelemetryDisabledPerRequest(t *testing.T) {
	tr := &telemetryTransport{respond: respondWith(http.StatusOK, queryOK, nil)}
	ns := namespaceOf(newTelemetryClient(t, tr, option.WithTelemetry(true)), "my-ns")

	mustQuery(t, ns)
	mustQuery(t, ns, option.WithTelemetry(false))
	mustQuery(t, ns)

	if tr.records(t, 1) != nil || tr.requestIDs()[1] != "" {
		t.Error("expected the request with telemetry disabled to carry nothing")
	}
	if records := tr.records(t, 2); len(records) != 1 {
		t.Errorf("expected only the first request's record, got %v", records)
	}
}

func TestTelemetrySharedByAPIKeyAndBaseURL(t *testing.T) {
	tr := &telemetryTransport{respond: respondWith(http.StatusOK, queryOK, nil)}
	key := option.WithAPIKey(uniqueAPIKey(t))

	mustQuery(t, namespaceOf(newClient(tr, key, option.WithTelemetry(true)), "a"))
	mustQuery(t, namespaceOf(newClient(tr, option.WithAPIKey(uniqueAPIKey(t)), option.WithTelemetry(true)), "a"))
	mustQuery(t, namespaceOf(newClient(tr, key, option.WithTelemetry(true), option.WithBaseURL("http://localhost:5001")), "a"))
	mustQuery(t, namespaceOf(newClient(tr, key, option.WithTelemetry(true)), "b"))

	if tr.records(t, 1) != nil || tr.records(t, 2) != nil {
		t.Error("expected another API key or base URL not to carry the first request's record")
	}
	if records := tr.records(t, 3); len(records) != 1 {
		t.Errorf("expected another client with the same API key and base URL to carry the first record, got %v", records)
	}
}

func TestTelemetryQueryParamsAreNotRecorded(t *testing.T) {
	tr := &telemetryTransport{respond: respondWith(http.StatusOK, `{"namespaces":[]}`, nil)}
	client := newTelemetryClient(t, tr, option.WithTelemetry(true))

	for range 2 {
		if _, err := client.Namespaces(context.Background(), turbopuffer.NamespacesParams{Prefix: turbopuffer.String("secret-prefix")}); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	if records := tr.records(t, 1); len(records) != 1 || records[0]["route"] != "GET /v1/namespaces" {
		t.Errorf("expected a route without query params, got %v", records)
	}
}

func TestTelemetryFailuresAreRecorded(t *testing.T) {
	tr := &telemetryTransport{respond: func(n int, req *http.Request) (*http.Response, error) {
		switch n {
		case 0:
			return nil, &net.OpError{Op: "read", Net: "tcp", Err: os.ErrDeadlineExceeded}
		case 1:
			return nil, &net.OpError{Op: "dial", Net: "tcp", Err: errors.New("connection refused")}
		case 2:
			return makeResponse(http.StatusBadRequest, `{"error":"bad request"}`, nil), nil
		default:
			return makeResponse(http.StatusOK, queryOK, nil), nil
		}
	}}
	ns := namespaceOf(newTelemetryClient(t, tr, option.WithTelemetry(true), option.WithMaxRetries(0)), "my-ns")

	for range 3 {
		if err := query(t, ns); err == nil {
			t.Fatal("expected an error")
		}
	}
	mustQuery(t, ns)

	if records := tr.records(t, 0); records != nil {
		t.Errorf("expected the first request to carry nothing, got %v", records)
	}
	checkFailures(t, tr.records(t, 1), [2]any{0.0, "timeout"})
	// The connect error never reached the server, so its batch is sent again with the new record.
	checkFailures(t, tr.records(t, 2), [2]any{0.0, "timeout"}, [2]any{0.0, "connect"})
	checkFailures(t, tr.records(t, 3), [2]any{400.0, nil})
}

func TestTelemetryParseFailureIsRecorded(t *testing.T) {
	tr := &telemetryTransport{respond: func(n int, req *http.Request) (*http.Response, error) {
		if n == 0 {
			return makeResponse(http.StatusOK, "not json", nil), nil
		}
		return makeResponse(http.StatusOK, queryOK, nil), nil
	}}
	ns := namespaceOf(newTelemetryClient(t, tr, option.WithTelemetry(true)), "my-ns")

	if err := query(t, ns); err == nil {
		t.Fatal("expected an error")
	}
	mustQuery(t, ns)

	checkFailures(t, tr.records(t, 1), [2]any{200.0, "other"})
}

func TestTelemetryCanceledRequestIsRecorded(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	tr := &telemetryTransport{respond: func(n int, req *http.Request) (*http.Response, error) {
		if n == 0 {
			cancel()
			<-req.Context().Done()
			return nil, req.Context().Err()
		}
		return makeResponse(http.StatusOK, queryOK, nil), nil
	}}
	ns := namespaceOf(newTelemetryClient(t, tr, option.WithTelemetry(true)), "my-ns")

	if _, err := ns.Query(ctx, turbopuffer.NamespaceQueryParams{TopK: turbopuffer.Int(1)}); !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled, got %v", err)
	}
	mustQuery(t, ns)

	checkFailures(t, tr.records(t, 1), [2]any{0.0, "canceled"})
}

func TestTelemetryBatchIsResentAfterConnectError(t *testing.T) {
	tr := &telemetryTransport{respond: func(n int, req *http.Request) (*http.Response, error) {
		if n == 1 {
			return nil, &net.OpError{Op: "dial", Net: "tcp", Err: errors.New("connection refused")}
		}
		return makeResponse(http.StatusOK, queryOK, nil), nil
	}}
	ns := namespaceOf(newTelemetryClient(t, tr, option.WithTelemetry(true)), "my-ns")

	mustQuery(t, ns)
	mustQuery(t, ns)

	tr.mu.Lock()
	defer tr.mu.Unlock()
	if tr.headers[1] == "" || tr.headers[2] != tr.headers[1] {
		t.Errorf("expected the retry to resend the batch, got %q then %q", tr.headers[1], tr.headers[2])
	}
}

func TestTelemetryBatchIsNotResentAfterServerError(t *testing.T) {
	tr := &telemetryTransport{respond: func(n int, req *http.Request) (*http.Response, error) {
		if n == 1 {
			return makeResponse(http.StatusServiceUnavailable, `{}`, http.Header{"Retry-After-Ms": {"0"}}), nil
		}
		return makeResponse(http.StatusOK, queryOK, nil), nil
	}}
	ns := namespaceOf(newTelemetryClient(t, tr, option.WithTelemetry(true)), "my-ns")

	for range 3 {
		mustQuery(t, ns)
	}

	if tr.records(t, 1) == nil || tr.records(t, 2) != nil {
		t.Error("expected only the attempt that reached the server to carry the batch")
	}
	// The retried query reports how many retries it took.
	if records := tr.records(t, 3); len(records) != 1 || records[0]["retries"] != 1.0 {
		t.Errorf("expected one record with 1 retry, got %v", records)
	}
}

func TestTelemetryServerSampleRateOfZeroStopsRecording(t *testing.T) {
	tr := &telemetryTransport{respond: respondWith(http.StatusOK, queryOK, http.Header{
		http.CanonicalHeaderKey(telemetry.SampleRateHeader): {"0"},
	})}
	ns := namespaceOf(newTelemetryClient(t, tr, option.WithTelemetry(true)), "my-ns")

	for range 3 {
		mustQuery(t, ns)
	}

	if records := tr.allRecords(t); records != nil {
		t.Errorf("expected nothing to be recorded, got %v", records)
	}
	// Unrecorded requests have no request IDs to link.
	if ids := tr.requestIDs(); ids[1] != "" || ids[2] != "" {
		t.Errorf("expected no request IDs after the first request, got %q", ids)
	}
}

func TestTelemetryRequestIDsRideOnRequestsAndRecords(t *testing.T) {
	tr := &telemetryTransport{respond: respondWith(http.StatusOK, queryOK, nil)}
	ns := namespaceOf(newTelemetryClient(t, tr, option.WithTelemetry(true)), "my-ns")

	mustQuery(t, ns)
	mustQuery(t, ns)

	ids := tr.requestIDs()
	if !requestIDPattern.MatchString(ids[0]) || !requestIDPattern.MatchString(ids[1]) || ids[0] == ids[1] {
		t.Errorf("expected two different request IDs, got %q", ids)
	}
	if records := tr.records(t, 1); len(records) != 1 || records[0]["request_id"] != ids[0] {
		t.Errorf("expected the record to repeat request ID %s, got %v", ids[0], records)
	}
}

func TestTelemetryRetriesKeepRequestIDs(t *testing.T) {
	tr := &telemetryTransport{respond: func(n int, req *http.Request) (*http.Response, error) {
		if n == 0 {
			return makeResponse(http.StatusServiceUnavailable, `{}`, http.Header{"Retry-After-Ms": {"0"}}), nil
		}
		return makeResponse(http.StatusOK, queryOK, nil), nil
	}}
	ns := namespaceOf(newTelemetryClient(t, tr, option.WithTelemetry(true)), "my-ns")

	mustQuery(t, ns)
	mustQuery(t, ns)

	ids := tr.requestIDs()
	if ids[0] == "" || ids[1] != ids[0] || ids[2] == ids[0] {
		t.Errorf("expected the retry to keep the request's ID, got %q", ids)
	}
	if records := tr.records(t, 2); len(records) != 1 || records[0]["request_id"] != ids[0] || records[0]["retries"] != 1.0 {
		t.Errorf("expected a record with request ID %s and 1 retry, got %v", ids[0], records)
	}
}

func TestTelemetryKeepsRequestIDsSetByCallers(t *testing.T) {
	header := option.WithHeader("X-Request-Id", "my-request")
	for _, tc := range []struct {
		name                    string
		clientOpts, requestOpts []option.RequestOption
	}{
		{"default header", []option.RequestOption{header}, nil},
		{"request header", nil, []option.RequestOption{header}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tr := &telemetryTransport{respond: respondWith(http.StatusOK, queryOK, nil)}
			opts := append([]option.RequestOption{option.WithTelemetry(true)}, tc.clientOpts...)
			ns := namespaceOf(newTelemetryClient(t, tr, opts...), "my-ns")

			mustQuery(t, ns, tc.requestOpts...)
			mustQuery(t, ns)

			if id := tr.requestIDs()[0]; id != "my-request" {
				t.Errorf("expected the caller's request ID, got %q", id)
			}
			if records := tr.records(t, 1); len(records) != 1 || records[0]["request_id"] != nil {
				t.Errorf("expected a record without a request ID, got %v", records)
			}
		})
	}
}

func TestTelemetryConcurrentRequests(t *testing.T) {
	tr := &telemetryTransport{respond: respondWith(http.StatusOK, queryOK, nil)}
	ns := namespaceOf(newTelemetryClient(t, tr, option.WithTelemetry(true)), "my-ns")

	const workers, queries = 8, 25
	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range queries {
				if err := query(t, ns); err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}
		}()
	}
	wg.Wait()

	// Every record is sent at most once, and a header carries at most 20.
	sent := 0
	for n := range workers * queries {
		records := tr.records(t, n)
		if len(records) > 20 {
			t.Errorf("expected at most 20 records in a header, got %d", len(records))
		}
		sent += len(records)
	}
	if sent == 0 || sent >= workers*queries {
		t.Errorf("expected fewer records sent than requests made, got %d", sent)
	}
}

// respondAsync accepts writes, answers polls with polls in turn, and responds to queries.
func respondAsync(polls ...string) func(int, *http.Request) (*http.Response, error) {
	var mu sync.Mutex
	n := 0
	return func(_ int, req *http.Request) (*http.Response, error) {
		switch req.Method {
		case http.MethodGet:
			mu.Lock()
			defer mu.Unlock()
			n++
			return makeResponse(http.StatusOK, polls[min(n-1, len(polls)-1)], nil), nil
		case http.MethodPost:
			if req.URL.Path == "/v2/namespaces/test" {
				return asyncAcceptedResponse(pollURL), nil
			}
		}
		return makeResponse(http.StatusOK, queryOK, nil), nil
	}
}

// respondAsyncRecords writes, then returns every record sent, including those sent by a query
// made afterwards.
func respondAsyncRecords(t *testing.T, tr *telemetryTransport, opts ...option.RequestOption) ([]telemetryRecord, error) {
	t.Helper()
	ns := namespaceOf(newTelemetryClient(t, tr, append([]option.RequestOption{option.WithTelemetry(true)}, opts...)...), "test")
	_, err := ns.Write(context.Background(), writeParams())
	mustQuery(t, ns)
	records := tr.allRecords(t)
	if len(records) < 2 {
		t.Fatalf("expected polls and the write to be recorded, got %v", records)
	}
	return records, err
}

func TestTelemetryRespondAsync(t *testing.T) {
	fastPollInterval(t)
	tr := &telemetryTransport{respond: respondAsync(
		`{"status":"running"}`,
		`{"status":"finished","result":{"success":`+writeOK+`}}`,
	)}

	records, err := respondAsyncRecords(t, tr)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	polls, write := records[:len(records)-1], records[len(records)-1]
	if len(polls) != 2 {
		t.Fatalf("expected 2 polls, got %v", polls)
	}
	var pollTotal float64
	for i, poll := range polls {
		if poll["route"] != pollRoute || poll["status"] != 200.0 {
			t.Errorf("poll #%d: unexpected record %v", i, poll)
		}
		pollTotal += poll["total_us"].(float64)
	}
	// The write reports the final outcome, and its total spans the polling.
	if write["route"] != writeRoute || write["status"] != 200.0 || write["error"] != nil {
		t.Errorf("unexpected write record %v", write)
	}
	if write["total_us"].(float64) < pollTotal {
		t.Errorf("expected the write's total to span its polls, got %v", write)
	}
	// The write and each of its polls are sent with an ID of their own, which their records repeat.
	ids := tr.requestIDs()
	if ids[0] == ids[1] || ids[1] == ids[2] ||
		write["request_id"] != ids[0] || polls[0]["request_id"] != ids[1] || polls[1]["request_id"] != ids[2] {
		t.Errorf("expected records to repeat their own request IDs %q, got %v", ids, records)
	}
}

func TestTelemetryRespondAsyncFailedOperation(t *testing.T) {
	fastPollInterval(t)
	tr := &telemetryTransport{respond: respondAsync(
		`{"status":"finished","result":{"error":{"status_code":404,"detail":{"error":"gone"}}}}`,
	)}

	records, err := respondAsyncRecords(t, tr)
	var apierr *turbopuffer.Error
	if !errors.As(err, &apierr) {
		t.Fatalf("expected *turbopuffer.Error, got %T: %v", err, err)
	}

	write := records[len(records)-1]
	if write["route"] != writeRoute || write["status"] != 404.0 || write["error"] != nil {
		t.Errorf("unexpected write record %v", write)
	}
}

func TestTelemetryRespondAsyncPollTimeout(t *testing.T) {
	prev := respondasync.PollInterval
	respondasync.PollInterval = 5 * time.Millisecond
	t.Cleanup(func() { respondasync.PollInterval = prev })
	tr := &telemetryTransport{respond: respondAsync(`{"status":"running"}`)}

	records, err := respondAsyncRecords(t, tr, option.WithRequestTimeout(50*time.Millisecond))
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected context.DeadlineExceeded, got %v", err)
	}

	write := records[len(records)-1]
	if write["route"] != writeRoute || write["status"] != 202.0 || write["error"] != "timeout" {
		t.Errorf("unexpected write record %v", write)
	}
}

// telemetryServer is a real server, since custom transports never open connections to trace.
type telemetryServer struct {
	*httptest.Server
	mu      sync.Mutex
	headers []string
}

func newTelemetryServer(t *testing.T, respond func(w http.ResponseWriter, r *http.Request)) *telemetryServer {
	t.Helper()
	s := &telemetryServer{}
	s.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.Copy(io.Discard, r.Body)
		s.mu.Lock()
		s.headers = append(s.headers, r.Header.Get(telemetry.Header))
		s.mu.Unlock()
		respond(w, r)
	}))
	t.Cleanup(s.Close)
	return s
}

// client uses the default HTTP client, which is traced.
func (s *telemetryServer) client(t *testing.T) turbopuffer.Client {
	return turbopuffer.NewClient(
		option.WithBaseURL(s.URL),
		option.WithAPIKey(uniqueAPIKey(t)),
		option.WithTelemetry(true),
	)
}

func (s *telemetryServer) records(t *testing.T, n int) []telemetryRecord {
	t.Helper()
	s.mu.Lock()
	defer s.mu.Unlock()
	return decodeRecords(t, s.headers[n])
}

func writeJSON(w http.ResponseWriter, status int, body string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(body)))
	w.WriteHeader(status)
	_, _ = io.WriteString(w, body)
}

func checkConnectionPhases(t *testing.T, rec telemetryRecord, reused bool) {
	t.Helper()
	if rec["reused"] != reused || (rec["connect_us"] == nil) == !reused {
		t.Errorf("expected reused %t, and connecting only on a new connection, got %v", reused, rec)
	}
	for _, key := range []string{"pool_us", "send_us", "wait_us", "body_read_us"} {
		if v, ok := rec[key].(float64); !ok || v < 0 {
			t.Errorf("expected %s to be non-negative, got %v", key, rec)
		}
	}
	if rec["wait_us"].(float64)+rec["body_read_us"].(float64) > rec["total_us"].(float64) {
		t.Errorf("expected the total to cover waiting and reading the body, got %v", rec)
	}
}

func TestTelemetryConnectionPhases(t *testing.T) {
	srv := newTelemetryServer(t, func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, queryOK)
	})
	client := srv.client(t)
	ns := client.Namespace("my-ns")
	for range 3 {
		mustQuery(t, ns)
	}

	first, second := srv.records(t, 1), srv.records(t, 2)
	if len(first) != 1 || len(second) != 1 {
		t.Fatalf("expected one record per request, got %v and %v", first, second)
	}
	checkConnectionPhases(t, first[0], false)
	checkConnectionPhases(t, second[0], true)
}

func TestTelemetryRespondAsyncConnectionPhases(t *testing.T) {
	fastPollInterval(t)
	srv := newTelemetryServer(t, func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet:
			writeJSON(w, http.StatusOK, `{"status":"finished","result":{"success":`+writeOK+`}}`)
		case r.URL.Path == "/v2/namespaces/test":
			w.Header().Set("Preference-Applied", "respond-async")
			w.Header().Set("Location", "/v1/namespaces/test/operations/op-abc")
			writeJSON(w, http.StatusAccepted, `{"status":"accepted"}`)
		default:
			writeJSON(w, http.StatusOK, queryOK)
		}
	})
	client := srv.client(t)
	ns := client.Namespace("test")
	if _, err := ns.Write(context.Background(), writeParams()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	mustQuery(t, ns)

	records := srv.records(t, 2)
	if len(records) != 2 {
		t.Fatalf("expected the poll and the write to be recorded, got %v", records)
	}
	poll, write := records[0], records[1]
	if poll["route"] != pollRoute || write["route"] != writeRoute || write["status"] != 200.0 {
		t.Errorf("unexpected records %v", records)
	}
	// The write's network phases are from the 202 exchange, before polling reused its connection.
	checkConnectionPhases(t, write, false)
	checkConnectionPhases(t, poll, true)
	if write["total_us"].(float64) < poll["total_us"].(float64)+write["wait_us"].(float64) {
		t.Errorf("expected the write's total to span its poll, got %v", records)
	}
}
