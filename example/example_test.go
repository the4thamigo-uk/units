package example

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Speed(t *testing.T) {
	d := NewDistance(100)
	tm := NewDuration(50)
	s := d.DivideDuration(tm)
	t.Logf("%v (%v = %v)", s.Value(), s.Unit(), s.BaseUnit())
	require.Equal(t, float64(2), s.Value())

	s_kn, err := s.Convert(kn)
	require.NoError(t, err)
	t.Logf("%v (%s=%v)", s_kn, "kn", kn)

	s_mps, err := s.Convert(mps)
	require.NoError(t, err)
	t.Logf("%v (%v)", s_mps, mps)
}

func Test_Abs(t *testing.T) {
	d := NewLength(-1)
	require.Equal(t, NewLength(1), d.Abs())
}

func Test_Comparison(t *testing.T) {
	small := NewLength(1)
	big := NewLength(2)
	medium := NewLength(1.5)
	require.True(t, big.Eq(big))
	require.True(t, big.Gt(small))
	require.True(t, big.GtEq(small))
	require.False(t, big.Lt(small))
	require.False(t, big.LtEq(small))
	require.False(t, big.Eq(small))
	require.True(t, medium.Between(small, big))
	require.False(t, medium.Between(big, small))
	require.True(t, medium.Inside(small, big))
	require.False(t, medium.Inside(big, small))
	require.True(t, small.Between(small, big))
	require.False(t, small.Between(big, small))
	require.False(t, small.Inside(small, big))
	require.False(t, small.Inside(big, small))
	require.True(t, big.Between(small, big))
	require.False(t, big.Between(big, small))
	require.False(t, big.Inside(small, big))
	require.False(t, big.Inside(big, small))
}
