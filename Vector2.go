package types

import (
	"math"
	"strconv"
)

type Vector2 struct {
	X, Y float32
}

func NewVector2(x, y float64) Vector2 {
	return Vector2{X: float32(x), Y: float32(y)}
}

func (v Vector2) Magnitude() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

func (v Vector2) Unit() Vector2 {
	m := float32(v.Magnitude())
	return Vector2{X: v.X / m, Y: v.Y / m}
}

func (v Vector2) Lerp(goal Vector2, alpha float64) Vector2 {
	a := float32(alpha)
	na := 1 - a
	return Vector2{
		X: na*v.X + a*goal.X,
		Y: na*v.Y + a*goal.Y,
	}
}

func (v Vector2) Dot(op Vector2) float64 {
	return float64(v.X*op.X + v.Y*op.Y)
}

func (v Vector2) Cross(op Vector2) float64 {
	return float64(v.X*op.Y - v.Y*op.X)
}

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

func (v Vector2) Add(op Vector2) Vector2 {
	return Vector2{X: v.X + op.X, Y: v.Y + op.Y}
}

func (v Vector2) Sub(op Vector2) Vector2 {
	return Vector2{X: v.X - op.X, Y: v.Y - op.Y}
}

func (v Vector2) Mul(op Vector2) Vector2 {
	return Vector2{X: v.X * op.X, Y: v.Y * op.Y}
}

func (v Vector2) Div(op Vector2) Vector2 {
	return Vector2{X: v.X / op.X, Y: v.Y / op.Y}
}

func (v Vector2) MulN(op float64) Vector2 {
	return Vector2{X: v.X * float32(op), Y: v.Y * float32(op)}
}

func (v Vector2) DivN(op float64) Vector2 {
	return Vector2{X: v.X / float32(op), Y: v.Y / float32(op)}
}

func (v Vector2) Neg() Vector2 {
	return Vector2{X: -v.X, Y: -v.Y}
}

// Type returns a string identifying the type.
func (Vector2) Type() string {
	return "Vector2"
}

func (v Vector2) String() string {
	var b []byte
	b = strconv.AppendFloat(b, float64(v.X), 'g', -1, 32)
	b = append(b, ", "...)
	b = strconv.AppendFloat(b, float64(v.Y), 'g', -1, 32)
	return string(b)
}

func (v Vector2) Copy() PropValue {
	return v
}
