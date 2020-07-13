package types

type Ray struct {
	Origin    Vector3
	Direction Vector3
}

func (r Ray) ClosestPoint(point Vector3) Vector3 {
	panic("not implemented")
}

func (r Ray) Distance(point Vector3) float64 {
	panic("not implemented")
}
