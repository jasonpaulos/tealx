package element

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestXmlMarshalRoundtrip(t *testing.T) {
	t.Parallel()
	element := &Program{
		Version: 8,
		Container: Container{
			Children: []Element{
				&Bytes{Value: []byte("testing 1 2 3 4 5")},
				&Return{Value: &Int{Value: 1234}},
			},
		},
	}

	encoded, err := MarshalXml(element)
	require.NoError(t, err)

	decoded, err := UnmarshalXmlBytes(encoded)
	require.NoError(t, err)

	roundTripEncoded, err := MarshalXml(decoded)
	require.NoError(t, err)

	require.Equal(t, encoded, roundTripEncoded)

	t.Log(string(encoded))
}
