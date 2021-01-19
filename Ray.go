package types

type Ray struct {
	Origin    Vector3
	Direction Vector3
}

func (r Ray) ClosestPoint(point Vector3) Vector3 {
	e := r.Direction.Dot(point.Sub(r.Origin))
	if e <= 0 {
		return r.Origin
	}
	return r.Direction.MulN(e).Add(r.Origin)
}

func (r Ray) Distance(point Vector3) float64 {
	return r.ClosestPoint(point).Sub(r.Origin).Magnitude()
}

// Type returns a string identifying the type.
func (Ray) Type() string {
	return "Ray"
}

func (r Ray) String() string {
	var b []byte
	b = append(b, r.Origin.String()...)
	b = append(b, "; "...)
	b = append(b, r.Direction.String()...)
	return string(b)
}

func (r Ray) Copy() PropValue {
	return r
}
