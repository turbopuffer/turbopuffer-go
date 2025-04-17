// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turbopuffer_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/turbopuffer-go"
	"github.com/stainless-sdks/turbopuffer-go/internal/testutil"
	"github.com/stainless-sdks/turbopuffer-go/option"
)

func TestNamespaceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Namespaces.List(context.TODO(), turbopuffer.NamespaceListParams{
		Cursor:   turbopuffer.String("cursor"),
		PageSize: turbopuffer.Int(1),
		Prefix:   turbopuffer.String("prefix"),
	})
	if err != nil {
		var apierr *turbopuffer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

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
	_, err := client.Namespaces.DeleteAll(context.TODO(), "namespace")
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
	_, err := client.Namespaces.GetSchema(context.TODO(), "namespace")
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
	_, err := client.Namespaces.Query(
		context.TODO(),
		"namespace",
		turbopuffer.NamespaceQueryParams{
			Consistency: turbopuffer.NamespaceQueryParamsConsistency{
				Level: turbopuffer.NamespaceQueryParamsConsistencyLevelStrong,
			},
			DistanceMetric: turbopuffer.DistanceMetricCosineDistance,
			Filters:        map[string]interface{}{},
			IncludeAttributes: turbopuffer.NamespaceQueryParamsIncludeAttributesUnion{
				OfBool: turbopuffer.Bool(true),
			},
			IncludeVectors: turbopuffer.Bool(true),
			RankBy:         map[string]interface{}{},
			TopK:           turbopuffer.Int(0),
			Vector:         []float64{0},
		},
	)
	if err != nil {
		var apierr *turbopuffer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNamespaceUpsertWithOptionalParams(t *testing.T) {
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
	_, err := client.Namespaces.Upsert(
		context.TODO(),
		"namespace",
		turbopuffer.NamespaceUpsertParams{
			Documents: turbopuffer.NamespaceUpsertParamsDocumentsUpsertColumnar{
				DocumentColumnsParam: turbopuffer.DocumentColumnsParam{
					Attributes: map[string][]map[string]any{
						"foo": {{
							"foo": "bar",
						}},
					},
					IDs: []turbopuffer.IDUnionParam{{
						OfString: turbopuffer.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					}},
					Vectors: [][]float64{{0}},
				},
				DistanceMetric: turbopuffer.DistanceMetricCosineDistance,
				Schema: map[string][]turbopuffer.AttributeSchemaParam{
					"foo": {{
						Filterable: turbopuffer.Bool(true),
						FullTextSearch: turbopuffer.AttributeSchemaFullTextSearchUnionParam{
							OfBool: turbopuffer.Bool(true),
						},
						Type: turbopuffer.AttributeSchemaTypeString,
					}},
				},
			},
		},
	)
	if err != nil {
		var apierr *turbopuffer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
