package types

import (
	"strconv"
)

type Int int32

func (Int) Type() string {
	return "int"
}

func (i Int) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Int) Copy() PropValue {
	return i
}

func (i Int) Numberlike() float64 {
	return float64(i)
}

func (i Int) Intlike() int64 {
	return int64(i)
}
