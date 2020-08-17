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
		Subs(map[string]Unit) Unit
		Invert() Unit
		Multiply(Unit) Unit
		Divide(Unit) Unit
		UnmarshalText(s []byte) error
	}
)

func Scalar() Unit {
	return unit{}
}

func NewUnit(name string, dim int) Unit {
	if name == "" || dim == 0 {
		return unit{}
	}
	return unit{
		name: dim,
	}
}

func (u unit) Subs(us map[string]Unit) Unit {
	var subs bool
	out := unit{}
	for k1, v1 := range u {
		v2, ok := us[k1]
		if !ok {
			out[k1] += v1
			continue
		}

		subs = true
		for k3, v3 := range v2.(unit) {
			out[k3] += v1 * v3
		}
		fmt.Printf("%v", out)
	}
	if subs {
		// some substitutions were made so reprocess
		return out.Subs(us)
	}

	return out
}

func (u unit) clear() {
	for k := range u {
		delete(u, k)
	}
}

func (u unit) UnmarshalText(s []byte) error {
	u2, err := Parse(string(s))
	if err != nil {
		return err
	}
	u.clear()
	for k, v := range u2.(unit) {
		u[k] = v
	}

	return nil
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
