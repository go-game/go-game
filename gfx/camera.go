package gfx

import (
	"github.com/go-gl/gl/v2.1/gl"
)

var activeCamera *Camera

// Camera represents an area on the screen.
type Camera struct {
	pixelSize     int32
	X, Y          int32
	Width, Height int32
	posX          float64
	posY          float64
}

// NewCamera returns a new Camera
func NewCamera(w, h, x, y, ps int32) *Camera {
	c := Camera{Width: w, Height: h, X: x, Y: y, pixelSize: ps}
	c.setGLViewPort()
	return &c
}

func (c *Camera) RenderXY(r renderer) {
	//c.Re
}

// Render renders the given Renderer with the given Params.
func (c *Camera) Render(r renderer, p *Params) {
	if activeCamera != c {
		activeCamera = c
		c.setGLViewPort()
	}
	transform(p.X, p.Y, p.Scale.X, p.Scale.Y, p.Scale.Factor, p.Rot.X, p.Rot.Y, p.Rot.Angle)
	gl.Color4d(p.R, p.G, p.B, p.A)
	r.render()
}

// SetPixelSize scales all graphics by the given factor.
func (c *Camera) SetPixelSize(i int32) {
	c.pixelSize = i
	c.setGLViewPort()
}

// Move moves the camera bey the given values.
func (c *Camera) Move(x, y float64) {
	c.posX += x
	c.posY += y
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
	var fX, fY int32 = 1, 1
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.Viewport(c.X, c.Y, c.Width*fX, c.Height*fY)

	gl.MatrixMode(gl.MODELVIEW)
}
