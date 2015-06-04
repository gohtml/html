package html

import (
	"fmt"
	"testing"

	"github.com/golangplus/testing/assert"

	. "github.com/gohtml/url"
)

func ExampleHtml_NoOmit() {
	h := HTML("en")
	fmt.Println(NodeToHTMLNode(h, RenderOptions{DisableOmit: true, SortAttr: true}))
	// OUTPUT:
	// <!DOCTYPE html>
	// <html lang="en"><head><meta charset="utf-8"></head><body></body></html>
}

func ExampleHtml_Simple() {
	h := HTML("en")
	h.Title("Title of Page")
	h.Favicon("favicon.png", "image/png")
	h.Css(U("", "main.css", ""))

	h.Body().T(`Hello, "world"`)
	fmt.Println(NodeToHTMLNode(h, RenderOptions{SortAttr: true}))
	// OUTPUT:
	// <!DOCTYPE html>
	// <html lang="en"><meta charset="utf-8"><title>Title of Page</title><link href="favicon.png" rel="shortcut icon" type="image/png"><link href="main.css" rel="stylesheet" type="text/css">Hello, &quot;world&quot;
}

func TestHtml_Child(t *testing.T) {
	h := HTML("")
	assert.Panic(t, "h.Child", h.Child)
}
