package example

import (
	"fmt"
	"github.com/the4thamigo-uk/units"
	"math"
)

// ensure we use math package to avoid unused imports error for some code generations
var _ = math.Abs(0)

// units
var (
	BF   = units.MustParse(units.Parse("BF"))
	H    = units.MustParse(units.Parse("3.60000000000000000e+03*S"))
	KM   = units.MustParse(units.Parse("1.00000000000000000e+03*M"))
	KMpH = units.MustParse(units.Parse("2.77777777777777790e-01*M*S^-1"))
	KN   = units.MustParse(units.Parse("5.14444444444444482e-01*M*S^-1"))
	M    = units.MustParse(units.Parse("M"))
	Min  = units.MustParse(units.Parse("6.00000000000000000e+01*S"))
	MpS  = units.MustParse(units.Parse("M*S^-1"))
	NM   = units.MustParse(units.Parse("1.85200000000000000e+03*M"))
	S    = units.MustParse(units.Parse("S"))
)

// quantities
type (
	Area      float64
	Beaufort  int
	Distance  float64
	Duration  float64
	Frequency float64
	Length    float64
	RPM       float64
	Scalar    float64
	Speed     float64
	Time      float64
)

// quantity units
var (
	_unit_Area           = units.MustParse(units.Parse("M^2"))
	_base_unit_Area      = units.MustParse(units.Parse("M^2"))
	_unit_Beaufort       = units.MustParse(units.Parse("BF"))
	_base_unit_Beaufort  = units.MustParse(units.Parse("BF"))
	_unit_Distance       = units.MustParse(units.Parse("KM"))
	_base_unit_Distance  = units.MustParse(units.Parse("1.00000000000000000e+03*M"))
	_unit_Duration       = units.MustParse(units.Parse("H"))
	_base_unit_Duration  = units.MustParse(units.Parse("3.60000000000000000e+03*S"))
	_unit_Frequency      = units.MustParse(units.Parse("S^-1"))
	_base_unit_Frequency = units.MustParse(units.Parse("S^-1"))
	_unit_Length         = units.MustParse(units.Parse("M"))
	_base_unit_Length    = units.MustParse(units.Parse("M"))
	_unit_RPM            = units.MustParse(units.Parse("Min^-1"))
	_base_unit_RPM       = units.MustParse(units.Parse("1.66666666666666664e-02*S^-1"))
	_unit_Scalar         = units.MustParse(units.Parse(""))
	_base_unit_Scalar    = units.MustParse(units.Parse(""))
	_unit_Speed          = units.MustParse(units.Parse("H^-1*KM"))
	_base_unit_Speed     = units.MustParse(units.Parse("2.77777777777777790e-01*M*S^-1"))
	_unit_Time           = units.MustParse(units.Parse("S"))
	_base_unit_Time      = units.MustParse(units.Parse("S"))
)

// quantity zero values
var (
	_zero_value_Area      = NewArea(0)
	_zero_value_Beaufort  = NewBeaufort(0)
	_zero_value_Distance  = NewDistance(0)
	_zero_value_Duration  = NewDuration(0)
	_zero_value_Frequency = NewFrequency(0)
	_zero_value_Length    = NewLength(0)
	_zero_value_RPM       = NewRPM(0)
	_zero_value_Scalar    = NewScalar(0)
	_zero_value_Speed     = NewSpeed(0)
	_zero_value_Time      = NewTime(0)
)

// quantity slices
type (
	AreaSlice        []*Area
	AreaMapper       func(q *Area) *Area
	AreaFilter       func(q *Area) bool
	AreaReducer      func(acc *Area, q *Area) *Area
	BeaufortSlice    []*Beaufort
	BeaufortMapper   func(q *Beaufort) *Beaufort
	BeaufortFilter   func(q *Beaufort) bool
	BeaufortReducer  func(acc *Beaufort, q *Beaufort) *Beaufort
	DistanceSlice    []*Distance
	DistanceMapper   func(q *Distance) *Distance
	DistanceFilter   func(q *Distance) bool
	DistanceReducer  func(acc *Distance, q *Distance) *Distance
	DurationSlice    []*Duration
	DurationMapper   func(q *Duration) *Duration
	DurationFilter   func(q *Duration) bool
	DurationReducer  func(acc *Duration, q *Duration) *Duration
	FrequencySlice   []*Frequency
	FrequencyMapper  func(q *Frequency) *Frequency
	FrequencyFilter  func(q *Frequency) bool
	FrequencyReducer func(acc *Frequency, q *Frequency) *Frequency
	LengthSlice      []*Length
	LengthMapper     func(q *Length) *Length
	LengthFilter     func(q *Length) bool
	LengthReducer    func(acc *Length, q *Length) *Length
	RPMSlice         []*RPM
	RPMMapper        func(q *RPM) *RPM
	RPMFilter        func(q *RPM) bool
	RPMReducer       func(acc *RPM, q *RPM) *RPM
	ScalarSlice      []*Scalar
	ScalarMapper     func(q *Scalar) *Scalar
	ScalarFilter     func(q *Scalar) bool
	ScalarReducer    func(acc *Scalar, q *Scalar) *Scalar
	SpeedSlice       []*Speed
	SpeedMapper      func(q *Speed) *Speed
	SpeedFilter      func(q *Speed) bool
	SpeedReducer     func(acc *Speed, q *Speed) *Speed
	TimeSlice        []*Time
	TimeMapper       func(q *Time) *Time
	TimeFilter       func(q *Time) bool
	TimeReducer      func(acc *Time, q *Time) *Time
)

// declared conversions between quantities
var (
	_convertFrequencyToRPM = units.MustConvert(units.NewConverter(_base_unit_Frequency, _base_unit_RPM))
	_convertRPMToFrequency = units.MustConvert(units.NewConverter(_base_unit_RPM, _base_unit_Frequency))
)

func NewArea(val float64) Area {
	return Area(val)
}

func NewAreaPtr(val float64) *Area {
	q := NewArea(val)
	return &q
}

func NewAreaFromPtr(val *float64) *Area {
	if val == nil {
		return nil
	}
	return NewAreaPtr(*val)
}

func (q *Area) Value() *float64 {
	if q == nil {
		return nil
	}
	v := float64(*q)
	return &v
}

func (q *Area) ValueOrDefault(dft float64) float64 {
	if q == nil {
		return dft
	}
	return float64(*q)
}

func (q *Area) Convert(u units.Unit) (*float64, error) {
	if q == nil {
		return nil, nil
	}
	u2 := q.BaseUnit().Divide(u)
	if !u2.IsScalar() {
		return nil, fmt.Errorf("cannot convert '%s' to given units '%s'", q.Unit(), u)
	}
	v := float64(u2.Scale() * float64(*q.Value()))
	return &v, nil
}

func (q *Area) ConvertOrDefault(u units.Unit, dft float64) (*float64, error) {
	if q == nil {
		return &dft, nil
	}
	return q.Convert(u)
}

func (q *Area) Unit() units.Unit {
	return _unit_Area
}

func (q *Area) BaseUnit() units.Unit {
	return _base_unit_Area
}

func (q *Area) IsZero() bool {
	return q.Eq(&_zero_value_Area)
}

func (q *Area) Eq(q2 *Area) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() == *q2.Value()
}

func (q *Area) Gt(q2 *Area) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() > *q2.Value()
}

func (q *Area) GtEq(q2 *Area) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() >= *q2.Value()
}

func (q *Area) Lt(q2 *Area) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() < *q2.Value()
}

func (q *Area) LtEq(q2 *Area) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() <= *q2.Value()
}

func (q *Area) Between(q1, q2 *Area) bool {
	if q == nil || q1 == nil || q2 == nil {
		return false
	}
	return q.GtEq(q1) && q.LtEq(q2)
}

func (q *Area) Inside(q1, q2 *Area) bool {
	if q == nil || q1 == nil || q2 == nil {
		return false
	}
	return q.Gt(q1) && q.Lt(q2)
}

func (q *Area) Negate() *Area {
	if q == nil {
		return nil
	}
	return NewAreaPtr(-*q.Value())
}

func (q *Area) Abs() *Area {
	if q == nil {
		return nil
	}
	if q.GtEq(&_zero_value_Area) {
		return q
	}
	return q.Negate()
}

func (q *Area) Min(q2 *Area) *Area {
	if q == nil || q2 == nil {
		return nil
	}
	if q.Lt(q2) {
		return q
	}
	return q2
}

func (q *Area) Max(q2 *Area) *Area {
	if q == nil || q2 == nil {
		return nil
	}
	if q.Gt(q2) {
		return q
	}
	return q2
}

func (qq AreaSlice) Map(f AreaMapper) AreaSlice {
	out := make(AreaSlice, len(qq))
	for i, q := range qq {
		out[i] = f(q)
	}
	return out
}

func (qq AreaSlice) Filter(f AreaFilter) AreaSlice {
	out := make(AreaSlice, 0, len(qq))
	for _, q := range qq {
		if f(q) {
			out = append(out, q)
		}
	}
	return out
}

func (qq AreaSlice) Reduce(acc *Area, f AreaReducer) *Area {
	for _, q := range qq {
		acc = f(acc, q)
	}
	return acc
}

func NewBeaufort(val int) Beaufort {
	return Beaufort(val)
}

func NewBeaufortPtr(val int) *Beaufort {
	q := NewBeaufort(val)
	return &q
}

func NewBeaufortFromPtr(val *int) *Beaufort {
	if val == nil {
		return nil
	}
	return NewBeaufortPtr(*val)
}

func (q *Beaufort) Value() *int {
	if q == nil {
		return nil
	}
	v := int(*q)
	return &v
}

func (q *Beaufort) ValueOrDefault(dft int) int {
	if q == nil {
		return dft
	}
	return int(*q)
}

func (q *Beaufort) Convert(u units.Unit) (*int, error) {
	if q == nil {
		return nil, nil
	}
	u2 := q.BaseUnit().Divide(u)
	if !u2.IsScalar() {
		return nil, fmt.Errorf("cannot convert '%s' to given units '%s'", q.Unit(), u)
	}
	v := int(u2.Scale() * float64(*q.Value()))
	return &v, nil
}

func (q *Beaufort) ConvertOrDefault(u units.Unit, dft int) (*int, error) {
	if q == nil {
		return &dft, nil
	}
	return q.Convert(u)
}

func (q *Beaufort) Unit() units.Unit {
	return _unit_Beaufort
}

func (q *Beaufort) BaseUnit() units.Unit {
	return _base_unit_Beaufort
}

func (q *Beaufort) IsZero() bool {
	return q.Eq(&_zero_value_Beaufort)
}

func (q *Beaufort) Eq(q2 *Beaufort) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() == *q2.Value()
}

func (q *Beaufort) Gt(q2 *Beaufort) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() > *q2.Value()
}

func (q *Beaufort) GtEq(q2 *Beaufort) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() >= *q2.Value()
}

func (q *Beaufort) Lt(q2 *Beaufort) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() < *q2.Value()
}

func (q *Beaufort) LtEq(q2 *Beaufort) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() <= *q2.Value()
}

func (q *Beaufort) Between(q1, q2 *Beaufort) bool {
	if q == nil || q1 == nil || q2 == nil {
		return false
	}
	return q.GtEq(q1) && q.LtEq(q2)
}

func (q *Beaufort) Inside(q1, q2 *Beaufort) bool {
	if q == nil || q1 == nil || q2 == nil {
		return false
	}
	return q.Gt(q1) && q.Lt(q2)
}

func (q *Beaufort) Negate() *Beaufort {
	if q == nil {
		return nil
	}
	return NewBeaufortPtr(-*q.Value())
}

func (q *Beaufort) Abs() *Beaufort {
	if q == nil {
		return nil
	}
	if q.GtEq(&_zero_value_Beaufort) {
		return q
	}
	return q.Negate()
}

func (q *Beaufort) Min(q2 *Beaufort) *Beaufort {
	if q == nil || q2 == nil {
		return nil
	}
	if q.Lt(q2) {
		return q
	}
	return q2
}

func (q *Beaufort) Max(q2 *Beaufort) *Beaufort {
	if q == nil || q2 == nil {
		return nil
	}
	if q.Gt(q2) {
		return q
	}
	return q2
}

func (qq BeaufortSlice) Map(f BeaufortMapper) BeaufortSlice {
	out := make(BeaufortSlice, len(qq))
	for i, q := range qq {
		out[i] = f(q)
	}
	return out
}

func (qq BeaufortSlice) Filter(f BeaufortFilter) BeaufortSlice {
	out := make(BeaufortSlice, 0, len(qq))
	for _, q := range qq {
		if f(q) {
			out = append(out, q)
		}
	}
	return out
}

func (qq BeaufortSlice) Reduce(acc *Beaufort, f BeaufortReducer) *Beaufort {
	for _, q := range qq {
		acc = f(acc, q)
	}
	return acc
}

func (q *Beaufort) ModBeaufort(q2 *Beaufort) *Beaufort {
	if q == nil || q2 == nil {
		return nil
	}

	if q2.IsZero() {
		return nil
	}

	return NewBeaufortPtr(*q.Value() % *q2.Value())
}

func NewDistance(val float64) Distance {
	return Distance(val)
}

func NewDistancePtr(val float64) *Distance {
	q := NewDistance(val)
	return &q
}

func NewDistanceFromPtr(val *float64) *Distance {
	if val == nil {
		return nil
	}
	return NewDistancePtr(*val)
}

func (q *Distance) Value() *float64 {
	if q == nil {
		return nil
	}
	v := float64(*q)
	return &v
}

func (q *Distance) ValueOrDefault(dft float64) float64 {
	if q == nil {
		return dft
	}
	return float64(*q)
}

func (q *Distance) Convert(u units.Unit) (*float64, error) {
	if q == nil {
		return nil, nil
	}
	u2 := q.BaseUnit().Divide(u)
	if !u2.IsScalar() {
		return nil, fmt.Errorf("cannot convert '%s' to given units '%s'", q.Unit(), u)
	}
	v := float64(u2.Scale() * float64(*q.Value()))
	return &v, nil
}

func (q *Distance) ConvertOrDefault(u units.Unit, dft float64) (*float64, error) {
	if q == nil {
		return &dft, nil
	}
	return q.Convert(u)
}

func (q *Distance) Unit() units.Unit {
	return _unit_Distance
}

func (q *Distance) BaseUnit() units.Unit {
	return _base_unit_Distance
}

func (q *Distance) IsZero() bool {
	return q.Eq(&_zero_value_Distance)
}

func (q *Distance) Eq(q2 *Distance) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() == *q2.Value()
}

func (q *Distance) Gt(q2 *Distance) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() > *q2.Value()
}

func (q *Distance) GtEq(q2 *Distance) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() >= *q2.Value()
}

func (q *Distance) Lt(q2 *Distance) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() < *q2.Value()
}

func (q *Distance) LtEq(q2 *Distance) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() <= *q2.Value()
}

func (q *Distance) Between(q1, q2 *Distance) bool {
	if q == nil || q1 == nil || q2 == nil {
		return false
	}
	return q.GtEq(q1) && q.LtEq(q2)
}

func (q *Distance) Inside(q1, q2 *Distance) bool {
	if q == nil || q1 == nil || q2 == nil {
		return false
	}
	return q.Gt(q1) && q.Lt(q2)
}

func (q *Distance) Negate() *Distance {
	if q == nil {
		return nil
	}
	return NewDistancePtr(-*q.Value())
}

func (q *Distance) Abs() *Distance {
	if q == nil {
		return nil
	}
	if q.GtEq(&_zero_value_Distance) {
		return q
	}
	return q.Negate()
}

func (q *Distance) Min(q2 *Distance) *Distance {
	if q == nil || q2 == nil {
		return nil
	}
	if q.Lt(q2) {
		return q
	}
	return q2
}

func (q *Distance) Max(q2 *Distance) *Distance {
	if q == nil || q2 == nil {
		return nil
	}
	if q.Gt(q2) {
		return q
	}
	return q2
}

func (qq DistanceSlice) Map(f DistanceMapper) DistanceSlice {
	out := make(DistanceSlice, len(qq))
	for i, q := range qq {
		out[i] = f(q)
	}
	return out
}

func (qq DistanceSlice) Filter(f DistanceFilter) DistanceSlice {
	out := make(DistanceSlice, 0, len(qq))
	for _, q := range qq {
		if f(q) {
			out = append(out, q)
		}
	}
	return out
}

func (qq DistanceSlice) Reduce(acc *Distance, f DistanceReducer) *Distance {
	for _, q := range qq {
		acc = f(acc, q)
	}
	return acc
}

func (q *Distance) DivideDuration(q2 *Duration) *Speed {
	if q == nil || q2 == nil {
		return nil
	}

	if q2.IsZero() {
		return nil
	}

	return NewSpeedPtr(*q.Value() / *q2.Value())
}

func NewDuration(val float64) Duration {
	return Duration(val)
}

func NewDurationPtr(val float64) *Duration {
	q := NewDuration(val)
	return &q
}

func NewDurationFromPtr(val *float64) *Duration {
	if val == nil {
		return nil
	}
	return NewDurationPtr(*val)
}

func (q *Duration) Value() *float64 {
	if q == nil {
		return nil
	}
	v := float64(*q)
	return &v
}

func (q *Duration) ValueOrDefault(dft float64) float64 {
	if q == nil {
		return dft
	}
	return float64(*q)
}

func (q *Duration) Convert(u units.Unit) (*float64, error) {
	if q == nil {
		return nil, nil
	}
	u2 := q.BaseUnit().Divide(u)
	if !u2.IsScalar() {
		return nil, fmt.Errorf("cannot convert '%s' to given units '%s'", q.Unit(), u)
	}
	v := float64(u2.Scale() * float64(*q.Value()))
	return &v, nil
}

func (q *Duration) ConvertOrDefault(u units.Unit, dft float64) (*float64, error) {
	if q == nil {
		return &dft, nil
	}
	return q.Convert(u)
}

func (q *Duration) Unit() units.Unit {
	return _unit_Duration
}

func (q *Duration) BaseUnit() units.Unit {
	return _base_unit_Duration
}

func (q *Duration) IsZero() bool {
	return q.Eq(&_zero_value_Duration)
}

func (q *Duration) Eq(q2 *Duration) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() == *q2.Value()
}

func (q *Duration) Gt(q2 *Duration) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() > *q2.Value()
}

func (q *Duration) GtEq(q2 *Duration) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() >= *q2.Value()
}

func (q *Duration) Lt(q2 *Duration) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() < *q2.Value()
}

func (q *Duration) LtEq(q2 *Duration) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() <= *q2.Value()
}

func (q *Duration) Between(q1, q2 *Duration) bool {
	if q == nil || q1 == nil || q2 == nil {
		return false
	}
	return q.GtEq(q1) && q.LtEq(q2)
}

func (q *Duration) Inside(q1, q2 *Duration) bool {
	if q == nil || q1 == nil || q2 == nil {
		return false
	}
	return q.Gt(q1) && q.Lt(q2)
}

func (q *Duration) Negate() *Duration {
	if q == nil {
		return nil
	}
	return NewDurationPtr(-*q.Value())
}

func (q *Duration) Abs() *Duration {
	if q == nil {
		return nil
	}
	if q.GtEq(&_zero_value_Duration) {
		return q
	}
	return q.Negate()
}

func (q *Duration) Min(q2 *Duration) *Duration {
	if q == nil || q2 == nil {
		return nil
	}
	if q.Lt(q2) {
		return q
	}
	return q2
}

func (q *Duration) Max(q2 *Duration) *Duration {
	if q == nil || q2 == nil {
		return nil
	}
	if q.Gt(q2) {
		return q
	}
	return q2
}

func (qq DurationSlice) Map(f DurationMapper) DurationSlice {
	out := make(DurationSlice, len(qq))
	for i, q := range qq {
		out[i] = f(q)
	}
	return out
}

func (qq DurationSlice) Filter(f DurationFilter) DurationSlice {
	out := make(DurationSlice, 0, len(qq))
	for _, q := range qq {
		if f(q) {
			out = append(out, q)
		}
	}
	return out
}

func (qq DurationSlice) Reduce(acc *Duration, f DurationReducer) *Duration {
	for _, q := range qq {
		acc = f(acc, q)
	}
	return acc
}

func NewFrequency(val float64) Frequency {
	return Frequency(val)
}

func NewFrequencyPtr(val float64) *Frequency {
	q := NewFrequency(val)
	return &q
}

func NewFrequencyFromPtr(val *float64) *Frequency {
	if val == nil {
		return nil
	}
	return NewFrequencyPtr(*val)
}

func (q *Frequency) Value() *float64 {
	if q == nil {
		return nil
	}
	v := float64(*q)
	return &v
}

func (q *Frequency) ValueOrDefault(dft float64) float64 {
	if q == nil {
		return dft
	}
	return float64(*q)
}

func (q *Frequency) Convert(u units.Unit) (*float64, error) {
	if q == nil {
		return nil, nil
	}
	u2 := q.BaseUnit().Divide(u)
	if !u2.IsScalar() {
		return nil, fmt.Errorf("cannot convert '%s' to given units '%s'", q.Unit(), u)
	}
	v := float64(u2.Scale() * float64(*q.Value()))
	return &v, nil
}

func (q *Frequency) ConvertOrDefault(u units.Unit, dft float64) (*float64, error) {
	if q == nil {
		return &dft, nil
	}
	return q.Convert(u)
}

func (q *Frequency) Unit() units.Unit {
	return _unit_Frequency
}

func (q *Frequency) BaseUnit() units.Unit {
	return _base_unit_Frequency
}

func (q *Frequency) IsZero() bool {
	return q.Eq(&_zero_value_Frequency)
}

func (q *Frequency) Eq(q2 *Frequency) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() == *q2.Value()
}

func (q *Frequency) Gt(q2 *Frequency) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() > *q2.Value()
}

func (q *Frequency) GtEq(q2 *Frequency) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() >= *q2.Value()
}

func (q *Frequency) Lt(q2 *Frequency) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() < *q2.Value()
}

func (q *Frequency) LtEq(q2 *Frequency) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() <= *q2.Value()
}

func (q *Frequency) Between(q1, q2 *Frequency) bool {
	if q == nil || q1 == nil || q2 == nil {
		return false
	}
	return q.GtEq(q1) && q.LtEq(q2)
}

func (q *Frequency) Inside(q1, q2 *Frequency) bool {
	if q == nil || q1 == nil || q2 == nil {
		return false
	}
	return q.Gt(q1) && q.Lt(q2)
}

func (q *Frequency) Negate() *Frequency {
	if q == nil {
		return nil
	}
	return NewFrequencyPtr(-*q.Value())
}

func (q *Frequency) Abs() *Frequency {
	if q == nil {
		return nil
	}
	if q.GtEq(&_zero_value_Frequency) {
		return q
	}
	return q.Negate()
}

func (q *Frequency) Min(q2 *Frequency) *Frequency {
	if q == nil || q2 == nil {
		return nil
	}
	if q.Lt(q2) {
		return q
	}
	return q2
}

func (q *Frequency) Max(q2 *Frequency) *Frequency {
	if q == nil || q2 == nil {
		return nil
	}
	if q.Gt(q2) {
		return q
	}
	return q2
}

func (qq FrequencySlice) Map(f FrequencyMapper) FrequencySlice {
	out := make(FrequencySlice, len(qq))
	for i, q := range qq {
		out[i] = f(q)
	}
	return out
}

func (qq FrequencySlice) Filter(f FrequencyFilter) FrequencySlice {
	out := make(FrequencySlice, 0, len(qq))
	for _, q := range qq {
		if f(q) {
			out = append(out, q)
		}
	}
	return out
}

func (qq FrequencySlice) Reduce(acc *Frequency, f FrequencyReducer) *Frequency {
	for _, q := range qq {
		acc = f(acc, q)
	}
	return acc
}

func (q *Frequency) ToRPM() *RPM {
	if q == nil {
		return nil
	}
	return NewRPMFromPtr(_convertFrequencyToRPM.ConvertPtr(q.Value()))
}

func NewLength(val float64) Length {
	return Length(val)
}

func NewLengthPtr(val float64) *Length {
	q := NewLength(val)
	return &q
}

func NewLengthFromPtr(val *float64) *Length {
	if val == nil {
		return nil
	}
	return NewLengthPtr(*val)
}

func (q *Length) Value() *float64 {
	if q == nil {
		return nil
	}
	v := float64(*q)
	return &v
}

func (q *Length) ValueOrDefault(dft float64) float64 {
	if q == nil {
		return dft
	}
	return float64(*q)
}

func (q *Length) Convert(u units.Unit) (*float64, error) {
	if q == nil {
		return nil, nil
	}
	u2 := q.BaseUnit().Divide(u)
	if !u2.IsScalar() {
		return nil, fmt.Errorf("cannot convert '%s' to given units '%s'", q.Unit(), u)
	}
	v := float64(u2.Scale() * float64(*q.Value()))
	return &v, nil
}

func (q *Length) ConvertOrDefault(u units.Unit, dft float64) (*float64, error) {
	if q == nil {
		return &dft, nil
	}
	return q.Convert(u)
}

func (q *Length) Unit() units.Unit {
	return _unit_Length
}

func (q *Length) BaseUnit() units.Unit {
	return _base_unit_Length
}

func (q *Length) IsZero() bool {
	return q.Eq(&_zero_value_Length)
}

func (q *Length) Eq(q2 *Length) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() == *q2.Value()
}

func (q *Length) Gt(q2 *Length) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() > *q2.Value()
}

func (q *Length) GtEq(q2 *Length) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() >= *q2.Value()
}

func (q *Length) Lt(q2 *Length) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() < *q2.Value()
}

func (q *Length) LtEq(q2 *Length) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() <= *q2.Value()
}

func (q *Length) Between(q1, q2 *Length) bool {
	if q == nil || q1 == nil || q2 == nil {
		return false
	}
	return q.GtEq(q1) && q.LtEq(q2)
}

func (q *Length) Inside(q1, q2 *Length) bool {
	if q == nil || q1 == nil || q2 == nil {
		return false
	}
	return q.Gt(q1) && q.Lt(q2)
}

func (q *Length) Negate() *Length {
	if q == nil {
		return nil
	}
	return NewLengthPtr(-*q.Value())
}

func (q *Length) Abs() *Length {
	if q == nil {
		return nil
	}
	if q.GtEq(&_zero_value_Length) {
		return q
	}
	return q.Negate()
}

func (q *Length) Min(q2 *Length) *Length {
	if q == nil || q2 == nil {
		return nil
	}
	if q.Lt(q2) {
		return q
	}
	return q2
}

func (q *Length) Max(q2 *Length) *Length {
	if q == nil || q2 == nil {
		return nil
	}
	if q.Gt(q2) {
		return q
	}
	return q2
}

func (qq LengthSlice) Map(f LengthMapper) LengthSlice {
	out := make(LengthSlice, len(qq))
	for i, q := range qq {
		out[i] = f(q)
	}
	return out
}

func (qq LengthSlice) Filter(f LengthFilter) LengthSlice {
	out := make(LengthSlice, 0, len(qq))
	for _, q := range qq {
		if f(q) {
			out = append(out, q)
		}
	}
	return out
}

func (qq LengthSlice) Reduce(acc *Length, f LengthReducer) *Length {
	for _, q := range qq {
		acc = f(acc, q)
	}
	return acc
}

func (q *Length) MultiplyLength(q2 *Length) *Area {
	if q == nil || q2 == nil {
		return nil
	}

	return NewAreaPtr(*q.Value() * *q2.Value())
}

func (q *Length) ModLength(q2 *Length) *Length {
	if q == nil || q2 == nil {
		return nil
	}

	if q2.IsZero() {
		return nil
	}

	return NewLengthPtr(float64(math.Mod(float64(*q.Value()), float64(*q2.Value()))))
}

func (q *Length) DivideLength(q2 *Length) *Scalar {
	if q == nil || q2 == nil {
		return nil
	}

	if q2.IsZero() {
		return nil
	}

	return NewScalarPtr(*q.Value() / *q2.Value())
}

func (q *Length) AddLength(q2 *Length) *Length {
	if q == nil || q2 == nil {
		return nil
	}

	return NewLengthPtr(*q.Value() + *q2.Value())
}

func (q *Length) SubtractLength(q2 *Length) *Length {
	if q == nil || q2 == nil {
		return nil
	}

	return NewLengthPtr(*q.Value() - *q2.Value())
}

func (q *Length) MultiplyScalar(q2 *Scalar) *Length {
	if q == nil || q2 == nil {
		return nil
	}

	return NewLengthPtr(*q.Value() * *q2.Value())
}

func NewRPM(val float64) RPM {
	return RPM(val)
}

func NewRPMPtr(val float64) *RPM {
	q := NewRPM(val)
	return &q
}

func NewRPMFromPtr(val *float64) *RPM {
	if val == nil {
		return nil
	}
	return NewRPMPtr(*val)
}

func (q *RPM) Value() *float64 {
	if q == nil {
		return nil
	}
	v := float64(*q)
	return &v
}

func (q *RPM) ValueOrDefault(dft float64) float64 {
	if q == nil {
		return dft
	}
	return float64(*q)
}

func (q *RPM) Convert(u units.Unit) (*float64, error) {
	if q == nil {
		return nil, nil
	}
	u2 := q.BaseUnit().Divide(u)
	if !u2.IsScalar() {
		return nil, fmt.Errorf("cannot convert '%s' to given units '%s'", q.Unit(), u)
	}
	v := float64(u2.Scale() * float64(*q.Value()))
	return &v, nil
}

func (q *RPM) ConvertOrDefault(u units.Unit, dft float64) (*float64, error) {
	if q == nil {
		return &dft, nil
	}
	return q.Convert(u)
}

func (q *RPM) Unit() units.Unit {
	return _unit_RPM
}

func (q *RPM) BaseUnit() units.Unit {
	return _base_unit_RPM
}

func (q *RPM) IsZero() bool {
	return q.Eq(&_zero_value_RPM)
}

func (q *RPM) Eq(q2 *RPM) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() == *q2.Value()
}

func (q *RPM) Gt(q2 *RPM) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() > *q2.Value()
}

func (q *RPM) GtEq(q2 *RPM) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() >= *q2.Value()
}

func (q *RPM) Lt(q2 *RPM) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() < *q2.Value()
}

func (q *RPM) LtEq(q2 *RPM) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() <= *q2.Value()
}

func (q *RPM) Between(q1, q2 *RPM) bool {
	if q == nil || q1 == nil || q2 == nil {
		return false
	}
	return q.GtEq(q1) && q.LtEq(q2)
}

func (q *RPM) Inside(q1, q2 *RPM) bool {
	if q == nil || q1 == nil || q2 == nil {
		return false
	}
	return q.Gt(q1) && q.Lt(q2)
}

func (q *RPM) Negate() *RPM {
	if q == nil {
		return nil
	}
	return NewRPMPtr(-*q.Value())
}

func (q *RPM) Abs() *RPM {
	if q == nil {
		return nil
	}
	if q.GtEq(&_zero_value_RPM) {
		return q
	}
	return q.Negate()
}

func (q *RPM) Min(q2 *RPM) *RPM {
	if q == nil || q2 == nil {
		return nil
	}
	if q.Lt(q2) {
		return q
	}
	return q2
}

func (q *RPM) Max(q2 *RPM) *RPM {
	if q == nil || q2 == nil {
		return nil
	}
	if q.Gt(q2) {
		return q
	}
	return q2
}

func (qq RPMSlice) Map(f RPMMapper) RPMSlice {
	out := make(RPMSlice, len(qq))
	for i, q := range qq {
		out[i] = f(q)
	}
	return out
}

func (qq RPMSlice) Filter(f RPMFilter) RPMSlice {
	out := make(RPMSlice, 0, len(qq))
	for _, q := range qq {
		if f(q) {
			out = append(out, q)
		}
	}
	return out
}

func (qq RPMSlice) Reduce(acc *RPM, f RPMReducer) *RPM {
	for _, q := range qq {
		acc = f(acc, q)
	}
	return acc
}

func (q *RPM) ToFrequency() *Frequency {
	if q == nil {
		return nil
	}
	return NewFrequencyFromPtr(_convertRPMToFrequency.ConvertPtr(q.Value()))
}

func NewScalar(val float64) Scalar {
	return Scalar(val)
}

func NewScalarPtr(val float64) *Scalar {
	q := NewScalar(val)
	return &q
}

func NewScalarFromPtr(val *float64) *Scalar {
	if val == nil {
		return nil
	}
	return NewScalarPtr(*val)
}

func (q *Scalar) Value() *float64 {
	if q == nil {
		return nil
	}
	v := float64(*q)
	return &v
}

func (q *Scalar) ValueOrDefault(dft float64) float64 {
	if q == nil {
		return dft
	}
	return float64(*q)
}

func (q *Scalar) Convert(u units.Unit) (*float64, error) {
	if q == nil {
		return nil, nil
	}
	u2 := q.BaseUnit().Divide(u)
	if !u2.IsScalar() {
		return nil, fmt.Errorf("cannot convert '%s' to given units '%s'", q.Unit(), u)
	}
	v := float64(u2.Scale() * float64(*q.Value()))
	return &v, nil
}

func (q *Scalar) ConvertOrDefault(u units.Unit, dft float64) (*float64, error) {
	if q == nil {
		return &dft, nil
	}
	return q.Convert(u)
}

func (q *Scalar) Unit() units.Unit {
	return _unit_Scalar
}

func (q *Scalar) BaseUnit() units.Unit {
	return _base_unit_Scalar
}

func (q *Scalar) IsZero() bool {
	return q.Eq(&_zero_value_Scalar)
}

func (q *Scalar) Eq(q2 *Scalar) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() == *q2.Value()
}

func (q *Scalar) Gt(q2 *Scalar) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() > *q2.Value()
}

func (q *Scalar) GtEq(q2 *Scalar) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() >= *q2.Value()
}

func (q *Scalar) Lt(q2 *Scalar) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() < *q2.Value()
}

func (q *Scalar) LtEq(q2 *Scalar) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() <= *q2.Value()
}

func (q *Scalar) Between(q1, q2 *Scalar) bool {
	if q == nil || q1 == nil || q2 == nil {
		return false
	}
	return q.GtEq(q1) && q.LtEq(q2)
}

func (q *Scalar) Inside(q1, q2 *Scalar) bool {
	if q == nil || q1 == nil || q2 == nil {
		return false
	}
	return q.Gt(q1) && q.Lt(q2)
}

func (q *Scalar) Negate() *Scalar {
	if q == nil {
		return nil
	}
	return NewScalarPtr(-*q.Value())
}

func (q *Scalar) Abs() *Scalar {
	if q == nil {
		return nil
	}
	if q.GtEq(&_zero_value_Scalar) {
		return q
	}
	return q.Negate()
}

func (q *Scalar) Min(q2 *Scalar) *Scalar {
	if q == nil || q2 == nil {
		return nil
	}
	if q.Lt(q2) {
		return q
	}
	return q2
}

func (q *Scalar) Max(q2 *Scalar) *Scalar {
	if q == nil || q2 == nil {
		return nil
	}
	if q.Gt(q2) {
		return q
	}
	return q2
}

func (qq ScalarSlice) Map(f ScalarMapper) ScalarSlice {
	out := make(ScalarSlice, len(qq))
	for i, q := range qq {
		out[i] = f(q)
	}
	return out
}

func (qq ScalarSlice) Filter(f ScalarFilter) ScalarSlice {
	out := make(ScalarSlice, 0, len(qq))
	for _, q := range qq {
		if f(q) {
			out = append(out, q)
		}
	}
	return out
}

func (qq ScalarSlice) Reduce(acc *Scalar, f ScalarReducer) *Scalar {
	for _, q := range qq {
		acc = f(acc, q)
	}
	return acc
}

func (q *Scalar) DivideTime(q2 *Time) *Frequency {
	if q == nil || q2 == nil {
		return nil
	}

	if q2.IsZero() {
		return nil
	}

	return NewFrequencyPtr(*q.Value() / *q2.Value())
}

func NewSpeed(val float64) Speed {
	return Speed(val)
}

func NewSpeedPtr(val float64) *Speed {
	q := NewSpeed(val)
	return &q
}

func NewSpeedFromPtr(val *float64) *Speed {
	if val == nil {
		return nil
	}
	return NewSpeedPtr(*val)
}

func (q *Speed) Value() *float64 {
	if q == nil {
		return nil
	}
	v := float64(*q)
	return &v
}

func (q *Speed) ValueOrDefault(dft float64) float64 {
	if q == nil {
		return dft
	}
	return float64(*q)
}

func (q *Speed) Convert(u units.Unit) (*float64, error) {
	if q == nil {
		return nil, nil
	}
	u2 := q.BaseUnit().Divide(u)
	if !u2.IsScalar() {
		return nil, fmt.Errorf("cannot convert '%s' to given units '%s'", q.Unit(), u)
	}
	v := float64(u2.Scale() * float64(*q.Value()))
	return &v, nil
}

func (q *Speed) ConvertOrDefault(u units.Unit, dft float64) (*float64, error) {
	if q == nil {
		return &dft, nil
	}
	return q.Convert(u)
}

func (q *Speed) Unit() units.Unit {
	return _unit_Speed
}

func (q *Speed) BaseUnit() units.Unit {
	return _base_unit_Speed
}

func (q *Speed) IsZero() bool {
	return q.Eq(&_zero_value_Speed)
}

func (q *Speed) Eq(q2 *Speed) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() == *q2.Value()
}

func (q *Speed) Gt(q2 *Speed) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() > *q2.Value()
}

func (q *Speed) GtEq(q2 *Speed) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() >= *q2.Value()
}

func (q *Speed) Lt(q2 *Speed) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() < *q2.Value()
}

func (q *Speed) LtEq(q2 *Speed) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() <= *q2.Value()
}

func (q *Speed) Between(q1, q2 *Speed) bool {
	if q == nil || q1 == nil || q2 == nil {
		return false
	}
	return q.GtEq(q1) && q.LtEq(q2)
}

func (q *Speed) Inside(q1, q2 *Speed) bool {
	if q == nil || q1 == nil || q2 == nil {
		return false
	}
	return q.Gt(q1) && q.Lt(q2)
}

func (q *Speed) Negate() *Speed {
	if q == nil {
		return nil
	}
	return NewSpeedPtr(-*q.Value())
}

func (q *Speed) Abs() *Speed {
	if q == nil {
		return nil
	}
	if q.GtEq(&_zero_value_Speed) {
		return q
	}
	return q.Negate()
}

func (q *Speed) Min(q2 *Speed) *Speed {
	if q == nil || q2 == nil {
		return nil
	}
	if q.Lt(q2) {
		return q
	}
	return q2
}

func (q *Speed) Max(q2 *Speed) *Speed {
	if q == nil || q2 == nil {
		return nil
	}
	if q.Gt(q2) {
		return q
	}
	return q2
}

func (qq SpeedSlice) Map(f SpeedMapper) SpeedSlice {
	out := make(SpeedSlice, len(qq))
	for i, q := range qq {
		out[i] = f(q)
	}
	return out
}

func (qq SpeedSlice) Filter(f SpeedFilter) SpeedSlice {
	out := make(SpeedSlice, 0, len(qq))
	for _, q := range qq {
		if f(q) {
			out = append(out, q)
		}
	}
	return out
}

func (qq SpeedSlice) Reduce(acc *Speed, f SpeedReducer) *Speed {
	for _, q := range qq {
		acc = f(acc, q)
	}
	return acc
}

func (q *Speed) MultiplyDuration(q2 *Duration) *Distance {
	if q == nil || q2 == nil {
		return nil
	}

	return NewDistancePtr(*q.Value() * *q2.Value())
}

func NewTime(val float64) Time {
	return Time(val)
}

func NewTimePtr(val float64) *Time {
	q := NewTime(val)
	return &q
}

func NewTimeFromPtr(val *float64) *Time {
	if val == nil {
		return nil
	}
	return NewTimePtr(*val)
}

func (q *Time) Value() *float64 {
	if q == nil {
		return nil
	}
	v := float64(*q)
	return &v
}

func (q *Time) ValueOrDefault(dft float64) float64 {
	if q == nil {
		return dft
	}
	return float64(*q)
}

func (q *Time) Convert(u units.Unit) (*float64, error) {
	if q == nil {
		return nil, nil
	}
	u2 := q.BaseUnit().Divide(u)
	if !u2.IsScalar() {
		return nil, fmt.Errorf("cannot convert '%s' to given units '%s'", q.Unit(), u)
	}
	v := float64(u2.Scale() * float64(*q.Value()))
	return &v, nil
}

func (q *Time) ConvertOrDefault(u units.Unit, dft float64) (*float64, error) {
	if q == nil {
		return &dft, nil
	}
	return q.Convert(u)
}

func (q *Time) Unit() units.Unit {
	return _unit_Time
}

func (q *Time) BaseUnit() units.Unit {
	return _base_unit_Time
}

func (q *Time) IsZero() bool {
	return q.Eq(&_zero_value_Time)
}

func (q *Time) Eq(q2 *Time) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() == *q2.Value()
}

func (q *Time) Gt(q2 *Time) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() > *q2.Value()
}

func (q *Time) GtEq(q2 *Time) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() >= *q2.Value()
}

func (q *Time) Lt(q2 *Time) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() < *q2.Value()
}

func (q *Time) LtEq(q2 *Time) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() <= *q2.Value()
}

func (q *Time) Between(q1, q2 *Time) bool {
	if q == nil || q1 == nil || q2 == nil {
		return false
	}
	return q.GtEq(q1) && q.LtEq(q2)
}

func (q *Time) Inside(q1, q2 *Time) bool {
	if q == nil || q1 == nil || q2 == nil {
		return false
	}
	return q.Gt(q1) && q.Lt(q2)
}

func (q *Time) Negate() *Time {
	if q == nil {
		return nil
	}
	return NewTimePtr(-*q.Value())
}

func (q *Time) Abs() *Time {
	if q == nil {
		return nil
	}
	if q.GtEq(&_zero_value_Time) {
		return q
	}
	return q.Negate()
}

func (q *Time) Min(q2 *Time) *Time {
	if q == nil || q2 == nil {
		return nil
	}
	if q.Lt(q2) {
		return q
	}
	return q2
}

func (q *Time) Max(q2 *Time) *Time {
	if q == nil || q2 == nil {
		return nil
	}
	if q.Gt(q2) {
		return q
	}
	return q2
}

func (qq TimeSlice) Map(f TimeMapper) TimeSlice {
	out := make(TimeSlice, len(qq))
	for i, q := range qq {
		out[i] = f(q)
	}
	return out
}

func (qq TimeSlice) Filter(f TimeFilter) TimeSlice {
	out := make(TimeSlice, 0, len(qq))
	for _, q := range qq {
		if f(q) {
			out = append(out, q)
		}
	}
	return out
}

func (qq TimeSlice) Reduce(acc *Time, f TimeReducer) *Time {
	for _, q := range qq {
		acc = f(acc, q)
	}
	return acc
}

func (q *Time) MultiplyScalar(q2 *Scalar) *Time {
	if q == nil || q2 == nil {
		return nil
	}

	return NewTimePtr(*q.Value() * *q2.Value())
}
