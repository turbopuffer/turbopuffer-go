// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package requestconfig

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/turbopuffer/turbopuffer-go/v2/internal/apierror"
	"github.com/turbopuffer/turbopuffer-go/v2/internal/observer"
)

// requestPrep is what NewRequestConfig records for a request observer, which isn't known until
// options are applied.
type requestPrep struct {
	url            string
	start          time.Time
	serializeStart time.Time
	serializeEnd   time.Time
}

// observation drives a request's scope through its attempts. Its methods do nothing when it's
// nil, which it is when the request isn't observed.
type observation struct {
	cfg     *RequestConfig
	scope   observer.RequestScope
	attempt int
	current *attemptState
	req     *http.Request
	res     *http.Response
}

func startObservation(cfg *RequestConfig) *observation {
	if cfg.RequestObserver == nil {
		return nil
	}
	scope := startScope(cfg.RequestObserver, observer.RequestInfo{
		Method: cfg.Request.Method,
		URL:    observedURL(cfg),
	})
	if scope == nil {
		return nil
	}
	o := &observation{cfg: cfg, scope: scope, current: &attemptState{}}
	t := &o.current.timings
	if cfg.prep.start.IsZero() {
		t.AttemptStart = time.Now()
	} else {
		t.AttemptStart = cfg.prep.start
		t.SerializeStart = cfg.prep.serializeStart
		t.SerializeEnd = cfg.prep.serializeEnd
	}
	return o
}

func startScope(o observer.RequestObserver, info observer.RequestInfo) observer.RequestScope {
	defer recoverObserver()
	return o.Start(info)
}

// observedURL is the URL the request was made with, with a leading slash when it's relative.
func observedURL(cfg *RequestConfig) string {
	u := cfg.prep.url
	if u == "" {
		u = cfg.Request.URL.String()
	}
	if parsed, err := url.Parse(u); err == nil && !parsed.IsAbs() && !strings.HasPrefix(u, "/") {
		return "/" + u
	}
	return u
}

// startAttempt returns ctx carrying the attempt's timings, so custom code can add marks.
func (o *observation) startAttempt(ctx context.Context, attempt int) context.Context {
	if o == nil {
		return ctx
	}
	// The first attempt began when the request was built.
	if attempt > 0 {
		o.attempt = attempt
		o.current = &attemptState{}
		o.current.timings.AttemptStart = time.Now()
	}
	return observer.WithTimings(ctx, &o.current.timings)
}

// traced attaches the attempt's trace below any middleware, so requests that middleware makes
// itself aren't traced as part of the attempt.
func (o *observation) traced(next middlewareNext) middlewareNext {
	if o == nil {
		return next
	}
	return func(req *http.Request) (*http.Response, error) {
		ctx := httptrace.WithClientTrace(req.Context(), o.current.clientTrace())
		return next(req.WithContext(ctx))
	}
}

func (o *observation) attemptBegin(req *http.Request) {
	if o == nil {
		return
	}
	o.current.timings.RequestReady = time.Now()
	o.req = req
	defer recoverObserver()
	o.scope.AttemptBegin(o.attempt, req)
}

func (o *observation) attemptEnd(res *http.Response, err error) {
	if o == nil {
		return
	}
	a := o.current
	a.finalize()
	a.responded = res != nil
	a.failedAfterResponse = res != nil && err != nil
	o.res = res
	result := &observer.AttemptResult{
		Attempt:  o.attempt,
		Request:  o.req,
		Response: res,
		Error:    err,
		Timings:  &a.timings,
	}
	if err != nil {
		phase := errorPhase(err, a)
		result.ErrorKind = errorKind(err, phase)
		if res == nil && (phase == phasePool || phase == phaseConnect) {
			result.ReachedServer = boolPtr(false)
		}
	}
	if res != nil {
		result.ReachedServer = boolPtr(true)
	}
	defer recoverObserver()
	o.scope.AttemptEnd(result)
}

// bodyRead marks the response body as read, unless custom code that replaced the response
// already marked when it read the original.
func (o *observation) bodyRead() {
	if o != nil && o.current.timings.BodyEnd.IsZero() {
		o.current.timings.BodyEnd = time.Now()
	}
}

func (o *observation) decodeStart() {
	if o != nil {
		o.current.timings.DeserializeStart = time.Now()
	}
}

func (o *observation) finish(err error) {
	if o == nil {
		return
	}
	a := o.current
	a.finalize()
	t := &a.timings
	// A response whose body the caller reads itself doesn't finish.
	if err == nil && t.DeserializeStart.IsZero() {
		return
	}
	if !t.DeserializeStart.IsZero() {
		t.DeserializeEnd = time.Now()
	}
	result := &observer.RequestResult{
		Retries:  o.attempt,
		Response: o.res,
		Error:    err,
		Timings:  t,
	}
	var apiErr *apierror.Error
	switch {
	case err == nil:
		result.Parsed = o.cfg.ResponseBodyInto
	case errors.As(err, &apiErr):
		// An error response, which Response describes.
	case !t.DeserializeStart.IsZero():
		// The body couldn't be decoded.
		result.ErrorKind = observer.ErrorKindOther
	default:
		result.ErrorKind = errorKind(err, errorPhase(err, a))
	}
	defer recoverObserver()
	o.scope.Finish(result)
}

// recoverObserver logs a panic from an observer, which must never break a request.
func recoverObserver() {
	if r := recover(); r != nil {
		slog.Warn("request observer panicked", "panic", r, "stack", string(debug.Stack()))
	}
}

func boolPtr(b bool) *bool { return &b }

// attemptState records one attempt. Trace hooks can run on other goroutines and fire late, so they
// hold mu and stop once the attempt is done.
type attemptState struct {
	mu      sync.Mutex
	timings observer.Timings
	// traced is set once the transport asked for a connection, so trace facts are known.
	traced    bool
	dialing   bool
	connected bool
	responded bool
	// failedAfterResponse is set when middleware failed after the response arrived, such as while
	// polling for the result of an accepted request.
	failedAfterResponse bool
	done                bool
}

func (a *attemptState) clientTrace() *httptrace.ClientTrace {
	return &httptrace.ClientTrace{
		GetConn:              func(string) { a.getConn() },
		DNSStart:             func(httptrace.DNSStartInfo) { a.connectStart() },
		ConnectStart:         func(string, string) { a.connectStart() },
		TLSHandshakeStart:    func() { a.connectStart() },
		ConnectDone:          func(_, _ string, err error) { a.connectDone(err) },
		TLSHandshakeDone:     func(_ tls.ConnectionState, err error) { a.connectDone(err) },
		GotConn:              a.gotConn,
		WroteRequest:         a.wroteRequest,
		GotFirstResponseByte: a.gotFirstResponseByte,
	}
}

// update calls f with the timings unless the attempt is done.
func (a *attemptState) update(f func(t *observer.Timings)) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.done {
		f(&a.timings)
	}
}

func (a *attemptState) finalize() {
	a.mu.Lock()
	a.done = true
	a.mu.Unlock()
}

// getConn resets what an earlier connection recorded, such as before a redirect or after a
// reused connection turned out to be closed.
func (a *attemptState) getConn() {
	a.update(func(t *observer.Timings) {
		a.traced, a.dialing, a.connected = true, false, false
		t.ConnectStart, t.ConnectEnd = time.Time{}, time.Time{}
		t.SendStart, t.SendEnd, t.ResponseHeaders = time.Time{}, time.Time{}, time.Time{}
		t.Reused = nil
	})
}

func (a *attemptState) connectStart() {
	now := time.Now()
	a.update(func(t *observer.Timings) {
		// A dial can continue after the attempt got another connection.
		if a.connected {
			return
		}
		a.dialing = true
		if t.ConnectStart.IsZero() {
			t.ConnectStart = now
		}
	})
}

func (a *attemptState) connectDone(err error) {
	now := time.Now()
	a.update(func(t *observer.Timings) {
		if err == nil && !a.connected {
			t.ConnectEnd = now
		}
	})
}

func (a *attemptState) gotConn(info httptrace.GotConnInfo) {
	now := time.Now()
	a.update(func(t *observer.Timings) {
		a.connected = true
		existing := info.Reused || info.WasIdle
		t.Reused = &existing
		if existing {
			t.ConnectStart, t.ConnectEnd = time.Time{}, time.Time{}
		} else if !t.ConnectStart.IsZero() && t.ConnectEnd.IsZero() {
			t.ConnectEnd = now
		}
		t.SendStart = now
	})
}

func (a *attemptState) wroteRequest(info httptrace.WroteRequestInfo) {
	now := time.Now()
	a.update(func(t *observer.Timings) {
		if info.Err == nil {
			t.SendEnd = now
		}
	})
}

func (a *attemptState) gotFirstResponseByte() {
	now := time.Now()
	a.update(func(t *observer.Timings) {
		t.ResponseHeaders = now
	})
}

type attemptPhase int

const (
	phaseUnknown attemptPhase = iota
	phasePool
	phaseConnect
	phaseWrite
	phaseRead
)

// phase reports how far the attempt got. It must only be called once the attempt is done.
func (a *attemptState) phase() attemptPhase {
	switch {
	case a.failedAfterResponse:
		return phaseUnknown
	case a.responded || !a.timings.ResponseHeaders.IsZero():
		return phaseRead
	case !a.traced:
		return phaseUnknown
	case !a.connected && a.dialing:
		return phaseConnect
	case !a.connected:
		return phasePool
	case a.timings.SendEnd.IsZero():
		return phaseWrite
	default:
		return phaseRead
	}
}

// errorPhase is the phase err happened in, from its type if that tells, or else from the trace.
func errorPhase(err error, a *attemptState) attemptPhase {
	var opErr *net.OpError
	if isConnectError(err) {
		return phaseConnect
	}
	if errors.As(err, &opErr) {
		switch opErr.Op {
		case "read":
			return phaseRead
		case "write":
			return phaseWrite
		}
	}
	return a.phase()
}

func isConnectError(err error) bool {
	var (
		dnsErr       *net.DNSError
		opErr        *net.OpError
		verifyErr    *tls.CertificateVerificationError
		recordErr    tls.RecordHeaderError
		authorityErr x509.UnknownAuthorityError
		hostnameErr  x509.HostnameError
		invalidErr   x509.CertificateInvalidError
	)
	return errors.As(err, &dnsErr) ||
		(errors.As(err, &opErr) && (opErr.Op == "dial" || opErr.Op == "proxyconnect")) ||
		errors.As(err, &verifyErr) ||
		errors.As(err, &recordErr) ||
		errors.As(err, &authorityErr) ||
		errors.As(err, &hostnameErr) ||
		errors.As(err, &invalidErr)
}

func isTimeout(err error) bool {
	var netErr net.Error
	return errors.Is(err, context.DeadlineExceeded) || (errors.As(err, &netErr) && netErr.Timeout())
}

func errorKind(err error, phase attemptPhase) observer.ErrorKind {
	if errors.Is(err, context.Canceled) {
		return observer.ErrorKindCanceled
	}
	timeout := isTimeout(err)
	switch {
	case phase == phasePool && timeout:
		return observer.ErrorKindPoolTimeout
	case phase == phaseConnect && timeout:
		return observer.ErrorKindConnectTimeout
	case phase == phaseConnect:
		return observer.ErrorKindConnectError
	case phase == phaseWrite && timeout:
		return observer.ErrorKindWriteTimeout
	case phase == phaseRead && timeout:
		return observer.ErrorKindReadTimeout
	case phase == phaseRead:
		return observer.ErrorKindReadError
	case timeout:
		return observer.ErrorKindTimeout
	default:
		return observer.ErrorKindOther
	}
}
