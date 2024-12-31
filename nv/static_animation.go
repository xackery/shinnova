package nv

import (
	"fmt"
	"strconv"

	"github.com/xackery/shinnova/nxml"
)

// StaticAnimation represents a static animation in the UI
type StaticAnimation struct {
	Item       string
	ScreenID   *string
	LocationXY [2]*int
	SizeWH     [2]*int
	Animation  *string
}

// NewStaticAnimation converts an nxml.StaticAnimation to a nv.StaticAnimation
func NewStaticAnimation(src *nxml.StaticAnimation) (StaticAnimation, error) {
	var v1, v2 *int

	dst := StaticAnimation{Item: src.Item}
	if src.ScreenID != nil {
		dst.ScreenID = src.ScreenID
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

	if src.Animation != nil {
		dst.Animation = src.Animation
	}

	return dst, nil
}
