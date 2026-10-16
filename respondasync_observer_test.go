package turbopuffer_test

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/turbopuffer/turbopuffer-go/v2"
	"github.com/turbopuffer/turbopuffer-go/v2/internal/observer"
	"github.com/turbopuffer/turbopuffer-go/v2/internal/requestconfig"
	"github.com/turbopuffer/turbopuffer-go/v2/internal/respondasync"
	"github.com/turbopuffer/turbopuffer-go/v2/option"
)

type finishedRequest struct {
	info   observer.RequestInfo
	result *observer.RequestResult
}

// finishRecorder records every request that finishes, in order.
type finishRecorder struct {
	mu       sync.Mutex
	finished []finishedRequest
}

func (r *finishRecorder) Start(info observer.RequestInfo) observer.RequestScope {
	return &finishScope{recorder: r, info: info}
}

// write returns the observed write request, which finishes after its polls.
func (r *finishRecorder) write(t *testing.T) (polls []finishedRequest, write finishedRequest) {
	t.Helper()
	if len(r.finished) == 0 || r.finished[len(r.finished)-1].info.Method != http.MethodPost {
		t.Fatalf("expected the write to finish last, got %+v", r.finished)
	}
	return r.finished[:len(r.finished)-1], r.finished[len(r.finished)-1]
}

type finishScope struct {
	recorder *finishRecorder
	info     observer.RequestInfo
}

func (*finishScope) AttemptBegin(int, *http.Request)    {}
func (*finishScope) AttemptEnd(*observer.AttemptResult) {}

func (s *finishScope) Finish(result *observer.RequestResult) {
	s.recorder.mu.Lock()
	defer s.recorder.mu.Unlock()
	s.recorder.finished = append(s.recorder.finished, finishedRequest{info: s.info, result: result})
}

func withObserver(o observer.RequestObserver) option.RequestOption {
	return requestconfig.RequestOptionFunc(func(c *requestconfig.RequestConfig) error {
		c.RequestObserver = o
		return nil
	})
}

// newObservedClient responds to the nth request with respond(n).
func newObservedClient(rec *finishRecorder, respond func(n int) *http.Response, opts ...option.RequestOption) turbopuffer.Client {
	calls := 0
	return newClient(&closureTransport{
		fn: func(req *http.Request) (*http.Response, error) {
			calls++
			return respond(calls), nil
		},
	}, append([]option.RequestOption{withObserver(rec)}, opts...)...)
}

// accepted responds to the first request with a 202, and to later ones with polls.
func accepted(polls ...string) func(n int) *http.Response {
	return func(n int) *http.Response {
		if n == 1 {
			return asyncAcceptedResponse(pollURL)
		}
		return makeResponse(http.StatusOK, polls[min(n-2, len(polls)-1)], nil)
	}
}

func TestRespondAsyncObservedPolls(t *testing.T) {
	fastPollInterval(t)
	rec := &finishRecorder{}
	client := newObservedClient(rec, accepted(
		`{"status":"running"}`,
		`{"status":"finished","result":{"success":`+writeOK+`}}`,
	))
	ns := client.Namespace("test")
	if _, err := ns.Write(context.Background(), writeParams()); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	polls, write := rec.write(t)
	if len(polls) != 2 {
		t.Fatalf("expected 2 polls to finish, got %d", len(polls))
	}
	for i, poll := range polls {
		if poll.info.Method != http.MethodGet || poll.info.URL != pollURL {
			t.Errorf("poll #%d: unexpected request info %+v", i, poll.info)
		}
		if poll.result.Error != nil || poll.result.Response.StatusCode != http.StatusOK || poll.result.Parsed == nil {
			t.Errorf("poll #%d: unexpected result %+v", i, poll.result)
		}
	}

	// The write ends with the polled result, and its body was read before polling began.
	if write.result.Error != nil || write.result.Response.StatusCode != http.StatusOK || write.result.Parsed == nil {
		t.Errorf("unexpected write result %+v", write.result)
	}
	timings := write.result.Timings
	if timings.BodyEnd.IsZero() || timings.BodyEnd.After(polls[0].result.Timings.AttemptBegin) {
		t.Error("expected the write's body to be read before the first poll")
	}
	if timings.DeserializeEnd.Before(polls[1].result.Timings.DeserializeEnd) {
		t.Error("expected the write to finish after its polls")
	}
}

func TestRespondAsyncObservedFailedOperation(t *testing.T) {
	fastPollInterval(t)
	rec := &finishRecorder{}
	client := newObservedClient(rec, accepted(
		`{"status":"finished","result":{"error":{"status_code":404,"detail":{"message":"namespace not found"}}}}`,
	))
	ns := client.Namespace("test")
	_, err := ns.Write(context.Background(), writeParams())
	var apierr *turbopuffer.Error
	if !errors.As(err, &apierr) {
		t.Fatalf("expected *turbopuffer.Error, got %T: %v", err, err)
	}

	_, write := rec.write(t)
	if write.result.Error != err || write.result.ErrorKind != "" || write.result.Response.StatusCode != http.StatusNotFound {
		t.Errorf("unexpected write result %+v", write.result)
	}
}

func TestRespondAsyncObservedPollTimeout(t *testing.T) {
	prev := respondasync.PollInterval
	respondasync.PollInterval = 5 * time.Millisecond
	t.Cleanup(func() { respondasync.PollInterval = prev })

	rec := &finishRecorder{}
	client := newObservedClient(rec, accepted(`{"status":"running"}`), option.WithRequestTimeout(50*time.Millisecond))
	ns := client.Namespace("test")
	_, err := ns.Write(context.Background(), writeParams())
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected context.DeadlineExceeded, got %v", err)
	}

	_, write := rec.write(t)
	if write.result.ErrorKind != observer.ErrorKindTimeout || write.result.Response.StatusCode != http.StatusAccepted {
		t.Errorf("unexpected write result %+v", write.result)
	}
}
