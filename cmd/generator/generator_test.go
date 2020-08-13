package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerator(t *testing.T) {

	f, err := parse(testCfg)
	require.NoError(t, err)

	out, err := generate(f)
	require.NoError(t, err)
	t.Log(f.Units)
	t.Log(out)
}
