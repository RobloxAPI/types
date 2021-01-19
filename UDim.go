package types

import (
	"strconv"
)

// UDim represents one dimension with a dynamic and constant component.
type UDim struct {
	Scale  float32
	Offset int32
}

// Adds returns the sum of two UDims.
func (u UDim) Add(v UDim) UDim {
	return UDim{
		Scale:  u.Scale + v.Scale,
		Offset: u.Offset + v.Offset,
	}
}

// Sub returns the difference of two UDims.
func (u UDim) Sub(v UDim) UDim {
	return UDim{
		Scale:  u.Scale - v.Scale,
		Offset: u.Offset - v.Offset,
	}
}

// Neg returns the negation of the UDim.
func (u UDim) Neg() UDim {
	return UDim{
		Scale:  -u.Scale,
		Offset: -u.Offset,
	}
}

// Type returns a string that identifies the type.
func (UDim) Type() string {
	return "UDim"
}

// String returns a human-readable string representation of the value.
func (u UDim) String() string {
	var b []byte
	b = strconv.AppendFloat(b, float64(u.Scale), 'g', -1, 32)
	b = append(b, ", "...)
	b = strconv.AppendInt(b, int64(u.Offset), 10)
	return string(b)
}

// Copy returns a copy of the value.
func (u UDim) Copy() PropValue {
	return u
}
