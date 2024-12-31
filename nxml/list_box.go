package nxml

type ListBox struct {
	Item             string `xml:"item,attr"`
	ScreenID         string `xml:"ScreenID"`
	Font             string `xml:"Font"`
	DrawTemplate     string `xml:"DrawTemplate"`
	RelativePosition string `xml:"RelativePosition"`
	LocationX        string `xml:"Location>X"`
	LocationY        string `xml:"Location>Y"`
	SizeCX           string `xml:"Size>CX"`
	SizeCY           string `xml:"Size>CY"`
	StyleBorder      string `xml:"Style_Border"`
	StyleTransparent string `xml:"Style_Transparent"`
	StyleVScroll     string `xml:"Style_VScroll"`
	Columns          []struct {
		Width   string `xml:"Width"`
		Heading string `xml:"Heading"`
	} `xml:"Columns"`
}

/*
<Listbox item = "MACRO_MacroList">
		<ScreenID>MacroList</ScreenID>
		<Font>4</Font>
		<DrawTemplate>WDT_EQLS_Def_Bordered</DrawTemplate>
		<RelativePosition>true</RelativePosition>
		<Location>
			<X>10</X>
			<Y>10</Y>
		</Location>
		<Size>
			<CX>480</CX>
			<CY>200</CY>
		</Size>
		<Style_Border>true</Style_Border>
		<Style_Transparent>false</Style_Transparent>
		<Style_VScroll>true</Style_VScroll>
		<Columns>
			<Width>50</Width>
			<Heading>ID:</Heading>
		</Columns>
		<Columns>
			<Width>150</Width>
			<Heading>Name:</Heading>
		</Columns>
		<Columns>
			<Width>3000</Width>
			<Heading>Macro:</Heading>
		</Columns>
	</Listbox>
*/
