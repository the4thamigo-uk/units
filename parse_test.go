package units

import (
	"github.com/openlyinc/pointy"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParse_Empty(t *testing.T) {
	u, err := Parse("")
	require.NoError(t, err)
	require.Equal(t, Scalar(), u)
}

func TestParse_Unit(t *testing.T) {
	u, err := Parse("u")
	require.NoError(t, err)
	require.Equal(t, NewUnit("u", 1), u)
}

func TestParse_UnitsMultiplied(t *testing.T) {
	u, err := Parse("u1 * u2")
	require.NoError(t, err)
	require.Equal(t, NewUnit("u1", 1).Multiply(NewUnit("u2", 1)), u)
}

func TestParse_UnitsExponentsAdded(t *testing.T) {
	u, err := Parse("u * u^1 * u^2 * u^3")
	require.NoError(t, err)
	require.Equal(t, NewUnit("u", 7), u)
}

func TestParse_UnitsExponentsSubtracted(t *testing.T) {
	u, err := Parse("u * u^1 * u^2 * u^-3")
	require.NoError(t, err)
	require.Equal(t, NewUnit("u", 1), u)
}

func TestParse_CompoundUnit(t *testing.T) {
	u, err := Parse("u1^1 * u2^-1 * u1^-2 * u2^2")
	require.NoError(t, err)
	require.Equal(t, NewUnit("u1", -1).Multiply(NewUnit("u2", 1)), u)
}

func TestParse_CompoundUnitCancelsToScalar(t *testing.T) {
	u, err := Parse("u1^1 * u2^-1 * u1^-1 * u2^1")
	require.NoError(t, err)
	require.Equal(t, Scalar(), u)
}

func TestParse_EvalUnitExpression_BlankIsScalar(t *testing.T) {
	ue := Expression{}
	u, err := ue.Unit()
	require.NoError(t, err)
	require.Equal(t, Scalar(), u)
}

func TestParse_EvalUnitExpression_SingleTerm(t *testing.T) {
	ue := Expression{
		Lhs: Term{Name: "u", Exponent: pointy.Int(2)},
	}
	u, err := ue.Unit()
	require.NoError(t, err)
	require.Equal(t, NewUnit("u", 2), u)
}

func TestParse_EvalUnitExpression_TwoTerms(t *testing.T) {
	ue := Expression{
		Lhs:      Term{Name: "u1", Exponent: pointy.Int(1)},
		Operator: pointy.String("/"),
		Rhs: &Expression{
			Lhs: Term{Name: "u2", Exponent: pointy.Int(2)},
		},
	}
	u, err := ue.Unit()
	require.NoError(t, err)
	require.Equal(t, NewUnit("u1", 1).Divide(NewUnit("u2", 2)), u)
}

func TestParse_EvalUnitExpression_ThreeTerms(t *testing.T) {
	ue := Expression{
		Lhs:      Term{Name: "u1", Exponent: pointy.Int(3)},
		Operator: pointy.String("/"),
		Rhs: &Expression{
			Lhs:      Term{Name: "u2", Exponent: pointy.Int(2)},
			Operator: pointy.String("/"),
			Rhs: &Expression{
				Lhs: Term{Name: "u3", Exponent: pointy.Int(1)},
			},
		},
	}
	u, err := ue.Unit()
	require.NoError(t, err)
	require.Equal(t, NewUnit("u1", 3).Divide(NewUnit("u2", 2)).Divide(NewUnit("u3", 1)), u)
}
