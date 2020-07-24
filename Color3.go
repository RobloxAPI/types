package types

import (
	"strconv"
)

type Color3 struct {
	R, G, B float32
}

func NewColor3FromRGB(r, g, b int) Color3 {
	return Color3{
		R: float32(r) / 255,
		G: float32(g) / 255,
		B: float32(b) / 255,
	}
}

func NewColor3FromHSV(h, s, v float64) Color3 {
	panic("not implemented")
}

func (c Color3) Lerp(goal Color3, alpha float64) Color3 {
	panic("not implemented")
}

func (c Color3) ToHSV() (h, s, v float64) {
	panic("not implemented")
}

// Type returns a string identifying the type.
func (Color3) Type() string {
	return "Color3"
}

func (c Color3) String() string {
	var b []byte
	b = strconv.AppendFloat(b, float64(c.R), 'g', -1, 32)
	b = append(b, ", "...)
	b = strconv.AppendFloat(b, float64(c.G), 'g', -1, 32)
	b = append(b, ", "...)
	b = strconv.AppendFloat(b, float64(c.B), 'g', -1, 32)
	return string(b)
}

func (c Color3) Copy() PropValue {
	return c
}
