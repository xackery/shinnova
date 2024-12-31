package nxml

// Ui2DAnimation represents a 2D animation in the UI
type Ui2DAnimation struct {
	Item   string      `xml:"item,attr"`
	Cycle  *string     `xml:"Cycle"`
	Grid   *string     `xml:"Grid"`
	Frames *Ui2DFrames `xml:"Frames"`
}

// Ui2DFrames represents the frames of a 2D animation in the UI
type Ui2DFrames struct {
	Texture   *string `xml:"Texture"`
	LocationX *string `xml:"Location>X"`
	LocationY *string `xml:"Location>Y"`
	SizeCX    *string `xml:"Size>CX"`
	SizeCY    *string `xml:"Size>CY"`
}

/*
<Ui2DAnimation item="RaceClassOverlay">
		<Cycle>false</Cycle>
		<Grid>false</Grid>
		<Frames>
			<Texture>FemaleRaceDisabled.tga</Texture>
			<Location>
				<X>0</X>
				<Y>191</Y>
			</Location>
			<Size>
				<CX>48</CX>
				<CY>48</CY>
			</Size>
		</Frames>
	</Ui2DAnimation>
*/
