package nxml

// Button represents a button in the UI
type Button struct {
	Item               string                    `xml:"item,attr"`
	ScreenID           *string                   `xml:"ScreenID"`
	Font               *string                   `xml:"Font"`
	RadioGroup         *string                   `xml:"RadioGroup"`
	RelativePosition   *string                   `xml:"RelativePosition"`
	LocationX          *string                   `xml:"Location>X"`
	LocationY          *string                   `xml:"Location>Y"`
	SizeX              *string                   `xml:"Size>CX"`
	SizeY              *string                   `xml:"Size>CY"`
	StyleTransparent   *string                   `xml:"Style_Transparent"`
	TooltipReference   *string                   `xml:"TooltipReference"`
	StyleCheckbox      *string                   `xml:"Style_Checkbox"`
	Text               *string                   `xml:"Text"`
	TextColorR         *string                   `xml:"TextColor>R"`
	TextColorG         *string                   `xml:"TextColor>G"`
	TextColorB         *string                   `xml:"TextColor>B"`
	Template           *string                   `xml:"Template"`
	DecalOffsetX       *string                   `xml:"DecalOffset>X"`
	DecalOffsetY       *string                   `xml:"DecalOffset>Y"`
	DecalSizeCX        *string                   `xml:"DecalSize>CX"`
	DecalSizeCY        *string                   `xml:"DecalSize>CY"`
	ButtonDrawTemplate *ButtonButtonDrawTemplate `xml:"ButtonDrawTemplate"`
}

type ButtonButtonDrawTemplate struct {
	Normal            *string `xml:"Normal"`
	Pressed           *string `xml:"Pressed"`
	Flyby             *string `xml:"Flyby"`
	Disabled          *string `xml:"Disabled"`
	PressedFlyby      *string `xml:"PressedFlyby"`
	NormalDecal       *string `xml:"NormalDecal"`
	PressedDecal      *string `xml:"PressedDecal"`
	FlybyDecal        *string `xml:"FlybyDecal"`
	DisabledDecal     *string `xml:"DisabledDecal"`
	PressedFlybyDecal *string `xml:"PressedFlybyDecal"`
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
		<ButtonDrawTemplate>
			<Normal>A_SquareBtnNormal</Normal>
			<Pressed>A_SquareBtnPressed</Pressed>
			<Flyby>A_SquareBtnFlyby</Flyby>
			<Disabled>A_SquareBtnDisabled</Disabled>
			<PressedFlyby>A_SquareBtnPressedFlyby</PressedFlyby>
			<NormalDecal>MaleIcon</NormalDecal>
			<PressedDecal>MaleIcon</PressedDecal>
			<FlybyDecal>MaleIcon</FlybyDecal>
			<DisabledDecal>MaleIcon</DisabledDecal>
			<PressedFlybyDecal>MaleIcon</PressedFlybyDecal>
		</ButtonDrawTemplate>
		<Template>BDT_Normal</Template>
		<!--<SoundPressed/>-->
		<!--<SoundUp/>-->
		<!--<SoundFlyby/>-->
  </Button>
*/
