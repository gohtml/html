package html

import (
	"github.com/golangplus/sort"
)

type attrInfo struct {
	name  HTMLNode
	value HTMLNode
}

type Attributes []attrInfo

func (attrs Attributes) WriteTo(b Writer, sortAttr bool) {
	if len(attrs) == 0 {
		return
	}

	if sortAttr && len(attrs) > 1 {
		sortp.SortF(len(attrs), func(i, j int) bool {
			return attrs[i].name < attrs[j].name
		}, func(i, j int) {
			attrs[i], attrs[j] = attrs[j], attrs[i]
		})
	}

	for i := range attrs {
		b.WriteByte(' ')
		attrs[i].name.WriteRaw(b)
		if len(attrs[i].value) > 0 {
			b.WriteString(`="`)
			attrs[i].value.WriteRaw(b)
			b.WriteByte('"')
		}
	}
}

func (attrs Attributes) index(name HTMLNode) int {
	for i := range attrs {
		if attrs[i].name == name {
			return i
		}
	}
	return -1
}

func (attrs *Attributes) Put(name, value HTMLNode) {
	i := attrs.index(name)
	if i >= 0 {
		(*attrs)[i].value = value
		return
	}

	*attrs = append(*attrs, attrInfo{name, value})
}
