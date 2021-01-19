package types

// UDim2 is a UDim on two dimensions.
type UDim2 struct {
	X, Y UDim
}

// Lerp returns a UDim2 linearly interpolated from the UDim2 to *goal* according
// to *alpha*, which has an interval of [0, 1].
func (u UDim2) Lerp(goal UDim2, alpha float64) UDim2 {
	a := float32(alpha)
	na := 1 - a
	return UDim2{
		X: UDim{
			Scale:  na*u.X.Scale + a*goal.X.Scale,
			Offset: int32(na*float32(u.X.Offset) + a*float32(goal.X.Offset)),
		},
		Y: UDim{
			Scale:  na*u.Y.Scale + a*goal.Y.Scale,
			Offset: int32(na*float32(u.Y.Offset) + a*float32(goal.Y.Offset)),
		},
	}
}

// Add returns the sum of two UDim2s.
func (u UDim2) Add(v UDim2) UDim2 {
	return UDim2{
		X: u.X.Add(v.X),
		Y: u.Y.Add(v.Y),
	}
}

// Sub returns the difference of two UDim2s.
func (u UDim2) Sub(v UDim2) UDim2 {
	return UDim2{
		X: u.X.Sub(v.X),
		Y: u.Y.Sub(v.Y),
	}
}

// Neg returns the negation of the UDim2.
func (u UDim2) Neg() UDim2 {
	return UDim2{
		X: u.X.Neg(),
		Y: u.Y.Neg(),
	}
}

// Type returns a string that identifies the type.
func (UDim2) Type() string {
	return "UDim2"
}

// String returns a human-readable string representation of the value.
func (u UDim2) String() string {
	var b []byte
	b = append(b, u.X.String()...)
	b = append(b, "; "...)
	b = append(b, u.Y.String()...)
	return string(b)
}

// Copy returns a copy of the value.
func (u UDim2) Copy() PropValue {
	return u
}
