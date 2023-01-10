package element

import (
	"fmt"

	"github.com/jasonpaulos/tealx/language"
)

type Element interface {
	Inner() []Element
	Codegen(ctx CodegenContext) language.ControlFlowGraph
	xml() xmlElement
}

type CodegenContext struct {
	CurrentSubroutine *Subroutine
}

type xmlElement interface {
	element() (Element, error)
}

type emptyElement struct{}

func (e emptyElement) Inner() []Element {
	return nil
}

type Container struct {
	Children []Element
}

func (c Container) Inner() []Element {
	return c.Children
}

func (c Container) Codegen(ctx CodegenContext) language.ControlFlowGraph {
	graph := language.MakeControlFlowGraph(nil)
	for _, child := range c.Children {
		graph.Append(child.Codegen(ctx))
	}
	return graph
}

func (c Container) xmlContainer() xmlContainer {
	xmlChildren := make([]xmlChildElement, len(c.Children))
	for i, child := range c.Children {
		xmlChildren[i] = xmlChildElement{child.xml()}
	}
	return xmlContainer{
		Children: xmlChildren,
	}
}

func (c Container) xml() xmlElement {
	return c.xmlContainer()
}

type xmlChildElement struct {
	xmlElement
}

type xmlContainer struct {
	Children []xmlChildElement `xml:",any"`
}

func makeXmlContainer(elements ...xmlElement) xmlContainer {
	children := make([]xmlChildElement, len(elements))
	for i, element := range elements {
		children[i] = xmlChildElement{element}
	}
	return xmlContainer{Children: children}
}

func (x xmlContainer) expectNElements(n int) ([]Element, error) {
	containerElement, err := x.containerElement()
	if err != nil {
		return nil, err
	}
	if len(containerElement.Children) != n {
		return nil, fmt.Errorf("expected container to have %d child element(s), but got %d", n, len(containerElement.Children))
	}
	return containerElement.Children, nil
}

func (x xmlContainer) expectSingleElement() (Element, error) {
	expected, err := x.expectNElements(1)
	if err != nil {
		return nil, err
	}
	return expected[0], err
}

func (x xmlContainer) expectTwoElements() (Element, Element, error) {
	expected, err := x.expectNElements(2)
	if err != nil {
		return nil, nil, err
	}
	return expected[0], expected[1], err
}

func (x xmlContainer) containerElement() (Container, error) {
	elementChildren := make([]Element, len(x.Children))
	for i, inner := range x.Children {
		var err error
		elementChildren[i], err = inner.element()
		if err != nil {
			return Container{}, err
		}
	}
	return Container{
		Children: elementChildren,
	}, nil
}

func (x xmlContainer) element() (Element, error) {
	containerElement, err := x.containerElement()
	if err != nil {
		return nil, err
	}
	return containerElement, nil
}
