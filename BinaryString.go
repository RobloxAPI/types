package types

type BinaryString []byte

func (BinaryString) Type() string {
	return "BinaryString"
}

func (s BinaryString) String() string {
	return string(s)
}

func (s BinaryString) Copy() PropValue {
	c := make(BinaryString, len(s))
	copy(c, s)
	return c
}

func (s BinaryString) Stringlike() string {
	return string(s)
}
