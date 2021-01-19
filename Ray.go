package types

// Ray is a line extending infinitely in one direction.
type Ray struct {
	Origin    Vector3
	Direction Vector3
}

// ClosestPoint returns the position on the ray that is nearest to *point*.
func (r Ray) ClosestPoint(point Vector3) Vector3 {
	e := r.Direction.Dot(point.Sub(r.Origin))
	if e <= 0 {
		return r.Origin
	}
	return r.Direction.MulN(e).Add(r.Origin)
}

// Distance returns the distance between *point* and the point on the ray
// nearest to *point*.
func (r Ray) Distance(point Vector3) float64 {
	return r.ClosestPoint(point).Sub(r.Origin).Magnitude()
}

// Type returns a string that identifies the type.
func (Ray) Type() string {
	return "Ray"
}

// String returns a human-readable string representation of the value.
func (r Ray) String() string {
	var b []byte
	b = append(b, r.Origin.String()...)
	b = append(b, "; "...)
	b = append(b, r.Direction.String()...)
	return string(b)
}

// Copy returns a copy of the value.
func (r Ray) Copy() PropValue {
	return r
}
