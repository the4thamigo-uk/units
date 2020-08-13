package example

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Speed(t *testing.T) {
	l := NewLength(100)
	tm := NewTime(50)
	s := l.DivideTime(tm)
	t.Logf("%v (%v)", s.Value(), s.Unit())
	require.Equal(t, float64(2), s.Value())

}
