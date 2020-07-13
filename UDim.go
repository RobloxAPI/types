package types

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
