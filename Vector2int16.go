package types

import (
	"strconv"
)

type Vector2int16 struct {
	X, Y int16
}

func NewVector2int16(x, y int) Vector2int16 {
	return Vector2int16{X: int16(x), Y: int16(y)}
}

func (v Vector2int16) Add(op Vector2int16) Vector2int16 {
	return Vector2int16{X: v.X + op.X, Y: v.Y + op.Y}
}

func (v Vector2int16) Sub(op Vector2int16) Vector2int16 {
	return Vector2int16{X: v.X - op.X, Y: v.Y - op.Y}
}

func (v Vector2int16) Mul(op Vector2int16) Vector2int16 {
	return Vector2int16{X: v.X * op.X, Y: v.Y * op.Y}
}

func (v Vector2int16) Div(op Vector2int16) Vector2int16 {
	return Vector2int16{X: v.X / op.X, Y: v.Y / op.Y}
}

func (v Vector2int16) MulN(op float64) Vector2int16 {
	return Vector2int16{X: int16(float64(v.X) * op), Y: int16(float64(v.Y) * op)}
}

func (v Vector2int16) DivN(op float64) Vector2int16 {
	return Vector2int16{X: int16(float64(v.X) / op), Y: int16(float64(v.Y) / op)}
}

func (v Vector2int16) Neg() Vector2int16 {
	return Vector2int16{X: -v.X, Y: -v.Y}
}

// Type returns a string identifying the type.
func (Vector2int16) Type() string {
	return "Vector2int16"
}

func (v Vector2int16) String() string {
	var b []byte
	b = strconv.AppendInt(b, int64(v.X), 10)
	b = append(b, ", "...)
	b = strconv.AppendInt(b, int64(v.Y), 10)
	return string(b)
}
