package nv

import (
	"fmt"
	"strconv"

	"github.com/xackery/shinnova/nxml"
)

// TextureInfo represents a texture info in the UI
type TextureInfo struct {
	Item   string
	SizeXY [2]*int
}

// NewTextureInfo converts an nxml.TextureInfo to a nv.TextureInfo
func NewTextureInfo(src *nxml.TextureInfo) (TextureInfo, error) {
	var v1, v2 *int

	dst := TextureInfo{Item: src.Item}

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

	return dst, nil
}
