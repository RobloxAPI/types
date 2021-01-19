package types

import (
	"strconv"
)

// Float is a 32-bit floating-point number.
type Float float32

// Type returns a string that identifies the type.
func (Float) Type() string {
	return "float"
}

// String returns a human-readable string representation of the value.
func (f Float) String() string {
	return strconv.FormatFloat(float64(f), 'g', -1, 32)
}

// Copy returns a copy of the value.
func (f Float) Copy() PropValue {
	return f
}

// Numberlike returns the value as a float.
func (f Float) Numberlike() float64 {
	return float64(f)
}

// Intlike returns the value as an integer.
func (f Float) Intlike() int64 {
	return int64(f)
}
