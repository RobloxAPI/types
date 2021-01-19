package types

// ColorSequence represents an interpolated sequence of colors.
type ColorSequence []ColorSequenceKeypoint

// Type returns a string that identifies the type.
func (ColorSequence) Type() string {
	return "ColorSequence"
}

// String returns a human-readable string representation of the value.
func (c ColorSequence) String() string {
	var b []byte
	for i, k := range c {
		if i > 0 {
			b = append(b, "; "...)
		}
		b = append(b, k.String()...)
	}
	return string(b)
}

// Copy returns a copy of the value.
func (c ColorSequence) Copy() PropValue {
	d := make(ColorSequence, len(c))
	copy(d, c)
	return d
}
