// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turbopuffer_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/turbopuffer-go"
	"github.com/stainless-sdks/turbopuffer-go/internal/testutil"
	"github.com/stainless-sdks/turbopuffer-go/option"
)

func TestAutoPagination(t *testing.T) {
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
	)
	iter := client.Namespaces.ListAutoPaging(context.TODO(), turbopuffer.NamespaceListParams{
		Prefix: turbopuffer.String("products"),
	})
	// Prism mock isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		namespace := iter.Current()
		t.Logf("%+v\n", namespace.Namespace.ID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
