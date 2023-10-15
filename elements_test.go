package godom_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/suite"
	. "github.com/tbe/godom"
	"github.com/tbe/godom/types"
)

func TestElements(t *testing.T) {
	suite.Run(t, new(ElementsTestSuite))
}

type ElementsTestSuite struct {
	suite.Suite
	buf bytes.Buffer
}

func (suite *ElementsTestSuite) SetupTest() {
	suite.buf.Reset()
}

func (suite *ElementsTestSuite) testElement(element types.Element, expected string) {
	suite.Assert().NoError(element.Render(&suite.buf))
	suite.Assert().Equal(expected, suite.buf.String())

}

func (suite *ElementsTestSuite) TestDoctype() {
	suite.testElement(Doctype(), "<!DOCTYPE html>")
}
func (suite *ElementsTestSuite) TestA() {
	suite.testElement(A()(), "<a></a>")
}

func (suite *ElementsTestSuite) TestAbbr() {
	suite.testElement(Abbr()(), "<abbr></abbr>")
}

func (suite *ElementsTestSuite) TestAddress() {
	suite.testElement(Address()(), "<address></address>")
}

func (suite *ElementsTestSuite) TestArea() {
	suite.testElement(Area()(), "<area></area>")
}

func (suite *ElementsTestSuite) TestArticle() {
	suite.testElement(Article()(), "<article></article>")
}

func (suite *ElementsTestSuite) TestAside() {
	suite.testElement(Aside()(), "<aside></aside>")
}

func (suite *ElementsTestSuite) TestAudio() {
	suite.testElement(Audio()(), "<audio></audio>")
}

func (suite *ElementsTestSuite) TestB() {
	suite.testElement(B()(), "<b></b>")
}

func (suite *ElementsTestSuite) TestBase() {
	suite.testElement(Base(), "<base/>")
}

func (suite *ElementsTestSuite) TestBDI() {
	suite.testElement(BDI()(), "<bdi></bdi>")
}

func (suite *ElementsTestSuite) TestBDO() {
	suite.testElement(BDO()(), "<bdo></bdo>")
}

func (suite *ElementsTestSuite) TestBlockquote() {
	suite.testElement(Blockquote()(), "<blockquote></blockquote>")
}

func (suite *ElementsTestSuite) TestBody() {
	suite.testElement(Body()(), "<body></body>")
}

func (suite *ElementsTestSuite) TestBr() {
	suite.testElement(Br(), "<br/>")
}

func (suite *ElementsTestSuite) TestButton() {
	suite.testElement(Button()(), "<button></button>")
}

func (suite *ElementsTestSuite) TestCanvas() {
	suite.testElement(Canvas()(), "<canvas></canvas>")
}

func (suite *ElementsTestSuite) TestCaption() {
	suite.testElement(Caption()(), "<caption></caption>")
}

func (suite *ElementsTestSuite) TestCite() {
	suite.testElement(Cite()(), "<cite></cite>")
}

func (suite *ElementsTestSuite) TestCode() {
	suite.testElement(Code()(), "<code></code>")
}

func (suite *ElementsTestSuite) TestCol() {
	suite.testElement(Col(), "<col/>")
}

func (suite *ElementsTestSuite) TestColGroup() {
	suite.testElement(ColGroup()(), "<colgroup></colgroup>")
}

func (suite *ElementsTestSuite) TestData() {
	suite.testElement(Data()(), "<data></data>")
}

func (suite *ElementsTestSuite) TestDataList() {
	suite.testElement(DataList()(), "<datalist></datalist>")
}

func (suite *ElementsTestSuite) TestDD() {
	suite.testElement(DD()(), "<dd></dd>")
}

func (suite *ElementsTestSuite) TestDel() {
	suite.testElement(Del()(), "<del></del>")
}

func (suite *ElementsTestSuite) TestDetails() {
	suite.testElement(Details()(), "<details></details>")
}

func (suite *ElementsTestSuite) TestDfn() {
	suite.testElement(Dfn()(), "<dfn></dfn>")
}

func (suite *ElementsTestSuite) TestDialog() {
	suite.testElement(Dialog()(), "<dialog></dialog>")
}

func (suite *ElementsTestSuite) TestDiv() {
	suite.testElement(Div()(), "<div></div>")
}

func (suite *ElementsTestSuite) TestDL() {
	suite.testElement(DL()(), "<dl></dl>")
}

func (suite *ElementsTestSuite) TestDT() {
	suite.testElement(DT()(), "<dt></dt>")
}

func (suite *ElementsTestSuite) TestEm() {
	suite.testElement(Em()(), "<em></em>")
}

func (suite *ElementsTestSuite) TestEmbed() {
	suite.testElement(Embed(), "<embed/>")
}

func (suite *ElementsTestSuite) TestFieldSet() {
	suite.testElement(FieldSet()(), "<fieldset></fieldset>")
}

func (suite *ElementsTestSuite) TestFigCaption() {
	suite.testElement(FigCaption()(), "<figcaption></figcaption>")
}

func (suite *ElementsTestSuite) TestFigure() {
	suite.testElement(Figure()(), "<figure></figure>")
}

func (suite *ElementsTestSuite) TestFooter() {
	suite.testElement(Footer()(), "<footer></footer>")
}

func (suite *ElementsTestSuite) TestForm() {
	suite.testElement(Form()(), "<form></form>")
}

func (suite *ElementsTestSuite) TestH1() {
	suite.testElement(H1()(), "<h1></h1>")
}

func (suite *ElementsTestSuite) TestH2() {
	suite.testElement(H2()(), "<h2></h2>")
}

func (suite *ElementsTestSuite) TestH3() {
	suite.testElement(H3()(), "<h3></h3>")
}

func (suite *ElementsTestSuite) TestH4() {
	suite.testElement(H4()(), "<h4></h4>")
}

func (suite *ElementsTestSuite) TestH5() {
	suite.testElement(H5()(), "<h5></h5>")
}

func (suite *ElementsTestSuite) TestH6() {
	suite.testElement(H6()(), "<h6></h6>")
}

func (suite *ElementsTestSuite) TestHead() {
	suite.testElement(Head()(), "<head></head>")
}

func (suite *ElementsTestSuite) TestHeader() {
	suite.testElement(Header()(), "<header></header>")
}

func (suite *ElementsTestSuite) TestHR() {
	suite.testElement(HR(), "<hr/>")
}

func (suite *ElementsTestSuite) TestHTML() {
	suite.testElement(HTML()(), "<html></html>")
}

func (suite *ElementsTestSuite) TestI() {
	suite.testElement(I()(), "<i></i>")
}

func (suite *ElementsTestSuite) TestIFrame() {
	suite.testElement(IFrame(), "<iframe></iframe>")
}

func (suite *ElementsTestSuite) TestImg() {
	suite.testElement(Img(), "<img/>")
}

func (suite *ElementsTestSuite) TestInput() {
	suite.testElement(Input(), "<input/>")
}

func (suite *ElementsTestSuite) TestIns() {
	suite.testElement(Ins()(), "<ins></ins>")
}

func (suite *ElementsTestSuite) TestKbd() {
	suite.testElement(Kbd()(), "<kbd></kbd>")
}

func (suite *ElementsTestSuite) TestLabel() {
	suite.testElement(Label()(), "<label></label>")
}

func (suite *ElementsTestSuite) TestLegend() {
	suite.testElement(Legend()(), "<legend></legend>")
}

func (suite *ElementsTestSuite) TestLi() {
	suite.testElement(Li()(), "<li></li>")
}

func (suite *ElementsTestSuite) TestLink() {
	suite.testElement(Link(), "<link/>")
}

func (suite *ElementsTestSuite) TestMain() {
	suite.testElement(Main()(), "<main></main>")
}

func (suite *ElementsTestSuite) TestMap() {
	suite.testElement(Map()(), "<map></map>")
}

func (suite *ElementsTestSuite) TestMark() {
	suite.testElement(Mark()(), "<mark></mark>")
}

func (suite *ElementsTestSuite) TestMeta() {
	suite.testElement(Meta(), "<meta/>")
}

func (suite *ElementsTestSuite) TestMeter() {
	suite.testElement(Meter()(), "<meter></meter>")
}

func (suite *ElementsTestSuite) TestNav() {
	suite.testElement(Nav()(), "<nav></nav>")
}

func (suite *ElementsTestSuite) TestNoScript() {
	suite.testElement(NoScript()(), "<noscript></noscript>")
}

func (suite *ElementsTestSuite) TestObject() {
	suite.testElement(Object()(), "<object></object>")
}

func (suite *ElementsTestSuite) TestOL() {
	suite.testElement(OL()(), "<ol></ol>")
}

func (suite *ElementsTestSuite) TestOptGroup() {
	suite.testElement(OptGroup()(), "<optgroup></optgroup>")
}

func (suite *ElementsTestSuite) TestOption() {
	suite.testElement(Option()(), "<option></option>")
}

func (suite *ElementsTestSuite) TestOutput() {
	suite.testElement(Output()(), "<output></output>")
}

func (suite *ElementsTestSuite) TestP() {
	suite.testElement(P()(), "<p></p>")
}

func (suite *ElementsTestSuite) TestParam() {
	suite.testElement(Param(), "<param/>")
}

func (suite *ElementsTestSuite) TestPicture() {
	suite.testElement(Picture()(), "<picture></picture>")
}

func (suite *ElementsTestSuite) TestPre() {
	suite.testElement(Pre()(), "<pre></pre>")
}

func (suite *ElementsTestSuite) TestProgress() {
	suite.testElement(Progress()(), "<progress></progress>")
}

func (suite *ElementsTestSuite) TestQ() {
	suite.testElement(Q()(), "<q></q>")
}

func (suite *ElementsTestSuite) TestRP() {
	suite.testElement(RP()(), "<rp></rp>")
}

func (suite *ElementsTestSuite) TestRT() {
	suite.testElement(RT()(), "<rt></rt>")
}

func (suite *ElementsTestSuite) TestRuby() {
	suite.testElement(Ruby()(), "<ruby></ruby>")
}

func (suite *ElementsTestSuite) TestS() {
	suite.testElement(S()(), "<s></s>")
}

func (suite *ElementsTestSuite) TestSamp() {
	suite.testElement(Samp()(), "<samp></samp>")
}

func (suite *ElementsTestSuite) TestScript() {
	suite.testElement(Script()(), "<script></script>")
}

func (suite *ElementsTestSuite) TestSection() {
	suite.testElement(Section()(), "<section></section>")
}

func (suite *ElementsTestSuite) TestSelect() {
	suite.testElement(Select()(), "<select></select>")
}

func (suite *ElementsTestSuite) TestSmall() {
	suite.testElement(Small()(), "<small></small>")
}

func (suite *ElementsTestSuite) TestSource() {
	suite.testElement(Source(), "<source/>")
}

func (suite *ElementsTestSuite) TestSpan() {
	suite.testElement(Span()(), "<span></span>")
}

func (suite *ElementsTestSuite) TestStrong() {
	suite.testElement(Strong()(), "<strong></strong>")
}

func (suite *ElementsTestSuite) TestStyle() {
	suite.testElement(Style()(), "<style></style>")
}

func (suite *ElementsTestSuite) TestSub() {
	suite.testElement(Sub()(), "<sub></sub>")
}

func (suite *ElementsTestSuite) TestSummary() {
	suite.testElement(Summary()(), "<summary></summary>")
}

func (suite *ElementsTestSuite) TestSup() {
	suite.testElement(Sup()(), "<sup></sup>")
}

func (suite *ElementsTestSuite) TestSVG() {
	suite.testElement(SVG()(), "<svg></svg>")
}

func (suite *ElementsTestSuite) TestTable() {
	suite.testElement(Table()(), "<table></table>")
}

func (suite *ElementsTestSuite) TestTBody() {
	suite.testElement(TBody()(), "<tbody></tbody>")
}

func (suite *ElementsTestSuite) TestTD() {
	suite.testElement(TD()(), "<td></td>")
}

func (suite *ElementsTestSuite) TestTemplate() {
	suite.testElement(Template()(), "<template></template>")
}

func (suite *ElementsTestSuite) TestTextArea() {
	suite.testElement(TextArea()(), "<textarea></textarea>")
}

func (suite *ElementsTestSuite) TestTFoot() {
	suite.testElement(TFoot()(), "<tfoot></tfoot>")
}

func (suite *ElementsTestSuite) TestTH() {
	suite.testElement(TH()(), "<th></th>")
}

func (suite *ElementsTestSuite) TestTHead() {
	suite.testElement(THead()(), "<thead></thead>")
}

func (suite *ElementsTestSuite) TestTime() {
	suite.testElement(Time()(), "<time></time>")
}

func (suite *ElementsTestSuite) TestTitle() {
	suite.testElement(Title()(), "<title></title>")
}

func (suite *ElementsTestSuite) TestTR() {
	suite.testElement(TR()(), "<tr></tr>")
}

func (suite *ElementsTestSuite) TestTrack() {
	suite.testElement(Track(), "<track/>")
}

func (suite *ElementsTestSuite) TestU() {
	suite.testElement(U()(), "<u></u>")
}

func (suite *ElementsTestSuite) TestUL() {
	suite.testElement(UL()(), "<ul></ul>")
}

func (suite *ElementsTestSuite) TestVar() {
	suite.testElement(Var()(), "<var></var>")
}

func (suite *ElementsTestSuite) TestVideo() {
	suite.testElement(Video()(), "<video></video>")
}

func (suite *ElementsTestSuite) TestWBr() {
	suite.testElement(WBr(), "<wbr/>")
}

func (suite *ElementsTestSuite) TestDocument() {
	suite.testElement(Group(Div()(), Div()()), "<div></div><div></div>")
}

func (suite *ElementsTestSuite) TestContent() {
	suite.testElement(Content("some <content>"), "some &lt;content&gt;")
}
