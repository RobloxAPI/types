package types

import (
	"strconv"
)

// Vector2int16 is a two-dimensional Euclidean vector with 16-bit integer
// precision.
type Vector2int16 struct {
	X, Y int16
}

// NewVector2int16 returns a vector initialized with the given components.
func NewVector2int16(x, y int) Vector2int16 {
	return Vector2int16{X: int16(x), Y: int16(y)}
}

// Add returns the sum of two vectors.
func (v Vector2int16) Add(op Vector2int16) Vector2int16 {
	return Vector2int16{X: v.X + op.X, Y: v.Y + op.Y}
}

// Sub returns the difference of two vectors.
func (v Vector2int16) Sub(op Vector2int16) Vector2int16 {
	return Vector2int16{X: v.X - op.X, Y: v.Y - op.Y}
}

// Mul returns the product of two vectors.
func (v Vector2int16) Mul(op Vector2int16) Vector2int16 {
	return Vector2int16{X: v.X * op.X, Y: v.Y * op.Y}
}

// Div returns the quotient of two vectors.
func (v Vector2int16) Div(op Vector2int16) Vector2int16 {
	return Vector2int16{X: v.X / op.X, Y: v.Y / op.Y}
}

// MulN returns the vector with each component multiplied by a number.
func (v Vector2int16) MulN(op float64) Vector2int16 {
	return Vector2int16{X: int16(float64(v.X) * op), Y: int16(float64(v.Y) * op)}
}

// DivN returns the vector with each component divided by a number.
func (v Vector2int16) DivN(op float64) Vector2int16 {
	return Vector2int16{X: int16(float64(v.X) / op), Y: int16(float64(v.Y) / op)}
}

// Neg returns the negation of the vector.
func (v Vector2int16) Neg() Vector2int16 {
	return Vector2int16{X: -v.X, Y: -v.Y}
}

// Type returns a string that identifies the type.
func (Vector2int16) Type() string {
	return "Vector2int16"
}

// String returns a human-readable string representation of the value.
func (v Vector2int16) String() string {
	var b []byte
	b = strconv.AppendInt(b, int64(v.X), 10)
	b = append(b, ", "...)
	b = strconv.AppendInt(b, int64(v.Y), 10)
	return string(b)
}

// Copy returns a copy of the value.
func (v Vector2int16) Copy() PropValue {
	return v
}
