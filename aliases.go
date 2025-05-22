// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turbopuffer

import (
	"github.com/turbopuffer/turbopuffer-go/internal/apierror"
	"github.com/turbopuffer/turbopuffer-go/packages/param"
	"github.com/turbopuffer/turbopuffer-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// The schema for an attribute attached to a document.
//
// This is an alias to an internal type.
type AttributeSchema = shared.AttributeSchema

// Whether this attribute can be used as part of a BM25 full-text search. Requires
// the `string` or `[]string` type, and by default, BM25-enabled attributes are not
// filterable. You can override this by setting `filterable: true`.
//
// This is an alias to an internal type.
type AttributeSchemaFullTextSearchUnion = shared.AttributeSchemaFullTextSearchUnion

// The language of the text. Defaults to `english`.
//
// This is an alias to an internal type.
type AttributeSchemaFullTextSearchLanguage = shared.AttributeSchemaFullTextSearchLanguage

// Equals "arabic"
const AttributeSchemaFullTextSearchLanguageArabic = shared.AttributeSchemaFullTextSearchLanguageArabic

// Equals "danish"
const AttributeSchemaFullTextSearchLanguageDanish = shared.AttributeSchemaFullTextSearchLanguageDanish

// Equals "dutch"
const AttributeSchemaFullTextSearchLanguageDutch = shared.AttributeSchemaFullTextSearchLanguageDutch

// Equals "english"
const AttributeSchemaFullTextSearchLanguageEnglish = shared.AttributeSchemaFullTextSearchLanguageEnglish

// Equals "finnish"
const AttributeSchemaFullTextSearchLanguageFinnish = shared.AttributeSchemaFullTextSearchLanguageFinnish

// Equals "french"
const AttributeSchemaFullTextSearchLanguageFrench = shared.AttributeSchemaFullTextSearchLanguageFrench

// Equals "german"
const AttributeSchemaFullTextSearchLanguageGerman = shared.AttributeSchemaFullTextSearchLanguageGerman

// Equals "greek"
const AttributeSchemaFullTextSearchLanguageGreek = shared.AttributeSchemaFullTextSearchLanguageGreek

// Equals "hungarian"
const AttributeSchemaFullTextSearchLanguageHungarian = shared.AttributeSchemaFullTextSearchLanguageHungarian

// Equals "italian"
const AttributeSchemaFullTextSearchLanguageItalian = shared.AttributeSchemaFullTextSearchLanguageItalian

// Equals "norwegian"
const AttributeSchemaFullTextSearchLanguageNorwegian = shared.AttributeSchemaFullTextSearchLanguageNorwegian

// Equals "portuguese"
const AttributeSchemaFullTextSearchLanguagePortuguese = shared.AttributeSchemaFullTextSearchLanguagePortuguese

// Equals "romanian"
const AttributeSchemaFullTextSearchLanguageRomanian = shared.AttributeSchemaFullTextSearchLanguageRomanian

// Equals "russian"
const AttributeSchemaFullTextSearchLanguageRussian = shared.AttributeSchemaFullTextSearchLanguageRussian

// Equals "spanish"
const AttributeSchemaFullTextSearchLanguageSpanish = shared.AttributeSchemaFullTextSearchLanguageSpanish

// Equals "swedish"
const AttributeSchemaFullTextSearchLanguageSwedish = shared.AttributeSchemaFullTextSearchLanguageSwedish

// Equals "tamil"
const AttributeSchemaFullTextSearchLanguageTamil = shared.AttributeSchemaFullTextSearchLanguageTamil

// Equals "turkish"
const AttributeSchemaFullTextSearchLanguageTurkish = shared.AttributeSchemaFullTextSearchLanguageTurkish

// The data type of the attribute.
//
// This is an alias to an internal type.
type AttributeSchemaType = shared.AttributeSchemaType

// A string. Equals "string"
const AttributeSchemaTypeString = shared.AttributeSchemaTypeString

// An unsigned integer. Equals "uint"
const AttributeSchemaTypeUint = shared.AttributeSchemaTypeUint

// A UUID. Equals "uuid"
const AttributeSchemaTypeUuid = shared.AttributeSchemaTypeUuid

// A boolean. Equals "bool"
const AttributeSchemaTypeBool = shared.AttributeSchemaTypeBool

// A date and time. Equals "datetime"
const AttributeSchemaTypeDatetime = shared.AttributeSchemaTypeDatetime

// An array of strings. Equals "[]string"
const AttributeSchemaTypeStringArray = shared.AttributeSchemaTypeStringArray

// An array of unsigned integers. Equals "[]uint"
const AttributeSchemaTypeUintArray = shared.AttributeSchemaTypeUintArray

// An array of UUIDs. Equals "[]uuid"
const AttributeSchemaTypeUuidArray = shared.AttributeSchemaTypeUuidArray

// An array of date and time values. Equals "[]datetime"
const AttributeSchemaTypeDatetimeArray = shared.AttributeSchemaTypeDatetimeArray

// The schema for an attribute attached to a document.
//
// This is an alias to an internal type.
type AttributeSchemaParam = shared.AttributeSchemaParam

// Whether this attribute can be used as part of a BM25 full-text search. Requires
// the `string` or `[]string` type, and by default, BM25-enabled attributes are not
// filterable. You can override this by setting `filterable: true`.
//
// This is an alias to an internal type.
type AttributeSchemaFullTextSearchUnionParam = shared.AttributeSchemaFullTextSearchUnionParam

// A function used to calculate vector similarity.
//
// This is an alias to an internal type.
type DistanceMetric = shared.DistanceMetric

// Defined as `1 - cosine_similarity` and ranges from 0 to 2. Lower is better.
// Equals "cosine_distance"
const DistanceMetricCosineDistance = shared.DistanceMetricCosineDistance

// Defined as `sum((x - y)^2)`. Lower is better. Equals "euclidean_squared"
const DistanceMetricEuclideanSquared = shared.DistanceMetricEuclideanSquared

// A list of documents in columnar format. The keys are the column names.
//
// This is an alias to an internal type.
type DocumentColumnsParam = shared.DocumentColumnsParam

// A single document, in a row-based format.
//
// This is an alias to an internal type.
type DocumentRow = shared.DocumentRow

// A vector describing the document.
//
// This is an alias to an internal type.
type DocumentRowVectorUnion = shared.DocumentRowVectorUnion

// A single document, in a row-based format.
//
// This is an alias to an internal type.
type DocumentRowParam = shared.DocumentRowParam

// A vector describing the document.
//
// This is an alias to an internal type.
type DocumentRowVectorUnionParam = shared.DocumentRowVectorUnionParam

// Detailed configuration options for BM25 full-text search.
//
// This is an alias to an internal type.
type FullTextSearchConfig = shared.FullTextSearchConfig

// The language of the text. Defaults to `english`.
//
// This is an alias to an internal type.
type FullTextSearchConfigLanguage = shared.FullTextSearchConfigLanguage

// Equals "arabic"
const FullTextSearchConfigLanguageArabic = shared.FullTextSearchConfigLanguageArabic

// Equals "danish"
const FullTextSearchConfigLanguageDanish = shared.FullTextSearchConfigLanguageDanish

// Equals "dutch"
const FullTextSearchConfigLanguageDutch = shared.FullTextSearchConfigLanguageDutch

// Equals "english"
const FullTextSearchConfigLanguageEnglish = shared.FullTextSearchConfigLanguageEnglish

// Equals "finnish"
const FullTextSearchConfigLanguageFinnish = shared.FullTextSearchConfigLanguageFinnish

// Equals "french"
const FullTextSearchConfigLanguageFrench = shared.FullTextSearchConfigLanguageFrench

// Equals "german"
const FullTextSearchConfigLanguageGerman = shared.FullTextSearchConfigLanguageGerman

// Equals "greek"
const FullTextSearchConfigLanguageGreek = shared.FullTextSearchConfigLanguageGreek

// Equals "hungarian"
const FullTextSearchConfigLanguageHungarian = shared.FullTextSearchConfigLanguageHungarian

// Equals "italian"
const FullTextSearchConfigLanguageItalian = shared.FullTextSearchConfigLanguageItalian

// Equals "norwegian"
const FullTextSearchConfigLanguageNorwegian = shared.FullTextSearchConfigLanguageNorwegian

// Equals "portuguese"
const FullTextSearchConfigLanguagePortuguese = shared.FullTextSearchConfigLanguagePortuguese

// Equals "romanian"
const FullTextSearchConfigLanguageRomanian = shared.FullTextSearchConfigLanguageRomanian

// Equals "russian"
const FullTextSearchConfigLanguageRussian = shared.FullTextSearchConfigLanguageRussian

// Equals "spanish"
const FullTextSearchConfigLanguageSpanish = shared.FullTextSearchConfigLanguageSpanish

// Equals "swedish"
const FullTextSearchConfigLanguageSwedish = shared.FullTextSearchConfigLanguageSwedish

// Equals "tamil"
const FullTextSearchConfigLanguageTamil = shared.FullTextSearchConfigLanguageTamil

// Equals "turkish"
const FullTextSearchConfigLanguageTurkish = shared.FullTextSearchConfigLanguageTurkish

// Detailed configuration options for BM25 full-text search.
//
// This is an alias to an internal type.
type FullTextSearchConfigParam = shared.FullTextSearchConfigParam

// An identifier for a document.
//
// This is an alias to an internal type.
type IDUnion = shared.IDUnion

// An identifier for a document.
//
// This is an alias to an internal type.
type IDUnionParam = shared.IDUnionParam
