package util

import (
	"bytes"

	"github.com/tbe/godom/types"
)

// RenderToString renders a given types.Element to its string representation.
// This is useful when the HTML generated by godoc needs to be used in another context or module.
func RenderToString(element types.Element) (string, error) {
	var buf bytes.Buffer
	if err := element.Render(&buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}
