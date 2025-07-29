// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turbopuffer_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/turbopuffer/turbopuffer-go"
	"github.com/turbopuffer/turbopuffer-go/internal"
	"github.com/turbopuffer/turbopuffer-go/option"
)

type closureTransport struct {
	fn func(req *http.Request) (*http.Response, error)
}

func (t *closureTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.fn(req)
}

func TestUserAgentHeader(t *testing.T) {
	var userAgent string
	client := turbopuffer.NewClient(
		option.WithAPIKey("tpuf_A1..."),
		option.WithRegion("testing"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					userAgent = req.Header.Get("User-Agent")
					return &http.Response{
						StatusCode: http.StatusOK,
					}, nil
				},
			},
		}),
	)
	client.Namespaces(context.Background(), turbopuffer.NamespacesParams{
		Prefix: turbopuffer.String("foo"),
	})
	if userAgent != fmt.Sprintf("Turbopuffer/Go %s", internal.PackageVersion) {
		t.Errorf("Expected User-Agent to be correct, but got: %#v", userAgent)
	}
}

func TestRetryAfter(t *testing.T) {
	retryCountHeaders := make([]string, 0)
	client := turbopuffer.NewClient(
		option.WithAPIKey("tpuf_A1..."),
		option.WithRegion("testing"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					retryCountHeaders = append(retryCountHeaders, req.Header.Get("X-Stainless-Retry-Count"))
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After"): []string{"0.1"},
						},
					}, nil
				},
			},
		}),
	)
	_, err := client.Namespaces(context.Background(), turbopuffer.NamespacesParams{
		Prefix: turbopuffer.String("foo"),
	})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}

	attempts := len(retryCountHeaders)
	if attempts != 5 {
		t.Errorf("Expected %d attempts, got %d", 5, attempts)
	}

	expectedRetryCountHeaders := []string{"0", "1", "2", "3", "4"}
	if !reflect.DeepEqual(retryCountHeaders, expectedRetryCountHeaders) {
		t.Errorf("Expected %v retry count headers, got %v", expectedRetryCountHeaders, retryCountHeaders)
	}
}

func TestDeleteRetryCountHeader(t *testing.T) {
	retryCountHeaders := make([]string, 0)
	client := turbopuffer.NewClient(
		option.WithAPIKey("tpuf_A1..."),
		option.WithRegion("testing"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					retryCountHeaders = append(retryCountHeaders, req.Header.Get("X-Stainless-Retry-Count"))
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After"): []string{"0.1"},
						},
					}, nil
				},
			},
		}),
		option.WithHeaderDel("X-Stainless-Retry-Count"),
	)
	_, err := client.Namespaces(context.Background(), turbopuffer.NamespacesParams{
		Prefix: turbopuffer.String("foo"),
	})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}

	expectedRetryCountHeaders := []string{"", "", "", "", ""}
	if !reflect.DeepEqual(retryCountHeaders, expectedRetryCountHeaders) {
		t.Errorf("Expected %v retry count headers, got %v", expectedRetryCountHeaders, retryCountHeaders)
	}
}

func TestOverwriteRetryCountHeader(t *testing.T) {
	retryCountHeaders := make([]string, 0)
	client := turbopuffer.NewClient(
		option.WithAPIKey("tpuf_A1..."),
		option.WithRegion("testing"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					retryCountHeaders = append(retryCountHeaders, req.Header.Get("X-Stainless-Retry-Count"))
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After"): []string{"0.1"},
						},
					}, nil
				},
			},
		}),
		option.WithHeader("X-Stainless-Retry-Count", "42"),
	)
	_, err := client.Namespaces(context.Background(), turbopuffer.NamespacesParams{
		Prefix: turbopuffer.String("foo"),
	})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}

	expectedRetryCountHeaders := []string{"42", "42", "42", "42", "42"}
	if !reflect.DeepEqual(retryCountHeaders, expectedRetryCountHeaders) {
		t.Errorf("Expected %v retry count headers, got %v", expectedRetryCountHeaders, retryCountHeaders)
	}
}

func TestRetryAfterMs(t *testing.T) {
	attempts := 0
	client := turbopuffer.NewClient(
		option.WithAPIKey("tpuf_A1..."),
		option.WithRegion("testing"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					attempts++
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After-Ms"): []string{"100"},
						},
					}, nil
				},
			},
		}),
	)
	_, err := client.Namespaces(context.Background(), turbopuffer.NamespacesParams{
		Prefix: turbopuffer.String("foo"),
	})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}
	if want := 5; attempts != want {
		t.Errorf("Expected %d attempts, got %d", want, attempts)
	}
}

func TestContextCancel(t *testing.T) {
	client := turbopuffer.NewClient(
		option.WithAPIKey("tpuf_A1..."),
		option.WithRegion("testing"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					<-req.Context().Done()
					return nil, req.Context().Err()
				},
			},
		}),
	)
	cancelCtx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := client.Namespaces(cancelCtx, turbopuffer.NamespacesParams{
		Prefix: turbopuffer.String("foo"),
	})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}
}

func TestContextCancelDelay(t *testing.T) {
	client := turbopuffer.NewClient(
		option.WithAPIKey("tpuf_A1..."),
		option.WithRegion("testing"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					<-req.Context().Done()
					return nil, req.Context().Err()
				},
			},
		}),
	)
	cancelCtx, cancel := context.WithTimeout(context.Background(), 2*time.Millisecond)
	defer cancel()
	_, err := client.Namespaces(cancelCtx, turbopuffer.NamespacesParams{
		Prefix: turbopuffer.String("foo"),
	})
	if err == nil {
		t.Error("expected there to be a cancel error")
	}
}

func TestContextDeadline(t *testing.T) {
	testTimeout := time.After(3 * time.Second)
	testDone := make(chan struct{})

	deadline := time.Now().Add(100 * time.Millisecond)
	deadlineCtx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go func() {
		client := turbopuffer.NewClient(
			option.WithAPIKey("tpuf_A1..."),
			option.WithRegion("testing"),
			option.WithHTTPClient(&http.Client{
				Transport: &closureTransport{
					fn: func(req *http.Request) (*http.Response, error) {
						<-req.Context().Done()
						return nil, req.Context().Err()
					},
				},
			}),
		)
		_, err := client.Namespaces(deadlineCtx, turbopuffer.NamespacesParams{
			Prefix: turbopuffer.String("foo"),
		})
		if err == nil {
			t.Error("expected there to be a deadline error")
		}
		close(testDone)
	}()

	select {
	case <-testTimeout:
		t.Fatal("client didn't finish in time")
	case <-testDone:
		if diff := time.Since(deadline); diff < -30*time.Millisecond || 30*time.Millisecond < diff {
			t.Fatalf("client did not return within 30ms of context deadline, got %s", diff)
		}
	}
}

func TestRegionSubstitutionWithDefaultURL(t *testing.T) {
	var host string
	client := turbopuffer.NewClient(
		option.WithAPIKey("tpuf_A1..."),
		option.WithRegion("my-cool-region"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					host = req.URL.Host
					return &http.Response{
						StatusCode: http.StatusOK,
					}, nil
				},
			},
		}),
	)
	client.Namespaces(context.Background(), turbopuffer.NamespacesParams{})
	if host != "my-cool-region.turbopuffer.com" {
		t.Errorf("expected host to be my-cool-region.turbopuffer.com, got %s", host)
	}
}

func TestRegionUnspecifiedWithDefaultURL(t *testing.T) {
	client := turbopuffer.NewClient(
		option.WithAPIKey("tpuf_A1..."),
	)
	_, err := client.Namespaces(context.Background(), turbopuffer.NamespacesParams{})
	if err == nil {
		t.Error("expected error when region is not provided with default URL")
	}
	expectedMsg := "region is required, but not set (baseURL has a REGION placeholder: https://REGION.turbopuffer.com/)"
	if err.Error() != expectedMsg {
		t.Errorf("expected error message: %s, got: %s", expectedMsg, err.Error())
	}
}

func TestRegionUnspecifiedWithCompleteOverriddenURL(t *testing.T) {
	var host string
	client := turbopuffer.NewClient(
		option.WithAPIKey("tpuf_A1..."),
		option.WithBaseURL("https://tpuf.example.com"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					host = req.URL.Host
					return &http.Response{
						StatusCode: http.StatusOK,
					}, nil
				},
			},
		}),
	)
	client.Namespaces(context.Background(), turbopuffer.NamespacesParams{})
	if host != "tpuf.example.com" {
		t.Errorf("expected host to be tpuf.example.com, got %s", host)
	}
}

func TestRegionSpecifiedWithIncompleteOverriddenURL(t *testing.T) {
	client := turbopuffer.NewClient(
		option.WithAPIKey("tpuf_A1..."),
		option.WithBaseURL("https://REGION.custom.turbopuffer.com"),
	)
	_, err := client.Namespaces(context.Background(), turbopuffer.NamespacesParams{})
	if err == nil {
		t.Error("expected error when region is provided but URL has no placeholder")
	}
	expectedMsg := "region is required, but not set (baseURL has a REGION placeholder: https://REGION.custom.turbopuffer.com)"
	if err.Error() != expectedMsg {
		t.Errorf("expected error message: %s, got: %s", expectedMsg, err.Error())
	}
}

func TestRegionSpecifiedWithCompleteOverriddenURL(t *testing.T) {
	client := turbopuffer.NewClient(
		option.WithAPIKey("tpuf_A1..."),
		option.WithBaseURL("https://tpuf.example.com"),
		option.WithRegion("gcp-us-central1"),
	)
	_, err := client.Namespaces(context.Background(), turbopuffer.NamespacesParams{})
	if err == nil {
		t.Error("expected error when region is provided but URL has no placeholder")
	}
	expectedMsg := "region is set, but would be ignored (baseURL does not contain REGION placeholder: https://tpuf.example.com)"
	if err.Error() != expectedMsg {
		t.Errorf("expected error message: %s, got: %s", expectedMsg, err.Error())
	}
}
