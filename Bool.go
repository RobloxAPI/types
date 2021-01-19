package types

import (
	"strconv"
)

// Bool represents one of two possible states.
type Bool bool

const (
	True  Bool = true
	False Bool = false
)

// Type returns a string that identifies the type.
func (Bool) Type() string {
	return "bool"
}

// String returns a human-readable string representation of the value.
func (b Bool) String() string {
	return strconv.FormatBool(bool(b))
}

// Copy returns a copy of the value.
func (b Bool) Copy() PropValue {
	return b
}
