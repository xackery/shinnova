package nova

type TabBox struct {
	Item               string   `xml:"item,attr"`
	ScreenID           string   `xml:"ScreenID"`
	RelativePosition   string   `xml:"RelativePosition"`
	AutoStretch        string   `xml:"AutoStretch"`
	TopAnchorToTop     string   `xml:"TopAnchorToTop"`
	BottomAnchorToTop  string   `xml:"BottomAnchorToTop"`
	LeftAnchorToLeft   string   `xml:"LeftAnchorToLeft"`
	RightAnchorToLeft  string   `xml:"RightAnchorToLeft"`
	TabBorderTemplate  string   `xml:"TabBorderTemplate"`
	PageBorderTemplate string   `xml:"PageBorderTemplate"`
	Pages              []string `xml:"Pages"`
	LeftAnchorOffset   string   `xml:"LeftAnchorOffset"`
	RightAnchorOffset  string   `xml:"RightAnchorOffset"`
	TopAnchorOffset    string   `xml:"TopAnchorOffset"`
	BottomAnchorOffset string   `xml:"BottomAnchorOffset"`
}

/*
<TabBox item="IDW_ItemDescriptionTabBox">
    <ScreenID>ItemDescriptionTabBox</ScreenID>
    <!--<Font>3</Font>-->
    <RelativePosition>true</RelativePosition>
    <AutoStretch>true</AutoStretch>
    <TopAnchorToTop>true</TopAnchorToTop>
    <BottomAnchorToTop>false</BottomAnchorToTop>
    <LeftAnchorToLeft>true</LeftAnchorToLeft>
    <RightAnchorToLeft>false</RightAnchorToLeft>
    <TabBorderTemplate>FT_DefTabBorder</TabBorderTemplate>
    <PageBorderTemplate>FT_DefPageBorder</PageBorderTemplate>
    <Pages>IDW_ItemDescriptionTab</Pages>
    <Pages>IDW_ItemLoreTab</Pages>
    <LeftAnchorOffset>2</LeftAnchorOffset>
    <RightAnchorOffset>-2</RightAnchorOffset>
    <TopAnchorOffset>2</TopAnchorOffset>
    <BottomAnchorOffset>-9</BottomAnchorOffset>
  </TabBox>
*/
