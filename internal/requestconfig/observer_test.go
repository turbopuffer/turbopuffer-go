// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package requestconfig_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/http/httptrace"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/turbopuffer/turbopuffer-go/v2/internal/apierror"
	"github.com/turbopuffer/turbopuffer-go/v2/internal/observer"
	"github.com/turbopuffer/turbopuffer-go/v2/internal/requestconfig"
	"github.com/turbopuffer/turbopuffer-go/v2/option"
)

type recorder struct {
	events   []string
	infos    []observer.RequestInfo
	attempts []*observer.AttemptResult
	results  []*observer.RequestResult
}

func (r *recorder) Start(info observer.RequestInfo) observer.RequestScope {
	r.events = append(r.events, "start")
	r.infos = append(r.infos, info)
	return r
}

func (r *recorder) AttemptBegin(attempt int, req *http.Request) {
	r.events = append(r.events, fmt.Sprintf("before %d", attempt))
	req.Header.Set("X-Observed-Attempt", strconv.Itoa(attempt))
}

func (r *recorder) AttemptEnd(result *observer.AttemptResult) {
	r.events = append(r.events, fmt.Sprintf("after %d", result.Attempt))
	r.attempts = append(r.attempts, result)
}

func (r *recorder) Finish(result *observer.RequestResult) {
	r.events = append(r.events, "finish")
	r.results = append(r.results, result)
}

func (r *recorder) checkEvents(t *testing.T, want ...string) {
	t.Helper()
	if got := strings.Join(r.events, ", "); got != strings.Join(want, ", ") {
		t.Fatalf("expected events %q, got %q", want, r.events)
	}
}

// writeAttempt responds with the attempt the observer stamped on the request.
func writeAttempt(w http.ResponseWriter, r *http.Request) {
	_, _ = io.Copy(io.Discard, r.Body)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"attempt": r.Header.Get("X-Observed-Attempt")})
}

func withObserver(o observer.RequestObserver) option.RequestOption {
	return requestconfig.RequestOptionFunc(func(r *requestconfig.RequestConfig) error {
		r.RequestObserver = o
		return nil
	})
}

func executeObserved(srv *httptest.Server, o observer.RequestObserver, dst any, opts ...option.RequestOption) error {
	opts = append([]option.RequestOption{
		option.WithBaseURL(srv.URL),
		option.WithHTTPClient(srv.Client()),
		option.WithMaxRetries(0),
		withObserver(o),
	}, opts...)
	return requestconfig.ExecuteNewRequest(context.Background(), http.MethodPost, "v1/things", map[string]string{"name": "thing"}, dst, opts...)
}

type mark struct {
	name string
	at   time.Time
}

// checkOrdered fails unless every mark is set and none is before the one preceding it.
func checkOrdered(t *testing.T, marks ...mark) {
	t.Helper()
	for i, m := range marks {
		if m.at.IsZero() {
			t.Errorf("expected %s to be set", m.name)
		} else if i > 0 && m.at.Before(marks[i-1].at) {
			t.Errorf("expected %s not to be before %s", m.name, marks[i-1].name)
		}
	}
}

func TestRequestObserver(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(writeAttempt))
	defer srv.Close()

	rec := &recorder{}
	middlewareTraced := false
	middleware := func(req *http.Request, next option.MiddlewareNext) (*http.Response, error) {
		middlewareTraced = httptrace.ContextClientTrace(req.Context()) != nil
		if timings := observer.ContextTimings(req.Context()); timings != nil {
			timings.Mark("middleware")
		}
		return next(req)
	}
	var dst map[string]string
	if err := executeObserved(srv, rec, &dst, option.WithMiddleware(middleware)); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	rec.checkEvents(t, "start", "before 0", "after 0", "finish")
	if info := rec.infos[0]; info.Method != http.MethodPost || info.URL != "/v1/things" {
		t.Errorf("unexpected request info: %+v", info)
	}
	if dst["attempt"] != "0" {
		t.Errorf("expected the header set in AttemptBegin to be sent, got %q", dst["attempt"])
	}
	if middlewareTraced {
		t.Error("expected middleware not to see the attempt's trace")
	}

	attempt := rec.attempts[0]
	if attempt.Response == nil || attempt.Response.StatusCode != http.StatusOK || attempt.Error != nil || attempt.ErrorKind != "" {
		t.Errorf("unexpected attempt result: %+v", attempt)
	}
	if attempt.ReachedServer == nil || !*attempt.ReachedServer {
		t.Error("expected the attempt to have reached the server")
	}

	result := rec.results[0]
	if result.Retries != 0 || result.Error != nil || result.ErrorKind != "" || result.Response != attempt.Response {
		t.Errorf("unexpected request result: %+v", result)
	}
	if parsed, ok := result.Parsed.(*map[string]string); !ok || parsed != &dst {
		t.Errorf("expected the decoded destination, got %#v", result.Parsed)
	}
	if result.Timings != attempt.Timings {
		t.Error("expected the request result to have the final attempt's timings")
	}

	timings := result.Timings
	if timings.Reused == nil || *timings.Reused {
		t.Error("expected a new connection")
	}
	if _, ok := timings.Extra["middleware"]; !ok {
		t.Error("expected the mark added by middleware")
	}
	checkOrdered(t,
		mark{"AttemptBegin", timings.AttemptBegin},
		mark{"SerializeBegin", timings.SerializeBegin},
		mark{"SerializeEnd", timings.SerializeEnd},
		mark{"RequestReady", timings.RequestReady},
		mark{"ConnectBegin", timings.ConnectBegin},
		mark{"ConnectEnd", timings.ConnectEnd},
		mark{"SendBegin", timings.SendBegin},
		mark{"ResponseHeaders", timings.ResponseHeaders},
		mark{"BodyEnd", timings.BodyEnd},
		mark{"DeserializeBegin", timings.DeserializeBegin},
		mark{"DeserializeEnd", timings.DeserializeEnd},
	)
	// Sending can finish after the response arrives.
	checkOrdered(t, mark{"SendBegin", timings.SendBegin}, mark{"SendEnd", timings.SendEnd})
}

func TestRequestObserverReplacedResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(writeAttempt))
	defer srv.Close()

	// Like middleware that polls for the result of an accepted request, then returns that instead.
	var replaced *http.Response
	var drained time.Time
	middleware := func(req *http.Request, next option.MiddlewareNext) (*http.Response, error) {
		res, err := next(req)
		if err != nil {
			return res, err
		}
		_, _ = io.Copy(io.Discard, res.Body)
		_ = res.Body.Close()
		drained = time.Now()
		observer.ContextTimings(req.Context()).BodyEnd = drained
		time.Sleep(10 * time.Millisecond)
		replaced = &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": {"application/json"}},
			Body:       io.NopCloser(strings.NewReader(`{"attempt":"replaced"}`)),
		}
		return replaced, nil
	}
	rec := &recorder{}
	var dst map[string]string
	if err := executeObserved(srv, rec, &dst, option.WithMiddleware(middleware)); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	result := rec.results[0]
	if result.Response != replaced || dst["attempt"] != "replaced" {
		t.Errorf("expected the request to finish with the replaced response, got %+v", result)
	}
	timings := result.Timings
	if !timings.BodyEnd.Equal(drained) {
		t.Error("expected BodyEnd to be kept from when the original response was read")
	}
	checkOrdered(t,
		mark{"BodyEnd", timings.BodyEnd},
		mark{"DeserializeBegin", timings.DeserializeBegin},
		mark{"DeserializeEnd", timings.DeserializeEnd},
	)

	// Failing after the response arrived isn't a read error.
	failing := func(req *http.Request, next option.MiddlewareNext) (*http.Response, error) {
		res, err := next(req)
		if err != nil {
			return res, err
		}
		_ = res.Body.Close()
		return res, context.DeadlineExceeded
	}
	rec = &recorder{}
	if err := executeObserved(srv, rec, &dst, option.WithMiddleware(failing)); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected the middleware's error, got %v", err)
	}
	if kind := rec.results[0].ErrorKind; kind != observer.ErrorKindTimeout {
		t.Errorf("expected a timeout, got %q", kind)
	}
}

func TestRequestObserverRetries(t *testing.T) {
	var requests atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if requests.Add(1) == 1 {
			_, _ = io.Copy(io.Discard, r.Body)
			w.Header().Set("Retry-After-Ms", "1")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		writeAttempt(w, r)
	}))
	defer srv.Close()

	rec := &recorder{}
	var dst map[string]string
	if err := executeObserved(srv, rec, &dst, option.WithMaxRetries(1)); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	rec.checkEvents(t, "start", "before 0", "after 0", "before 1", "after 1", "finish")
	if dst["attempt"] != "1" {
		t.Errorf("expected the second attempt's header to be sent, got %q", dst["attempt"])
	}

	first, second := rec.attempts[0], rec.attempts[1]
	if first.Response == nil || first.Response.StatusCode != http.StatusInternalServerError || first.ReachedServer == nil || !*first.ReachedServer {
		t.Errorf("unexpected first attempt result: %+v", first)
	}
	if first.Timings.Reused == nil || *first.Timings.Reused || first.Timings.ConnectBegin.IsZero() {
		t.Error("expected the first attempt to open a connection")
	}
	if second.Timings.Reused == nil || !*second.Timings.Reused {
		t.Error("expected the second attempt to reuse the connection")
	}
	if !second.Timings.ConnectBegin.IsZero() || !second.Timings.ConnectEnd.IsZero() {
		t.Error("expected no connect timings on a reused connection")
	}
	if !second.Timings.SerializeBegin.IsZero() {
		t.Error("expected serialize timings only on the first attempt")
	}
	checkOrdered(t,
		mark{"first ResponseHeaders", first.Timings.ResponseHeaders},
		mark{"second AttemptBegin", second.Timings.AttemptBegin},
		mark{"second RequestReady", second.Timings.RequestReady},
		mark{"second SendBegin", second.Timings.SendBegin},
		mark{"second ResponseHeaders", second.Timings.ResponseHeaders},
	)

	result := rec.results[0]
	if result.Retries != 1 || result.Timings != second.Timings || result.Parsed == nil {
		t.Errorf("unexpected request result: %+v", result)
	}
}

func TestRequestObserverConnectError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(writeAttempt))
	srv.Close()

	rec := &recorder{}
	var dst map[string]string
	err := executeObserved(srv, rec, &dst)
	if err == nil {
		t.Fatal("expected an error")
	}

	rec.checkEvents(t, "start", "before 0", "after 0", "finish")
	attempt := rec.attempts[0]
	if attempt.Error == nil || attempt.ErrorKind != observer.ErrorKindConnectError {
		t.Errorf("expected a connect error, got %q: %v", attempt.ErrorKind, attempt.Error)
	}
	if attempt.ReachedServer == nil || *attempt.ReachedServer {
		t.Error("expected the attempt not to have reached the server")
	}
	if result := rec.results[0]; result.Error != err || result.ErrorKind != observer.ErrorKindConnectError || result.Response != nil {
		t.Errorf("unexpected request result: %+v", result)
	}
}

func TestRequestObserverTimeout(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.Copy(io.Discard, r.Body)
		<-r.Context().Done()
	}))
	defer srv.Close()

	rec := &recorder{}
	var dst map[string]string
	err := executeObserved(srv, rec, &dst, option.WithRequestTimeout(50*time.Millisecond))
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected a deadline error, got %v", err)
	}

	rec.checkEvents(t, "start", "before 0", "after 0", "finish")
	attempt := rec.attempts[0]
	if attempt.ErrorKind != observer.ErrorKindReadTimeout || attempt.ReachedServer != nil {
		t.Errorf("expected a read timeout with an unknown outcome, got %q, %v", attempt.ErrorKind, attempt.ReachedServer)
	}
	if result := rec.results[0]; result.ErrorKind != observer.ErrorKindReadTimeout {
		t.Errorf("expected a read timeout, got %q", result.ErrorKind)
	}
}

func TestRequestObserverErrorResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.Copy(io.Discard, r.Body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("{}"))
	}))
	defer srv.Close()

	rec := &recorder{}
	var dst map[string]string
	err := executeObserved(srv, rec, &dst)
	if err == nil {
		t.Fatal("expected an error")
	}

	rec.checkEvents(t, "start", "before 0", "after 0", "finish")
	if attempt := rec.attempts[0]; attempt.Error != nil || attempt.ErrorKind != "" {
		t.Errorf("expected an error response to be no attempt error, got %q: %v", attempt.ErrorKind, attempt.Error)
	}
	result := rec.results[0]
	if result.Error != err || result.Response == nil || result.Response.StatusCode != http.StatusNotFound || result.Parsed != nil {
		t.Errorf("unexpected request result: %+v", result)
	}
	var apiErr *apierror.Error
	if errors.As(err, &apiErr) && result.ErrorKind != "" {
		t.Errorf("expected an API error to have no kind, got %q", result.ErrorKind)
	}
	if result.Timings.BodyEnd.IsZero() || !result.Timings.DeserializeBegin.IsZero() {
		t.Error("expected the error body to be read but not decoded")
	}
}

type panickingObserver struct{ inStart bool }

func (o panickingObserver) Start(observer.RequestInfo) observer.RequestScope {
	if o.inStart {
		panic("start")
	}
	return o
}

func (panickingObserver) AttemptBegin(int, *http.Request)    { panic("before attempt") }
func (panickingObserver) AttemptEnd(*observer.AttemptResult) { panic("after attempt") }
func (panickingObserver) Finish(*observer.RequestResult)     { panic("finish") }

func TestRequestObserverPanics(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(writeAttempt))
	defer srv.Close()

	var logs bytes.Buffer
	defaultLogger := slog.Default()
	slog.SetDefault(slog.New(slog.NewTextHandler(&logs, nil)))
	defer slog.SetDefault(defaultLogger)

	for _, o := range []panickingObserver{{inStart: true}, {inStart: false}} {
		var dst map[string]string
		if err := executeObserved(srv, o, &dst); err != nil || dst == nil {
			t.Fatalf("expected the request to succeed despite the observer, got %v, %v", err, dst)
		}
	}
	if n := strings.Count(logs.String(), "request observer panicked"); n != 4 {
		t.Errorf("expected 4 panics to be logged, got %d", n)
	}
}

func TestRequestObserverUnobserved(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(writeAttempt))
	defer srv.Close()

	observed := false
	middleware := func(req *http.Request, next option.MiddlewareNext) (*http.Response, error) {
		observed = observer.ContextTimings(req.Context()) != nil
		return next(req)
	}

	starts := 0
	unobserving := observer.RequestObserverFunc(func(observer.RequestInfo) observer.RequestScope {
		starts++
		return nil
	})
	for _, o := range []observer.RequestObserver{nil, unobserving} {
		var dst map[string]string
		if err := executeObserved(srv, o, &dst, option.WithMiddleware(middleware)); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if observed {
			t.Error("expected an unobserved request not to carry timings")
		}
	}
	if starts != 1 {
		t.Errorf("expected Start to be called once, got %d", starts)
	}

	// Building a request config without executing it doesn't start an observation.
	rec := &recorder{}
	_, err := requestconfig.NewRequestConfig(context.Background(), http.MethodGet, "v1/things", nil, nil, withObserver(rec))
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	rec.checkEvents(t)
}
