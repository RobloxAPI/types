package types

// Rect represents an axis-aligned two-dimensional rectangle with a lower and
// upper boundary.
type Rect struct {
	Min, Max Vector2
}

// Width returns the width of the rectangle.
func (r Rect) Width() float64 {
	return float64(r.Max.X - r.Min.X)
}

// Height returns the height of the rectangle.
func (r Rect) Height() float64 {
	return float64(r.Max.Y - r.Min.Y)
}

// Type returns a string that identifies the type.
func (Rect) Type() string {
	return "Rect"
}

// String returns a human-readable string representation of the value.
func (r Rect) String() string {
	var b []byte
	b = append(b, r.Min.String()...)
	b = append(b, "; "...)
	b = append(b, r.Max.String()...)
	return string(b)
}

// Copy returns a copy of the value.
func (r Rect) Copy() PropValue {
	return r
}
