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

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := turbopuffer.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("tpuf_A1..."),
		option.WithRegion("gcp-us-central1"),
	)
	ns := client.Namespace("products")
	response, err := ns.Write(context.TODO(), turbopuffer.NamespaceWriteParams{
		DistanceMetric: turbopuffer.DistanceMetricCosineDistance,
		UpsertRows: []turbopuffer.RowParam{{
			ID: turbopuffer.IDParam{
				String: turbopuffer.String("2108ed60-6851-49a0-9016-8325434f3845"),
			},
			Vector: turbopuffer.VectorParam{
				FloatArray: []float64{0.1, 0.2},
			},
		}},
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.RowsAffected)
}
