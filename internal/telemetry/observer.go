package telemetry

import (
	"math"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/tidwall/gjson"
	"github.com/turbopuffer/turbopuffer-go/v2/internal/observer"
	"github.com/turbopuffer/turbopuffer-go/v2/internal/requestconfig"
)

// Option installs an observer that records telemetry for requests, or removes it.
func Option(enabled bool) requestconfig.RequestOption {
	return requestconfig.RequestOptionFunc(func(cfg *requestconfig.RequestConfig) error {
		if enabled {
			cfg.RequestObserver = &requestObserver{cfg: cfg}
		} else if _, ok := cfg.RequestObserver.(*requestObserver); ok {
			cfg.RequestObserver = nil
		}
		return nil
	})
}

// EnvOptions returns the options that enable telemetry when TURBOPUFFER_TELEMETRY does. A client
// applies them before its own options, which can override them.
func EnvOptions() []requestconfig.RequestOption {
	if !enabledByEnv() {
		return nil
	}
	return []requestconfig.RequestOption{Option(true)}
}

// requestObserver records telemetry for the requests made with cfg. It reads cfg as each request
// starts, once all options are applied, so records stay with the API key and base URL they were
// made with.
type requestObserver struct {
	cfg *requestconfig.RequestConfig
}

func (o *requestObserver) Start(info observer.RequestInfo) observer.RequestScope {
	base := o.cfg.BaseURL
	if base == nil {
		base = o.cfg.DefaultBaseURL
	}
	baseURL := ""
	if base != nil {
		baseURL = base.String()
	}
	return newScope(sharedBuffer(o.cfg.APIKey, baseURL), info)
}

type requestScope struct {
	buffer *buffer
	info   observer.RequestInfo
	// Kept for later attempts until one of them may have reached the server.
	batch *batch
}

func newScope(b *buffer, info observer.RequestInfo) *requestScope {
	return &requestScope{buffer: b, info: info}
}

func (s *requestScope) AttemptBegin(attempt int, req *http.Request) {
	// Taken only once a request is about to be sent, so one that fails before then can't lose it.
	if attempt == 0 {
		s.batch = s.buffer.take()
	}
	if s.batch != nil {
		req.Header.Set(Header, s.batch.header)
	}
}

func (s *requestScope) AttemptEnd(result *observer.AttemptResult) {
	if result.ReachedServer == nil || *result.ReachedServer {
		s.batch = nil
	}
	if result.Response != nil {
		s.buffer.observeResponse(result.Response.Header)
	}
}

func (s *requestScope) Finish(result *observer.RequestResult) {
	// A batch that never reached the server goes back for a later request.
	if s.batch != nil {
		s.buffer.restore(s.batch)
		s.batch = nil
	}
	rate, ok := s.buffer.sample()
	if !ok {
		return
	}
	s.buffer.add(newRecord(s.info, result, rate))
}

func newRecord(info observer.RequestInfo, result *observer.RequestResult, rate float64) record {
	rec := record{
		Route:   routeTemplate(info.Method, info.URL),
		Error:   errorName(result.ErrorKind),
		Retries: result.Retries,
	}
	if result.Response != nil {
		rec.Status = result.Response.StatusCode
	}
	if rate < 1 {
		rec.SampleRate = rate
	}

	t := result.Timings
	end := t.DeserializeEnd
	if end.IsZero() {
		end = time.Now()
	}
	rec.setDurations(t, end)
	if t.Reused != nil {
		reused := *t.Reused
		rec.Reused = &reused
	}

	if ms, ok := serverTotalMs(result.Parsed); ok {
		us := int64(math.Round(ms * 1000))
		rec.ServerTotalUs = &us
	}
	return rec
}

// errorName is the error reported for a failed request, or empty.
func errorName(kind observer.ErrorKind) string {
	switch {
	case kind == "":
		return ""
	case kind == observer.ErrorKindCanceled:
		return "canceled"
	case kind == observer.ErrorKindOther:
		return "other"
	case strings.HasSuffix(string(kind), "timeout"):
		return "timeout"
	default:
		return "connect"
	}
}

// setDurations fills in the duration of each phase of an attempt ending at end. Phases that
// weren't observed are left out.
func (r *record) setDurations(t *observer.Timings, end time.Time) {
	if t.AttemptStart.IsZero() {
		return
	}
	r.TotalUs = micros(end.Sub(t.AttemptStart))

	if !t.SerializeStart.IsZero() && !t.SerializeEnd.IsZero() {
		r.SerializeUs = micros(t.SerializeEnd.Sub(t.SerializeStart))
	}

	// Set by middleware that compresses request bodies. Zero means the body wasn't compressed.
	compressStart, compressEnd := t.Extra["compress_start"], t.Extra["compress_end"]
	if !compressStart.IsZero() && !compressEnd.IsZero() {
		if d := compressEnd.Sub(compressStart); d != 0 {
			r.CompressUs = micros(d)
		}
	}

	// Waiting for a free connection, until connecting a new one or sending on a reused one.
	poolStart := t.RequestReady
	if !compressEnd.IsZero() {
		poolStart = compressEnd
	}
	connectionReady := t.SendStart
	if !t.ConnectStart.IsZero() {
		connectionReady = t.ConnectStart
	}
	if !poolStart.IsZero() && !connectionReady.IsZero() {
		r.PoolUs = micros(max(connectionReady.Sub(poolStart), 0))
	}

	if !t.ConnectStart.IsZero() && !t.ConnectEnd.IsZero() {
		r.ConnectUs = micros(t.ConnectEnd.Sub(t.ConnectStart))
	}

	if !t.SendStart.IsZero() && !t.SendEnd.IsZero() {
		r.SendUs = micros(t.SendEnd.Sub(t.SendStart))
	}

	// Covers the network round trip and the server. Sending can finish after the response arrives.
	if !t.SendEnd.IsZero() && !t.ResponseHeaders.IsZero() {
		r.WaitUs = micros(max(t.ResponseHeaders.Sub(t.SendEnd), 0))
	}

	if !t.BodyEnd.IsZero() && !t.ResponseHeaders.IsZero() {
		r.BodyReadUs = micros(t.BodyEnd.Sub(t.ResponseHeaders))
	}

	if !t.DeserializeStart.IsZero() && !t.DeserializeEnd.IsZero() {
		r.DeserializeUs = micros(t.DeserializeEnd.Sub(t.DeserializeStart))
	}
}

func micros(d time.Duration) *int64 {
	us := d.Microseconds()
	return &us
}

// serverTotalMs returns performance.server_total_ms of a decoded response. The field is looked up
// by name since response types share no interface for it.
func serverTotalMs(parsed any) (float64, bool) {
	v := reflect.ValueOf(parsed)
	for v.Kind() == reflect.Pointer || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return 0, false
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return 0, false
	}
	performance := v.FieldByName("Performance")
	if !performance.IsValid() || !performance.CanInterface() ||
		(performance.Kind() == reflect.Pointer && performance.IsNil()) {
		return 0, false
	}
	// The raw JSON tells a missing field from a zero one.
	raw, ok := performance.Interface().(interface{ RawJSON() string })
	if !ok {
		return 0, false
	}
	ms := gjson.Get(raw.RawJSON(), "server_total_ms")
	if ms.Type != gjson.Number {
		return 0, false
	}
	return ms.Float(), true
}
