package types

type NumberSequence []NumberSequenceKeypoint

// Type returns a string identifying the type.
func (NumberSequence) Type() string {
	return "NumberSequence"
}

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

func (n NumberSequence) Copy() PropValue {
	c := make(NumberSequence, len(n))
	copy(c, n)
	return c
}
