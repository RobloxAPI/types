package types

import (
	"strconv"
)

// NumberSequenceKeypoint is a keypoint in a NumberSequence.
type NumberSequenceKeypoint struct {
	Time     float32
	Value    float32
	Envelope float32
}

// Type returns a string identifying the type.
func (NumberSequenceKeypoint) Type() string {
	return "NumberSequenceKeypoint"
}

// String returns a human-readable string representation of the value.
func (n NumberSequenceKeypoint) String() string {
	var b []byte
	b = strconv.AppendFloat(b, float64(n.Time), 'g', -1, 32)
	b = append(b, ", "...)
	b = strconv.AppendFloat(b, float64(n.Value), 'g', -1, 32)
	b = append(b, ", "...)
	b = strconv.AppendFloat(b, float64(n.Time), 'g', -1, 32)
	return string(b)
}

// Copy returns a copy of the value.
func (n NumberSequenceKeypoint) Copy() PropValue {
	return n
}
