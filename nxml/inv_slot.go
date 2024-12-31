package nxml

// InvSlot represents an inventory slot in the UI
type InvSlot struct {
	Item             string  `xml:"item,attr"`
	ScreenID         *string `xml:"ScreenID"`
	Font             *string `xml:"Font"`
	RelativePosition *string `xml:"RelativePosition"`
	LocationX        *string `xml:"Location>X"`
	LocationY        *string `xml:"Location>Y"`
	SizeCX           *string `xml:"Size>CX"`
	SizeCY           *string `xml:"Size>CY"`
	Background       *string `xml:"Background"`
	EQType           *string `xml:"EQType"`
}

/* <InvSlot item="BW_BankSlot15">
	<ScreenID>BW_BankSlot15</ScreenID>
	<!--<Font>3</Font>-->
	<RelativePosition>true</RelativePosition>
	<Location>
		<X>139</X>
		<Y>151</Y>
	</Location>
	<Size>
		<CX>40</CX>
		<CY>40</CY>
	</Size>
	<Background>A_RecessedBox</Background>
	<EQType>inventory/Bank 15</EQType>
</InvSlot> */
