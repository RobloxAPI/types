package types

// String is a sequence of characters.
type String string

// Type returns a string that identifies the type.
func (String) Type() string {
	return "string"
}

// String returns a human-readable string representation of the value.
func (s String) String() string {
	return string(s)
}

// Copy returns a copy of the value.
func (s String) Copy() PropValue {
	return s
}

// Stringlike returns the value as a string.
func (s String) Stringlike() string {
	return string(s)
}
