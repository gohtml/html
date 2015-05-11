package html

import (
	"testing"

	"github.com/golangplus/testing/assert"

	. "github.com/gohtml/elements"
)

func TestVoid(t *testing.T) {
	div := &Void{
		tagType: DIVTag,
	}

	div.TabIndex(1024)

	div.AddClass("line")

	assert.StringEqual(t, "div", NodeToHTMLNode(div, DefaultOptions),
		`<div class="line" tabindex="1024">`)

	div.DelClass("line")

	assert.StringEqual(t, "div", NodeToHTMLNode(div, DefaultOptions),
		`<div tabindex="1024">`)
}
