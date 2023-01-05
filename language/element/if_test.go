package element

import (
	"testing"

	"github.com/jasonpaulos/tealx/language"
	"github.com/stretchr/testify/require"
)

func TestIfMarshal(t *testing.T) {
	t.Parallel()
	element := &If{
		Condition: &Int{Value: 1},
		Then:      Container{Children: []Element{&Int{Value: 2}}},
		Else:      Container{Children: []Element{&Int{Value: 3}, &Int{Value: 4}}},
	}

	actual, err := MarshalXml(element)
	require.NoError(t, err)
	expected := `<if><condition><int value="1"></int></condition><then><int value="2"></int></then><else><int value="3"></int><int value="4"></int></else></if>`
	require.Equal(t, expected, string(actual))

	decoded, err := UnmarshalXmlBytes(actual)
	require.NoError(t, err)
	encoded, err := MarshalXml(decoded)
	require.NoError(t, err)
	require.Equal(t, expected, string(encoded))
}

func TestIfCodegen(t *testing.T) {
	t.Parallel()
	element := &If{
		Condition: &Int{Value: 1},
		Then:      Container{Children: []Element{&Int{Value: 2}}},
		Else:      Container{Children: []Element{&Int{Value: 3}, &Int{Value: 4}}},
	}

	graph := element.Codegen()

	// this is just a temporary test to manually inspect output

	blocks := graph.Sort()
	ops := language.Flatten(blocks)
	compiled := language.Serialize(ops)

	t.Log(compiled)
}
