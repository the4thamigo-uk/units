package units

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConverter_Incompatible(t *testing.T) {
	m1 := NewUnit("m", 1)
	m2 := NewUnit("m", 2)
	_, err := NewConverter(m1, m2)
	require.Error(t, err)
}

func TestConverter_Convert(t *testing.T) {
	m := NewUnit("m", 1)
	km := m.Multiply(Scalar(1000))
	c, err := NewConverter(km, m)
	require.NoError(t, err)
	require.Equal(t, float64(1000), c.Convert(1))
}

func TestConverter_ConvertPtr(t *testing.T) {
	m := NewUnit("m", 1)
	km := m.Multiply(Scalar(1000))
	c, err := NewConverter(km, m)
	require.NoError(t, err)
	val := float64(1)
	require.Equal(t, float64(1000), *c.ConvertPtr(&val))
}

func TestConverter_ConvertPtr_Nil(t *testing.T) {
	u1 := NewUnit("m", 1)
	u2 := NewUnit("m", 1)
	c, err := NewConverter(u1, u2)
	require.NoError(t, err)
	require.Nil(t, c.ConvertPtr(nil))
}
