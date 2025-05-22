// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turbopuffer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/turbopuffer/turbopuffer-go/internal/apijson"
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
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.Namespace, precfg.DefaultNamespace)
	if body.Namespace.Value == "" {
		err = errors.New("missing required namespace parameter")
		return
	}
	path := fmt.Sprintf("v2/namespaces/%s", body.Namespace.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get namespace schema.
func (r *NamespaceService) GetSchema(ctx context.Context, query NamespaceGetSchemaParams, opts ...option.RequestOption) (res *NamespaceGetSchemaResponse, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.Namespace, precfg.DefaultNamespace)
	if query.Namespace.Value == "" {
		err = errors.New("missing required namespace parameter")
		return
	}
	path := fmt.Sprintf("v1/namespaces/%s/schema", query.Namespace.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Query, filter, full-text search and vector search documents.
func (r *NamespaceService) Query(ctx context.Context, params NamespaceQueryParams, opts ...option.RequestOption) (res *NamespaceQueryResponse, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.Namespace, precfg.DefaultNamespace)
	if params.Namespace.Value == "" {
		err = errors.New("missing required namespace parameter")
		return
	}
	path := fmt.Sprintf("v2/namespaces/%s/query", params.Namespace.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Evaluate recall.
func (r *NamespaceService) Recall(ctx context.Context, query NamespaceRecallParams, opts ...option.RequestOption) (res *NamespaceRecallResponse, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.Namespace, precfg.DefaultNamespace)
	if query.Namespace.Value == "" {
		err = errors.New("missing required namespace parameter")
		return
	}
	path := fmt.Sprintf("v1/namespaces/%s/_debug/recall", query.Namespace.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update namespace schema.
func (r *NamespaceService) UpdateSchema(ctx context.Context, params NamespaceUpdateSchemaParams, opts ...option.RequestOption) (res *NamespaceUpdateSchemaResponse, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.Namespace, precfg.DefaultNamespace)
	if params.Namespace.Value == "" {
		err = errors.New("missing required namespace parameter")
		return
	}
	path := fmt.Sprintf("v1/namespaces/%s/schema", params.Namespace.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Warm the cache for a namespace.
func (r *NamespaceService) WarmCache(ctx context.Context, query NamespaceWarmCacheParams, opts ...option.RequestOption) (res *NamespaceWarmCacheResponse, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.Namespace, precfg.DefaultNamespace)
	if query.Namespace.Value == "" {
		err = errors.New("missing required namespace parameter")
		return
	}
	path := fmt.Sprintf("v1/namespaces/%s/hint_cache_warm", query.Namespace.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Create, update, or delete documents.
func (r *NamespaceService) Write(ctx context.Context, params NamespaceWriteParams, opts ...option.RequestOption) (res *NamespaceWriteResponse, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.Namespace, precfg.DefaultNamespace)
	if params.Namespace.Value == "" {
		err = errors.New("missing required namespace parameter")
		return
	}
	path := fmt.Sprintf("v2/namespaces/%s", params.Namespace.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
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
	// Any of "string", "uint", "uuid", "bool", "datetime", "[]string", "[]uint",
	// "[]uuid", "[]datetime".
	Type AttributeSchemaType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filterable     respjson.Field
		FullTextSearch respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
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
// AttributeSchemaParam.Overrides()
func (r AttributeSchema) ToParam() AttributeSchemaParam {
	return param.Override[AttributeSchemaParam](r.RawJSON())
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
		OfBool          respjson.Field
		CaseSensitive   respjson.Field
		Language        respjson.Field
		RemoveStopwords respjson.Field
		Stemming        respjson.Field
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
	AttributeSchemaTypeString        AttributeSchemaType = "string"
	AttributeSchemaTypeUint          AttributeSchemaType = "uint"
	AttributeSchemaTypeUuid          AttributeSchemaType = "uuid"
	AttributeSchemaTypeBool          AttributeSchemaType = "bool"
	AttributeSchemaTypeDatetime      AttributeSchemaType = "datetime"
	AttributeSchemaTypeStringArray   AttributeSchemaType = "[]string"
	AttributeSchemaTypeUintArray     AttributeSchemaType = "[]uint"
	AttributeSchemaTypeUuidArray     AttributeSchemaType = "[]uuid"
	AttributeSchemaTypeDatetimeArray AttributeSchemaType = "[]datetime"
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
	// Any of "string", "uint", "uuid", "bool", "datetime", "[]string", "[]uint",
	// "[]uuid", "[]datetime".
	Type AttributeSchemaType `json:"type,omitzero"`
	paramObj
}

func (r AttributeSchemaParam) MarshalJSON() (data []byte, err error) {
	type shadow AttributeSchemaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributeSchemaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AttributeSchemaFullTextSearchUnionParam struct {
	OfBool                 param.Opt[bool]            `json:",omitzero,inline"`
	OfFullTextSearchConfig *FullTextSearchConfigParam `json:",omitzero,inline"`
	paramUnion
}

func (u AttributeSchemaFullTextSearchUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[AttributeSchemaFullTextSearchUnionParam](u.OfBool, u.OfFullTextSearchConfig)
}
func (u *AttributeSchemaFullTextSearchUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
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

// A list of documents in columnar format. The keys are the column names.
type DocumentColumnsParam struct {
	// The IDs of the documents.
	ID          []IDUnionParam   `json:"id,omitzero" format:"uuid"`
	ExtraFields map[string][]any `json:"-"`
	paramObj
}

func (r DocumentColumnsParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentColumnsParam
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *DocumentColumnsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single document, in a row-based format.
type DocumentRow struct {
	// An identifier for a document.
	ID IDUnion `json:"id" format:"uuid"`
	// A vector describing the document.
	Vector      DocumentRowVectorUnion `json:"vector,nullable"`
	ExtraFields map[string]any         `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Vector      respjson.Field
		ExtraFields map[string]respjson.Field
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
// DocumentRowParam.Overrides()
func (r DocumentRow) ToParam() DocumentRowParam {
	return param.Override[DocumentRowParam](r.RawJSON())
}

// DocumentRowVectorUnion contains all possible properties and values from
// [[]float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloatArray OfString]
type DocumentRowVectorUnion struct {
	// This field will be present if the value is a [[]float64] instead of an object.
	OfFloatArray []float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloatArray respjson.Field
		OfString     respjson.Field
		raw          string
	} `json:"-"`
}

func (u DocumentRowVectorUnion) AsFloatArray() (v []float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentRowVectorUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DocumentRowVectorUnion) RawJSON() string { return u.JSON.raw }

func (r *DocumentRowVectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single document, in a row-based format.
type DocumentRowParam struct {
	// A vector describing the document.
	Vector DocumentRowVectorUnionParam `json:"vector,omitzero"`
	// An identifier for a document.
	ID          IDUnionParam   `json:"id,omitzero" format:"uuid"`
	ExtraFields map[string]any `json:"-"`
	paramObj
}

func (r DocumentRowParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentRowParam
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *DocumentRowParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DocumentRowVectorUnionParam struct {
	OfFloatArray []float64         `json:",omitzero,inline"`
	OfString     param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u DocumentRowVectorUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[DocumentRowVectorUnionParam](u.OfFloatArray, u.OfString)
}
func (u *DocumentRowVectorUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DocumentRowVectorUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloatArray) {
		return &u.OfFloatArray
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CaseSensitive   respjson.Field
		Language        respjson.Field
		RemoveStopwords respjson.Field
		Stemming        respjson.Field
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
	return param.Override[FullTextSearchConfigParam](r.RawJSON())
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

func (r FullTextSearchConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow FullTextSearchConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FullTextSearchConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
		OfString respjson.Field
		OfInt    respjson.Field
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
// IDUnionParam.Overrides()
func (r IDUnion) ToParam() IDUnionParam {
	return param.Override[IDUnionParam](r.RawJSON())
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IDUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u IDUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[IDUnionParam](u.OfString, u.OfInt)
}
func (u *IDUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *IDUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
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

type NamespaceGetSchemaResponse map[string]AttributeSchema

// The result of a query.
type NamespaceQueryResponse struct {
	Aggregations []map[string]any `json:"aggregations"`
	Rows         []DocumentRow    `json:"rows"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Aggregations respjson.Field
		Rows         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgAnnCount        respjson.Field
		AvgExhaustiveCount respjson.Field
		AvgRecall          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NamespaceRecallResponse) RawJSON() string { return r.JSON.raw }
func (r *NamespaceRecallResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NamespaceUpdateSchemaResponse map[string]AttributeSchema

// The response to a successful cache warm request.
type NamespaceWarmCacheResponse struct {
	// The status of the request.
	Status  constant.Ok `json:"status,required"`
	Message string      `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NamespaceWarmCacheResponse) RawJSON() string { return r.JSON.raw }
func (r *NamespaceWarmCacheResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response to a successful upsert request.
type NamespaceWriteResponse struct {
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
func (r NamespaceWriteResponse) RawJSON() string { return r.JSON.raw }
func (r *NamespaceWriteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NamespaceDeleteAllParams struct {
	Namespace param.Opt[string] `path:"namespace,omitzero,required" json:"-"`
	paramObj
}

type NamespaceGetSchemaParams struct {
	Namespace param.Opt[string] `path:"namespace,omitzero,required" json:"-"`
	paramObj
}

type NamespaceQueryParams struct {
	Namespace param.Opt[string] `path:"namespace,omitzero,required" json:"-"`
	// How to rank the documents in the namespace.
	RankBy any `json:"rank_by,omitzero,required"`
	// The number of results to return.
	TopK int64 `json:"top_k,required"`
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
	// The encoding to use for vectors in the response.
	//
	// Any of "float", "base64".
	VectorEncoding NamespaceQueryParamsVectorEncoding `json:"vector_encoding,omitzero"`
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type NamespaceQueryParamsIncludeAttributesUnion struct {
	OfBool        param.Opt[bool] `json:",omitzero,inline"`
	OfStringArray []string        `json:",omitzero,inline"`
	paramUnion
}

func (u NamespaceQueryParamsIncludeAttributesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[NamespaceQueryParamsIncludeAttributesUnion](u.OfBool, u.OfStringArray)
}
func (u *NamespaceQueryParamsIncludeAttributesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *NamespaceQueryParamsIncludeAttributesUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// The encoding to use for vectors in the response.
type NamespaceQueryParamsVectorEncoding string

const (
	NamespaceQueryParamsVectorEncodingFloat  NamespaceQueryParamsVectorEncoding = "float"
	NamespaceQueryParamsVectorEncodingBase64 NamespaceQueryParamsVectorEncoding = "base64"
)

type NamespaceRecallParams struct {
	Namespace param.Opt[string] `path:"namespace,omitzero,required" json:"-"`
	paramObj
}

type NamespaceUpdateSchemaParams struct {
	Namespace param.Opt[string] `path:"namespace,omitzero,required" json:"-"`
	// The desired schema for the namespace.
	Body map[string]AttributeSchemaParam
	paramObj
}

func (r NamespaceUpdateSchemaParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *NamespaceUpdateSchemaParams) UnmarshalJSON(data []byte) error {
	return r.Body.UnmarshalJSON(data)
}

type NamespaceWarmCacheParams struct {
	Namespace param.Opt[string] `path:"namespace,omitzero,required" json:"-"`
	paramObj
}

type NamespaceWriteParams struct {
	Namespace param.Opt[string] `path:"namespace,omitzero,required" json:"-"`
	// The namespace to copy documents from.
	CopyFromNamespace param.Opt[string] `json:"copy_from_namespace,omitzero"`
	// The filter specifying which documents to delete.
	DeleteByFilter any            `json:"delete_by_filter,omitzero"`
	Deletes        []IDUnionParam `json:"deletes,omitzero" format:"uuid"`
	// A function used to calculate vector similarity.
	//
	// Any of "cosine_distance", "euclidean_squared".
	DistanceMetric DistanceMetric `json:"distance_metric,omitzero"`
	// A list of documents in columnar format. The keys are the column names.
	PatchColumns DocumentColumnsParam `json:"patch_columns,omitzero"`
	PatchRows    []DocumentRowParam   `json:"patch_rows,omitzero"`
	// The schema of the attributes attached to the documents.
	Schema map[string]AttributeSchemaParam `json:"schema,omitzero"`
	// A list of documents in columnar format. The keys are the column names.
	UpsertColumns DocumentColumnsParam `json:"upsert_columns,omitzero"`
	UpsertRows    []DocumentRowParam   `json:"upsert_rows,omitzero"`
	paramObj
}

func (r NamespaceWriteParams) MarshalJSON() (data []byte, err error) {
	type shadow NamespaceWriteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NamespaceWriteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
