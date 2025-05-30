// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turbopuffer_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/turbopuffer/turbopuffer-go"
	"github.com/turbopuffer/turbopuffer-go/internal/testutil"
	"github.com/turbopuffer/turbopuffer-go/option"
)

func TestNamespaceDeleteAll(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Namespaces.DeleteAll(context.TODO(), turbopuffer.NamespaceDeleteAllParams{
		Namespace: turbopuffer.String("namespace"),
	})
	if err != nil {
		var apierr *turbopuffer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNamespaceHintCacheWarm(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Namespaces.HintCacheWarm(context.TODO(), turbopuffer.NamespaceHintCacheWarmParams{
		Namespace: turbopuffer.String("namespace"),
	})
	if err != nil {
		var apierr *turbopuffer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNamespaceQueryWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Namespaces.Query(context.TODO(), turbopuffer.NamespaceQueryParams{
		Namespace: turbopuffer.String("namespace"),
		AggregateBy: map[string]any{
			"foo": "bar",
		},
		Consistency: turbopuffer.NamespaceQueryParamsConsistency{
			Level: turbopuffer.NamespaceQueryParamsConsistencyLevelStrong,
		},
		DistanceMetric: turbopuffer.DistanceMetricCosineDistance,
		Filters:        map[string]interface{}{},
		IncludeAttributes: turbopuffer.IncludeAttributesParam{
			Bool: turbopuffer.Bool(true),
		},
		RankBy:         map[string]interface{}{},
		TopK:           turbopuffer.Int(0),
		VectorEncoding: turbopuffer.VectorEncodingFloat,
	})
	if err != nil {
		var apierr *turbopuffer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNamespaceRecallWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Namespaces.Recall(context.TODO(), turbopuffer.NamespaceRecallParams{
		Namespace: turbopuffer.String("namespace"),
		Filters:   map[string]interface{}{},
		Num:       turbopuffer.Int(0),
		Queries:   []float64{0},
		TopK:      turbopuffer.Int(0),
	})
	if err != nil {
		var apierr *turbopuffer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNamespaceSchema(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Namespaces.Schema(context.TODO(), turbopuffer.NamespaceSchemaParams{
		Namespace: turbopuffer.String("namespace"),
	})
	if err != nil {
		var apierr *turbopuffer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNamespaceUpdateSchemaWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Namespaces.UpdateSchema(context.TODO(), turbopuffer.NamespaceUpdateSchemaParams{
		Namespace: turbopuffer.String("namespace"),
		Schema: map[string]turbopuffer.AttributeSchemaParam{
			"foo": {
				String: turbopuffer.String("string"),
			},
		},
	})
	if err != nil {
		var apierr *turbopuffer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNamespaceWriteWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Namespaces.Write(context.TODO(), turbopuffer.NamespaceWriteParams{
		Namespace:         turbopuffer.String("namespace"),
		CopyFromNamespace: turbopuffer.String("copy_from_namespace"),
		DeleteByFilter:    map[string]interface{}{},
		Deletes: []turbopuffer.IDParam{{
			String: turbopuffer.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		}},
		DistanceMetric: turbopuffer.DistanceMetricCosineDistance,
		Encryption: turbopuffer.NamespaceWriteParamsEncryption{
			Cmek: turbopuffer.NamespaceWriteParamsEncryptionCmek{
				KeyName: "key_name",
			},
		},
		PatchColumns: turbopuffer.DocumentColumnsParam{
			ID: []turbopuffer.IDParam{{
				String: turbopuffer.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}},
			Vector: turbopuffer.DocumentColumnsVectorParam{
				VectorArray: []turbopuffer.VectorParam{{
					FloatArray: []float64{0},
				}},
			},
		},
		PatchRows: []turbopuffer.DocumentRowParam{{
			ID: turbopuffer.IDParam{
				String: turbopuffer.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			},
			Vector: turbopuffer.VectorParam{
				FloatArray: []float64{0},
			},
		}},
		Schema: map[string]turbopuffer.AttributeSchemaParam{
			"foo": {
				String: turbopuffer.String("string"),
			},
		},
		UpsertColumns: turbopuffer.DocumentColumnsParam{
			ID: []turbopuffer.IDParam{{
				String: turbopuffer.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}},
			Vector: turbopuffer.DocumentColumnsVectorParam{
				VectorArray: []turbopuffer.VectorParam{{
					FloatArray: []float64{0},
				}},
			},
		},
		UpsertRows: []turbopuffer.DocumentRowParam{{
			ID: turbopuffer.IDParam{
				String: turbopuffer.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			},
			Vector: turbopuffer.VectorParam{
				FloatArray: []float64{0},
			},
		}},
	})
	if err != nil {
		var apierr *turbopuffer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
