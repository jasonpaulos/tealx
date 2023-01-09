package element

import (
	"encoding/xml"

	"github.com/jasonpaulos/tealx/language"
)

type Return struct {
	Value Element
}

func (r *Return) Codegen() language.ControlFlowGraph {
	value := r.Value.Codegen()
	returnStmt := language.MakeControlFlowGraph([]language.Operation{
		{
			Opcode: "return",
		},
	})

	value.Append(returnStmt)
	return value
}

func (r *Return) Inner() []Element {
	return []Element{r.Value}
}

func (r *Return) xml() xmlElement {
	return &xmlReturn{
		xmlContainer: makeXmlContainer(r.Value.xml()),
	}
}

type xmlReturn struct {
	xmlContainer

	XMLName xml.Name `xml:"return"`
}

func (x *xmlReturn) element() (Element, error) {
	value, err := x.xmlContainer.expectSingleElement()
	if err != nil {
		return nil, err
	}

	return &Return{
		Value: value,
	}, nil
}
