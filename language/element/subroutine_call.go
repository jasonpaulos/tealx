package element

import (
	"encoding/xml"

	"github.com/jasonpaulos/tealx/language"
)

type SubroutineCall struct {
	Name      string
	Arguments []Element
}

func (c *SubroutineCall) Codegen(ctx CodegenContext) language.ControlFlowGraph {
	graph := language.MakeControlFlowGraph(nil)
	for _, arg := range c.Arguments {
		graph.Append(arg.Codegen(ctx))
	}

	graph.Append(language.MakeControlFlowGraph([]language.Operation{
		{Opcode: "callsub", Arguments: []string{"sub_" + c.Name}},
	}))

	return graph
}

func (c *SubroutineCall) Inner() []Element {
	return c.Arguments
}

func (c *SubroutineCall) xml() xmlElement {
	args := make([]xmlContainer, len(c.Arguments))
	for i, arg := range c.Arguments {
		args[i] = makeXmlContainer(arg.xml())
	}
	return &xmlSubroutineCall{
		Name:      c.Name,
		Arguments: args,
	}
}

type xmlSubroutineCall struct {
	XMLName   xml.Name       `xml:"subroutine-call"`
	Name      string         `xml:"name,attr"`
	Arguments []xmlContainer `xml:"argument"`
}

func (x *xmlSubroutineCall) element() (Element, error) {
	args := make([]Element, len(x.Arguments))
	for i, arg := range x.Arguments {
		var err error
		args[i], err = arg.expectSingleElement()
		if err != nil {
			return nil, err
		}
	}
	return &SubroutineCall{
		Name:      x.Name,
		Arguments: args,
	}, nil
}
