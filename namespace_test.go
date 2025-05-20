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
		option.WithAPIKey("My API Key"),
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
		option.WithAPIKey("My API Key"),
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

func TestNamespaceMultiQueryWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Namespaces.MultiQuery(context.TODO(), turbopuffer.NamespaceMultiQueryParams{
		Namespace: turbopuffer.String("namespace"),
		Consistency: turbopuffer.NamespaceMultiQueryParamsConsistency{
			Level: turbopuffer.NamespaceMultiQueryParamsConsistencyLevelStrong,
		},
		Queries: []turbopuffer.NamespaceMultiQueryParamsQuery{{
			DistanceMetric: turbopuffer.DistanceMetricCosineDistance,
			Filters:        map[string]interface{}{},
			IncludeAttributes: turbopuffer.NamespaceMultiQueryParamsQueryIncludeAttributesUnion{
				OfBool: turbopuffer.Bool(true),
			},
			RankBy: map[string]interface{}{},
			TopK:   turbopuffer.Int(0),
		}},
		VectorEncoding: turbopuffer.NamespaceMultiQueryParamsVectorEncodingFloat,
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
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Namespaces.Query(context.TODO(), turbopuffer.NamespaceQueryParams{
		Namespace: turbopuffer.String("namespace"),
		Consistency: turbopuffer.NamespaceQueryParamsConsistency{
			Level: turbopuffer.NamespaceQueryParamsConsistencyLevelStrong,
		},
		DistanceMetric: turbopuffer.DistanceMetricCosineDistance,
		Filters:        map[string]interface{}{},
		IncludeAttributes: turbopuffer.NamespaceQueryParamsIncludeAttributesUnion{
			OfBool: turbopuffer.Bool(true),
		},
		RankBy:         map[string]interface{}{},
		TopK:           turbopuffer.Int(0),
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
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Namespaces.Write(context.TODO(), turbopuffer.NamespaceWriteParams{
		Namespace:         turbopuffer.String("namespace"),
		CopyFromNamespace: turbopuffer.String("copy_from_namespace"),
		DeleteByFilter:    map[string]interface{}{},
		Deletes: []turbopuffer.IDUnionParam{{
			OfString: turbopuffer.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		}},
		DistanceMetric: turbopuffer.DistanceMetricCosineDistance,
		PatchColumns: turbopuffer.DocumentColumnsParam{
			ID: []turbopuffer.IDUnionParam{{
				OfString: turbopuffer.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}},
		},
		PatchRows: []turbopuffer.DocumentRowParam{{
			ID: turbopuffer.IDUnionParam{
				OfString: turbopuffer.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			},
			Vector: turbopuffer.DocumentRowVectorUnionParam{
				OfFloatArray: []float64{0},
			},
		}},
		Schema: map[string]turbopuffer.AttributeSchemaParam{
			"foo": {
				Filterable: turbopuffer.Bool(true),
				FullTextSearch: turbopuffer.AttributeSchemaFullTextSearchUnionParam{
					OfBool: turbopuffer.Bool(true),
				},
				Type: turbopuffer.AttributeSchemaTypeString,
			},
		},
		UpsertColumns: turbopuffer.DocumentColumnsParam{
			ID: []turbopuffer.IDUnionParam{{
				OfString: turbopuffer.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}},
		},
		UpsertRows: []turbopuffer.DocumentRowParam{{
			ID: turbopuffer.IDUnionParam{
				OfString: turbopuffer.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			},
			Vector: turbopuffer.DocumentRowVectorUnionParam{
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
