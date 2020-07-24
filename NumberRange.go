package types

import (
	"strconv"
)

type NumberRange struct {
	Min, Max float32
}

// Type returns a string identifying the type.
func (NumberRange) Type() string {
	return "NumberRange"
}

func (n NumberRange) String() string {
	var b []byte
	b = strconv.AppendFloat(b, float64(n.Min), 'g', -1, 32)
	b = append(b, ", "...)
	b = strconv.AppendFloat(b, float64(n.Max), 'g', -1, 32)
	return string(b)
}

func (n NumberRange) Copy() PropValue {
	return n
}
