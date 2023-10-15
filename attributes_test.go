package godom_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	. "github.com/tbe/godom"
	"github.com/tbe/godom/types"
)

func TestAttributes(t *testing.T) {
	suite.Run(t, new(AttributesTestSuite))
}

type AttributesTestSuite struct {
	suite.Suite
	state      attrState
	refTime    time.Time
	refTimeStr string
}

func (suite *AttributesTestSuite) SetupSuite() {
	// we construct a time for us
	tz, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		panic("can not load timezone")
	}

	suite.refTime = time.Date(2023, 10, 13, 18, 30, 00, 00, tz)
	suite.refTimeStr = "2023-10-13T18:30:00+02"
}

func (suite *AttributesTestSuite) SetupTest() {
	suite.state.reset()
}

func (suite *AttributesTestSuite) TestAbbr() {
	suite.testAttr(AbbrAttr("test"), "abbr", "test")
}

func (suite *AttributesTestSuite) TestAccept() {
	suite.testAttr(Accept("test"), "accept", "test")
}

func (suite *AttributesTestSuite) TestAcceptCharset() {
	suite.testAttr(AcceptCharset("test"), "accept-charset", "test")
}

func (suite *AttributesTestSuite) TestAccessKey() {
	suite.testAttr(AccessKey('a'), "accesskey", "a")
}

func (suite *AttributesTestSuite) TestAction() {
	suite.testAttr(Action("test"), "action", "test")
}

func (suite *AttributesTestSuite) TestAllow() {
	suite.testAttr(Allow("test"), "allow", "test")
}

func (suite *AttributesTestSuite) TestAlt() {
	suite.testAttr(Alt("test"), "alt", "test")
}

func (suite *AttributesTestSuite) TestAsync() {
	suite.testFlag(Async(), "async")
}

func (suite *AttributesTestSuite) TestAutocompleteOn() {
	suite.testAttr(Autocomplete(true), "autocomplete", "on")
}

func (suite *AttributesTestSuite) TestAutocompleteOff() {
	suite.testAttr(Autocomplete(false), "autocomplete", "off")
}

func (suite *AttributesTestSuite) TestAutofocus() {
	suite.testFlag(Autofocus(), "autofocus")
}

func (suite *AttributesTestSuite) TestAutoplay() {
	suite.testFlag(Autoplay(), "autoplay")
}

func (suite *AttributesTestSuite) TestCharset() {
	suite.testAttr(Charset("test"), "charset", "test")
}

func (suite *AttributesTestSuite) TestChecked() {
	suite.testFlag(Checked(), "checked")
}

func (suite *AttributesTestSuite) TestCite() {
	suite.testAttr(CiteAttr("test"), "cite", "test")
}

func (suite *AttributesTestSuite) TestSingleClass() {
	suite.testAttr(Class("test"), "class", "test")
}

func (suite *AttributesTestSuite) TestMultipleClasses() {
	suite.testAttr(Class("testA"), "class", "testA")
	suite.testAttr(Class("testB", "testC"), "class", "testA testB testC")
}

func (suite *AttributesTestSuite) TestCols() {
	suite.testAttr(Cols(10), "cols", "10")
}

func (suite *AttributesTestSuite) TestColSpan() {
	suite.testAttr(ColSpan(10), "colspan", "10")
}

func (suite *AttributesTestSuite) TestContent() {
	suite.testAttr(ContentAttr("test"), "content", "test")
}

func (suite *AttributesTestSuite) TestContentEditableTrue() {
	suite.testAttr(ContentEditable(true), "contenteditable", "true")
}

func (suite *AttributesTestSuite) TestContentEditableFalse() {
	suite.testAttr(ContentEditable(false), "contenteditable", "false")
}

func (suite *AttributesTestSuite) TestControls() {
	suite.testFlag(Controls(), "controls")
}

func (suite *AttributesTestSuite) TestCoords() {
	suite.testAttr(Coords("test"), "coords", "test")
}

func (suite *AttributesTestSuite) TestCrossOrigin() {
	suite.testAttr(CrossOrigin("test"), "crossorigin", "test")
}

func (suite *AttributesTestSuite) TestData() {
	suite.testAttr(DataAttr("test"), "data", "test")
}

func (suite *AttributesTestSuite) TestData_() {
	suite.testAttr(Data_("x", "test"), "data-x", "test")
}

func (suite *AttributesTestSuite) TestDateTime() {

	suite.testAttr(DateTime(suite.refTime), "datetime", suite.refTimeStr)
}

func (suite *AttributesTestSuite) TestDefault() {
	suite.testFlag(Default(), "default")
}

func (suite *AttributesTestSuite) TestDefer() {
	suite.testFlag(Defer(), "defer")
}

func (suite *AttributesTestSuite) TestDir() {
	suite.testAttr(Dir("test"), "dir", "test")
}

func (suite *AttributesTestSuite) TestDirName() {
	suite.testAttr(DirName("test"), "dirname", "test")
}

func (suite *AttributesTestSuite) TestDisabled() {
	suite.testFlag(Disabled(), "disabled")
}

func (suite *AttributesTestSuite) TestDownloadFlag() {
	suite.testFlag(Download(), "download")
}

func (suite *AttributesTestSuite) TestDownloadString() {
	suite.testAttr(Download("test"), "download", "test")
}

func (suite *AttributesTestSuite) TestDraggableTrue() {
	suite.testAttr(Draggable(true), "draggable", "true")
}

func (suite *AttributesTestSuite) TestDraggableFalse() {
	suite.testAttr(Draggable(false), "draggable", "false")
}

func (suite *AttributesTestSuite) TestEnctype() {
	suite.testAttr(EncType("test"), "enctype", "test")
}

func (suite *AttributesTestSuite) TestFor() {
	suite.testAttr(For("test"), "for", "test")
}

func (suite *AttributesTestSuite) TestForm() {
	suite.testAttr(FormID("test"), "form", "test")
}

func (suite *AttributesTestSuite) TestFormAttr() {
	suite.testAttr(FormAttr("test"), "form", "test")
}

func (suite *AttributesTestSuite) TestFormAction() {
	suite.testAttr(FormAction("test"), "formaction", "test")
}

func (suite *AttributesTestSuite) TestFormEncType() {
	suite.testAttr(FormEncType("test"), "formenctype", "test")
}

func (suite *AttributesTestSuite) TestFormMethod() {
	suite.testAttr(FormMethod("test"), "formmethod", "test")
}

func (suite *AttributesTestSuite) TestFormNoValidate() {
	suite.testFlag(FormNoValidate(), "formnovalidate")
}

func (suite *AttributesTestSuite) TestFormTarget() {
	suite.testAttr(FormTarget("test"), "formtaget", "test")
}

func (suite *AttributesTestSuite) TestHeaders() {
	suite.testAttr(Headers("test"), "headers", "test")
}

func (suite *AttributesTestSuite) TestHeadersMulti() {
	suite.testAttr(Headers("testA", "testB"), "headers", "testA testB")
	suite.testAttr(Headers("testC"), "headers", "testA testB testC")
}

func (suite *AttributesTestSuite) TestHeight() {
	suite.testAttr(Height(10), "height", "10")
}

func (suite *AttributesTestSuite) TestHidden() {
	suite.testFlag(Hidden(), "hidden")
}

func (suite *AttributesTestSuite) TestHigh() {
	suite.testAttr(High(10), "high", "10")
}

func (suite *AttributesTestSuite) TestHRef() {
	suite.testAttr(HRef("test"), "href", "test")
}

func (suite *AttributesTestSuite) TestHRefLang() {
	suite.testAttr(HRefLang("test"), "hreflang", "test")
}

func (suite *AttributesTestSuite) TestHTTPEquiv() {
	suite.testAttr(HTTPEquiv("test"), "http-equiv", "test")
}

func (suite *AttributesTestSuite) TestID() {
	suite.testAttr(ID("test"), "id", "test")
}

func (suite *AttributesTestSuite) TestIntegrity() {
	suite.testAttr(Integrity("test"), "integrity", "test")
}

func (suite *AttributesTestSuite) TestIsMap() {
	suite.testFlag(IsMap(), "ismap")
}

func (suite *AttributesTestSuite) TestKind() {
	suite.testAttr(Kind("test"), "kind", "test")
}

func (suite *AttributesTestSuite) TestLabel() {
	suite.testAttr(LabelAttr("test"), "label", "test")
}

func (suite *AttributesTestSuite) TestLang() {
	suite.testAttr(Lang("test"), "lang", "test")
}

func (suite *AttributesTestSuite) TestList() {
	suite.testAttr(List("test"), "list", "test")
}

func (suite *AttributesTestSuite) TestLoading() {
	suite.testAttr(Loading("test"), "loading", "test")
}

func (suite *AttributesTestSuite) TestLongDesc() {
	suite.testAttr(LongDesc("test"), "longdesc", "test")
}

func (suite *AttributesTestSuite) TestLoop() {
	suite.testFlag(Loop(), "loop")
}

func (suite *AttributesTestSuite) TestLow() {
	suite.testAttr(Low(10), "low", "10")
}

func (suite *AttributesTestSuite) TestMaxInt() {
	suite.testAttr(Max(10), "max", "10")
}

func (suite *AttributesTestSuite) TestMaxTime() {
	suite.testAttr(Max(suite.refTime), "max", suite.refTimeStr)
}

func (suite *AttributesTestSuite) TestMaxFloat32() {
	suite.testAttr(Max(float32(1.2)), "max", "1.200000")
}

func (suite *AttributesTestSuite) TestMaxFloat64() {
	suite.testAttr(Max(float64(1.2)), "max", "1.200000")
}

func (suite *AttributesTestSuite) TestMaxString() {
	suite.testAttr(Max("test"), "max", "test")
}

func (suite *AttributesTestSuite) TestMaxLength() {
	suite.testAttr(MaxLength(10), "maxlength", "10")
}

func (suite *AttributesTestSuite) TestMedia() {
	suite.testAttr(Media("test"), "media", "test")
}

func (suite *AttributesTestSuite) TestMethod() {
	suite.testAttr(Method("test"), "method", "test")
}

func (suite *AttributesTestSuite) TestMinInt() {
	suite.testAttr(Min(10), "min", "10")
}

func (suite *AttributesTestSuite) TestMinTime() {
	suite.testAttr(Min(suite.refTime), "min", suite.refTimeStr)
}

func (suite *AttributesTestSuite) TestMinFloat32() {
	suite.testAttr(Min(float32(1.2)), "min", "1.200000")
}

func (suite *AttributesTestSuite) TestMinFloat64() {
	suite.testAttr(Min(float64(1.2)), "min", "1.200000")
}

func (suite *AttributesTestSuite) TestMinString() {
	suite.testAttr(Min("test"), "min", "test")
}

func (suite *AttributesTestSuite) TestMinLength() {
	suite.testAttr(MinLength(10), "minlength", "10")
}

func (suite *AttributesTestSuite) TestMultiple() {
	suite.testFlag(Multiple(), "multiple")
}

func (suite *AttributesTestSuite) TestMuted() {
	suite.testFlag(Muted(), "muted")
}

func (suite *AttributesTestSuite) TestName() {
	suite.testAttr(Name("test"), "name", "test")
}

func (suite *AttributesTestSuite) TestNoModuleTrue() {
	suite.testAttr(NoModule(true), "nomodule", "true")
}

func (suite *AttributesTestSuite) TestNoModuleFalse() {
	suite.testAttr(NoModule(false), "nomodule", "false")
}

func (suite *AttributesTestSuite) TestNoValidate() {
	suite.testFlag(NoValidate(), "novalidate")
}

func (suite *AttributesTestSuite) TestOnAfterPrint() {
	suite.testAttr(OnAfterPrint("test"), "onafterprint", "test")
}

func (suite *AttributesTestSuite) TestOnBeforeUnload() {
	suite.testAttr(OnBeforeUnload("test"), "onbeforeunload", "test")
}

func (suite *AttributesTestSuite) TestOnBlur() {
	suite.testAttr(OnBlur("test"), "onblur", "test")
}

func (suite *AttributesTestSuite) TestOnCanPlay() {
	suite.testAttr(OnCanPlay("test"), "oncanplay", "test")
}

func (suite *AttributesTestSuite) TestOnCanPlayThrough() {
	suite.testAttr(OnCanPlayThrough("test"), "oncanplaythrough", "test")
}

func (suite *AttributesTestSuite) TestOnChange() {
	suite.testAttr(OnChange("test"), "onchange", "test")
}

func (suite *AttributesTestSuite) TestOnClick() {
	suite.testAttr(OnClick("test"), "onclick", "test")
}

func (suite *AttributesTestSuite) TestOnContextMenu() {
	suite.testAttr(OnContextMenu("test"), "oncontextmenu", "test")
}

func (suite *AttributesTestSuite) TestOnCopy() {
	suite.testAttr(OnCopy("test"), "oncopy", "test")
}

func (suite *AttributesTestSuite) TestOnCueChange() {
	suite.testAttr(OnCueChange("test"), "oncuechange", "test")
}

func (suite *AttributesTestSuite) TestOnCut() {
	suite.testAttr(OnCut("test"), "oncut", "test")
}

func (suite *AttributesTestSuite) TestOnDoubleClick() {
	suite.testAttr(OnDoubleClick("test"), "ondblclick", "test")
}

func (suite *AttributesTestSuite) TestOnDrag() {
	suite.testAttr(OnDrag("test"), "ondrag", "test")
}

func (suite *AttributesTestSuite) TestOnDragEnd() {
	suite.testAttr(OnDragEnd("test"), "ondragend", "test")
}

func (suite *AttributesTestSuite) TestOnDragEnter() {
	suite.testAttr(OnDragEnter("test"), "ondragenter", "test")
}

func (suite *AttributesTestSuite) TestOnDragLeave() {
	suite.testAttr(OnDragLeave("test"), "ondragleave", "test")
}

func (suite *AttributesTestSuite) TestOnDragOver() {
	suite.testAttr(OnDragOver("test"), "ondragover", "test")
}

func (suite *AttributesTestSuite) TestOnDragStart() {
	suite.testAttr(OnDragStart("test"), "ondragstart", "test")
}

func (suite *AttributesTestSuite) TestOnDrop() {
	suite.testAttr(OnDrop("test"), "ondrop", "test")
}

func (suite *AttributesTestSuite) TestOnDurationChange() {
	suite.testAttr(OnDurationChange("test"), "ondurationchange", "test")
}

func (suite *AttributesTestSuite) TestOnEmptied() {
	suite.testAttr(OnEmptied("test"), "onemptied", "test")
}

func (suite *AttributesTestSuite) TestOnEnded() {
	suite.testAttr(OnEnded("test"), "onended", "test")
}

func (suite *AttributesTestSuite) TestOnError() {
	suite.testAttr(OnError("test"), "onerror", "test")
}

func (suite *AttributesTestSuite) TestOnFocus() {
	suite.testAttr(OnFocus("test"), "onfocus", "test")
}

func (suite *AttributesTestSuite) TestOnHashChange() {
	suite.testAttr(OnHashChange("test"), "onhashchange", "test")
}

func (suite *AttributesTestSuite) TestOnInput() {
	suite.testAttr(OnInput("test"), "oninput", "test")
}

func (suite *AttributesTestSuite) TestOnInvalid() {
	suite.testAttr(OnInvalid("test"), "oninvalid", "test")
}

func (suite *AttributesTestSuite) TestOnKeyDown() {
	suite.testAttr(OnKeyDown("test"), "onkeydown", "test")
}

func (suite *AttributesTestSuite) TestOnKeyPress() {
	suite.testAttr(OnKeyPress("test"), "onkeypress", "test")
}

func (suite *AttributesTestSuite) TestOnKeyUp() {
	suite.testAttr(OnKeyUp("test"), "onkeyup", "test")
}

func (suite *AttributesTestSuite) TestOnLoad() {
	suite.testAttr(OnLoad("test"), "onload", "test")
}

func (suite *AttributesTestSuite) TestOnLoadedData() {
	suite.testAttr(OnLoadedData("test"), "onloadeddata", "test")
}

func (suite *AttributesTestSuite) TestOnLoadedMetaData() {
	suite.testAttr(OnLoadedMetaData("test"), "onloadedmetadata", "test")
}

func (suite *AttributesTestSuite) TestOnLoadStart() {
	suite.testAttr(OnLoadStart("test"), "onloadstart", "test")
}

func (suite *AttributesTestSuite) TestOnMouseDown() {
	suite.testAttr(OnMouseDown("test"), "onmousedown", "test")
}

func (suite *AttributesTestSuite) TestOnMouseMove() {
	suite.testAttr(OnMouseMove("test"), "onmousemove", "test")
}

func (suite *AttributesTestSuite) TestOnMouseOut() {
	suite.testAttr(OnMouseOut("test"), "onmouseout", "test")
}

func (suite *AttributesTestSuite) TestOnMouseOver() {
	suite.testAttr(OnMouseOver("test"), "onmouseover", "test")
}

func (suite *AttributesTestSuite) TestOnMouseUp() {
	suite.testAttr(OnMouseUp("test"), "onmouseup", "test")
}

func (suite *AttributesTestSuite) TestOnOffline() {
	suite.testAttr(OnOffline("test"), "onoffline", "test")
}

func (suite *AttributesTestSuite) TestOnOnline() {
	suite.testAttr(OnOnline("test"), "ononline", "test")
}

func (suite *AttributesTestSuite) TestOnPageHide() {
	suite.testAttr(OnPageHide("test"), "onpagehide", "test")
}

func (suite *AttributesTestSuite) TestOnPageShow() {
	suite.testAttr(OnPageShow("test"), "onpageshow", "test")
}

func (suite *AttributesTestSuite) TestOnPaste() {
	suite.testAttr(OnPaste("test"), "onpaste", "test")
}

func (suite *AttributesTestSuite) TestOnPause() {
	suite.testAttr(OnPause("test"), "onpause", "test")
}

func (suite *AttributesTestSuite) TestOnPlay() {
	suite.testAttr(OnPlay("test"), "onplay", "test")
}

func (suite *AttributesTestSuite) TestOnPlaying() {
	suite.testAttr(OnPlaying("test"), "onplaying", "test")
}

func (suite *AttributesTestSuite) TestOnPopState() {
	suite.testAttr(OnPopState("test"), "onpopstate", "test")
}

func (suite *AttributesTestSuite) TestOnProgress() {
	suite.testAttr(OnProgress("test"), "onprogress", "test")
}

func (suite *AttributesTestSuite) TestOnRateChange() {
	suite.testAttr(OnRateChange("test"), "onratechange", "test")
}

func (suite *AttributesTestSuite) TestOnReset() {
	suite.testAttr(OnReset("test"), "onreset", "test")
}

func (suite *AttributesTestSuite) TestOnResize() {
	suite.testAttr(OnResize("test"), "onresize", "test")
}

func (suite *AttributesTestSuite) TestOnScroll() {
	suite.testAttr(OnScroll("test"), "onscroll", "test")
}

func (suite *AttributesTestSuite) TestOnSearch() {
	suite.testAttr(OnSearch("test"), "onsearch", "test")
}

func (suite *AttributesTestSuite) TestOnSeeked() {
	suite.testAttr(OnSeeked("test"), "onseeked", "test")
}

func (suite *AttributesTestSuite) TestOnSeeking() {
	suite.testAttr(OnSeeking("test"), "onseeking", "test")
}

func (suite *AttributesTestSuite) TestOnSelect() {
	suite.testAttr(OnSelect("test"), "onselect", "test")
}

func (suite *AttributesTestSuite) TestOnStalled() {
	suite.testAttr(OnStalled("test"), "onstalled", "test")
}

func (suite *AttributesTestSuite) TestOnStorage() {
	suite.testAttr(OnStorage("test"), "onstorage", "test")
}

func (suite *AttributesTestSuite) TestOnSubmit() {
	suite.testAttr(OnSubmit("test"), "onsubmit", "test")
}

func (suite *AttributesTestSuite) TestOnSuspend() {
	suite.testAttr(OnSuspend("test"), "onsuspend", "test")
}

func (suite *AttributesTestSuite) TestOnTimeUpdate() {
	suite.testAttr(OnTimeUpdate("test"), "ontimeupdate", "test")
}

func (suite *AttributesTestSuite) TestOnToggle() {
	suite.testAttr(OnToggle("test"), "ontoggle", "test")
}

func (suite *AttributesTestSuite) TestOnUnload() {
	suite.testAttr(OnUnload("test"), "onunload", "test")
}

func (suite *AttributesTestSuite) TestOnVolumeChange() {
	suite.testAttr(OnVolumeChange("test"), "onvolumechange", "test")
}

func (suite *AttributesTestSuite) TestOnWaiting() {
	suite.testAttr(OnWaiting("test"), "onwaiting", "test")
}

func (suite *AttributesTestSuite) TestOnWheel() {
	suite.testAttr(OnWheel("test"), "onwheel", "test")
}

func (suite *AttributesTestSuite) TestOpen() {
	suite.testFlag(Open(), "open")
}

func (suite *AttributesTestSuite) TestOptimum() {
	suite.testAttr(Optimum(10), "optimum", "10")
}

func (suite *AttributesTestSuite) TestPattern() {
	suite.testAttr(Pattern("test"), "pattern", "test")
}

func (suite *AttributesTestSuite) TestPing() {
	suite.testAttr(Ping("test"), "ping", "test")
}

func (suite *AttributesTestSuite) TestPingMulti() {
	suite.testAttr(Ping("testA", "testB"), "ping", "testA testB")
	suite.testAttr(Ping("testC"), "ping", "testA testB testC")
}

func (suite *AttributesTestSuite) TestPlaceholder() {
	suite.testAttr(Placeholder("test"), "placeholder", "test")
}

func (suite *AttributesTestSuite) TestPoster() {
	suite.testAttr(Poster("test"), "poster", "test")
}

func (suite *AttributesTestSuite) TestPreload() {
	suite.testAttr(Preload("test"), "preload", "test")
}

func (suite *AttributesTestSuite) TestReadOnly() {
	suite.testFlag(ReadOnly(), "readonly")
}

func (suite *AttributesTestSuite) TestReferrerPolicy() {
	suite.testAttr(ReferrerPolicy("test"), "referrerpolicy", "test")
}

func (suite *AttributesTestSuite) TestRel() {
	suite.testAttr(Rel("test"), "rel", "test")
}

func (suite *AttributesTestSuite) TestRequired() {
	suite.testFlag(Required(), "required")
}

func (suite *AttributesTestSuite) TestReversed() {
	suite.testFlag(Reversed(), "reversed")
}

func (suite *AttributesTestSuite) TestRows() {
	suite.testAttr(Rows(10), "rows", "10")
}

func (suite *AttributesTestSuite) TestRowSpan() {
	suite.testAttr(RowSpan(10), "rowspan", "10")
}

func (suite *AttributesTestSuite) TestSandbox() {
	suite.testFlag(Sandbox(), "sandbox")
}

func (suite *AttributesTestSuite) TestSandboxList() {
	suite.testAttr(Sandbox("testA", "testB"), "sandbox", "testA testB")
}

func (suite *AttributesTestSuite) TestScope() {
	suite.testAttr(Scope("test"), "scope", "test")
}

func (suite *AttributesTestSuite) TestSelected() {
	suite.testFlag(Selected(), "selected")
}

func (suite *AttributesTestSuite) TestShape() {
	suite.testAttr(Shape("test"), "shape", "test")
}

func (suite *AttributesTestSuite) TestSize() {
	suite.testAttr(Size(10), "size", "10")
}

func (suite *AttributesTestSuite) TestSizes() {
	suite.testAttr(Sizes("test"), "sizes", "test")
}

func (suite *AttributesTestSuite) TestSpan() {
	suite.testAttr(SpanAttr(10), "span", "10")
}

func (suite *AttributesTestSuite) TestSpellcheckTrue() {
	suite.testAttr(Spellcheck(true), "spellcheck", "true")
}

func (suite *AttributesTestSuite) TestSpellcheckFalse() {
	suite.testAttr(Spellcheck(false), "spellcheck", "false")
}

func (suite *AttributesTestSuite) TestSrc() {
	suite.testAttr(Src("test"), "src", "test")
}

func (suite *AttributesTestSuite) TestSrcDoc() {
	suite.testAttr(SrcDoc(P()()), "srcdoc", "<p></p>")
}

func (suite *AttributesTestSuite) TestSrcLang() {
	suite.testAttr(SrcLang("test"), "srclang", "test")
}

func (suite *AttributesTestSuite) TestSrcSet() {
	suite.testAttr(SrcSet("test"), "srcset", "test")
}

func (suite *AttributesTestSuite) TestStart() {
	suite.testAttr(Start(10), "start", "10")
}

func (suite *AttributesTestSuite) TestStepInt() {
	suite.testAttr(Step(10), "step", "10")
}

func (suite *AttributesTestSuite) TestStepFloat() {
	suite.testAttr(Step(10.0), "step", "10.000000")
}

func (suite *AttributesTestSuite) TestStepString() {
	suite.testAttr(Step("test"), "step", "test")
}

func (suite *AttributesTestSuite) TestStyle() {
	suite.testAttr(StyleAttr("test"), "style", "test")
}

func (suite *AttributesTestSuite) TestTabIndex() {
	suite.testAttr(TabIndex(10), "tabindex", "10")
}

func (suite *AttributesTestSuite) TestTarget() {
	suite.testAttr(Target("test"), "target", "test")
}

func (suite *AttributesTestSuite) TestTitle() {
	suite.testAttr(TitleAttr("test"), "title", "test")
}

func (suite *AttributesTestSuite) TestTranslateYes() {
	suite.testAttr(Translate(true), "translate", "yes")
}

func (suite *AttributesTestSuite) TestTranslateNo() {
	suite.testAttr(Translate(false), "translate", "no")
}

func (suite *AttributesTestSuite) TestType() {
	suite.testAttr(Type("test"), "type", "test")
}

func (suite *AttributesTestSuite) TestTypeMustMatchTrue() {
	suite.testAttr(TypeMustMatch(true), "typemustmatch", "true")
}

func (suite *AttributesTestSuite) TestTypeMustMatchFalse() {
	suite.testAttr(TypeMustMatch(false), "typemustmatch", "false")
}

func (suite *AttributesTestSuite) TestUseMap() {
	suite.testAttr(UseMap("test"), "usemap", "test")
}

func (suite *AttributesTestSuite) TestValueInt() {
	suite.testAttr(Value(10), "value", "10")
}

func (suite *AttributesTestSuite) TestValueStr() {
	suite.testAttr(Value("test"), "value", "test")
}

func (suite *AttributesTestSuite) TestWidth() {
	suite.testAttr(Width(10), "width", "10")
}

func (suite *AttributesTestSuite) TestWrap() {
	suite.testAttr(Wrap("test"), "wrap", "test")
}

func (suite *AttributesTestSuite) TestXMLNS() {
	suite.testAttr(XMLNS("test"), "xmlns", "test")
}

func (suite *AttributesTestSuite) testAttr(attribute types.Attribute, key, value string) {
	suite.testAttrAndFlag(attribute, map[string]string{key: value}, nil)
}

func (suite *AttributesTestSuite) testFlag(attribute types.Attribute, key string) {
	suite.testAttrAndFlag(attribute, nil, []string{key})
}

func (suite *AttributesTestSuite) testAttrAndFlag(attribute types.Attribute, expectedAttributes map[string]string, expectedFlags []string) {
	attribute(suite.state.attributes, &suite.state.flags, nil)

	if expectedAttributes != nil {
		suite.Assert().Equal(expectedAttributes, suite.state.attributes)
	} else {
		suite.Assert().Empty(suite.state.attributes)
	}

	if expectedFlags != nil {
		suite.Assert().Equal(expectedFlags, suite.state.flags)
	} else {
		suite.Assert().Empty(suite.state.flags)
	}
}

type attrState struct {
	attributes map[string]string
	flags      []string
}

func (s *attrState) reset() {
	s.attributes = make(map[string]string)
	s.flags = []string{}
}
