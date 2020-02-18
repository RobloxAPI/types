package types

// Face represents a 3-dimensional egocentric direction, each value
// corresponding to one of the 3 orthographic axes.
type Face int8

const (
	_          Face = iota - 1
	FaceRight       // +X
	FaceTop         // +Y
	FaceBack        // +Z
	FaceLeft        // -X
	FaceBottom      // -X
	FaceFront       // -Z
)

// Valid returns whether the value is valid.
func (face Face) Valid() bool {
	return FaceRight <= face && face <= FaceFront
}

// Axis returns the Axis corresponding to the face.
func (face Face) Axis() Axis {
	switch face {
	case FaceRight, FaceLeft:
		return AxisX
	case FaceTop, FaceBottom:
		return AxisY
	case FaceBack, FaceFront:
		return AxisZ
	}
	return Invalid
}

// Direction returns 1 with the sign corresponding to the direction of the face
// on its axis. Returns 0 for an invalid value.
func (face Face) Direction() float64 {
	switch face {
	case FaceRight, FaceTop, FaceBack:
		return 1
	case FaceLeft, FaceBottom, FaceFront:
		return -1
	}
	return 0
}
