package nv

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xackery/shinnova/nxml"
)

// Label represents a label in the UI
type Label struct {
	Item               string
	ScreenID           *string
	Font               *int
	IsRelativePosition *int
	LocationXY         [2]*int
	SizeWH             [2]*int
	Text               *string
	TextColorRGB       [3]*int
	IsNoWrap           *int
	IsAlignCenter      *int
	IsAlignRight       *int
}

func NewLabel(src *nxml.Label) (Label, error) {
	dst := Label{}
	dst.Item = src.Item
	if src.ScreenID != nil {
		dst.ScreenID = src.ScreenID
	}
	var v1, v2, v3 *int
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
	dst.SizeWH = [2]*int{v1, v2}
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

	v1 = nil
	if src.NoWrap != nil {
		val := 0
		if strings.ToLower(*src.NoWrap) == "true" {
			val = 1
		}
		v1 = &val
	}
	dst.IsNoWrap = v1

	v1 = nil
	if src.AlignCenter != nil {
		val := 0
		if strings.ToLower(*src.AlignCenter) == "true" {
			val = 1
		}
		v1 = &val
	}
	dst.IsAlignCenter = v1

	v1 = nil
	if src.AlignRight != nil {
		val := 0
		if strings.ToLower(*src.AlignRight) == "true" {
			val = 1
		}
		v1 = &val
	}
	dst.IsAlignRight = v1

	return dst, nil
}
