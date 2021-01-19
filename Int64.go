package types

import (
	"strconv"
)

// Int64 is a 64-bit signed integer.
type Int64 int64

// Type returns a string that identifies the type.
func (Int64) Type() string {
	return "int64"
}

// String returns a human-readable string representation of the value.
func (i Int64) String() string {
	return strconv.FormatInt(int64(i), 10)
}

// Copy returns a copy of the value.
func (i Int64) Copy() PropValue {
	return i
}

// Numberlike returns the value as a float.
func (i Int64) Numberlike() float64 {
	return float64(i)
}

// Intlike returns the value as an integer.
func (i Int64) Intlike() int64 {
	return int64(i)
}
