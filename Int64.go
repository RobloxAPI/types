package types

import (
	"strconv"
)

type Int64 int64

func (Int64) Type() string {
	return "int64"
}

func (i Int64) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Int64) Copy() PropValue {
	return i
}

func (i Int64) Numberlike() float64 {
	return float64(i)
}

func (i Int64) Intlike() int64 {
	return int64(i)
}
