package types

import (
	"strconv"
)

type Float float32

func (Float) Type() string {
	return "float"
}

func (f Float) String() string {
	return strconv.FormatFloat(float64(f), 'g', -1, 32)
}

func (f Float) Copy() PropValue {
	return f
}

func (f Float) Numberlike() float64 {
	return float64(f)
}

func (f Float) Intlike() int64 {
	return int64(f)
}
