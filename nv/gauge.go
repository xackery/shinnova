package nv

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xackery/shinnova/nxml"
)

// Gauge represents a gauge in the UI
type Gauge struct {
	Item                         string
	ScreenID                     *string
	IsRelativePosition           *int
	GaugeOffsetXY                [2]*int
	FillTintRGB                  [3]*int
	EQType                       *string
	GaugeDrawTemplateBackground  *string
	GaugeDrawTemplateFill        *string
	GaugeDrawTemplateLines       *string
	GaugeDrawTemplateEndCapLeft  *string
	GaugeDrawTemplateEndCapRight *string
	IsAutoStretch                *int
	TopAnchorOffset              *int
	BottomAnchorOffset           *int
	LeftAnchorOffset             *int
	RightAnchorOffset            *int
	IsRightAnchorToLeft          *int
}

// NewGauge converts an nxml.Gauge to a nv.Gauge
func NewGauge(src *nxml.Gauge) (Gauge, error) {
	var v1, v2, v3 *int

	dst := Gauge{Item: src.Item}

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

	v1, v2 = nil, nil
	if src.GaugeOffsetX != nil {
		val, err := strconv.Atoi(*src.GaugeOffsetX)
		if err != nil {
			return dst, fmt.Errorf("gaugeoffsetx: %w", err)
		}
		v1 = &val
	}
	if src.GaugeOffsetY != nil {
		val, err := strconv.Atoi(*src.GaugeOffsetY)
		if err != nil {
			return dst, fmt.Errorf("gaugeoffsety: %w", err)
		}
		v2 = &val
	}
	dst.GaugeOffsetXY = [2]*int{v1, v2}

	v1, v2, v3 = nil, nil, nil
	if src.FillTintR != nil {
		val, err := strconv.Atoi(*src.FillTintR)
		if err != nil {
			return dst, fmt.Errorf("filltintr: %w", err)
		}
		v1 = &val
	}
	if src.FillTintG != nil {
		val, err := strconv.Atoi(*src.FillTintG)
		if err != nil {
			return dst, fmt.Errorf("filltintg: %w", err)
		}
		v2 = &val
	}
	if src.FillTintB != nil {
		val, err := strconv.Atoi(*src.FillTintB)
		if err != nil {
			return dst, fmt.Errorf("filltintb: %w", err)
		}
		v3 = &val
	}
	dst.FillTintRGB = [3]*int{v1, v2, v3}

	if src.EQType != nil {
		dst.EQType = src.EQType
	}

	if src.GaugeDrawTemplate != nil {
		if src.GaugeDrawTemplate.Background != nil {
			dst.GaugeDrawTemplateBackground = src.GaugeDrawTemplate.Background
		}
		if src.GaugeDrawTemplate.Fill != nil {
			dst.GaugeDrawTemplateFill = src.GaugeDrawTemplate.Fill
		}
		if src.GaugeDrawTemplate.Lines != nil {
			dst.GaugeDrawTemplateLines = src.GaugeDrawTemplate.Lines
		}
		if src.GaugeDrawTemplate.EndCapLeft != nil {
			dst.GaugeDrawTemplateEndCapLeft = src.GaugeDrawTemplate.EndCapLeft
		}
		if src.GaugeDrawTemplate.EndCapRight != nil {
			dst.GaugeDrawTemplateEndCapRight = src.GaugeDrawTemplate.EndCapRight
		}
	}

	if src.AutoStretch != nil {
		val := 0
		if strings.ToLower(*src.AutoStretch) == "true" {
			val = 1
		}
		dst.IsAutoStretch = &val
	}

	if src.TopAnchorOffset != nil {
		val, err := strconv.Atoi(*src.TopAnchorOffset)
		if err != nil {
			return dst, fmt.Errorf("topanchoroffset: %w", err)
		}
		dst.TopAnchorOffset = &val
	}

	if src.BottomAnchorOffset != nil {
		val, err := strconv.Atoi(*src.BottomAnchorOffset)
		if err != nil {
			return dst, fmt.Errorf("bottomanchoroffset: %w", err)
		}
		dst.BottomAnchorOffset = &val
	}

	if src.LeftAnchorOffset != nil {
		val, err := strconv.Atoi(*src.LeftAnchorOffset)
		if err != nil {
			return dst, fmt.Errorf("leftanchoroffset: %w", err)
		}
		dst.LeftAnchorOffset = &val
	}

	if src.RightAnchorOffset != nil {
		val, err := strconv.Atoi(*src.RightAnchorOffset)
		if err != nil {
			return dst, fmt.Errorf("rightanchoroffset: %w", err)
		}
		dst.RightAnchorOffset = &val
	}

	if src.RightAnchorToLeft != nil {
		val := 0
		if strings.ToLower(*src.RightAnchorToLeft) == "true" {
			val = 1
		}
		dst.IsRightAnchorToLeft = &val
	}

	return dst, nil
}
