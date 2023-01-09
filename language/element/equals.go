package element

import (
	"encoding/xml"

	"github.com/jasonpaulos/tealx/language"
)

type Equals struct {
	Left  Element
	Right Element
}

func (e *Equals) Codegen() language.ControlFlowGraph {
	left := e.Left.Codegen()
	right := e.Right.Codegen()
	equalityStmt := language.MakeControlFlowGraph([]language.Operation{
		{
			Opcode: "==",
		},
	})

	left.Append(right)
	left.Append(equalityStmt)
	return left
}

func (e *Equals) Inner() []Element {
	return []Element{e.Left, e.Right}
}

func (e *Equals) xml() xmlElement {
	return &xmlEquals{
		xmlContainer: makeXmlContainer(e.Left.xml(), e.Right.xml()),
	}
}

type xmlEquals struct {
	xmlContainer

	XMLName xml.Name `xml:"equals"`
}

func (x *xmlEquals) element() (Element, error) {
	left, right, err := x.xmlContainer.expectTwoElements()
	if err != nil {
		return nil, err
	}

	return &Equals{
		Left:  left,
		Right: right,
	}, nil
}
