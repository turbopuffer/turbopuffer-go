package requestconfig

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

// trackingBody wraps an io.Reader and records whether the body was fully
// drained and whether Close was called.
type trackingBody struct {
	io.Reader
	closed  bool
	drained bool
}

func newTrackingBody(data string) *trackingBody {
	return &trackingBody{Reader: strings.NewReader(data)}
}

func (b *trackingBody) Read(p []byte) (int, error) {
	n, err := b.Reader.Read(p)
	if err == io.EOF {
		b.drained = true
	}
	return n, err
}

func (b *trackingBody) Close() error {
	b.closed = true
	return nil
}

// roundTripFunc adapts a function to the http.RoundTripper interface.
type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestExecute_ResponseBody(t *testing.T) {
	newTestRequestConfig := func(responseBodyInto any) (*trackingBody, *RequestConfig) {
		body := newTrackingBody(`{"status":"ok"}`)
		return body, &RequestConfig{
			MaxRetries: 0,
			Request: &http.Request{
				Method: http.MethodGet,
				URL:    &url.URL{Path: "/test"},
				Header: http.Header{"Content-Type": []string{"application/json"}},
			},
			BaseURL: &url.URL{Scheme: "http", Host: "localhost"},
			HTTPClient: &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Header:     http.Header{"Content-Type": []string{"application/json"}},
					Body:       body,
				}, nil
			})},
			ResponseBodyInto: responseBodyInto,
		}
	}

	// If the client doesn't care to the read the response body, Execute should
	// drain and close it.
	t.Run("nil", func(t *testing.T) {
		body, cfg := newTestRequestConfig(nil)
		if err := cfg.Execute(); err != nil {
			t.Fatalf("Execute() returned error: %v", err)
		}
		if !body.drained {
			t.Error("expected response body to be drained, but it was not")
		}
		if !body.closed {
			t.Error("expected response body to be closed, but it was not")
		}
	})

	// If the client wants to read the response body deserialized into a Go
	// type, Execute should read and close the http.Response.Body.
	t.Run("map", func(t *testing.T) {
		var dst map[string]any
		body, cfg := newTestRequestConfig(&dst)

		if err := cfg.Execute(); err != nil {
			t.Fatalf("Execute() returned error: %v", err)
		}
		if !body.closed {
			t.Error("expected response body to be closed, but it was not")
		}
	})

	// If the client wants to read the response body into a *http.Response,
	// Execute should NOT drain or close the http.Response.Body, and the caller
	// takes responsibility for closing it.
	t.Run("**http.Response", func(t *testing.T) {
		var dst *http.Response
		body, cfg := newTestRequestConfig(&dst)
		if err := cfg.Execute(); err != nil {
			t.Fatalf("Execute() returned error: %v", err)
		}
		if body.closed {
			t.Error("expected response body NOT to be closed when caller takes ownership via **http.Response")
		}
		if dst == nil {
			t.Fatal("expected dst to be set to the *http.Response")
		}

		got, err := io.ReadAll(dst.Body)
		if err != nil {
			t.Fatalf("reading body from caller-owned response: %v", err)
		}
		dst.Body.Close()
		if !bytes.Equal(got, []byte(`{"status":"ok"}`)) {
			t.Errorf("unexpected body contents: %q", got)
		}
	})
}
