package nv

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xackery/shinnova/nxml"
)

// Button represents a button in the UI
type Button struct {
	Item                                string
	ScreenID                            *string
	Font                                *string
	RadioGroup                          *string
	IsRelativePosition                  *int
	LocationXY                          [2]*int
	SizeWH                              [2]*int
	IsStyleTransparent                  *int
	TooltipReference                    *string
	IsStyleCheckbox                     *int
	Text                                *string
	TextColorRGB                        [3]*int
	Template                            *string
	DecalOffsetXY                       [2]*int
	DecalSizeWH                         [2]*int
	ButtonDrawTemplateNormal            *string
	ButtonDrawTemplatePressed           *string
	ButtonDrawTemplateFlyby             *string
	ButtonDrawTemplateDisabled          *string
	ButtonDrawTemplatePressedFlyby      *string
	ButtonDrawTemplateNormalDecal       *string
	ButtonDrawTemplatePressedDecal      *string
	ButtonDrawTemplateFlybyDecal        *string
	ButtonDrawTemplateDisabledDecal     *string
	ButtonDrawTemplatePressedFlybyDecal *string
}

func NewButton(src *nxml.Button) (Button, error) {
	var v1, v2, v3 *int

	dst := Button{Item: src.Item}
	if src.ScreenID != nil {
		dst.ScreenID = src.ScreenID
	}
	if src.Font != nil {
		dst.Font = src.Font
	}
	if src.RadioGroup != nil {
		dst.RadioGroup = src.RadioGroup
	}
	if src.RelativePosition != nil {
		val := 0
		if strings.ToLower(*src.RelativePosition) == "true" {
			val = 1
		}
		dst.IsRelativePosition = &val
	}
	v1 = nil
	if src.SizeX != nil {
		val, err := strconv.Atoi(*src.SizeX)
		if err != nil {
			return dst, fmt.Errorf("sizex: %w", err)
		}
		v1 = &val
	}
	v2 = nil
	if src.SizeY != nil {
		val, err := strconv.Atoi(*src.SizeY)
		if err != nil {
			return dst, fmt.Errorf("sizey: %w", err)
		}
		v2 = &val
	}
	dst.SizeWH = [2]*int{v1, v2}

	if src.StyleTransparent != nil {
		val := 0
		if strings.ToLower(*src.StyleTransparent) == "true" {
			val = 1
		}
		dst.IsStyleTransparent = &val
	}

	if src.TooltipReference != nil {
		dst.TooltipReference = src.TooltipReference
	}

	if src.StyleCheckbox != nil {
		val := 0
		if strings.ToLower(*src.StyleCheckbox) == "true" {
			val = 1
		}
		dst.IsStyleCheckbox = &val
	}

	if src.Text != nil {
		dst.Text = src.Text
	}
	v1, v2, v3 = nil, nil, nil
	if src.TextColorR != nil {
		val, err := strconv.Atoi(*src.TextColorR)
		if err != nil {
			return dst, fmt.Errorf("textcolorr: %w", err)
		}
		v1 = &val
	}
	if src.TextColorG != nil {
		val, err := strconv.Atoi(*src.TextColorG)
		if err != nil {
			return dst, fmt.Errorf("textcolorg: %w", err)
		}
		v2 = &val
	}
	if src.TextColorB != nil {
		val, err := strconv.Atoi(*src.TextColorB)
		if err != nil {
			return dst, fmt.Errorf("textcolorb: %w", err)
		}
		v3 = &val
	}
	dst.TextColorRGB = [3]*int{v1, v2, v3}

	if src.Template != nil {
		dst.Template = src.Template
	}

	if src.ButtonDrawTemplate != nil {
		if src.ButtonDrawTemplate.Normal != nil {
			dst.ButtonDrawTemplateNormal = src.ButtonDrawTemplate.Normal
		}
		if src.ButtonDrawTemplate.Pressed != nil {
			dst.ButtonDrawTemplatePressed = src.ButtonDrawTemplate.Pressed
		}
		if src.ButtonDrawTemplate.Flyby != nil {
			dst.ButtonDrawTemplateFlyby = src.ButtonDrawTemplate.Flyby
		}
		if src.ButtonDrawTemplate.Disabled != nil {
			dst.ButtonDrawTemplateDisabled = src.ButtonDrawTemplate.Disabled
		}
		if src.ButtonDrawTemplate.PressedFlyby != nil {
			dst.ButtonDrawTemplatePressedFlyby = src.ButtonDrawTemplate.PressedFlyby
		}
		if src.ButtonDrawTemplate.NormalDecal != nil {
			dst.ButtonDrawTemplateNormalDecal = src.ButtonDrawTemplate.NormalDecal
		}
		if src.ButtonDrawTemplate.PressedDecal != nil {
			dst.ButtonDrawTemplatePressedDecal = src.ButtonDrawTemplate.PressedDecal
		}
		if src.ButtonDrawTemplate.FlybyDecal != nil {
			dst.ButtonDrawTemplateFlybyDecal = src.ButtonDrawTemplate.FlybyDecal
		}
		if src.ButtonDrawTemplate.DisabledDecal != nil {
			dst.ButtonDrawTemplateDisabledDecal = src.ButtonDrawTemplate.DisabledDecal
		}
		if src.ButtonDrawTemplate.PressedFlybyDecal != nil {
			dst.ButtonDrawTemplatePressedFlybyDecal = src.ButtonDrawTemplate.PressedFlybyDecal
		}
	}

	v1, v2 = nil, nil
	if src.LocationX != nil {
		val, err := strconv.Atoi(*src.LocationX)
		if err != nil {
			return dst, fmt.Errorf("locationx: %w", err)
		}
		v1 = &val
	}
	if src.LocationY != nil {
		val, err := strconv.Atoi(*src.LocationY)
		if err != nil {
			return dst, fmt.Errorf("locationy: %w", err)
		}
		v2 = &val
	}
	dst.LocationXY = [2]*int{v1, v2}

	v1, v2 = nil, nil
	if src.DecalOffsetX != nil {
		val, err := strconv.Atoi(*src.DecalOffsetX)
		if err != nil {
			return dst, fmt.Errorf("decaloffsetx: %w", err)
		}
		v1 = &val
	}
	if src.DecalOffsetY != nil {
		val, err := strconv.Atoi(*src.DecalOffsetY)
		if err != nil {
			return dst, fmt.Errorf("decaloffsety: %w", err)
		}
		v2 = &val
	}
	dst.DecalOffsetXY = [2]*int{v1, v2}

	v1, v2 = nil, nil
	if src.DecalSizeCX != nil {
		val, err := strconv.Atoi(*src.DecalSizeCX)
		if err != nil {
			return dst, fmt.Errorf("decalsizecx: %w", err)
		}
		v1 = &val
	}
	if src.DecalSizeCY != nil {
		val, err := strconv.Atoi(*src.DecalSizeCY)
		if err != nil {
			return dst, fmt.Errorf("decalsizecy: %w", err)
		}
		v2 = &val
	}
	dst.DecalSizeWH = [2]*int{v1, v2}

	return dst, nil
}
