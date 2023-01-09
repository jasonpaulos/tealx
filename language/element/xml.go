package element

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
)

func MarshalXml(element Element) ([]byte, error) {
	return xml.Marshal(element.xml())
}

func UnmarshalXmlBytes(data []byte) (Element, error) {
	return UnmarshalXml(bytes.NewReader(data))
}

func UnmarshalXml(r io.Reader) (Element, error) {
	decoder := xml.NewDecoder(r)
	var start xml.StartElement
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			return nil, errors.New("no element found")
		}
		if err != nil {
			return nil, err
		}

		var ok bool
		start, ok = token.(xml.StartElement)
		if ok {
			break
		}
	}

	element, err := decodeXmlElement(decoder, start)
	if err != nil {
		return nil, err
	}
	return element.element()
}

func (x *xmlChildElement) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	return encoder.Encode(x.xmlElement)
}

func (x *xmlChildElement) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	element, err := decodeXmlElement(decoder, start)
	if err == nil {
		x.xmlElement = element
	}
	return err
}

func decodeXmlElement(decoder *xml.Decoder, start xml.StartElement) (xmlElement, error) {
	element := xmlTagToXmlElement(start.Name.Local)
	if element == nil {
		return nil, fmt.Errorf("unknown tag: %s", start.Name.Local)
	}
	err := decoder.DecodeElement(&element, &start)
	if err != nil {
		return nil, err
	}
	return element, nil
}

func xmlTagToXmlElement(tag string) xmlElement {
	switch tag {
	case "program":
		return &xmlProgram{}
	case "int":
		return &xmlInt{}
	case "bytes":
		return &xmlBytes{}
	case "equals":
		return &xmlEquals{}
	case "return":
		return &xmlReturn{}
	case "if":
		return &xmlIf{}
	case "match":
		return &xmlMatch{}
	default:
		return nil
	}
}
