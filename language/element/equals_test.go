package element

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEqualsMarshal(t *testing.T) {
	t.Parallel()
	element := &Equals{
		Left:  &Int{Value: 1},
		Right: &Int{Value: 2},
	}

	actual, err := MarshalXml(element)
	require.NoError(t, err)
	expected := `<equals><int value="1"></int><int value="2"></int></equals>`
	require.Equal(t, expected, string(actual))

	decoded, err := UnmarshalXmlBytes(actual)
	require.NoError(t, err)
	encoded, err := MarshalXml(decoded)
	require.NoError(t, err)
	require.Equal(t, expected, string(encoded))
}
