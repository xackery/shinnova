package nxml

// Label represents a label in the UI
type Label struct {
	Item             string  `xml:"item,attr"`
	ScreenID         *string `xml:"ScreenID"`
	Font             *string `xml:"Font"`
	RelativePosition *string `xml:"RelativePosition"`
	LocationX        *string `xml:"Location>X"`
	LocationY        *string `xml:"Location>Y"`
	SizeX            *string `xml:"Size>CX"`
	SizeY            *string `xml:"Size>CY"`
	Text             *string `xml:"Text"`
	TextColorR       *string `xml:"TextColor>R"`
	TextColorG       *string `xml:"TextColor>G"`
	TextColorB       *string `xml:"TextColor>B"`
	NoWrap           *string `xml:"NoWrap"`
	AlignCenter      *string `xml:"AlignCenter"`
	AlignRight       *string `xml:"AlignRight"`
}

/*
<Label item = "IDW_ModButtonLabel">
	<ScreenID>IDW_ModButtonLabel</ScreenID>
	<RelativePosition>true</RelativePosition>
	<Location>
		<X>10</X>
		<Y>60</Y>
	</Location>
	<Size>
		<CX>40</CX>
		<CY>20</CY>
	</Size>
	<Text></Text>
	<TextColor>
		<R>255</R>
		<G>255</G>
		<B>255</B>
	</TextColor>
	<NoWrap>true</NoWrap>
	<AlignCenter>true</AlignCenter>
	<AlignRight>false</AlignRight>
  </Label>
*/
