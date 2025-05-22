// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"

	"github.com/turbopuffer/turbopuffer-go/internal/apijson"
	"github.com/turbopuffer/turbopuffer-go/packages/param"
	"github.com/turbopuffer/turbopuffer-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

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
