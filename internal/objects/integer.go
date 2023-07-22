package objects

import (
	"strconv"
)

type Integer struct {
	Value int
}

func NewInteger(value int) *Integer {
	return &Integer{Value: value}
}

func (*Integer) Type() ObjectType {
	return IntegerObject
}

func (i *Integer) Bytes() []byte {
	return []byte(strconv.Itoa(i.Value))
}

func (i *Integer) Add(other *Integer) *Integer {
	return NewInteger(i.Value + other.Value)
}

func (i *Integer) Sub(other *Integer) *Integer {
	return NewInteger(i.Value - other.Value)
}

func (i *Integer) Mul(other *Integer) *Integer {
	return NewInteger(i.Value * other.Value)
}

func (i *Integer) Div(other *Integer) *Integer {
	return NewInteger(i.Value / other.Value)
}
