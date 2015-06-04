package html

import (
	"fmt"
	"testing"

	"github.com/golangplus/testing/assert"

	. "github.com/gohtml/elements"
	. "github.com/gohtml/url"
)

func TestAllTagsOneLine(t *testing.T) {
	// Slice of element, expected string results.
	elementExpected := []interface{}{
		`HTML`, HTML("en"), `<!DOCTYPE html>
<html lang="en"><meta charset="utf-8">`,
		`HEAD`, HEAD(), `<meta charset="utf-8">`,
		`TITLE`, TITLE(`The <title>`), `<title>The &lt;title&gt;</title>`,
		`BASE`, BASE(U("", "/sub", ""), "_blank"), `<base href="/sub" target="_blank">`,
		`LINK`, LINK(U("", "/main.css", ""), "stylesheet"), `<link href="/main.css" rel="stylesheet">`,
	}

	assert.ValueShould(t, "len(elementExpected)", len(elementExpected), len(elementExpected)%3 == 0,
		"is not a multiple of 3")
	for i := 2; i < len(elementExpected); i += 3 {
		el := elementExpected[i-1].(Node)
		assert.Equal(t, elementExpected[i-2].(string), string(NodeToHTMLNode(el, DefaultOptions)),
			elementExpected[i])
	}
}

func ExampleMAP() {
	h := HTML("")
	body := h.Body()
	body.Child(MAP("menu", AREA("fun.html", "Fun", "circle", []int{1, 2, 3, 4})))
	fmt.Println(NodeToHTMLNode(h, RenderOptions{SortAttr: true}))

	// OUTPUT:
	// <!DOCTYPE html>
	// <meta charset="utf-8"><map name="menu"><area alt="Fun" coords="12,3,4," href="fun.html" shape="circle"></map>
}

func ExampleBASE() {
	h := HTML("")
	head := h.Head()
	head.Child(BASE("/images", "_blank"))
	fmt.Println(NodeToHTMLNode(h, RenderOptions{SortAttr: true}))

	// OUTPUT:
	// <!DOCTYPE html>
	// <meta charset="utf-8"><base href="/images" target="_blank">
}

func ExampleTextBR() {
	h := HTML("")
	body := h.Body()
	body.Child(T("ABC"), BR(), T("Def"))
	fmt.Println(NodeToHTMLNode(h, RenderOptions{SortAttr: true}))

	// OUTPUT:
	// <!DOCTYPE html>
	// <meta charset="utf-8">ABC<br>Def
}

func ExampleTABLE() {
	h := HTML("")
	body := h.Body()
	body.Child(TABLE(
		CAPTION(T("Hello")),
		COLGROUP(0,
			COL(2), COL(1),
		),
		COLGROUP(5),
		THEAD(
			TR(TH()),
		),
		TBODY(
			TR(TD()),
		),
		TFOOT(),
	))
	fmt.Println(NodeToHTMLNode(h, RenderOptions{SortAttr: true}))

	// OUTPUT:
	// <!DOCTYPE html>
	// <meta charset="utf-8"><table><caption>Hello</caption><col span="2"><col><colgroup span="5"><thead><tr><th><tbody><tr><td><tfoot></table>
}

func ExampleEMBED() {
	h := HTML("")
	h.Body().Child(
		EMBED(URL("http://example.com/a.swf"), "flash", 100, 200),
		EMBED(URL("http://example.com/a.swf"), "", -1, -1),
	)
	fmt.Println(NodeToHTMLNode(h, RenderOptions{SortAttr: true}))

	// OUTPUT:
	// <!DOCTYPE html>
	// <meta charset="utf-8"><embed height="200" src="http://example.com/a.swf" type="flash" width="100"><embed src="http://example.com/a.swf">
}

func ExampleHR() {
	h := HTML("")
	h.Body().Child(HR())
	fmt.Println(NodeToHTMLNode(h, RenderOptions{SortAttr: true}))

	// OUTPUT:
	// <!DOCTYPE html>
	// <meta charset="utf-8"><hr>
}
