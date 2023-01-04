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

type parentElement struct {
	Children []Element
}

func (p parentElement) Inner() []Element {
	return p.Children
}

func (p parentElement) xmlParentElement() xmlParentElement {
	xmlChildren := make([]xmlChildElement, len(p.Children))
	for i, child := range p.Children {
		xmlChildren[i] = xmlChildElement{child.xml()}
	}
	return xmlParentElement{
		Children: xmlChildren,
	}
}

func (p parentElement) xml() xmlElement {
	return p.xmlParentElement()
}

type xmlChildElement struct {
	xmlElement
}

type xmlParentElement struct {
	Children []xmlChildElement `xml:",any"`
}

func (x xmlParentElement) parentElement() (parentElement, error) {
	elementChildren := make([]Element, len(x.Children))
	for i, inner := range x.Children {
		var err error
		elementChildren[i], err = inner.element()
		if err != nil {
			return parentElement{}, err
		}
	}
	return parentElement{
		Children: elementChildren,
	}, nil
}

func (x xmlParentElement) element() (Element, error) {
	parentElement, err := x.parentElement()
	if err != nil {
		return nil, err
	}
	return parentElement, nil
}
