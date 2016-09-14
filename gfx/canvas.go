package gfx

import (
	"fmt"

	"github.com/go-gl/gl/v2.1/gl"
)

var currentName uint32

// Canvas is a offscreen area which can be rendered to and rendered on to the screen.
type Canvas struct {
	tex           *texture
	frameBufferID uint32
	width, height int32
}

// NewCanvas returns a new canvas.
func NewCanvas(width, height int) (*Canvas, error) {
	c := Canvas{width: int32(width), height: int32(height)}

	c.tex = newTexture(width, height, nil)

	gl.GenFramebuffersEXT(1, &c.frameBufferID)
	gl.BindFramebufferEXT(gl.FRAMEBUFFER_EXT, c.frameBufferID)
	gl.FramebufferTexture2DEXT(gl.FRAMEBUFFER_EXT, gl.COLOR_ATTACHMENT0_EXT, gl.TEXTURE_2D, c.tex.id, 0)

	status := gl.CheckFramebufferStatusEXT(gl.FRAMEBUFFER_EXT)
	if status != gl.FRAMEBUFFER_COMPLETE_EXT {
		return nil, fmt.Errorf("Cannot use Canvas: Framebuffer not supported")
	}

	return &c, nil
}

// Render uses a renderer to put pixels onto the canvas.
func (c *Canvas) Render(r Renderer, o *RenderOptions) {
	gl.Disable(gl.TEXTURE_2D)
	gl.Disable(gl.BLEND)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(c.width), -float64(c.height), 0, -1, 1)

	gl.MatrixMode(gl.MODELVIEW)

	transform(o)

	r.render(o)

	gl.BindFramebufferEXT(gl.FRAMEBUFFER_EXT, c.frameBufferID)
}

// Delete removes the canvas from memory.
func (c *Canvas) Delete() {
	gl.DeleteTextures(1, &c.tex.id)
	gl.BindFramebufferEXT(gl.FRAMEBUFFER_EXT, 0)
	gl.DeleteFramebuffersEXT(1, &c.frameBufferID)
}

// Clear clears the canvas
func (c *Canvas) Clear() {
	gl.Viewport(0, 0, c.width, c.height)
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	gl.ClearDepth(1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.BindFramebufferEXT(gl.FRAMEBUFFER_EXT, c.frameBufferID)
}

func (c *Canvas) render(o *RenderOptions) {
	gl.Enable(gl.TEXTURE_2D)
	// gl.Disable(gl.BLEND) // Needed?

	gl.BindTexture(gl.TEXTURE_2D, c.tex.id)

	gl.Begin(gl.QUADS)

	gl.Color4d(o.R, o.G, o.B, o.A)
	gl.TexCoord2f(0, 0)
	gl.Vertex3d(0, -float64(c.height), 0)

	gl.Color4d(o.R, o.G, o.B, o.A)
	gl.TexCoord2f(0, 1)
	gl.Vertex3d(0, 0, 0)

	gl.Color4d(o.R, o.G, o.B, o.A)
	gl.TexCoord2f(1, 1)
	gl.Vertex3d(float64(c.width), 0, 0)

	gl.Color4d(o.R, o.G, o.B, o.A)
	gl.TexCoord2f(1, 0)
	gl.Vertex3d(float64(c.width), -float64(c.height), 0)

	gl.End()
}
