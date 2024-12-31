package nxml

// ButtonDrawTemplate represents a button draw template in the UI
type ButtonDrawTemplate struct {
	Item              string  `xml:"item,attr"`
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

/* <ButtonDrawTemplate item="CC_BDT_RaceClassOverlay">
<Normal>RaceClassOverlay</Normal>
<Pressed>RaceClassOverlay</Pressed>
<Flyby>RaceClassOverlay</Flyby>
<Disabled>RaceClassOverlay</Disabled>
<PressedFlyby>RaceClassOverlay</PressedFlyby>
</ButtonDrawTemplate> */
