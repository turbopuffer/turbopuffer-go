// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turbopuffer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/turbopuffer/turbopuffer-go/internal/apijson"
	shimjson "github.com/turbopuffer/turbopuffer-go/internal/encoding/json"
	"github.com/turbopuffer/turbopuffer-go/internal/requestconfig"
	"github.com/turbopuffer/turbopuffer-go/option"
	"github.com/turbopuffer/turbopuffer-go/packages/param"
	"github.com/turbopuffer/turbopuffer-go/packages/respjson"
	"github.com/turbopuffer/turbopuffer-go/shared/constant"
)

// NamespaceService contains methods and other services that help with interacting
// with the turbopuffer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNamespaceService] method instead.
type NamespaceService struct {
	Options []option.RequestOption
}

// NewNamespaceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNamespaceService(opts ...option.RequestOption) (r NamespaceService) {
	r = NamespaceService{}
	r.Options = opts
	return
}

// Delete namespace.
func (r *NamespaceService) DeleteAll(ctx context.Context, body NamespaceDeleteAllParams, opts ...option.RequestOption) (res *NamespaceDeleteAllResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.Namespace, precfg.DefaultNamespace)
	if body.Namespace.Value == "" {
		err = errors.New("missing required namespace parameter")
		return
	}
	path := fmt.Sprintf("v2/namespaces/%s", url.PathEscape(body.Namespace.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Explain a query plan.
func (r *NamespaceService) ExplainQuery(ctx context.Context, params NamespaceExplainQueryParams, opts ...option.RequestOption) (res *NamespaceExplainQueryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.Namespace, precfg.DefaultNamespace)
	if params.Namespace.Value == "" {
		err = errors.New("missing required namespace parameter")
		return
	}
	path := fmt.Sprintf("v2/namespaces/%s/explain_query", url.PathEscape(params.Namespace.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Signal turbopuffer to prepare for low-latency requests.
func (r *NamespaceService) HintCacheWarm(ctx context.Context, query NamespaceHintCacheWarmParams, opts ...option.RequestOption) (res *NamespaceHintCacheWarmResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.Namespace, precfg.DefaultNamespace)
	if query.Namespace.Value == "" {
		err = errors.New("missing required namespace parameter")
		return
	}
	path := fmt.Sprintf("v1/namespaces/%s/hint_cache_warm", url.PathEscape(query.Namespace.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get metadata about a namespace.
func (r *NamespaceService) Metadata(ctx context.Context, query NamespaceMetadataParams, opts ...option.RequestOption) (res *NamespaceMetadata, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.Namespace, precfg.DefaultNamespace)
	if query.Namespace.Value == "" {
		err = errors.New("missing required namespace parameter")
		return
	}
	path := fmt.Sprintf("v1/namespaces/%s/metadata", url.PathEscape(query.Namespace.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Issue multiple concurrent queries filter or search documents.
func (r *NamespaceService) MultiQuery(ctx context.Context, params NamespaceMultiQueryParams, opts ...option.RequestOption) (res *NamespaceMultiQueryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.Namespace, precfg.DefaultNamespace)
	if params.Namespace.Value == "" {
		err = errors.New("missing required namespace parameter")
		return
	}
	path := fmt.Sprintf("v2/namespaces/%s/query?stainless_overload=multiQuery", url.PathEscape(params.Namespace.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Query, filter, full-text search and vector search documents.
func (r *NamespaceService) Query(ctx context.Context, params NamespaceQueryParams, opts ...option.RequestOption) (res *NamespaceQueryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.Namespace, precfg.DefaultNamespace)
	if params.Namespace.Value == "" {
		err = errors.New("missing required namespace parameter")
		return
	}
	path := fmt.Sprintf("v2/namespaces/%s/query", url.PathEscape(params.Namespace.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Evaluate recall.
func (r *NamespaceService) Recall(ctx context.Context, params NamespaceRecallParams, opts ...option.RequestOption) (res *NamespaceRecallResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.Namespace, precfg.DefaultNamespace)
	if params.Namespace.Value == "" {
		err = errors.New("missing required namespace parameter")
		return
	}
	path := fmt.Sprintf("v1/namespaces/%s/_debug/recall", url.PathEscape(params.Namespace.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get namespace schema.
func (r *NamespaceService) Schema(ctx context.Context, query NamespaceSchemaParams, opts ...option.RequestOption) (res *NamespaceSchemaResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.Namespace, precfg.DefaultNamespace)
	if query.Namespace.Value == "" {
		err = errors.New("missing required namespace parameter")
		return
	}
	path := fmt.Sprintf("v1/namespaces/%s/schema", url.PathEscape(query.Namespace.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update namespace schema.
func (r *NamespaceService) UpdateSchema(ctx context.Context, params NamespaceUpdateSchemaParams, opts ...option.RequestOption) (res *NamespaceUpdateSchemaResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.Namespace, precfg.DefaultNamespace)
	if params.Namespace.Value == "" {
		err = errors.New("missing required namespace parameter")
		return
	}
	path := fmt.Sprintf("v1/namespaces/%s/schema", url.PathEscape(params.Namespace.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Create, update, or delete documents.
func (r *NamespaceService) Write(ctx context.Context, params NamespaceWriteParams, opts ...option.RequestOption) (res *NamespaceWriteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.Namespace, precfg.DefaultNamespace)
	if params.Namespace.Value == "" {
		err = errors.New("missing required namespace parameter")
		return
	}
	path := fmt.Sprintf("v2/namespaces/%s", url.PathEscape(params.Namespace.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AggregationGroup map[string]any

// Detailed configuration for an attribute attached to a document.
type AttributeSchemaConfig struct {
	// Whether to create an approximate nearest neighbor index for the attribute.
	Ann bool `json:"ann"`
	// Whether or not the attributes can be used in filters.
	Filterable bool `json:"filterable"`
	// Whether this attribute can be used as part of a BM25 full-text search. Requires
	// the `string` or `[]string` type, and by default, BM25-enabled attributes are not
	// filterable. You can override this by setting `filterable: true`.
	FullTextSearch FullTextSearchConfig `json:"full_text_search"`
	// Whether to enable Regex filters on this attribute.
	Regex bool `json:"regex"`
	// The data type of the attribute. Valid values: string, int, uint, uuid, datetime,
	// bool, []string, []int, []uint, []uuid, []datetime, [DIMS]f16, [DIMS]f32.
	Type AttributeType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ann            respjson.Field
		Filterable     respjson.Field
		FullTextSearch respjson.Field
		Regex          respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributeSchemaConfig) RawJSON() string { return r.JSON.raw }
func (r *AttributeSchemaConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AttributeSchemaConfig to a AttributeSchemaConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AttributeSchemaConfigParam.Overrides()
func (r AttributeSchemaConfig) ToParam() AttributeSchemaConfigParam {
	return param.Override[AttributeSchemaConfigParam](json.RawMessage(r.RawJSON()))
}

// Detailed configuration for an attribute attached to a document.
type AttributeSchemaConfigParam struct {
	// Whether to create an approximate nearest neighbor index for the attribute.
	Ann param.Opt[bool] `json:"ann,omitzero"`
	// Whether or not the attributes can be used in filters.
	Filterable param.Opt[bool] `json:"filterable,omitzero"`
	// Whether to enable Regex filters on this attribute.
	Regex param.Opt[bool] `json:"regex,omitzero"`
	// The data type of the attribute. Valid values: string, int, uint, uuid, datetime,
	// bool, []string, []int, []uint, []uuid, []datetime, [DIMS]f16, [DIMS]f32.
	Type param.Opt[AttributeType] `json:"type,omitzero"`
	// Whether this attribute can be used as part of a BM25 full-text search. Requires
	// the `string` or `[]string` type, and by default, BM25-enabled attributes are not
	// filterable. You can override this by setting `filterable: true`.
	FullTextSearch FullTextSearchConfigParam `json:"full_text_search,omitzero"`
	paramObj
}

func (r AttributeSchemaConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow AttributeSchemaConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributeSchemaConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributeType = string

// Additional (optional) parameters for a single BM25 query clause.
type Bm25ClauseParams struct {
	// Whether to treat the last token in the query input as a literal prefix.
	LastAsPrefix param.Opt[bool] `json:"last_as_prefix,omitzero"`
	paramObj
}

func (r Bm25ClauseParams) MarshalJSON() (data []byte, err error) {
	type shadow Bm25ClauseParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Bm25ClauseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A list of documents in columnar format. Each key is a column name, mapped to an
// array of values for that column.
//
// The property ID is required.
type ColumnsParam struct {
	// The IDs of the documents.
	ID []IDParam `json:"id,omitzero,required" format:"uuid"`
	// The vector embeddings of the documents.
	Vector      ColumnsVectorParam `json:"vector,omitzero"`
	ExtraFields map[string][]any   `json:"-"`
	paramObj
}

func (r ColumnsParam) MarshalJSON() (data []byte, err error) {
	type shadow ColumnsParam
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ColumnsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ColumnsVectorParam struct {
	VectorArray []VectorParam     `json:",omitzero,inline"`
	FloatArray  []float64         `json:",omitzero,inline"`
	String      param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ColumnsVectorParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.VectorArray, u.FloatArray, u.String)
}
func (u *ColumnsVectorParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ColumnsVectorParam) asAny() any {
	if !param.IsOmitted(u.VectorArray) {
		return &u.VectorArray
	} else if !param.IsOmitted(u.FloatArray) {
		return &u.FloatArray
	} else if !param.IsOmitted(u.String) {
		return &u.String.Value
	}
	return nil
}

// Additional (optional) parameters for the ContainsAllTokens filter.
type ContainsAllTokensFilterParams struct {
	// Whether to treat the last token in the query input as a literal prefix.
	LastAsPrefix param.Opt[bool] `json:"last_as_prefix,omitzero"`
	paramObj
}

func (r ContainsAllTokensFilterParams) MarshalJSON() (data []byte, err error) {
	type shadow ContainsAllTokensFilterParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContainsAllTokensFilterParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A function used to calculate vector similarity.
type DistanceMetric string

const (
	DistanceMetricCosineDistance   DistanceMetric = "cosine_distance"
	DistanceMetricEuclideanSquared DistanceMetric = "euclidean_squared"
)

// Configuration options for full-text search.
type FullTextSearchConfig struct {
	// The `b` document length normalization parameter for BM25. Defaults to `0.75`.
	B float64 `json:"b"`
	// Whether searching is case-sensitive. Defaults to `false` (i.e.
	// case-insensitive).
	CaseSensitive bool `json:"case_sensitive"`
	// The `k1` term saturation parameter for BM25. Defaults to `1.2`.
	K1 float64 `json:"k1"`
	// Describes the language of a text attribute. Defaults to `english`.
	//
	// Any of "arabic", "danish", "dutch", "english", "finnish", "french", "german",
	// "greek", "hungarian", "italian", "norwegian", "portuguese", "romanian",
	// "russian", "spanish", "swedish", "tamil", "turkish".
	Language Language `json:"language"`
	// Maximum length of a token in bytes. Tokens larger than this value during
	// tokenization will be filtered out. Has to be between `1` and `254` (inclusive).
	// Defaults to `39`.
	MaxTokenLength int64 `json:"max_token_length"`
	// Removes common words from the text based on language. Defaults to `true` (i.e.
	// remove common words).
	RemoveStopwords bool `json:"remove_stopwords"`
	// Language-specific stemming for the text. Defaults to `false` (i.e., do not
	// stem).
	Stemming bool `json:"stemming"`
	// The tokenizer to use for full-text search on an attribute. Defaults to
	// `word_v2`.
	//
	// Any of "pre_tokenized_array", "word_v0", "word_v1", "word_v2".
	Tokenizer Tokenizer `json:"tokenizer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		B               respjson.Field
		CaseSensitive   respjson.Field
		K1              respjson.Field
		Language        respjson.Field
		MaxTokenLength  respjson.Field
		RemoveStopwords respjson.Field
		Stemming        respjson.Field
		Tokenizer       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FullTextSearchConfig) RawJSON() string { return r.JSON.raw }
func (r *FullTextSearchConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FullTextSearchConfig to a FullTextSearchConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FullTextSearchConfigParam.Overrides()
func (r FullTextSearchConfig) ToParam() FullTextSearchConfigParam {
	return param.Override[FullTextSearchConfigParam](json.RawMessage(r.RawJSON()))
}

// Configuration options for full-text search.
type FullTextSearchConfigParam struct {
	// The `b` document length normalization parameter for BM25. Defaults to `0.75`.
	B param.Opt[float64] `json:"b,omitzero"`
	// Whether searching is case-sensitive. Defaults to `false` (i.e.
	// case-insensitive).
	CaseSensitive param.Opt[bool] `json:"case_sensitive,omitzero"`
	// The `k1` term saturation parameter for BM25. Defaults to `1.2`.
	K1 param.Opt[float64] `json:"k1,omitzero"`
	// Maximum length of a token in bytes. Tokens larger than this value during
	// tokenization will be filtered out. Has to be between `1` and `254` (inclusive).
	// Defaults to `39`.
	MaxTokenLength param.Opt[int64] `json:"max_token_length,omitzero"`
	// Removes common words from the text based on language. Defaults to `true` (i.e.
	// remove common words).
	RemoveStopwords param.Opt[bool] `json:"remove_stopwords,omitzero"`
	// Language-specific stemming for the text. Defaults to `false` (i.e., do not
	// stem).
	Stemming param.Opt[bool] `json:"stemming,omitzero"`
	// Describes the language of a text attribute. Defaults to `english`.
	//
	// Any of "arabic", "danish", "dutch", "english", "finnish", "french", "german",
	// "greek", "hungarian", "italian", "norwegian", "portuguese", "romanian",
	// "russian", "spanish", "swedish", "tamil", "turkish".
	Language Language `json:"language,omitzero"`
	// The tokenizer to use for full-text search on an attribute. Defaults to
	// `word_v2`.
	//
	// Any of "pre_tokenized_array", "word_v0", "word_v1", "word_v2".
	Tokenizer Tokenizer `json:"tokenizer,omitzero"`
	paramObj
}

func (r FullTextSearchConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow FullTextSearchConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FullTextSearchConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ID contains all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: String Int]
type ID struct {
	// This field will be present if the value is a [string] instead of an object.
	String string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	Int  int64 `json:",inline"`
	JSON struct {
		String respjson.Field
		Int    respjson.Field
		raw    string
	} `json:"-"`
}

func (u ID) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ID) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ID) RawJSON() string { return u.JSON.raw }

func (r *ID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ID to a IDParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// IDParam.Overrides()
func (r ID) ToParam() IDParam {
	return param.Override[IDParam](json.RawMessage(r.RawJSON()))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IDParam struct {
	String param.Opt[string] `json:",omitzero,inline"`
	Int    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u IDParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.String, u.Int)
}
func (u *IDParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *IDParam) asAny() any {
	if !param.IsOmitted(u.String) {
		return &u.String.Value
	} else if !param.IsOmitted(u.Int) {
		return &u.Int.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IncludeAttributesParam struct {
	Bool        param.Opt[bool] `json:",omitzero,inline"`
	StringArray []string        `json:",omitzero,inline"`
	paramUnion
}

func (u IncludeAttributesParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.Bool, u.StringArray)
}
func (u *IncludeAttributesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *IncludeAttributesParam) asAny() any {
	if !param.IsOmitted(u.Bool) {
		return &u.Bool.Value
	} else if !param.IsOmitted(u.StringArray) {
		return &u.StringArray
	}
	return nil
}

// Describes the language of a text attribute. Defaults to `english`.
type Language string

const (
	LanguageArabic     Language = "arabic"
	LanguageDanish     Language = "danish"
	LanguageDutch      Language = "dutch"
	LanguageEnglish    Language = "english"
	LanguageFinnish    Language = "finnish"
	LanguageFrench     Language = "french"
	LanguageGerman     Language = "german"
	LanguageGreek      Language = "greek"
	LanguageHungarian  Language = "hungarian"
	LanguageItalian    Language = "italian"
	LanguageNorwegian  Language = "norwegian"
	LanguagePortuguese Language = "portuguese"
	LanguageRomanian   Language = "romanian"
	LanguageRussian    Language = "russian"
	LanguageSpanish    Language = "spanish"
	LanguageSwedish    Language = "swedish"
	LanguageTamil      Language = "tamil"
	LanguageTurkish    Language = "turkish"
)

// Metadata about a namespace.
type NamespaceMetadata struct {
	// The approximate number of logical bytes in the namespace.
	ApproxLogicalBytes int64 `json:"approx_logical_bytes,required"`
	// The approximate number of rows in the namespace.
	ApproxRowCount int64 `json:"approx_row_count,required"`
	// The timestamp when the namespace was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The schema of the namespace.
	Schema map[string]AttributeSchemaConfig `json:"schema,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApproxLogicalBytes respjson.Field
		ApproxRowCount     respjson.Field
		CreatedAt          respjson.Field
		Schema             respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NamespaceMetadata) RawJSON() string { return r.JSON.raw }
func (r *NamespaceMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Query, filter, full-text search and vector search documents.
type QueryParam struct {
	// The number of results to return.
	TopK param.Opt[int64] `json:"top_k,omitzero"`
	// Aggregations to compute over all documents in the namespace that match the
	// filters.
	AggregateBy map[string]any `json:"aggregate_by,omitzero"`
	// A function used to calculate vector similarity.
	//
	// Any of "cosine_distance", "euclidean_squared".
	DistanceMetric DistanceMetric `json:"distance_metric,omitzero"`
	// List of attribute names to exclude from the response. All other attributes will
	// be included in the response.
	ExcludeAttributes []string `json:"exclude_attributes,omitzero"`
	// Exact filters for attributes to refine search results for. Think of it as a SQL
	// WHERE clause.
	Filters any `json:"filters,omitzero"`
	// Groups documents by the specified attributes (the "group key") before computing
	// aggregates. Aggregates are computed separately for each group.
	GroupBy []string `json:"group_by,omitzero"`
	// Whether to include attributes in the response.
	IncludeAttributes IncludeAttributesParam `json:"include_attributes,omitzero"`
	// How to rank the documents in the namespace.
	RankBy any `json:"rank_by,omitzero"`
	paramObj
}

func (r QueryParam) MarshalJSON() (data []byte, err error) {
	type shadow QueryParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The billing information for a query.
type QueryBilling struct {
	// The number of billable logical bytes queried from the namespace.
	BillableLogicalBytesQueried int64 `json:"billable_logical_bytes_queried,required"`
	// The number of billable logical bytes returned from the query.
	BillableLogicalBytesReturned int64 `json:"billable_logical_bytes_returned,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillableLogicalBytesQueried  respjson.Field
		BillableLogicalBytesReturned respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryBilling) RawJSON() string { return r.JSON.raw }
func (r *QueryBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The performance information for a query.
type QueryPerformance struct {
	// the approximate number of documents in the namespace.
	ApproxNamespaceSize int64 `json:"approx_namespace_size,required"`
	// The ratio of cache hits to total cache lookups.
	CacheHitRatio float64 `json:"cache_hit_ratio,required"`
	// A qualitative description of the cache hit ratio (`hot`, `warm`, or `cold`).
	CacheTemperature string `json:"cache_temperature,required"`
	// The number of unindexed documents processed by the query.
	ExhaustiveSearchCount int64 `json:"exhaustive_search_count,required"`
	// Request time measured on the server, excluding time spent waiting due to the
	// namespace concurrency limit.
	QueryExecutionMs int64 `json:"query_execution_ms,required"`
	// Request time measured on the server, including time spent waiting for other
	// queries to complete if the namespace was at its concurrency limit.
	ServerTotalMs int64 `json:"server_total_ms,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApproxNamespaceSize   respjson.Field
		CacheHitRatio         respjson.Field
		CacheTemperature      respjson.Field
		ExhaustiveSearchCount respjson.Field
		QueryExecutionMs      respjson.Field
		ServerTotalMs         respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryPerformance) RawJSON() string { return r.JSON.raw }
func (r *QueryPerformance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single document, in a row-based format.
type Row struct {
	// An identifier for a document.
	ID ID `json:"id,required" format:"uuid"`
	// A vector embedding associated with a document.
	Vector      Vector         `json:"vector"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Vector      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Row) RawJSON() string { return r.JSON.raw }
func (r *Row) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Row to a RowParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RowParam.Overrides()
func (r Row) ToParam() RowParam {
	return param.Override[RowParam](json.RawMessage(r.RawJSON()))
}

// A single document, in a row-based format.
//
// The property ID is required.
type RowParam struct {
	// An identifier for a document.
	ID IDParam `json:"id,omitzero,required" format:"uuid"`
	// A vector embedding associated with a document.
	Vector      VectorParam    `json:"vector,omitzero"`
	ExtraFields map[string]any `json:"-"`
	paramObj
}

func (r RowParam) MarshalJSON() (data []byte, err error) {
	type shadow RowParam
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *RowParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The tokenizer to use for full-text search on an attribute. Defaults to
// `word_v2`.
type Tokenizer string

const (
	TokenizerPreTokenizedArray Tokenizer = "pre_tokenized_array"
	TokenizerWordV0            Tokenizer = "word_v0"
	TokenizerWordV1            Tokenizer = "word_v1"
	TokenizerWordV2            Tokenizer = "word_v2"
)

// Vector contains all possible properties and values from [[]float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: FloatArray String]
type Vector struct {
	// This field will be present if the value is a [[]float64] instead of an object.
	FloatArray []float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	String string `json:",inline"`
	JSON   struct {
		FloatArray respjson.Field
		String     respjson.Field
		raw        string
	} `json:"-"`
}

func (u Vector) AsFloatArray() (v []float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u Vector) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u Vector) RawJSON() string { return u.JSON.raw }

func (r *Vector) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Vector to a VectorParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// VectorParam.Overrides()
func (r Vector) ToParam() VectorParam {
	return param.Override[VectorParam](json.RawMessage(r.RawJSON()))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VectorParam struct {
	FloatArray []float64         `json:",omitzero,inline"`
	String     param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u VectorParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.FloatArray, u.String)
}
func (u *VectorParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorParam) asAny() any {
	if !param.IsOmitted(u.FloatArray) {
		return &u.FloatArray
	} else if !param.IsOmitted(u.String) {
		return &u.String.Value
	}
	return nil
}

// The encoding to use for vectors in the response.
type VectorEncoding string

const (
	VectorEncodingFloat  VectorEncoding = "float"
	VectorEncodingBase64 VectorEncoding = "base64"
)

// The billing information for a write request.
type WriteBilling struct {
	// The number of billable logical bytes written to the namespace.
	BillableLogicalBytesWritten int64 `json:"billable_logical_bytes_written,required"`
	// The billing information for a query.
	Query QueryBilling `json:"query"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillableLogicalBytesWritten respjson.Field
		Query                       respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WriteBilling) RawJSON() string { return r.JSON.raw }
func (r *WriteBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response to a successful namespace deletion request.
type NamespaceDeleteAllResponse struct {
	// The status of the request.
	Status constant.Ok `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NamespaceDeleteAllResponse) RawJSON() string { return r.JSON.raw }
func (r *NamespaceDeleteAllResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response to a successful query explain.
type NamespaceExplainQueryResponse struct {
	// The textual representation of the query plan.
	PlanText string `json:"plan_text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PlanText    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NamespaceExplainQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *NamespaceExplainQueryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response to a successful cache warm request.
type NamespaceHintCacheWarmResponse struct {
	// The status of the request.
	Status  constant.Accepted `json:"status,required"`
	Message string            `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NamespaceHintCacheWarmResponse) RawJSON() string { return r.JSON.raw }
func (r *NamespaceHintCacheWarmResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The result of a multi-query.
type NamespaceMultiQueryResponse struct {
	// The billing information for a query.
	Billing QueryBilling `json:"billing,required"`
	// The performance information for a query.
	Performance QueryPerformance                    `json:"performance,required"`
	Results     []NamespaceMultiQueryResponseResult `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Billing     respjson.Field
		Performance respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NamespaceMultiQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *NamespaceMultiQueryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NamespaceMultiQueryResponseResult struct {
	AggregationGroups []AggregationGroup `json:"aggregation_groups"`
	Aggregations      map[string]any     `json:"aggregations"`
	Rows              []Row              `json:"rows"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AggregationGroups respjson.Field
		Aggregations      respjson.Field
		Rows              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NamespaceMultiQueryResponseResult) RawJSON() string { return r.JSON.raw }
func (r *NamespaceMultiQueryResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The result of a query.
type NamespaceQueryResponse struct {
	// The billing information for a query.
	Billing QueryBilling `json:"billing,required"`
	// The performance information for a query.
	Performance       QueryPerformance   `json:"performance,required"`
	AggregationGroups []AggregationGroup `json:"aggregation_groups"`
	Aggregations      map[string]any     `json:"aggregations"`
	Rows              []Row              `json:"rows"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Billing           respjson.Field
		Performance       respjson.Field
		AggregationGroups respjson.Field
		Aggregations      respjson.Field
		Rows              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NamespaceQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *NamespaceQueryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response to a successful cache warm request.
type NamespaceRecallResponse struct {
	// The average number of documents retrieved by the approximate nearest neighbor
	// searches.
	AvgAnnCount float64 `json:"avg_ann_count,required"`
	// The average number of documents retrieved by the exhaustive searches.
	AvgExhaustiveCount float64 `json:"avg_exhaustive_count,required"`
	// The average recall of the queries.
	AvgRecall float64 `json:"avg_recall,required"`
	// Ground truth data including query vectors and true nearest neighbors. Only
	// included when include_ground_truth is true.
	GroundTruth []NamespaceRecallResponseGroundTruth `json:"ground_truth"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgAnnCount        respjson.Field
		AvgExhaustiveCount respjson.Field
		AvgRecall          respjson.Field
		GroundTruth        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NamespaceRecallResponse) RawJSON() string { return r.JSON.raw }
func (r *NamespaceRecallResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NamespaceRecallResponseGroundTruth struct {
	// The true nearest neighbors with their distances and vectors.
	NearestNeighbors []Row `json:"nearest_neighbors,required"`
	// The query vector used for this search.
	QueryVector []float64 `json:"query_vector,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NearestNeighbors respjson.Field
		QueryVector      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NamespaceRecallResponseGroundTruth) RawJSON() string { return r.JSON.raw }
func (r *NamespaceRecallResponseGroundTruth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NamespaceSchemaResponse map[string]AttributeSchemaConfig

type NamespaceUpdateSchemaResponse map[string]AttributeSchemaConfig

// The response to a successful write request.
type NamespaceWriteResponse struct {
	// The billing information for a write request.
	Billing WriteBilling `json:"billing,required"`
	// A message describing the result of the write request.
	Message string `json:"message,required"`
	// The number of rows affected by the write request.
	RowsAffected int64 `json:"rows_affected,required"`
	// The status of the request.
	Status constant.Ok `json:"status,required"`
	// The number of rows deleted by the write request.
	RowsDeleted int64 `json:"rows_deleted"`
	// The number of rows patched by the write request.
	RowsPatched int64 `json:"rows_patched"`
	// The number of rows upserted by the write request.
	RowsUpserted int64 `json:"rows_upserted"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Billing      respjson.Field
		Message      respjson.Field
		RowsAffected respjson.Field
		Status       respjson.Field
		RowsDeleted  respjson.Field
		RowsPatched  respjson.Field
		RowsUpserted respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NamespaceWriteResponse) RawJSON() string { return r.JSON.raw }
func (r *NamespaceWriteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NamespaceDeleteAllParams struct {
	Namespace param.Opt[string] `path:"namespace,omitzero,required" json:"-"`
	paramObj
}

type NamespaceExplainQueryParams struct {
	Namespace param.Opt[string] `path:"namespace,omitzero,required" json:"-"`
	// The number of results to return.
	TopK param.Opt[int64] `json:"top_k,omitzero"`
	// Aggregations to compute over all documents in the namespace that match the
	// filters.
	AggregateBy map[string]any `json:"aggregate_by,omitzero"`
	// The consistency level for a query.
	Consistency NamespaceExplainQueryParamsConsistency `json:"consistency,omitzero"`
	// A function used to calculate vector similarity.
	//
	// Any of "cosine_distance", "euclidean_squared".
	DistanceMetric DistanceMetric `json:"distance_metric,omitzero"`
	// List of attribute names to exclude from the response. All other attributes will
	// be included in the response.
	ExcludeAttributes []string `json:"exclude_attributes,omitzero"`
	// Exact filters for attributes to refine search results for. Think of it as a SQL
	// WHERE clause.
	Filters any `json:"filters,omitzero"`
	// Groups documents by the specified attributes (the "group key") before computing
	// aggregates. Aggregates are computed separately for each group.
	GroupBy []string `json:"group_by,omitzero"`
	// Whether to include attributes in the response.
	IncludeAttributes IncludeAttributesParam `json:"include_attributes,omitzero"`
	// How to rank the documents in the namespace.
	RankBy any `json:"rank_by,omitzero"`
	// The encoding to use for vectors in the response.
	//
	// Any of "float", "base64".
	VectorEncoding VectorEncoding `json:"vector_encoding,omitzero"`
	paramObj
}

func (r NamespaceExplainQueryParams) MarshalJSON() (data []byte, err error) {
	type shadow NamespaceExplainQueryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NamespaceExplainQueryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The consistency level for a query.
type NamespaceExplainQueryParamsConsistency struct {
	// The query's consistency level.
	//
	// Any of "strong", "eventual".
	Level NamespaceExplainQueryParamsConsistencyLevel `json:"level,omitzero"`
	paramObj
}

func (r NamespaceExplainQueryParamsConsistency) MarshalJSON() (data []byte, err error) {
	type shadow NamespaceExplainQueryParamsConsistency
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NamespaceExplainQueryParamsConsistency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The query's consistency level.
type NamespaceExplainQueryParamsConsistencyLevel string

const (
	NamespaceExplainQueryParamsConsistencyLevelStrong   NamespaceExplainQueryParamsConsistencyLevel = "strong"
	NamespaceExplainQueryParamsConsistencyLevelEventual NamespaceExplainQueryParamsConsistencyLevel = "eventual"
)

type NamespaceHintCacheWarmParams struct {
	Namespace param.Opt[string] `path:"namespace,omitzero,required" json:"-"`
	paramObj
}

type NamespaceMetadataParams struct {
	Namespace param.Opt[string] `path:"namespace,omitzero,required" json:"-"`
	paramObj
}

type NamespaceMultiQueryParams struct {
	Namespace param.Opt[string] `path:"namespace,omitzero,required" json:"-"`
	Queries   []QueryParam      `json:"queries,omitzero,required"`
	// The consistency level for a query.
	Consistency NamespaceMultiQueryParamsConsistency `json:"consistency,omitzero"`
	// The encoding to use for vectors in the response.
	//
	// Any of "float", "base64".
	VectorEncoding VectorEncoding `json:"vector_encoding,omitzero"`
	paramObj
}

func (r NamespaceMultiQueryParams) MarshalJSON() (data []byte, err error) {
	type shadow NamespaceMultiQueryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NamespaceMultiQueryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The consistency level for a query.
type NamespaceMultiQueryParamsConsistency struct {
	// The query's consistency level.
	//
	// Any of "strong", "eventual".
	Level NamespaceMultiQueryParamsConsistencyLevel `json:"level,omitzero"`
	paramObj
}

func (r NamespaceMultiQueryParamsConsistency) MarshalJSON() (data []byte, err error) {
	type shadow NamespaceMultiQueryParamsConsistency
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NamespaceMultiQueryParamsConsistency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The query's consistency level.
type NamespaceMultiQueryParamsConsistencyLevel string

const (
	NamespaceMultiQueryParamsConsistencyLevelStrong   NamespaceMultiQueryParamsConsistencyLevel = "strong"
	NamespaceMultiQueryParamsConsistencyLevelEventual NamespaceMultiQueryParamsConsistencyLevel = "eventual"
)

type NamespaceQueryParams struct {
	Namespace param.Opt[string] `path:"namespace,omitzero,required" json:"-"`
	// The number of results to return.
	TopK param.Opt[int64] `json:"top_k,omitzero"`
	// Aggregations to compute over all documents in the namespace that match the
	// filters.
	AggregateBy map[string]any `json:"aggregate_by,omitzero"`
	// The consistency level for a query.
	Consistency NamespaceQueryParamsConsistency `json:"consistency,omitzero"`
	// A function used to calculate vector similarity.
	//
	// Any of "cosine_distance", "euclidean_squared".
	DistanceMetric DistanceMetric `json:"distance_metric,omitzero"`
	// List of attribute names to exclude from the response. All other attributes will
	// be included in the response.
	ExcludeAttributes []string `json:"exclude_attributes,omitzero"`
	// Exact filters for attributes to refine search results for. Think of it as a SQL
	// WHERE clause.
	Filters any `json:"filters,omitzero"`
	// Groups documents by the specified attributes (the "group key") before computing
	// aggregates. Aggregates are computed separately for each group.
	GroupBy []string `json:"group_by,omitzero"`
	// Whether to include attributes in the response.
	IncludeAttributes IncludeAttributesParam `json:"include_attributes,omitzero"`
	// How to rank the documents in the namespace.
	RankBy any `json:"rank_by,omitzero"`
	// The encoding to use for vectors in the response.
	//
	// Any of "float", "base64".
	VectorEncoding VectorEncoding `json:"vector_encoding,omitzero"`
	paramObj
}

func (r NamespaceQueryParams) MarshalJSON() (data []byte, err error) {
	type shadow NamespaceQueryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NamespaceQueryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The consistency level for a query.
type NamespaceQueryParamsConsistency struct {
	// The query's consistency level.
	//
	// Any of "strong", "eventual".
	Level NamespaceQueryParamsConsistencyLevel `json:"level,omitzero"`
	paramObj
}

func (r NamespaceQueryParamsConsistency) MarshalJSON() (data []byte, err error) {
	type shadow NamespaceQueryParamsConsistency
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NamespaceQueryParamsConsistency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The query's consistency level.
type NamespaceQueryParamsConsistencyLevel string

const (
	NamespaceQueryParamsConsistencyLevelStrong   NamespaceQueryParamsConsistencyLevel = "strong"
	NamespaceQueryParamsConsistencyLevelEventual NamespaceQueryParamsConsistencyLevel = "eventual"
)

type NamespaceRecallParams struct {
	Namespace param.Opt[string] `path:"namespace,omitzero,required" json:"-"`
	// Include ground truth data (query vectors and true nearest neighbors) in the
	// response.
	IncludeGroundTruth param.Opt[bool] `json:"include_ground_truth,omitzero"`
	// The number of searches to run.
	Num param.Opt[int64] `json:"num,omitzero"`
	// Search for `top_k` nearest neighbors.
	TopK param.Opt[int64] `json:"top_k,omitzero"`
	// Filter by attributes. Same syntax as the query endpoint.
	Filters any `json:"filters,omitzero"`
	// Use specific query vectors for the measurement. If omitted, sampled from the
	// index.
	Queries []float64 `json:"queries,omitzero"`
	paramObj
}

func (r NamespaceRecallParams) MarshalJSON() (data []byte, err error) {
	type shadow NamespaceRecallParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NamespaceRecallParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NamespaceSchemaParams struct {
	Namespace param.Opt[string] `path:"namespace,omitzero,required" json:"-"`
	paramObj
}

type NamespaceUpdateSchemaParams struct {
	Namespace param.Opt[string] `path:"namespace,omitzero,required" json:"-"`
	// The desired schema for the namespace.
	Schema map[string]AttributeSchemaConfigParam
	paramObj
}

func (r NamespaceUpdateSchemaParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Schema)
}
func (r *NamespaceUpdateSchemaParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Schema)
}

type NamespaceWriteParams struct {
	Namespace param.Opt[string] `path:"namespace,omitzero,required" json:"-"`
	// The namespace to copy documents from.
	CopyFromNamespace param.Opt[string] `json:"copy_from_namespace,omitzero"`
	// Disables write throttling (HTTP 429 responses) during high-volume ingestion.
	DisableBackpressure param.Opt[bool] `json:"disable_backpressure,omitzero"`
	// The filter specifying which documents to delete.
	DeleteByFilter any `json:"delete_by_filter,omitzero"`
	// A condition evaluated against the current value of each document targeted by a
	// delete write. Only documents that pass the condition are deleted.
	DeleteCondition any       `json:"delete_condition,omitzero"`
	Deletes         []IDParam `json:"deletes,omitzero" format:"uuid"`
	// A function used to calculate vector similarity.
	//
	// Any of "cosine_distance", "euclidean_squared".
	DistanceMetric DistanceMetric `json:"distance_metric,omitzero"`
	// The encryption configuration for a namespace.
	Encryption NamespaceWriteParamsEncryption `json:"encryption,omitzero"`
	// A list of documents in columnar format. Each key is a column name, mapped to an
	// array of values for that column.
	PatchColumns ColumnsParam `json:"patch_columns,omitzero"`
	// A condition evaluated against the current value of each document targeted by a
	// patch write. Only documents that pass the condition are patched.
	PatchCondition any        `json:"patch_condition,omitzero"`
	PatchRows      []RowParam `json:"patch_rows,omitzero"`
	// The schema of the attributes attached to the documents.
	Schema map[string]AttributeSchemaConfigParam `json:"schema,omitzero"`
	// A list of documents in columnar format. Each key is a column name, mapped to an
	// array of values for that column.
	UpsertColumns ColumnsParam `json:"upsert_columns,omitzero"`
	// A condition evaluated against the current value of each document targeted by an
	// upsert write. Only documents that pass the condition are upserted.
	UpsertCondition any        `json:"upsert_condition,omitzero"`
	UpsertRows      []RowParam `json:"upsert_rows,omitzero"`
	paramObj
}

func (r NamespaceWriteParams) MarshalJSON() (data []byte, err error) {
	type shadow NamespaceWriteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NamespaceWriteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The encryption configuration for a namespace.
type NamespaceWriteParamsEncryption struct {
	Cmek NamespaceWriteParamsEncryptionCmek `json:"cmek,omitzero"`
	paramObj
}

func (r NamespaceWriteParamsEncryption) MarshalJSON() (data []byte, err error) {
	type shadow NamespaceWriteParamsEncryption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NamespaceWriteParamsEncryption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property KeyName is required.
type NamespaceWriteParamsEncryptionCmek struct {
	// The identifier of the CMEK key to use for encryption. For GCP, the
	// fully-qualified resource name of the key. For AWS, the ARN of the key.
	KeyName string `json:"key_name,required"`
	paramObj
}

func (r NamespaceWriteParamsEncryptionCmek) MarshalJSON() (data []byte, err error) {
	type shadow NamespaceWriteParamsEncryptionCmek
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NamespaceWriteParamsEncryptionCmek) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
