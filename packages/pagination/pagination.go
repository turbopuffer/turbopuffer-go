// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/turbopuffer/turbopuffer-go/internal/apijson"
	"github.com/turbopuffer/turbopuffer-go/internal/requestconfig"
	"github.com/turbopuffer/turbopuffer-go/option"
	"github.com/turbopuffer/turbopuffer-go/packages/param"
	"github.com/turbopuffer/turbopuffer-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type NamespacePage[T any] struct {
	Namespaces []T    `json:"namespaces"`
	NextCursor string `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Namespaces  respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r NamespacePage[T]) RawJSON() string { return r.JSON.raw }
func (r *NamespacePage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *NamespacePage[T]) GetNextPage() (res *NamespacePage[T], err error) {
	if len(r.Namespaces) == 0 {
		return nil, nil
	}
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *NamespacePage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &NamespacePage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type NamespacePageAutoPager[T any] struct {
	page *NamespacePage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewNamespacePageAutoPager[T any](page *NamespacePage[T], err error) *NamespacePageAutoPager[T] {
	return &NamespacePageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *NamespacePageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Namespaces) == 0 {
		return false
	}
	if r.idx >= len(r.page.Namespaces) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Namespaces) == 0 {
			return false
		}
	}
	r.cur = r.page.Namespaces[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *NamespacePageAutoPager[T]) Current() T {
	return r.cur
}

func (r *NamespacePageAutoPager[T]) Err() error {
	return r.err
}

func (r *NamespacePageAutoPager[T]) Index() int {
	return r.run
}
