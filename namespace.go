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
	"github.com/turbopuffer/turbopuffer-go/shared"
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

type NamespaceGetSchemaResponse map[string]shared.AttributeSchema

// The result of a query.
type NamespaceQueryResponse struct {
	Aggregations []map[string]any `json:"aggregations"`
	// The billing information for a query.
	Billing NamespaceQueryResponseBilling `json:"billing"`
	// The performance information for a query.
	Performance NamespaceQueryResponsePerformance `json:"performance"`
	Rows        []shared.DocumentRow              `json:"rows"`
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

type NamespaceUpdateSchemaResponse map[string]shared.AttributeSchema

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
	RankBy any `json:"rank_by,omitzero,required"`
	// The number of results to return.
	TopK int64 `json:"top_k,required"`
	// The consistency level for a query.
	Consistency NamespaceQueryParamsConsistency `json:"consistency,omitzero"`
	// A function used to calculate vector similarity.
	//
	// Any of "cosine_distance", "euclidean_squared".
	DistanceMetric shared.DistanceMetric `json:"distance_metric,omitzero"`
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
	// The number of searches to run.
	Num param.Opt[int64] `json:"num,omitzero"`
	// Search for `top_k` nearest neighbors.
	TopK param.Opt[int64] `json:"top_k,omitzero"`
	// Filter by attributes. Same syntax as the query endpoint.
	Filters any `json:"filters,omitzero"`
	// Use specific query vectors for the measurement. If omitted, sampled from the
	// index.
	Queries []any `json:"queries,omitzero"`
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
	Schema map[string]shared.AttributeSchemaParam
	paramObj
}

func (r NamespaceUpdateSchemaParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Schema)
}
func (r *NamespaceUpdateSchemaParams) UnmarshalJSON(data []byte) error {
	return r.Schema.UnmarshalJSON(data)
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
	DeleteByFilter any                   `json:"delete_by_filter,omitzero"`
	Deletes        []shared.IDUnionParam `json:"deletes,omitzero" format:"uuid"`
	// A function used to calculate vector similarity.
	//
	// Any of "cosine_distance", "euclidean_squared".
	DistanceMetric shared.DistanceMetric `json:"distance_metric,omitzero"`
	// A list of documents in columnar format. The keys are the column names.
	PatchColumns shared.DocumentColumnsParam `json:"patch_columns,omitzero"`
	PatchRows    []shared.DocumentRowParam   `json:"patch_rows,omitzero"`
	// The schema of the attributes attached to the documents.
	Schema map[string]shared.AttributeSchemaParam `json:"schema,omitzero"`
	// A list of documents in columnar format. The keys are the column names.
	UpsertColumns shared.DocumentColumnsParam `json:"upsert_columns,omitzero"`
	UpsertRows    []shared.DocumentRowParam   `json:"upsert_rows,omitzero"`
	paramObj
}

func (r NamespaceWriteParams) MarshalJSON() (data []byte, err error) {
	type shadow NamespaceWriteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NamespaceWriteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
