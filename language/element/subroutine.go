package element

import (
	"encoding/xml"
	"strconv"

	"github.com/jasonpaulos/tealx/language"
)

type Subroutine struct {
	Name      string
	Body      Container
	Return    *SubroutineReturnInfo
	Arguments []SubroutineArgumentInfo
}

func (s *Subroutine) Codegen(ctx CodegenContext) language.ControlFlowGraph {
	numReturns := 0
	if s.Return != nil {
		numReturns = 1
	}
	graph := language.MakeControlFlowGraph([]language.Operation{
		{
			Opcode:    "proto",
			Arguments: []string{strconv.Itoa(len(s.Arguments)), strconv.Itoa(numReturns)},
		},
	})

	ctx.CurrentSubroutine = s

	graph.Append(s.Body.Codegen(ctx))
	return graph
}

func (s *Subroutine) Inner() []Element {
	return []Element{s.Body}
}

func (s *Subroutine) xmlSubroutine() xmlSubroutine {
	args := make([]xmlSubroutineArgumentInfo, len(s.Arguments))
	for i, arg := range s.Arguments {
		args[i] = arg.xml()
	}
	var subroutineReturn *xmlSubroutineReturnInfo
	if s.Return != nil {
		tmp := s.Return.xml()
		subroutineReturn = &tmp
	}
	return xmlSubroutine{
		Name:      s.Name,
		Body:      s.Body.xmlContainer(),
		Return:    subroutineReturn,
		Arguments: args,
	}
}

func (s *Subroutine) xml() xmlElement {
	value := s.xmlSubroutine()
	return &value
}

type SubroutineArgumentInfo struct {
	Name string
	Type string
}

func (a *SubroutineArgumentInfo) xml() xmlSubroutineArgumentInfo {
	return xmlSubroutineArgumentInfo{
		Name: a.Name,
		Type: a.Type,
	}
}

type SubroutineReturnInfo struct {
	Type string
}

func (r *SubroutineReturnInfo) xml() xmlSubroutineReturnInfo {
	return xmlSubroutineReturnInfo{
		Type: r.Type,
	}
}

type xmlSubroutine struct {
	XMLName   xml.Name                    `xml:"subroutine"`
	Name      string                      `xml:"name,attr"`
	Body      xmlContainer                `xml:"body"`
	Return    *xmlSubroutineReturnInfo    `xml:"returns,omitempty"`
	Arguments []xmlSubroutineArgumentInfo `xml:"argument"`
}

func (x *xmlSubroutine) subroutine() (*Subroutine, error) {
	body, err := x.Body.containerElement()
	if err != nil {
		return nil, err
	}
	args := make([]SubroutineArgumentInfo, len(x.Arguments))
	for i, arg := range x.Arguments {
		args[i] = arg.subroutineArgumentInfo()
	}
	var subroutineReturn *SubroutineReturnInfo
	if x.Return != nil {
		tmp := x.Return.subroutineReturnInfo()
		subroutineReturn = &tmp
	}

	return &Subroutine{
		Name:      x.Name,
		Body:      body,
		Return:    subroutineReturn,
		Arguments: args,
	}, nil
}

func (x *xmlSubroutine) element() (Element, error) {
	element, err := x.subroutine()
	if err != nil {
		return nil, err
	}
	return element, nil
}

type xmlSubroutineArgumentInfo struct {
	XMLName xml.Name `xml:"argument"`
	Name    string   `xml:"name,attr"`
	Type    string   `xml:"type,attr"`
}

func (x *xmlSubroutineArgumentInfo) subroutineArgumentInfo() SubroutineArgumentInfo {
	return SubroutineArgumentInfo{
		Name: x.Name,
		Type: x.Type,
	}
}

type xmlSubroutineReturnInfo struct {
	XMLName xml.Name `xml:"returns"`
	Type    string   `xml:"type,attr"`
}

func (x *xmlSubroutineReturnInfo) subroutineReturnInfo() SubroutineReturnInfo {
	return SubroutineReturnInfo{
		Type: x.Type,
	}
}
