package nv

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xackery/shinnova/nxml"
)

// Screen represents a screen in the UI
type Screen struct {
	Item               string
	ScreenID           *string
	Font               *string
	IsRelativePosition *int
	LocationXY         [2]*int
	SizeXY             [2]*int
	IsStyleVScroll     *int
	IsStyleHScroll     *int
	IsStyleTransparent *int
	TooltipReference   *string
	DrawTemplate       *string
	IsStyleTitlebar    *int
	IsStyleClosebox    *int
	IsStyleMinimizebox *int
	IsStyleBorder      *int
	IsStyleSizable     *int
	IsStyleQmarkbox    *int
	Pieces             []string
}

// NewScreen converts an nxml.Screen to a nv.Screen
func NewScreen(src *nxml.Screen) (Screen, error) {
	var v1, v2 *int

	dst := Screen{
		Item: src.Item,
	}

	if src.ScreenID != nil {
		dst.ScreenID = src.ScreenID
	}
	if src.Font != nil {
		dst.Font = src.Font
	}
	if src.RelativePosition != nil {
		val := 0
		if strings.ToLower(*src.RelativePosition) == "true" {
			val = 1
		}
		dst.IsRelativePosition = &val
	}

	v1 = nil
	if src.LocationX != nil {
		val, err := strconv.Atoi(*src.LocationX)
		if err != nil {
			return dst, fmt.Errorf("locationx: %w", err)
		}
		v1 = &val
	}
	v2 = nil
	if src.LocationY != nil {
		val, err := strconv.Atoi(*src.LocationY)
		if err != nil {
			return dst, fmt.Errorf("locationy: %w", err)
		}
		v2 = &val
	}
	dst.LocationXY = [2]*int{v1, v2}

	v1, v2 = nil, nil
	if src.SizeX != nil {
		val, err := strconv.Atoi(*src.SizeX)
		if err != nil {
			return dst, fmt.Errorf("sizex: %w", err)
		}
		v1 = &val
	}
	if src.SizeY != nil {
		val, err := strconv.Atoi(*src.SizeY)
		if err != nil {
			return dst, fmt.Errorf("sizey: %w", err)
		}
		v2 = &val
	}
	dst.SizeXY = [2]*int{v1, v2}

	if src.StyleVScroll != nil {
		val := 0
		if strings.ToLower(*src.StyleVScroll) == "true" {
			val = 1
		}
		dst.IsStyleVScroll = &val
	}
	if src.StyleHScroll != nil {
		val := 0
		if strings.ToLower(*src.StyleHScroll) == "true" {
			val = 1
		}
		dst.IsStyleHScroll = &val
	}
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
	if src.DrawTemplate != nil {
		dst.DrawTemplate = src.DrawTemplate
	}
	if src.StyleTitlebar != nil {
		val := 0
		if strings.ToLower(*src.StyleTitlebar) == "true" {
			val = 1
		}
		dst.IsStyleTitlebar = &val
	}
	if src.StyleClosebox != nil {
		val := 0
		if strings.ToLower(*src.StyleClosebox) == "true" {
			val = 1
		}
		dst.IsStyleClosebox = &val
	}
	if src.StyleMinimizebox != nil {
		val := 0
		if strings.ToLower(*src.StyleMinimizebox) == "true" {
			val = 1
		}
		dst.IsStyleMinimizebox = &val
	}
	if src.StyleBorder != nil {
		val := 0
		if strings.ToLower(*src.StyleBorder) == "true" {
			val = 1
		}
		dst.IsStyleBorder = &val
	}
	if src.StyleSizable != nil {
		val := 0
		if strings.ToLower(*src.StyleSizable) == "true" {
			val = 1
		}
		dst.IsStyleSizable = &val
	}
	if src.StyleQmarkbox != nil {
		val := 0
		if strings.ToLower(*src.StyleQmarkbox) == "true" {
			val = 1
		}
		dst.IsStyleQmarkbox = &val
	}
	dst.Pieces = src.Pieces

	return dst, nil
}
