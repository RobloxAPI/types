package types

import (
	"strconv"
)

// Vector3int16 is a three-dimensional Euclidean vector with 16-bit integer
// precision.
type Vector3int16 struct {
	X, Y, Z int16
}

// NewVector3int16 returns a vector initialized with the given components.
func NewVector3int16(x, y, z int) Vector3int16 {
	return Vector3int16{X: int16(x), Y: int16(y), Z: int16(z)}
}

// Add returns the sum of two vectors.
func (v Vector3int16) Add(op Vector3int16) Vector3int16 {
	return Vector3int16{X: v.X + op.X, Y: v.Y + op.Y, Z: v.Z + op.Z}
}

// Sub returns the difference of two vectors.
func (v Vector3int16) Sub(op Vector3int16) Vector3int16 {
	return Vector3int16{X: v.X - op.X, Y: v.Y - op.Y, Z: v.Z - op.Z}
}

// Mul returns the product of two vectors.
func (v Vector3int16) Mul(op Vector3int16) Vector3int16 {
	return Vector3int16{X: v.X * op.X, Y: v.Y * op.Y, Z: v.Z * op.Z}
}

// Div returns the quotient of two vectors.
func (v Vector3int16) Div(op Vector3int16) Vector3int16 {
	return Vector3int16{X: v.X / op.X, Y: v.Y / op.Y, Z: v.Z / op.Z}
}

// MulN returns the vector with each component multiplied by a number.
func (v Vector3int16) MulN(op float64) Vector3int16 {
	return Vector3int16{X: int16(float64(v.X) * op), Y: int16(float64(v.Y) * op), Z: int16(float64(v.Z) * op)}
}

// DivN returns the vector with each component divided by a number.
func (v Vector3int16) DivN(op float64) Vector3int16 {
	return Vector3int16{X: int16(float64(v.X) / op), Y: int16(float64(v.Y) / op), Z: int16(float64(v.Z) / op)}
}

// Neg returns the negation of the vector.
func (v Vector3int16) Neg() Vector3int16 {
	return Vector3int16{X: -v.X, Y: -v.Y, Z: -v.Z}
}

// Type returns a string that identifies the type.
func (Vector3int16) Type() string {
	return "Vector3int16"
}

// String returns a human-readable string representation of the value.
func (v Vector3int16) String() string {
	var b []byte
	b = strconv.AppendInt(b, int64(v.X), 10)
	b = append(b, ", "...)
	b = strconv.AppendInt(b, int64(v.Y), 10)
	b = append(b, ", "...)
	b = strconv.AppendInt(b, int64(v.Z), 10)
	return string(b)
}

// Copy returns a copy of the value.
func (v Vector3int16) Copy() PropValue {
	return v
}
