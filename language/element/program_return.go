package element

import (
	"encoding/xml"

	"github.com/jasonpaulos/tealx/language"
)

type ProgramReturn struct {
	Value Element
}

func (r *ProgramReturn) Codegen(ctx CodegenContext) language.ControlFlowGraph {
	value := r.Value.Codegen(ctx)
	returnStmt := language.MakeControlFlowGraph([]language.Operation{
		{
			Opcode: "return",
		},
	})

	value.Append(returnStmt)
	return value
}

func (r *ProgramReturn) Inner() []Element {
	return []Element{r.Value}
}

func (r *ProgramReturn) xml() xmlElement {
	return &xmlProgramReturn{
		xmlContainer: makeXmlContainer(r.Value.xml()),
	}
}

type xmlProgramReturn struct {
	xmlContainer

	XMLName xml.Name `xml:"program-return"`
}

func (x *xmlProgramReturn) element() (Element, error) {
	value, err := x.xmlContainer.expectSingleElement()
	if err != nil {
		return nil, err
	}

	return &ProgramReturn{
		Value: value,
	}, nil
}
