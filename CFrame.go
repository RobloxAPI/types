package types

import (
	"math"
	"strconv"
)

// CFrame is a matrix representing a combined position and orientation.
type CFrame struct {
	Position Vector3
	Rotation [9]float32
}

// cframeID is the identity matrix.
var cframeID = CFrame{
	Position: Vector3{X: 0, Y: 0, Z: 0},
	Rotation: [9]float32{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	},
}

// NewCFrame returns a CFrame with the default orientation located at the
// origin.
func NewCFrame() CFrame {
	return cframeID
}

// Constructors and methods adapted from
// https://raw.githubusercontent.com/EgoMoose/Articles/master/CFrames/cframe class.lua

// NewCFrameFromVector3 returns a CFrame with the default orientation located at
// *v*.
func NewCFrameFromVector3(v Vector3) CFrame {
	return CFrame{
		Position: v,
		Rotation: cframeID.Rotation,
	}
}

// NewCFrameFromLook returns a CFrame located at *pos*, facing towards *lookAt*.
// The local upward direction is (0, 1, 0).
func NewCFrameFromLook(pos, lookAt Vector3) CFrame {
	return NewCFrameFromLookAt(pos, lookAt, Vector3{0, 1, 0})
}

// NewCFrameFromLookAt returns a CFrame located at *pos*, facing towards
// *lookAt*, with the local upward direction determined by *up*.
func NewCFrameFromLookAt(pos, lookAt, up Vector3) CFrame {
	dir := lookAt.Sub(pos)
	right := dir.Cross(up)
	if right.Dot(right) > 0 {
		u := right.Cross(dir).Unit()
		b := dir.Unit().Neg()
		r := right.Unit()
		return CFrame{
			Position: pos,
			Rotation: [9]float32{
				r.X, u.X, b.X,
				r.Y, u.Y, b.Y,
				r.Z, u.Z, b.Z,
			},
		}
	} else if (Vector3{0, 1, 0}).Dot(dir) > 0 {
		return CFrame{
			Position: pos,
			Rotation: [9]float32{
				0, 1, 0,
				0, 0, -1,
				-1, 0, 0,
			},
		}
	} else {
		return CFrame{
			Position: pos,
			Rotation: [9]float32{
				0, 1, 0,
				0, 0, 1,
				1, 0, 0,
			},
		}
	}
}

// NewCFrameFromPosition returns a CFrame located at (x, y, z), with the default
// orientation.
func NewCFrameFromPosition(x, y, z float64) CFrame {
	return CFrame{
		Position: Vector3{X: float32(x), Y: float32(y), Z: float32(z)},
		Rotation: cframeID.Rotation,
	}
}

// NewCFrameFromQuat returns a CFrame located at (x, y, z), oriented according
// to the quaternion (qx, qy, qz, qw).
func NewCFrameFromQuat(x, y, z, qx, qy, qz, qw float64) CFrame {
	m := float32(1 / math.Sqrt(qx*qx+qy*qy+qz*qz+qw*qw))
	Qx, Qy, Qz, Qw := float32(qx)*m, float32(qy)*m, float32(qz)*m, float32(qw)*m
	return CFrame{
		Position: Vector3{X: float32(x), Y: float32(y), Z: float32(z)},
		Rotation: [9]float32{
			1 - 2*(Qy*Qy+Qz*Qz), 2 * (Qy*Qx - Qw*Qz), 2 * (Qw*Qy + Qz*Qx),
			2 * (Qw*Qz + Qy*Qx), 1 - 2*(Qx*Qx+Qz*Qz), 2 * (Qz*Qy - Qw*Qx),
			2 * (Qz*Qx - Qw*Qy), 2 * (Qw*Qx + Qz*Qy), 1 - 2*(Qx*Qx+Qy*Qy),
		},
	}
}

// NewCFrameFromComponents returns a CFrame located at (x, y, z), oriented
// according to the rotation matrix
//
//     [r00, r01, r02]
//     [r10, r11, r12]
//     [r20, r21, r22]
//
func NewCFrameFromComponents(x, y, z, r00, r01, r02, r10, r11, r12, r20, r21, r22 float64) CFrame {
	return CFrame{
		Position: Vector3{X: float32(x), Y: float32(y), Z: float32(z)},
		Rotation: [9]float32{
			float32(r00), float32(r01), float32(r02),
			float32(r10), float32(r11), float32(r12),
			float32(r20), float32(r21), float32(r22),
		},
	}
}

// NewCFrameFromAngles returns a CFrame located at the origin, oriented
// according to the angles (rx, ry, rz), in radians. Rotations are ordered Z, Y,
// X.
func NewCFrameFromAngles(rx, ry, rz float64) CFrame {
	cx, sx := float32(math.Cos(rx)), float32(math.Sin(rx))
	cy, sy := float32(math.Cos(ry)), float32(math.Sin(ry))
	cz, sz := float32(math.Cos(rz)), float32(math.Sin(rz))
	return CFrame{
		Rotation: [9]float32{
			cy * cz, -cy * sz, sy,
			sx*sy*cz + cx*sz, -sx*sy*sz + cx*cz, -sx * cy,
			-cx*sy*cz + sx*sz, cx*sy*sz + sx*cz, cx * cy,
		},
	}
}

// NewCFrameFromOrientation returns a CFrame located at the origin, oriented
// according to the angles (rx, ry, rz), in radians. Rotations are ordered Y, X,
// Z.
func NewCFrameFromOrientation(rx, ry, rz float64) CFrame {
	cx, sx := float32(math.Cos(rx)), float32(math.Sin(rx))
	cy, sy := float32(math.Cos(ry)), float32(math.Sin(ry))
	cz, sz := float32(math.Cos(rz)), float32(math.Sin(rz))
	return CFrame{
		Rotation: [9]float32{
			cy*cz + sy*sx*sz, -cy*sz + sy*sx*cz, sy * cx,
			cx * sz, cx * cz, -sx,
			-sy*cz + cy*sx*sz, sy*sz + cy*sx*cz, cy * cx,
		},
	}
}

// NewCFrameFromAxisAngle returns a CFrame located at the origin, rotated *r*
// radians around *axis*.
func NewCFrameFromAxisAngle(axis Vector3, rotation float64) CFrame {
	axis = axis.Unit()
	c, s := float32(math.Cos(rotation)), float32(math.Sin(rotation))
	return CFrame{
		Rotation: [9]float32{
			c + (1-c)*axis.X*axis.X, -s*axis.Z + (1-c)*axis.X*axis.Y, s*axis.Y + (1-c)*axis.X*axis.Z,
			s*axis.Z + (1-c)*axis.Y*axis.X, c + (1-c)*axis.Y*axis.Y, -s*axis.X + (1-c)*axis.Y*axis.Z,
			-s*axis.Y + (1-c)*axis.Z*axis.X, s*axis.X + (1-c)*axis.Z*axis.Y, c + (1-c)*axis.Z*axis.Z,
		},
	}
}

// NewCFrameFromMatrix returns a CFrame located at *p*, rotated according to the
// rotation matrix
//
//     [vx.X, vy.X, vz.X]
//     [vx.Y, vy.Y, vz.Y]
//     [vx.Z, vy.Z, vz.Z]
//
// If *vz* is the zero vector, then *vz* is calculated as the unit of the cross
// product of vx and vy.
func NewCFrameFromMatrix(p, vx, vy, vz Vector3) CFrame {
	if vz == (Vector3{}) {
		vz = vx.Cross(vy).Unit()
	}

	return CFrame{
		Position: p,
		Rotation: [9]float32{
			vx.X, vy.X, vz.X,
			vx.Y, vy.Y, vz.Y,
			vx.Z, vy.Z, vz.Z,
		},
	}
}

// X returns the X component of the Position.
func (c CFrame) X() float64 {
	return float64(c.Position.X)
}

// Y returns the Y component of the Position.
func (c CFrame) Y() float64 {
	return float64(c.Position.Y)
}

// Z returns the Z component of the Position.
func (c CFrame) Z() float64 {
	return float64(c.Position.Z)
}

// RightVector returns the right-direction, or first column of the rotation
// matrix.
func (c CFrame) RightVector() Vector3 {
	return Vector3{c.Rotation[0], c.Rotation[3], c.Rotation[6]}
}

// UpVector returns the up-direction, or second column of the rotation matrix.
func (c CFrame) UpVector() Vector3 {
	return Vector3{c.Rotation[1], c.Rotation[4], c.Rotation[7]}
}

// LookVector returns the forward-direction, or the negation of the third column
// of the rotation matrix.
func (c CFrame) LookVector() Vector3 {
	return Vector3{-c.Rotation[2], -c.Rotation[5], -c.Rotation[8]}
}

// XVector returns the first row of the rotation matrix.
func (c CFrame) XVector() Vector3 {
	return Vector3{c.Rotation[0], c.Rotation[1], c.Rotation[2]}
}

// YVector returns the second row the rotation matrix.
func (c CFrame) YVector() Vector3 {
	return Vector3{c.Rotation[3], c.Rotation[4], c.Rotation[5]}
}

// ZVector returns the third row of the rotation matrix.
func (c CFrame) ZVector() Vector3 {
	return Vector3{c.Rotation[6], c.Rotation[7], c.Rotation[8]}
}

// Inverse returns the inverse of the CFrame.
func (c CFrame) Inverse() CFrame {
	return CFrame{
		Position: Vector3{
			X: -(c.Rotation[0]*c.Position.X + c.Rotation[3]*c.Position.Y + c.Rotation[6]*c.Position.Z),
			Y: -(c.Rotation[1]*c.Position.X + c.Rotation[4]*c.Position.Y + c.Rotation[7]*c.Position.Z),
			Z: -(c.Rotation[2]*c.Position.X + c.Rotation[5]*c.Position.Y + c.Rotation[8]*c.Position.Z),
		},
		Rotation: [9]float32{
			c.Rotation[0], c.Rotation[3], c.Rotation[6],
			c.Rotation[1], c.Rotation[4], c.Rotation[7],
			c.Rotation[2], c.Rotation[5], c.Rotation[8],
		},
	}
}

// Lerp returns a CFrame linearly interpolated from the CFrame to *goal*
// according to *alpha*, which has an interval of [0, 1].
func (c CFrame) Lerp(goal CFrame, alpha float64) CFrame {
	p := c.Position.MulN(1 - alpha).Add(goal.Position.MulN(alpha))
	diff := c.Inverse().Mul(goal)
	axis, theta := diff.AxisAngle()
	c = c.Mul(NewCFrameFromAxisAngle(axis, theta*alpha))
	return CFrame{Position: p, Rotation: c.Rotation}
}

// ToWorldSpace returns the CFrame transformed from local to world space of
// *cf*.
func (c CFrame) ToWorldSpace(cf CFrame) CFrame {
	return c.Mul(cf)
}

// ToObjectSpace returns the CFrame transformed from world to local space of
// *cf*.
func (c CFrame) ToObjectSpace(cf CFrame) CFrame {
	return c.Inverse().Mul(cf)
}

// PointToWorldSpace returns a Vector3 transformed from local to world space of
// the CFrame.
func (c CFrame) PointToWorldSpace(v Vector3) Vector3 {
	return c.MulVec(v)
}

// PointToObjectSpace returns a Vector3 transformed from world to local space of
// the CFrame.
func (c CFrame) PointToObjectSpace(v Vector3) Vector3 {
	return c.Inverse().MulVec(v)
}

// VectorToWorldSpace returns a Vector3 rotated from local to world space of the
// CFrame.
func (c CFrame) VectorToWorldSpace(v Vector3) Vector3 {
	return CFrame{Rotation: c.Rotation}.MulVec(v)
}

// VectorToObjectSpace returns a Vector3 rotated from world to local space of
// the CFrame.
func (c CFrame) VectorToObjectSpace(v Vector3) Vector3 {
	return CFrame{Rotation: c.Rotation}.Inverse().MulVec(v)
}

// Components returns the components of the CFrame's position and rotation
// matrix.
func (c CFrame) Components() (x, y, z, r00, r01, r02, r10, r11, r12, r20, r21, r22 float64) {
	return float64(c.Position.X),
		float64(c.Position.Y),
		float64(c.Position.Z),
		float64(c.Rotation[0]),
		float64(c.Rotation[1]),
		float64(c.Rotation[2]),
		float64(c.Rotation[3]),
		float64(c.Rotation[4]),
		float64(c.Rotation[5]),
		float64(c.Rotation[6]),
		float64(c.Rotation[7]),
		float64(c.Rotation[8])
}

// Angles returns the approximate angles of the CFrame's orientation, in
// radians, if ordered Z, Y, X.
func (c CFrame) Angles() (rx, ry, rz float64) {
	rx = math.Atan2(-float64(c.Rotation[5]), float64(c.Rotation[8]))
	ry = math.Asin(float64(c.Rotation[2]))
	rz = math.Atan2(-float64(c.Rotation[1]), float64(c.Rotation[0]))
	return rx, ry, rz
}

// Orientation returns the approximate angles of the CFrame's orientation, in
// radians, if ordered Y, X, Z.
func (c CFrame) Orientation() (rx, ry, rz float64) {
	rx = math.Asin(-float64(c.Rotation[5]))
	ry = math.Atan2(float64(c.Rotation[2]), float64(c.Rotation[8]))
	rz = math.Atan2(float64(c.Rotation[3]), float64(c.Rotation[4]))
	return rx, ry, rz
}

// Quaternion returns the rotation matrix as a quaternion.
func (c CFrame) Quaternion() (qx, qy, qz, qw float64) {
	if c.Rotation[0]+c.Rotation[4]+c.Rotation[8] > 0 {
		qw = math.Sqrt(float64(1+c.Rotation[0]+c.Rotation[4]+c.Rotation[8])) * 0.5
		qx = float64((c.Rotation[7] - c.Rotation[5]) / (4 * float32(qw)))
		qy = float64((c.Rotation[2] - c.Rotation[6]) / (4 * float32(qw)))
		qz = float64((c.Rotation[3] - c.Rotation[1]) / (4 * float32(qw)))
	} else if c.Rotation[0] > c.Rotation[4] && c.Rotation[0] > c.Rotation[8] {
		qx = math.Sqrt(float64(1+c.Rotation[0]-c.Rotation[4]-c.Rotation[8])) * 0.5
		qy = float64((c.Rotation[3] + c.Rotation[1]) / (4 * float32(qx)))
		qz = float64((c.Rotation[6] + c.Rotation[2]) / (4 * float32(qx)))
		qw = float64((c.Rotation[7] - c.Rotation[5]) / (4 * float32(qx)))
	} else if c.Rotation[4] > c.Rotation[8] {
		qy = math.Sqrt(float64(1+c.Rotation[4]-c.Rotation[0]-c.Rotation[8])) * 0.5
		qx = float64((c.Rotation[3] + c.Rotation[1]) / (4 * float32(qy)))
		qz = float64((c.Rotation[7] + c.Rotation[5]) / (4 * float32(qy)))
		qw = float64((c.Rotation[2] - c.Rotation[6]) / (4 * float32(qy)))
	} else {
		qz = math.Sqrt(float64(1+c.Rotation[8]-c.Rotation[0]-c.Rotation[4])) * 0.5
		qx = float64((c.Rotation[6] + c.Rotation[2]) / (4 * float32(qz)))
		qy = float64((c.Rotation[7] + c.Rotation[5]) / (4 * float32(qz)))
		qw = float64((c.Rotation[3] - c.Rotation[1]) / (4 * float32(qz)))
	}
	return qw, qx, qy, qz
}

// AxisAngle returns the orientation of the CFrame as an angle, in radians,
// rotated around an axis.
func (c CFrame) AxisAngle() (axis Vector3, rotation float64) {
	qw, qx, qy, qz := c.Quaternion()
	if qw <= 0 {
		qw, qx, qy, qz = -qw, -qx, -qy, -qz
	}
	theta := math.Acos(qw) * 2
	axis = Vector3{float32(qx), float32(qy), float32(qz)}.DivN(math.Sin(theta * 0.5))
	if axis.Dot(axis) > 0 {
		return axis.Unit(), theta
	}
	return Vector3{1, 0, 0}, theta
}

// Mul returns the composition of two CFrames.
func (a CFrame) Mul(b CFrame) CFrame {
	a1 := Vector3{a.Rotation[0], a.Rotation[1], a.Rotation[2]}
	a2 := Vector3{a.Rotation[3], a.Rotation[4], a.Rotation[5]}
	a3 := Vector3{a.Rotation[6], a.Rotation[7], a.Rotation[8]}
	b1 := Vector3{b.Rotation[0], b.Rotation[3], b.Rotation[6]}
	b2 := Vector3{b.Rotation[1], b.Rotation[4], b.Rotation[7]}
	b3 := Vector3{b.Rotation[2], b.Rotation[5], b.Rotation[8]}
	return CFrame{
		Position: Vector3{
			float32(a1.Dot(b.Position)) + a.Position.X,
			float32(a2.Dot(b.Position)) + a.Position.Y,
			float32(a3.Dot(b.Position)) + a.Position.Z,
		},
		Rotation: [9]float32{
			float32(a1.Dot(b1)), float32(a1.Dot(b2)), float32(a1.Dot(b3)),
			float32(a2.Dot(b1)), float32(a2.Dot(b2)), float32(a2.Dot(b3)),
			float32(a3.Dot(b1)), float32(a3.Dot(b2)), float32(a3.Dot(b3)),
		},
	}
}

// MulVec returns *op* transformed from local to world space of the CFrame.
func (c CFrame) MulVec(op Vector3) Vector3 {
	return Vector3{
		float32(Vector3{c.Rotation[0], c.Rotation[1], c.Rotation[2]}.Dot(op)) + c.Position.X,
		float32(Vector3{c.Rotation[3], c.Rotation[4], c.Rotation[5]}.Dot(op)) + c.Position.Y,
		float32(Vector3{c.Rotation[6], c.Rotation[7], c.Rotation[8]}.Dot(op)) + c.Position.Z,
	}
}

// AddVec returns the CFrame translated in world space by *op*.
func (c CFrame) AddVec(op Vector3) CFrame {
	return CFrame{
		Position: c.Position.Add(op),
		Rotation: c.Rotation,
	}
}

// SubVec returns the CFrame translated in world space by the negation of *op*.
func (c CFrame) SubVec(op Vector3) CFrame {
	return CFrame{
		Position: c.Position.Sub(op),
		Rotation: c.Rotation,
	}
}

// Type returns a string that identifies the type.
func (CFrame) Type() string {
	return "CFrame"
}

// String returns a human-readable string representation of the value.
func (c CFrame) String() string {
	var b []byte
	b = strconv.AppendFloat(b, float64(c.Position.X), 'g', -1, 32)
	b = append(b, ", "...)
	b = strconv.AppendFloat(b, float64(c.Position.Y), 'g', -1, 32)
	b = append(b, ", "...)
	b = strconv.AppendFloat(b, float64(c.Position.Z), 'g', -1, 32)
	for _, r := range c.Rotation {
		b = append(b, ", "...)
		b = strconv.AppendFloat(b, float64(r), 'g', -1, 32)
	}
	return string(b)
}

// Copy returns a copy of the value.
func (c CFrame) Copy() PropValue {
	return c
}
