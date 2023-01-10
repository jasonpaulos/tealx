package element

import (
	"encoding/xml"

	"github.com/jasonpaulos/tealx/language"
)

type SubroutineCall struct {
	Name      string
	Arguments []SubroutineArgument
}

func (c *SubroutineCall) Codegen(ctx CodegenContext) language.ControlFlowGraph {
	graph := language.MakeControlFlowGraph(nil)
	for _, arg := range c.Arguments {
		graph.Append(arg.Value.Codegen(ctx))
	}

	graph.Append(language.MakeControlFlowGraph([]language.Operation{
		{Opcode: "callsub", Arguments: []string{"sub_" + c.Name}},
	}))

	return graph
}

func (c *SubroutineCall) Inner() []Element {
	inners := make([]Element, len(c.Arguments))
	for i, arg := range c.Arguments {
		inners[i] = arg.Value
	}
	return inners
}

func (c *SubroutineCall) xml() xmlElement {
	args := make([]xmlSubroutineArgument, len(c.Arguments))
	for i, arg := range c.Arguments {
		args[i] = arg.xml()
	}
	return &xmlSubroutineCall{
		Name:      c.Name,
		Arguments: args,
	}
}

type SubroutineArgument struct {
	Value Element
}

func (a *SubroutineArgument) xml() xmlSubroutineArgument {
	return xmlSubroutineArgument{
		xmlContainer: makeXmlContainer(a.Value.xml()),
	}
}

type xmlSubroutineCall struct {
	XMLName   xml.Name                `xml:"subroutine-call"`
	Name      string                  `xml:"name,attr"`
	Arguments []xmlSubroutineArgument `xml:"argument"`
}

func (x *xmlSubroutineCall) element() (Element, error) {
	args := make([]SubroutineArgument, len(x.Arguments))
	for i, arg := range x.Arguments {
		var err error
		args[i], err = arg.subroutineArgument()
		if err != nil {
			return nil, err
		}
	}
	return &SubroutineCall{
		Name:      x.Name,
		Arguments: args,
	}, nil
}

type xmlSubroutineArgument struct {
	xmlContainer

	XMLName xml.Name `xml:"argument"`
}

func (x *xmlSubroutineArgument) subroutineArgument() (SubroutineArgument, error) {
	value, err := x.xmlContainer.expectSingleElement()
	if err != nil {
		return SubroutineArgument{}, err
	}
	return SubroutineArgument{
		Value: value,
	}, nil
}
