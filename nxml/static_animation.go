package nxml

// StaticAnimation represents a static animation in the UI
type StaticAnimation struct {
	Item      string  `xml:"item,attr"`
	ScreenID  *string `xml:"ScreenID"`
	LocationX *string `xml:"Location>X"`
	LocationY *string `xml:"Location>Y"`
	SizeCX    *string `xml:"Size>CX"`
	SizeCY    *string `xml:"Size>CY"`
	Animation *string `xml:"Animation"`
}

/*
<StaticAnimation item="CC_LogoAnim">
		<ScreenID>CC_LogoAnim</ScreenID>
		<Location>
			<X>0</X>
			<Y>8</Y>
		</Location>
		<Size>
			<CX>225</CX>
			<CY>100</CY>
		</Size>
		<Animation>CC_EQIcon</Animation>
	</StaticAnimation> */
