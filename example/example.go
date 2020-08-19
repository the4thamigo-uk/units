package example

import (
	"fmt"
	"github.com/the4thamigo-uk/units"
)

// units
var (
	bf   = units.Must(units.Parse("bf"))
	h    = units.Must(units.Parse("3.60000000000000000e+03*s"))
	km   = units.Must(units.Parse("1.00000000000000000e+03*m"))
	kmph = units.Must(units.Parse("2.77777777777777790e-01*m*s^-1"))
	kn   = units.Must(units.Parse("5.14444444444444482e-01*m*s^-1"))
	m    = units.Must(units.Parse("m"))
	mps  = units.Must(units.Parse("m*s^-1"))
	nm   = units.Must(units.Parse("1.85200000000000000e+03*m"))
	s    = units.Must(units.Parse("s"))
)

// quantities
type (
	_Area float64
	Area  interface {
		Value() float64
		Convert(units.Unit) (float64, error)
		Unit() units.Unit
		BaseUnit() units.Unit
		Eq(q2 Area) bool
		Gt(q2 Area) bool
		GtEq(q2 Area) bool
		Lt(q2 Area) bool
		LtEq(q2 Area) bool
		Between(q1, q2 Area) bool
		Inside(q1, q2 Area) bool
		Abs() Area
		Min(q2 Area) Area
		Max(q2 Area) Area
	}

	_Beaufort int
	Beaufort  interface {
		Value() int
		Convert(units.Unit) (int, error)
		Unit() units.Unit
		BaseUnit() units.Unit
		Eq(q2 Beaufort) bool
		Gt(q2 Beaufort) bool
		GtEq(q2 Beaufort) bool
		Lt(q2 Beaufort) bool
		LtEq(q2 Beaufort) bool
		Between(q1, q2 Beaufort) bool
		Inside(q1, q2 Beaufort) bool
		Abs() Beaufort
		Min(q2 Beaufort) Beaufort
		Max(q2 Beaufort) Beaufort
	}

	_Distance float64
	Distance  interface {
		Value() float64
		Convert(units.Unit) (float64, error)
		Unit() units.Unit
		BaseUnit() units.Unit
		Eq(q2 Distance) bool
		Gt(q2 Distance) bool
		GtEq(q2 Distance) bool
		Lt(q2 Distance) bool
		LtEq(q2 Distance) bool
		Between(q1, q2 Distance) bool
		Inside(q1, q2 Distance) bool
		Abs() Distance
		Min(q2 Distance) Distance
		Max(q2 Distance) Distance
		DivideDuration(val Duration) Speed
	}

	_Duration float64
	Duration  interface {
		Value() float64
		Convert(units.Unit) (float64, error)
		Unit() units.Unit
		BaseUnit() units.Unit
		Eq(q2 Duration) bool
		Gt(q2 Duration) bool
		GtEq(q2 Duration) bool
		Lt(q2 Duration) bool
		LtEq(q2 Duration) bool
		Between(q1, q2 Duration) bool
		Inside(q1, q2 Duration) bool
		Abs() Duration
		Min(q2 Duration) Duration
		Max(q2 Duration) Duration
	}

	_Frequency float64
	Frequency  interface {
		Value() float64
		Convert(units.Unit) (float64, error)
		Unit() units.Unit
		BaseUnit() units.Unit
		Eq(q2 Frequency) bool
		Gt(q2 Frequency) bool
		GtEq(q2 Frequency) bool
		Lt(q2 Frequency) bool
		LtEq(q2 Frequency) bool
		Between(q1, q2 Frequency) bool
		Inside(q1, q2 Frequency) bool
		Abs() Frequency
		Min(q2 Frequency) Frequency
		Max(q2 Frequency) Frequency
	}

	_Length float64
	Length  interface {
		Value() float64
		Convert(units.Unit) (float64, error)
		Unit() units.Unit
		BaseUnit() units.Unit
		Eq(q2 Length) bool
		Gt(q2 Length) bool
		GtEq(q2 Length) bool
		Lt(q2 Length) bool
		LtEq(q2 Length) bool
		Between(q1, q2 Length) bool
		Inside(q1, q2 Length) bool
		Abs() Length
		Min(q2 Length) Length
		Max(q2 Length) Length
		MultiplyLength(val Length) Area
		MultiplyScalar(val Scalar) Length
	}

	_Scalar float64
	Scalar  interface {
		Value() float64
		Convert(units.Unit) (float64, error)
		Unit() units.Unit
		BaseUnit() units.Unit
		Eq(q2 Scalar) bool
		Gt(q2 Scalar) bool
		GtEq(q2 Scalar) bool
		Lt(q2 Scalar) bool
		LtEq(q2 Scalar) bool
		Between(q1, q2 Scalar) bool
		Inside(q1, q2 Scalar) bool
		Abs() Scalar
		Min(q2 Scalar) Scalar
		Max(q2 Scalar) Scalar
		DivideTime(val Time) Frequency
	}

	_Speed float64
	Speed  interface {
		Value() float64
		Convert(units.Unit) (float64, error)
		Unit() units.Unit
		BaseUnit() units.Unit
		Eq(q2 Speed) bool
		Gt(q2 Speed) bool
		GtEq(q2 Speed) bool
		Lt(q2 Speed) bool
		LtEq(q2 Speed) bool
		Between(q1, q2 Speed) bool
		Inside(q1, q2 Speed) bool
		Abs() Speed
		Min(q2 Speed) Speed
		Max(q2 Speed) Speed
		MultiplyDuration(val Duration) Distance
	}

	_Time float64
	Time  interface {
		Value() float64
		Convert(units.Unit) (float64, error)
		Unit() units.Unit
		BaseUnit() units.Unit
		Eq(q2 Time) bool
		Gt(q2 Time) bool
		GtEq(q2 Time) bool
		Lt(q2 Time) bool
		LtEq(q2 Time) bool
		Between(q1, q2 Time) bool
		Inside(q1, q2 Time) bool
		Abs() Time
		Min(q2 Time) Time
		Max(q2 Time) Time
		MultiplyScalar(val Scalar) Time
	}
)

// quantity units
var (
	_unit_Area           = units.Must(units.Parse("m^2"))
	_base_unit_Area      = units.Must(units.Parse("m^2"))
	_unit_Beaufort       = units.Must(units.Parse("bf"))
	_base_unit_Beaufort  = units.Must(units.Parse("bf"))
	_unit_Distance       = units.Must(units.Parse("km"))
	_base_unit_Distance  = units.Must(units.Parse("1.00000000000000000e+03*m"))
	_unit_Duration       = units.Must(units.Parse("h"))
	_base_unit_Duration  = units.Must(units.Parse("3.60000000000000000e+03*s"))
	_unit_Frequency      = units.Must(units.Parse("s^-1"))
	_base_unit_Frequency = units.Must(units.Parse("s^-1"))
	_unit_Length         = units.Must(units.Parse("m"))
	_base_unit_Length    = units.Must(units.Parse("m"))
	_unit_Scalar         = units.Must(units.Parse(""))
	_base_unit_Scalar    = units.Must(units.Parse(""))
	_unit_Speed          = units.Must(units.Parse("h^-1*km"))
	_base_unit_Speed     = units.Must(units.Parse("2.77777777777777790e-01*m*s^-1"))
	_unit_Time           = units.Must(units.Parse("s"))
	_base_unit_Time      = units.Must(units.Parse("s"))
)

// quantity zero values
var (
	_zero_value_Area      = NewArea(0)
	_zero_value_Beaufort  = NewBeaufort(0)
	_zero_value_Distance  = NewDistance(0)
	_zero_value_Duration  = NewDuration(0)
	_zero_value_Frequency = NewFrequency(0)
	_zero_value_Length    = NewLength(0)
	_zero_value_Scalar    = NewScalar(0)
	_zero_value_Speed     = NewSpeed(0)
	_zero_value_Time      = NewTime(0)
)

func NewArea(val float64) Area {
	return _Area(val)
}

func (q _Area) Value() float64 {
	return float64(q)
}

func (q _Area) Convert(u units.Unit) (float64, error) {
	u2 := q.BaseUnit().Divide(u)
	if !u2.IsScalar() {
		return 0, fmt.Errorf("cannot convert '%s' to given units '%s'", q.Unit(), u)
	}
	return float64(u2.Scale() * float64(q.Value())), nil
}

func (q _Area) Unit() units.Unit {
	return _unit_Area
}

func (q _Area) BaseUnit() units.Unit {
	return _base_unit_Area
}

func (q _Area) Eq(q2 Area) bool {
	return q.Value() == q2.Value()
}

func (q _Area) Gt(q2 Area) bool {
	return q.Value() > q2.Value()
}

func (q _Area) GtEq(q2 Area) bool {
	return q.Value() >= q2.Value()
}

func (q _Area) Lt(q2 Area) bool {
	return q.Value() < q2.Value()
}

func (q _Area) LtEq(q2 Area) bool {
	return q.Value() <= q2.Value()
}

func (q _Area) Between(q1, q2 Area) bool {
	return q.GtEq(q1) && q.LtEq(q2)
}

func (q _Area) Inside(q1, q2 Area) bool {
	return q.Gt(q1) && q.Lt(q2)
}

func (q _Area) Negate() Area {
	return NewArea(-q.Value())
}

func (q _Area) Abs() Area {
	if q.GtEq(_zero_value_Area) {
		return q
	}
	return q.Negate()
}

func (q _Area) Min(q2 Area) Area {
	if q.Lt(q2) {
		return q
	}
	return q2
}

func (q _Area) Max(q2 Area) Area {
	if q.Gt(q2) {
		return q
	}
	return q2
}

func NewBeaufort(val int) Beaufort {
	return _Beaufort(val)
}

func (q _Beaufort) Value() int {
	return int(q)
}

func (q _Beaufort) Convert(u units.Unit) (int, error) {
	u2 := q.BaseUnit().Divide(u)
	if !u2.IsScalar() {
		return 0, fmt.Errorf("cannot convert '%s' to given units '%s'", q.Unit(), u)
	}
	return int(u2.Scale() * float64(q.Value())), nil
}

func (q _Beaufort) Unit() units.Unit {
	return _unit_Beaufort
}

func (q _Beaufort) BaseUnit() units.Unit {
	return _base_unit_Beaufort
}

func (q _Beaufort) Eq(q2 Beaufort) bool {
	return q.Value() == q2.Value()
}

func (q _Beaufort) Gt(q2 Beaufort) bool {
	return q.Value() > q2.Value()
}

func (q _Beaufort) GtEq(q2 Beaufort) bool {
	return q.Value() >= q2.Value()
}

func (q _Beaufort) Lt(q2 Beaufort) bool {
	return q.Value() < q2.Value()
}

func (q _Beaufort) LtEq(q2 Beaufort) bool {
	return q.Value() <= q2.Value()
}

func (q _Beaufort) Between(q1, q2 Beaufort) bool {
	return q.GtEq(q1) && q.LtEq(q2)
}

func (q _Beaufort) Inside(q1, q2 Beaufort) bool {
	return q.Gt(q1) && q.Lt(q2)
}

func (q _Beaufort) Negate() Beaufort {
	return NewBeaufort(-q.Value())
}

func (q _Beaufort) Abs() Beaufort {
	if q.GtEq(_zero_value_Beaufort) {
		return q
	}
	return q.Negate()
}

func (q _Beaufort) Min(q2 Beaufort) Beaufort {
	if q.Lt(q2) {
		return q
	}
	return q2
}

func (q _Beaufort) Max(q2 Beaufort) Beaufort {
	if q.Gt(q2) {
		return q
	}
	return q2
}

func NewDistance(val float64) Distance {
	return _Distance(val)
}

func (q _Distance) Value() float64 {
	return float64(q)
}

func (q _Distance) Convert(u units.Unit) (float64, error) {
	u2 := q.BaseUnit().Divide(u)
	if !u2.IsScalar() {
		return 0, fmt.Errorf("cannot convert '%s' to given units '%s'", q.Unit(), u)
	}
	return float64(u2.Scale() * float64(q.Value())), nil
}

func (q _Distance) Unit() units.Unit {
	return _unit_Distance
}

func (q _Distance) BaseUnit() units.Unit {
	return _base_unit_Distance
}

func (q _Distance) Eq(q2 Distance) bool {
	return q.Value() == q2.Value()
}

func (q _Distance) Gt(q2 Distance) bool {
	return q.Value() > q2.Value()
}

func (q _Distance) GtEq(q2 Distance) bool {
	return q.Value() >= q2.Value()
}

func (q _Distance) Lt(q2 Distance) bool {
	return q.Value() < q2.Value()
}

func (q _Distance) LtEq(q2 Distance) bool {
	return q.Value() <= q2.Value()
}

func (q _Distance) Between(q1, q2 Distance) bool {
	return q.GtEq(q1) && q.LtEq(q2)
}

func (q _Distance) Inside(q1, q2 Distance) bool {
	return q.Gt(q1) && q.Lt(q2)
}

func (q _Distance) Negate() Distance {
	return NewDistance(-q.Value())
}

func (q _Distance) Abs() Distance {
	if q.GtEq(_zero_value_Distance) {
		return q
	}
	return q.Negate()
}

func (q _Distance) Min(q2 Distance) Distance {
	if q.Lt(q2) {
		return q
	}
	return q2
}

func (q _Distance) Max(q2 Distance) Distance {
	if q.Gt(q2) {
		return q
	}
	return q2
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

func (q _Duration) Convert(u units.Unit) (float64, error) {
	u2 := q.BaseUnit().Divide(u)
	if !u2.IsScalar() {
		return 0, fmt.Errorf("cannot convert '%s' to given units '%s'", q.Unit(), u)
	}
	return float64(u2.Scale() * float64(q.Value())), nil
}

func (q _Duration) Unit() units.Unit {
	return _unit_Duration
}

func (q _Duration) BaseUnit() units.Unit {
	return _base_unit_Duration
}

func (q _Duration) Eq(q2 Duration) bool {
	return q.Value() == q2.Value()
}

func (q _Duration) Gt(q2 Duration) bool {
	return q.Value() > q2.Value()
}

func (q _Duration) GtEq(q2 Duration) bool {
	return q.Value() >= q2.Value()
}

func (q _Duration) Lt(q2 Duration) bool {
	return q.Value() < q2.Value()
}

func (q _Duration) LtEq(q2 Duration) bool {
	return q.Value() <= q2.Value()
}

func (q _Duration) Between(q1, q2 Duration) bool {
	return q.GtEq(q1) && q.LtEq(q2)
}

func (q _Duration) Inside(q1, q2 Duration) bool {
	return q.Gt(q1) && q.Lt(q2)
}

func (q _Duration) Negate() Duration {
	return NewDuration(-q.Value())
}

func (q _Duration) Abs() Duration {
	if q.GtEq(_zero_value_Duration) {
		return q
	}
	return q.Negate()
}

func (q _Duration) Min(q2 Duration) Duration {
	if q.Lt(q2) {
		return q
	}
	return q2
}

func (q _Duration) Max(q2 Duration) Duration {
	if q.Gt(q2) {
		return q
	}
	return q2
}

func NewFrequency(val float64) Frequency {
	return _Frequency(val)
}

func (q _Frequency) Value() float64 {
	return float64(q)
}

func (q _Frequency) Convert(u units.Unit) (float64, error) {
	u2 := q.BaseUnit().Divide(u)
	if !u2.IsScalar() {
		return 0, fmt.Errorf("cannot convert '%s' to given units '%s'", q.Unit(), u)
	}
	return float64(u2.Scale() * float64(q.Value())), nil
}

func (q _Frequency) Unit() units.Unit {
	return _unit_Frequency
}

func (q _Frequency) BaseUnit() units.Unit {
	return _base_unit_Frequency
}

func (q _Frequency) Eq(q2 Frequency) bool {
	return q.Value() == q2.Value()
}

func (q _Frequency) Gt(q2 Frequency) bool {
	return q.Value() > q2.Value()
}

func (q _Frequency) GtEq(q2 Frequency) bool {
	return q.Value() >= q2.Value()
}

func (q _Frequency) Lt(q2 Frequency) bool {
	return q.Value() < q2.Value()
}

func (q _Frequency) LtEq(q2 Frequency) bool {
	return q.Value() <= q2.Value()
}

func (q _Frequency) Between(q1, q2 Frequency) bool {
	return q.GtEq(q1) && q.LtEq(q2)
}

func (q _Frequency) Inside(q1, q2 Frequency) bool {
	return q.Gt(q1) && q.Lt(q2)
}

func (q _Frequency) Negate() Frequency {
	return NewFrequency(-q.Value())
}

func (q _Frequency) Abs() Frequency {
	if q.GtEq(_zero_value_Frequency) {
		return q
	}
	return q.Negate()
}

func (q _Frequency) Min(q2 Frequency) Frequency {
	if q.Lt(q2) {
		return q
	}
	return q2
}

func (q _Frequency) Max(q2 Frequency) Frequency {
	if q.Gt(q2) {
		return q
	}
	return q2
}

func NewLength(val float64) Length {
	return _Length(val)
}

func (q _Length) Value() float64 {
	return float64(q)
}

func (q _Length) Convert(u units.Unit) (float64, error) {
	u2 := q.BaseUnit().Divide(u)
	if !u2.IsScalar() {
		return 0, fmt.Errorf("cannot convert '%s' to given units '%s'", q.Unit(), u)
	}
	return float64(u2.Scale() * float64(q.Value())), nil
}

func (q _Length) Unit() units.Unit {
	return _unit_Length
}

func (q _Length) BaseUnit() units.Unit {
	return _base_unit_Length
}

func (q _Length) Eq(q2 Length) bool {
	return q.Value() == q2.Value()
}

func (q _Length) Gt(q2 Length) bool {
	return q.Value() > q2.Value()
}

func (q _Length) GtEq(q2 Length) bool {
	return q.Value() >= q2.Value()
}

func (q _Length) Lt(q2 Length) bool {
	return q.Value() < q2.Value()
}

func (q _Length) LtEq(q2 Length) bool {
	return q.Value() <= q2.Value()
}

func (q _Length) Between(q1, q2 Length) bool {
	return q.GtEq(q1) && q.LtEq(q2)
}

func (q _Length) Inside(q1, q2 Length) bool {
	return q.Gt(q1) && q.Lt(q2)
}

func (q _Length) Negate() Length {
	return NewLength(-q.Value())
}

func (q _Length) Abs() Length {
	if q.GtEq(_zero_value_Length) {
		return q
	}
	return q.Negate()
}

func (q _Length) Min(q2 Length) Length {
	if q.Lt(q2) {
		return q
	}
	return q2
}

func (q _Length) Max(q2 Length) Length {
	if q.Gt(q2) {
		return q
	}
	return q2
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

func (q _Scalar) Convert(u units.Unit) (float64, error) {
	u2 := q.BaseUnit().Divide(u)
	if !u2.IsScalar() {
		return 0, fmt.Errorf("cannot convert '%s' to given units '%s'", q.Unit(), u)
	}
	return float64(u2.Scale() * float64(q.Value())), nil
}

func (q _Scalar) Unit() units.Unit {
	return _unit_Scalar
}

func (q _Scalar) BaseUnit() units.Unit {
	return _base_unit_Scalar
}

func (q _Scalar) Eq(q2 Scalar) bool {
	return q.Value() == q2.Value()
}

func (q _Scalar) Gt(q2 Scalar) bool {
	return q.Value() > q2.Value()
}

func (q _Scalar) GtEq(q2 Scalar) bool {
	return q.Value() >= q2.Value()
}

func (q _Scalar) Lt(q2 Scalar) bool {
	return q.Value() < q2.Value()
}

func (q _Scalar) LtEq(q2 Scalar) bool {
	return q.Value() <= q2.Value()
}

func (q _Scalar) Between(q1, q2 Scalar) bool {
	return q.GtEq(q1) && q.LtEq(q2)
}

func (q _Scalar) Inside(q1, q2 Scalar) bool {
	return q.Gt(q1) && q.Lt(q2)
}

func (q _Scalar) Negate() Scalar {
	return NewScalar(-q.Value())
}

func (q _Scalar) Abs() Scalar {
	if q.GtEq(_zero_value_Scalar) {
		return q
	}
	return q.Negate()
}

func (q _Scalar) Min(q2 Scalar) Scalar {
	if q.Lt(q2) {
		return q
	}
	return q2
}

func (q _Scalar) Max(q2 Scalar) Scalar {
	if q.Gt(q2) {
		return q
	}
	return q2
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

func (q _Speed) Convert(u units.Unit) (float64, error) {
	u2 := q.BaseUnit().Divide(u)
	if !u2.IsScalar() {
		return 0, fmt.Errorf("cannot convert '%s' to given units '%s'", q.Unit(), u)
	}
	return float64(u2.Scale() * float64(q.Value())), nil
}

func (q _Speed) Unit() units.Unit {
	return _unit_Speed
}

func (q _Speed) BaseUnit() units.Unit {
	return _base_unit_Speed
}

func (q _Speed) Eq(q2 Speed) bool {
	return q.Value() == q2.Value()
}

func (q _Speed) Gt(q2 Speed) bool {
	return q.Value() > q2.Value()
}

func (q _Speed) GtEq(q2 Speed) bool {
	return q.Value() >= q2.Value()
}

func (q _Speed) Lt(q2 Speed) bool {
	return q.Value() < q2.Value()
}

func (q _Speed) LtEq(q2 Speed) bool {
	return q.Value() <= q2.Value()
}

func (q _Speed) Between(q1, q2 Speed) bool {
	return q.GtEq(q1) && q.LtEq(q2)
}

func (q _Speed) Inside(q1, q2 Speed) bool {
	return q.Gt(q1) && q.Lt(q2)
}

func (q _Speed) Negate() Speed {
	return NewSpeed(-q.Value())
}

func (q _Speed) Abs() Speed {
	if q.GtEq(_zero_value_Speed) {
		return q
	}
	return q.Negate()
}

func (q _Speed) Min(q2 Speed) Speed {
	if q.Lt(q2) {
		return q
	}
	return q2
}

func (q _Speed) Max(q2 Speed) Speed {
	if q.Gt(q2) {
		return q
	}
	return q2
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

func (q _Time) Convert(u units.Unit) (float64, error) {
	u2 := q.BaseUnit().Divide(u)
	if !u2.IsScalar() {
		return 0, fmt.Errorf("cannot convert '%s' to given units '%s'", q.Unit(), u)
	}
	return float64(u2.Scale() * float64(q.Value())), nil
}

func (q _Time) Unit() units.Unit {
	return _unit_Time
}

func (q _Time) BaseUnit() units.Unit {
	return _base_unit_Time
}

func (q _Time) Eq(q2 Time) bool {
	return q.Value() == q2.Value()
}

func (q _Time) Gt(q2 Time) bool {
	return q.Value() > q2.Value()
}

func (q _Time) GtEq(q2 Time) bool {
	return q.Value() >= q2.Value()
}

func (q _Time) Lt(q2 Time) bool {
	return q.Value() < q2.Value()
}

func (q _Time) LtEq(q2 Time) bool {
	return q.Value() <= q2.Value()
}

func (q _Time) Between(q1, q2 Time) bool {
	return q.GtEq(q1) && q.LtEq(q2)
}

func (q _Time) Inside(q1, q2 Time) bool {
	return q.Gt(q1) && q.Lt(q2)
}

func (q _Time) Negate() Time {
	return NewTime(-q.Value())
}

func (q _Time) Abs() Time {
	if q.GtEq(_zero_value_Time) {
		return q
	}
	return q.Negate()
}

func (q _Time) Min(q2 Time) Time {
	if q.Lt(q2) {
		return q
	}
	return q2
}

func (q _Time) Max(q2 Time) Time {
	if q.Gt(q2) {
		return q
	}
	return q2
}

func (q _Time) MultiplyScalar(val Scalar) Time {
	return NewTime(q.Value() * val.Value())
}
