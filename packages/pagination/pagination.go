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

type Export[T any] struct {
	IDs        []T    `json:"ids"`
	NextCursor string `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IDs         respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r Export[T]) RawJSON() string { return r.JSON.raw }
func (r *Export[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *Export[T]) GetNextPage() (res *Export[T], err error) {
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

func (r *Export[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &Export[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ExportAutoPager[T any] struct {
	page *Export[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewExportAutoPager[T any](page *Export[T], err error) *ExportAutoPager[T] {
	return &ExportAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ExportAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.IDs) == 0 {
		return false
	}
	if r.idx >= len(r.page.IDs) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.IDs) == 0 {
			return false
		}
	}
	r.cur = r.page.IDs[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ExportAutoPager[T]) Current() T {
	return r.cur
}

func (r *ExportAutoPager[T]) Err() error {
	return r.err
}

func (r *ExportAutoPager[T]) Index() int {
	return r.run
}

type ListNamespaces[T any] struct {
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
func (r ListNamespaces[T]) RawJSON() string { return r.JSON.raw }
func (r *ListNamespaces[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *ListNamespaces[T]) GetNextPage() (res *ListNamespaces[T], err error) {
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

func (r *ListNamespaces[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &ListNamespaces[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ListNamespacesAutoPager[T any] struct {
	page *ListNamespaces[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewListNamespacesAutoPager[T any](page *ListNamespaces[T], err error) *ListNamespacesAutoPager[T] {
	return &ListNamespacesAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ListNamespacesAutoPager[T]) Next() bool {
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

func (r *ListNamespacesAutoPager[T]) Current() T {
	return r.cur
}

func (r *ListNamespacesAutoPager[T]) Err() error {
	return r.err
}

func (r *ListNamespacesAutoPager[T]) Index() int {
	return r.run
}
