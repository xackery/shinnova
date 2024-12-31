package nv

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xackery/shinnova/nxml"
)

// Combobox represents a combobox in the UI
type Combobox struct {
	Item               string
	ScreenID           *string
	IsRelativePosition *int
	LocationXY         [2]*int
	SizeWH             [2]*int
	DrawTemplate       *string
	ListHeight         *int
	Button             *string
	IsStyleBorder      *int
}

// NewCombobox converts an nxml.Combobox to a nv.Combobox
func NewCombobox(src *nxml.Combobox) (Combobox, error) {
	var v1, v2 *int

	dst := Combobox{Item: src.Item}
	if src.ScreenID != nil {
		dst.ScreenID = src.ScreenID
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
	if src.SizeCX != nil {
		val, err := strconv.Atoi(*src.SizeCX)
		if err != nil {
			return dst, fmt.Errorf("sizecx: %w", err)
		}
		v1 = &val
	}
	if src.SizeCY != nil {
		val, err := strconv.Atoi(*src.SizeCY)
		if err != nil {
			return dst, fmt.Errorf("sizecy: %w", err)
		}
		v2 = &val
	}
	dst.SizeWH = [2]*int{v1, v2}

	if src.DrawTemplate != nil {
		dst.DrawTemplate = src.DrawTemplate
	}

	if src.ListHeight != nil {
		val, err := strconv.Atoi(*src.ListHeight)
		if err != nil {
			return dst, fmt.Errorf("listheight: %w", err)
		}
		dst.ListHeight = &val
	}

	if src.Button != nil {
		dst.Button = src.Button
	}

	if src.StyleBorder != nil {
		val := 0
		if strings.ToLower(*src.StyleBorder) == "true" {
			val = 1
		}
		dst.IsStyleBorder = &val
	}

	return dst, nil
}
