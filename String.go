package types

type String string

func (String) Type() string {
	return "string"
}

func (s String) String() string {
	return string(s)
}

func (s String) Copy() PropValue {
	return s
}

func (s String) Stringlike() string {
	return string(s)
}
