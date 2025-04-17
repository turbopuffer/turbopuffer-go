// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turbopuffer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/turbopuffer-go/internal/apijson"
	"github.com/stainless-sdks/turbopuffer-go/internal/apiquery"
	"github.com/stainless-sdks/turbopuffer-go/internal/requestconfig"
	"github.com/stainless-sdks/turbopuffer-go/option"
	"github.com/stainless-sdks/turbopuffer-go/packages/pagination"
	"github.com/stainless-sdks/turbopuffer-go/packages/param"
	"github.com/stainless-sdks/turbopuffer-go/packages/resp"
	"github.com/stainless-sdks/turbopuffer-go/shared/constant"
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

// List namespaces.
func (r *NamespaceService) List(ctx context.Context, query NamespaceListParams, opts ...option.RequestOption) (res *pagination.ListNamespaces[NamespaceSummary], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/namespaces"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List namespaces.
func (r *NamespaceService) ListAutoPaging(ctx context.Context, query NamespaceListParams, opts ...option.RequestOption) *pagination.ListNamespacesAutoPager[NamespaceSummary] {
	return pagination.NewListNamespacesAutoPager(r.List(ctx, query, opts...))
}

// Delete namespace.
func (r *NamespaceService) DeleteAll(ctx context.Context, namespace string, opts ...option.RequestOption) (res *NamespaceDeleteAllResponse, err error) {
	opts = append(r.Options[:], opts...)
	if namespace == "" {
		err = errors.New("missing required namespace parameter")
		return
	}
	path := fmt.Sprintf("v1/namespaces/%s", namespace)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get namespace schema.
func (r *NamespaceService) GetSchema(ctx context.Context, namespace string, opts ...option.RequestOption) (res *NamespaceGetSchemaResponse, err error) {
	opts = append(r.Options[:], opts...)
	if namespace == "" {
		err = errors.New("missing required namespace parameter")
		return
	}
	path := fmt.Sprintf("v1/namespaces/%s/schema", namespace)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Query, filter, full-text search and vector search documents.
func (r *NamespaceService) Query(ctx context.Context, namespace string, body NamespaceQueryParams, opts ...option.RequestOption) (res *[]DocumentRowWithScore, err error) {
	opts = append(r.Options[:], opts...)
	if namespace == "" {
		err = errors.New("missing required namespace parameter")
		return
	}
	path := fmt.Sprintf("v1/namespaces/%s/query", namespace)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create, update, or delete documents.
func (r *NamespaceService) Upsert(ctx context.Context, namespace string, body NamespaceUpsertParams, opts ...option.RequestOption) (res *NamespaceUpsertResponse, err error) {
	opts = append(r.Options[:], opts...)
	if namespace == "" {
		err = errors.New("missing required namespace parameter")
		return
	}
	path := fmt.Sprintf("v1/namespaces/%s", namespace)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The schema for an attribute attached to a document.
type AttributeSchema struct {
	// Whether or not the attributes can be used in filters/WHERE clauses.
	Filterable bool `json:"filterable"`
	// Whether this attribute can be used as part of a BM25 full-text search. Requires
	// the `string` or `[]string` type, and by default, BM25-enabled attributes are not
	// filterable. You can override this by setting `filterable: true`.
	FullTextSearch AttributeSchemaFullTextSearchUnion `json:"full_text_search"`
	// The data type of the attribute.
	//
	// Any of "string", "uint", "uuid", "bool", "[]string", "[]uint", "[]uuid".
	Type AttributeSchemaType `json:"type"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Filterable     resp.Field
		FullTextSearch resp.Field
		Type           resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributeSchema) RawJSON() string { return r.JSON.raw }
func (r *AttributeSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AttributeSchema to a AttributeSchemaParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AttributeSchemaParam.IsOverridden()
func (r AttributeSchema) ToParam() AttributeSchemaParam {
	return param.OverrideObj[AttributeSchemaParam](r.RawJSON())
}

// AttributeSchemaFullTextSearchUnion contains all possible properties and values
// from [bool], [FullTextSearchConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool]
type AttributeSchemaFullTextSearchUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field is from variant [FullTextSearchConfig].
	CaseSensitive bool `json:"case_sensitive"`
	// This field is from variant [FullTextSearchConfig].
	Language FullTextSearchConfigLanguage `json:"language"`
	// This field is from variant [FullTextSearchConfig].
	RemoveStopwords bool `json:"remove_stopwords"`
	// This field is from variant [FullTextSearchConfig].
	Stemming bool `json:"stemming"`
	JSON     struct {
		OfBool          resp.Field
		CaseSensitive   resp.Field
		Language        resp.Field
		RemoveStopwords resp.Field
		Stemming        resp.Field
		raw             string
	} `json:"-"`
}

func (u AttributeSchemaFullTextSearchUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributeSchemaFullTextSearchUnion) AsFullTextSearchConfig() (v FullTextSearchConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributeSchemaFullTextSearchUnion) RawJSON() string { return u.JSON.raw }

func (r *AttributeSchemaFullTextSearchUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The language of the text. Defaults to `english`.
type AttributeSchemaFullTextSearchLanguage string

const (
	AttributeSchemaFullTextSearchLanguageArabic     AttributeSchemaFullTextSearchLanguage = "arabic"
	AttributeSchemaFullTextSearchLanguageDanish     AttributeSchemaFullTextSearchLanguage = "danish"
	AttributeSchemaFullTextSearchLanguageDutch      AttributeSchemaFullTextSearchLanguage = "dutch"
	AttributeSchemaFullTextSearchLanguageEnglish    AttributeSchemaFullTextSearchLanguage = "english"
	AttributeSchemaFullTextSearchLanguageFinnish    AttributeSchemaFullTextSearchLanguage = "finnish"
	AttributeSchemaFullTextSearchLanguageFrench     AttributeSchemaFullTextSearchLanguage = "french"
	AttributeSchemaFullTextSearchLanguageGerman     AttributeSchemaFullTextSearchLanguage = "german"
	AttributeSchemaFullTextSearchLanguageGreek      AttributeSchemaFullTextSearchLanguage = "greek"
	AttributeSchemaFullTextSearchLanguageHungarian  AttributeSchemaFullTextSearchLanguage = "hungarian"
	AttributeSchemaFullTextSearchLanguageItalian    AttributeSchemaFullTextSearchLanguage = "italian"
	AttributeSchemaFullTextSearchLanguageNorwegian  AttributeSchemaFullTextSearchLanguage = "norwegian"
	AttributeSchemaFullTextSearchLanguagePortuguese AttributeSchemaFullTextSearchLanguage = "portuguese"
	AttributeSchemaFullTextSearchLanguageRomanian   AttributeSchemaFullTextSearchLanguage = "romanian"
	AttributeSchemaFullTextSearchLanguageRussian    AttributeSchemaFullTextSearchLanguage = "russian"
	AttributeSchemaFullTextSearchLanguageSpanish    AttributeSchemaFullTextSearchLanguage = "spanish"
	AttributeSchemaFullTextSearchLanguageSwedish    AttributeSchemaFullTextSearchLanguage = "swedish"
	AttributeSchemaFullTextSearchLanguageTamil      AttributeSchemaFullTextSearchLanguage = "tamil"
	AttributeSchemaFullTextSearchLanguageTurkish    AttributeSchemaFullTextSearchLanguage = "turkish"
)

// The data type of the attribute.
type AttributeSchemaType string

const (
	AttributeSchemaTypeString      AttributeSchemaType = "string"
	AttributeSchemaTypeUint        AttributeSchemaType = "uint"
	AttributeSchemaTypeUuid        AttributeSchemaType = "uuid"
	AttributeSchemaTypeBool        AttributeSchemaType = "bool"
	AttributeSchemaTypeStringArray AttributeSchemaType = "[]string"
	AttributeSchemaTypeUintArray   AttributeSchemaType = "[]uint"
	AttributeSchemaTypeUuidArray   AttributeSchemaType = "[]uuid"
)

// The schema for an attribute attached to a document.
type AttributeSchemaParam struct {
	// Whether or not the attributes can be used in filters/WHERE clauses.
	Filterable param.Opt[bool] `json:"filterable,omitzero"`
	// Whether this attribute can be used as part of a BM25 full-text search. Requires
	// the `string` or `[]string` type, and by default, BM25-enabled attributes are not
	// filterable. You can override this by setting `filterable: true`.
	FullTextSearch AttributeSchemaFullTextSearchUnionParam `json:"full_text_search,omitzero"`
	// The data type of the attribute.
	//
	// Any of "string", "uint", "uuid", "bool", "[]string", "[]uint", "[]uuid".
	Type AttributeSchemaType `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AttributeSchemaParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r AttributeSchemaParam) MarshalJSON() (data []byte, err error) {
	type shadow AttributeSchemaParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AttributeSchemaFullTextSearchUnionParam struct {
	OfBool                 param.Opt[bool]            `json:",omitzero,inline"`
	OfFullTextSearchConfig *FullTextSearchConfigParam `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u AttributeSchemaFullTextSearchUnionParam) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u AttributeSchemaFullTextSearchUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[AttributeSchemaFullTextSearchUnionParam](u.OfBool, u.OfFullTextSearchConfig)
}

func (u *AttributeSchemaFullTextSearchUnionParam) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfFullTextSearchConfig) {
		return u.OfFullTextSearchConfig
	}
	return nil
}

// A function used to calculate vector similarity.
type DistanceMetric string

const (
	DistanceMetricCosineDistance   DistanceMetric = "cosine_distance"
	DistanceMetricEuclideanSquared DistanceMetric = "euclidean_squared"
)

// A list of documents in columnar format.
type DocumentColumnsParam struct {
	// The attributes attached to each of the documents.
	Attributes map[string][]map[string]any `json:"attributes,omitzero"`
	// The IDs of the documents.
	IDs []IDUnionParam `json:"ids,omitzero" format:"uuid"`
	// Vectors describing each of the documents.
	Vectors [][]float64 `json:"vectors,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f DocumentColumnsParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r DocumentColumnsParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentColumnsParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// A single document, in a row-based format.
type DocumentRow struct {
	// An identifier for a document.
	ID IDUnion `json:"id" format:"uuid"`
	// The attributes attached to the document.
	Attributes map[string]any `json:"attributes"`
	// A vector describing the document.
	Vector []float64 `json:"vector,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		Attributes  resp.Field
		Vector      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentRow) RawJSON() string { return r.JSON.raw }
func (r *DocumentRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DocumentRow to a DocumentRowParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DocumentRowParam.IsOverridden()
func (r DocumentRow) ToParam() DocumentRowParam {
	return param.OverrideObj[DocumentRowParam](r.RawJSON())
}

// A single document, in a row-based format.
type DocumentRowParam struct {
	// A vector describing the document.
	Vector []float64 `json:"vector,omitzero"`
	// An identifier for a document.
	ID IDUnionParam `json:"id,omitzero" format:"uuid"`
	// The attributes attached to the document.
	Attributes map[string]any `json:"attributes,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f DocumentRowParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r DocumentRowParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentRowParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// A single document, in a row-based format.
type DocumentRowWithScore struct {
	// For vector search, the distance between the query vector and the document
	// vector. For BM25 full-text search, the score of the document. Not present for
	// other types of queries.
	Dist float64 `json:"dist"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Dist        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
	DocumentRow
}

// Returns the unmodified JSON received from the API
func (r DocumentRowWithScore) RawJSON() string { return r.JSON.raw }
func (r *DocumentRowWithScore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed configuration options for BM25 full-text search.
type FullTextSearchConfig struct {
	// Whether searching is case-sensitive. Defaults to `false` (i.e.
	// case-insensitive).
	CaseSensitive bool `json:"case_sensitive"`
	// The language of the text. Defaults to `english`.
	//
	// Any of "arabic", "danish", "dutch", "english", "finnish", "french", "german",
	// "greek", "hungarian", "italian", "norwegian", "portuguese", "romanian",
	// "russian", "spanish", "swedish", "tamil", "turkish".
	Language FullTextSearchConfigLanguage `json:"language"`
	// Removes common words from the text based on language. Defaults to `true` (i.e.
	// remove common words).
	RemoveStopwords bool `json:"remove_stopwords"`
	// Language-specific stemming for the text. Defaults to `false` (i.e., do not
	// stem).
	Stemming bool `json:"stemming"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CaseSensitive   resp.Field
		Language        resp.Field
		RemoveStopwords resp.Field
		Stemming        resp.Field
		ExtraFields     map[string]resp.Field
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
// FullTextSearchConfigParam.IsOverridden()
func (r FullTextSearchConfig) ToParam() FullTextSearchConfigParam {
	return param.OverrideObj[FullTextSearchConfigParam](r.RawJSON())
}

// The language of the text. Defaults to `english`.
type FullTextSearchConfigLanguage string

const (
	FullTextSearchConfigLanguageArabic     FullTextSearchConfigLanguage = "arabic"
	FullTextSearchConfigLanguageDanish     FullTextSearchConfigLanguage = "danish"
	FullTextSearchConfigLanguageDutch      FullTextSearchConfigLanguage = "dutch"
	FullTextSearchConfigLanguageEnglish    FullTextSearchConfigLanguage = "english"
	FullTextSearchConfigLanguageFinnish    FullTextSearchConfigLanguage = "finnish"
	FullTextSearchConfigLanguageFrench     FullTextSearchConfigLanguage = "french"
	FullTextSearchConfigLanguageGerman     FullTextSearchConfigLanguage = "german"
	FullTextSearchConfigLanguageGreek      FullTextSearchConfigLanguage = "greek"
	FullTextSearchConfigLanguageHungarian  FullTextSearchConfigLanguage = "hungarian"
	FullTextSearchConfigLanguageItalian    FullTextSearchConfigLanguage = "italian"
	FullTextSearchConfigLanguageNorwegian  FullTextSearchConfigLanguage = "norwegian"
	FullTextSearchConfigLanguagePortuguese FullTextSearchConfigLanguage = "portuguese"
	FullTextSearchConfigLanguageRomanian   FullTextSearchConfigLanguage = "romanian"
	FullTextSearchConfigLanguageRussian    FullTextSearchConfigLanguage = "russian"
	FullTextSearchConfigLanguageSpanish    FullTextSearchConfigLanguage = "spanish"
	FullTextSearchConfigLanguageSwedish    FullTextSearchConfigLanguage = "swedish"
	FullTextSearchConfigLanguageTamil      FullTextSearchConfigLanguage = "tamil"
	FullTextSearchConfigLanguageTurkish    FullTextSearchConfigLanguage = "turkish"
)

// Detailed configuration options for BM25 full-text search.
type FullTextSearchConfigParam struct {
	// Whether searching is case-sensitive. Defaults to `false` (i.e.
	// case-insensitive).
	CaseSensitive param.Opt[bool] `json:"case_sensitive,omitzero"`
	// Removes common words from the text based on language. Defaults to `true` (i.e.
	// remove common words).
	RemoveStopwords param.Opt[bool] `json:"remove_stopwords,omitzero"`
	// Language-specific stemming for the text. Defaults to `false` (i.e., do not
	// stem).
	Stemming param.Opt[bool] `json:"stemming,omitzero"`
	// The language of the text. Defaults to `english`.
	//
	// Any of "arabic", "danish", "dutch", "english", "finnish", "french", "german",
	// "greek", "hungarian", "italian", "norwegian", "portuguese", "romanian",
	// "russian", "spanish", "swedish", "tamil", "turkish".
	Language FullTextSearchConfigLanguage `json:"language,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FullTextSearchConfigParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r FullTextSearchConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow FullTextSearchConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// IDUnion contains all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type IDUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString resp.Field
		OfInt    resp.Field
		raw      string
	} `json:"-"`
}

func (u IDUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u IDUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u IDUnion) RawJSON() string { return u.JSON.raw }

func (r *IDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this IDUnion to a IDUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// IDUnionParam.IsOverridden()
func (r IDUnion) ToParam() IDUnionParam {
	return param.OverrideObj[IDUnionParam](r.RawJSON())
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IDUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u IDUnionParam) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }
func (u IDUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[IDUnionParam](u.OfString, u.OfInt)
}

func (u *IDUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// A summary of a namespace.
type NamespaceSummary struct {
	// The namespace ID.
	ID string `json:"id,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NamespaceSummary) RawJSON() string { return r.JSON.raw }
func (r *NamespaceSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response to a successful namespace deletion request.
type NamespaceDeleteAllResponse struct {
	// The status of the request.
	Status constant.Ok `json:"status,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Status      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NamespaceDeleteAllResponse) RawJSON() string { return r.JSON.raw }
func (r *NamespaceDeleteAllResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NamespaceGetSchemaResponse map[string][]AttributeSchema

// The response to a successful upsert request.
type NamespaceUpsertResponse struct {
	// The status of the request.
	Status constant.Ok `json:"status,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Status      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NamespaceUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *NamespaceUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NamespaceListParams struct {
	// Retrieve the next page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Limit the number of results per page.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Retrieve only the namespaces that match the prefix.
	Prefix param.Opt[string] `query:"prefix,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NamespaceListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [NamespaceListParams]'s query parameters as `url.Values`.
func (r NamespaceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NamespaceQueryParams struct {
	// Whether to return vectors for the search results. Vectors are large and slow to
	// deserialize on the client, so use this option only if you need them.
	IncludeVectors param.Opt[bool] `json:"include_vectors,omitzero"`
	// The number of results to return.
	TopK param.Opt[int64] `json:"top_k,omitzero"`
	// The consistency level for a query.
	Consistency NamespaceQueryParamsConsistency `json:"consistency,omitzero"`
	// A function used to calculate vector similarity.
	//
	// Any of "cosine_distance", "euclidean_squared".
	DistanceMetric DistanceMetric `json:"distance_metric,omitzero"`
	// Exact filters for attributes to refine search results for. Think of it as a SQL
	// WHERE clause.
	Filters any `json:"filters,omitzero"`
	// Whether to include attributes in the response.
	IncludeAttributes NamespaceQueryParamsIncludeAttributesUnion `json:"include_attributes,omitzero"`
	// The attribute to rank the results by. Cannot be specified with `vector`.
	RankBy any `json:"rank_by,omitzero"`
	// A vector to search for. It must have the same number of dimensions as the
	// vectors in the namespace. Cannot be specified with `rank_by`.
	Vector []float64 `json:"vector,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NamespaceQueryParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r NamespaceQueryParams) MarshalJSON() (data []byte, err error) {
	type shadow NamespaceQueryParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// The consistency level for a query.
type NamespaceQueryParamsConsistency struct {
	// The query's consistency level.
	//
	// Any of "strong", "eventual".
	Level NamespaceQueryParamsConsistencyLevel `json:"level,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NamespaceQueryParamsConsistency) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r NamespaceQueryParamsConsistency) MarshalJSON() (data []byte, err error) {
	type shadow NamespaceQueryParamsConsistency
	return param.MarshalObject(r, (*shadow)(&r))
}

// The query's consistency level.
type NamespaceQueryParamsConsistencyLevel string

const (
	NamespaceQueryParamsConsistencyLevelStrong   NamespaceQueryParamsConsistencyLevel = "strong"
	NamespaceQueryParamsConsistencyLevelEventual NamespaceQueryParamsConsistencyLevel = "eventual"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type NamespaceQueryParamsIncludeAttributesUnion struct {
	OfBool                                  param.Opt[bool] `json:",omitzero,inline"`
	OfNamespaceQuerysIncludeAttributesArray []string        `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u NamespaceQueryParamsIncludeAttributesUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u NamespaceQueryParamsIncludeAttributesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[NamespaceQueryParamsIncludeAttributesUnion](u.OfBool, u.OfNamespaceQuerysIncludeAttributesArray)
}

func (u *NamespaceQueryParamsIncludeAttributesUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfNamespaceQuerysIncludeAttributesArray) {
		return &u.OfNamespaceQuerysIncludeAttributesArray
	}
	return nil
}

type NamespaceUpsertParams struct {
	// Upsert documents in columnar format.
	Documents NamespaceUpsertParamsDocumentsUnion
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NamespaceUpsertParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r NamespaceUpsertParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Documents)
}

// Upsert documents in columnar format.
//
// Satisfied by [NamespaceUpsertParamsDocumentsUpsertColumnar],
// [NamespaceUpsertParamsDocumentsUpsertRowBased],
// [NamespaceUpsertParamsDocumentsCopyFromNamespace] and
// [NamespaceUpsertParamsDocumentsDeleteByFilter]
type NamespaceUpsertParamsDocumentsUnion interface {
	implNamespaceUpsertParamsDocumentsUnion()
}

func (NamespaceUpsertParamsDocumentsUpsertColumnar) implNamespaceUpsertParamsDocumentsUnion()    {}
func (NamespaceUpsertParamsDocumentsUpsertRowBased) implNamespaceUpsertParamsDocumentsUnion()    {}
func (NamespaceUpsertParamsDocumentsCopyFromNamespace) implNamespaceUpsertParamsDocumentsUnion() {}
func (NamespaceUpsertParamsDocumentsDeleteByFilter) implNamespaceUpsertParamsDocumentsUnion()    {}

// Upsert documents in columnar format.
type NamespaceUpsertParamsDocumentsUpsertColumnar struct {
	// A function used to calculate vector similarity.
	DistanceMetric DistanceMetric `json:"distance_metric,omitzero,required"`
	// The schema of the attributes attached to the documents.
	Schema map[string][]AttributeSchemaParam `json:"schema,omitzero"`
	DocumentColumnsParam
}

func (r NamespaceUpsertParamsDocumentsUpsertColumnar) MarshalJSON() (data []byte, err error) {
	type shadow NamespaceUpsertParamsDocumentsUpsertColumnar
	return param.MarshalObject(r, (*shadow)(&r))
}

// Upsert documents in row-based format.
//
// The properties DistanceMetric, Upserts are required.
type NamespaceUpsertParamsDocumentsUpsertRowBased struct {
	// A function used to calculate vector similarity.
	//
	// Any of "cosine_distance", "euclidean_squared".
	DistanceMetric DistanceMetric     `json:"distance_metric,omitzero,required"`
	Upserts        []DocumentRowParam `json:"upserts,omitzero,required"`
	// The schema of the attributes attached to the documents.
	Schema map[string][]AttributeSchemaParam `json:"schema,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NamespaceUpsertParamsDocumentsUpsertRowBased) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r NamespaceUpsertParamsDocumentsUpsertRowBased) MarshalJSON() (data []byte, err error) {
	type shadow NamespaceUpsertParamsDocumentsUpsertRowBased
	return param.MarshalObject(r, (*shadow)(&r))
}

// Copy documents from another namespace.
//
// The property CopyFromNamespace is required.
type NamespaceUpsertParamsDocumentsCopyFromNamespace struct {
	// The namespace to copy documents from.
	CopyFromNamespace string `json:"copy_from_namespace,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NamespaceUpsertParamsDocumentsCopyFromNamespace) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r NamespaceUpsertParamsDocumentsCopyFromNamespace) MarshalJSON() (data []byte, err error) {
	type shadow NamespaceUpsertParamsDocumentsCopyFromNamespace
	return param.MarshalObject(r, (*shadow)(&r))
}

// Delete documents by filter.
//
// The property DeleteByFilter is required.
type NamespaceUpsertParamsDocumentsDeleteByFilter struct {
	// The filter specifying which documents to delete.
	DeleteByFilter any `json:"delete_by_filter,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NamespaceUpsertParamsDocumentsDeleteByFilter) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r NamespaceUpsertParamsDocumentsDeleteByFilter) MarshalJSON() (data []byte, err error) {
	type shadow NamespaceUpsertParamsDocumentsDeleteByFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
