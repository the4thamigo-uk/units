package units

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strings"
)

type (
	dims map[string]int
	unit struct {
		scale float64
		dims  dims
	}

	Unit interface {
		String() string
		Equal(Unit) bool
		Subs(map[string]Unit) Unit
		Validate(map[string]Unit) error
		Invert() Unit
		Multiply(Unit) Unit
		Divide(Unit) Unit
		Scale() float64
	}
)

func Scalar(scale float64) Unit {
	return makeUnit(scale, dims{})
}

func NewUnit(name string, dim int) Unit {
	if name == "" || dim == 0 {
		return Scalar(1)
	}
	return makeUnit(1.0, dims{name: dim})
}

func makeUnit(scale float64, dms dims) unit {
	if dms == nil {
		dms = dims{}
	}
	if scale == 0.0 {
		scale = 1.0
	}
	return unit{
		scale: scale,
		dims:  dms,
	}
}

func (u unit) Subs(us map[string]Unit) Unit {
	out := makeUnit(u.scale, nil)

	for k1, v1 := range u.dims {
		v2, ok := us[k1]
		if !ok {
			out.dims[k1] += v1
			continue
		}

		u2 := v2.(unit)
		out.scale *= math.Pow(u2.scale, float64(v1))
		for k3, v3 := range u2.dims {
			out.dims[k3] += v1 * v3
		}
	}
	if !u.Equal(out) {
		// some substitutions were made so reprocess
		return out.Subs(us)
	}

	return out
}

func (u unit) Validate(us map[string]Unit) error {
	for k1 := range u.dims {
		_, ok := us[k1]
		if !ok {
			return fmt.Errorf("the unit '%s' is not defined", k1)
		}
	}
	return nil
}

func (u unit) Equal(u2 Unit) bool {
	return u.Scale() == u2.Scale() && reflect.DeepEqual(u, u2)
}

func (u unit) String() string {
	var ss []string
	for k, v := range u.dims {
		if v == 1 {
			ss = append(ss, k)
			continue
		}
		ss = append(ss, fmt.Sprintf("%s^%d", k, v))
	}
	sort.Strings(ss)
	if u.scale == 1.0 {
		return strings.Join(ss, "*")
	}
	return fmt.Sprintf("%e*", u.scale) + strings.Join(ss, "*")
}

func (u unit) Invert() Unit {
	out := makeUnit(1.0/u.scale, nil)
	for k, v := range u.dims {
		out.dims[k] = -v
	}
	return out
}

func (u unit) Multiply(u2 Unit) Unit {
	out := makeUnit(u.scale*u2.Scale(), nil)
	for k, v := range u.dims {
		out.dims[k] = v
	}
	for k, v := range u2.(unit).dims {
		val := out.dims[k] + v
		if val == 0 {
			delete(out.dims, k)
			continue
		}
		out.dims[k] = val
	}
	return out
}

func (u unit) Divide(ut Unit) Unit {
	return u.Multiply(ut.Invert())
}

func (u unit) Scale() float64 {
	return u.scale
}
