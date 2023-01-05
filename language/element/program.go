package element

import "encoding/xml"

type Program struct {
	Container

	Version uint64
}

func (p *Program) xml() xmlElement {
	return &xmlProgram{
		xmlContainer: p.Container.xmlContainer(),
		Version:      p.Version,
	}
}

type xmlProgram struct {
	xmlContainer

	XMLName xml.Name `xml:"program"`
	Version uint64   `xml:"version,attr"`
}

func (x *xmlProgram) element() (Element, error) {
	parentElement, err := x.xmlContainer.containerElement()
	if err != nil {
		return nil, err
	}

	return &Program{
		Container: parentElement,
		Version:   x.Version,
	}, nil
}
