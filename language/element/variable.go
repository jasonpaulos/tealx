package element

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/jasonpaulos/tealx/language"
)

type VariableType int

const (
	VariableTypeUint64 VariableType = iota
	VariableTypeBytes
)

func VariableTypeFromString(s string) (VariableType, error) {
	switch s {
	case "uint64":
		return VariableTypeUint64, nil
	case "bytes":
		return VariableTypeBytes, nil
	default:
		return -1, fmt.Errorf("unknown VariableType string: %s", s)
	}
}

func (t VariableType) String() string {
	switch t {
	case VariableTypeUint64:
		return "uint64"
	case VariableTypeBytes:
		return "bytes"
	default:
		return fmt.Sprintf("unknown(%d)", t)
	}
}

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

type VariableSet struct {
	emptyElement

	Name  string
	Value Element
}

func (v *VariableSet) Codegen(ctx CodegenContext) language.ControlFlowGraph {
	// search the context to figure out how to load the variable

	graph := v.Value.Codegen(ctx)

	if ctx.CurrentSubroutine != nil {
		for i, arg := range ctx.CurrentSubroutine.Arguments {
			if arg.Name == v.Name {
				frameIndex := i - len(ctx.CurrentSubroutine.Arguments)
				graph.Append(language.MakeControlFlowGraph([]language.Operation{
					{
						Opcode:    "frame_bury",
						Arguments: []string{strconv.Itoa(frameIndex)},
					},
				}))
				return graph
			}
		}
	}

	panic(fmt.Sprintf("variable not found: %s", v.Name))
}

func (v *VariableSet) xml() xmlElement {
	return &xmlVariableSet{Name: v.Name, xmlContainer: makeXmlContainer(v.Value.xml())}
}

type xmlVariableSet struct {
	xmlContainer

	XMLName xml.Name `xml:"variable-set"`
	Name    string   `xml:"name,attr"`
}

func (x *xmlVariableSet) element() (Element, error) {
	value, err := x.xmlContainer.expectSingleElement()
	if err != nil {
		return nil, err
	}
	return &VariableSet{
		Name:  x.Name,
		Value: value,
	}, nil
}
