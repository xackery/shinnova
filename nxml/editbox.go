package nxml

type Editbox struct {
	Item             string `xml:"item,attr"`
	ScreenID         string `xml:"ScreenID"`
	Font             string `xml:"Font"`
	DrawTemplate     string `xml:"DrawTemplate"`
	StyleBorder      string `xml:"Style_Border"`
	RelativePosition string `xml:"RelativePosition"`
	LocationX        string `xml:"Location>X"`
	LocationY        string `xml:"Location>Y"`
	SizeCX           string `xml:"Size>CX"`
	SizeCY           string `xml:"Size>CY"`
	TextColorR       string `xml:"TextColor>R"`
	TextColorG       string `xml:"TextColor>G"`
	TextColorB       string `xml:"TextColor>B"`
}

/*
<Editbox item = "MACRO_IDEdit">
		<ScreenID>IDEdit</ScreenID>
		<Font>4</Font>
		<DrawTemplate>WDT_EQLS_Def_Bordered</DrawTemplate>
		<Style_Border>true</Style_Border>
		<RelativePosition>true</RelativePosition>
		<Location>
			<X>80</X>
			<Y>220</Y>
		</Location>
		<Size>
			<CX>300</CX>
			<CY>35</CY>
		</Size>
		<TextColor>
				<R>255</R>
				<G>255</G>
				<B>255</B>
		</TextColor>
	</Editbox>
*/
