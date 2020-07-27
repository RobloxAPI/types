package types

// Value is a value that has a type.
type Value interface {
	Type() string
}

// Stringer is any Value that has a string representation.
type Stringer interface {
	Value
	String() string
}

// PropValue is a Value that can be copied.
type PropValue interface {
	Value
	Copy() PropValue
}

// Aliaser is a Value that has an underlying Value of a different type.
type Aliaser interface {
	Value
	Alias() Value
}

// Stringlike is any Value that is string-like. Note that this is distinct from
// a string-representation of the value. Rather, Stringlike indicates that the
// value has string-like properties.
type Stringlike interface {
	Value
	Stringlike() string
}

// Numberlike is any value that can be converted to a floating-point number.
type Numberlike interface {
	Value
	Numberlike() float64
}

// Intlike is any value that can be converted to an integer.
type Intlike interface {
	Value
	Intlike() int64
}
