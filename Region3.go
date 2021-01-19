package types

import (
	"math"
)

// Region3 represents an axis-aligned three-dimensional rectangular cuboid with
// a lower and upper boundary.
type Region3 struct {
	Min, Max Vector3
}

// CFrame returns a CFrame with the default orientation, located at the center
// of the region.
func (r Region3) CFrame() CFrame {
	cf := NewCFrame()
	cf.Position = r.Min.Add(r.Max).DivN(2)
	return cf
}

// Size returns the size of the region.
func (r Region3) Size() Vector3 {
	return r.Max.Sub(r.Min)
}

// sign returns the sign of x.
func sign(x float32) float32 {
	switch {
	case x < 0:
		return -1
	case x > 0:
		return 1
	default:
		return 0
	}
}

// rafz3 returns v with each component rounded, away from zero, to the nearest
// integer.
func rafz3(v Vector3) Vector3 {
	v.X = sign(v.X) * float32(math.Ceil(math.Abs(float64(v.X))))
	v.Y = sign(v.Y) * float32(math.Ceil(math.Abs(float64(v.Y))))
	v.Z = sign(v.Z) * float32(math.Ceil(math.Abs(float64(v.Z))))
	return v
}

// ExpandToGrid returns the region expanded to align to the given resolution. If
// *res* is 0, the region is returned unchanged.
func (r Region3) ExpandToGrid(res float64) Region3 {
	if res == 0 {
		return r
	}
	r.Min = rafz3(r.Min.DivN(res)).MulN(res)
	r.Max = rafz3(r.Max.DivN(res)).MulN(res)
	return r
}

// Type returns a string that identifies the type.
func (Region3) Type() string {
	return "Region3"
}

// String returns a human-readable string representation of the value.
func (r Region3) String() string {
	var b []byte
	b = append(b, r.Min.String()...)
	b = append(b, "; "...)
	b = append(b, r.Max.String()...)
	return string(b)
}

// Copy returns a copy of the value.
func (r Region3) Copy() PropValue {
	return r
}
