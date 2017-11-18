package gfx

import "github.com/go-gl/gl/v2.1/gl"

var activeCamera *Camera

// Camera represents an area on the screen.
type Camera struct {
	pixelSize     int
	X, Y          int32
	Width, Height int
}

// NewCamera returns a new Camera
func NewCamera(w, h, x, y, ps int) *Camera {
	c := Camera{Width: w, Height: h, X: int32(x), Y: int32(y), pixelSize: ps}
	c.setGLViewPort()
	return &c
}

// Render renders the given Renderer with the given Params.
func (c *Camera) Render(r Renderer, p *Params) {
	if activeCamera != c {
		c.setGLViewPort()
		activeCamera = c
	}
	gl.LoadIdentity()
	transform(p)
	r.render(p)
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
	gl.Ortho(0, float64(c.Width/c.pixelSize), -float64(c.Height/c.pixelSize), 0, -1, 1)
	// This is for retina stuff on macs
	// var c.width, c.height = desktop.CurrentWindow.GlfwWindow.GetFramebufferSize()
	// fX, fY := int32(width/desktop.CurrentWindow.Mode.Width), int32(height/desktop.CurrentWindow.Mode.Height)
	fX, fY := 1, 1
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.Viewport(c.X, c.Y, int32(c.Width*fX), int32(c.Height*fY))

	gl.MatrixMode(gl.MODELVIEW)
}
