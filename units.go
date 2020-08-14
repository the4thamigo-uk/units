package units

import (
	"fmt"
	"reflect"
)

type (
	unit map[string]int

	Unit interface {
		String() string
		Equal(Unit) bool
		Invert() Unit
		Multiply(Unit) Unit
		Divide(Unit) Unit
	}
)

func Scalar() Unit {
	return unit{}
}

func NewUnit(name string) Unit {
	return newUnit(name, 1)
}

func newUnit(name string, dim int) Unit {
	if name == "" || dim == 0 {
		return unit{}
	}
	return unit{
		name: dim,
	}
}

func (u unit) Equal(u2 Unit) bool {
	return reflect.DeepEqual(u, u2)
}

func (u unit) String() string {
	// todo: format this better
	return fmt.Sprintf("%v", map[string]int(u))
}

func (u unit) Invert() Unit {
	out := unit{}
	for k, v := range u {
		out[k] = -v
	}
	return out
}

func (u unit) Multiply(ut Unit) Unit {
	out := unit{}
	for k, v := range u {
		out[k] = v
	}
	for k, v := range ut.(unit) {
		val := out[k] + v
		if val == 0 {
			delete(out, k)
			continue
		}
		out[k] = val
	}
	return out
}

func (u unit) Divide(ut Unit) Unit {
	return u.Multiply(ut.Invert())
}
