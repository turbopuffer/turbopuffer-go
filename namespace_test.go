// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turbopuffer_test

import (
	"context"
	"errors"
	"net/url"
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
		option.WithRegion("gcp-us-central1"),
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
		option.WithRegion("gcp-us-central1"),
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
		option.WithRegion("gcp-us-central1"),
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
		option.WithRegion("gcp-us-central1"),
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
		option.WithRegion("gcp-us-central1"),
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
		option.WithRegion("gcp-us-central1"),
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

// Tests for URL path encoding fix for namespace names with special characters
func TestNamespacePathEncoding(t *testing.T) {
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

	// Test cases with various special characters that need URL encoding
	testCases := []struct {
		name      string
		namespace string
	}{
		{
			name:      "simple_namespace",
			namespace: "simple-namespace",
		},
		{
			name:      "namespace_with_spaces",
			namespace: "namespace with spaces",
		},
		{
			name:      "namespace_with_slashes",
			namespace: "namespace/with/slashes",
		},
		{
			name:      "namespace_with_hash",
			namespace: "namespace#with#hash",
		},
		{
			name:      "namespace_with_percent",
			namespace: "namespace%already%encoded",
		},
		{
			name:      "unicode_namespace",
			namespace: "用户数据",
		},
		{
			name:      "complex_namespace",
			namespace: "test namespace/with@special:chars?param=value",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create namespace with special characters - this should not panic or error
			ns := client.Namespace(tc.namespace)

			// Test that we can get the namespace ID (which internally uses the encoded path)
			namespaceID := ns.ID()
			if namespaceID != tc.namespace {
				t.Errorf("Expected namespace ID %s, got %s", tc.namespace, namespaceID)
			}
		})
	}
}

// Test URL encoding behavior specifically
func TestURLPathEncodingBehavior(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"simple", "simple"},
		{"with spaces", "with%20spaces"},
		{"with/slashes", "with%2Fslashes"},
		{"with#hash", "with%23hash"},
		{"with%percent", "with%25percent"},
		{"with?query", "with%3Fquery"},
		{"with&ampersand", "with&ampersand"},
		{"with:colon", "with:colon"},
		{"用户数据", "%E7%94%A8%E6%88%B7%E6%95%B0%E6%8D%AE"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			encoded := url.PathEscape(tc.input)
			if encoded != tc.expected {
				t.Errorf("url.PathEscape(%q) = %q, want %q", tc.input, encoded, tc.expected)
			}
		})
	}
}

// Test that namespace operations handle encoded paths correctly
func TestNamespaceOperationsWithEncodedPaths(t *testing.T) {
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

	// Test with namespace containing spaces
	ns := client.Namespace("test namespace with spaces")

	// Test Schema operation
	_, err := ns.Schema(context.TODO(), turbopuffer.NamespaceSchemaParams{})
	if err != nil {
		var apierr *turbopuffer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("Schema operation failed with encoded namespace: %s", err.Error())
	}
}

// Test that all namespace operations work with special characters
func TestAllNamespaceOperationsWithSpecialChars(t *testing.T) {
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

	// Test with namespace containing multiple special characters
	specialNamespace := "test/namespace with@special:chars?param=value#hash"
	ns := client.Namespace(specialNamespace)

	// Test DeleteAll operation
	_, err := ns.DeleteAll(context.TODO(), turbopuffer.NamespaceDeleteAllParams{})
	if err != nil {
		var apierr *turbopuffer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("DeleteAll operation failed with encoded namespace: %s", err.Error())
	}

	// Test Write operation
	_, err = ns.Write(context.TODO(), turbopuffer.NamespaceWriteParams{
		UpsertColumns: turbopuffer.ColumnsParam{},
		UpsertRows:    []turbopuffer.RowParam{},
	})
	if err != nil {
		var apierr *turbopuffer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("Write operation failed with encoded namespace: %s", err.Error())
	}

	// Test Query operation
	_, err = ns.Query(context.TODO(), turbopuffer.NamespaceQueryParams{
		TopK:   turbopuffer.Int(10),
		RankBy: turbopuffer.NewRankByVector("vector", []float32{0.1, 0.2, 0.3}),
	})
	if err != nil {
		var apierr *turbopuffer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("Query operation failed with encoded namespace: %s", err.Error())
	}
}
