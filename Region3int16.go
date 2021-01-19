package types

// Region3int16 represents an axis-aligned three-dimensional rectangular cuboid
// with a lower and upper boundary, with 16-bit integer precision.
type Region3int16 struct {
	Min, Max Vector3int16
}

// Type returns a string that identifies the type.
func (Region3int16) Type() string {
	return "Region3int16"
}

// String returns a human-readable string representation of the value.
func (r Region3int16) String() string {
	var b []byte
	b = append(b, r.Min.String()...)
	b = append(b, "; "...)
	b = append(b, r.Max.String()...)
	return string(b)
}

// Copy returns a copy of the value.
func (r Region3int16) Copy() PropValue {
	return r
}
