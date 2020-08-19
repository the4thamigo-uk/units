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
