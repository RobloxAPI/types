package types

type Rect struct {
	Min, Max Vector2
}

func (r Rect) Width() float64 {
	return float64(r.Max.X - r.Min.X)
}

func (r Rect) Height() float64 {
	return float64(r.Max.Y - r.Min.Y)
}

// Type returns a string identifying the type.
func (Rect) Type() string {
	return "Rect"
}

func (r Rect) String() string {
	var b []byte
	b = append(b, r.Min.String()...)
	b = append(b, "; "...)
	b = append(b, r.Max.String()...)
	return string(b)
}

func (r Rect) Copy() PropValue {
	return r
}
