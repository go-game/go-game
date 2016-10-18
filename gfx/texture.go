package gfx

import "github.com/go-gl/gl/v2.1/gl"

type texture struct {
	id            uint32
	width, height float64
}

var currentlyBoundTextureID uint32

func (t *texture) render(o *RenderOptions) {
	gl.Enable(gl.TEXTURE_2D)

	if currentlyBoundTextureID != t.id {
		gl.BindTexture(gl.TEXTURE_2D, t.id)
		currentlyBoundTextureID = t.id
	}

	gl.Begin(gl.QUADS)

	gl.Color4d(o.R, o.G, o.B, o.A)
	gl.TexCoord2f(0, 0)
	gl.Vertex3d(0, 0, 0)

	gl.Color4d(o.R, o.G, o.B, o.A)
	gl.TexCoord2f(0, 1)
	gl.Vertex3d(0, -float64(t.height), 0)

	gl.Color4d(o.R, o.G, o.B, o.A)
	gl.TexCoord2f(1, 1)
	gl.Vertex3d(float64(t.width), -float64(t.height), 0)

	gl.Color4d(o.R, o.G, o.B, o.A)
	gl.TexCoord2f(1, 0)
	gl.Vertex3d(float64(t.width), 0, 0)

	gl.End()
}

func (t *texture) delete() {
	gl.DeleteTextures(1, &t.id)
}

func newTexture(width, height int, pixelData interface{}) *texture {
	var id uint32
	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &id)
	gl.BindTexture(gl.TEXTURE_2D, id)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(width),
		int32(height),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(pixelData))
	return &texture{id: id, width: float64(width), height: float64(height)}
}
