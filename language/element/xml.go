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
	case "subroutine":
		return &xmlSubroutine{}
	case "int":
		return &xmlInt{}
	case "bytes":
		return &xmlBytes{}
	case "log":
		return &xmlLog{}
	case "subroutine-call":
		return &xmlSubroutineCall{}
	case "subroutine-return":
		return &xmlSubroutineReturn{}
	case "program-return":
		return &xmlProgramReturn{}
	case "variable-get":
		return &xmlVariableGet{}
	case "if":
		return &xmlIf{}
	case "match":
		return &xmlMatch{}
		// binary types
	case "subtract":
		return &xmlBinary{XMLName: BinaryTypeSubtraction.XmlName()}
	case "divide":
		return &xmlBinary{XMLName: BinaryTypeDivision.XmlName()}
	case "mod":
		return &xmlBinary{XMLName: BinaryTypeModulus.XmlName()}
	case "exp":
		return &xmlBinary{XMLName: BinaryTypeExponential.XmlName()}
	case "bitwise-and":
		return &xmlBinary{XMLName: BinaryTypeBitwiseAnd.XmlName()}
	case "bitwise-or":
		return &xmlBinary{XMLName: BinaryTypeBitwiseOr.XmlName()}
	case "bitwise-xor":
		return &xmlBinary{XMLName: BinaryTypeBitwiseXor.XmlName()}
	case "shift-left":
		return &xmlBinary{XMLName: BinaryTypeShiftLeft.XmlName()}
	case "shift-right":
		return &xmlBinary{XMLName: BinaryTypeShiftRight.XmlName()}
	case "equal":
		return &xmlBinary{XMLName: BinaryTypeEqual.XmlName()}
	case "not-equal":
		return &xmlBinary{XMLName: BinaryTypeNotEqual.XmlName()}
	case "less-than":
		return &xmlBinary{XMLName: BinaryTypeLessThan.XmlName()}
	case "less-than-or-equal":
		return &xmlBinary{XMLName: BinaryTypeLessThanOrEqual.XmlName()}
	case "greater-than":
		return &xmlBinary{XMLName: BinaryTypeGreaterThan.XmlName()}
	case "greater-than-or-equal":
		return &xmlBinary{XMLName: BinaryTypeGreaterThanOrEqual.XmlName()}
	case "get-bit":
		return &xmlBinary{XMLName: BinaryTypeGetBit.XmlName()}
	case "get-byte":
		return &xmlBinary{XMLName: BinaryTypeGetByte.XmlName()}
	case "bytes-add":
		return &xmlBinary{XMLName: BinaryTypeBytesAddition.XmlName()}
	case "bytes-subtract":
		return &xmlBinary{XMLName: BinaryTypeBytesSubtraction.XmlName()}
	case "bytes-divide":
		return &xmlBinary{XMLName: BinaryTypeBytesDivision.XmlName()}
	case "bytes-multiply":
		return &xmlBinary{XMLName: BinaryTypeBytesMultiplication.XmlName()}
	case "bytes-mod":
		return &xmlBinary{XMLName: BinaryTypeBytesModulus.XmlName()}
	case "bytes-and":
		return &xmlBinary{XMLName: BinaryTypeBytesAnd.XmlName()}
	case "bytes-or":
		return &xmlBinary{XMLName: BinaryTypeBytesOr.XmlName()}
	case "bytes-xor":
		return &xmlBinary{XMLName: BinaryTypeBytesXor.XmlName()}
	case "bytes-equal":
		return &xmlBinary{XMLName: BinaryTypeBytesEqual.XmlName()}
	case "bytes-not-equal":
		return &xmlBinary{XMLName: BinaryTypeBytesNotEqual.XmlName()}
	case "bytes-less-than":
		return &xmlBinary{XMLName: BinaryTypeBytesLessThan.XmlName()}
	case "bytes-less-than-or-equal":
		return &xmlBinary{XMLName: BinaryTypeBytesLessThanOrEqual.XmlName()}
	case "bytes-greater-than":
		return &xmlBinary{XMLName: BinaryTypeBytesGreaterThan.XmlName()}
	case "bytes-greater-than-or-equal":
		return &xmlBinary{XMLName: BinaryTypeBytesGreaterThanOrEqual.XmlName()}
	case "extract-uint16":
		return &xmlBinary{XMLName: BinaryTypeExtractUint16.XmlName()}
	case "extract-uint32":
		return &xmlBinary{XMLName: BinaryTypeExtractUint32.XmlName()}
	case "extract-uint64":
		return &xmlBinary{XMLName: BinaryTypeExtractUint64.XmlName()}
	case "add":
		return &xmlBinary{XMLName: BinaryTypeAddition.XmlName()}
	case "multiply":
		return &xmlBinary{XMLName: BinaryTypeMultiplication.XmlName()}
	case "and":
		return &xmlBinary{XMLName: BinaryTypeAnd.XmlName()}
	case "or":
		return &xmlBinary{XMLName: BinaryTypeOr.XmlName()}
	case "xor":
		return &xmlBinary{XMLName: BinaryTypeXor.XmlName()}
	case "concat":
		return &xmlBinary{XMLName: BinaryTypeConcat.XmlName()}
	default:
		return nil
	}
}
