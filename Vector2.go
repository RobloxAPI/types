package types

import (
	"math"
	"strconv"
)

// Vector2 is a two-dimensional Euclidean vector.
type Vector2 struct {
	X, Y float32
}

// NewVector2 returns a vector initialized with the given components.
func NewVector2(x, y float64) Vector2 {
	return Vector2{X: float32(x), Y: float32(y)}
}

// Magnitude returns the length of the vector.
func (v Vector2) Magnitude() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

// Unit returns the vector with the same direction and a length of 1.
func (v Vector2) Unit() Vector2 {
	m := float32(v.Magnitude())
	return Vector2{X: v.X / m, Y: v.Y / m}
}

// Lerp returns a Vector2 linearly interpolated from the Vector2 to *goal*
// according to *alpha*, which has an interval of [0, 1].
func (v Vector2) Lerp(goal Vector2, alpha float64) Vector2 {
	a := float32(alpha)
	na := 1 - a
	return Vector2{
		X: na*v.X + a*goal.X,
		Y: na*v.Y + a*goal.Y,
	}
}

// Dot returns the dot product of two vectors.
func (v Vector2) Dot(op Vector2) float64 {
	return float64(v.X*op.X + v.Y*op.Y)
}

// Cross returns the cross product of two vectors extended into three dimensions
// with z components of 0.
func (v Vector2) Cross(op Vector2) float64 {
	return float64(v.X*op.Y - v.Y*op.X)
}

// FuzzyEq returns whether the two vectors are approximately equal.
func (v Vector2) FuzzyEq(op Vector2, epsilon float64) bool {
	switch {
	case epsilon == 0:
		return v == op
	case epsilon < 0:
		// Default.
		epsilon = 1e-5
	}
	x := v.X - op.X
	y := v.Y - op.Y
	return x*x+y*y <= float32(epsilon)
}

// Add returns the sum of two vectors.
func (v Vector2) Add(op Vector2) Vector2 {
	return Vector2{X: v.X + op.X, Y: v.Y + op.Y}
}

// Sub returns the difference of two vectors.
func (v Vector2) Sub(op Vector2) Vector2 {
	return Vector2{X: v.X - op.X, Y: v.Y - op.Y}
}

// Mul returns the product of two vectors.
func (v Vector2) Mul(op Vector2) Vector2 {
	return Vector2{X: v.X * op.X, Y: v.Y * op.Y}
}

// Div returns the quotient of two vectors.
func (v Vector2) Div(op Vector2) Vector2 {
	return Vector2{X: v.X / op.X, Y: v.Y / op.Y}
}

// MulN returns the vector with each component multiplied by a number.
func (v Vector2) MulN(op float64) Vector2 {
	return Vector2{X: v.X * float32(op), Y: v.Y * float32(op)}
}

// DivN returns a Vector2 with each component divided by a number.
func (v Vector2) DivN(op float64) Vector2 {
	return Vector2{X: v.X / float32(op), Y: v.Y / float32(op)}
}

// Neg returns the negation of the vector.
func (v Vector2) Neg() Vector2 {
	return Vector2{X: -v.X, Y: -v.Y}
}

// Type returns a string that identifies the type.
func (Vector2) Type() string {
	return "Vector2"
}

// String returns a human-readable string representation of the value.
func (v Vector2) String() string {
	var b []byte
	b = strconv.AppendFloat(b, float64(v.X), 'g', -1, 32)
	b = append(b, ", "...)
	b = strconv.AppendFloat(b, float64(v.Y), 'g', -1, 32)
	return string(b)
}

// Copy returns a copy of the value.
func (v Vector2) Copy() PropValue {
	return v
}
