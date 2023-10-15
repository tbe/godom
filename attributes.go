package godom

import (
	"bytes"
	"fmt"
	"strconv"
	"time"

	"github.com/tbe/godom/helpers"
	"github.com/tbe/godom/types"
)

const timeFormat = "2006-01-02T15:04:05-07"

type MinMaxType interface {
	int | int32 | int64 | float32 | float64 | time.Time | string
}

type NumOrString interface {
	int | int32 | int64 | float32 | float64 | string
}

// The AbbrAttr attribute specifies an abbreviated version of the content in a header cell.
//
// This attribute is allowed for:
// - TH
func AbbrAttr(abbreviation string) types.Attribute {
	return helpers.SingleAttribute("abbr", abbreviation)
}

// The Accept attribute specifies a filter for what file types the user can pick from the file input dialog box.
//
// This attribute is allowed for:
// - Input
func Accept(accept string) types.Attribute {
	return helpers.MultiValueAttribute("accept", accept)
}

// The AcceptCharset attribute specifies the character encodings that are to be used for the form submission.
//
// This attribute is allowed for:
// - Form
func AcceptCharset(characterSets ...string) types.Attribute {
	return helpers.MultiValueAttribute("accept-charset", characterSets...)
}

// The AccessKey attribute specifies a shortcut key to activate/focus an element.
//
// This is a global types.Attribute.
func AccessKey(key rune) types.Attribute {
	return helpers.SingleAttribute("accesskey", string(key))
}

// The Action attribute specifies where to send the form-data when a form is submitted.
//
// This attribute is allowed for:
// - Form
func Action(url string) types.Attribute {
	return helpers.SingleAttribute("action", url)
}

// The Allow attribute specifies a feature policy for the IFrame
//
// This attribute is allowed for:
// - IFrame
func Allow(policy string) types.Attribute {
	return helpers.SingleAttribute("allow", policy)
}

// The Alt attribute specifies an alternate text for the Element.
//
// This attribute is allowed for:
// - Area
// - Img
// - Input
func Alt(text string) types.Attribute {
	return helpers.SingleAttribute("alt", text)
}

// The Async flag specifies that the Script is downloaded in parallel to parsing the page,
// and executed as soon as it is available (before parsing completes)
//
// This flag is allowed for:
// - Script
func Async() types.Attribute {
	return helpers.FlagAttribute("async")
}

// The Autocomplete attribute specifies whether a form should have autocomplete on or off
//
// This attribute is allowed for:
// - Form
// - Input
func Autocomplete(on bool) types.Attribute {
	var flagStr string
	if on {
		flagStr = "on"
	} else {
		flagStr = "off"
	}
	return helpers.SingleAttribute("autocomplete", flagStr)
}

// The Autofocus flag specifies that an Element should automatically get focus when the page loads.
//
// This flag is allowed for:
// - Button
// - Input
// - Select
// - TextArea
func Autofocus() types.Attribute {
	return helpers.FlagAttribute("autofocus")
}

// The Autoplay flag specifies that the audio or video will start playing as soon as it is ready.
//
// This flag is allowed for:
// - Audio
// - Video
func Autoplay() types.Attribute {
	return helpers.FlagAttribute("autoplay")
}

// The Charset attribute specifies the character encoding for the HTML document.
//
// This attribute is allowed for:
// - Meta
func Charset(charset string) types.Attribute {
	return helpers.SingleAttribute("charset", charset)
}

// The Checked flag specifies that an Input element should be pre-selected when the page loads.
//
// This flag is allowed for:
// - Input
func Checked() types.Attribute {
	return helpers.FlagAttribute("checked")
}

// The CiteAttr attribute specifies the source of the quotation or a URL to a document that explains the reason why
// the text inside a Del tag was deleted.
//
// This attribute is allowed for:
// - Blockquote
// - Del
// - Ins
// - Q
//
// Note: This attribute will render to `cite="<url>"`
func CiteAttr(url string) types.Attribute {
	return helpers.SingleAttribute("cite", url)
}

// The Class attribute specifies one or more classnames for an element (refers to a class in a style sheet)
//
// This is a global types.Attribute.
func Class(classes ...string) types.Attribute {
	return helpers.MultiValueAttribute("class", classes...)
}

// The Cols attribute specifies  the visible width of a TextArea
//
// This attribute is allowed for:
// - TextArea
func Cols(width int) types.Attribute {
	return helpers.SingleAttribute("cols", strconv.Itoa(width))
}

// The ColSpan attribute specifies the number of columns a cell should span
//
// This attribute is allowed for:
// - TD
// - TH
func ColSpan(columns int) types.Attribute {
	return helpers.SingleAttribute("colspan", strconv.Itoa(columns))
}

// The ContentAttr attribute specifies the value associated with the http-equiv or name attribute.
//
// This attribute is allowed for:
// - Meta
func ContentAttr(charset string) types.Attribute {
	return helpers.SingleAttribute("content", charset)
}

// The ContentEditable attribute specifies whether the content of an element is editable or not.
//
// This is a global types.Attribute.
func ContentEditable(editable bool) types.Attribute {
	return helpers.MultiValueAttribute("contenteditable", strconv.FormatBool(editable))
}

// The Controls flag specifies that audio or video controls should be displayed.
//
// This flag is allowed for:
// - Audio
// - Video
func Controls() types.Attribute {
	return helpers.FlagAttribute("controls")
}

// The Coords attribute specifies the coordinates of the area.
//
// This attribute is allowed for:
// - Area
func Coords(coords string) types.Attribute {
	return helpers.SingleAttribute("coords", coords)
}

// The CrossOrigin allows images from third-party sites that allow cross-origin access.
//
// This attribute is allowed for:
// - Img
// - Link
// - Script
func CrossOrigin(coords string) types.Attribute {
	return helpers.SingleAttribute("crossorigin", coords)
}

// The Data_ attributes are used to store custom data private to the page or application.
//
// This is a global types.Attribute.
func Data_(key, value string) types.Attribute {
	return helpers.SingleAttribute(fmt.Sprintf("data-%s", key), value)
}

// The DataAttr specifies the URL of the resource to be used by the Object.
//
// This attribute is allowed for:
// - Object
func DataAttr(url string) types.Attribute {
	return helpers.SingleAttribute("data", url)
}

// The DateTime attribute specifies the date and time of when the text was deleted/changed.
//
// This attribute is allowed for:
// - Del
// - Ins
// - Time
func DateTime(datetime time.Time) types.Attribute {
	return helpers.SingleAttribute("datetime", datetime.Format(timeFormat))
}

// The Defer flag specifies that the script is downloaded in parallel to parsing the page, and executed after the page has finished parsing.
//
// This flag is allowed for:
// - Script
func Defer() types.Attribute {
	return helpers.FlagAttribute("defer")
}

// The Default flag specifies that the track is to be enabled if the user's preferences do not indicate that another track would be more appropriate.
//
// This flag is allowed for:
// - Track
func Default() types.Attribute {
	return helpers.FlagAttribute("default")
}

// The Dir attribute specifies the text direction for the content in an Element
//
// This is a global types.Attribute.
func Dir(direction string) types.Attribute {
	return helpers.SingleAttribute("dir", direction)
}

// The DirName attribute specifies t that the text direction will be submitted.
//
// This attribute is allowed for:
// - Input
// - TextArea
func DirName(direction string) types.Attribute {
	return helpers.SingleAttribute("dirname", direction)
}

// The Disabled flag specifies that an element should be disabled.
//
// This flag is allowed for:
// - Button
// - FieldSet
// - Input
// - OptGroup
// - Option
// - Select
// - TextArea
func Disabled() types.Attribute {
	return helpers.FlagAttribute("disabled")
}

// The Download attribute specifies that the target will be downloaded when a user clicks on the hyperlink.
//
// This attribute is allowed for:
// - A
// - Area
func Download(filename ...string) types.Attribute {
	if len(filename) == 0 {
		return helpers.FlagAttribute("download")
	} else {
		if len(filename) > 1 {
			panic("only one filename allowed")
		}
		return helpers.SingleAttribute("download", filename[0])
	}
}

// The Draggable attribute specifies whether an Element is draggable or not.
//
// This is a global types.Attribute.
func Draggable(draggable bool) types.Attribute {
	return helpers.SingleAttribute("draggable", strconv.FormatBool(draggable))
}

// The EncType attribute specifies how the form-data should be encoded when submitting it to the server.
//
// This attribute is allowed for:
// - Form
func EncType(encoding string) types.Attribute {
	return helpers.SingleAttribute("enctype", encoding)
}

// The For attribute specifies the id of the Form Element the Label or Option should be bound to.
//
// This attribute is allowed for:
// - Label
// - Output
func For(elements ...string) types.Attribute {
	return helpers.MultiValueAttribute("for", elements...)
}

// The FormID attribute specifies which form the button belongs to.
//
// This attribute is allowed for:
// - Button
// - FieldSet
// - Input
// - Label
// - Meter
// - Object
// - Output
// - Select
// - TextArea
//
// Note: This attribute will render to `form="<formID>"`
func FormID(formID string) types.Attribute {
	return helpers.SingleAttribute("form", formID)
}

var FormAttr = FormID

// The FormAction attribute specifies where to send the form-data when a form is submitted.
//
// This attribute is allowed for:
// - Button
// - Input
func FormAction(url string) types.Attribute {
	return helpers.SingleAttribute("formaction", url)
}

// The FormEncType attribute specifies how form-data should be encoded before sending it to a server.
//
// This attribute is allowed for:
// - Button
// - Input
func FormEncType(encoding string) types.Attribute {
	return helpers.SingleAttribute("formenctype", encoding)
}

// The FormMethod attribute specifies how to send the form-data (which HTTP method to use).
//
// This attribute is allowed for:
// - Button
// - Input
func FormMethod(method string) types.Attribute {
	return helpers.SingleAttribute("formmethod", method)
}

// The FormNoValidate flag specifies that the form-data should not be validated on submission.
//
// This attribute is allowed for:
// - Button
// - Input
func FormNoValidate() types.Attribute {
	return helpers.FlagAttribute("formnovalidate")
}

// The FormTarget attribute specifies where to display the response after submitting the form.
//
// This attribute is allowed for:
// - Button
// - Input
func FormTarget(target string) types.Attribute {
	return helpers.SingleAttribute("formtaget", target)
}

// The Headers attribute specifies one or more header cells a cell is related to
//
// This attribute is allowed for:
// - TD
// - TH
func Headers(headers ...string) types.Attribute {
	return helpers.MultiValueAttribute("headers", headers...)
}

// The Height attribute specifies the height of an Element.
//
// This attribute is allowed for:
// - Canvas
// - Embed
// - IFrame
// - Img
// - Input
// - Object
// - Video
func Height(pixels int) types.Attribute {
	return helpers.SingleAttribute("height", strconv.Itoa(pixels))
}

// The Hidden flag specifies the range that is considered to be a high value.
//
// This is a global types.Attribute.
func Hidden() types.Attribute {
	return helpers.FlagAttribute("hidden")
}

// The High attribute specifies the range that is considered to be a high value.
//
// This attribute is allowed for:
// - Meter
func High(high int) types.Attribute {
	return helpers.SingleAttribute("high", strconv.Itoa(high))
}

// The HRef attribute specifies the URL of the page the link goes to.
//
// This attribute is allowed for:
// - A
// - Area
// - Base
// - Link
func HRef(url string) types.Attribute {
	return helpers.SingleAttribute("href", url)
}

// The HRefLang attribute specifies the language of the linked document.
//
// This attribute is allowed for:
// - A
// - Area
// - Link
func HRefLang(languageCode string) types.Attribute {
	return helpers.SingleAttribute("hreflang", languageCode)
}

// The HTTPEquiv attribute provides an HTTP header for the information/value of the content attribute.
//
// This attribute is allowed for:
// - Meta
func HTTPEquiv(equiv string) types.Attribute {
	return helpers.SingleAttribute("http-equiv", equiv)
}

// The ID attribute specifies a unique id for an element
//
// This is a global types.Attribute.
func ID(id string) types.Attribute {
	return helpers.SingleAttribute("id", id)
}

// The Integrity attribute allows a browser to check the fetched script to ensure that the code is never loaded if the source has been manipulated.
//
// This attribute is allowed for:
// - Script
func Integrity(hash string) types.Attribute {
	return helpers.SingleAttribute("integrity", hash)
}

// The IsMap flag specifies an image as a server-side image map.
//
// This attribute is allowed for:
// - Img
func IsMap() types.Attribute {
	return helpers.FlagAttribute("ismap")
}

// The Kind attribute specifies the kind of text track.
//
// This attribute is allowed for:
// - Track
func Kind(kind string) types.Attribute {
	return helpers.SingleAttribute("kind", kind)
}

// The LabelAttr attribute specifies a label for an OptGroup, Option or Track.
//
// This attribute is allowed for:
// - OptGroup
// - Option
// - Track
func LabelAttr(label string) types.Attribute {
	return helpers.SingleAttribute("label", label)
}

// The Lang attribute specifies the language of the Element's content.
//
// This is a global types.Attribute.
func Lang(lang string) types.Attribute {
	return helpers.SingleAttribute("lang", lang)
}

// The List attribute refers to a DataList element that contains pre-defined options for an Input element
//
// This attribute is allowed for:
// - Input
func List(dataListID string) types.Attribute {
	return helpers.SingleAttribute("list", dataListID)
}

// The Loading attribute specifies whether a browser should load an Element immediately or to
// defer loading of iframes until some conditions are met.
//
// This attribute is allowed for:
// - IFrame
// - Img
func Loading(loading string) types.Attribute {
	return helpers.SingleAttribute("loading", loading)
}

// The LongDesc attribute specifies a URL to a detailed description of an image.
//
// This attribute is allowed for:
// - Img
func LongDesc(url string) types.Attribute {
	return helpers.SingleAttribute("longdesc", url)
}

// The Loop flag specifies that the audio or video will start over again, every time it is finished.
//
// This attribute is allowed for:
// - Audio
// - Video
func Loop() types.Attribute {
	return helpers.FlagAttribute("loop")
}

// The Low attribute specifies the range that is considered to be a low value.
//
// This attribute is allowed for:
// - Meter
func Low(low int) types.Attribute {
	return helpers.SingleAttribute("low", strconv.Itoa(low))
}

// The Max attribute specifies the maximum value for an Input or Meter element.
//
// This attribute is allowed for:
// - Input
// - Meter
// - Progress
func Max[T MinMaxType](max T) types.Attribute {
	return minMaxHelper("max", max)
}

// The MaxLength attribute specifies the maximum number of characters allowed in an Input element.
//
// This attribute is allowed for:
// - Input
// - TextArea
func MaxLength(length int) types.Attribute {
	return helpers.SingleAttribute("maxlength", strconv.Itoa(length))
}

// The Media attribute specifies what media/device the linked document is optimized for.
//
// This attribute is allowed for:
// - A
// - Area
// - Link
// - Source
// - Style
func Media(mediaQuery string) types.Attribute {
	return helpers.SingleAttribute("media", mediaQuery)
}

// The Method attribute specifies the HTTP method to use when sending form-data.
//
// This attribute is allowed for:
// - Form
func Method(method string) types.Attribute {
	return helpers.SingleAttribute("method", method)
}

// The Min attribute specifies a minimum value for an Input or Meter element
//
// This attribute is allowed for:
// - Input
// - Meter
func Min[T MinMaxType](min T) types.Attribute {
	return minMaxHelper("min", min)
}

// The MinLength attribute specifies the minimum number of characters required in an Input element.
//
// This attribute is allowed for:
// - Input
func MinLength(length int) types.Attribute {
	return helpers.SingleAttribute("minlength", strconv.Itoa(length))
}

// The Multiple flag specifies that a user can enter more than one value in an Input element.
//
// This attribute is allowed for:
// - Input
// - Select
func Multiple() types.Attribute {
	return helpers.FlagAttribute("multiple")
}

// The Muted flag specifies that the audio output should be muted.
//
// This flag is allowed for:
// - Audio
// - Video
func Muted() types.Attribute {
	return helpers.FlagAttribute("muted")
}

// The Name attribute specifies a name for an Element.
//
// This flag is allowed for:
// - Button
// - FieldSet
// - Form
// - IFrame
// - Input
// - Map
// - Meta
// - Object
// - Output
// - Param
// - Select
// - TextArea
func Name(name string) types.Attribute {
	return helpers.SingleAttribute("name", name)
}

// The NoModule attribute specifies that the script should not be executed in browsers supporting ES2015 modules.
//
// This flag is allowed for:
// - Script
func NoModule(noModule bool) types.Attribute {
	return helpers.SingleAttribute("nomodule", strconv.FormatBool(noModule))
}

// The NoValidate flag specifies that the form should not be validated when submitted.
//
// This attribute is allowed for:
// - Form
func NoValidate() types.Attribute {
	return helpers.FlagAttribute("novalidate")
}

// The Open flag specifies that an Element should be visible (open) to the user.
//
// This flag is allowed for:
// - Details
// - Dialog
func Open() types.Attribute {
	return helpers.FlagAttribute("open")
}

// The Optimum attribute specifies what value is the optimal value for the gauge.
//
// This attribute is allowed for:
// - Meter
func Optimum(optimum int) types.Attribute {
	return helpers.SingleAttribute("optimum", strconv.Itoa(optimum))
}

// The Pattern specifies a JavaScript regular expression that an Input element's value is checked against.
//
// This attribute is allowed for:
// - Input
func Pattern(pattern string) types.Attribute {
	return helpers.SingleAttribute("pattern", pattern)
}

// The Ping attribute specifies a list of URLs to which, when the link is followed,
// post requests with the body ping will be sent by the browser.
//
// This attribute is allowed for:
// - A
func Ping(urls ...string) types.Attribute {
	return helpers.MultiValueAttribute("ping", urls...)
}

// The Placeholder attribute specifies a short hint that describes the expected value of an Input element.
//
// This attribute is allowed for:
// - Input
// - TextArea
func Placeholder(placeholder string) types.Attribute {
	return helpers.SingleAttribute("placeholder", placeholder)
}

// The Poster attribute specifies an image to be shown while the Video is downloading,
// or until the user hits the play button.
//
// This attribute is allowed for:
// - Video
func Poster(url string) types.Attribute {
	return helpers.SingleAttribute("poster", url)
}

// The Preload attribute specifies if and how the author thinks the audio or video should be loaded when the page loads.
//
// This attribute is allowed for:
// - Audio
// - Video
func Preload(preload string) types.Attribute {
	return helpers.SingleAttribute("preload", preload)
}

// The ReadOnly flag specifies that an Input field is read-only
//
// This flag is allowed for:
// - Input
// - TextArea
func ReadOnly() types.Attribute {
	return helpers.FlagAttribute("readonly")
}

// The ReferrerPolicy attribute specifies which referrer information to send with the link, IFrame or Img.
//
// This attribute is allowed for:
// - A
// - Area
// - IFrame
// - Img
// - Link
// - Script
func ReferrerPolicy(policy string) types.Attribute {
	return helpers.SingleAttribute("referrerpolicy", policy)
}

// The Rel attribute specifies the relationship between the current document and the linked document.
//
// This attribute is allowed for:
// - A
// - Area
// - Form
// - Link
func Rel(relationship string) types.Attribute {
	return helpers.SingleAttribute("rel", relationship)
}

// The Required flag specifies that an Input field must be filled out before submitting the form.
//
// This flag is allowed for:
// - Input
// - Select
// - TextArea
func Required() types.Attribute {
	return helpers.FlagAttribute("required")
}

// The Reversed flag specifies that the list order should be reversed (9,8,7...)
//
// This flag is allowed for:
// - OL
func Reversed() types.Attribute {
	return helpers.FlagAttribute("reversed")
}

// The Rows attribute sets the visible number of lines in a TextArea.
//
// This attribute is allowed for:
// - TextArea
func Rows(rows int) types.Attribute {
	return helpers.SingleAttribute("rows", strconv.Itoa(rows))
}

// The RowSpan attribute sets the number of rows a cell should span.
//
// This attribute is allowed for:
// - TD
// - TH
func RowSpan(rows int) types.Attribute {
	return helpers.SingleAttribute("rowspan", strconv.Itoa(rows))
}

// The Sandbox attribute enables an extra set of restrictions for the content in an IFrame.
//
// This attribute is allowed for:
// - IFrame
func Sandbox(restrictions ...string) types.Attribute {
	if len(restrictions) == 0 {
		return helpers.FlagAttribute("sandbox")
	}
	return helpers.MultiValueAttribute("sandbox", restrictions...)
}

// The Scope attribute specifies whether a header cell is a header for a column, row, or group of columns or rows.
//
// This attribute is allowed for:
// - TH
func Scope(scope string) types.Attribute {
	return helpers.SingleAttribute("scope", scope)
}

// The Selected flag specifies that an Option should be pre-selected when the page loads
//
// This attribute is allowed for:
// - Option
func Selected() types.Attribute {
	return helpers.FlagAttribute("selected")
}

// The Shape attribute specifies the shape of the area.
//
// This attribute is allowed for:
// - Area
func Shape(shape string) types.Attribute {
	return helpers.SingleAttribute("shape", shape)
}

// The Size attribute specifies the width, in characters, of an Input element or the number of visible options in the
// Select list
//
// This attribute is allowed for:
// - Input
// - Select
func Size(size int) types.Attribute {
	return helpers.SingleAttribute("size", strconv.Itoa(size))
}

// The Sizes attribute specifies image sizes for different page layouts.
//
// This attribute is allowed for:
// - Img
// - Link
// - Source
func Sizes(sizes string) types.Attribute {
	return helpers.SingleAttribute("sizes", sizes)
}

// The SpanAttr attribute specifies the number of columns a Col or ColGroup should span
//
// This attribute is allowed for:
// - Col
// - ColGroup
//
// Note: This attribute will render to `span="<columns>"`
func SpanAttr(columns int) types.Attribute {
	return helpers.SingleAttribute("span", strconv.Itoa(columns))
}

// The Spellcheck attribute specifies whether the element is to have its spelling and grammar checked or not
//
// This is a global types.Attribute.
func Spellcheck(check bool) types.Attribute {
	return helpers.SingleAttribute("spellcheck", strconv.FormatBool(check))
}

// The Src attribute specifies the URL of external content for an Element.
//
// This attribute is allowed for:
// - Audio
// - Embed
// - IFrame
// - Img
// - Input
// - Script
// - Source
// - Track
// - Video
func Src(url string) types.Attribute {
	return helpers.SingleAttribute("src", url)
}

// The SrcDoc attribute specifies the HTML content of the page to show in the IFrame
//
// This attribute is allowed for:
// - IFrame
func SrcDoc(elements ...types.Element) types.Attribute {
	// this is special. The srcdoc attribute accepts HTML ...
	buf := new(bytes.Buffer)
	for _, elem := range elements {
		if err := elem.Render(buf); err != nil {
			panic(fmt.Sprintf("failed to render srcdoc attribute: %v", err))
		}
	}
	return helpers.SingleAttribute("srcdoc", buf.String())
}

// The SrcLang attribute specifies the language of the track text data.
//
// This attribute is allowed for:
// - Track
func SrcLang(lang string) types.Attribute {
	return helpers.SingleAttribute("srclang", lang)
}

// The SrcSet attribute specifies a list of image files to use in different situations.
//
// This attribute is allowed for:
// - Img
// - Source
func SrcSet(urlList string) types.Attribute {
	return helpers.SingleAttribute("srcset", urlList)
}

// The Start attribute specifies the start value of an ordered list.
//
// This attribute is allowed for:
// - OL
func Start(start int) types.Attribute {
	return helpers.SingleAttribute("start", strconv.Itoa(start))
}

// The Step attribute specifies the interval between legal numbers in an Input field.
//
// This attribute is allowed for:
// - Input
func Step[T NumOrString](step T) types.Attribute {
	return numOrStringHelper("step", step)
}

// The StyleAttr attribute specifies an inline CSS style for an Element.
//
// This is a global types.Attribute.
func StyleAttr(style string) types.Attribute {
	return helpers.SingleAttribute("style", style)
}

// The TabIndex attribute specifies the tabbing order of an Element.
//
// This is a global types.Attribute.
func TabIndex(idx int) types.Attribute {
	return helpers.SingleAttribute("tabindex", strconv.Itoa(idx))
}

// The Target attribute specifies where to open the linked document.
//
// This attribute is allowed for:
// - A
// - Area
// - Base
// - Form
func Target(target string) types.Attribute {
	return helpers.SingleAttribute("target", target)
}

// The TitleAttr attribute specifies extra information about an element.
//
// This is a global types.Attribute.
func TitleAttr(title string) types.Attribute {
	return helpers.SingleAttribute("title", title)
}

// The Translate attribute specifies extra information about an element.
//
// This is a global types.Attribute.
func Translate(translate bool) types.Attribute {
	var translateStr string
	if translate {
		translateStr = "yes"
	} else {
		translateStr = "no"
	}
	return helpers.SingleAttribute("translate", translateStr)
}

// The Type attribute specifies the media type of the linked document, or the type of an Element.
//
// This attribute is allowed for:
// - A
// - Area
// - Button
// - Embed
// - Input
// - Link
// - Object
// - OL
// - Script
// - Source
// - Style
func Type(typ string) types.Attribute {
	return helpers.SingleAttribute("type", typ)
}

// The TypeMustMatch attribute specifies whether the type attribute and the actual content
// of the resource must match to be displayed.
//
// This attribute is allowed for:
// - Object
func TypeMustMatch(mustMatch bool) types.Attribute {
	return helpers.SingleAttribute("typemustmatch", strconv.FormatBool(mustMatch))
}

// The UseMap attribute specifies an image as a client-side image map.
//
// This attribute is allowed for:
// - Img
// - Object
func UseMap(mapName string) types.Attribute {
	return helpers.SingleAttribute("usemap", mapName)
}

// The Value attribute specifies an initial value for the Element or the machine-readable translation of the content of a Data element.
//
// This attribute is allowed for:
// - Button
// - Data
// - Input
// - Li
// - Meter
// - Option
// - Param
// - Progress
func Value[T NumOrString](value T) types.Attribute {
	return numOrStringHelper("value", value)
}

// The Width attribute specifies the width of an Element.
//
// This attribute is allowed for:
// - Canvas
// - Embed
// - IFrame
// - Img
// - Input
// - Object
// - Video
func Width(pixels int) types.Attribute {
	return helpers.SingleAttribute("width", strconv.Itoa(pixels))
}

// The Wrap attribute specifies how the text in a TextArea is to be wrapped when submitted in a form.
//
// This attribute is allowed for:
// - TextArea
func Wrap(wrap string) types.Attribute {
	return helpers.SingleAttribute("wrap", wrap)
}

// The XMLNS attribute specifies the XML namespace attribute.
//
// This attribute is allowed for:
// - HTML
func XMLNS(namespace string) types.Attribute {
	return helpers.SingleAttribute("xmlns", namespace)
}

func numOrStringHelper[T NumOrString](attr string, value T) types.Attribute {
	// we use a generic here as a compile time barrier.
	// But, as generics don't allow direct type switches, we have to do this stunt ...
	switch any(value).(type) {
	case int, int32, int64:
		return helpers.SingleAttribute(attr, intPrinter(value))
	case float32, float64:
		return helpers.SingleAttribute(attr, floatPrinter(value))
	case string:
		return helpers.SingleAttribute(attr, any(value).(string))
	default:
		panic("unknown type")
	}
}

func minMaxHelper[T MinMaxType](attr string, value T) types.Attribute {
	// same as with the numOrStringHelper
	switch any(value).(type) {
	case int, int32, int64:
		return helpers.SingleAttribute(attr, intPrinter(value))
	case time.Time:
		return helpers.SingleAttribute(attr, any(value).(time.Time).Format(timeFormat))
	case float32, float64:
		return helpers.SingleAttribute(attr, floatPrinter(value))
	case string:
		return helpers.SingleAttribute(attr, any(value).(string))
	default:
		panic("unknown type")
	}
}

func intPrinter(i any) string {
	return fmt.Sprintf("%d", i)
}

func floatPrinter(f any) string {
	return fmt.Sprintf("%f", f)
}
