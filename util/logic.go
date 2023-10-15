package util

import (
	"github.com/tbe/godom"
	"github.com/tbe/godom/types"
)

// IfAttr returns a godom.Attribute that is applied to the godom.Element
// based on the given condition. If the condition is true, the provided attribute
// is applied; otherwise, a no-op attribute is applied.
func IfAttr(condition bool, attribute types.Attribute) types.Attribute {
	if condition {
		return attribute
	} else {
		return func(_ map[string]string, _ *[]string, _ *[]types.Attribute) {

		}
	}
}

// IfElem returns a godom.Element based on the given condition. If the condition
// is true, the provided element is returned; otherwise, an empty element is returned.
func IfElem(condition bool, element types.Element) types.Element {
	if condition {
		return element
	} else {
		return godom.Group()
	}
}

// Map takes a slice of data and a creator function that transforms each data item into
// a godom.Element. It returns a group of elements created from the data slice.
func Map[T any](data []T, creator func(T) types.Element) types.Element {
	var elems []types.Element
	for _, e := range data {
		elems = append(elems, creator(e))
	}
	return godom.Group(elems...)
}
