package html

import (
	"testing"

	"github.com/golangplus/bytes"
	"github.com/golangplus/testing/assert"
)

func TestAttributes(t *testing.T) {
	var attrs Attributes

	attrs.Put("abc", "")

	var b bytesp.ByteSlice

	attrs.WriteTo(&b, false)
	assert.Equal(t, "attrs", string(b), ` abc`)

	b = nil
	attrs.WriteTo(&b, true)
	assert.Equal(t, "attrs", string(b), ` abc`)
}
