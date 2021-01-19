package types

import (
	"math"
)

// BrickColor represents a color from a predefined collection.
type BrickColor uint32

// NewBrickColor returns the BrickColor that corresponds to the given numeric
// value.
//
// The the given number has no corresponding BrickColor, then BrickColorDefault
// is returned.
func NewBrickColor(number int) BrickColor {
	if _, ok := bcIndex[BrickColor(number)]; ok {
		return BrickColor(number)
	}
	return BrickColorDefault
}

// NewBrickColorFromName returns the BrickColor that corresponds the given name.
//
// The the given name has no corresponding BrickColor, then BrickColorDefault is
// returned.
func NewBrickColorFromName(name string) BrickColor {
	if bc, ok := bcNameIndex[name]; ok {
		return bc
	}
	return BrickColorDefault
}

// NewBrickColorFromColor returns the BrickColor that is nearest to the given
// color components. Each component is in the interval [0, 1].
func NewBrickColorFromColor(r, g, b float64) BrickColor {
	return NewBrickColorFromColor3(Color3{
		R: float32(r),
		G: float32(g),
		B: float32(b),
	})
}

// NewBrickColorFromColor3 returns the BrickColor that is nearest to the given
// color.
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

// _NewBrickColorFromColor3_euc returns the BrickColor that is nearest to the
// given color according to Euclidean distance.
func _NewBrickColorFromColor3_euc(color Color3) BrickColor {
	// Euclidean distance.
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

// NewBrickColorFromPalette returns the BrickColor that corresponds to the given
// palette index.
//
// If the index is less than 0 or greater than or equal to
// BrickColorPaletteSize, then BrickColorDefault is returned.
func NewBrickColorFromPalette(index int) BrickColor {
	if index >= 0 && index < BrickColorPaletteSize {
		return bcPalette[index]
	}
	return BrickColorDefault
}

// NewBrickColorFromIndex returns the BrickColor that corresponds to the given
// index.
//
// If the index is less than 0 or greater than or equal to BrickColorIndexSize,
// then BrickColorDefault is returned.
func NewBrickColorFromIndex(index int) BrickColor {
	if index >= 0 && index < BrickColorIndexSize {
		return brickColors[index]
	}
	return BrickColorDefault
}

// Index returns the index of the BrickColor.
func (b BrickColor) Index() int {
	if i, ok := bcIndex[b]; ok {
		return i
	}
	return bcDefaultIndex
}

// R returns the red component of the BrickColor's color.
func (b BrickColor) R() float64 {
	return float64(bcColors[b.Index()].R)
}

// G returns the green component of the BrickColor's color.
func (b BrickColor) G() float64 {
	return float64(bcColors[b.Index()].G)
}

// B returns the blue component of the BrickColor's color.
func (b BrickColor) B() float64 {
	return float64(bcColors[b.Index()].B)
}

// Number returns the numeric value that identifies the BrickColor.
func (b BrickColor) Number() int {
	return int(b)
}

// Color returns the color of the BrickColor as a Color3.
func (b BrickColor) Color() Color3 {
	return bcColors[b.Index()]
}

// Name returns the name of the BrickColor.
func (b BrickColor) Name() string {
	i := b.Index()
	return bcNames[bcNamesIndex[i]:bcNamesIndex[i+1]]
}

// Type returns a string that identifies the type.
func (BrickColor) Type() string {
	return "BrickColor"
}

// String returns a human-readable string representation of the value.
func (b BrickColor) String() string {
	return b.Name()
}

// Copy returns a copy of the value.
func (b BrickColor) Copy() PropValue {
	return b
}
