package units

import (
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
	require.Equal(t, NewUnit("u"), u)
}

func TestParse_UnitsMultiplied(t *testing.T) {
	u, err := Parse("u1 * u2")
	require.NoError(t, err)
	require.Equal(t, NewUnit("u1").Multiply(NewUnit("u2")), u)
}

func TestParse_UnitsExponentsAdded(t *testing.T) {
	u, err := Parse("u * u^1 * u^2 * u^3")
	require.NoError(t, err)
	require.Equal(t, newUnit("u", 7), u)
}

func TestParse_UnitsExponentsSubtracted(t *testing.T) {
	u, err := Parse("u * u^1 * u^2 * u^-3")
	require.NoError(t, err)
	require.Equal(t, newUnit("u", 1), u)
}

func TestParse_CompoundUnit(t *testing.T) {
	u, err := Parse("u1^1 * u2^-1 * u1^-2 * u2^2")
	require.NoError(t, err)
	require.Equal(t, newUnit("u1", -1).Multiply(newUnit("u2", 1)), u)
}

func TestParse_CompoundUnitCancelsToScalar(t *testing.T) {
	u, err := Parse("u1^1 * u2^-1 * u1^-1 * u2^1")
	require.NoError(t, err)
	require.Equal(t, Scalar(), u)
}
