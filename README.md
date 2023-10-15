# GoDOM

GoDOM is a Go library that allows for the creation and rendering of HTML elements in a type-safe and idiomatic manner.
It draws inspiration from the [gomponents](https://github.com/maragudk/gomponents) library,
providing an alternative approach for constructing HTML in Go.

## Features

1. **Type-Safe Elements and Attributes:** GoDOM offers a strong typing system to prevent runtime errors due to incorrect
   usage of elements or attributes.
2. **Easily Extensible:** Use custom elements and attributes, or extend the existing ones.
3. **Conditional Rendering:** GoDOM supports conditional attributes and elements, which allows elements or attributes to
   be included or excluded based on certain conditions.
4. **Delayed Rendering:** Elements and attributes can be constructed immediately before rendering, allowing for dynamic
   changes between the time of construction and the rendering phase.
5. **html/template integration:** GoDOM can be used together with

## Differences from Gomponents

While both GoDOM and gomponents provide tools for constructing HTML in Go, there are notable differences:

- **Design Philosophy:** While gomponents focuses on creating a pure-functional way of building HTML, GoDOM provides a
  more direct way of working with elements and attributes.
- **Additional Features:** GoDOM offers features like delayed and conditional rendering out of the box.

## When to use which?

If you're looking for a purely functional approach, gomponents might be more suitable.
If you want a direct way of manipulating and rendering HTML with additional utilities or want to implement a reusable
set of advanced components, GoDOM is the way to go.

## Basic Usage

Here's a basic example of using GoDOM:

Imagine you're building a blog post page. The blog post might have an optional author bio, and if the blog post is
featured, it might have special stylings.

```go
package main

import (
	"bytes"
	"fmt"
	. "github.com/tbe/godom"
	"github.com/tbe/godom/helpers"
	"github.com/tbe/godom/util"
)

func main() {
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
```

## Integration with GoDOM's template package

The template package within GoDOM provides a powerful bridge between GoDOM elements and traditional HTML templating.
It offers a seamless way to combine static template constructs with dynamic GoDOM elements, enabling efficient and
optimized rendering.

### Example

Consider a scenario where you have a GoDOM element for an article, but you want the article's content and attributes to
be populated dynamically:

```go
package main

import (
	"bytes"
	. "github.com/tbe/godom"
	"github.com/tbe/godom/helpers"
	"github.com/tbe/godom/template"
)

func main() {
	// Create a new template
	tmpl := template.New("article")

	// Define the structure using GoDOM and the template placeholders
	articleElement := Div(Class("article"), tmpl.Attribute("data-id"))(
		H2()(tmpl.Placeholder("title")),
		P()(tmpl.Placeholder("content")),
	)

	// Parse the GoDOM element into the template
	template.Must(tmpl.Parse(articleElement))

	// Render the template with dynamic content
	var buf bytes.Buffer
	tmpl.Execute(&buf, &template.Context{
		Placeholders: map[string]types.Element{
			"title":   helpers.NewStringElement("Dynamic Article Title"),
			"content": helpers.NewStringElement("This is the dynamically inserted content for our article."),
		},
		Attributes: map[string]types.Attribute{
			"data-id": Data_("id", "12345"),
		},
	})

	// Output will be:
	// <div class="article" data-id="12345">
	//     <h2>Dynamic Article Title</h2>
	//     <p>This is the dynamically inserted content for our article.</p>
	// </div>
}
```

## Contribute

Contributions to GoDOM are welcome! Feel free to open issues or submit pull requests.
