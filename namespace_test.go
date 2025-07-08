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
		option.WithAPIKey("tpuf_A1..."),
	)
	ns := client.Namespace("ns")
	_, err := ns.MultiQuery(context.TODO(), turbopuffer.NamespaceMultiQueryParams{
		Namespace: turbopuffer.String("namespace"),
		Queries: []turbopuffer.QueryParam{{
			DistanceMetric: turbopuffer.DistanceMetricCosineDistance,
			IncludeAttributes: turbopuffer.IncludeAttributesParam{
				Bool: turbopuffer.Bool(true),
			},
			RankBy: turbopuffer.NewRankByVector("vector", []float32{0}),
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
	)
	ns := client.Namespace("ns")
	_, err := ns.Query(context.TODO(), turbopuffer.NamespaceQueryParams{
		Namespace: turbopuffer.String("namespace"),
		Consistency: turbopuffer.NamespaceQueryParamsConsistency{
			Level: turbopuffer.NamespaceQueryParamsConsistencyLevelStrong,
		},
		DistanceMetric: turbopuffer.DistanceMetricCosineDistance,
		IncludeAttributes: turbopuffer.IncludeAttributesParam{
			Bool: turbopuffer.Bool(true),
		},
		RankBy:         turbopuffer.NewRankByVector("vector", []float32{0}),
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
	)
	ns := client.Namespace("ns")
	_, err := ns.Recall(context.TODO(), turbopuffer.NamespaceRecallParams{
		Namespace: turbopuffer.String("namespace"),
		Filters:   map[string]interface{}{},
		Num:       turbopuffer.Int(0),
		Queries:   []float32{0},
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
	)
	ns := client.Namespace("ns")
	_, err := ns.Schema(context.TODO(), turbopuffer.NamespaceSchemaParams{})
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
	)
	ns := client.Namespace("ns")
	_, err := ns.UpdateSchema(context.TODO(), turbopuffer.NamespaceUpdateSchemaParams{
		Schema: map[string]turbopuffer.AttributeSchemaConfigParam{
			"foo": {Ann: turbopuffer.Bool(true), Filterable: turbopuffer.Bool(true), FullTextSearch: &turbopuffer.FullTextSearchConfigParam{B: turbopuffer.Float(0), CaseSensitive: turbopuffer.Bool(true), K1: turbopuffer.Float(0), Language: turbopuffer.LanguageArabic, RemoveStopwords: turbopuffer.Bool(true), Stemming: turbopuffer.Bool(true), Tokenizer: turbopuffer.TokenizerPreTokenizedArray}, Type: turbopuffer.String("string")},
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
	)
	ns := client.Namespace("ns")
	_, err := ns.Write(context.TODO(), turbopuffer.NamespaceWriteParams{
		CopyFromNamespace: turbopuffer.String("copy_from_namespace"),
		DeleteByFilter:    turbopuffer.NewFilterEq("id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Deletes:           []any{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		DistanceMetric:    turbopuffer.DistanceMetricCosineDistance,
		Encryption: turbopuffer.NamespaceWriteParamsEncryption{
			Cmek: turbopuffer.NamespaceWriteParamsEncryptionCmek{
				KeyName: "key_name",
			},
		},
		PatchColumns: turbopuffer.ColumnsParam{},
		PatchRows:    []turbopuffer.RowParam{},
		Schema: map[string]turbopuffer.AttributeSchemaConfigParam{
			"foo": {Ann: turbopuffer.Bool(true), Filterable: turbopuffer.Bool(true), FullTextSearch: &turbopuffer.FullTextSearchConfigParam{B: turbopuffer.Float(0), CaseSensitive: turbopuffer.Bool(true), K1: turbopuffer.Float(0), Language: turbopuffer.LanguageArabic, RemoveStopwords: turbopuffer.Bool(true), Stemming: turbopuffer.Bool(true), Tokenizer: turbopuffer.TokenizerPreTokenizedArray}, Type: turbopuffer.String("string")},
		},
		UpsertColumns:   turbopuffer.ColumnsParam{},
		UpsertRows:      []turbopuffer.RowParam{},
		DeleteCondition: turbopuffer.Filter(nil),
		PatchCondition:  turbopuffer.Filter(nil),
		UpsertCondition: turbopuffer.Filter(nil),
	})
	if err != nil {
		var apierr *turbopuffer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
