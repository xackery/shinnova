package nv

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xackery/shinnova/nxml"
)

// InvSlot represents an inventory slot in the UI
type InvSlot struct {
	Item               string
	ScreenID           *string
	Font               *string
	IsRelativePosition *int
	LocationXY         [2]*int
	SizeWH             [2]*int
	Background         *string
	EQType             *string
}

// NewInvSlot converts an nxml.InvSlot to a nv.InvSlot
func NewInvSlot(src *nxml.InvSlot) (InvSlot, error) {
	var v1, v2 *int

	dst := InvSlot{Item: src.Item}
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
	if src.SizeCX != nil {
		val, err := strconv.Atoi(*src.SizeCX)
		if err != nil {
			return dst, fmt.Errorf("sizex: %w", err)
		}
		v1 = &val
	}
	if src.SizeCY != nil {
		val, err := strconv.Atoi(*src.SizeCY)
		if err != nil {
			return dst, fmt.Errorf("sizey: %w", err)
		}
		v2 = &val
	}
	dst.SizeWH = [2]*int{v1, v2}

	dst.Background = src.Background
	dst.EQType = src.EQType

	return dst, nil
}
