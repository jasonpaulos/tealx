package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompileExample(t *testing.T) {
	inFile, err := os.OpenFile("../../examples/example.xml", os.O_RDONLY, 0666)
	require.NoError(t, err)
	outFile, err := os.OpenFile("../../examples/example.teal", os.O_RDONLY, 0666)
	require.NoError(t, err)

	expected, err := io.ReadAll(outFile)
	require.NoError(t, err)

	var actual strings.Builder
	err = compileProgram(inFile, &actual)
	require.NoError(t, err)

	require.Equal(t, string(expected), actual.String())
}
