package element

type Element interface {
	Inner() []Element
	xml() xmlElement
}

type xmlElement interface {
	element() (Element, error)
}

type emptyElement struct{}

func (e emptyElement) Inner() []Element {
	return nil
}

type Container struct {
	Children []Element
}

func (p Container) Inner() []Element {
	return p.Children
}

func (p Container) xmlContainer() xmlContainer {
	xmlChildren := make([]xmlChildElement, len(p.Children))
	for i, child := range p.Children {
		xmlChildren[i] = xmlChildElement{child.xml()}
	}
	return xmlContainer{
		Children: xmlChildren,
	}
}

func (p Container) xml() xmlElement {
	return p.xmlContainer()
}

type xmlChildElement struct {
	xmlElement
}

type xmlContainer struct {
	Children []xmlChildElement `xml:",any"`
}

func makeXmlContainer(elements ...xmlElement) xmlContainer {
	children := make([]xmlChildElement, len(elements))
	for i, element := range elements {
		children[i] = xmlChildElement{element}
	}
	return xmlContainer{Children: children}
}

func (x xmlContainer) containerElement() (Container, error) {
	elementChildren := make([]Element, len(x.Children))
	for i, inner := range x.Children {
		var err error
		elementChildren[i], err = inner.element()
		if err != nil {
			return Container{}, err
		}
	}
	return Container{
		Children: elementChildren,
	}, nil
}

func (x xmlContainer) element() (Element, error) {
	parentElement, err := x.containerElement()
	if err != nil {
		return nil, err
	}
	return parentElement, nil
}
