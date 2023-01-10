package element

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinaryMarshal(t *testing.T) {
	t.Parallel()

	for binaryType := BinaryType(0); binaryType < binaryTypeFinal; binaryType++ {
		name := binaryType.XmlName().Local
		binaryType := binaryType
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			element := &Binary{
				Type:  binaryType,
				Left:  &Int{Value: 1},
				Right: &Int{Value: 2},
			}

			actual, err := MarshalXml(element)
			require.NoError(t, err)
			expected := fmt.Sprintf(`<%s><int value="1"></int><int value="2"></int></%s>`, name, name)
			require.Equal(t, expected, string(actual))

			decoded, err := UnmarshalXmlBytes(actual)
			require.NoError(t, err)
			encoded, err := MarshalXml(decoded)
			require.NoError(t, err)
			require.Equal(t, expected, string(encoded))
		})
	}
}
