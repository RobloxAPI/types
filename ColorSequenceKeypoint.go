package types

import (
	"strconv"
)

// ColorSequenceKeypoint is a keypoint in a ColorSequence.
type ColorSequenceKeypoint struct {
	Time     float32
	Value    Color3
	Envelope float32
}

// Type returns a string that identifies the type.
func (ColorSequenceKeypoint) Type() string {
	return "ColorSequenceKeypoint"
}

// String returns a human-readable string representation of the value.
func (c ColorSequenceKeypoint) String() string {
	var b []byte
	b = strconv.AppendFloat(b, float64(c.Time), 'g', -1, 32)
	b = append(b, ", ("...)
	b = append(b, c.Value.String()...)
	b = append(b, "), "...)
	b = strconv.AppendFloat(b, float64(c.Time), 'g', -1, 32)
	return string(b)
}

// Copy returns a copy of the value.
func (c ColorSequenceKeypoint) Copy() PropValue {
	return c
}
