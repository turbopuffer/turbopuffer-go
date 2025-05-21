// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/turbopuffer/turbopuffer-go/internal/apijson"
	"github.com/turbopuffer/turbopuffer-go/packages/param"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FilterUnionParam struct {
	OfAnyArray    []any `json:",omitzero,inline"`
	OfFilterArray []any `json:",omitzero,inline"`
	OfVariant2    []any `json:",omitzero,inline"`
	OfVariant3    []any `json:",omitzero,inline"`
	paramUnion
}

func (u FilterUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[FilterUnionParam](u.OfAnyArray, u.OfFilterArray, u.OfVariant2, u.OfVariant3)
}
func (u *FilterUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FilterUnionParam) asAny() any {
	if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	} else if !param.IsOmitted(u.OfFilterArray) {
		return &u.OfFilterArray
	} else if !param.IsOmitted(u.OfVariant2) {
		return &u.OfVariant2
	} else if !param.IsOmitted(u.OfVariant3) {
		return &u.OfVariant3
	}
	return nil
}
