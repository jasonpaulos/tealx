package element

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProgramMarshal(t *testing.T) {
	t.Parallel()
	element := &Program{
		Version: 8,
		Subroutines: []*Subroutine{
			{
				Name: "test_subroutine",
				Arguments: []SubroutineArgumentInfo{
					{Name: "a", Type: VariableTypeUint64},
					{Name: "b", Type: VariableTypeBytes},
				},
				Return: &SubroutineReturnInfo{Type: VariableTypeUint64},
				Body: Container{
					Children: []Element{&Int{Value: 1}},
				},
			},
		},
		Main: Container{
			Children: []Element{
				&Bytes{Value: []byte("testing 1 2 3 4 5")},
				&ProgramReturn{Value: &Int{Value: 1234}},
			},
		},
	}

	actual, err := MarshalXml(element)
	require.NoError(t, err)
	expected := `<program version="8"><main><bytes value="74657374696e6720312032203320342035" format="hex"></bytes><program-return><int value="1234"></int></program-return></main><subroutine name="test_subroutine"><body><int value="1"></int></body><returns type="uint64"></returns><argument name="a" type="uint64"></argument><argument name="b" type="bytes"></argument></subroutine></program>`
	require.Equal(t, expected, string(actual))

	decoded, err := UnmarshalXmlBytes(actual)
	require.NoError(t, err)
	encoded, err := MarshalXml(decoded)
	require.NoError(t, err)
	require.Equal(t, expected, string(encoded))
}
