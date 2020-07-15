package types

import (
	"strconv"
)

type NumberRange struct {
	Min, Max float32
}

func (n NumberRange) String() string {
	var b []byte
	b = strconv.AppendFloat(b, float64(n.Min), 'g', -1, 32)
	b = append(b, ", "...)
	b = strconv.AppendFloat(b, float64(n.Max), 'g', -1, 32)
	return string(b)
}
