package types

// Axes represents a set of coordinate axes that are considered active.
type Axes struct {
	X, Y, Z bool
}

// NewAxesFromAxis returns an Axes where the axes corresponding to the given
// Axis values are active.
func NewAxesFromAxis(axes ...Axis) Axes {
	a := Axes{}
	for _, axis := range axes {
		switch axis {
		case AxisX:
			a.X = true
		case AxisY:
			a.Y = true
		case AxisZ:
			a.Z = true
		}
	}
	return a
}

// NewAxesFromFace returns an Axes where the axes corresponding to the given
// Face values are active.
func NewAxesFromFace(faces ...Face) Axes {
	a := Axes{}
	for _, face := range faces {
		switch face {
		case FaceRight, FaceLeft:
			a.X = true
		case FaceTop, FaceBottom:
			a.Y = true
		case FaceBack, FaceFront:
			a.Z = true
		}
	}
	return a
}

// Face returns whether the axis corresponding to the given Face is active.
func (a Axes) Face(face Face) bool {
	switch face {
	case FaceRight, FaceLeft:
		return a.X
	case FaceTop, FaceBottom:
		return a.Y
	case FaceBack, FaceFront:
		return a.Z
	}
	return false
}
