package turbopuffer_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/turbopuffer/turbopuffer-go"
)

type testContext struct {
	ctx    context.Context
	client turbopuffer.Client
	prefix string
}

func setup(t *testing.T) testContext {
	return testContext{
		ctx:    context.Background(),
		client: turbopuffer.NewClient(),
		prefix: fmt.Sprintf("test-go-%s", t.Name()),
	}
}

func TestTurbopufferFullTextSearchSchema(t *testing.T) {
	tctx := setup(t)

	roundtripSchema := func(
		t *testing.T,
		schema map[string]turbopuffer.AttributeSchemaConfigParam,
	) turbopuffer.NamespaceSchemaResponse {
		ns := tctx.client.Namespace(fmt.Sprintf("%s-fts-default", tctx.prefix))
		_, err := ns.Write(tctx.ctx, turbopuffer.NamespaceWriteParams{
			UpsertRows: []turbopuffer.RowParam{
				{
					ID: turbopuffer.IDParam{
						Int: turbopuffer.Int(1),
					},
					ExtraFields: map[string]any{"test-attr": "test-value"},
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
				Type: turbopuffer.String("string"),
			},
		})
		if schema["test-attr"].JSON.FullTextSearch.Valid() {
			t.Fatal("FullTextSearchConfig should be omitted")
		}
	})

	t.Run("present-but-empty-full-text-search", func(t *testing.T) {
		schema := roundtripSchema(t, map[string]turbopuffer.AttributeSchemaConfigParam{
			"test-attr": {
				Type:           turbopuffer.String("string"),
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

	t.Run("present-but-empty-full-text-search", func(t *testing.T) {
		schema := roundtripSchema(t, map[string]turbopuffer.AttributeSchemaConfigParam{
			"test-attr": {
				Type: turbopuffer.String("string"),
				FullTextSearch: &turbopuffer.FullTextSearchConfigParam{
					Language: turbopuffer.LanguageDutch,
				},
			},
		})
		if !schema["test-attr"].JSON.FullTextSearch.Valid() {
			t.Fatal("FullTextSearchConfig should be valid")
		}
		if schema["test-attr"].FullTextSearch.Language != turbopuffer.LanguageDutch {
			t.Fatal("FullTextSearchConfig should contain default values")
		}
	})

}
