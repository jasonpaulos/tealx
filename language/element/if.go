package element

import (
	"encoding/xml"

	"github.com/jasonpaulos/tealx/language"
)

type If struct {
	Condition Element
	Then      Container
	Else      Container
}

func (i *If) Codegen() language.ControlFlowGraph {
	conditionSubgraph := i.Condition.Codegen()
	thenSubgraph := i.Then.Codegen()
	elseSubgraph := i.Else.Codegen()

	conditionSubgraph.AppendConditional(thenSubgraph, elseSubgraph)
	return conditionSubgraph
}

func (i *If) Inner() []Element {
	return []Element{i.Condition, i.Then, i.Else}
}

func (i *If) xml() xmlElement {
	return &xmlIf{
		Condition: makeXmlContainer(i.Condition.xml()),
		Then:      i.Then.xmlContainer(),
		Else:      i.Else.xmlContainer(),
	}
}

type xmlIf struct {
	XMLName   xml.Name     `xml:"if"`
	Condition xmlContainer `xml:"condition"`
	Then      xmlContainer `xml:"then"`
	Else      xmlContainer `xml:"else"`
}

func (x *xmlIf) element() (Element, error) {
	condition, err := x.Condition.expectSingleElement()
	if err != nil {
		return nil, err
	}
	thenBranch, err := x.Then.containerElement()
	if err != nil {
		return nil, err
	}
	elseBranch, err := x.Else.containerElement()
	if err != nil {
		return nil, err
	}

	return &If{
		Condition: condition,
		Then:      thenBranch,
		Else:      elseBranch,
	}, nil
}
