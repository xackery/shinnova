package nv

import "github.com/xackery/shinnova/nxml"

// ButtonDrawTemplate represents a button draw template in the UI
type ButtonDrawTemplate struct {
	Item              string
	Normal            *string
	Pressed           *string
	Flyby             *string
	Disabled          *string
	PressedFlyby      *string
	NormalDecal       *string
	PressedDecal      *string
	FlybyDecal        *string
	DisabledDecal     *string
	PressedFlybyDecal *string
}

// NewButtonDrawTemplate creates a new button draw template
func NewButtonDrawTemplate(src *nxml.ButtonDrawTemplate) (ButtonDrawTemplate, error) {
	dst := ButtonDrawTemplate{Item: src.Item}
	if src.Normal != nil {
		dst.Normal = src.Normal
	}
	if src.Pressed != nil {
		dst.Pressed = src.Pressed
	}
	if src.Flyby != nil {
		dst.Flyby = src.Flyby
	}
	if src.Disabled != nil {
		dst.Disabled = src.Disabled
	}
	if src.PressedFlyby != nil {
		dst.PressedFlyby = src.PressedFlyby
	}
	if src.NormalDecal != nil {
		dst.NormalDecal = src.NormalDecal
	}
	if src.PressedDecal != nil {
		dst.PressedDecal = src.PressedDecal
	}
	if src.FlybyDecal != nil {
		dst.FlybyDecal = src.FlybyDecal
	}
	if src.DisabledDecal != nil {
		dst.DisabledDecal = src.DisabledDecal
	}
	if src.PressedFlybyDecal != nil {
		dst.PressedFlybyDecal = src.PressedFlybyDecal
	}
	return dst, nil
}
