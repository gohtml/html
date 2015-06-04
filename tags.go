package html

import (
	"strconv"

	. "github.com/gohtml/elements"
	"github.com/gohtml/utils"
)

/*
The root element.
http://www.w3.org/TR/html5/semantics.html#the-root-element
*/

// HTML creates an HTML element with type Html.
// http://www.w3.org/TR/html5/semantics.html#the-html-element
func HTML(lang string) *Html {
	return (&Html{
		Element: Element{
			Void: Void{tagType: HTMLTag},
			children: []Node{
				HEAD(),
				BODY(),
			},
		},
	}).Lang(lang)
}

/*
Document metadata.
http://www.w3.org/TR/html5/document-metadata.html#document-metadata
*/

// HEAD creates an HEAD element.
// http://www.w3.org/TR/html5/document-metadata.html#the-head-element
func HEAD() *Element {
	return &Element{
		Void: Void{tagType: HEADTag},
		children: []Node{
			META().Attr("charset", "utf-8"),
		},
	}
}

// TITLE creats a TITLE element.
// http://www.w3.org/TR/html5/document-metadata.html#the-title-element
// TODO maybe this is not necessary, force user to use Html.Title since there
//      should not be more than one title
func TITLE(title string) *Element {
	return &Element{
		Void: Void{tagType: TITLETag},
		children: []Node{
			T(title),
		},
	}
}

// BASE creates a base element.
// http://www.w3.org/TR/html5/document-metadata.html#the-base-element
// TODO maybe this is not necessary, force user to use Html.Base since there
//      should not be more than one base
func BASE(href URL, target string) *Void {
	return (&Void{
		tagType: BASETag,
	}).AttrIfNotEmpty("href", string(href)).AttrIfNotEmpty("target", target)
}

// LINK creates a link element.
// http://www.w3.org/TR/html5/document-metadata.html#the-link-element
func LINK(href URL, rel string) *Void {
	return (&Void{
		tagType: LINKTag,
	}).Attr("href", string(href)).Attr("rel", rel)
}

// The area element
func AREA(href, alt string, shape string, coords []int) *Void {
	area := (&Void{
		tagType: AREATag,
	}).AttrIfNotEmpty("href", href).AttrIfNotEmpty("alt", alt).AttrIfNotEmpty("shape", shape)
	if len(coords) > 0 {
		area.attrOfEscaped("coords", HTMLNode(utils.IntSliceToBytes(coords)))
	}

	return area
}

// The br element
func BR() *Void {
	return &Void{
		tagType: BRTag,
	}
}

// The col element.
// span is the number of columns spanned by the element.
func COL(span int) *Void {
	col := &Void{
		tagType: COLTag,
	}
	if span > 1 {
		col.attrOfEscaped("span", HTMLNode(strconv.Itoa(span)))
	}
	return col
}

// The embed element.
//
// tp is the type attribute. Empty string will be ignored.
// width or height with a negative value will be ignored.
func EMBED(src URL, tp string, width, height int) *Void {
	v := (&Void{
		tagType: EMBEDTag,
	}).AttrIfNotEmpty("type", tp)

	v.attrOfEscaped("src", HTMLNode(src))

	if width >= 0 {
		v.attrOfEscaped("width", HTMLNode(strconv.Itoa(width)))
	}
	if height >= 0 {
		v.attrOfEscaped("height", HTMLNode(strconv.Itoa(height)))
	}

	return v
}

// The br element
func HR() *Void {
	return &Void{
		tagType: HRTag,
	}
}

func IMG(src URL, alt string) *Void {
	return (&Void{
		tagType: IMGTag,
	}).attrOfEscaped("src", HTMLNode(src)).AttrIfNotEmpty("alt", alt)
}

func INPUT(tp, name, value string) *Void {
	v := &Void{
		tagType: INPUTTag,
	}

	if tp != "" {
		v.Attr("type", tp)
	}

	if name != "" {
		v.Attr("name", name)
	}

	if value != "" {
		v.Attr("value", value)
	}

	return v
}

func META() *Void {
	return &Void{
		tagType: METATag,
	}
}

func PARAM(name, value string) *Void {
	return (&Void{
		tagType: PARAMTag,
	}).Attr("name", name).Attr("value", value)
}

/* Normal elements */

func A(href string, children ...Node) *Element {
	t := (&Element{
		Void: Void{tagType: ATag},
	}).Child(children...)

	if href != "" {
		t.Attr("href", href)
	}

	return t
}

func B(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: BTag},
	}).Child(children...)
}

func BODY() *Element {
	return &Element{
		Void: Void{tagType: BODYTag},
	}
}

func BUTTON(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: BUTTONTag},
	}).Child(children...)
}

// The caption element
func CAPTION(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: CAPTIONTag},
	}).Child(children...)
}

// The colgroup element
//
// If span > 0, cols are ignored. Otherwise, cols (col tags) are appended as children.
func COLGROUP(span int, cols ...*Void) *Element {
	colgroup := &Element{
		Void: Void{tagType: COLGROUPTag},
	}
	if span > 0 {
		colgroup.attrOfEscaped("span", HTMLNode(strconv.Itoa(span)))
	} else {
		colgroup.ChildVoids(cols...)
	}
	return colgroup
}

func DD(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: DDTag},
	}).Child(children...)
}

func DIV(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: DIVTag},
	}).Child(children...)
}

func DL(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: DLTag},
	}).Child(children...)
}

func DT(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: DTTag},
	}).Child(children...)
}

func FOOTER(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: FOOTERTag},
	}).Child(children...)
}

func FORM(method, action string, children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: FORMTag},
	}).Attr("method", method).Attr("action", action).Child(children...)
}

func H1(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: H1Tag},
	}).Child(children...)
}

func H2(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: H2Tag},
	}).Child(children...)
}

func H3(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: H3Tag},
	}).Child(children...)
}

func H4(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: H4Tag},
	}).Child(children...)
}

func H5(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: H5Tag},
	}).Child(children...)
}

func H6(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: H6Tag},
	}).Child(children...)
}

func LABEL(For string, children ...Node) *Element {
	t := (&Element{
		Void: Void{tagType: LABELTag},
	}).Child(children...)
	if For != "" {
		t.Attr("for", For)
	}
	return t
}

func LI(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: LITag},
	}).Child(children...)
}

// The map element.
func MAP(name string, children ...Node) *Element {
	mp := (&Element{
		Void: Void{tagType: MAPTag},
	}).Child(children...)
	mp.AttrIfNotEmpty("name", name)
	return mp
}

func NAV(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: NAVTag},
	}).Child(children...)
}

func NOSCRIPT(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: NOSCRIPTTag},
	}).Child(children...)
}

func OBJECT(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: OBJECTTag},
	}).Child(children...)
}

func OL(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: OLTag},
	}).Child(children...)
}

func OPTGROUP(label string, children ...*Element) *Element {
	return (&Element{
		Void: Void{tagType: OPTGROUPTag},
	}).Attr("label", label).ChildEls(children...)
}

func OPTION(value, text string) *Element {
	return (&Element{
		Void: Void{tagType: OPTIONTag},
	}).Attr("value", value).Child(T(text))
}

func P(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: PTag},
	}).Child(children...)
}

func PRE(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: PRETag},
	}).Child(children...)
}

func RB(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: RBTag},
	}).Child(children...)
}

func RP(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: RPTag},
	}).Child(children...)
}

func RT(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: RTTag},
	}).Child(children...)
}

func RTC(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: RTCTag},
	}).Child(children...)
}

func RUBY(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: RUBYTag},
	}).Child(children...)
}

func SCRIPT(src URL, content string) *Element {
	t := &Element{Void: Void{tagType: SCRIPTTag}}
	t.NonEmptyAttr("src", string(src))

	if content != "" {
		t.Child(HTMLNode(content))
	}

	return t
}

func SELECT(children ...*Element) *Element {
	return (&Element{
		Void: Void{tagType: SELECTTag},
	}).ChildEls(children...)
}

func SMALL(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: SMALLTag},
	}).Child(children...)
}

func SPAN(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: SPANTag},
	}).Child(children...)
}

// The table element
func TABLE(children ...*Element) *Element {
	return (&Element{
		Void: Void{tagType: TABLETag},
	}).ChildEls(children...)
}

// The tbody element.
func TBODY(trs ...*Element) *Element {
	return (&Element{
		Void: Void{tagType: TBODYTag},
	}).ChildEls(trs...)
}

// The td element.
func TD(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: TDTag},
	}).Child(children...)
}

func TEXTAREA(name, value string, children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: TEXTAREATag},
	}).Attr("name", name).Attr("value", value).Child(children...)
}

// The tfoot element.
func TFOOT(trs ...*Element) *Element {
	return (&Element{
		Void: Void{tagType: TFOOTTag},
	}).ChildEls(trs...)
}

// The th element.
func TH(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: THTag},
	}).Child(children...)
}

// The thead element.
func THEAD(trs ...*Element) *Element {
	return (&Element{
		Void: Void{tagType: THEADTag},
	}).ChildEls(trs...)
}

func TR(children ...*Element) *Element {
	return (&Element{
		Void: Void{tagType: TRTag},
	}).ChildEls(children...)
}

func UL(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: ULTag},
	}).Child(children...)
}
