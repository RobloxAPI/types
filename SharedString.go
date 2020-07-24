package types

type SharedString []byte

func (SharedString) Type() string {
	return "SharedString"
}

func (s SharedString) String() string {
	return string(s)
}

func (s SharedString) Copy() PropValue {
	c := make(SharedString, len(s))
	copy(c, s)
	return c
}

func (s SharedString) Stringlike() string {
	return string(s)
}
