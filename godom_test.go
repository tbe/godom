package godom_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/tbe/godom"
	"github.com/tbe/godom/types"
	"github.com/tbe/godom/util"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
)

func ExampleElement_Render() {
	isFeatured := true // this can be dynamically set based on data
	showAuthorBio := true

	// Conditional class for featured post
	featuredClass := util.IfAttr(isFeatured, Class("featured"))

	// Conditional author bio
	authorBio := util.IfElem(showAuthorBio, Div(Class("author-bio"))(
		P()(Content("This is the author's bio. It provides information about the author.")),
	))

	// Delayed attribute example: A function that decides the attribute based on some condition
	dynamicAttribute := util.DelayedAttribute(func(attrs map[string]string, _ *[]string, _ *[]types.Attribute) {
		if isFeatured {
			attrs["data-highlight"] = "true"
		}
	})

	// Delayed element: Maybe we decide to render a special note for featured articles
	featuredNote := util.DelayedElement(func() types.Element {
		if isFeatured {
			return P(Class("featured-note"))(Content("This is a featured article!"))
		}
		return util.IfElem(false, nil) // returns an empty group if not featured
	})

	// Construct the full document
	doc := Group(
		Doctype(),
		HTML()(
			Header()(
				Meta(Charset("utf-8")),
				Title()(Content("Blog Post Title")),
			),
			Body()(
				Div(Class("blog-post"), featuredClass, dynamicAttribute)(
					H1()(Content("Blog Post Title")),
					P()(Content("This is the introduction of the blog post.")),
					featuredNote,
					authorBio,
				),
			),
		),
	)

	var buf bytes.Buffer
	doc.Render(&buf)

	fmt.Println(buf.String())
}

func TestDocRender(t *testing.T) {
	var buf bytes.Buffer

	// we create a full example document and render it
	doc := Group(
		Doctype(),
		HTML()(
			Header()(
				Meta(Charset("utf-8")),
				Title()(Content("Testdocument")),
			),
			Body(Class("someclass"))(
				Div(Class("someotherclass"), ID("mydiv"))(
					H1()(Content("Some Heading here")),
					Form()(
						Button(Type("abort"), Disabled())(Content("My Button")),
					),
				),
			),
		),
	)

	assert.NoError(t, doc.Render(&buf))

	// to keep this readable, we minify everything
	expected := `
<!DOCTYPE html>
<html>
<header>
	<meta charset="utf-8">
	<title>Testdocument</title>
</header>
<body class="someclass">
	<div class="someotherclass" id="mydiv">
		<h1>Some Heading here</h1>
		<form>
			<button disabled type="abort">My Button</button>
		</form>
	</div>
</body>
</html>`

	m := minify.New()
	m.AddFunc("text/html", html.Minify)

	minifedExpected, err := m.String("text/html", expected)
	if err != nil {
		panic(err)
	}

	minifiedRendered, err := m.String("text/html", buf.String())
	if err != nil {
		panic(err)
	}
	assert.Equal(t, minifedExpected, minifiedRendered)
}

func TestDuplicateAttribute(t *testing.T) {
	attrs := make(map[string]string)
	attrs["href"] = "somelink"
	assert.Panics(t, func() { HRef("abcd")(attrs, nil, nil) })
}
