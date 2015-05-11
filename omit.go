package html

import (
	. "github.com/gohtml/elements"
	"github.com/gohtml/utils"
)

func canElementOmitStartTag(e, parent *Element, childIndex int) bool {
	if len(e.attributes) > 0 || len(e.classes) > 0 {
		return false
	}

	switch e.tagType {
	case HTMLTag, HEADTag:
		return true

	case BODYTag:
		if len(e.children) == 0 {
			return true
		}
		switch e.children[0].Type() {
		case TextType:
			return !utils.StartWithSpace(string(e.children[0].(HTMLNode)))

		case METATag, LINKTag, SCRIPTTag, TEMPLATETag:
			return false
		}
		return true

	case COLGROUPTag:
		if len(e.children) == 0 {
			return false
		}

		if e.children[0].Type() != COLTag {
			return false
		}

		if childIndex > 0 && parent != nil && parent.children[childIndex-1].Type() == COLGROUPTag {
			return false
		}

		return true

	case TBODYTag:
		if len(e.children) == 0 {
			return false
		}

		if e.children[0].Type() != TRTag {
			return false
		}

		if childIndex > 0 && parent != nil {
			switch parent.children[childIndex-1].Type() {
			case TBODYTag, THEADTag, TFOOTTag:
				return false
			}
		}

		return true
	}

	return false
}

var pOmittedAfter = []bool{
	ADDRESSTag:    true,
	ARTICLETag:    true,
	ASIDETag:      true,
	BLOCKQUOTETag: true,
	DIVTag:        true,
	DLTag:         true,
	FIELDSETTag:   true,
	FOOTERTag:     true,
	FORMTag:       true,
	H1Tag:         true,
	H2Tag:         true,
	H3Tag:         true,
	H4Tag:         true,
	H5Tag:         true,
	H6Tag:         true,
	HEADERTag:     true,
	HGROUPTag:     true,
	HRTag:         true,
	MAINTag:       true,
	NAVTag:        true,
	OLTag:         true,
	PTag:          true,
	PRETag:        true,
	SECTIONTag:    true,
	TABLETag:      true,
	ULTag:         true,
}

func canElementOmitEndTag(e, parent *Element, childIndex int) bool {
	switch tp := e.Type(); tp {
	case HTMLTag, HEADTag, BODYTag:
		return true

	case LITag:
		if parent == nil {
			return false
		}
		if childIndex == len(parent.children)-1 {
			return true
		}
		if parent.children[childIndex+1].Type() == LITag {
			return true
		}

	case DTTag, DDTag:
		if childIndex == len(parent.children)-1 {
			return tp == DDTag
		}
		switch parent.children[childIndex+1].Type() {
		case DTTag, DDTag:
			return true
		}

	case PTag:
		if childIndex == len(parent.children)-1 {
			return parent.Type() != ATag
		}

		nextTp := parent.children[childIndex+1].Type()
		if nextTp < 0 || int(nextTp) >= len(pOmittedAfter) {
			return false
		}
		return pOmittedAfter[nextTp]

	case RBTag, RTTag, RTCTag, RPTag:
		if childIndex == len(parent.children)-1 {
			return true
		}

		switch parent.children[childIndex+1].Type() {
		case RBTag, RTCTag, RPTag:
			return true
		case RTTag:
			return tp != RTCTag
		}

	case OPTGROUPTag, OPTIONTag:
		if childIndex == len(parent.children)-1 {
			return true
		}

		switch parent.children[childIndex+1].Type() {
		case OPTGROUPTag:
			return true
		case OPTIONTag:
			return tp == OPTIONTag
		}

	case COLGROUPTag:
		if childIndex == len(parent.children)-1 {
			return true
		}

		if parent.children[childIndex+1].Type() == TextType {
			return !utils.StartWithSpace(string(e.children[childIndex+1].(HTMLNode)))
		}

		return true

	case THEADTag:
		if childIndex == len(parent.children)-1 {
			return false
		}

		switch parent.children[childIndex+1].Type() {
		case TBODYTag, TFOOTTag:
			return true
		}

	case TBODYTag, TFOOTTag:
		if childIndex == len(parent.children)-1 {
			return true
		}

		switch parent.children[childIndex+1].Type() {
		case TBODYTag:
			return true

		case TFOOTTag:
			return tp == TBODYTag
		}

	case TRTag:
		if childIndex == len(parent.children)-1 {
			return true
		}

		return parent.children[childIndex+1].Type() == TRTag

	case THTag, TDTag:
		if childIndex == len(parent.children)-1 {
			return true
		}

		switch parent.children[childIndex+1].Type() {
		case THTag, TDTag:
			return true
		}
	}

	return false
}
