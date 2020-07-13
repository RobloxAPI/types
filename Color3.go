package types

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
