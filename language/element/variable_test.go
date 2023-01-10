package element

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVariableGetMarshal(t *testing.T) {
	t.Parallel()
	element := &VariableGet{
		Name: "example",
	}

	actual, err := MarshalXml(element)
	require.NoError(t, err)
	expected := `<variable-get name="example"></variable-get>`
	require.Equal(t, expected, string(actual))

	decoded, err := UnmarshalXmlBytes(actual)
	require.NoError(t, err)
	encoded, err := MarshalXml(decoded)
	require.NoError(t, err)
	require.Equal(t, expected, string(encoded))
}
