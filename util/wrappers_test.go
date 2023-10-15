package util_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tbe/godom"
	"github.com/tbe/godom/types"
	"github.com/tbe/godom/util"
)

func TestDelayedElement(t *testing.T) {
	// we take a simple string for our dynamic data
	testClass := "testA"
	creator := func() types.Element {
		return godom.P(godom.Class(testClass))()
	}

	// here is our document
	doc := godom.Div()(
		util.DelayedElement(creator),
	)

	// now, we render this two times, and see that everything matches
	{
		var buf bytes.Buffer
		assert.NoError(t, doc.Render(&buf))

		assert.Equal(t, `<div><p class="testA"></p></div>`, buf.String())
	}
	// set the testClass
	testClass = "testB"
	{
		var buf bytes.Buffer
		assert.NoError(t, doc.Render(&buf))

		assert.Equal(t, `<div><p class="testB"></p></div>`, buf.String())
	}
}

func TestDelayedAttribute(t *testing.T) {
	// we take a simple string for our dynamic data
	testClass := "testA"
	// and here is our very simple attribute
	attr := func(attrs map[string]string, flags *[]string, _ *[]types.Attribute) {
		// we ignore any testing here
		attrs["class"] = testClass
	}

	// here is our document
	doc := godom.Div(util.DelayedAttribute(attr))()

	// now, we render this two times, and see that everything matches
	{
		var buf bytes.Buffer
		assert.NoError(t, doc.Render(&buf))

		assert.Equal(t, `<div class="testA"></div>`, buf.String())
	}
	// set the testClass
	testClass = "testB"
	{
		var buf bytes.Buffer
		assert.NoError(t, doc.Render(&buf))

		assert.Equal(t, `<div class="testB"></div>`, buf.String())
	}
}
