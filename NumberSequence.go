package types

// NumberSequence represents an interpolated sequence of numbers.
type NumberSequence []NumberSequenceKeypoint

// Type returns a string that identifies the type.
func (NumberSequence) Type() string {
	return "NumberSequence"
}

// String returns a human-readable string representation of the value.
func (n NumberSequence) String() string {
	var b []byte
	for i, k := range n {
		if i > 0 {
			b = append(b, "; "...)
		}
		b = append(b, k.String()...)
	}
	return string(b)
}

// Copy returns a copy of the value.
func (n NumberSequence) Copy() PropValue {
	c := make(NumberSequence, len(n))
	copy(c, n)
	return c
}
