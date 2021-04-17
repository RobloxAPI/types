package types

type BinaryString string

// Type returns a string that identifies the type.
func (BinaryString) Type() string {
	return "BinaryString"
}

// String returns a human-readable string representation of the value.
func (s BinaryString) String() string {
	return string(s)
}

// Copy returns a copy of the value.
func (s BinaryString) Copy() PropValue {
	return s
}

// Stringlike returns the value as a string.
func (s BinaryString) Stringlike() string {
	return string(s)
}
