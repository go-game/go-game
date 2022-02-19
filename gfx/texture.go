package gfx

import (
	"github.com/go-gl/gl/v2.1/gl"
)

// FilterMode represents the interpolation mode for texture rendering.
type FilterMode int

const (
	// NearestFilter scales images with nearest neighbor interpolation.
	NearestFilter FilterMode = iota
	// LinearFilter scales image with linear interpolation.
	LinearFilter
)

type texture struct {
	id            uint32
	width, height float64
}

var currentlyBoundTextureID uint32

func (t *texture) activate(mode FilterMode) {
	gl.Enable(gl.TEXTURE_2D)

	if currentlyBoundTextureID != t.id {
		gl.BindTexture(gl.TEXTURE_2D, t.id)
		currentlyBoundTextureID = t.id
	}

	switch mode {
	case NearestFilter:
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	case LinearFilter:
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	}
}

func (t *texture) render() {
	gl.Begin(gl.QUADS)

	gl.TexCoord2f(0, 0)
	gl.Vertex3d(0, 0, 0)

	gl.TexCoord2f(0, 1)
	gl.Vertex3d(0, -t.height, 0)

	gl.TexCoord2f(1, 1)
	gl.Vertex3d(t.width, -t.height, 0)

	gl.TexCoord2f(1, 0)
	gl.Vertex3d(t.width, 0, 0)

	gl.End()
}

func (t *texture) delete() {
	gl.DeleteTextures(1, &t.id)
}

func newTexture(width, height int, pixelData []byte) *texture {
	numBytes := width * height / len(pixelData)

	format := gl.RGBA
	switch numBytes {
	case 1:
		format = gl.ALPHA
		for i, p := range pixelData {
			if p > 0 {
				pixelData[i] *= 255
			}
		}
	case 3:
		format = gl.RGB
	case 4:
		format = gl.RGBA
	}

	var id uint32
	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &id)
	gl.BindTexture(gl.TEXTURE_2D, id)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		int32(format),
		int32(width),
		int32(height),
		0,
		uint32(format),
		gl.UNSIGNED_BYTE,
		gl.Ptr(pixelData))
	return &texture{id: id, width: float64(width), height: float64(height)}
}
