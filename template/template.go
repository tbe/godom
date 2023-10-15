/*
Package template provides a high-level wrapper around the standard html/template package, enabling seamless integration with the godom library.

The template package bridges the gap between godom and standard templates, focusing on improving performance and user experience.
Unlike the upstream html/template package, it currently does not support parsing multiple related templates.

Notably, when using this package, users must provide an access their own data within the `UserData` member of the
Context, limiting customization.

Features:
- Seamless integration with godom elements.
- Provides utility functions to manage placeholders and attributes.
- Offers a mechanism to define fallback content and attributes.
*/
package template

import (
	"fmt"
	htmltemplate "html/template"
	"io"

	"github.com/tbe/godom"
	"github.com/tbe/godom/helpers"
	"github.com/tbe/godom/types"
	"github.com/tbe/godom/util"
)

// Context represents the data structure used during the template execution.
// It defines placeholders, attributes, and user data to customize the template rendering.
type Context struct {
	Placeholders map[string]types.Element
	Attributes   map[string]types.Attribute
	UserData     any
}

// Must is a utility function that accepts a template and an error. If the error is non-nil, it panics; otherwise, it returns the template.
// This is useful for ensuring that template creation operations are successful without requiring explicit error handling.
func Must(t *Template, err error) *Template {
	if err != nil {
		panic(err)
	}
	return t
}

// Template is a high-level structure that wraps the standard html/template.Template.
// It integrates godom elements, allowing them to be treated as HTML templates.
type Template struct {
	html *htmltemplate.Template

	contents   map[string]types.Element
	attributes map[string]types.Attribute
}

// New initializes and returns a new Template with the specified name.
// It sets up the necessary helper functions and data structures for rendering godom elements.
func New(name string) *Template {
	// we create a new HTML template
	html := htmltemplate.New(name)

	t := &Template{
		html:       html,
		contents:   make(map[string]types.Element),
		attributes: make(map[string]types.Attribute),
	}

	// we provide our render helper functions
	html.Funcs(htmltemplate.FuncMap{
		"godoc_content":   t.getContent,
		"godoc_attribute": t.getAttribute,
	})

	return t
}

// HTML provides direct access to the underlying html/template.Template of the wrapper.
func (t *Template) HTML() *htmltemplate.Template {
	return t.html
}

// Parse accepts a godom element and parses it as the template body.
// The element is immediately rendered into the template.
func (t *Template) Parse(element types.Element) (*Template, error) {
	// parse into our HTML template
	htmlStr, err := util.RenderToString(element)
	if err != nil {
		return nil, err
	}
	_, err = t.html.Parse(htmlStr)
	return t, err
}

// Execute wraps the underlying template's Execute method.
// It renders the template with the provided context data and writes the output to the specified writer.
func (t *Template) Execute(wr io.Writer, data *Context) error {
	if data == nil {
		data = &Context{}
	}
	return t.html.Execute(wr, data)
}

// Placeholder defines a placeholder in the template that can later be replaced with actual content.
// The returned godom element represents the placeholder in template syntax. Initially, it's set to an empty element.
func (t *Template) Placeholder(key string) types.Element {
	if _, exists := t.contents[key]; exists {
		panic(fmt.Sprintf("placeholder %q already exists", key))
	}
	// we set our placeholder to an empty element
	t.contents[key] = godom.Group()

	return helpers.NewStringElement(fmt.Sprintf("{{ godoc_content %q .Placeholders }}", key))
}

// Attribute defines an attribute placeholder in the template.
// The returned godom attribute represents the placeholder in template syntax. Initially, it's set to an empty attribute.
func (t *Template) Attribute(key string) types.Attribute {
	if _, exists := t.attributes[key]; exists {
		panic(fmt.Sprintf("attribute %q already exists", key))
	}
	// we set our placeholder to an empty element
	t.attributes[key] = func(_ map[string]string, _ *[]string, _ *[]types.Attribute) {

	}

	return helpers.FlagAttribute(fmt.Sprintf("{{ godoc_attribute %q .Attributes }}", key))
}

// SetFallbackContent specifies the default content for a placeholder.
// If no content is provided during the template execution, this fallback content will be used.
// The provided godom element will be rendered each time the template is executed.
func (t *Template) SetFallbackContent(key string, element types.Element) error {
	if _, exists := t.contents[key]; !exists {
		return fmt.Errorf("key %q does not exist", key)
	}
	t.contents[key] = element
	return nil
}

// SetFallbackAttribute specifies the default attribute for a placeholder.
// If no attribute is provided during the template execution, this fallback attribute will be used.
// The provided godom attribute will be rendered each time the template is executed.
func (t *Template) SetFallbackAttribute(key string, attribute types.Attribute) error {
	if _, exists := t.attributes[key]; !exists {
		return fmt.Errorf("key %q does not exist", key)
	}
	t.attributes[key] = attribute
	return nil
}

// getContent fetches and returns the content associated with a given key.
// It looks for the content in the provided data and, if not found, falls back to the template's default contents.
func (t *Template) getContent(key string, data map[string]types.Element) htmltemplate.HTML {
	var content types.Element
	exists := false

	if data != nil {
		// check if we have content for the element
		content, exists = data[key]
	}
	if !exists {
		content, exists = t.contents[key]
		if !exists {
			panic(fmt.Sprintf("missing content for placeholder %q", key))
		}
	}
	rendered, err := util.RenderToString(content)
	if err != nil {
		panic(fmt.Sprintf("failed to render the contents of %q", key))
	}
	return htmltemplate.HTML(rendered)
}

// getAttribute fetches and returns the attribute associated with a given key.
// It looks for the attribute in the provided data and, if not found, falls back to the template's default attributes.
func (t *Template) getAttribute(key string, data map[string]types.Attribute) htmltemplate.HTMLAttr {
	var attr types.Attribute
	exists := false

	if data != nil {
		// check if we have content for the element
		attr, exists = data[key]
	}
	if !exists {
		attr, exists = t.attributes[key]
		if !exists {
			panic(fmt.Sprintf("missing content for placeholder %q", key))
		}
	}

	// we need to render our attributes into a string.
	attributes := make(map[string]string)
	var flags []string

	// execute the attribute function
	attr(attributes, &flags, nil)

	// render to a string
	var allAttrs []string
	for k, v := range attributes {
		allAttrs = append(allAttrs, fmt.Sprintf("%s=%q", k, v))
	}
	for _, k := range flags {
		allAttrs = append(allAttrs, k)
	}

	// make sure we only have one attribute here
	if len(allAttrs) > 1 {
		panic("a template attribute is not allowed to create multiple attributes")
	}

	if len(allAttrs) == 1 {
		return htmltemplate.HTMLAttr(allAttrs[0])
	}

	return ""
}
