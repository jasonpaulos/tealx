package element

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLogMarshal(t *testing.T) {
	t.Parallel()
	element := &Log{
		Value: &Int{Value: 1},
	}

	actual, err := MarshalXml(element)
	require.NoError(t, err)
	expected := `<log><int value="1"></int></log>`
	require.Equal(t, expected, string(actual))

	decoded, err := UnmarshalXmlBytes(actual)
	require.NoError(t, err)
	encoded, err := MarshalXml(decoded)
	require.NoError(t, err)
	require.Equal(t, expected, string(encoded))
}
