package template_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	. "github.com/tbe/godom"
	"github.com/tbe/godom/helpers"
	"github.com/tbe/godom/template"
	"github.com/tbe/godom/types"
)

type TemplateTestSuite struct {
	suite.Suite

	buf  bytes.Buffer
	tmpl *template.Template
}

// TestTemplateTestSuite initializes the test suite.
func TestTemplateTestSuite(t *testing.T) {
	suite.Run(t, new(TemplateTestSuite))
}

func (s *TemplateTestSuite) SetupTest() {
	s.buf.Reset()
	s.tmpl = template.New("root")
}

func (s *TemplateTestSuite) TestTemplateElement() {
	root := Div()(
		P()(Content("This is a counter: "), s.tmpl.Placeholder("counter")),
	)
	template.Must(s.tmpl.Parse(root))

	// Check rendering without the counter value
	s.NoError(s.tmpl.Execute(&s.buf, nil))
	s.Equal("<div><p>This is a counter: </p></div>", s.buf.String())

	// Check rendering with a static text content
	counterText := helpers.NewStringElement("42")
	s.buf.Reset()
	assert.NoError(s.T(), s.tmpl.Execute(&s.buf, &template.Context{
		Placeholders: map[string]types.Element{"counter": counterText},
	}))
	s.Equal("<div><p>This is a counter: 42</p></div>", s.buf.String())
}

func (s *TemplateTestSuite) TestTemplateAttribute() {
	root := Div(s.tmpl.Attribute("class"))()
	template.Must(s.tmpl.Parse(root))

	// Check rendering without the class attribute
	s.NoError(s.tmpl.Execute(&s.buf, nil))
	s.Equal("<div ></div>", s.buf.String())

	// Check rendering with a static class attribute
	testclass := Class("test")
	s.buf.Reset()
	s.NoError(s.tmpl.Execute(&s.buf, &template.Context{
		Attributes: map[string]types.Attribute{"class": testclass},
	}))
	assert.Equal(s.T(), "<div class=\"test\"></div>", s.buf.String())
}

func (s *TemplateTestSuite) TestMultiplePlaceholders() {
	root := Div()(
		P()(s.tmpl.Placeholder("first"), s.tmpl.Placeholder("second")),
	)
	template.Must(s.tmpl.Parse(root))

	// Check rendering with both placeholders
	firstText := helpers.NewStringElement("Hello")
	secondText := helpers.NewStringElement("World")
	s.buf.Reset()
	s.NoError(s.tmpl.Execute(&s.buf, &template.Context{
		Placeholders: map[string]types.Element{
			"first":  firstText,
			"second": secondText,
		},
	}))
	s.Equal("<div><p>HelloWorld</p></div>", s.buf.String())
}

func (s *TemplateTestSuite) TestMultipleAttributes() {
	root := Div(s.tmpl.Attribute("class"), s.tmpl.Attribute("id"))()
	template.Must(s.tmpl.Parse(root))

	// Check rendering with both attributes
	testclass := Class("test")
	testid := ID("unique")
	s.buf.Reset()
	s.NoError(s.tmpl.Execute(&s.buf, &template.Context{
		Attributes: map[string]types.Attribute{
			"class": testclass,
			"id":    testid,
		},
	}))
	assert.Equal(s.T(), `<div class="test" id="unique"></div>`, s.buf.String())
}

func (s *TemplateTestSuite) TestFallbackContentAndAttributes() {
	root := Div()(
		P()(s.tmpl.Placeholder("placeholder")),
	)
	template.Must(s.tmpl.Parse(root))
	fallback := helpers.NewStringElement("Fallback Content")
	s.NoError(s.tmpl.SetFallbackContent("placeholder", fallback))

	// Check rendering with fallback content
	s.buf.Reset()
	s.NoError(s.tmpl.Execute(&s.buf, nil))
	s.Equal("<div><p>Fallback Content</p></div>", s.buf.String())
}

func (s *TemplateTestSuite) TestDuplicatePlaceholdersPanics() {
	s.Panics(func() {
		Div()(
			s.tmpl.Placeholder("duplicate"), s.tmpl.Placeholder("duplicate"),
		)
	}, "Expected panic for duplicate placeholders")

}

func (s *TemplateTestSuite) TestDuplicateAttributesPanics() {
	s.Panics(func() { Div(s.tmpl.Attribute("duplicate"), s.tmpl.Attribute("duplicate"))() }, "Expected panic for duplicate attributes")
}
