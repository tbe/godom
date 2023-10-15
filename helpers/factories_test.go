package helpers_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tbe/godom/helpers"
)

type HelpersTestSuite struct {
	suite.Suite
}

// TestHelpersTestSuite initializes the test suite.
func TestHelpersTestSuite(t *testing.T) {
	suite.Run(t, new(HelpersTestSuite))
}

func (s *HelpersTestSuite) TestChildlessElement() {
	brElement := helpers.NewChildlessElement("br")
	var buf bytes.Buffer
	assert.NoError(s.T(), brElement.Render(&buf))
	assert.Equal(s.T(), "<br/>", buf.String())
}

func (s *HelpersTestSuite) TestElementWithChildren() {
	Div := helpers.NewElement("div")
	childlessChild := helpers.NewChildlessElement("br")
	divElement := Div(childlessChild)
	var buf bytes.Buffer
	assert.NoError(s.T(), divElement.Render(&buf))
	assert.Equal(s.T(), "<div><br/></div>", buf.String())
}

func (s *HelpersTestSuite) TestStringElement() {
	staticContent := "Hello, World!"
	stringElem := helpers.NewStringElement(staticContent)
	var buf bytes.Buffer
	assert.NoError(s.T(), stringElem.Render(&buf))
	assert.Equal(s.T(), staticContent, buf.String())
}

func (s *HelpersTestSuite) TestAttributes() {
	Div := helpers.NewElement("div", helpers.SingleAttribute("id", "test"))
	divElem := Div()

	var buf bytes.Buffer
	assert.NoError(s.T(), divElem.Render(&buf))
	assert.Equal(s.T(), `<div id="test"></div>`, buf.String())

	buf.Reset()

	// Now, with a multi-value attribute.
	DivWithClass := helpers.NewElement("div", helpers.MultiValueAttribute("class", "btn", "btn-primary"))
	divElemWithClass := DivWithClass()
	assert.NoError(s.T(), divElemWithClass.Render(&buf))
	assert.Equal(s.T(), `<div class="btn btn-primary"></div>`, buf.String())

	buf.Reset()

	// Now, a flag attribute.
	DivWithFlag := helpers.NewElement("div", helpers.FlagAttribute("hidden"))
	divElemWithFlag := DivWithFlag()
	assert.NoError(s.T(), divElemWithFlag.Render(&buf))
	assert.Equal(s.T(), `<div hidden></div>`, buf.String())
}
