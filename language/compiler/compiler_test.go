package compiler

import (
	"strings"
	"testing"

	. "github.com/jasonpaulos/tealx/language/element"
	"github.com/stretchr/testify/require"
)

func TestProgramCompile(t *testing.T) {
	t.Parallel()
	element := &Program{
		Version: 8,
		Subroutines: []*Subroutine{
			{
				Name: "test_subroutine",
				Arguments: []SubroutineArgumentInfo{
					{Name: "a", Type: "uint64"},
					{Name: "b", Type: "bytes"},
				},
				Return: &SubroutineReturnInfo{Type: "uint64"},
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

	var builder strings.Builder
	err := Compile(*element, &builder)
	require.NoError(t, err)

	t.Log(builder.String())
}

func TestIfCodegen(t *testing.T) {
	t.Parallel()
	element := &If{
		Condition: &Int{Value: 1},
		Then:      Container{Children: []Element{&Int{Value: 2}}},
		Else:      Container{Children: []Element{&Int{Value: 3}, &Int{Value: 4}}},
	}

	graph := element.Codegen(CodegenContext{})

	// this is just a temporary test to manually inspect output

	blocks := graph.Sort()
	ops := flatten(blocks, "")
	compiled := serialize(ops)

	t.Log(compiled)
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

	graph := element.Codegen(CodegenContext{})

	// this is just a temporary test to manually inspect output

	blocks := graph.Sort()
	ops := flatten(blocks, "")
	compiled := serialize(ops)

	t.Log(compiled)
}
