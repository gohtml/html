package html

import (
	"strconv"

	"github.com/golangplus/bytes"
)

// HTML Entities
var (
	// HTML Character Entity copyright
	COPY = HTMLNode("&copy;")
	// HTML Character Entity ampersand
	AMP = HTMLNode("&amp;")
	// HTML Character Entity less than
	LT = HTMLNode("&lt;")
	// HTML Character Entity greater than
	GT = HTMLNode("&gt;")
	// HTML Character Entity cent
	CENT = HTMLNode("&cent;")
	// HTML Character Entity pound
	POUND = HTMLNode("&pound;")
	// HTML Character Entity yen
	YEN = HTMLNode("&yen;")
	// HTML Character Entity euro
	EURO = HTMLNode("&euro;")
	// HTML Character Entity registered trademark
	REG = HTMLNode("&reg;")
	// HTML Character Entity non-breaking space
	NBSP  = HTMLNode("&nbsp;")
	TIMES = HTMLNode("&times;")
	LAQUO = HTMLNode("&laquo;")
	RAQUO = HTMLNode("&raquo;")

	//TODO define all entities
)

var nePrefix = "&#"

// NumEnt returns HTMLBytes a numerical entity.
func NumEnt(num int) HTMLNode {
	var b bytesp.ByteSlice

	b.WriteString(nePrefix)
	b = strconv.AppendInt([]byte(b), int64(num), 10)
	b.WriteByte(';')

	return HTMLNode(b)
}
