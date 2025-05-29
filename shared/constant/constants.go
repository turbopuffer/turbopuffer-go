// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	"encoding/json"
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

type Datetime string // Always "[]datetime"
type String string   // Always "[]string"
type Uint string     // Always "[]uint"
type Uuid string     // Always "[]uuid"
type Bool string     // Always "bool"
type Datetime string // Always "datetime"
type Ok string       // Always "OK"
type String string   // Always "string"
type Uint string     // Always "uint"
type Uuid string     // Always "uuid"

func (c Datetime) Default() Datetime { return "[]datetime" }
func (c String) Default() String     { return "[]string" }
func (c Uint) Default() Uint         { return "[]uint" }
func (c Uuid) Default() Uuid         { return "[]uuid" }
func (c Bool) Default() Bool         { return "bool" }
func (c Datetime) Default() Datetime { return "datetime" }
func (c Ok) Default() Ok             { return "OK" }
func (c String) Default() String     { return "string" }
func (c Uint) Default() Uint         { return "uint" }
func (c Uuid) Default() Uuid         { return "uuid" }

func (c Datetime) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c String) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c Uint) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c Uuid) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c Bool) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c Datetime) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Ok) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c String) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c Uint) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c Uuid) MarshalJSON() ([]byte, error)     { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return json.Marshal(string(v))
}
