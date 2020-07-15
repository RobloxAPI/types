package types

import (
	"strconv"
)

type NumberSequenceKeypoint struct {
	Time     float32
	Value    float32
	Envelope float32
}

func (c NumberSequenceKeypoint) String() string {
	var b []byte
	b = strconv.AppendFloat(b, float64(c.Time), 'g', -1, 32)
	b = append(b, ", "...)
	b = strconv.AppendFloat(b, float64(c.Value), 'g', -1, 32)
	b = append(b, ", "...)
	b = strconv.AppendFloat(b, float64(c.Time), 'g', -1, 32)
	return string(b)
}
