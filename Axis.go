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
