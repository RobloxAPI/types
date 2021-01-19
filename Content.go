package types

// Content is a URI pointing to a resource.
type Content string

// Type returns a string that identifies the type.
func (Content) Type() string {
	return "Content"
}

// String returns a human-readable string representation of the value.
func (s Content) String() string {
	return string(s)
}

// Copy returns a copy of the value.
func (s Content) Copy() PropValue {
	return s
}

// Stringlike returns the value as a string.
func (s Content) Stringlike() string {
	return string(s)
}
