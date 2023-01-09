package element

import (
	"encoding/xml"

	"github.com/jasonpaulos/tealx/language"
)

type Match struct {
	Value       Element
	DefaultCase *Container
	Cases       []MatchCase
}

func (m *Match) Codegen() language.ControlFlowGraph {
	graph := m.Value.Codegen()
	targets := make([]language.ControlFlowGraph, len(m.Cases))
	for i, c := range m.Cases {
		graph.Append(c.Value.Codegen())
		targets[i] = c.Body.Codegen()
	}
	var defaultTarget language.ControlFlowGraph
	if m.DefaultCase == nil {
		defaultTarget = language.MakeControlFlowGraph(nil)
	} else {
		defaultTarget = m.DefaultCase.Codegen()
	}
	graph.AppendMatch(targets, defaultTarget)
	return graph
}

func (m *Match) Inner() []Element {
	numInners := 2*len(m.Cases) + 1
	if m.DefaultCase != nil {
		numInners++
	}
	inners := make([]Element, 0, numInners)
	inners = append(inners, m.Value)
	if m.DefaultCase != nil {
		inners = append(inners, *m.DefaultCase)
	}
	for _, c := range m.Cases {
		inners = append(inners, c.Value, c.Body)
	}
	return inners
}

func (m *Match) xml() xmlElement {
	cases := make([]xmlMatchCase, len(m.Cases))
	for i, c := range m.Cases {
		cases[i] = c.xml()
	}
	var defaultCase *xmlContainer
	if m.DefaultCase != nil {
		tmp := m.DefaultCase.xmlContainer()
		defaultCase = &tmp
	}
	return &xmlMatch{
		Value:       makeXmlContainer(m.Value.xml()),
		Cases:       cases,
		DefaultCase: defaultCase,
	}
}

type xmlMatch struct {
	XMLName     xml.Name       `xml:"match"`
	Value       xmlContainer   `xml:"value"`
	DefaultCase *xmlContainer  `xml:"default-case,omitempty"`
	Cases       []xmlMatchCase `xml:"case"`
}

func (x *xmlMatch) element() (Element, error) {
	value, err := x.Value.expectSingleElement()
	if err != nil {
		return nil, err
	}
	cases := make([]MatchCase, len(x.Cases))
	for i, c := range x.Cases {
		mc, err := c.matchCase()
		if err != nil {
			return nil, err
		}
		cases[i] = mc
	}
	var defaultCase *Container
	if x.DefaultCase != nil {
		tmp, err := x.DefaultCase.containerElement()
		if err != nil {
			return nil, err
		}
		defaultCase = &tmp
	}

	return &Match{
		Value:       value,
		Cases:       cases,
		DefaultCase: defaultCase,
	}, nil
}

type MatchCase struct {
	Value Element
	Body  Container
}

func (c *MatchCase) xml() xmlMatchCase {
	return xmlMatchCase{
		Value: makeXmlContainer(c.Value.xml()),
		Body:  c.Body.xmlContainer(),
	}
}

type xmlMatchCase struct {
	XMLName xml.Name     `xml:"case"`
	Value   xmlContainer `xml:"value,omitempty"`
	Body    xmlContainer `xml:"body"`
}

func (x *xmlMatchCase) matchCase() (MatchCase, error) {
	value, err := x.Value.expectSingleElement()
	if err != nil {
		return MatchCase{}, err
	}

	body, err := x.Body.containerElement()
	if err != nil {
		return MatchCase{}, err
	}

	return MatchCase{
		Value: value,
		Body:  body,
	}, nil
}
