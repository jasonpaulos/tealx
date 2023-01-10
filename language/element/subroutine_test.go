package element

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubroutineMarshal(t *testing.T) {
	t.Parallel()
	element := &Subroutine{
		Name: "test_subroutine",
		Arguments: []SubroutineArgumentInfo{
			{Name: "a", Type: "uint64"},
			{Name: "b", Type: "bytes"},
		},
		Return: &SubroutineReturnInfo{Type: "uint64"},
		Body: Container{
			Children: []Element{
				&SubroutineReturn{
					Value: &VariableGet{Name: "a"},
				},
			},
		},
	}

	actual, err := MarshalXml(element)
	require.NoError(t, err)
	expected := `<subroutine name="test_subroutine"><body><subroutine-return><variable-get name="a"></variable-get></subroutine-return></body><returns type="uint64"></returns><argument name="a" type="uint64"></argument><argument name="b" type="bytes"></argument></subroutine>`
	require.Equal(t, expected, string(actual))

	decoded, err := UnmarshalXmlBytes(actual)
	require.NoError(t, err)
	encoded, err := MarshalXml(decoded)
	require.NoError(t, err)
	require.Equal(t, expected, string(encoded))
}
