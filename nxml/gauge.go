package nxml

// Gauge represents a gauge in the UI
type Gauge struct {
	Item               string             `xml:"item,attr"`
	ScreenID           *string            `xml:"ScreenID"`
	RelativePosition   *string            `xml:"RelativePosition"`
	GaugeOffsetX       *string            `xml:"GaugeOffsetX"`
	GaugeOffsetY       *string            `xml:"GaugeOffsetY"`
	FillTintR          *string            `xml:"FillTint>R"`
	FillTintG          *string            `xml:"FillTint>G"`
	FillTintB          *string            `xml:"FillTint>B"`
	EQType             *string            `xml:"EQType"`
	GaugeDrawTemplate  *GaugeDrawTemplate `xml:"GaugeDrawTemplate"`
	AutoStretch        *string            `xml:"AutoStretch"`
	TopAnchorOffset    *string            `xml:"TopAnchorOffset"`
	BottomAnchorOffset *string            `xml:"BottomAnchorOffset"`
	LeftAnchorOffset   *string            `xml:"LeftAnchorOffset"`
	RightAnchorOffset  *string            `xml:"RightAnchorOffset"`
	RightAnchorToLeft  *string            `xml:"RightAnchorToLeft"`
}

type GaugeDrawTemplate struct {
	Background  *string `xml:"Background"`
	Fill        *string `xml:"Fill"`
	Lines       *string `xml:"Lines"`
	EndCapLeft  *string `xml:"EndCapLeft"`
	EndCapRight *string `xml:"EndCapRight"`
}

/*
<Gauge item="Player_CombatTimer">
		<ScreenID>Player_CombatTimer</ScreenID>
		<RelativePosition>true</RelativePosition>
		<GaugeOffsetY>0</GaugeOffsetY>
		<FillTint>
			<R>169</R>
			<G>169</G>
			<B>169</B>
		</FillTint>
		<EQType>29</EQType>
		<GaugeDrawTemplate>
			<Background>A_GaugeBackground</Background>
			<Fill>A_GaugeFill</Fill>
			<Lines>A_GaugeLines</Lines>
			<EndCapLeft>A_GaugeEndCapLeft</EndCapLeft>
			<EndCapRight>A_GaugeEndCapRight</EndCapRight>
		</GaugeDrawTemplate>
		<AutoStretch>true</AutoStretch>
		<TopAnchorOffset>82</TopAnchorOffset>
		<BottomAnchorOffset>92</BottomAnchorOffset>
		<LeftAnchorOffset>2</LeftAnchorOffset>
		<RightAnchorOffset>0</RightAnchorOffset>
		<RightAnchorToLeft>false</RightAnchorToLeft>
	</Gauge> */
