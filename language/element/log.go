package element

import (
	"encoding/xml"

	"github.com/jasonpaulos/tealx/language"
)

type Log struct {
	Value Element
}

func (l *Log) Codegen(ctx CodegenContext) language.ControlFlowGraph {
	graph := l.Value.Codegen(ctx)
	logStmt := language.MakeControlFlowGraph([]language.Operation{
		{
			Opcode: "log",
		},
	})
	graph.Append(logStmt)
	return graph
}

func (l *Log) Inner() []Element {
	return []Element{l.Value}
}

func (l *Log) xml() xmlElement {
	return &xmlLog{
		xmlContainer: makeXmlContainer(l.Value.xml()),
	}
}

type xmlLog struct {
	xmlContainer

	XMLName xml.Name `xml:"log"`
}

func (x *xmlLog) element() (Element, error) {
	value, err := x.xmlContainer.expectSingleElement()
	if err != nil {
		return nil, err
	}

	return &Log{
		Value: value,
	}, nil
}
