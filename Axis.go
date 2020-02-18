package types

// Axis represents an orthographic axis in a 3-dimensional coordinate system.
type Axis int8

const (
	_ Axis = iota - 1
	AxisX
	AxisY
	AxisZ
)

// Valid returns whether the value is valid.
func (axis Axis) Valid() bool {
	return AxisX <= axis && axis <= AxisZ
}

// Face returns a face from the axis and a given direction. Returns Invalid if
// axis is invalid, or if direction is 0.
func (axis Axis) Face(direction float64) Face {
	if direction > 0 {
		switch axis {
		case AxisX:
			return FaceRight
		case AxisY:
			return FaceTop
		case AxisZ:
			return FaceBack
		}
	} else if direction < 0 {
		switch axis {
		case AxisX:
			return FaceLeft
		case AxisY:
			return FaceBottom
		case AxisZ:
			return FaceFront
		}
	}
	return Invalid
}
