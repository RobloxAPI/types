package types

import (
	"strconv"
)

type Token uint32

func (Token) Type() string {
	return "token"
}

func (t Token) String() string {
	return strconv.FormatUint(uint64(t), 10)
}

func (t Token) Copy() PropValue {
	return t
}

func (t Token) Numberlike() float64 {
	return float64(t)
}

func (t Token) Intlike() int64 {
	return int64(t)
}
