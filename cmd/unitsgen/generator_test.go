package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerator(t *testing.T) {

	ast, err := parse(testCfg)
	require.NoError(t, err)

	s, err := buildModel(ast)
	require.NoError(t, err)

	_, err = generate(s)
	require.NoError(t, err)
}
