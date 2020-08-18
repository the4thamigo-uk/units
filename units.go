package units

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type (
	unit map[string]int

	Unit interface {
		String() string
		Equal(Unit) bool
		Subs(map[string]Unit) Unit
		Validate(map[string]Unit) error
		Invert() Unit
		Multiply(Unit) Unit
		Divide(Unit) Unit
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
	out := unit{}
	for k1, v1 := range u {
		v2, ok := us[k1]
		if !ok {
			out[k1] += v1
			continue
		}

		for k3, v3 := range v2.(unit) {
			out[k3] += v1 * v3
		}
	}
	if !u.Equal(out) {
		// some substitutions were made so reprocess
		return out.Subs(us)
	}

	return out
}

func (u unit) Validate(us map[string]Unit) error {
	for k1 := range u {
		_, ok := us[k1]
		if !ok {
			return fmt.Errorf("the unit '%s' is not defined", k1)
		}
	}
	return nil
}

func (u unit) clear() {
	for k := range u {
		delete(u, k)
	}
}

func (u unit) Equal(u2 Unit) bool {
	return reflect.DeepEqual(u, u2)
}

func (u unit) String() string {
	var ss []string
	for k, v := range u {
		if v == 1 {
			ss = append(ss, k)
			continue
		}
		ss = append(ss, fmt.Sprintf("%s^%d", k, v))
	}
	sort.Strings(ss)
	return strings.Join(ss, "*")
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
