package types

// ProtectedString indicates the source of a script.
type ProtectedString string

// Type returns a string that identifies the type.
func (ProtectedString) Type() string {
	return "ProtectedString"
}

// String returns a human-readable string representation of the value.
func (s ProtectedString) String() string {
	return string(s)
}

// Copy returns a copy of the value.
func (s ProtectedString) Copy() PropValue {
	return s
}

// Stringlike returns the value as a string.
func (s ProtectedString) Stringlike() string {
	return string(s)
}
