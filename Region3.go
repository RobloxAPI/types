package types

import (
	"math"
)

type Region3 struct {
	Min, Max Vector3
}

func (r Region3) CFrame() CFrame {
	cf := NewCFrame()
	cf.Position = r.Min.Add(r.Max).DivN(2)
	return cf
}

func (r Region3) Size() Vector3 {
	return r.Max.Sub(r.Min)
}

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

func rafz3(v Vector3) Vector3 {
	v.X = sign(v.X) * float32(math.Ceil(math.Abs(float64(v.X))))
	v.Y = sign(v.Y) * float32(math.Ceil(math.Abs(float64(v.Y))))
	v.Z = sign(v.Z) * float32(math.Ceil(math.Abs(float64(v.Z))))
	return v
}

func (r Region3) ExpandToGrid(res float64) Region3 {
	if res == 0 {
		return r
	}
	r.Min = rafz3(r.Min.DivN(res)).MulN(res)
	r.Max = rafz3(r.Max.DivN(res)).MulN(res)
	return r
}

// Type returns a string identifying the type.
func (Region3) Type() string {
	return "Region3"
}

func (r Region3) String() string {
	var b []byte
	b = append(b, r.Min.String()...)
	b = append(b, "; "...)
	b = append(b, r.Max.String()...)
	return string(b)
}

func (r Region3) Copy() PropValue {
	return r
}
