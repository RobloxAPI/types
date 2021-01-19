package types

import (
	"math"
	"strconv"
)

// Vector2 is a three-dimensional Euclidean vector.
type Vector3 struct {
	X, Y, Z float32
}

// NewVector3 returns a vector initialized with the given components.
func NewVector3(x, y, z float64) Vector3 {
	return Vector3{X: float32(x), Y: float32(y), Z: float32(z)}
}

// NewVector3FromFace returns the vector that corresponds to the given Face.
func NewVector3FromFace(normal Face) Vector3 {
	switch normal {
	case FaceRight:
		return Vector3{X: 1, Y: 0, Z: 0}
	case FaceTop:
		return Vector3{X: 0, Y: 1, Z: 0}
	case FaceBack:
		return Vector3{X: 0, Y: 0, Z: 1}
	case FaceLeft:
		return Vector3{X: -1, Y: 0, Z: 0}
	case FaceBottom:
		return Vector3{X: 0, Y: -1, Z: 0}
	case FaceFront:
		return Vector3{X: 0, Y: 0, Z: -1}
	}
	return Vector3{}
}

// NewVector3FromAxis returns the vector that corresponds to the given Axis.
func NewVector3FromAxis(axis Axis) Vector3 {
	switch axis {
	case AxisX:
		return Vector3{X: 1, Y: 0, Z: 0}
	case AxisY:
		return Vector3{X: 0, Y: 1, Z: 0}
	case AxisZ:
		return Vector3{X: 0, Y: 0, Z: 1}
	}
	return Vector3{}
}

// Magnitude returns the length of the vector.
func (v Vector3) Magnitude() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
}

// Unit returns the vector with the same direction and a length of 1.
func (v Vector3) Unit() Vector3 {
	m := float32(v.Magnitude())
	return Vector3{X: v.X / m, Y: v.Y / m, Z: v.Z / m}
}

// Lerp returns a Vector3 linearly interpolated from the Vector3 to *goal*
// according to *alpha*, which has an interval of [0, 1].
func (v Vector3) Lerp(goal Vector3, alpha float64) Vector3 {
	a := float32(alpha)
	na := 1 - a
	return Vector3{
		X: na*v.X + a*goal.X,
		Y: na*v.Y + a*goal.Y,
		Z: na*v.Z + a*goal.Z,
	}
}

// Dot returns the dot product of two vectors.
func (v Vector3) Dot(op Vector3) float64 {
	return float64(v.X*op.X + v.Y*op.Y + v.Z*op.Z)
}

// Cross returns the cross product of two vectors.
func (v Vector3) Cross(op Vector3) Vector3 {
	return Vector3{
		v.Y*op.Z - v.Z*op.Y,
		v.Z*op.X - v.X*op.Z,
		v.X*op.Y - v.Y*op.X,
	}
}

// FuzzyEq returns whether the two vectors are approximately equal.
func (v Vector3) FuzzyEq(op Vector3, epsilon float64) bool {
	switch {
	case epsilon == 0:
		return v == op
	case epsilon < 0:
		// Default.
		epsilon = 1e-5
	}
	x := v.X - op.X
	y := v.Y - op.Y
	z := v.Z - op.Z
	return x*x+y*y+z*z <= float32(epsilon)
}

// Add returns the sum of two vectors.
func (v Vector3) Add(op Vector3) Vector3 {
	return Vector3{X: v.X + op.X, Y: v.Y + op.Y, Z: v.Z + op.Z}
}

// Sub returns the difference of two vectors.
func (v Vector3) Sub(op Vector3) Vector3 {
	return Vector3{X: v.X - op.X, Y: v.Y - op.Y, Z: v.Z - op.Z}
}

// Mul returns the product of two vectors.
func (v Vector3) Mul(op Vector3) Vector3 {
	return Vector3{X: v.X * op.X, Y: v.Y * op.Y, Z: v.Z * op.Z}
}

// Div returns the quotient of two vectors.
func (v Vector3) Div(op Vector3) Vector3 {
	return Vector3{X: v.X / op.X, Y: v.Y / op.Y, Z: v.Z / op.Z}
}

// MulN returns the vector with each component multiplied by a number.
func (v Vector3) MulN(op float64) Vector3 {
	return Vector3{X: v.X * float32(op), Y: v.Y * float32(op), Z: v.Z * float32(op)}
}

// DivN returns the vector with each component divided by a number.
func (v Vector3) DivN(op float64) Vector3 {
	return Vector3{X: v.X / float32(op), Y: v.Y / float32(op), Z: v.Z / float32(op)}
}

// Neg returns the negation of the vector.
func (v Vector3) Neg() Vector3 {
	return Vector3{X: -v.X, Y: -v.Y, Z: -v.Z}
}

// Type returns a string that identifies the type.
func (Vector3) Type() string {
	return "Vector3"
}

// String returns a human-readable string representation of the value.
func (v Vector3) String() string {
	var b []byte
	b = strconv.AppendFloat(b, float64(v.X), 'g', -1, 32)
	b = append(b, ", "...)
	b = strconv.AppendFloat(b, float64(v.Y), 'g', -1, 32)
	b = append(b, ", "...)
	b = strconv.AppendFloat(b, float64(v.Z), 'g', -1, 32)
	return string(b)
}

// Copy returns a copy of the value.
func (v Vector3) Copy() PropValue {
	return v
}
