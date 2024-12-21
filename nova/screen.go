package nova

// STMLBox represents a box in the UI
type Screen struct {
	Item             string   `xml:"item,attr"`
	ScreenID         string   `xml:"ScreenID"`
	Font             string   `xml:"Font"`
	RelativePosition string   `xml:"RelativePosition"`
	LocationX        string   `xml:"Location>X"`
	LocationY        string   `xml:"Location>Y"`
	SizeX            string   `xml:"Size>CX"`
	SizeY            string   `xml:"Size>CY"`
	StyleVScroll     string   `xml:"Style_VScroll"`
	StyleHScroll     string   `xml:"Style_HScroll"`
	StyleTransparent string   `xml:"Style_Transparent"`
	TooltipReference string   `xml:"TooltipReference"`
	DrawTemplate     string   `xml:"DrawTemplate"`
	StyleTitlebar    string   `xml:"Style_Titlebar"`
	StyleClosebox    string   `xml:"Style_Closebox"`
	StyleMinimizebox string   `xml:"Style_Minimizebox"`
	StyleBorder      string   `xml:"Style_Border"`
	StyleSizable     string   `xml:"Style_Sizable"`
	StyleQmarkbox    string   `xml:"Style_Qmarkbox"`
	Pieces           []string `xml:"Pieces"`
}

/*
<Screen item="ItemDisplayWindow">
    <!--<ScreenID/>-->
    <!--<Font/>-->
    <RelativePosition>false</RelativePosition>
    <Location>
      <X>0</X>
      <Y>80</Y>
    </Location>
    <Size>
      <CX>400</CX>
      <CY>230</CY>
    </Size>
    <Style_VScroll>false</Style_VScroll>
    <Style_HScroll>false</Style_HScroll>
    <Style_Transparent>false</Style_Transparent>
    <TooltipReference>This is an Item Display window</TooltipReference>
    <DrawTemplate>WDT_Filigree</DrawTemplate>
    <Style_Titlebar>true</Style_Titlebar>
    <Style_Closebox>true</Style_Closebox>
    <Style_Minimizebox>true</Style_Minimizebox>
    <Style_Border>true</Style_Border>
    <Style_Sizable>true</Style_Sizable>
    <Style_Qmarkbox>true</Style_Qmarkbox>
    <Pieces>IDW_ItemDescriptionTabBox</Pieces>
    <Pieces>IDW_ItemDescriptionTabBox</Pieces>
    <Pieces>IDW_ItemDescriptionTabBox</Pieces>
    <Pieces>IDW_ItemDescriptionTabBox</Pieces>
  </Screen>
*/
