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
	ns := client.Namespace("ns")
	_, err := ns.DeleteAll(context.TODO(), turbopuffer.NamespaceDeleteAllParams{})
	if err != nil {
		var apierr *turbopuffer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNamespaceGetSchema(t *testing.T) {
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
	ns := client.Namespace("ns")
	_, err := ns.GetSchema(context.TODO(), turbopuffer.NamespaceGetSchemaParams{})
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
	ns := client.Namespace("ns")
	_, err := ns.HintCacheWarm(context.TODO(), turbopuffer.NamespaceHintCacheWarmParams{})
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
	ns := client.Namespace("ns")
	_, err := ns.Query(context.TODO(), turbopuffer.NamespaceQueryParams{
		RankBy: turbopuffer.NewRankByVector("vector", []float64{0}),
		TopK:   0,
		Consistency: turbopuffer.NamespaceQueryParamsConsistency{
			Level: turbopuffer.NamespaceQueryParamsConsistencyLevelStrong,
		},
		DistanceMetric: turbopuffer.DistanceMetricCosineDistance,
		IncludeAttributes: turbopuffer.IncludeAttributesParam{
			Bool: turbopuffer.Bool(true),
		},
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
	ns := client.Namespace("ns")
	_, err := ns.Recall(context.TODO(), turbopuffer.NamespaceRecallParams{
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
	ns := client.Namespace("ns")
	_, err := ns.UpdateSchema(context.TODO(), turbopuffer.NamespaceUpdateSchemaParams{
		Schema: map[string]turbopuffer.AttributeSchemaParam{
			"foo": {
				Ann:        turbopuffer.Bool(true),
				Filterable: turbopuffer.Bool(true),
				FullTextSearch: turbopuffer.FullTextSearchParam{
					Bool: turbopuffer.Bool(true),
				},
				Type: turbopuffer.AttributeTypeString,
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
	ns := client.Namespace("ns")
	_, err := ns.Write(context.TODO(), turbopuffer.NamespaceWriteParams{
		CopyFromNamespace: turbopuffer.String("copy_from_namespace"),
		DeleteByFilter:    turbopuffer.NewFilterEq("id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
				Ann:        turbopuffer.Bool(true),
				Filterable: turbopuffer.Bool(true),
				FullTextSearch: turbopuffer.FullTextSearchParam{
					Bool: turbopuffer.Bool(true),
				},
				Type: turbopuffer.AttributeTypeString,
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
