package element

import (
	"encoding/xml"

	"github.com/jasonpaulos/tealx/language"
)

type SubroutineReturn struct {
	Value Element
}

func (r *SubroutineReturn) Codegen(ctx CodegenContext) language.ControlFlowGraph {
	graph := r.Value.Codegen(ctx)

	// TODO: if local variables are used, need to do `frame_bury 0` to put the
	// return value in the right place

	returnStmt := language.MakeControlFlowGraph([]language.Operation{
		{
			Opcode: "retsub",
		},
	})

	graph.Append(returnStmt)
	return graph
}

func (r *SubroutineReturn) Inner() []Element {
	return []Element{r.Value}
}

func (r *SubroutineReturn) xml() xmlElement {
	return &xmlSubroutineReturn{
		xmlContainer: makeXmlContainer(r.Value.xml()),
	}
}

type xmlSubroutineReturn struct {
	xmlContainer

	XMLName xml.Name `xml:"subroutine-return"`
}

func (x *xmlSubroutineReturn) element() (Element, error) {
	value, err := x.xmlContainer.expectSingleElement()
	if err != nil {
		return nil, err
	}

	return &SubroutineReturn{
		Value: value,
	}, nil
}
