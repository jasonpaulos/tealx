package element

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/jasonpaulos/tealx/language"
)

type VariableGet struct {
	emptyElement

	Name string
}

func (v *VariableGet) Codegen(ctx CodegenContext) language.ControlFlowGraph {
	// search the context to figure out how to load the variable

	if ctx.CurrentSubroutine != nil {
		for i, arg := range ctx.CurrentSubroutine.Arguments {
			if arg.Name == v.Name {
				frameIndex := i - len(ctx.CurrentSubroutine.Arguments)
				return language.MakeControlFlowGraph([]language.Operation{
					{
						Opcode:    "frame_dig",
						Arguments: []string{strconv.Itoa(frameIndex)},
					},
				})
			}
		}
	}

	panic(fmt.Sprintf("variable not found: %s", v.Name))
}

func (v *VariableGet) xml() xmlElement {
	return &xmlVariableGet{Name: v.Name}
}

type xmlVariableGet struct {
	XMLName xml.Name `xml:"variable-get"`
	Name    string   `xml:"name,attr"`
}

func (x *xmlVariableGet) element() (Element, error) {
	return &VariableGet{
		Name: x.Name,
	}, nil
}
