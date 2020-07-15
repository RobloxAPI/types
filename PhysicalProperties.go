package types

import (
	"strconv"
)

type PhysicalProperties struct {
	CustomPhysics    bool
	Density          float32
	Friction         float32
	Elasticity       float32
	FrictionWeight   float32
	ElasticityWeight float32
}

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
