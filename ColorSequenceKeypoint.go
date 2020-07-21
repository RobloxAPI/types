package types

import (
	"strconv"
)

type ColorSequenceKeypoint struct {
	Time     float32
	Value    Color3
	Envelope float32
}

// Type returns a string identifying the type.
func (ColorSequenceKeypoint) Type() string {
	return "ColorSequenceKeypoint"
}

func (c ColorSequenceKeypoint) String() string {
	var b []byte
	b = strconv.AppendFloat(b, float64(c.Time), 'g', -1, 32)
	b = append(b, ", ("...)
	b = append(b, c.Value.String()...)
	b = append(b, "), "...)
	b = strconv.AppendFloat(b, float64(c.Time), 'g', -1, 32)
	return string(b)
}
