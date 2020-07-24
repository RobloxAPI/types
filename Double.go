package types

import (
	"strconv"
)

type Double float64

func (Double) Type() string {
	return "double"
}

func (d Double) String() string {
	return strconv.FormatFloat(float64(d), 'g', -1, 64)
}

func (d Double) Copy() PropValue {
	return d
}

func (d Double) Numberlike() float64 {
	return float64(d)
}

func (d Double) Intlike() int64 {
	return int64(d)
}
