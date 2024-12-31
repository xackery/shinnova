package nv

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xackery/shinnova/nxml"
)

// Ui2DAnimation represents a 2D animation in the UI
type Ui2DAnimation struct {
	Item            string
	IsCycle         *int
	IsGrid          *int
	FrameTexture    *string
	FrameLocationXY [2]*int
	FrameSizeWH     [2]*int
}

// NewUi2DAnimation converts an nxml.Ui2DAnimation to a nv.Ui2DAnimation
func NewUi2DAnimation(src *nxml.Ui2DAnimation) (Ui2DAnimation, error) {
	var v1, v2, v3 *int

	dst := Ui2DAnimation{Item: src.Item}
	if src.Cycle != nil {
		val := 0
		if strings.ToLower(*src.Cycle) == "true" {
			val = 1
		}
		dst.IsCycle = &val
	}
	if src.Grid != nil {
		val := 0
		if strings.ToLower(*src.Grid) == "true" {
			val = 1
		}
		dst.IsGrid = &val
	}
	if src.Frames != nil {
		dst.FrameTexture = src.Frames.Texture
		v1 = nil
		if src.Frames.LocationX != nil {
			val, err := strconv.Atoi(*src.Frames.LocationX)
			if err != nil {
				return dst, fmt.Errorf("frames.locationx: %w", err)
			}
			v1 = &val
		}
		v2 = nil
		if src.Frames.LocationY != nil {
			val, err := strconv.Atoi(*src.Frames.LocationY)
			if err != nil {
				return dst, fmt.Errorf("frames.locationy: %w", err)
			}
			v2 = &val
		}
		dst.FrameLocationXY = [2]*int{v1, v2}

		v1, v2 = nil, nil
		if src.Frames.SizeCX != nil {
			val, err := strconv.Atoi(*src.Frames.SizeCX)
			if err != nil {
				return dst, fmt.Errorf("frames.sizecx: %w", err)
			}
			v1 = &val
		}
		if src.Frames.SizeCY != nil {
			val, err := strconv.Atoi(*src.Frames.SizeCY)
			if err != nil {
				return dst, fmt.Errorf("frames.sizecy: %w", err)
			}
			v2 = &val
		}
		dst.FrameSizeWH = [2]*int{v1, v2}

	}
	if v3 != nil {

	}
	return dst, nil
}
