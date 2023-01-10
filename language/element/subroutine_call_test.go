package element

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubroutineCallMarshal(t *testing.T) {
	t.Parallel()
	element := &SubroutineCall{
		Name: "test",
		Arguments: []Element{
			&Int{Value: 1},
			&Int{Value: 2},
		},
	}

	actual, err := MarshalXml(element)
	require.NoError(t, err)
	expected := `<subroutine-call name="test"><argument><int value="1"></int></argument><argument><int value="2"></int></argument></subroutine-call>`
	require.Equal(t, expected, string(actual))

	decoded, err := UnmarshalXmlBytes(actual)
	require.NoError(t, err)
	encoded, err := MarshalXml(decoded)
	require.NoError(t, err)
	require.Equal(t, expected, string(encoded))
}
