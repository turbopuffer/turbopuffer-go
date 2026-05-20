package turbopuffer_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"maps"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/turbopuffer/turbopuffer-go/v2"
	"github.com/turbopuffer/turbopuffer-go/v2/internal/respondasync"
	"github.com/turbopuffer/turbopuffer-go/v2/option"
)

const (
	baseURL = "http://localhost:5000"
	pollURL = baseURL + "/v1/namespaces/test/operations/op-abc"
)

var writeOK = `{
	"billing": {"billable_logical_bytes_written": 0, "billable_logical_bytes_returned": 0},
	"message": "OK",
	"rows_affected": 1,
	"status": "OK"
}`

func fastPollInterval(t *testing.T) {
	t.Helper()
	prev := respondasync.PollInterval
	respondasync.PollInterval = 0
	t.Cleanup(func() { respondasync.PollInterval = prev })
}

func makeResponse(status int, body string, headers http.Header) *http.Response {
	h := http.Header{"Content-Type": {"application/json"}}
	maps.Copy(h, headers)
	return &http.Response{
		StatusCode: status,
		Header:     h,
		Body:       io.NopCloser(bytes.NewBufferString(body)),
	}
}

func asyncAcceptedResponse(location string) *http.Response {
	return makeResponse(http.StatusAccepted, "", http.Header{
		"Preference-Applied": {"respond-async"},
		"Location":           {location},
	})
}

func newClient(rt http.RoundTripper, opts ...option.RequestOption) turbopuffer.Client {
	base := []option.RequestOption{
		option.WithBaseURL(baseURL),
		option.WithAPIKey("tpuf_A1..."),
		option.WithHTTPClient(&http.Client{Transport: rt}),
	}
	return turbopuffer.NewClient(append(base, opts...)...)
}

func writeParams() turbopuffer.NamespaceWriteParams {
	return turbopuffer.NamespaceWriteParams{
		UpsertRows: []turbopuffer.RowParam{
			{"id": "2108ed60-6851-49a0-9016-8325434f3845", "vector": []float32{0.1, 0.2}},
		},
	}
}

func TestRespondAsyncPreferHeaderSent(t *testing.T) {
	var preferHeader string
	client := newClient(&closureTransport{
		fn: func(req *http.Request) (*http.Response, error) {
			preferHeader = req.Header.Get("Prefer")
			return makeResponse(http.StatusOK, writeOK, nil), nil
		},
	})
	ns := client.Namespace("test")
	if _, err := ns.Write(context.Background(), writeParams()); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if preferHeader != "respond-async" {
		t.Errorf("expected `prefer: respond-async`, got %q", preferHeader)
	}
}

func TestRespondAsyncSyncResponsePassthrough(t *testing.T) {
	calls := 0
	client := newClient(&closureTransport{
		fn: func(req *http.Request) (*http.Response, error) {
			calls++
			return makeResponse(http.StatusOK, writeOK, nil), nil
		},
	})
	ns := client.Namespace("test")
	res, err := ns.Write(context.Background(), writeParams())
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if calls != 1 {
		t.Errorf("expected 1 request, got %d", calls)
	}
	if res.RowsAffected != 1 {
		t.Errorf("expected rows_affected=1, got %d", res.RowsAffected)
	}
}

func TestRespondAsyncUnrelated202Passthrough(t *testing.T) {
	calls := 0
	client := newClient(&closureTransport{
		fn: func(req *http.Request) (*http.Response, error) {
			calls++
			return makeResponse(http.StatusAccepted, writeOK, nil), nil
		},
	})
	ns := client.Namespace("test")
	res, err := ns.Write(context.Background(), writeParams())
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if calls != 1 {
		t.Errorf("expected 1 request, got %d", calls)
	}
	if res.Status != "OK" {
		t.Errorf("expected status OK, got %q", res.Status)
	}
}

func TestRespondAsyncPolledToSuccess(t *testing.T) {
	fastPollInterval(t)
	responses := []*http.Response{
		asyncAcceptedResponse(pollURL),
		makeResponse(http.StatusOK, `{"status":"running"}`, nil),
		makeResponse(http.StatusOK, `{"status":"running"}`, nil),
		makeResponse(http.StatusOK, `{"status":"finished","result":{"success":`+writeOK+`}}`, nil),
	}
	var requests []*http.Request
	client := newClient(&closureTransport{
		fn: func(req *http.Request) (*http.Response, error) {
			if len(requests) >= len(responses) {
				t.Fatalf("unexpected request #%d: %s %s", len(requests)+1, req.Method, req.URL)
			}
			requests = append(requests, req)
			return responses[len(requests)-1], nil
		},
	})
	ns := client.Namespace("test")
	res, err := ns.Write(context.Background(), writeParams())
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if len(requests) != 4 {
		t.Errorf("expected 4 requests, got %d", len(requests))
	}
	if res.RowsAffected != 1 {
		t.Errorf("expected rows_affected=1, got %d", res.RowsAffected)
	}
	for i := 1; i < len(requests); i++ {
		req := requests[i]
		if got := req.URL.String(); got != pollURL {
			t.Errorf("poll #%d URL = %q, want %q", i, got, pollURL)
		}
		if req.Method != http.MethodGet {
			t.Errorf("poll #%d method = %s, want GET", i, req.Method)
		}
		if auth := req.Header.Get("Authorization"); !strings.HasPrefix(auth, "Bearer ") {
			t.Errorf("poll #%d missing Authorization header, got %q", i, auth)
		}
		if got := req.Header.Get("Prefer"); got != "" {
			t.Errorf("poll #%d should not send Prefer, got %q", i, got)
		}
	}
}

func TestRespondAsyncErrorResultThrowsStatusError(t *testing.T) {
	fastPollInterval(t)
	responses := []*http.Response{
		asyncAcceptedResponse(pollURL),
		makeResponse(http.StatusOK,
			`{"status":"finished","result":{"error":{"status_code":404,"detail":{"message":"namespace not found"}}}}`,
			nil),
	}
	calls := 0
	client := newClient(&closureTransport{
		fn: func(req *http.Request) (*http.Response, error) {
			if calls >= len(responses) {
				t.Fatalf("unexpected request #%d: %s %s", calls+1, req.Method, req.URL)
			}
			res := responses[calls]
			calls++
			return res, nil
		},
	})
	ns := client.Namespace("test")
	_, err := ns.Write(context.Background(), writeParams())
	var apierr *turbopuffer.Error
	if !errors.As(err, &apierr) {
		t.Fatalf("expected *turbopuffer.Error, got %T: %v", err, err)
	}
	if apierr.StatusCode != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", apierr.StatusCode)
	}
	if !strings.Contains(apierr.RawJSON(), "namespace not found") {
		t.Errorf("expected error detail in raw JSON, got %q", apierr.RawJSON())
	}
}

func TestRespondAsyncTransientPollErrorRetried(t *testing.T) {
	fastPollInterval(t)
	responses := []*http.Response{
		asyncAcceptedResponse(pollURL),
		makeResponse(http.StatusServiceUnavailable, "", http.Header{"Retry-After-Ms": {"0"}}),
		makeResponse(http.StatusServiceUnavailable, "", http.Header{"Retry-After-Ms": {"0"}}),
		makeResponse(http.StatusOK, `{"status":"finished","result":{"success":`+writeOK+`}}`, nil),
	}
	calls := 0
	client := newClient(&closureTransport{
		fn: func(req *http.Request) (*http.Response, error) {
			if calls >= len(responses) {
				t.Fatalf("unexpected request #%d: %s %s", calls+1, req.Method, req.URL)
			}
			res := responses[calls]
			calls++
			return res, nil
		},
	})
	ns := client.Namespace("test")
	res, err := ns.Write(context.Background(), writeParams())
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if calls != 4 {
		t.Errorf("expected 4 requests, got %d", calls)
	}
	if res.RowsAffected != 1 {
		t.Errorf("expected rows_affected=1, got %d", res.RowsAffected)
	}
}

func TestRespondAsyncPollTimeout(t *testing.T) {
	fastPollInterval(t)
	calls := 0
	client := newClient(&closureTransport{
		fn: func(req *http.Request) (*http.Response, error) {
			calls++
			if calls == 1 {
				return asyncAcceptedResponse(pollURL), nil
			}
			return makeResponse(http.StatusOK, `{"status":"running"}`, nil), nil
		},
	}, option.WithRequestTimeout(50*time.Millisecond))
	ns := client.Namespace("test")
	_, err := ns.Write(context.Background(), writeParams())
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected context.DeadlineExceeded, got %v", err)
	}
}

func TestRespondAsyncPersistentPollErrorThrows(t *testing.T) {
	fastPollInterval(t)
	responses := []*http.Response{
		asyncAcceptedResponse(pollURL),
		makeResponse(http.StatusInternalServerError, "", nil),
	}
	calls := 0
	client := newClient(&closureTransport{
		fn: func(req *http.Request) (*http.Response, error) {
			if calls >= len(responses) {
				t.Fatalf("unexpected request #%d: %s %s", calls+1, req.Method, req.URL)
			}
			res := responses[calls]
			calls++
			return res, nil
		},
	}, option.WithMaxRetries(0))
	ns := client.Namespace("test")
	_, err := ns.Write(context.Background(), writeParams())
	var apierr *turbopuffer.Error
	if !errors.As(err, &apierr) {
		t.Fatalf("expected *turbopuffer.Error, got %T: %v", err, err)
	}
	if apierr.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", apierr.StatusCode)
	}
	if calls != 2 {
		t.Errorf("expected 2 requests (no retries), got %d", calls)
	}
}
