package types

import (
	"strconv"
)

type Bool bool

const True = Bool(true)
const False = Bool(false)

func (Bool) Type() string {
	return "bool"
}

func (b Bool) String() string {
	return strconv.FormatBool(bool(b))
}

func (b Bool) Copy() PropValue {
	return b
}
