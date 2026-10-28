package turbopuffer_test

import (
	"context"
	"net/http"
	"sync"
	"testing"

	"github.com/turbopuffer/turbopuffer-go/v2"
	"github.com/turbopuffer/turbopuffer-go/v2/option"
)

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

// TestNamespaceConcurrentRace verifies that concurrent calls to Client.Namespace
// each resolve to the correct namespace. This guards against a slice-aliasing
// bug where append(r.Options, ...) reuses the backing array's spare capacity,
// causing concurrent callers to overwrite each other's namespace option.
func TestNamespaceConcurrentRace(t *testing.T) {
	client := turbopuffer.NewClient(
		option.WithAPIKey("tpuf_A1..."),
		option.WithRegion("testing"),
	)

	namespaces := []string{
		"ns_alpha",
		"ns_bravo",
		"ns_charlie",
		"ns_delta",
		"ns_echo",
		"ns_foxtrot",
	}

	for i := range 500 {
		var wg sync.WaitGroup
		handles := make([]turbopuffer.Namespace, len(namespaces))

		for j, ns := range namespaces {
			wg.Add(1)
			go func() {
				defer wg.Done()
				handles[j] = client.Namespace(ns)
			}()
		}
		wg.Wait()

		for j, ns := range namespaces {
			if got := handles[j].ID(); got != ns {
				t.Fatalf("iteration %d: Namespace(%q).ID() = %q; concurrent append aliased the backing array", i, ns, got)
			}
		}
	}
}
