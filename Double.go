package types

import (
	"strconv"
)

// Double is a 64-bit floating-point number.
type Double float64

// Type returns a string that identifies the type.
func (Double) Type() string {
	return "double"
}

// String returns a human-readable string representation of the value.
func (d Double) String() string {
	return strconv.FormatFloat(float64(d), 'g', -1, 64)
}

// Copy returns a copy of the value.
func (d Double) Copy() PropValue {
	return d
}

// Numberlike returns the value as a float.
func (d Double) Numberlike() float64 {
	return float64(d)
}

// Intlike returns the value as an integer.
func (d Double) Intlike() int64 {
	return int64(d)
}
