package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const testCfg = `package example;
base (
	m = "m";
	s = "s";
)

derived (
	mps = m / s;
)

quantity (
  Scalar();
  Length(m);
  Area(m * m);
  Time(s);
  Speed(mps);
	Frequency(s^-1);
)

operation (
  Area = Length * Length;
  Length = Length * Scalar;
  Time = Time * Scalar;
  Speed = Length / Time;
  Length = Speed * Time;
  Frequency = Scalar / Time;
)`

func TestParser(t *testing.T) {
	f, err := parse(testCfg)
	require.NoError(t, err)
	require.NotNil(t, f)

	s, err := analyse(f)
	require.NoError(t, err)

	require.Len(t, s.Quantities["Scalar"].Operations, 1)
	require.Len(t, s.Quantities["Length"].Operations, 3)
	require.Len(t, s.Quantities["Time"].Operations, 1)
	require.Len(t, s.Quantities["Speed"].Operations, 1)
	require.Len(t, s.Quantities["Length"].Operations, 3)
	require.Len(t, s.Quantities["Frequency"].Operations, 0)
}
