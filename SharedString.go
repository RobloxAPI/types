package types

// SharedString represents data that is shared by multiple objects.
type SharedString string

// Type returns a string that identifies the type.
func (SharedString) Type() string {
	return "SharedString"
}

// String returns a human-readable string representation of the value.
func (s SharedString) String() string {
	return string(s)
}

// Copy returns a copy of the value.
func (s SharedString) Copy() PropValue {
	return s
}

// Stringlike returns the value as a string.
func (s SharedString) Stringlike() string {
	return string(s)
}
