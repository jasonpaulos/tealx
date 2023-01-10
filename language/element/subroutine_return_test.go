package element

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubroutineReturnMarshal(t *testing.T) {
	t.Parallel()
	element := &SubroutineReturn{
		Value: &Int{Value: 7},
	}

	actual, err := MarshalXml(element)
	require.NoError(t, err)
	expected := `<subroutine-return><int value="7"></int></subroutine-return>`
	require.Equal(t, expected, string(actual))

	decoded, err := UnmarshalXmlBytes(actual)
	require.NoError(t, err)
	encoded, err := MarshalXml(decoded)
	require.NoError(t, err)
	require.Equal(t, expected, string(encoded))
}
