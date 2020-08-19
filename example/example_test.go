package example

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Speed(t *testing.T) {
	d := NewDistance(100)
	tm := NewDuration(50)
	s := d.DivideDuration(tm)
	t.Logf("%v (%v)", s.Value(), s.Unit())
	require.Equal(t, float64(2), s.Value())

	val, err := s.Convert(kn)
	require.NoError(t, err)
	t.Logf("%v", val)
}
