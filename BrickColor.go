package types

import (
	"math"
)

// BrickColor represents a color from a predefined collection.
type BrickColor uint32

func NewBrickColor(number int) BrickColor {
	if _, ok := bcIndex[BrickColor(number)]; ok {
		return BrickColor(number)
	}
	return BrickColorDefault
}

func NewBrickColorFromName(name string) BrickColor {
	if bc, ok := bcNameIndex[name]; ok {
		return bc
	}
	return BrickColorDefault
}

func NewBrickColorFromColor(r, g, b float64) BrickColor {
	return NewBrickColorFromColor3(Color3{
		R: float32(r),
		G: float32(g),
		B: float32(b),
	})
}

func NewBrickColorFromColor3(color Color3) BrickColor {
	index := bcDefaultIndex
	dist := float32(math.Inf(1))
	for i, c := range bcColors {
		r := math.Float32frombits(math.Float32bits(c.R-color.R) &^ (1 << 31))
		g := math.Float32frombits(math.Float32bits(c.G-color.G) &^ (1 << 31))
		b := math.Float32frombits(math.Float32bits(c.B-color.B) &^ (1 << 31))
		if d := r + g + b; d == 0 {
			return brickColors[i]
		} else if d < dist {
			dist, index = d, i
		}
	}
	return brickColors[index]
}

func _NewBrickColorFromColor3_euc(color Color3) BrickColor {
	// Euclidian distance.
	index := bcDefaultIndex
	dist := float32(math.Inf(1))
	for i, c := range bcColors {
		r := c.R - color.R
		g := c.G - color.G
		b := c.B - color.B
		if d := r*r + g*g + b*b; d == 0 {
			return brickColors[i]
		} else if d < dist {
			dist, index = d, i
		}
	}
	return brickColors[index]
}

func NewBrickColorFromPalette(index int) BrickColor {
	if index >= 0 && index < BrickColorPaletteSize {
		return bcPalette[index]
	}
	return BrickColorDefault
}

func NewBrickColorFromIndex(index int) BrickColor {
	if index >= 0 && index < BrickColorIndexSize {
		return brickColors[index]
	}
	return BrickColorDefault
}

func (b BrickColor) Index() int {
	if i, ok := bcIndex[b]; ok {
		return i
	}
	return bcDefaultIndex
}

func (b BrickColor) R() float64 {
	return float64(bcColors[b.Index()].R)
}

func (b BrickColor) G() float64 {
	return float64(bcColors[b.Index()].G)
}

func (b BrickColor) B() float64 {
	return float64(bcColors[b.Index()].B)
}

func (b BrickColor) Number() int {
	return int(b)
}

func (b BrickColor) Color() Color3 {
	return bcColors[b.Index()]
}

func (b BrickColor) Name() string {
	i := b.Index()
	return bcNames[bcNamesIndex[i]:bcNamesIndex[i+1]]
}
