package types

// Axes represents a set of coordinate axes that are considered active.
type Faces struct {
	Right, Top, Back, Left, Bottom, Front bool
}

// NewFacesFromAxis returns a Faces where the faces corresponding to the given
// Axis values are active.
func NewFacesFromAxis(axes ...Axis) Faces {
	f := Faces{}
	for _, axis := range axes {
		switch axis {
		case AxisX:
			f.Right = true
			f.Left = true
		case AxisY:
			f.Top = true
			f.Bottom = true
		case AxisZ:
			f.Back = true
			f.Front = true
		}
	}
	return f
}

// NewFacesFromFace returns a Faces where the faces corresponding to the given
// Face values are active.
func NewFacesFromFace(faces ...Face) Faces {
	f := Faces{}
	for _, face := range faces {
		switch face {
		case FaceRight:
			f.Right = true
		case FaceLeft:
			f.Left = true
		case FaceTop:
			f.Top = true
		case FaceBottom:
			f.Bottom = true
		case FaceBack:
			f.Back = true
		case FaceFront:
			f.Front = true
		}
	}
	return f
}
