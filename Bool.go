package types

import (
	"strconv"
)

type Bool bool

func (Bool) Type() string {
	return "bool"
}

func (b Bool) String() string {
	return strconv.FormatBool(bool(b))
}

func (b Bool) Copy() PropValue {
	return b
}
