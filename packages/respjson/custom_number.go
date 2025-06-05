package respjson

import (
	"encoding/json"
	"strconv"
)

// A Number represents a raw JSON number literal.
//
// The caller can decode the number by calling the Float64 or Int64 method, or
// interpret the string representation directly.
//
// When marshaled to JSON, the string representation is assumed to be valid
// JSON.
type Number float64

// String returns the literal text of the number.
func (n Number) String() string { return string(n) }

// Float64 returns the number as a float64.
func (n Number) Float64() (float64, error) {
	return strconv.ParseFloat(string(n), 64)
}

// Int64 returns the number as an int64.
func (n Number) Int64() (int64, error) {
	return strconv.ParseInt(string(n), 10, 64)
}

func (n Number) MarshalJSON() ([]byte, error) {
	return []byte(n), nil
}

func (n *Number) UnmarshalJSON(data []byte) error {
	var nj json.Number
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	*n = Number(nj)
	return nil
}
