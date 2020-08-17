package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const testCfg = `
package example;
unit (
  scalarUnit = "";
  timeUnit = "s";
  lengthUnit = "m";
  speedUnit = "m * s^-1";
  frequencyUnit = "s^-1";
  areaUnit = "m^2";
)

quantity (
  Scalar(scalarUnit);
  Length(lengthUnit);
  Time(timeUnit);
  Speed(speedUnit);
  Frequency(frequencyUnit);
  Area(areaUnit);
)

operation (
  Area = Length * Length;
  Length = Length * Scalar;
  Time = Time * Scalar;
  Speed = Length / Time;
  Length = Speed * Time;
  Frequency = Scalar / Time;
)
`

func TestParser(t *testing.T) {
	f, err := parse(testCfg)
	require.NoError(t, err)
	require.NotNil(t, f)
	t.Log(f.Quantities)
	require.Len(t, f.Quantities[0].OperationDefinitions, 1)
	require.Len(t, f.Quantities[1].OperationDefinitions, 3)
	require.Len(t, f.Quantities[2].OperationDefinitions, 1)
	require.Len(t, f.Quantities[3].OperationDefinitions, 1)
	require.Len(t, f.Quantities[4].OperationDefinitions, 0)
	require.Len(t, f.Quantities[5].OperationDefinitions, 0)
}
