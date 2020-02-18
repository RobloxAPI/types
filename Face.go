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
