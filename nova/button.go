package nova

// Button represents a button in the UI
type Button struct {
	Item             string `xml:"item,attr"`
	ScreenID         string `xml:"ScreenID"`
	Font             string `xml:"Font"`
	RelativePosition string `xml:"RelativePosition"`
	LocationX        string `xml:"Location>X"`
	LocationY        string `xml:"Location>Y"`
	SizeX            string `xml:"Size>CX"`
	SizeY            string `xml:"Size>CY"`
	Text             string `xml:"Text"`
	TextColorR       string `xml:"TextColor>R"`
	TextColorG       string `xml:"TextColor>G"`
	TextColorB       string `xml:"TextColor>B"`
	Template         string `xml:"Template"`
}

/*
<Button item="IDW_ModButton">
		<ScreenID>IDW_ModButton</ScreenID>
		<!--<Font>3</Font>-->
		<RelativePosition>true</RelativePosition>
		<Location>
			<X>10</X>
			<Y>135</Y>
		</Location>
		<Size>
			<CX>10</CX>
			<CY>10</CY>
		</Size>
		<Text></Text>
		<TextColor>
			<R>255</R>
			<G>255</G>
			<B>255</B>
		</TextColor>
		<Template>BDT_Normal</Template>
		<!--<SoundPressed/>-->
		<!--<SoundUp/>-->
		<!--<SoundFlyby/>-->
  </Button>
*/
