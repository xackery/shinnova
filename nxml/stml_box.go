package nxml

// STMLBox represents a box in the UI
type STMLBox struct {
	Item             string `xml:"item,attr"`
	ScreenID         string `xml:"ScreenID"`
	DrawTemplate     string `xml:"DrawTemplate"`
	LocationX        string `xml:"Location>X"`
	LocationY        string `xml:"Location>Y"`
	SizeX            string `xml:"Size>CX"`
	SizeY            string `xml:"Size>CY"`
	StyleBorder      string `xml:"Style_Border"`
	StyleTransparent string `xml:"Style_Transparent"`
	Text             string `xml:"Text"`
}

/*
<STMLbox item="IDW_Appearance_Socket_Description">
    <ScreenID>IDW_Appearance_Socket_Description</ScreenID>
    <DrawTemplate>WDT_Inner</DrawTemplate>
	<Location>
		<X>58</X>
		<Y>12</Y>
	</Location>
	<Size>
		<CX>700</CX>
		<CY>20</CY>
	</Size>
	<Style_Border>false</Style_Border>
	<Style_Transparent>true</Style_Transparent>
	<Text>Socket Description</Text>
  </STMLbox>
*/
