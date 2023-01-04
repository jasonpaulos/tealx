package element

import "encoding/xml"

type Program struct {
	parentElement

	Version uint64
}

func (p *Program) xml() xmlElement {
	return &xmlProgram{
		xmlParentElement: p.parentElement.xmlParentElement(),
		Version:          p.Version,
	}
}

type xmlProgram struct {
	xmlParentElement

	XMLName xml.Name `xml:"program"`
	Version uint64   `xml:"version,attr"`
}

func (x *xmlProgram) element() (Element, error) {
	parentElement, err := x.xmlParentElement.parentElement()
	if err != nil {
		return nil, err
	}

	return &Program{
		parentElement: parentElement,
		Version:       x.Version,
	}, nil
}
