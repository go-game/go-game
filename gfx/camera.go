package gfx

import "github.com/go-gl/gl/v2.1/gl"

var activeCamera *Camera

// Camera represents an area on the screen.
type Camera struct {
	pixelSize     int
	x, y          int32
	width, height int
}

// NewCamera returns a new Camera
func NewCamera(w, h int) *Camera {
	c := Camera{width: w, height: h, pixelSize: 1}
	c.setGLViewPort()
	return &c
}

// Render renders the given Renderer with the given RenderOptions.
func (c *Camera) Render(r Renderer, o *RenderOptions) {
	if activeCamera != c {
		c.setGLViewPort()
		activeCamera = c
	}
	gl.LoadIdentity()
	transform(o)
	r.render(o)
}

// SetPixelSize scales all graphics by the given factor.
func (c *Camera) SetPixelSize(i int) {
	c.pixelSize = i
	c.setGLViewPort()
}

func (c *Camera) setGLViewPort() {
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(c.width/c.pixelSize), -float64(c.height/c.pixelSize), 0, -1, 1)
	// This is for retina stuff on macs
	// var c.width, c.height = desktop.CurrentWindow.GlfwWindow.GetFramebufferSize()
	// fX, fY := int32(width/desktop.CurrentWindow.Mode.Width), int32(height/desktop.CurrentWindow.Mode.Height)
	fX, fY := 1, 1
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.Viewport(c.x, c.y, int32(c.width*fX), int32(c.height*fY))

	gl.MatrixMode(gl.MODELVIEW)
}
