package element

import (
	"encoding/xml"
	"fmt"

	"github.com/jasonpaulos/tealx/language"
)

type BinaryType int

const (
	BinaryTypeSubtraction BinaryType = iota
	BinaryTypeDivision
	BinaryTypeModulus
	BinaryTypeExponential
	BinaryTypeBitwiseAnd
	BinaryTypeBitwiseOr
	BinaryTypeBitwiseXor
	BinaryTypeShiftLeft
	BinaryTypeShiftRight
	BinaryTypeEqual
	BinaryTypeNotEqual
	BinaryTypeLessThan
	BinaryTypeLessThanOrEqual
	BinaryTypeGreaterThan
	BinaryTypeGreaterThanOrEqual
	BinaryTypeGetBit
	BinaryTypeGetByte
	BinaryTypeBytesAddition
	BinaryTypeBytesSubtraction
	BinaryTypeBytesDivision
	BinaryTypeBytesMultiplication
	BinaryTypeBytesModulus
	BinaryTypeBytesAnd
	BinaryTypeBytesOr
	BinaryTypeBytesXor
	BinaryTypeBytesEqual
	BinaryTypeBytesNotEqual
	BinaryTypeBytesLessThan
	BinaryTypeBytesLessThanOrEqual
	BinaryTypeBytesGreaterThan
	BinaryTypeBytesGreaterThanOrEqual
	BinaryTypeExtractUint16
	BinaryTypeExtractUint32
	BinaryTypeExtractUint64

	// these can be made into "n-ary" expressions, not just binary
	BinaryTypeAddition
	BinaryTypeMultiplication
	BinaryTypeAnd
	BinaryTypeOr
	BinaryTypeXor
	BinaryTypeConcat

	binaryTypeFinal
)

func BinaryTypeFromXmlName(name string) (BinaryType, error) {
	switch name {
	case "subtract":
		return BinaryTypeSubtraction, nil
	case "divide":
		return BinaryTypeDivision, nil
	case "mod":
		return BinaryTypeModulus, nil
	case "exp":
		return BinaryTypeExponential, nil
	case "bitwise-and":
		return BinaryTypeBitwiseAnd, nil
	case "bitwise-or":
		return BinaryTypeBitwiseOr, nil
	case "bitwise-xor":
		return BinaryTypeBitwiseXor, nil
	case "shift-left":
		return BinaryTypeShiftLeft, nil
	case "shift-right":
		return BinaryTypeShiftRight, nil
	case "equal":
		return BinaryTypeEqual, nil
	case "not-equal":
		return BinaryTypeNotEqual, nil
	case "less-than":
		return BinaryTypeLessThan, nil
	case "less-than-or-equal":
		return BinaryTypeLessThanOrEqual, nil
	case "greater-than":
		return BinaryTypeGreaterThan, nil
	case "greater-than-or-equal":
		return BinaryTypeGreaterThanOrEqual, nil
	case "get-bit":
		return BinaryTypeGetBit, nil
	case "get-byte":
		return BinaryTypeGetByte, nil
	case "bytes-add":
		return BinaryTypeBytesAddition, nil
	case "bytes-subtract":
		return BinaryTypeBytesSubtraction, nil
	case "bytes-divide":
		return BinaryTypeBytesDivision, nil
	case "bytes-multiply":
		return BinaryTypeBytesMultiplication, nil
	case "bytes-mod":
		return BinaryTypeBytesModulus, nil
	case "bytes-and":
		return BinaryTypeBytesAnd, nil
	case "bytes-or":
		return BinaryTypeBytesOr, nil
	case "bytes-xor":
		return BinaryTypeBytesXor, nil
	case "bytes-equal":
		return BinaryTypeBytesEqual, nil
	case "bytes-not-equal":
		return BinaryTypeBytesNotEqual, nil
	case "bytes-less-than":
		return BinaryTypeBytesLessThan, nil
	case "bytes-less-than-or-equal":
		return BinaryTypeBytesLessThanOrEqual, nil
	case "bytes-greater-than":
		return BinaryTypeBytesGreaterThan, nil
	case "bytes-greater-than-or-equal":
		return BinaryTypeBytesGreaterThanOrEqual, nil
	case "extract-uint16":
		return BinaryTypeExtractUint16, nil
	case "extract-uint32":
		return BinaryTypeExtractUint32, nil
	case "extract-uint64":
		return BinaryTypeExtractUint64, nil
	case "add":
		return BinaryTypeAddition, nil
	case "multiply":
		return BinaryTypeMultiplication, nil
	case "and":
		return BinaryTypeAnd, nil
	case "or":
		return BinaryTypeOr, nil
	case "xor":
		return BinaryTypeXor, nil
	case "concat":
		return BinaryTypeConcat, nil
	default:
		return -1, fmt.Errorf("unknown binary type name: %s", name)
	}
}

func (t BinaryType) Opcode() string {
	switch t {
	case BinaryTypeSubtraction:
		return "-"
	case BinaryTypeDivision:
		return "/"
	case BinaryTypeModulus:
		return "%"
	case BinaryTypeExponential:
		return "exp"
	case BinaryTypeBitwiseAnd:
		return "&"
	case BinaryTypeBitwiseOr:
		return "|"
	case BinaryTypeBitwiseXor:
		return "^"
	case BinaryTypeShiftLeft:
		return "<<"
	case BinaryTypeShiftRight:
		return ">>"
	case BinaryTypeEqual:
		return "=="
	case BinaryTypeNotEqual:
		return "!="
	case BinaryTypeLessThan:
		return "<"
	case BinaryTypeLessThanOrEqual:
		return "<="
	case BinaryTypeGreaterThan:
		return ">"
	case BinaryTypeGreaterThanOrEqual:
		return ">="
	case BinaryTypeGetBit:
		return "getbit"
	case BinaryTypeGetByte:
		return "getbyte"
	case BinaryTypeBytesAddition:
		return "b+"
	case BinaryTypeBytesSubtraction:
		return "b-"
	case BinaryTypeBytesDivision:
		return "b/"
	case BinaryTypeBytesMultiplication:
		return "b*"
	case BinaryTypeBytesModulus:
		return "b%"
	case BinaryTypeBytesAnd:
		return "b&"
	case BinaryTypeBytesOr:
		return "b|"
	case BinaryTypeBytesXor:
		return "b^"
	case BinaryTypeBytesEqual:
		return "b=="
	case BinaryTypeBytesNotEqual:
		return "b!="
	case BinaryTypeBytesLessThan:
		return "b<"
	case BinaryTypeBytesLessThanOrEqual:
		return "b<="
	case BinaryTypeBytesGreaterThan:
		return "b>"
	case BinaryTypeBytesGreaterThanOrEqual:
		return "b>="
	case BinaryTypeExtractUint16:
		return "extract_uint16"
	case BinaryTypeExtractUint32:
		return "extract_uint32"
	case BinaryTypeExtractUint64:
		return "extract_uint64"
	case BinaryTypeAddition:
		return "+"
	case BinaryTypeMultiplication:
		return "*"
	case BinaryTypeAnd:
		return "&&"
	case BinaryTypeOr:
		return "||"
	case BinaryTypeXor:
		return "^"
	case BinaryTypeConcat:
		return "concat"
	default:
		panic(fmt.Sprintf("unknown binary type: %v", t))
	}
}

func (t BinaryType) XmlName() xml.Name {
	var name string
	switch t {
	case BinaryTypeSubtraction:
		name = "subtract"
	case BinaryTypeDivision:
		name = "divide"
	case BinaryTypeModulus:
		name = "mod"
	case BinaryTypeExponential:
		name = "exp"
	case BinaryTypeBitwiseAnd:
		name = "bitwise-and"
	case BinaryTypeBitwiseOr:
		name = "bitwise-or"
	case BinaryTypeBitwiseXor:
		name = "bitwise-xor"
	case BinaryTypeShiftLeft:
		name = "shift-left"
	case BinaryTypeShiftRight:
		name = "shift-right"
	case BinaryTypeEqual:
		name = "equal"
	case BinaryTypeNotEqual:
		name = "not-equal"
	case BinaryTypeLessThan:
		name = "less-than"
	case BinaryTypeLessThanOrEqual:
		name = "less-than-or-equal"
	case BinaryTypeGreaterThan:
		name = "greater-than"
	case BinaryTypeGreaterThanOrEqual:
		name = "greater-than-or-equal"
	case BinaryTypeGetBit:
		name = "get-bit"
	case BinaryTypeGetByte:
		name = "get-byte"
	case BinaryTypeBytesAddition:
		name = "bytes-add"
	case BinaryTypeBytesSubtraction:
		name = "bytes-subtract"
	case BinaryTypeBytesDivision:
		name = "bytes-divide"
	case BinaryTypeBytesMultiplication:
		name = "bytes-multiply"
	case BinaryTypeBytesModulus:
		name = "bytes-mod"
	case BinaryTypeBytesAnd:
		name = "bytes-and"
	case BinaryTypeBytesOr:
		name = "bytes-or"
	case BinaryTypeBytesXor:
		name = "bytes-xor"
	case BinaryTypeBytesEqual:
		name = "bytes-equal"
	case BinaryTypeBytesNotEqual:
		name = "bytes-not-equal"
	case BinaryTypeBytesLessThan:
		name = "bytes-less-than"
	case BinaryTypeBytesLessThanOrEqual:
		name = "bytes-less-than-or-equal"
	case BinaryTypeBytesGreaterThan:
		name = "bytes-greater-than"
	case BinaryTypeBytesGreaterThanOrEqual:
		name = "bytes-greater-than-or-equal"
	case BinaryTypeExtractUint16:
		name = "extract-uint16"
	case BinaryTypeExtractUint32:
		name = "extract-uint32"
	case BinaryTypeExtractUint64:
		name = "extract-uint64"
	case BinaryTypeAddition:
		name = "add"
	case BinaryTypeMultiplication:
		name = "multiply"
	case BinaryTypeAnd:
		name = "and"
	case BinaryTypeOr:
		name = "or"
	case BinaryTypeXor:
		name = "xor"
	case BinaryTypeConcat:
		name = "concat"
	default:
		panic(fmt.Sprintf("unknown binary type: %v", t))
	}
	return xml.Name{Local: name}
}

type Binary struct {
	Type  BinaryType
	Left  Element
	Right Element
}

func (b *Binary) Codegen(ctx CodegenContext) language.ControlFlowGraph {
	left := b.Left.Codegen(ctx)
	right := b.Right.Codegen(ctx)
	binaryStmt := language.MakeControlFlowGraph([]language.Operation{
		{
			Opcode: b.Type.Opcode(),
		},
	})

	left.Append(right)
	left.Append(binaryStmt)
	return left
}

func (b *Binary) Inner() []Element {
	return []Element{b.Left, b.Right}
}

func (b *Binary) xml() xmlElement {
	return &xmlBinary{
		XMLName:      b.Type.XmlName(),
		xmlContainer: makeXmlContainer(b.Left.xml(), b.Right.xml()),
	}
}

type xmlBinary struct {
	xmlContainer

	XMLName xml.Name
}

func (x *xmlBinary) element() (Element, error) {
	binaryType, err := BinaryTypeFromXmlName(x.XMLName.Local)
	if err != nil {
		return nil, err
	}

	left, right, err := x.xmlContainer.expectTwoElements()
	if err != nil {
		return nil, err
	}

	return &Binary{
		Type:  binaryType,
		Left:  left,
		Right: right,
	}, nil
}
