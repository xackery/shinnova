package nxml

// Combobox represents a combobox in the UI
type Combobox struct {
	Item             string  `xml:"item,attr"`
	ScreenID         *string `xml:"ScreenID"`
	RelativePosition *string `xml:"RelativePosition"`
	LocationX        *string `xml:"Location>X"`
	LocationY        *string `xml:"Location>Y"`
	SizeCX           *string `xml:"Size>CX"`
	SizeCY           *string `xml:"Size>CY"`
	DrawTemplate     *string `xml:"DrawTemplate"`
	ListHeight       *string `xml:"ListHeight"`
	Button           *string `xml:"Button"`
	StyleBorder      *string `xml:"Style_Border"`
}

/*
<Combobox item = "CC_City_Combo">
		<ScreenID>CC_City_Combo</ScreenID>
		<RelativePosition>true</RelativePosition>
		<Location>
			<X>60</X>
			<Y>35</Y>
		</Location>
		<Size>
			<CX>110</CX>
			<CY>20</CY>
		</Size>
		<DrawTemplate>WDT_Inner</DrawTemplate>
		<ListHeight>80</ListHeight>
		<Button>BDT_Combo</Button>
		<Style_Border>true</Style_Border>
	</Combobox> */
