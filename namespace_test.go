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
	t.Skip("Prism tests are disabled")
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

func TestNamespaceExplainQueryWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	)
	_, err := client.Namespaces.ExplainQuery(context.TODO(), turbopuffer.NamespaceExplainQueryParams{
		Namespace: turbopuffer.String("namespace"),
		AggregateBy: map[string]any{
			"foo": "bar",
		},
		Consistency: turbopuffer.NamespaceExplainQueryParamsConsistency{
			Level: turbopuffer.NamespaceExplainQueryParamsConsistencyLevelStrong,
		},
		DistanceMetric:    turbopuffer.DistanceMetricCosineDistance,
		ExcludeAttributes: []string{"string"},
		Filters:           map[string]any{},
		GroupBy:           []string{"string"},
		IncludeAttributes: turbopuffer.IncludeAttributesParam{
			Bool: turbopuffer.Bool(true),
		},
		RankBy:         map[string]any{},
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

func TestNamespaceHintCacheWarm(t *testing.T) {
	t.Skip("Prism tests are disabled")
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

func TestNamespaceMetadata(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	)
	_, err := client.Namespaces.Metadata(context.TODO(), turbopuffer.NamespaceMetadataParams{
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
	t.Skip("Prism tests are disabled")
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
	)
	_, err := client.Namespaces.MultiQuery(context.TODO(), turbopuffer.NamespaceMultiQueryParams{
		Namespace: turbopuffer.String("namespace"),
		Queries: []turbopuffer.QueryParam{{
			AggregateBy: map[string]any{
				"foo": "bar",
			},
			DistanceMetric:    turbopuffer.DistanceMetricCosineDistance,
			ExcludeAttributes: []string{"string"},
			Filters:           map[string]any{},
			GroupBy:           []string{"string"},
			IncludeAttributes: turbopuffer.IncludeAttributesParam{
				Bool: turbopuffer.Bool(true),
			},
			RankBy: map[string]any{},
			TopK:   turbopuffer.Int(0),
		}},
		Consistency: turbopuffer.NamespaceMultiQueryParamsConsistency{
			Level: turbopuffer.NamespaceMultiQueryParamsConsistencyLevelStrong,
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

func TestNamespaceQueryWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	)
	_, err := client.Namespaces.Query(context.TODO(), turbopuffer.NamespaceQueryParams{
		Namespace: turbopuffer.String("namespace"),
		AggregateBy: map[string]any{
			"foo": "bar",
		},
		Consistency: turbopuffer.NamespaceQueryParamsConsistency{
			Level: turbopuffer.NamespaceQueryParamsConsistencyLevelStrong,
		},
		DistanceMetric:    turbopuffer.DistanceMetricCosineDistance,
		ExcludeAttributes: []string{"string"},
		Filters:           map[string]any{},
		GroupBy:           []string{"string"},
		IncludeAttributes: turbopuffer.IncludeAttributesParam{
			Bool: turbopuffer.Bool(true),
		},
		RankBy:         map[string]any{},
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
	t.Skip("Prism tests are disabled")
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
	)
	_, err := client.Namespaces.Recall(context.TODO(), turbopuffer.NamespaceRecallParams{
		Namespace:          turbopuffer.String("namespace"),
		Filters:            map[string]any{},
		IncludeGroundTruth: turbopuffer.Bool(true),
		Num:                turbopuffer.Int(0),
		Queries:            []float64{0},
		TopK:               turbopuffer.Int(0),
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
	t.Skip("Prism tests are disabled")
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
	t.Skip("Prism tests are disabled")
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
	)
	_, err := client.Namespaces.UpdateSchema(context.TODO(), turbopuffer.NamespaceUpdateSchemaParams{
		Namespace: turbopuffer.String("namespace"),
		Schema: map[string]turbopuffer.AttributeSchemaConfigParam{
			"foo": {
				Type: "string",
				Ann: turbopuffer.AttributeSchemaConfigAnnParam{
					DistanceMetric: turbopuffer.DistanceMetricCosineDistance,
				},
				Filterable: turbopuffer.Bool(true),
				FullTextSearch: turbopuffer.FullTextSearchConfigParam{
					AsciiFolding:    turbopuffer.Bool(true),
					B:               turbopuffer.Float(0),
					CaseSensitive:   turbopuffer.Bool(true),
					K1:              turbopuffer.Float(0),
					Language:        turbopuffer.LanguageArabic,
					MaxTokenLength:  turbopuffer.Int(0),
					RemoveStopwords: turbopuffer.Bool(true),
					Stemming:        turbopuffer.Bool(true),
					Tokenizer:       turbopuffer.TokenizerPreTokenizedArray,
				},
				Regex: turbopuffer.Bool(true),
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
	t.Skip("Prism tests are disabled")
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
	)
	_, err := client.Namespaces.Write(context.TODO(), turbopuffer.NamespaceWriteParams{
		Namespace: turbopuffer.String("namespace"),
		CopyFromNamespace: turbopuffer.NamespaceWriteParamsCopyFromNamespace{
			String: turbopuffer.String("string"),
		},
		DeleteByFilter:             map[string]any{},
		DeleteByFilterAllowPartial: turbopuffer.Bool(true),
		DeleteCondition:            map[string]any{},
		Deletes: []turbopuffer.IDParam{{
			String: turbopuffer.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		}},
		DisableBackpressure: turbopuffer.Bool(true),
		DistanceMetric:      turbopuffer.DistanceMetricCosineDistance,
		Encryption: turbopuffer.NamespaceWriteParamsEncryption{
			Cmek: turbopuffer.NamespaceWriteParamsEncryptionCmek{
				KeyName: "key_name",
			},
		},
		PatchByFilter: turbopuffer.NamespaceWriteParamsPatchByFilter{
			Filters: map[string]any{},
			Patch: map[string]any{
				"foo": "bar",
			},
		},
		PatchByFilterAllowPartial: turbopuffer.Bool(true),
		PatchColumns: turbopuffer.ColumnsParam{
			ID: []turbopuffer.IDParam{{
				String: turbopuffer.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}},
			Vector: turbopuffer.ColumnsVectorParam{
				VectorArray: []turbopuffer.VectorParam{{
					FloatArray: []float64{0},
				}},
			},
		},
		PatchCondition: map[string]any{},
		PatchRows: []turbopuffer.RowParam{{
			ID: turbopuffer.IDParam{
				String: turbopuffer.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			},
			Vector: turbopuffer.VectorParam{
				FloatArray: []float64{0},
			},
		}},
		Schema: map[string]turbopuffer.AttributeSchemaConfigParam{
			"foo": {
				Type: "string",
				Ann: turbopuffer.AttributeSchemaConfigAnnParam{
					DistanceMetric: turbopuffer.DistanceMetricCosineDistance,
				},
				Filterable: turbopuffer.Bool(true),
				FullTextSearch: turbopuffer.FullTextSearchConfigParam{
					AsciiFolding:    turbopuffer.Bool(true),
					B:               turbopuffer.Float(0),
					CaseSensitive:   turbopuffer.Bool(true),
					K1:              turbopuffer.Float(0),
					Language:        turbopuffer.LanguageArabic,
					MaxTokenLength:  turbopuffer.Int(0),
					RemoveStopwords: turbopuffer.Bool(true),
					Stemming:        turbopuffer.Bool(true),
					Tokenizer:       turbopuffer.TokenizerPreTokenizedArray,
				},
				Regex: turbopuffer.Bool(true),
			},
		},
		UpsertColumns: turbopuffer.ColumnsParam{
			ID: []turbopuffer.IDParam{{
				String: turbopuffer.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}},
			Vector: turbopuffer.ColumnsVectorParam{
				VectorArray: []turbopuffer.VectorParam{{
					FloatArray: []float64{0},
				}},
			},
		},
		UpsertCondition: map[string]any{},
		UpsertRows: []turbopuffer.RowParam{{
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
