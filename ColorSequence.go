package types

type ColorSequence []ColorSequenceKeypoint

// Type returns a string identifying the type.
func (ColorSequence) Type() string {
	return "ColorSequence"
}

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
