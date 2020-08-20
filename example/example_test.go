package example

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLength_NewLength(t *testing.T) {
	q := NewLength(1)
	require.Equal(t, float64(1), q.ValueOrDefault(2))
}

func TestLength_NewLengthPtr(t *testing.T) {
	require.Equal(t, float64(1), NewLengthPtr(1).ValueOrDefault(2))
}

func TestLength_NewFromPtr(t *testing.T) {
	var n *float64
	f := float64(1)
	require.Nil(t, NewLengthFromPtr(n))
	require.Equal(t, float64(1), NewLengthFromPtr(&f).ValueOrDefault(2))
}

func TestLength_Value(t *testing.T) {
	var n *Length
	require.Nil(t, n.Value())
	require.Nil(t, n.Value())
	require.Equal(t, float64(1), n.ValueOrDefault(1))

	q := NewLengthPtr(1)
	require.NotNil(t, q.Value())
	require.Equal(t, float64(1), *q.Value())
	require.Equal(t, float64(1), q.ValueOrDefault(0))
}

func TestLength_IsZero(t *testing.T) {
	var n *Length
	q1 := NewLengthPtr(0)
	q2 := NewLengthPtr(1)
	require.False(t, n.IsZero())
	require.True(t, q1.IsZero())
	require.False(t, q2.IsZero())
}

func TestLength_Eq(t *testing.T) {
	var n *Length
	q := NewLengthPtr(1)
	q2 := NewLengthPtr(2)
	require.False(t, q.Eq(n))
	require.False(t, n.Eq(q))
	require.True(t, q.Eq(q))
	require.False(t, q.Eq(q2))
}

func TestLength_Gt(t *testing.T) {
	var n *Length
	q := NewLengthPtr(1)
	q2 := NewLengthPtr(2)
	require.False(t, q.Gt(n))
	require.False(t, n.Gt(q))
	require.False(t, q.Gt(q))
	require.False(t, q.Gt(q2))
	require.True(t, q2.Gt(q))
}

func TestLength_Lt(t *testing.T) {
	var n *Length
	q := NewLengthPtr(1)
	q2 := NewLengthPtr(2)
	require.False(t, q.Lt(n))
	require.False(t, n.Lt(q))
	require.False(t, q.Lt(q))
	require.True(t, q.Lt(q2))
	require.False(t, q2.Lt(q))
}

func TestLength_GtEq(t *testing.T) {
	var n *Length
	q := NewLengthPtr(1)
	q2 := NewLengthPtr(2)
	require.False(t, q.GtEq(n))
	require.False(t, n.GtEq(q))
	require.True(t, q.GtEq(q))
	require.False(t, q.GtEq(q2))
	require.True(t, q2.GtEq(q))
}

func TestLength_LtEq(t *testing.T) {
	var n *Length
	q := NewLengthPtr(1)
	q2 := NewLengthPtr(2)
	require.False(t, q.LtEq(n))
	require.False(t, n.LtEq(q))
	require.True(t, q.LtEq(q))
	require.True(t, q.LtEq(q2))
	require.False(t, q2.LtEq(q))
}

func TestLength_Between(t *testing.T) {
	var n *Length
	l := NewLengthPtr(1)
	m := NewLengthPtr(2)
	u := NewLengthPtr(3)
	require.False(t, m.Between(n, u))
	require.False(t, m.Between(l, n))
	require.False(t, n.Between(l, u))
	require.False(t, n.Between(n, n))

	require.True(t, m.Between(l, u))
	require.True(t, l.Between(l, u))
	require.True(t, u.Between(l, u))
}

func TestLength_Inside(t *testing.T) {
	var n *Length
	l := NewLengthPtr(1)
	m := NewLengthPtr(2)
	u := NewLengthPtr(3)
	require.False(t, m.Inside(n, u))
	require.False(t, m.Inside(l, n))
	require.False(t, n.Inside(l, u))
	require.False(t, n.Inside(n, n))

	require.True(t, m.Inside(l, u))
	require.False(t, l.Inside(l, u))
	require.False(t, u.Inside(l, u))
}

func TestLength_Negate(t *testing.T) {
	var n *Length
	q := NewLengthPtr(1)
	nq := NewLengthPtr(-1)
	require.Nil(t, n.Negate())
	require.Equal(t, nq, q.Negate())
}

func TestLength_Abs(t *testing.T) {
	var n *Length
	q := NewLengthPtr(1)
	nq := NewLengthPtr(-1)
	require.Nil(t, n.Negate())
	require.Equal(t, q, q.Abs())
	require.Equal(t, q, nq.Negate())
}

func TestLength_Min(t *testing.T) {
	var n *Length
	l := NewLengthPtr(1)
	u := NewLengthPtr(2)
	require.Nil(t, n.Min(n))
	require.Nil(t, n.Min(l))
	require.Nil(t, n.Min(u))
	require.Nil(t, l.Min(n))
	require.Nil(t, u.Min(n))
	require.Equal(t, l, l.Min(u))
	require.Nil(t, u.Min(n))
	require.Equal(t, l, u.Min(l))
}

func TestLength_Max(t *testing.T) {
	var n *Length
	l := NewLengthPtr(1)
	u := NewLengthPtr(2)
	require.Nil(t, n.Max(n))
	require.Nil(t, n.Max(l))
	require.Nil(t, n.Max(u))
	require.Nil(t, l.Max(n))
	require.Nil(t, u.Max(n))
	require.Equal(t, u, l.Max(u))
	require.Nil(t, u.Max(n))
	require.Equal(t, u, u.Max(l))
}

func TestLength_AddLength(t *testing.T) {
	var n *Length
	q1 := NewLengthPtr(1)
	q2 := NewLengthPtr(2)
	z := NewLengthPtr(0)
	r := NewLengthPtr(3)
	require.Nil(t, n.AddLength(n))
	require.Nil(t, n.AddLength(q1))
	require.Nil(t, q1.AddLength(n))
	require.Equal(t, q1, q1.AddLength(z))
	require.Equal(t, r, q1.AddLength(q2))
}

func TestLength_SubtractLength(t *testing.T) {
	var n *Length
	q1 := NewLengthPtr(1)
	q2 := NewLengthPtr(2)
	z := NewLengthPtr(0)
	r := NewLengthPtr(-1)
	require.Nil(t, n.SubtractLength(n))
	require.Nil(t, n.SubtractLength(q1))
	require.Nil(t, q1.SubtractLength(n))
	require.Equal(t, q1, q1.SubtractLength(z))
	require.Equal(t, r, q1.SubtractLength(q2))
}

func TestLength_MultiplyLength(t *testing.T) {
	var n *Length
	q1 := NewLengthPtr(2)
	q2 := NewLengthPtr(4)
	z := NewLengthPtr(0)
	a := NewAreaPtr(8)
	za := NewAreaPtr(0)
	require.Nil(t, n.MultiplyLength(n))
	require.Nil(t, n.MultiplyLength(q1))
	require.Nil(t, q1.MultiplyLength(n))
	require.Equal(t, za, q1.MultiplyLength(z))
	require.Equal(t, a, q1.MultiplyLength(q2))
}

func TestLength_DivideLength(t *testing.T) {
	var n *Length
	q1 := NewLengthPtr(4)
	q2 := NewLengthPtr(2)
	z := NewLengthPtr(0)
	s := NewScalarPtr(2)
	require.Nil(t, n.DivideLength(n))
	require.Nil(t, n.DivideLength(q1))
	require.Nil(t, q1.DivideLength(n))
	require.Nil(t, q1.DivideLength(z))
	require.Equal(t, s, q1.DivideLength(q2))
}

func TestLength_DivideTime(t *testing.T) {
	d := NewDistance(100)
	tm := NewDuration(50)
	s := d.DivideDuration(&tm)
	t.Logf("%v (%v = %v)", s.Value(), s.Unit(), s.BaseUnit())
	require.Equal(t, float64(2), s.ValueOrDefault(0))

	s_kn, err := s.Convert(kn)
	require.NoError(t, err)
	t.Logf("%v (%s=%v)", s_kn, "kn", kn)

	s_mps, err := s.Convert(mps)
	require.NoError(t, err)
	t.Logf("%v (%v)", s_mps, mps)
}
