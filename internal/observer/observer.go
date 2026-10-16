// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package observer

import (
	"context"
	"net/http"
	"time"
)

// RequestObserver observes the requests a client makes, such as to record their timings. Custom
// code installs one by setting RequestConfig.RequestObserver from a request option.
//
// Observers are called synchronously at fixed points of each request, so they must be cheap. A
// panic in an observer is logged and recovered, so it can never break a request. An observer is
// shared by concurrent requests, so per request state belongs on the RequestScope it returns.
type RequestObserver interface {
	// Start is called once per request, before its first attempt. It returns a scope to observe
	// the request, or nil to leave it unobserved.
	Start(info RequestInfo) RequestScope
}

// RequestObserverFunc adapts a function to a RequestObserver.
type RequestObserverFunc func(info RequestInfo) RequestScope

func (f RequestObserverFunc) Start(info RequestInfo) RequestScope { return f(info) }

// RequestScope observes one request across its attempts.
type RequestScope interface {
	// AttemptBegin is called before each attempt is sent. It may change the request's headers.
	AttemptBegin(attempt int, req *http.Request)

	// AttemptEnd is called after each attempt, before its response body is closed.
	AttemptEnd(result *AttemptResult)

	// Finish is called at most once, after the response body is decoded or when the request fails
	// for good. It isn't called for a response whose body the caller reads itself.
	Finish(result *RequestResult)
}

// RequestInfo describes a request as it starts.
type RequestInfo struct {
	Method string
	// URL is the URL the request was made with, usually a path relative to the base URL such as
	// "/v1/things/abc". It doesn't include query parameters added from request params.
	URL string
}

// AttemptResult is the outcome of one attempt of a request.
type AttemptResult struct {
	// Attempt is the attempt's 0-based index.
	Attempt  int
	Request  *http.Request
	Response *http.Response
	// Error is set when the attempt failed without a response.
	Error     error
	ErrorKind ErrorKind
	// ReachedServer is false when the attempt certainly never reached the server, true when a
	// response arrived, and nil when unknown.
	ReachedServer *bool
	Timings       *Timings
}

// RequestResult is the final outcome of a request.
type RequestResult struct {
	// Retries is the number of retries made, which is also the final attempt's index.
	Retries  int
	Response *http.Response
	// Error is the error the request returned. For an error response it's the API error, which
	// has no ErrorKind.
	Error     error
	ErrorKind ErrorKind
	// Parsed is the value the response body was decoded into, or nil if it wasn't decoded.
	Parsed any
	// Timings are the final attempt's.
	Timings *Timings
}

// ErrorKind classifies why an attempt or request failed.
type ErrorKind string

const (
	ErrorKindPoolTimeout    ErrorKind = "pool_timeout"
	ErrorKindConnectTimeout ErrorKind = "connect_timeout"
	ErrorKindConnectError   ErrorKind = "connect_error"
	ErrorKindWriteTimeout   ErrorKind = "write_timeout"
	ErrorKindReadTimeout    ErrorKind = "read_timeout"
	ErrorKindReadError      ErrorKind = "read_error"
	// ErrorKindTimeout is any other timeout.
	ErrorKindTimeout  ErrorKind = "timeout"
	ErrorKindCanceled ErrorKind = "canceled"
	ErrorKindOther    ErrorKind = "other"
)

// Timings mark the progress of one attempt with readings from time.Now, so durations between
// them use the monotonic clock. A mark stays zero when its phase didn't happen or couldn't be
// observed. Phases may overlap.
type Timings struct {
	// AttemptStart is when the attempt began, before the request was built.
	AttemptStart time.Time
	// SerializeStart and SerializeEnd bound encoding the request body, which happens once before
	// the first attempt.
	SerializeStart time.Time
	SerializeEnd   time.Time
	// RequestReady is when the request was built and handed to the HTTP client.
	RequestReady time.Time
	// ConnectStart and ConnectEnd bound opening a new connection: DNS, TCP and TLS. They're zero
	// when an existing connection was used.
	ConnectStart time.Time
	ConnectEnd   time.Time
	// SendStart and SendEnd bound writing the request headers and body.
	SendStart time.Time
	SendEnd   time.Time
	// ResponseHeaders is when the first byte of the response arrived.
	ResponseHeaders time.Time
	// BodyEnd is when the response body was fully read. Only on the final attempt.
	BodyEnd time.Time
	// DeserializeStart and DeserializeEnd bound decoding the response body. Only on the final
	// attempt.
	DeserializeStart time.Time
	DeserializeEnd   time.Time
	// Reused is whether the attempt used an existing connection, or nil when unknown.
	Reused *bool
	// Extra holds marks added by custom code, such as middleware that compresses request bodies.
	Extra map[string]time.Time
}

// Mark records the current time in Extra under name. Call it before the attempt completes, from
// the goroutine making the request.
func (t *Timings) Mark(name string) {
	if t.Extra == nil {
		t.Extra = make(map[string]time.Time)
	}
	t.Extra[name] = time.Now()
}

type timingsKey struct{}

// WithTimings returns a copy of ctx that carries an attempt's timings.
func WithTimings(ctx context.Context, timings *Timings) context.Context {
	return context.WithValue(ctx, timingsKey{}, timings)
}

// ContextTimings returns the timings of the attempt that ctx belongs to, or nil if it isn't
// observed. Middleware and custom transports can use it to add marks.
func ContextTimings(ctx context.Context) *Timings {
	timings, _ := ctx.Value(timingsKey{}).(*Timings)
	return timings
}
