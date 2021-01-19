package types

import (
	"math"
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

func NewColor3FromHSV(h, s, v float64) (c Color3) {
	var k float64
	k = math.Mod(5+h*6, 6)
	c.R = float32(v - v*s*math.Max(0, math.Min(math.Min(k, 4-k), 1)))
	k = math.Mod(3+h*6, 6)
	c.G = float32(v - v*s*math.Max(0, math.Min(math.Min(k, 4-k), 1)))
	k = math.Mod(1+h*6, 6)
	c.B = float32(v - v*s*math.Max(0, math.Min(math.Min(k, 4-k), 1)))
	return c
}

func (c Color3) Lerp(goal Color3, alpha float64) Color3 {
	a := float32(alpha)
	na := 1 - a
	return Color3{
		R: na*c.R + a*goal.R,
		G: na*c.G + a*goal.G,
		B: na*c.B + a*goal.B,
	}
}

func (c Color3) ToHSV() (h, s, v float64) {
	var r, g, b = float64(c.R), float64(c.G), float64(c.B)
	min := math.Min(math.Min(r, g), b)
	max := math.Max(math.Max(r, g), b)
	delta := max - min
	v = max
	if max != 0 {
		s = delta / max
	} else {
		s = 0
		h = -1
	}
	switch max {
	case r:
		h = ((g - b) / delta) / 6
	case g:
		h = (2 + (b-r)/delta) / 6
	case b:
		h = (4 + (r-g)/delta) / 6
	}
	if h <= 0 {
		h++
	}
	return h, s, v
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
