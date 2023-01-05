package element

import (
	"encoding/xml"
	"fmt"
)

type If struct {
	Condition Element   `xml:"condition"`
	Then      Container `xml:"then"`
	Else      Container `xml:"else"`
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
	condition, err := x.Condition.containerElement()
	if err != nil {
		return nil, err
	}
	if len(condition.Children) != 1 {
		return nil, fmt.Errorf("expected condition to have exactly 1 child, but got %d", len(condition.Children))
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
		Condition: condition.Children[0],
		Then:      thenBranch,
		Else:      elseBranch,
	}, nil
}
