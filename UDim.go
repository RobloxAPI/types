package types

import (
	"strconv"
)

type UDim struct {
	Scale  float32
	Offset int32
}

func (u UDim) Add(v UDim) UDim {
	return UDim{
		Scale:  u.Scale + v.Scale,
		Offset: u.Offset + v.Offset,
	}
}

func (u UDim) Sub(v UDim) UDim {
	return UDim{
		Scale:  u.Scale - v.Scale,
		Offset: u.Offset - v.Offset,
	}
}

func (u UDim) Neg() UDim {
	return UDim{
		Scale:  -u.Scale,
		Offset: -u.Offset,
	}
}

// Type returns a string identifying the type.
func (UDim) Type() string {
	return "UDim"
}

func (u UDim) String() string {
	var b []byte
	b = strconv.AppendFloat(b, float64(u.Scale), 'g', -1, 32)
	b = append(b, ", "...)
	b = strconv.AppendInt(b, int64(u.Offset), 10)
	return string(b)
}

func (u UDim) Copy() PropValue {
	return u
}
