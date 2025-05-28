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

// Warm the cache for a namespace.
func (r *NamespaceService) HintCacheWarm(ctx context.Context, query NamespaceHintCacheWarmParams, opts ...option.RequestOption) (res *NamespaceHintCacheWarmResponse, err error) {
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
func (r *NamespaceService) Recall(ctx context.Context, params NamespaceRecallParams, opts ...option.RequestOption) (res *NamespaceRecallResponse, err error) {
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
	path := fmt.Sprintf("v1/namespaces/%s/_debug/recall", params.Namespace.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
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
	FullTextSearch FullTextSearch `json:"full_text_search"`
	// The data type of the attribute.
	//
	// Any of "string", "uint", "uuid", "bool", "datetime", "[]string", "[]uint",
	// "[]uuid", "[]datetime".
	Type AttributeType `json:"type"`
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

// The schema for an attribute attached to a document.
type AttributeSchemaParam struct {
	// Whether or not the attributes can be used in filters/WHERE clauses.
	Filterable param.Opt[bool] `json:"filterable,omitzero"`
	// Whether this attribute can be used as part of a BM25 full-text search. Requires
	// the `string` or `[]string` type, and by default, BM25-enabled attributes are not
	// filterable. You can override this by setting `filterable: true`.
	FullTextSearch FullTextSearchParam `json:"full_text_search,omitzero"`
	// The data type of the attribute.
	//
	// Any of "string", "uint", "uuid", "bool", "datetime", "[]string", "[]uint",
	// "[]uuid", "[]datetime".
	Type AttributeType `json:"type,omitzero"`
	paramObj
}

func (r AttributeSchemaParam) MarshalJSON() (data []byte, err error) {
	type shadow AttributeSchemaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributeSchemaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The data type of the attribute.
type AttributeType string

const (
	AttributeTypeString        AttributeType = "string"
	AttributeTypeUint          AttributeType = "uint"
	AttributeTypeUuid          AttributeType = "uuid"
	AttributeTypeBool          AttributeType = "bool"
	AttributeTypeDatetime      AttributeType = "datetime"
	AttributeTypeStringArray   AttributeType = "[]string"
	AttributeTypeUintArray     AttributeType = "[]uint"
	AttributeTypeUuidArray     AttributeType = "[]uuid"
	AttributeTypeDatetimeArray AttributeType = "[]datetime"
)

// A function used to calculate vector similarity.
type DistanceMetric string

const (
	DistanceMetricCosineDistance   DistanceMetric = "cosine_distance"
	DistanceMetricEuclideanSquared DistanceMetric = "euclidean_squared"
)

// A list of documents in columnar format. The keys are the column names.
//
// The property ID is required.
type DocumentColumnsParam struct {
	// The IDs of the documents.
	ID []IDParam `json:"id,omitzero,required" format:"uuid"`
	// The vector embeddings of the documents.
	Vector      DocumentColumnsVectorParam `json:"vector,omitzero"`
	ExtraFields map[string][]any           `json:"-"`
	paramObj
}

func (r DocumentColumnsParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentColumnsParam
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *DocumentColumnsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DocumentColumnsVectorParam struct {
	VectorArray []VectorParam     `json:",omitzero,inline"`
	FloatArray  []float64         `json:",omitzero,inline"`
	String      param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u DocumentColumnsVectorParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[DocumentColumnsVectorParam](u.VectorArray, u.FloatArray, u.String)
}
func (u *DocumentColumnsVectorParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DocumentColumnsVectorParam) asAny() any {
	if !param.IsOmitted(u.VectorArray) {
		return &u.VectorArray
	} else if !param.IsOmitted(u.FloatArray) {
		return &u.FloatArray
	} else if !param.IsOmitted(u.String) {
		return &u.String.Value
	}
	return nil
}

// A single document, in a row-based format.
type DocumentRow struct {
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

// A single document, in a row-based format.
//
// The property ID is required.
type DocumentRowParam struct {
	// An identifier for a document.
	ID IDParam `json:"id,omitzero,required" format:"uuid"`
	// A vector embedding associated with a document.
	Vector      VectorParam    `json:"vector,omitzero"`
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

// FullTextSearch contains all possible properties and values from [bool],
// [FullTextSearchConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: Bool]
type FullTextSearch struct {
	// This field will be present if the value is a [bool] instead of an object.
	Bool bool `json:",inline"`
	// This field is from variant [FullTextSearchConfig].
	CaseSensitive bool `json:"case_sensitive"`
	// This field is from variant [FullTextSearchConfig].
	Language Language `json:"language"`
	// This field is from variant [FullTextSearchConfig].
	RemoveStopwords bool `json:"remove_stopwords"`
	// This field is from variant [FullTextSearchConfig].
	Stemming bool `json:"stemming"`
	// This field is from variant [FullTextSearchConfig].
	Tokenizer Tokenizer `json:"tokenizer"`
	JSON      struct {
		Bool            respjson.Field
		CaseSensitive   respjson.Field
		Language        respjson.Field
		RemoveStopwords respjson.Field
		Stemming        respjson.Field
		Tokenizer       respjson.Field
		raw             string
	} `json:"-"`
}

func (u FullTextSearch) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FullTextSearch) AsFullTextSearchConfig() (v FullTextSearchConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FullTextSearch) RawJSON() string { return u.JSON.raw }

func (r *FullTextSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FullTextSearch to a FullTextSearchParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FullTextSearchParam.Overrides()
func (r FullTextSearch) ToParam() FullTextSearchParam {
	return param.Override[FullTextSearchParam](r.RawJSON())
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FullTextSearchParam struct {
	Bool                 param.Opt[bool]            `json:",omitzero,inline"`
	FullTextSearchConfig *FullTextSearchConfigParam `json:",omitzero,inline"`
	paramUnion
}

func (u FullTextSearchParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[FullTextSearchParam](u.Bool, u.FullTextSearchConfig)
}
func (u *FullTextSearchParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FullTextSearchParam) asAny() any {
	if !param.IsOmitted(u.Bool) {
		return &u.Bool.Value
	} else if !param.IsOmitted(u.FullTextSearchConfig) {
		return u.FullTextSearchConfig
	}
	return nil
}

// Configuration options for full-text search.
type FullTextSearchConfig struct {
	// Whether searching is case-sensitive. Defaults to `false` (i.e.
	// case-insensitive).
	CaseSensitive bool `json:"case_sensitive"`
	// Describes the language of a text attribute. Defaults to `english`.
	//
	// Any of "arabic", "danish", "dutch", "english", "finnish", "french", "german",
	// "greek", "hungarian", "italian", "norwegian", "portuguese", "romanian",
	// "russian", "spanish", "swedish", "tamil", "turkish".
	Language Language `json:"language"`
	// Removes common words from the text based on language. Defaults to `true` (i.e.
	// remove common words).
	RemoveStopwords bool `json:"remove_stopwords"`
	// Language-specific stemming for the text. Defaults to `false` (i.e., do not
	// stem).
	Stemming bool `json:"stemming"`
	// The tokenizer to use for full-text search on an attribute.
	//
	// Any of "pre_tokenized_array", "word_v0", "word_v1".
	Tokenizer Tokenizer `json:"tokenizer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CaseSensitive   respjson.Field
		Language        respjson.Field
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
	return param.Override[FullTextSearchConfigParam](r.RawJSON())
}

// Configuration options for full-text search.
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
	// Describes the language of a text attribute. Defaults to `english`.
	//
	// Any of "arabic", "danish", "dutch", "english", "finnish", "french", "german",
	// "greek", "hungarian", "italian", "norwegian", "portuguese", "romanian",
	// "russian", "spanish", "swedish", "tamil", "turkish".
	Language Language `json:"language,omitzero"`
	// The tokenizer to use for full-text search on an attribute.
	//
	// Any of "pre_tokenized_array", "word_v0", "word_v1".
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
	return param.Override[IDParam](r.RawJSON())
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
	return param.MarshalUnion[IDParam](u.String, u.Int)
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
	return param.MarshalUnion[IncludeAttributesParam](u.Bool, u.StringArray)
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

// The tokenizer to use for full-text search on an attribute.
type Tokenizer string

const (
	TokenizerPreTokenizedArray Tokenizer = "pre_tokenized_array"
	TokenizerWordV0            Tokenizer = "word_v0"
	TokenizerWordV1            Tokenizer = "word_v1"
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
	return param.Override[VectorParam](r.RawJSON())
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
	return param.MarshalUnion[VectorParam](u.FloatArray, u.String)
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

// The response to a successful cache warm request.
type NamespaceHintCacheWarmResponse struct {
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
func (r NamespaceHintCacheWarmResponse) RawJSON() string { return r.JSON.raw }
func (r *NamespaceHintCacheWarmResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The result of a query.
type NamespaceQueryResponse struct {
	// The billing information for a query.
	Billing NamespaceQueryResponseBilling `json:"billing,required"`
	// The performance information for a query.
	Performance  NamespaceQueryResponsePerformance `json:"performance,required"`
	Aggregations []map[string]any                  `json:"aggregations"`
	Rows         []DocumentRow                     `json:"rows"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Billing      respjson.Field
		Performance  respjson.Field
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

// The billing information for a query.
type NamespaceQueryResponseBilling struct {
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
func (r NamespaceQueryResponseBilling) RawJSON() string { return r.JSON.raw }
func (r *NamespaceQueryResponseBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The performance information for a query.
type NamespaceQueryResponsePerformance struct {
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
func (r NamespaceQueryResponsePerformance) RawJSON() string { return r.JSON.raw }
func (r *NamespaceQueryResponsePerformance) UnmarshalJSON(data []byte) error {
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

// The response to a successful write request.
type NamespaceWriteResponse struct {
	Billing NamespaceWriteResponseBilling `json:"billing,required"`
	// A message describing the result of the write request.
	Message string `json:"message,required"`
	// The number of rows affected by the write request.
	RowsAffected int64 `json:"rows_affected,required"`
	// The status of the request.
	Status constant.Ok `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Billing      respjson.Field
		Message      respjson.Field
		RowsAffected respjson.Field
		Status       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NamespaceWriteResponse) RawJSON() string { return r.JSON.raw }
func (r *NamespaceWriteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NamespaceWriteResponseBilling struct {
	// The number of billable logical bytes written to the namespace.
	BillableLogicalBytesWritten int64 `json:"billable_logical_bytes_written,required"`
	// The billing information for a query.
	Query NamespaceWriteResponseBillingQuery `json:"query"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillableLogicalBytesWritten respjson.Field
		Query                       respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NamespaceWriteResponseBilling) RawJSON() string { return r.JSON.raw }
func (r *NamespaceWriteResponseBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The billing information for a query.
type NamespaceWriteResponseBillingQuery struct {
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
func (r NamespaceWriteResponseBillingQuery) RawJSON() string { return r.JSON.raw }
func (r *NamespaceWriteResponseBillingQuery) UnmarshalJSON(data []byte) error {
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

type NamespaceHintCacheWarmParams struct {
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
	IncludeAttributes IncludeAttributesParam `json:"include_attributes,omitzero"`
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

// The encoding to use for vectors in the response.
type NamespaceQueryParamsVectorEncoding string

const (
	NamespaceQueryParamsVectorEncodingFloat  NamespaceQueryParamsVectorEncoding = "float"
	NamespaceQueryParamsVectorEncodingBase64 NamespaceQueryParamsVectorEncoding = "base64"
)

type NamespaceRecallParams struct {
	Namespace param.Opt[string] `path:"namespace,omitzero,required" json:"-"`
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

type NamespaceUpdateSchemaParams struct {
	Namespace param.Opt[string] `path:"namespace,omitzero,required" json:"-"`
	// The desired schema for the namespace.
	Schema map[string]AttributeSchemaParam
	paramObj
}

func (r NamespaceUpdateSchemaParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Schema)
}
func (r *NamespaceUpdateSchemaParams) UnmarshalJSON(data []byte) error {
	return r.Schema.UnmarshalJSON(data)
}

type NamespaceWriteParams struct {
	Namespace param.Opt[string] `path:"namespace,omitzero,required" json:"-"`
	// The namespace to copy documents from.
	CopyFromNamespace param.Opt[string] `json:"copy_from_namespace,omitzero"`
	// The filter specifying which documents to delete.
	DeleteByFilter any       `json:"delete_by_filter,omitzero"`
	Deletes        []IDParam `json:"deletes,omitzero" format:"uuid"`
	// A function used to calculate vector similarity.
	//
	// Any of "cosine_distance", "euclidean_squared".
	DistanceMetric DistanceMetric `json:"distance_metric,omitzero"`
	// The encryption configuration for a namespace.
	Encryption NamespaceWriteParamsEncryption `json:"encryption,omitzero"`
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
