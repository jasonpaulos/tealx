package element

import "encoding/xml"

type Int struct {
	emptyElement

	Value uint64
}

func (c *Int) xml() xmlElement {
	return &xmlInt{Value: c.Value}
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
