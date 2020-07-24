package types

import (
	"strconv"
)

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

// Type returns a string identifying the type.
func (Faces) Type() string {
	return "Faces"
}

// String returns a string representation of the Faces.
func (f Faces) String() string {
	var b []byte
	b = append(b, "Right:"...)
	b = strconv.AppendBool(b, f.Right)
	b = append(b, ", Top:"...)
	b = strconv.AppendBool(b, f.Top)
	b = append(b, ", Back:"...)
	b = strconv.AppendBool(b, f.Back)
	b = append(b, ", Left:"...)
	b = strconv.AppendBool(b, f.Left)
	b = append(b, ", Bottom:"...)
	b = strconv.AppendBool(b, f.Bottom)
	b = append(b, ", Front:"...)
	b = strconv.AppendBool(b, f.Front)
	return string(b)
}

func (f Faces) Copy() PropValue {
	return f
}
