package types

type UDim2 struct {
	X, Y UDim
}

func (u UDim2) Lerp(goal UDim2, alpha float64) UDim2 {
	panic("not implemented")
}

func (u UDim2) Add(v UDim2) UDim2 {
	return UDim2{
		X: u.X.Add(v.X),
		Y: u.Y.Add(v.Y),
	}
}

func (u UDim2) Sub(v UDim2) UDim2 {
	return UDim2{
		X: u.X.Sub(v.X),
		Y: u.Y.Sub(v.Y),
	}
}

func (u UDim2) Neg() UDim2 {
	return UDim2{
		X: u.X.Neg(),
		Y: u.Y.Neg(),
	}
}

// Type returns a string identifying the type.
func (UDim2) Type() string {
	return "UDim2"
}

func (u UDim2) String() string {
	var b []byte
	b = append(b, u.X.String()...)
	b = append(b, "; "...)
	b = append(b, u.Y.String()...)
	return string(b)
}
