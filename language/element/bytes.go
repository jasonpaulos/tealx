package element

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/jasonpaulos/tealx/language"
)

type Bytes struct {
	emptyElement

	Value []byte
}

func (b *Bytes) Codegen(ctx CodegenContext) language.ControlFlowGraph {
	return language.MakeControlFlowGraph([]language.Operation{
		{
			Opcode:    "byte",
			Arguments: []string{"0x" + hex.EncodeToString(b.Value)},
		},
	})
}

func (b *Bytes) xml() xmlElement {
	return &xmlBytes{Value: hex.EncodeToString(b.Value), Format: BytesFormatHex}
}

const (
	BytesFormatHex    string = "hex"
	BytesFormatBase64 string = "base64"
	BytesFormatUTF8   string = "utf-8"
)

type xmlBytes struct {
	XMLName xml.Name `xml:"bytes"`
	Value   string   `xml:"value,attr"`
	Format  string   `xml:"format,attr"`
}

func (x *xmlBytes) element() (Element, error) {
	var decoded []byte
	var err error

	switch strings.ToLower(x.Format) {
	case BytesFormatHex:
		decoded, err = hex.DecodeString(x.Value)
	case BytesFormatBase64:
		decoded, err = base64.StdEncoding.DecodeString(x.Value)
	case BytesFormatUTF8:
		decoded = []byte(x.Value)
	default:
		return nil, fmt.Errorf("unknown format for bytes: \"%s\"", x.Format)
	}

	if err != nil {
		return nil, err
	}

	return &Bytes{
		Value: decoded,
	}, nil
}
