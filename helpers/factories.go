// Package helpers provides utility functions and types to simplify and enhance
// the creation and manipulation of godom elements and attributes.
package helpers

import (
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/tbe/godom/types"
	"golang.org/x/exp/maps"
)

// childlessElement represents an HTML element that cannot have child elements.
// Examples include tags like <br> and <img>.
type childlessElement struct {
	tag               string
	attributes        map[string]string
	flags             []string
	delayedAttributes []types.Attribute
}

// NewChildlessElement creates a new types.Element that cannot have child elements.
// It initializes the element with the provided tag and any additional attributes.
// Example usage: NewChildlessElement("br", someAttribute)
func NewChildlessElement(tag string, attrs ...types.Attribute) types.Element {
	e := &childlessElement{
		tag:        tag,
		attributes: make(map[string]string),
	}
	for _, attr := range attrs {
		attr(e.attributes, &e.flags, &e.delayedAttributes)
	}

	return e
}

// renderAttributes writes the attributes of the element to the provided writer.
// It ensures that attributes are ordered consistently for predictable output.
func (ce *childlessElement) renderAttributes(writer io.Writer) error {
	var attrs map[string]string
	var flags []string

	if len(ce.delayedAttributes) > 0 {
		// we make a deep copy of our attributes and flags
		copy(flags, ce.flags)

		attrs = make(map[string]string)
		maps.Copy(attrs, ce.attributes)

		// and then apply all delayed attributes
		for _, delayed := range ce.delayedAttributes {
			delayed(attrs, &flags, nil)
		}
	} else {
		// no copy needed
		flags = ce.flags
		attrs = ce.attributes
	}

	var allAttrs []string
	for k, v := range attrs {
		allAttrs = append(allAttrs, fmt.Sprintf("%s=%q", k, v))
	}
	for _, k := range flags {
		allAttrs = append(allAttrs, k)
	}

	// we sort for a stable result
	slices.Sort(allAttrs)

	if len(allAttrs) > 0 {
		_, err := writer.Write([]byte(" " + strings.Join(allAttrs, " ")))
		return err
	}
	return nil
}

// Render writes the complete representation of the element to the provided writer.
func (ce *childlessElement) Render(writer io.Writer) error {
	if _, err := writer.Write([]byte(fmt.Sprintf("<%s", ce.tag))); err != nil {
		return err
	}
	if err := ce.renderAttributes(writer); err != nil {
		return err
	}

	_, err := writer.Write([]byte("/>"))
	return err
}

// element represents a standard HTML element that can have zero or more child elements.
type element struct {
	childlessElement
	children []types.Element
}

// NewElement creates a factory function for a new types.Element with the given tag and attributes.
// The factory function, when called, will produce an element with the specified child elements.
// Example usage: Div := NewElement("div"); divElement := Div(child1, child2)
func NewElement(tag string, attrs ...types.Attribute) types.ElementFactory {
	e := &element{
		childlessElement: childlessElement{
			tag:        tag,
			attributes: make(map[string]string),
		},
		children: nil,
	}
	for _, attr := range attrs {
		attr(e.attributes, &e.flags, &e.delayedAttributes)
	}
	return func(children ...types.Element) types.Element {
		e.children = children
		return e
	}
}

// Render writes the complete representation of the element, including its children, to the provided writer.
func (e *element) Render(writer io.Writer) error {
	if _, err := writer.Write([]byte(fmt.Sprintf("<%s", e.tag))); err != nil {
		return err
	}
	if err := e.renderAttributes(writer); err != nil {
		return err
	}
	if _, err := writer.Write([]byte(">")); err != nil {
		return err
	}
	for _, child := range e.children {
		if err := child.Render(writer); err != nil {
			return err
		}
	}
	_, err := writer.Write([]byte(fmt.Sprintf("</%s>", e.tag)))
	return err
}

// stringElement represents an element that only holds a static string.
type stringElement struct {
	data []byte
}

// Render writes the static string content of the element to the provided writer.
func (d *stringElement) Render(writer io.Writer) error {
	_, err := writer.Write(d.data)
	return err
}

// NewStringElement creates a new element that holds a static string as its content.
func NewStringElement(content string) types.Element {
	return &stringElement{data: []byte(content)}
}

// MultiValueAttribute creates an attribute with a key that can hold multiple space-separated values.
// If the attribute already exists, the new values are appended.
// Example usage for a class attribute: MultiValueAttribute("class", "btn", "btn-primary")
func MultiValueAttribute(key string, values ...string) types.Attribute {
	return func(attrs map[string]string, _ *[]string, _ *[]types.Attribute) {
		if cValue, exists := attrs[key]; exists {
			attrs[key] = strings.Join(append([]string{cValue}, values...), " ")
		} else {
			attrs[key] = strings.Join(values, " ")
		}
	}
}

// SingleAttribute creates an attribute with a single key-value pair.
// If the attribute with the same key already exists with a different value, it panics.
func SingleAttribute(key, value string) types.Attribute {
	return func(attrs map[string]string, _ *[]string, _ *[]types.Attribute) {
		if cValue, exists := attrs[key]; exists {
			if cValue == value {
				return
			}
			panic(fmt.Sprintf("Attribute %q already exists with value %q", key, cValue))
		}
		attrs[key] = value
	}
}

// FlagAttribute creates an attribute that acts as a flag (i.e., key without a value).
// Example usage for boolean attributes: FlagAttribute("disabled")
func FlagAttribute(flag string) types.Attribute {
	return func(_ map[string]string, flags *[]string, _ *[]types.Attribute) {
		*flags = append(*flags, flag)
	}
}
