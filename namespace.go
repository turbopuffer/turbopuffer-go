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
func newNamespaceService(opts ...option.RequestOption) (r NamespaceService) {
	r = NamespaceService{}
	r.Options = opts
	return
}

func (r *NamespaceService) ID() string {
	requestConfig, err := requestconfig.NewRequestConfig(context.Background(), "GET", "/", nil, nil, r.Options...)
	if err != nil {
		panic(fmt.Errorf("failed to create request config: %w", err))
	}
	return *requestConfig.DefaultNamespace
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
	FullTextSearch FullTextSearchConfigUnion `json:"full_text_search"`
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
	FullTextSearch FullTextSearchConfigUnionParam `json:"full_text_search,omitzero"`
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
	ID []IDUnionParam `json:"id,omitzero,required" format:"uuid"`
	// The vector embeddings of the documents.
	Vector      DocumentColumnsVectorUnionParam `json:"vector,omitzero"`
	ExtraFields map[string][]any                `json:"-"`
	paramObj
}

func (r DocumentColumnsParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentColumnsParam
	extraFields := make(map[string]any)
	for fieldName, fieldValue := range r.ExtraFields {
		extraFields[fieldName] = fieldValue
	}
	return param.MarshalWithExtras(r, (*shadow)(&r), extraFields)
}
func (r *DocumentColumnsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DocumentColumnsVectorUnionParam struct {
	OfVectorArray []VectorUnionParam `json:",omitzero,inline"`
	OfFloatArray  []float64          `json:",omitzero,inline"`
	OfString      param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u DocumentColumnsVectorUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[DocumentColumnsVectorUnionParam](u.OfVectorArray, u.OfFloatArray, u.OfString)
}
func (u *DocumentColumnsVectorUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DocumentColumnsVectorUnionParam) asAny() any {
	if !param.IsOmitted(u.OfVectorArray) {
		return &u.OfVectorArray
	} else if !param.IsOmitted(u.OfFloatArray) {
		return &u.OfFloatArray
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// A single document, in a row-based format.
type DocumentRow struct {
	// An identifier for a document.
	ID IDUnion `json:"id,required" format:"uuid"`
	// A vector embedding associated with a document.
	Vector      VectorUnion    `json:"vector"`
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
	ID IDUnionParam `json:"id,omitzero,required" format:"uuid"`
	// A vector embedding associated with a document.
	Vector      VectorUnionParam `json:"vector,omitzero"`
	ExtraFields map[string]any   `json:"-"`
	paramObj
}

func (r DocumentRowParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentRowParam
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *DocumentRowParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FullTextSearchConfigUnion contains all possible properties and values from
// [bool], [FullTextSearchConfigDetailed].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool]
type FullTextSearchConfigUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field is from variant [FullTextSearchConfigDetailed].
	CaseSensitive bool `json:"case_sensitive"`
	// This field is from variant [FullTextSearchConfigDetailed].
	Language FullTextSearchConfigDetailedLanguage `json:"language"`
	// This field is from variant [FullTextSearchConfigDetailed].
	RemoveStopwords bool `json:"remove_stopwords"`
	// This field is from variant [FullTextSearchConfigDetailed].
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

func (u FullTextSearchConfigUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FullTextSearchConfigUnion) AsDetailed() (v FullTextSearchConfigDetailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FullTextSearchConfigUnion) RawJSON() string { return u.JSON.raw }

func (r *FullTextSearchConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FullTextSearchConfigUnion to a
// FullTextSearchConfigUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FullTextSearchConfigUnionParam.Overrides()
func (r FullTextSearchConfigUnion) ToParam() FullTextSearchConfigUnionParam {
	return param.Override[FullTextSearchConfigUnionParam](r.RawJSON())
}

type FullTextSearchConfigDetailed struct {
	// Whether searching is case-sensitive. Defaults to `false` (i.e.
	// case-insensitive).
	CaseSensitive bool `json:"case_sensitive"`
	// The language of the text. Defaults to `english`.
	//
	// Any of "arabic", "danish", "dutch", "english", "finnish", "french", "german",
	// "greek", "hungarian", "italian", "norwegian", "portuguese", "romanian",
	// "russian", "spanish", "swedish", "tamil", "turkish".
	Language FullTextSearchConfigDetailedLanguage `json:"language"`
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
func (r FullTextSearchConfigDetailed) RawJSON() string { return r.JSON.raw }
func (r *FullTextSearchConfigDetailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The language of the text. Defaults to `english`.
type FullTextSearchConfigDetailedLanguage string

const (
	FullTextSearchConfigDetailedLanguageArabic     FullTextSearchConfigDetailedLanguage = "arabic"
	FullTextSearchConfigDetailedLanguageDanish     FullTextSearchConfigDetailedLanguage = "danish"
	FullTextSearchConfigDetailedLanguageDutch      FullTextSearchConfigDetailedLanguage = "dutch"
	FullTextSearchConfigDetailedLanguageEnglish    FullTextSearchConfigDetailedLanguage = "english"
	FullTextSearchConfigDetailedLanguageFinnish    FullTextSearchConfigDetailedLanguage = "finnish"
	FullTextSearchConfigDetailedLanguageFrench     FullTextSearchConfigDetailedLanguage = "french"
	FullTextSearchConfigDetailedLanguageGerman     FullTextSearchConfigDetailedLanguage = "german"
	FullTextSearchConfigDetailedLanguageGreek      FullTextSearchConfigDetailedLanguage = "greek"
	FullTextSearchConfigDetailedLanguageHungarian  FullTextSearchConfigDetailedLanguage = "hungarian"
	FullTextSearchConfigDetailedLanguageItalian    FullTextSearchConfigDetailedLanguage = "italian"
	FullTextSearchConfigDetailedLanguageNorwegian  FullTextSearchConfigDetailedLanguage = "norwegian"
	FullTextSearchConfigDetailedLanguagePortuguese FullTextSearchConfigDetailedLanguage = "portuguese"
	FullTextSearchConfigDetailedLanguageRomanian   FullTextSearchConfigDetailedLanguage = "romanian"
	FullTextSearchConfigDetailedLanguageRussian    FullTextSearchConfigDetailedLanguage = "russian"
	FullTextSearchConfigDetailedLanguageSpanish    FullTextSearchConfigDetailedLanguage = "spanish"
	FullTextSearchConfigDetailedLanguageSwedish    FullTextSearchConfigDetailedLanguage = "swedish"
	FullTextSearchConfigDetailedLanguageTamil      FullTextSearchConfigDetailedLanguage = "tamil"
	FullTextSearchConfigDetailedLanguageTurkish    FullTextSearchConfigDetailedLanguage = "turkish"
)

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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FullTextSearchConfigUnionParam struct {
	OfBool     param.Opt[bool]                    `json:",omitzero,inline"`
	OfDetailed *FullTextSearchConfigDetailedParam `json:",omitzero,inline"`
	paramUnion
}

func (u FullTextSearchConfigUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[FullTextSearchConfigUnionParam](u.OfBool, u.OfDetailed)
}
func (u *FullTextSearchConfigUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FullTextSearchConfigUnionParam) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfDetailed) {
		return u.OfDetailed
	}
	return nil
}

type FullTextSearchConfigDetailedParam struct {
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
	Language FullTextSearchConfigDetailedLanguage `json:"language,omitzero"`
	paramObj
}

func (r FullTextSearchConfigDetailedParam) MarshalJSON() (data []byte, err error) {
	type shadow FullTextSearchConfigDetailedParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FullTextSearchConfigDetailedParam) UnmarshalJSON(data []byte) error {
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IncludeAttributesUnionParam struct {
	OfBool        param.Opt[bool] `json:",omitzero,inline"`
	OfStringArray []string        `json:",omitzero,inline"`
	paramUnion
}

func (u IncludeAttributesUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[IncludeAttributesUnionParam](u.OfBool, u.OfStringArray)
}
func (u *IncludeAttributesUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *IncludeAttributesUnionParam) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// VectorUnion contains all possible properties and values from [[]float64],
// [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloatArray OfString]
type VectorUnion struct {
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

func (u VectorUnion) AsFloatArray() (v []float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VectorUnion) RawJSON() string { return u.JSON.raw }

func (r *VectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this VectorUnion to a VectorUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// VectorUnionParam.Overrides()
func (r VectorUnion) ToParam() VectorUnionParam {
	return param.Override[VectorUnionParam](r.RawJSON())
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VectorUnionParam struct {
	OfFloatArray []float64         `json:",omitzero,inline"`
	OfString     param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u VectorUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[VectorUnionParam](u.OfFloatArray, u.OfString)
}
func (u *VectorUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloatArray) {
		return &u.OfFloatArray
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
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
	// The billing information for a query.
	Billing NamespaceQueryResponseBilling `json:"billing"`
	// The performance information for a query.
	Performance NamespaceQueryResponsePerformance `json:"performance"`
	Rows        []DocumentRow                     `json:"rows"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Aggregations respjson.Field
		Billing      respjson.Field
		Performance  respjson.Field
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

type NamespaceQueryParams struct {
	Namespace param.Opt[string] `path:"namespace,omitzero,required" json:"-"`
	// How to rank the documents in the namespace.
	RankBy RankBy `json:"rank_by,omitzero,required"`
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
	Filters Filter `json:"filters,omitzero"`
	// Whether to include attributes in the response.
	IncludeAttributes IncludeAttributesUnionParam `json:"include_attributes,omitzero"`
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
	return apijson.UnmarshalRoot(data, r)
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
	DeleteByFilter Filter         `json:"delete_by_filter,omitzero"`
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
