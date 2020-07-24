package types

type Content string

func (Content) Type() string {
	return "Content"
}

func (s Content) String() string {
	return string(s)
}

func (s Content) Copy() PropValue {
	return s
}

func (s Content) Stringlike() string {
	return string(s)
}
