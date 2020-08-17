package units

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUnit_ScalarInvert(t *testing.T) {
	s := Scalar()
	require.Equal(t, Scalar(), s.Invert())
}

func TestUnit_ScalarMultiply(t *testing.T) {
	s := Scalar().Multiply(Scalar())
	require.Equal(t, Scalar(), s.Invert())
}

func TestUnit_ScalarDivide(t *testing.T) {
	s := Scalar().Divide(Scalar())
	require.Equal(t, Scalar(), s.Invert())
}

func TestUnit_Invert(t *testing.T) {
	u := NewUnit("u", 1)
	require.NotEqual(t, Scalar(), u)
	require.NotEqual(t, u, u.Invert())
	require.Equal(t, u, u.Invert().Invert())
}

func TestUnit_MultiplyByInverse(t *testing.T) {
	u := NewUnit("u", 1)
	require.NotEqual(t, u, u.Multiply(u.Invert()))
}

func TestUnit_MultiplyByScalarGivesSelf(t *testing.T) {
	u := NewUnit("u", 1)
	require.Equal(t, u, u.Multiply(Scalar()))
}

func TestUnit_ScalarDivideIsInvert(t *testing.T) {
	u := NewUnit("u", 1)
	require.Equal(t, Scalar().Divide(u), u.Invert())
}

func TestUnit_DivideBySelfGivesScalar(t *testing.T) {
	u := NewUnit("u", 1)
	require.Equal(t, Scalar(), u.Divide(u))
}

func TestUnit_DivideByScalarGivesSelf(t *testing.T) {
	u := NewUnit("u", 1)
	require.Equal(t, u, u.Divide(Scalar()))
}

func TestUnit_MultiplyByDemoninatorGivesNumerator(t *testing.T) {
	u1 := NewUnit("u1", 1)
	u2 := NewUnit("u2", 1)
	u := u1.Divide(u2)
	require.Equal(t, u1, u.Multiply(u2))
}

func TestUnit_MultiplyByInverseNumeratorGivesInverseDemoninator(t *testing.T) {
	u1 := NewUnit("u1", 1)
	u2 := NewUnit("u2", 1)
	u := u1.Divide(u2)
	require.Equal(t, u2.Invert(), u.Multiply(u1.Invert()))
}

func TestUnit_MultiplyCommutativity(t *testing.T) {
	u1 := NewUnit("u1", 1)
	u2 := NewUnit("u2", 1)
	u3 := u1.Multiply(u2)
	u4 := u2.Multiply(u1)
	require.Equal(t, u3, u4)
}

func TestUnit_MultiplySquaring(t *testing.T) {
	u := NewUnit("u", 1)
	u2 := NewUnit("u", 2)
	require.NotEqual(t, u, u2)
	require.Equal(t, u, u2.Divide(u))
}

func TestUnit_DivideSquaring(t *testing.T) {
	u := NewUnit("u", 1)
	u1 := NewUnit("u", -1)
	u2 := u1.Divide(u)
	require.NotEqual(t, u1, u2)
	require.Equal(t, u1, u2.Multiply(u))
}

func TestUnit_MultipliedUnitsCancel(t *testing.T) {
	u1 := NewUnit("u1", 1)
	u2 := NewUnit("u2", 1)
	u := u1.Multiply(u2)
	require.NotEqual(t, u, u1)
	require.NotEqual(t, u, u2)
	require.Equal(t, u1, u.Divide(u2))
	require.Equal(t, u2, u.Divide(u1))
	require.Equal(t, Scalar(), u.Divide(u1).Divide(u2))
}

func TestUnit_Unmarshal(t *testing.T) {
	u := Scalar()
	err := u.UnmarshalText([]byte("u"))
	require.NoError(t, err)
	require.Equal(t, NewUnit("u", 1), u)
}

func TestUnit_Substitute(t *testing.T) {
	u1 := NewUnit("u1", 1)
	u2 := NewUnit("u2", 1)

	u3 := u1.Subs(map[string]Unit{"u1": u2})
	require.Equal(t, NewUnit("u2", 1), u3)
}

func TestUnit_SubstituteSquared(t *testing.T) {
	u1 := NewUnit("u1", 2)
	u2 := NewUnit("u2", 1)

	u := u1.Subs(map[string]Unit{"u1": u2})
	require.Equal(t, NewUnit("u2", 2), u)
}

func TestUnit_SubstituteCompoundSquared(t *testing.T) {
	u1 := NewUnit("u1", 2)
	u2 := NewUnit("u1", 1)

	u := u1.Subs(map[string]Unit{"u2": u2})
	require.Equal(t, NewUnit("u1", 2), u)
}

func TestUnit_SubstituteRecursive(t *testing.T) {
	u1 := NewUnit("u1", 1).Multiply(NewUnit("u2", -2)).Multiply(NewUnit("u3", 3))

	u := u1.Subs(
		map[string]Unit{
			"u2": NewUnit("u1", 1),
			"u3": NewUnit("u2", 1),
		})
	require.Equal(t, NewUnit("u1", 2), u)
}

func TestUnit_SubstituteComplex(t *testing.T) {
	u1 := NewUnit("u1", 1).Multiply(NewUnit("u2", -2)).Multiply(NewUnit("u3", 3))

	u := u1.Subs(
		map[string]Unit{
			"u1": NewUnit("u2", 3),
			"u3": NewUnit("u1", -3),
		})
	require.Equal(t, NewUnit("u2", -26), u)
}
