package turbopuffer_test

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/turbopuffer/turbopuffer-go/v2"
	"github.com/turbopuffer/turbopuffer-go/v2/option"
)

var nonce = fmt.Sprintf("%d", time.Now().UnixMilli())

type testContext struct {
	ctx    context.Context
	client turbopuffer.Client
	prefix string
}

func setup(t *testing.T) testContext {
	return testContext{
		ctx:    context.Background(),
		client: turbopuffer.NewClient(option.WithRegion("gcp-us-central1")),
		prefix: fmt.Sprintf("test-go-%s-%s", nonce, t.Name()),
	}
}

// TestTurbopufferFullTextSearchSchema tests the custom behavior of the
// `FullTextSearch` field in the `AttributeSchemaConfigParam` type.
//
// If the field is omitted or set to nil, full-text search is disabled
// (equivalent of explicit `false`). If the field is present but empty,
// full-text search is enabled with default values (equivalent of explicit
// `true`). Otherwise the full-text search is configured as specified.
//
// We do it this way because it's more idiomatic Go than the union struct that
// Stainless will generate for the type in the OpenAPI specification (i.e.,
// `bool | FullTextSearchConfigParam`).
func TestTurbopufferFullTextSearchSchema(t *testing.T) {
	tctx := setup(t)

	roundtripSchema := func(
		t *testing.T,
		schema map[string]turbopuffer.AttributeSchemaConfigParam,
	) turbopuffer.NamespaceSchemaResponse {
		subtestName := strings.SplitN(t.Name(), "/", 2)[1]
		ns := tctx.client.Namespace(fmt.Sprintf("%s-fts-%s", tctx.prefix, subtestName))
		_, err := ns.Write(tctx.ctx, turbopuffer.NamespaceWriteParams{
			UpsertRows: []turbopuffer.RowParam{
				{
					"id":        1,
					"test-attr": "test-value",
				},
			},
			Schema: schema,
		})
		if err != nil {
			t.Fatal(err)
		}
		res, err := ns.Schema(tctx.ctx, turbopuffer.NamespaceSchemaParams{})
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("result schema: %s", turbopuffer.PrettyPrint(res))
		return *res
	}

	t.Run("omit-full-text-search", func(t *testing.T) {
		schema := roundtripSchema(t, map[string]turbopuffer.AttributeSchemaConfigParam{
			"test-attr": {
				Type: turbopuffer.AttributeType("string"),
			},
		})
		if schema["test-attr"].JSON.FullTextSearch.Valid() {
			t.Fatal("FullTextSearchConfig should be omitted")
		}
	})

	t.Run("present-but-empty-full-text-search", func(t *testing.T) {
		schema := roundtripSchema(t, map[string]turbopuffer.AttributeSchemaConfigParam{
			"test-attr": {
				Type:           turbopuffer.AttributeType("string"),
				FullTextSearch: &turbopuffer.FullTextSearchConfigParam{},
			},
		})
		if !schema["test-attr"].JSON.FullTextSearch.Valid() {
			t.Fatal("FullTextSearchConfig should be valid")
		}
		// Sanity check one of the defaults that's likely to be stable over
		// time. We don't want to check all the defaults because they're subject
		// to change over time.
		if schema["test-attr"].FullTextSearch.Language != turbopuffer.LanguageEnglish {
			t.Fatal("FullTextSearchConfig should contain default values")
		}
	})

	t.Run("present-nonempty-full-text-search", func(t *testing.T) {
		schema := roundtripSchema(t, map[string]turbopuffer.AttributeSchemaConfigParam{
			"test-attr": {
				Type: turbopuffer.AttributeType("string"),
				FullTextSearch: &turbopuffer.FullTextSearchConfigParam{
					Language: turbopuffer.LanguageDutch,
				},
			},
		})
		if !schema["test-attr"].JSON.FullTextSearch.Valid() {
			t.Fatal("FullTextSearchConfig should be valid")
		}
		if schema["test-attr"].FullTextSearch.Language != turbopuffer.LanguageDutch {
			t.Fatal("FullTextSearchConfig should contain specified values")
		}
	})
}

// The `ToParam` method used to be broken for unions. This test protects against
// regressions.
// See: https://turbopuffer-internal.slack.com/archives/C08EK13Q98E/p1748933005836679
func TestToParamSerialization(t *testing.T) {
	type testCase struct {
		name string
		json string
	}

	for _, tc := range []testCase{
		{name: "int", json: "123"},
		{name: "string", json: `"abc"`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var id turbopuffer.ID
			err := id.UnmarshalJSON([]byte(tc.json))
			if err != nil {
				t.Fatal(err)
			}
			idParam := id.ToParam()
			idParamJSON, err := json.Marshal(idParam)
			if err != nil {
				t.Fatal(err)
			}
			var idRoundtripped turbopuffer.ID
			err = idRoundtripped.UnmarshalJSON(idParamJSON)
			if err != nil {
				t.Fatal(err)
			}
			if id != idRoundtripped {
				t.Fatal("id and idRoundtripped should be the same")
			}
		})
	}
}

// TestComputeAttributesSerialization verifies that each variant of the
// ComputeAttributes union serializes to its expected turbolisp wire format
// when used as a value of the compute_attributes map. compute_attributes is a
// request-only param, so this asserts the marshaled JSON rather than a full
// round-trip.
func TestComputeAttributesSerialization(t *testing.T) {
	for _, tc := range []struct {
		name string
		val  turbopuffer.ComputeAttributes
		want string
	}{
		{"vector_dist", turbopuffer.NewComputeAttributesVectorDist("vec", []float32{0.5}), `["vec","VectorDist",[0.5]]`},
		{"highlight", turbopuffer.NewComputeAttributesHighlight("body"), `["Highlight","body"]`},
		{"highlight_with_config", turbopuffer.NewComputeAttributesHighlightWithConfig("body", turbopuffer.HighlightConfigParams{}), `["Highlight","body",{}]`},
		{"score_rank_by", turbopuffer.NewRankByAnn("vec", []float32{0.5}), `["vec","ANN",[0.5]]`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := json.Marshal(map[string]turbopuffer.ComputeAttributes{"attr": tc.val})
			if err != nil {
				t.Fatal(err)
			}
			want := `{"attr":` + tc.want + `}`
			if string(got) != want {
				t.Fatalf("compute_attributes[%s]: got %s, want %s", tc.name, got, want)
			}
		})
	}
}
