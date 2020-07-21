package types

type Region3int16 struct {
	Min, Max Vector3int16
}

// Type returns a string identifying the type.
func (Region3int16) Type() string {
	return "Region3int16"
}

func (r Region3int16) String() string {
	var b []byte
	b = append(b, r.Min.String()...)
	b = append(b, "; "...)
	b = append(b, r.Max.String()...)
	return string(b)
}
