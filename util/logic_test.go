package util_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tbe/godom"
	"github.com/tbe/godom/types"
	"github.com/tbe/godom/util"
)

func TestIfAttr(t *testing.T) {
	// we use a simple flag attribute for testing
	{
		var flags []string
		util.IfAttr(true, godom.Disabled())(nil, &flags, nil)
		assert.Equal(t, []string{"disabled"}, flags)
	}

	{
		var flags []string
		util.IfAttr(false, godom.Disabled())(nil, &flags, nil)
		assert.Empty(t, flags)
	}
}

func TestIfElem(t *testing.T) {
	{
		var buf bytes.Buffer
		godom.Group(util.IfElem(true, godom.P()())).Render(&buf)
		assert.Equal(t, "<p></p>", buf.String())
	}

	{
		var buf bytes.Buffer
		godom.Group(util.IfElem(false, godom.P()())).Render(&buf)
		assert.Empty(t, buf.String())
	}
}

func TestMap(t *testing.T) {
	data := []string{"first", "second", "third"}
	factory := func(str string) types.Element {
		return godom.P()(godom.Content(str))
	}

	var buf bytes.Buffer
	util.Map(data, factory).Render(&buf)
	assert.Equal(t, "<p>first</p><p>second</p><p>third</p>", buf.String())
}
