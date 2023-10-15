package godom

import (
	"html"
	"io"

	"github.com/tbe/godom/helpers"
	"github.com/tbe/godom/types"
)

// Group is a wrapper that can hold 0...N children. This allows to hold a full document
// or can be used in places where multiple elements are required, but only a single types.Element is allowed.
func Group(children ...types.Element) types.Element {
	return &container{children: children}
}

// The Doctype provides the DOCTYPE declaration
func Doctype() types.Element {
	return helpers.NewStringElement("<!DOCTYPE html>")
}

// The A tag defines a hyperlink, which is used to link from one page to another.
func A(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("a", attrs...)
}

// The Abbr tag defines an abbreviation or an acronym.
func Abbr(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("abbr", attrs...)
}

// The Address tag defines the contact information for the author/owner of a document or an article.
func Address(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("address", attrs...)
}

// The Area tag defines an area inside an image map.
func Area(attrs ...types.Attribute) types.ElementFactory {
	// TODO: the area element is only allowed inside a map, so we should limit this
	return helpers.NewElement("area", attrs...)
}

// The Article tag specifies independent, self-contained content.
func Article(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("article", attrs...)
}

// The Aside tag defines some content aside from the content it is placed in.
func Aside(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("aside", attrs...)
}

// The Audio tag is used to embed sound content in a document, such as music or other audio streams.
func Audio(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("audio", attrs...)
}

// The B tag specifies bold text without any extra importance.
func B(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("b", attrs...)
}

// The Base tag specifies the base URL and/or target for all relative URLs in a document.
func Base(attrs ...types.Attribute) types.Element {
	return helpers.NewChildlessElement("base", attrs...)
}

// The BDI tag isolates a part of text that might be formatted in a different direction from other text outside it.
func BDI(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("bdi", attrs...)
}

// The BDO tag is used to override the current text direction.
func BDO(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("bdo", attrs...)
}

// The Blockquote tag specifies a section that is quoted from another source.
func Blockquote(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("blockquote", attrs...)
}

// The Body tag defines the document's body.
func Body(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("body", attrs...)
}

// The Br tag inserts a single line break.
func Br(attrs ...types.Attribute) types.Element {
	return helpers.NewChildlessElement("br", attrs...)
}

// The Button tag defines a clickable button.
func Button(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("button", attrs...)
}

// The Canvas tag is used to draw graphics, on the fly, via scripting (usually JavaScript).
func Canvas(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("canvas", attrs...)
}

// The Caption tag defines a table caption.
func Caption(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("caption", attrs...)
}

// The Cite tag defines the title of a creative work (e.g. a book, a poem, a song, a movie, a painting, a sculpture, etc.).
func Cite(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("cite", attrs...)
}

// The Code tag is used to define a piece of computer code.
// The content inside is displayed in the browser's default monospace font.
func Code(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("code", attrs...)
}

// The Col tag specifies column properties for each column within a ColGroup element.
func Col(attrs ...types.Attribute) types.Element {
	return helpers.NewChildlessElement("col", attrs...)
}

// The ColGroup tag specifies a group of one or more columns in a table for formatting.
func ColGroup(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("colgroup", attrs...)
}

// The Data tag is used to add a machine-readable translation of a given content.
func Data(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("data", attrs...)
}

// The DataList tag specifies a list of pre-defined options for an <input> element.
func DataList(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("datalist", attrs...)
}

// The DD tag is used to describe a term/name in a description list.
func DD(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("dd", attrs...)
}

// The Del tag defines text that has been deleted from a document.
func Del(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("del", attrs...)
}

// The Details tag specifies additional details that the user can open and close on demand.
func Details(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("details", attrs...)
}

// The Dfn tag stands for the "definition element", and it specifies a term that is going to be defined within the content.
func Dfn(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("dfn", attrs...)
}

// The Dialog tag defines a dialog box or subwindow.
func Dialog(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("dialog", attrs...)
}

// The Div tag defines a division or a section in an HTML document.
func Div(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("div", attrs...)
}

// The DL tag defines a description list.
func DL(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("dl", attrs...)
}

// The DT tag defines a term/name in a description list.
func DT(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("dt", attrs...)
}

// The Em tag is used to define emphasized text. The content inside is typically displayed in italic.
func Em(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("em", attrs...)
}

// The Embed tag defines a container for an external resource, such as a web page, a picture, a media player, or a plug-in application.
func Embed(attrs ...types.Attribute) types.Element {
	return helpers.NewChildlessElement("embed", attrs...)
}

// The FieldSet tag is used to group related elements in a form.
func FieldSet(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("fieldset", attrs...)
}

// The FigCaption tag defines a caption for a Figure element.
func FigCaption(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("figcaption", attrs...)
}

// The Figure tag specifies self-contained content, like illustrations, diagrams, photos, code listings, etc.
func Figure(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("figure", attrs...)
}

// The Footer tag defines a footer for a document or section.
func Footer(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("footer", attrs...)
}

// The Form tag is used to create an HTML form for user input.
func Form(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("form", attrs...)
}

// The H1 tag is used to define a level 1 HTML Heading
func H1(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("h1", attrs...)
}

// The H2 tag is used to define a level 2 HTML Heading
func H2(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("h2", attrs...)
}

// The H3 tag is used to define a level 3 HTML Heading
func H3(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("h3", attrs...)
}

// The H4 tag is used to define a level 4 HTML Heading
func H4(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("h4", attrs...)
}

// The H5 tag is used to define a level 5 HTML Heading
func H5(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("h5", attrs...)
}

// The H6 tag is used to define a level 6 HTML Heading
func H6(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("h6", attrs...)
}

// The Head element is a container for metadata (data about data) and is placed between the HTML tag and the Body tag.
func Head(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("head", attrs...)
}

// The Header element represents a container for introductory content or a set of navigational links.
func Header(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("header", attrs...)
}

// The HR tag defines a thematic break in an HTML page (e.g. a shift of topic).
func HR(attrs ...types.Attribute) types.Element {
	return helpers.NewChildlessElement("hr", attrs...)
}

// The HTML tag represents the root of an HTML document.
func HTML(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("html", attrs...)
}

// The I tag defines a part of text in an alternate voice or mood. The content inside is typically displayed in italic.
func I(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("i", attrs...)
}

// The IFrame tag specifies an inline frame.
func IFrame(attrs ...types.Attribute) types.Element {
	return helpers.NewElement("iframe", attrs...)()
}

// The Img tag is used to embed an image in an HTML page.
func Img(attrs ...types.Attribute) types.Element {
	return helpers.NewChildlessElement("img", attrs...)
}

// The Input tag specifies an input field where the user can enter data.
func Input(attrs ...types.Attribute) types.Element {
	return helpers.NewChildlessElement("input", attrs...)
}

// The Ins tag defines a text that has been inserted into a document.
func Ins(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("ins", attrs...)
}

// The Kbd tag is used to define keyboard input.
// The content inside is displayed in the browser's default monospace font.
func Kbd(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("kbd", attrs...)
}

// The Label tag defines a label for several elements:
// - Input
// - Meter
// - Progress
// - Select
// - Textarea
func Label(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("label", attrs...)
}

// The Legend tag defines a caption for the FieldSet element.
func Legend(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("legend", attrs...)
}

// The Li tag defines a list item.
func Li(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("li", attrs...)
}

// The Link tag defines the relationship between the current document and an external resource.
func Link(attrs ...types.Attribute) types.Element {
	return helpers.NewChildlessElement("link", attrs...)
}

// The Main tag specifies the main content of a document.
func Main(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("main", attrs...)
}

// The Map tag is used to define an image map. An image map is an image with clickable areas.
func Map(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("map", attrs...)
}

// The Mark tag defines text that should be marked or highlighted.
func Mark(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("mark", attrs...)
}

// The Meta tag defines metadata about an HTML document. Metadata is data (information) about data.
func Meta(attrs ...types.Attribute) types.Element {
	return helpers.NewChildlessElement("meta", attrs...)
}

// The Meter tag defines a scalar measurement within a known range, or a fractional value. This is also known as a gauge.
func Meter(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("meter", attrs...)
}

// The Nav tag defines a set of navigation links.
func Nav(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("nav", attrs...)
}

// The NoScript tag defines an alternate content to be displayed to users that have disabled scripts in their browser
// or have a browser that doesn't support script.
func NoScript(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("noscript", attrs...)
}

// The Object tag defines a container for an external resource.
func Object(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("object", attrs...)
}

// The OL tag defines an ordered list. An ordered list can be numerical or alphabetical.
func OL(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("ol", attrs...)
}

// The OptGroup tag is used to group related options in a Select element (drop-down list).
func OptGroup(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("optgroup", attrs...)
}

// The Option tag defines an option in a Select list.
func Option(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("option", attrs...)
}

// The Output tag is used to represent the result of a calculation (like one performed by a script).
func Output(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("output", attrs...)
}

// The P tag defines a paragraph.
func P(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("p", attrs...)
}

// The Param tag is used to define parameters for an Object element.
func Param(attrs ...types.Attribute) types.Element {
	return helpers.NewChildlessElement("param", attrs...)
}

// The Picture tag gives web developers more flexibility in specifying image resources.
func Picture(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("picture", attrs...)
}

// The Pre tag defines preformatted text.
func Pre(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("pre", attrs...)
}

// The Progress tag represents the completion progress of a task.
func Progress(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("progress", attrs...)
}

// The Q tag defines a short quotation.
func Q(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("q", attrs...)
}

// The RP tag can be used to provide parentheses around a ruby text, to be shown by browsers that do not support ruby annotations.
func RP(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("rp", attrs...)
}

// The RT tag defines an explanation or pronunciation of characters (for East Asian typography) in a ruby annotation.
func RT(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("rt", attrs...)
}

// The Ruby tag specifies a ruby annotation.
func Ruby(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("ruby", attrs...)
}

// The S tag specifies text that is no longer correct, accurate or relevant.
// The text will be displayed with a line through it.
func S(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("s", attrs...)
}

// The Samp tag is used to define sample output from a computer program.
// The content inside is displayed in the browser's default monospace font.
func Samp(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("samp", attrs...)
}

// The Script tag is used to embed a client-side script (JavaScript).
func Script(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("script", attrs...)
}

// The Section tag defines a section in a document.
func Section(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("section", attrs...)
}

// The Select element is used to create a drop-down list.
func Select(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("select", attrs...)
}

// The Small tag defines smaller text (like copyright and other side-comments).
func Small(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("small", attrs...)
}

// The Source tag is used to specify multiple media resources for media elements, such as Video, Audio, and Picture.
func Source(attrs ...types.Attribute) types.Element {
	return helpers.NewChildlessElement("source", attrs...)
}

// The Span tag is an inline container used to mark up a part of a text, or a part of a document.
func Span(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("span", attrs...)
}

// The Strong tag is used to define text with strong importance. The content inside is typically displayed in bold.
func Strong(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("strong", attrs...)
}

// The Style tag is used to define style information (CSS) for a document.
func Style(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("style", attrs...)
}

// The Sub tag defines subscript text.
func Sub(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("sub", attrs...)
}

// The Summary tag defines a visible heading for the Details element.
// The heading can be clicked to view/hide the details.
func Summary(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("summary", attrs...)
}

// The Sup tag defines superscript text.
func Sup(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("sup", attrs...)
}

// The SVG tag defines a container for svg graphics.
func SVG(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("svg", attrs...)
}

// The Table tag defines an HTML table.
func Table(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("table", attrs...)
}

// The TBody tag is used to group the body content in an HTML Table.
func TBody(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("tbody", attrs...)
}

// The TD tag defines a standard data cell in an HTML Table.
func TD(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("td", attrs...)
}

// The Template tag is used as a container to hold some HTML content hidden from the user when the page loads.
func Template(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("template", attrs...)
}

// The TextArea tag defines a multi-line text input control.
func TextArea(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("textarea", attrs...)
}

// The TFoot tag is used to group footer content in an HTML Table.
func TFoot(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("tfoot", attrs...)
}

// The TH tag defines a header cell in an HTML Table.
func TH(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("th", attrs...)
}

// The THead tag is used to group header content in an HTML Table.
func THead(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("thead", attrs...)
}

// The Time tag defines a specific time (or datetime).
func Time(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("time", attrs...)
}

// The Title tag defines the title of the document.
func Title(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("title", attrs...)
}

// The TR tag defines a row in an HTML Table.
func TR(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("tr", attrs...)
}

// The Track tag specifies text tracks for Audio or Video elements.
func Track(attrs ...types.Attribute) types.Element {
	return helpers.NewChildlessElement("track", attrs...)
}

// The U tag represents some text that is unarticulated and styled differently from normal text.
func U(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("u", attrs...)
}

// The UL tag defines an unordered (bulleted) list.
func UL(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("ul", attrs...)
}

// The Var tag is used to defines a variable in programming or in a mathematical expression.
func Var(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("var", attrs...)
}

// The Video tag is used to embed video content in a document, such as a movie clip or other video streams.
func Video(attrs ...types.Attribute) types.ElementFactory {
	return helpers.NewElement("video", attrs...)
}

// The WBr  (Word Break Opportunity) tag specifies where in a text it would be ok to add a line-break.
func WBr(attrs ...types.Attribute) types.Element {
	return helpers.NewChildlessElement("wbr", attrs...)
}

// The Content types.Element renders the given content as HTML escaped text.
func Content(content string) types.Element {
	return helpers.NewStringElement(html.EscapeString(content))
}

type container struct {
	children []types.Element
}

func (c *container) Render(writer io.Writer) error {
	for _, child := range c.children {
		if err := child.Render(writer); err != nil {
			return err
		}
	}
	return nil
}
