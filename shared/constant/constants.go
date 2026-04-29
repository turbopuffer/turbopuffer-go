// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/turbopuffer/turbopuffer-go/v2/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type Accepted string        // Always "ACCEPTED"
type CustomerManaged string // Always "customer-managed"
type Default string         // Always "default"
type DotProduct string      // Always "dot_product"
type Ok string              // Always "OK"
type UpToDate string        // Always "up-to-date"
type Updating string        // Always "updating"

func (c Accepted) Default() Accepted               { return "ACCEPTED" }
func (c CustomerManaged) Default() CustomerManaged { return "customer-managed" }
func (c Default) Default() Default                 { return "default" }
func (c DotProduct) Default() DotProduct           { return "dot_product" }
func (c Ok) Default() Ok                           { return "OK" }
func (c UpToDate) Default() UpToDate               { return "up-to-date" }
func (c Updating) Default() Updating               { return "updating" }

func (c Accepted) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c CustomerManaged) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Default) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c DotProduct) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c Ok) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c UpToDate) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Updating) MarshalJSON() ([]byte, error)        { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
