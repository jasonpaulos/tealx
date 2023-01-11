package element

import (
	"encoding/xml"

	"github.com/jasonpaulos/tealx/language"
)

type Loop struct {
	Start     *Container
	Condition Element
	Step      *Container
	// Range *LoopRange
	Body Container
}

func (l *Loop) Codegen(ctx CodegenContext) language.ControlFlowGraph {
	graph := language.MakeControlFlowGraph(nil)

	if l.Start != nil {
		graph.Append(l.Start.Codegen(ctx))
	}

	conditionalGraph := l.Condition.Codegen(ctx)
	graph.AppendDistinct(conditionalGraph)

	bodyGraph := l.Body.Codegen(ctx)
	if l.Step != nil {
		// If step is present, perform it after the loop body
		bodyGraph.Append(l.Step.Codegen(ctx))
	}
	bodyGraph.AppendDistinct(conditionalGraph)

	graph.AppendConditionalRejoin(bodyGraph, true)

	return graph
}

func (l *Loop) Inner() []Element {
	inners := []Element{l.Body}
	if l.Start != nil {
		inners = append(inners, *l.Start)
	}
	if l.Condition != nil {
		inners = append(inners, l.Condition)
	}
	if l.Step != nil {
		inners = append(inners, *l.Step)
	}
	return inners
}

func (l *Loop) xml() xmlElement {
	var start *xmlContainer
	if l.Start != nil {
		tmp := l.Start.xmlContainer()
		start = &tmp
	}
	var condition *xmlContainer
	if l.Condition != nil {
		tmp := makeXmlContainer(l.Condition.xml())
		condition = &tmp
	}
	var step *xmlContainer
	if l.Step != nil {
		tmp := l.Step.xmlContainer()
		step = &tmp
	}
	return &xmlLoop{
		Start:     start,
		Condition: condition,
		Step:      step,
		Body:      l.Body.xmlContainer(),
	}
}

type xmlLoop struct {
	XMLName   xml.Name      `xml:"loop"`
	Start     *xmlContainer `xml:"start,omitempty"`
	Condition *xmlContainer `xml:"condition"`
	Step      *xmlContainer `xml:"step,omitempty"`
	Body      xmlContainer  `xml:"body"`
}

func (x *xmlLoop) element() (Element, error) {
	var start *Container
	if x.Start != nil {
		tmp, err := x.Start.containerElement()
		if err != nil {
			return nil, err
		}
		start = &tmp
	}
	condition, err := x.Condition.expectSingleElement()
	if err != nil {
		return nil, err
	}
	var step *Container
	if x.Step != nil {
		tmp, err := x.Step.containerElement()
		if err != nil {
			return nil, err
		}
		step = &tmp
	}
	body, err := x.Body.containerElement()
	if err != nil {
		return nil, err
	}

	return &Loop{
		Start:     start,
		Condition: condition,
		Step:      step,
		Body:      body,
	}, nil
}
