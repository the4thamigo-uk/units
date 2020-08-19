package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const testCfg = `package example;
unit (
	m;
	s;
	h = s * 3600;
	km = 1000 * m;
)

quantity (
  Scalar() float64;
  Length(m) float64;
  Area(m * m) float64;
  Time(s) float64;
	Frequency(s^-1) float64;
	Distance(km) float64;
	Duration(h) float64;
  Speed(km / h) float64;
)

operation (
  Area = Length * Length;
  Length = Length * Scalar;
  Time = Time * Scalar;
  Speed = Distance / Duration;
  Distance = Speed * Duration;
  Frequency = Scalar / Time;
)`

func TestParser(t *testing.T) {
	f, err := parse(testCfg)
	require.NoError(t, err)
	require.NotNil(t, f)

	s, err := analyse(f)
	require.NoError(t, err)

	require.Len(t, s.Quantities["Scalar"].Operations, 1)
	require.Len(t, s.Quantities["Length"].Operations, 2)
	require.Len(t, s.Quantities["Area"].Operations, 0)
	require.Len(t, s.Quantities["Time"].Operations, 1)
	require.Len(t, s.Quantities["Frequency"].Operations, 0)
	require.Len(t, s.Quantities["Distance"].Operations, 1)
	require.Len(t, s.Quantities["Duration"].Operations, 0)
	require.Len(t, s.Quantities["Speed"].Operations, 1)
}