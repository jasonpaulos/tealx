package element

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoopMarshal(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		element  *Loop
		expected string
	}{
		{
			name: "condition only",
			element: &Loop{
				Start:     nil,
				Condition: &Int{Value: 1},
				Step:      nil,
				Body:      Container{Children: []Element{&Int{Value: 2}}},
			},
			expected: `<loop><condition><int value="1"></int></condition><body><int value="2"></int></body></loop>`,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual, err := MarshalXml(testCase.element)
			require.NoError(t, err)
			require.Equal(t, testCase.expected, string(actual))

			decoded, err := UnmarshalXmlBytes(actual)
			require.NoError(t, err)
			encoded, err := MarshalXml(decoded)
			require.NoError(t, err)
			require.Equal(t, testCase.expected, string(encoded))
		})
	}
}
