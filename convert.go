package units

import (
	"fmt"
)

type (
	converter float64
	Converter interface {
		Convert(float64) float64
		ConvertPtr(*float64) *float64
	}
)

func NewConverter(from Unit, to Unit) (Converter, error) {
	scale, err := isConvertible(from, to)
	if err != nil {
		return nil, err
	}
	return converter(scale), nil
}

func MustConvert(c Converter, err error) Converter {
	if err != nil {
		panic(err.Error())
	}
	return c
}

func isConvertible(from Unit, to Unit) (float64, error) {
	u := from.Divide(to)
	if !u.IsScalar() {
		return 0, fmt.Errorf("cannot convert units from '%s' to '%s'", from, to)
	}
	return u.Scale(), nil
}

func (c converter) Convert(val float64) float64 {
	return float64(c) * val
}

func (c converter) ConvertPtr(val *float64) *float64 {
	if val == nil {
		return nil
	}
	val2 := c.Convert(*val)
	return &val2
}
