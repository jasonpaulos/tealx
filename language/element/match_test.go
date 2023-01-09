package element

import (
	"testing"

	"github.com/jasonpaulos/tealx/language"
	"github.com/stretchr/testify/require"
)

func TestMatchMarshal(t *testing.T) {
	t.Parallel()
	element := &Match{
		Value: &Int{Value: 1},
		DefaultCase: &Container{
			Children: []Element{&Bytes{Value: []byte("c")}},
		},
		Cases: []MatchCase{
			{
				Value: &Int{Value: 2},
				Body: Container{
					Children: []Element{&Bytes{Value: []byte("a")}},
				},
			},
			{
				Value: &Int{Value: 3},
				Body: Container{
					Children: []Element{&Bytes{Value: []byte("b")}},
				},
			},
		},
	}

	actual, err := MarshalXml(element)
	require.NoError(t, err)
	expected := `<match><value><int value="1"></int></value><default-case><bytes value="63" format="hex"></bytes></default-case><case><value><int value="2"></int></value><body><bytes value="61" format="hex"></bytes></body></case><case><value><int value="3"></int></value><body><bytes value="62" format="hex"></bytes></body></case></match>`
	require.Equal(t, expected, string(actual))

	decoded, err := UnmarshalXmlBytes(actual)
	require.NoError(t, err)
	encoded, err := MarshalXml(decoded)
	require.NoError(t, err)
	require.Equal(t, expected, string(encoded))
}

func TestMatchCodegen(t *testing.T) {
	t.Parallel()
	element := &Match{
		Value: &Int{Value: 1},
		DefaultCase: &Container{
			Children: []Element{&Bytes{Value: []byte("c")}},
		},
		Cases: []MatchCase{
			{
				Value: &Int{Value: 2},
				Body: Container{
					Children: []Element{&Bytes{Value: []byte("a")}},
				},
			},
			{
				Value: &Int{Value: 3},
				Body: Container{
					Children: []Element{&Bytes{Value: []byte("b")}},
				},
			},
		},
	}

	graph := element.Codegen()

	// this is just a temporary test to manually inspect output

	blocks := graph.Sort()
	ops := language.Flatten(blocks)
	compiled := language.Serialize(ops)

	t.Log(compiled)
}
