package types

import (
	"strconv"
)

type Vector3int16 struct {
	X, Y, Z int16
}

func NewVector3int16(x, y, z int) Vector3int16 {
	return Vector3int16{X: int16(x), Y: int16(y), Z: int16(z)}
}

func (v Vector3int16) Add(op Vector3int16) Vector3int16 {
	return Vector3int16{X: v.X + op.X, Y: v.Y + op.Y, Z: v.Z + op.Z}
}

func (v Vector3int16) Sub(op Vector3int16) Vector3int16 {
	return Vector3int16{X: v.X - op.X, Y: v.Y - op.Y, Z: v.Z - op.Z}
}

func (v Vector3int16) Mul(op Vector3int16) Vector3int16 {
	return Vector3int16{X: v.X * op.X, Y: v.Y * op.Y, Z: v.Z * op.Z}
}

func (v Vector3int16) Div(op Vector3int16) Vector3int16 {
	return Vector3int16{X: v.X / op.X, Y: v.Y / op.Y, Z: v.Z / op.Z}
}

func (v Vector3int16) MulN(op float64) Vector3int16 {
	return Vector3int16{X: int16(float64(v.X) * op), Y: int16(float64(v.Y) * op), Z: int16(float64(v.Z) * op)}
}

func (v Vector3int16) DivN(op float64) Vector3int16 {
	return Vector3int16{X: int16(float64(v.X) / op), Y: int16(float64(v.Y) / op), Z: int16(float64(v.Z) / op)}
}

func (v Vector3int16) Neg() Vector3int16 {
	return Vector3int16{X: -v.X, Y: -v.Y, Z: -v.Z}
}

// Type returns a string identifying the type.
func (Vector3int16) Type() string {
	return "Vector3int16"
}

func (v Vector3int16) String() string {
	var b []byte
	b = strconv.AppendInt(b, int64(v.X), 10)
	b = append(b, ", "...)
	b = strconv.AppendInt(b, int64(v.Y), 10)
	b = append(b, ", "...)
	b = strconv.AppendInt(b, int64(v.Z), 10)
	return string(b)
}
