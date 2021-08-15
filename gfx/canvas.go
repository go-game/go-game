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
	filterMode    FilterMode
}

// NewCanvas returns a new canvas.
func NewCanvas(width, height int) (*Canvas, error) {
	if !CanvasAvailable() {
		return nil, fmt.Errorf("cannot use Canvas: Framebuffer not supported")
	}

	c := Canvas{width: int32(width), height: int32(height), filterMode: defaultFilterMode}

	gl.GenFramebuffersEXT(1, &c.frameBufferID)
	gl.BindFramebufferEXT(gl.FRAMEBUFFER_EXT, c.frameBufferID)

	c.tex = newTexture(width, height, make([]byte, width, height*4))
	gl.FramebufferTexture2DEXT(gl.FRAMEBUFFER_EXT, gl.COLOR_ATTACHMENT0_EXT, gl.TEXTURE_2D, c.tex.id, 0)

	c.Clear()

	return &c, nil
}

// CanvasAvailable returns if the hardware supports a canvas.
func CanvasAvailable() bool {
	status := gl.CheckFramebufferStatusEXT(gl.FRAMEBUFFER_EXT)
	return status == gl.FRAMEBUFFER_COMPLETE_EXT
}

// Render uses a renderer to put pixels onto the canvas.
func (c *Canvas) Render(r renderer, p *Params) {
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(c.width), 0, -float64(c.height), -1, 1)
	gl.BindFramebufferEXT(gl.FRAMEBUFFER_EXT, c.frameBufferID)
	gl.Viewport(0, 0, c.width, c.height)
	gl.MatrixMode(gl.MODELVIEW)

	transform(p)
	r.render(p)

	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

// Delete removes the canvas from memory.
func (c *Canvas) Delete() {
	gl.DeleteTextures(1, &c.tex.id)
	gl.BindFramebufferEXT(gl.FRAMEBUFFER_EXT, 0)
	gl.DeleteFramebuffersEXT(1, &c.frameBufferID)
}

// Clear clears the canvas.
func (c *Canvas) Clear() {
	gl.Viewport(0, 0, c.width, c.height)
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	gl.ClearDepth(1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (c *Canvas) render(p *Params) {
	c.tex.activate(c.filterMode)
	c.tex.render(p)
}
