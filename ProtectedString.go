package types

type ProtectedString string

func (ProtectedString) Type() string {
	return "ProtectedString"
}

func (s ProtectedString) String() string {
	return string(s)
}

func (s ProtectedString) Copy() PropValue {
	return s
}

func (s ProtectedString) Stringlike() string {
	return string(s)
}
