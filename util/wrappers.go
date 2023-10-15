package util

import (
	"io"

	"github.com/tbe/godom/types"
)

// ElementCreator defines a function type that creates and returns a godom.Element.
type ElementCreator func() types.Element

type delayedElementWrapper struct {
	creator ElementCreator
}

// DelayedElement returns a godom.Element that is constructed immediately before rendering.
// This allows for dynamic adaptation of parts of the DOM between its construction and rendering phases.
func DelayedElement(creator ElementCreator) types.Element {
	return &delayedElementWrapper{creator: creator}
}

func (d *delayedElementWrapper) Render(writer io.Writer) error {
	return d.creator().Render(writer)
}

// DelayedAttribute wraps the provided types.Attribute such that it renders every time
// the containing types.Element is rendered. This allows attributes to be dynamically
// determined during the rendering phase.
func DelayedAttribute(attribute types.Attribute) types.Attribute {
	return func(_ map[string]string, _ *[]string, delayed *[]types.Attribute) {
		*delayed = append(*delayed, attribute)
	}
}
