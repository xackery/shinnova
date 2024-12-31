package nxml

// TextureInfo represents a texture info in the UI
type TextureInfo struct {
	Item  string  `xml:"item,attr"`
	SizeX *string `xml:"Size>CX"`
	SizeY *string `xml:"Size>CY"`
}

/*
<TextureInfo item="AttackIndicator.tga">
		<Size>
			<CX>128</CX>
			<CY>32</CY>
		</Size>
	</TextureInfo> */
