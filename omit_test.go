package html

import (
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestOmittedTags_li(t *testing.T) {
	h := HTML("")
	body := h.Body()
	body.Child(UL(
		LI(T("Hello")),
		LI(T("World")),
	))

	assert.StringEqual(t, "html", NodeToHTMLNode(h, RenderOptions{SortAttr: true}),
		`<!DOCTYPE html>
<meta charset="utf-8"><ul><li>Hello<li>World</ul>`)
}

func TestOmittedTags_dtdd(t *testing.T) {
	h := HTML("")
	body := h.Body()
	body.Child(DL(
		DT(T("Hello")),
		DT(T("Hello")),
		DD(T("World")),
		DD(T("World")),
		DT(T("Hello")),
		DD(T("World")),
	))
	assert.StringEqual(t, "html", NodeToHTMLNode(h, RenderOptions{SortAttr: true}),
		`<!DOCTYPE html>
<meta charset="utf-8"><dl><dt>Hello<dt>Hello<dd>World<dd>World<dt>Hello<dd>World</dl>`)
}

func TestOmittedTags_p(t *testing.T) {
	h := HTML("")
	body := h.Body()
	body.Child(
		P(T("Hello")),
		DIV(T("Hello")),
		P(T("World")),
	)

	assert.StringEqual(t, "html", NodeToHTMLNode(h, RenderOptions{SortAttr: true}),
		`<!DOCTYPE html>
<meta charset="utf-8"><p>Hello<div>Hello</div><p>World`)
}

func TestOmittedTags_ruby(t *testing.T) {
	h := HTML("")
	body := h.Body()
	body.Child(RUBY(
		T("中文"),
		RB(T("Hello")),
		RT(T("zhongwen")),
		RTC(T("World")),
		RP(T("World")),
	))

	assert.StringEqual(t, "html", NodeToHTMLNode(h, RenderOptions{SortAttr: true}),
		`<!DOCTYPE html>
<meta charset="utf-8"><ruby>中文<rb>Hello<rt>zhongwen<rtc>World<rp>World</ruby>`)
}

func TestOmittedTags_optgroup(t *testing.T) {
	h := HTML("")
	body := h.Body()
	body.Child(SELECT(
		OPTION("W2", "w2"),
		OPTGROUP("hello",
			OPTION("W1", "w1"),
		),
		OPTGROUP("world",
			OPTION("H1", "h1"),
		),
	))
	assert.StringEqual(t, "html", NodeToHTMLNode(h, RenderOptions{SortAttr: true}),
		`<!DOCTYPE html>
<meta charset="utf-8"><select><option value="W2">w2<optgroup label="hello"><option value="W1">w1<optgroup label="world"><option value="H1">h1</select>`)
}

func TestOmittedTags_table(t *testing.T) {
	h := HTML("")
	body := h.Body()
	body.Child(TABLE(
		COLGROUP(0,
			COL(2), COL(1),
		),
		THEAD(
			TR(TH()),
		),
		TBODY(
			TR(TD()),
		),
		TFOOT(),
	))

	assert.StringEqual(t, "html", NodeToHTMLNode(h, RenderOptions{SortAttr: true}),
		`<!DOCTYPE html>
<meta charset="utf-8"><table><col span="2"><col><thead><tr><th><tbody><tr><td><tfoot></table>`)
}
