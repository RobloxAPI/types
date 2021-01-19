package types

import (
	"strconv"
)

// Token represents the value of an enumerated type.
type Token uint32

// Type returns a string that identifies the type.
func (Token) Type() string {
	return "token"
}

// String returns a human-readable string representation of the value.
func (t Token) String() string {
	return strconv.FormatUint(uint64(t), 10)
}

// Copy returns a copy of the value.
func (t Token) Copy() PropValue {
	return t
}

// Numberlike returns the value as a float.
func (t Token) Numberlike() float64 {
	return float64(t)
}

// Intlike returns the value as an integer.
func (t Token) Intlike() int64 {
	return int64(t)
}
