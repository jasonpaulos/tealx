package element

import (
	"encoding/xml"

	"github.com/jasonpaulos/tealx/language"
)

type Program struct {
	Version     uint64
	Main        Container
	Subroutines []*Subroutine
}

func (p *Program) Codegen(ctx CodegenContext) language.ControlFlowGraph {
	// subroutines are handled separately
	return p.Main.Codegen(ctx)
}

func (p *Program) Inner() []Element {
	elements := make([]Element, 0, len(p.Subroutines)+1)
	elements = append(elements, p.Main)
	for _, subroutine := range p.Subroutines {
		elements = append(elements, subroutine)
	}
	return elements
}

func (p *Program) xml() xmlElement {
	subroutines := make([]xmlSubroutine, len(p.Subroutines))
	for i, sub := range p.Subroutines {
		subroutines[i] = sub.xmlSubroutine()
	}
	return &xmlProgram{
		Version:     p.Version,
		Main:        p.Main.xmlContainer(),
		Subroutines: subroutines,
	}
}

type xmlProgram struct {
	XMLName     xml.Name        `xml:"program"`
	Version     uint64          `xml:"version,attr"`
	Main        xmlContainer    `xml:"main"`
	Subroutines []xmlSubroutine `xml:"subroutine"`
}

func (x *xmlProgram) element() (Element, error) {
	main, err := x.Main.containerElement()
	if err != nil {
		return nil, err
	}
	subroutines := make([]*Subroutine, len(x.Subroutines))
	for i, sub := range x.Subroutines {
		var err error
		subroutines[i], err = sub.subroutine()
		if err != nil {
			return nil, err
		}
	}

	return &Program{
		Version:     x.Version,
		Main:        main,
		Subroutines: subroutines,
	}, nil
}
