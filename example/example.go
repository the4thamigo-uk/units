package example

import (
	"github.com/the4thamigo-uk/units"
)

// units
var (
	h  = units.Must(units.Parse("3.600000e+03*s"))
	km = units.Must(units.Parse("1.000000e+03*m"))
	m  = units.Must(units.Parse("m"))
	s  = units.Must(units.Parse("s"))
)

// quantities
type (
	_Area float64
	Area  interface {
		Value() float64
		Unit() units.Unit
	}

	_Distance float64
	Distance  interface {
		Value() float64
		Unit() units.Unit
		DivideDuration(val Duration) Speed
	}

	_Duration float64
	Duration  interface {
		Value() float64
		Unit() units.Unit
	}

	_Frequency float64
	Frequency  interface {
		Value() float64
		Unit() units.Unit
	}

	_Length float64
	Length  interface {
		Value() float64
		Unit() units.Unit
		MultiplyLength(val Length) Area
		MultiplyScalar(val Scalar) Length
	}

	_Scalar float64
	Scalar  interface {
		Value() float64
		Unit() units.Unit
		DivideTime(val Time) Frequency
	}

	_Speed float64
	Speed  interface {
		Value() float64
		Unit() units.Unit
		MultiplyDuration(val Duration) Distance
	}

	_Time float64
	Time  interface {
		Value() float64
		Unit() units.Unit
		MultiplyScalar(val Scalar) Time
	}
)

// quantity units
var (
	_unit_Area      = units.Must(units.Parse("m^2"))
	_unit_Distance  = units.Must(units.Parse("1.000000e+03*m"))
	_unit_Duration  = units.Must(units.Parse("3.600000e+03*s"))
	_unit_Frequency = units.Must(units.Parse("s^-1"))
	_unit_Length    = units.Must(units.Parse("m"))
	_unit_Scalar    = units.Must(units.Parse(""))
	_unit_Speed     = units.Must(units.Parse("2.777778e-01*m*s^-1"))
	_unit_Time      = units.Must(units.Parse("s"))
)

func NewArea(val float64) Area {
	return _Area(val)
}

func (q _Area) Value() float64 {
	return float64(q)
}

func (q _Area) Unit() units.Unit {
	return _unit_Area
}

func NewDistance(val float64) Distance {
	return _Distance(val)
}

func (q _Distance) Value() float64 {
	return float64(q)
}

func (q _Distance) Unit() units.Unit {
	return _unit_Distance
}

func (q _Distance) DivideDuration(val Duration) Speed {
	return NewSpeed(q.Value() / val.Value())
}

func NewDuration(val float64) Duration {
	return _Duration(val)
}

func (q _Duration) Value() float64 {
	return float64(q)
}

func (q _Duration) Unit() units.Unit {
	return _unit_Duration
}

func NewFrequency(val float64) Frequency {
	return _Frequency(val)
}

func (q _Frequency) Value() float64 {
	return float64(q)
}

func (q _Frequency) Unit() units.Unit {
	return _unit_Frequency
}

func NewLength(val float64) Length {
	return _Length(val)
}

func (q _Length) Value() float64 {
	return float64(q)
}

func (q _Length) Unit() units.Unit {
	return _unit_Length
}

func (q _Length) MultiplyLength(val Length) Area {
	return NewArea(q.Value() * val.Value())
}

func (q _Length) MultiplyScalar(val Scalar) Length {
	return NewLength(q.Value() * val.Value())
}

func NewScalar(val float64) Scalar {
	return _Scalar(val)
}

func (q _Scalar) Value() float64 {
	return float64(q)
}

func (q _Scalar) Unit() units.Unit {
	return _unit_Scalar
}

func (q _Scalar) DivideTime(val Time) Frequency {
	return NewFrequency(q.Value() / val.Value())
}

func NewSpeed(val float64) Speed {
	return _Speed(val)
}

func (q _Speed) Value() float64 {
	return float64(q)
}

func (q _Speed) Unit() units.Unit {
	return _unit_Speed
}

func (q _Speed) MultiplyDuration(val Duration) Distance {
	return NewDistance(q.Value() * val.Value())
}

func NewTime(val float64) Time {
	return _Time(val)
}

func (q _Time) Value() float64 {
	return float64(q)
}

func (q _Time) Unit() units.Unit {
	return _unit_Time
}

func (q _Time) MultiplyScalar(val Scalar) Time {
	return NewTime(q.Value() * val.Value())
}
