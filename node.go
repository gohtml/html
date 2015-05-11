package html

import (
	"fmt"
	"io"
	"strconv"

	"github.com/golangplus/bytes"
	"github.com/golangplus/strings"

	. "github.com/gohtml/elements"
	"github.com/gohtml/utils"
)

type RenderOptions struct {
	Ident HTMLNode
	// Don't apply HTML5 omitting.
	DisableOmit bool
	// Sort attributes names before export.
	// This is useful for testing because otherwise the exported attributes could be unpredictable.
	SortAttr bool
}

// The default RenderOptions
var DefaultOptions = RenderOptions{}

type Writer interface {
	io.ByteWriter
	io.Writer
	WriteString(string) (int, error)
}

type Node interface {
	Type() TagType
	WriteTo(w Writer, opt RenderOptions, parent *Element, childIndex int)
}

// NodeToHTMLBytes converts a Node into HTMLNode.
func NodeToHTMLNode(nd Node, opt RenderOptions) HTMLNode {
	var b bytesp.ByteSlice
	nd.WriteTo(&b, opt, nil, 0)
	return HTMLNode(b)
}

// A string type representing an escaped HTML code
type HTMLNode string

var _ Node = HTMLNode("")

func (h HTMLNode) WriteTo(b Writer, opt RenderOptions, parent *Element, childIndex int) {
	b.WriteString(string(h))
}

func (h HTMLNode) WriteRaw(b Writer) {
	b.WriteString(string(h))
}

func (h HTMLNode) Type() TagType {
	return TextType
}

// An HTML void element
type Void struct {
	tagType    TagType
	attributes Attributes
	classes    htmlNodeSet
}

var _ Node = (*Void)(nil)

func (v *Void) Type() TagType {
	return v.tagType
}

// Implementation of Node.WriteTo. This will be called to generate open tags of both void and normal elements
func (v *Void) WriteTo(b Writer, opt RenderOptions, parent *Element, childIndex int) {
	b.WriteByte('<')
	b.WriteString(TagNames[v.tagType])

	if len(v.classes) > 0 {
		b.WriteString(` class="`)
		v.classes[0].WriteRaw(b)
		for i, n := 1, len(v.classes); i < n; i++ {
			b.WriteByte(' ')
			v.classes[i].WriteRaw(b)
		}
		b.WriteByte('"')
	}
	v.attributes.WriteTo(b, opt.SortAttr)

	b.WriteByte('>')
}

func (v *Void) Name() HTMLNode {
	return HTMLNode(TagNames[v.tagType])
}

func (v *Void) Attr(name string, value string) *Void {
	return v.attrOfEscaped(HTMLNode(utils.NormAttrName(name)), HTMLNode(utils.EscapeAttr(value)))
}

func (v *Void) AttrIfNotEmpty(name, value string) *Void {
	if value == "" {
		return v
	}
	return v.Attr(name, value)
}

func (v *Void) attrOfEscaped(name, value HTMLNode) *Void {
	if len(name) == 0 {
		// ignore empty name
		return v
	}

	if name == "class" {
		// classes are stored in the Void.classes field
		v.classes = v.classes[:0]
		stringsp.CallbackFields(string(value), func(n int) {
			if cap(v.classes) < n {
				v.classes = make([]HTMLNode, 0, n)
			}
		}, func(f string) {
			v.classes = append(v.classes, HTMLNode(f))
		})
		return v
	}

	v.attributes.Put(name, value)
	return v
}

// Title sets the title attribute of the node.
func (v *Void) Title(title string) {
	v.Attr("title", title)
}

// TabIndex sets the "tabindex" attribute of the node
func (v *Void) TabIndex(tablInex int) {
	v.Attr("tabindex", strconv.Itoa(tablInex))
}

// NonEmptyAttr sets the attribute is value is not empty.
func (v *Void) NonEmptyAttr(name, value string) *Void {
	if value == "" {
		return v
	}
	return v.Attr(name, value)
}

// AddClass adds a class into the class list of the node.
func (v *Void) AddClass(classes ...string) *Void {
	for _, cls := range classes {
		v.classes.Put(HTMLNode(utils.EscapeAttr(cls)))
	}

	return v
}

// DelClass deletes a class from the class list of the node.
func (v *Void) DelClass(classes ...string) *Void {
	for _, cls := range classes {
		v.classes.Del(HTMLNode(utils.EscapeAttr(cls)))
	}

	return v
}

// ID sets the "id" attribute of the node.
func (t *Void) ID(id string) *Void {
	return t.Attr("id", id)
}

// Element is a Node with children.
type Element struct {
	Void
	children []Node
}

var _ Node = (*Element)(nil)

// Children returns the children of the element.
func (t *Element) Children() []Element {
	return t.Children()
}

// Attr is same as Void.Attr but returns a *Element.
func (e *Element) Attr(name string, value string) *Element {
	e.Void.Attr(name, value)
	return e
}

// NonEmptyAttr is same as Void.NonEmptyAttr but returns a *Element.
func (e *Element) NonEmptyAttr(name string, value string) *Element {
	e.Void.NonEmptyAttr(name, value)
	return e
}

// Child appends children of the Element.
func (e *Element) Child(el ...Node) *Element {
	e.children = append(e.children, el...)

	return e
}

// T appends a text as a child of the Element.
func (e *Element) T(txt string) *Element {
	return e.Child(T(txt))
}

// ChildEls appends a list of *Void's as the children of the Element.
func (e *Element) ChildVoids(vs ...*Void) *Element {
	for _, v := range vs {
		e.children = append(e.children, v)
	}
	return e
}

// ChildEls appends a list of *Element's as the children of the Element.
func (e *Element) ChildEls(els ...*Element) *Element {
	for _, el := range els {
		e.children = append(e.children, el)
	}
	return e
}

func shouldNewLine(e *Element) bool {
	switch e.tagType {
	case PRETag, TEXTAREATag:
		if len(e.children) == 0 {
			return false
		}

		t, ok := e.children[0].(HTMLNode)
		if !ok {
			return false
		}
		if len(t) == 0 {
			return false
		}
		return t[0] == '\n'
	}
	return false
}

func (e *Element) WriteTo(b Writer, opt RenderOptions, parent *Element, childIndex int) {
	// TODO omit and indent
	if opt.DisableOmit || !canElementOmitStartTag(e, parent, childIndex) {
		// Write the open tag including attributes
		e.Void.WriteTo(b, opt, parent, childIndex)
	}

	if shouldNewLine(e) {
		b.WriteByte('\n')
	}

	for i, child := range e.children {
		child.WriteTo(b, opt, e, i)
	}

	if !opt.DisableOmit && canElementOmitEndTag(e, parent, childIndex) {
		return
	}

	b.WriteString(`</`)
	b.WriteString(TagNames[e.tagType])
	b.WriteByte('>')
}

// T returns an escaped HTMLNode of the specified text as contents
func T(text string) HTMLNode {
	return HTMLNode(utils.EscapeHTML(text))
}

// T returns an escaped HTMLNode of the specified formated text as contents
func Tf(format string, args ...interface{}) HTMLNode {
	return HTMLNode(utils.EscapeHTML(fmt.Sprintf(format, args...)))
}
