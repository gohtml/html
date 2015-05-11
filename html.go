package html

import (
	. "github.com/gohtml/elements"
)

var (
	doctypeNode = "<!DOCTYPE html>"
)

type Html struct {
	Element
}

var _ Node = Html{}

// Implementation of Node interface
func (h Html) WriteTo(b Writer, opt RenderOptions, parent *Element, childIndex int) {
	b.WriteString(doctypeNode)
	b.WriteByte('\n')

	h.Element.WriteTo(b, opt, parent, childIndex)
}

// Implementation of Node interface
func (h Html) Type() TagType {
	return HTMLTag
}

func (h Html) Head() *Element {
	return h.children[0].(*Element)
}

func (h Html) Body() *Element {
	return h.children[1].(*Element)
}

func (h Html) Lang(lang string) Html {
	h.NonEmptyAttr("lang", lang)
	return h
}

func (h Html) Title(title string) Html {
	h.Head().Child(TITLE(title))
	return h
}

func (h Html) Favicon(href URL, tp string) Html {
	h.Head().Child(LINK(href, "shortcut icon").Attr("type", tp))
	return h
}

func (h Html) Css(href URL) Html {
	h.Head().Child(LINK(href, "stylesheet").Attr("type", "text/css"))
	return h
}

func (h Html) SCRIPT(src URL, content string) Html {
	h.Body().Child(SCRIPT(src, content))
	return h
}
