package types

import (
	"io"
)

// Element represents the basic interface that every HTML element must implement.
type Element interface {
	// Render writes the HTML representation of the element to the provided writer.
	Render(writer io.Writer) error
}

// ElementFactory is a function type that produces an Element with an optional list of child Elements.
type ElementFactory func(children ...Element) Element

// ElementConstructor is a function type that produces an ElementFactory for an Element.
// This constructor allows the specification of attributes for the resulting Element.
type ElementConstructor func(attributes ...Attribute) ElementFactory

// ChildlessElementConstructor is a function type that produces an Element that doesn't wrap any children.
// This constructor allows the specification of attributes for the resulting Element.
type ChildlessElementConstructor func(attributes ...Attribute) Element

// Attribute represents a function type that can inspect and manipulate the attributes and flags of an Element.
// This allows for dynamic behavior and customization of the Element's properties.
// - attrs is the current set of attributes for the element.
// - flags is a pointer to a list of flags for the element. Flags are attributes without values.
// - delayed is a pointer to a list of attributes that are evaluated during the rendering phase. (this is nil during the rendering phase!)
type Attribute func(attrs map[string]string, flags *[]string, delayed *[]Attribute)
