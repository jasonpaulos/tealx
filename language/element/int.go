package element

import (
	"encoding/xml"
	"strconv"

	"github.com/jasonpaulos/tealx/language"
)

type Int struct {
	emptyElement

	Value uint64
}

func (i *Int) Codegen(ctx CodegenContext) language.ControlFlowGraph {
	return language.MakeControlFlowGraph([]language.Operation{
		{
			Opcode:    "int",
			Arguments: []string{strconv.FormatUint(i.Value, 10)},
		},
	})
}

func (i *Int) xml() xmlElement {
	return &xmlInt{Value: i.Value}
}

type xmlInt struct {
	XMLName xml.Name `xml:"int"`
	Value   uint64   `xml:"value,attr"`
}

func (x *xmlInt) element() (Element, error) {
	return &Int{
		Value: x.Value,
	}, nil
}
