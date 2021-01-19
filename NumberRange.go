package types

import (
	"strconv"
)

// NumberRange represents a range between two numbers as a minimum and maximum.
type NumberRange struct {
	Min, Max float32
}

// Type returns a string that identifies the type.
func (NumberRange) Type() string {
	return "NumberRange"
}

// String returns a human-readable string representation of the value.
func (n NumberRange) String() string {
	var b []byte
	b = strconv.AppendFloat(b, float64(n.Min), 'g', -1, 32)
	b = append(b, ", "...)
	b = strconv.AppendFloat(b, float64(n.Max), 'g', -1, 32)
	return string(b)
}

// Copy returns a copy of the value.
func (n NumberRange) Copy() PropValue {
	return n
}
