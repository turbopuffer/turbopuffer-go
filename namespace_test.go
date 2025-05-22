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
	"github.com/turbopuffer/turbopuffer-go/shared"
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
	_, err := client.Namespaces.GetSchema(context.TODO(), turbopuffer.NamespaceGetSchemaParams{
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
		RankBy:    map[string]interface{}{},
		TopK:      0,
		Consistency: turbopuffer.NamespaceQueryParamsConsistency{
			Level: turbopuffer.NamespaceQueryParamsConsistencyLevelStrong,
		},
		DistanceMetric: shared.DistanceMetricCosineDistance,
		Filters:        map[string]interface{}{},
		IncludeAttributes: turbopuffer.NamespaceQueryParamsIncludeAttributesUnion{
			OfBool: turbopuffer.Bool(true),
		},
		VectorEncoding: turbopuffer.NamespaceQueryParamsVectorEncodingFloat,
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
		Queries:   []any{map[string]interface{}{}},
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
	_, err := client.Namespaces.UpdateSchema(context.TODO(), turbopuffer.NamespaceUpdateSchemaParams{
		Namespace: turbopuffer.String("namespace"),
		Schema: map[string]shared.AttributeSchemaParam{
			"foo": {
				Filterable: turbopuffer.Bool(true),
				FullTextSearch: shared.AttributeSchemaFullTextSearchUnionParam{
					OfBool: turbopuffer.Bool(true),
				},
				Type: shared.AttributeSchemaTypeString,
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

func TestNamespaceWarmCache(t *testing.T) {
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
	_, err := client.Namespaces.WarmCache(context.TODO(), turbopuffer.NamespaceWarmCacheParams{
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
		Deletes: []shared.IDUnionParam{{
			OfString: turbopuffer.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		}},
		DistanceMetric: shared.DistanceMetricCosineDistance,
		PatchColumns: shared.DocumentColumnsParam{
			ID: []shared.IDUnionParam{{
				OfString: turbopuffer.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}},
		},
		PatchRows: []shared.DocumentRowParam{{
			ID: shared.IDUnionParam{
				OfString: turbopuffer.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			},
			Vector: shared.DocumentRowVectorUnionParam{
				OfFloatArray: []float64{0},
			},
		}},
		Schema: map[string]shared.AttributeSchemaParam{
			"foo": {
				Filterable: turbopuffer.Bool(true),
				FullTextSearch: shared.AttributeSchemaFullTextSearchUnionParam{
					OfBool: turbopuffer.Bool(true),
				},
				Type: shared.AttributeSchemaTypeString,
			},
		},
		UpsertColumns: shared.DocumentColumnsParam{
			ID: []shared.IDUnionParam{{
				OfString: turbopuffer.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}},
		},
		UpsertRows: []shared.DocumentRowParam{{
			ID: shared.IDUnionParam{
				OfString: turbopuffer.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			},
			Vector: shared.DocumentRowVectorUnionParam{
				OfFloatArray: []float64{0},
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
