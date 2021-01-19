package types

import (
	"strconv"
)

// PhysicalProperties represent the physical properties of an object.
type PhysicalProperties struct {
	CustomPhysics    bool
	Density          float32
	Friction         float32
	Elasticity       float32
	FrictionWeight   float32
	ElasticityWeight float32
}

// Type returns a string that identifies the type.
func (PhysicalProperties) Type() string {
	return "PhysicalProperties"
}

// String returns a human-readable string representation of the value.
func (p PhysicalProperties) String() string {
	if !p.CustomPhysics {
		return "(default)"
	}
	var b []byte
	b = strconv.AppendFloat(b, float64(p.Density), 'g', -1, 32)
	b = append(b, ", "...)
	b = strconv.AppendFloat(b, float64(p.Friction), 'g', -1, 32)
	b = append(b, ", "...)
	b = strconv.AppendFloat(b, float64(p.Elasticity), 'g', -1, 32)
	b = append(b, ", "...)
	b = strconv.AppendFloat(b, float64(p.FrictionWeight), 'g', -1, 32)
	b = append(b, ", "...)
	b = strconv.AppendFloat(b, float64(p.ElasticityWeight), 'g', -1, 32)
	return string(b)
}

// Copy returns a copy of the value.
func (p PhysicalProperties) Copy() PropValue {
	return p
}
