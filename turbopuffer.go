// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turbopuffer

import (
	"net/url"

	"github.com/turbopuffer/turbopuffer-go/internal/apijson"
	"github.com/turbopuffer/turbopuffer-go/internal/apiquery"
	"github.com/turbopuffer/turbopuffer-go/packages/param"
	"github.com/turbopuffer/turbopuffer-go/packages/respjson"
)

// A summary of a namespace.
type NamespaceSummary struct {
	// The namespace ID.
	ID string `json:"id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NamespaceSummary) RawJSON() string { return r.JSON.raw }
func (r *NamespaceSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NamespacesParams struct {
	// Retrieve the next page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Limit the number of results per page.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Retrieve only the namespaces that match the prefix.
	Prefix param.Opt[string] `query:"prefix,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NamespacesParams]'s query parameters as `url.Values`.
func (r NamespacesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
