package example

import (
	"github.com/the4thamigo-uk/units"
)

// units
var (
	scalarUnit    = units.NewUnit("")
	timeUnit      = units.NewUnit("s")
	lengthUnit    = units.NewUnit("m")
	speedUnit     = lengthUnit.Divide(timeUnit)
	frequencyUnit = scalarUnit.Divide(timeUnit)
	areaUnit      = lengthUnit.Multiply(lengthUnit)
)

// quantities
type (
	_scalar float64
	Scalar  interface {
		Value() float64
		Unit() units.Unit
		DivideTime(val Time) Frequency
	}

	_length float64
	Length  interface {
		Value() float64
		Unit() units.Unit
		MultiplyLength(val Length) Area
		MultiplyScalar(val Scalar) Length
		DivideTime(val Time) Speed
	}

	_time float64
	Time  interface {
		Value() float64
		Unit() units.Unit
		MultiplyScalar(val Scalar) Time
	}

	_speed float64
	Speed  interface {
		Value() float64
		Unit() units.Unit
		MultiplyTime(val Time) Length
	}

	_frequency float64
	Frequency  interface {
		Value() float64
		Unit() units.Unit
	}

	_area float64
	Area  interface {
		Value() float64
		Unit() units.Unit
	}
)

func NewScalar(val float64) Scalar {
	return _scalar(val)
}

func (q _scalar) Value() float64 {
	return float64(q)
}

func (q _scalar) Unit() units.Unit {
	return scalarUnit
}

func (q _scalar) DivideTime(val Time) Frequency {
	return NewFrequency(q.Value() / val.Value())
}

func NewLength(val float64) Length {
	return _length(val)
}

func (q _length) Value() float64 {
	return float64(q)
}

func (q _length) Unit() units.Unit {
	return lengthUnit
}

func (q _length) MultiplyLength(val Length) Area {
	return NewArea(q.Value() * val.Value())
}

func (q _length) MultiplyScalar(val Scalar) Length {
	return NewLength(q.Value() * val.Value())
}

func (q _length) DivideTime(val Time) Speed {
	return NewSpeed(q.Value() / val.Value())
}

func NewTime(val float64) Time {
	return _time(val)
}

func (q _time) Value() float64 {
	return float64(q)
}

func (q _time) Unit() units.Unit {
	return timeUnit
}

func (q _time) MultiplyScalar(val Scalar) Time {
	return NewTime(q.Value() * val.Value())
}

func NewSpeed(val float64) Speed {
	return _speed(val)
}

func (q _speed) Value() float64 {
	return float64(q)
}

func (q _speed) Unit() units.Unit {
	return speedUnit
}

func (q _speed) MultiplyTime(val Time) Length {
	return NewLength(q.Value() * val.Value())
}

func NewFrequency(val float64) Frequency {
	return _frequency(val)
}

func (q _frequency) Value() float64 {
	return float64(q)
}

func (q _frequency) Unit() units.Unit {
	return frequencyUnit
}

func NewArea(val float64) Area {
	return _area(val)
}

func (q _area) Value() float64 {
	return float64(q)
}

func (q _area) Unit() units.Unit {
	return areaUnit
}
