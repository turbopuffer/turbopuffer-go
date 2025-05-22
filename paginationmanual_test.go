// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turbopuffer_test

import (
	"context"
	"os"
	"testing"

	"github.com/turbopuffer/turbopuffer-go"
	"github.com/turbopuffer/turbopuffer-go/internal/testutil"
	"github.com/turbopuffer/turbopuffer-go/option"
)

func TestManualPagination(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := turbopuffer.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithRegion("My-Region"),
	)
	page, err := client.ListNamespaces(context.TODO(), turbopuffer.ListNamespacesParams{
		Prefix: turbopuffer.String("products"),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, client := range page.Namespaces {
		t.Logf("%+v\n", client.ID)
	}
	// Prism mock isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, client := range page.Namespaces {
			t.Logf("%+v\n", client.ID)
		}
	}
}
