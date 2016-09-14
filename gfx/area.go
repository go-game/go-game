package gfx

import "github.com/go-gl/gl/v2.1/gl"

var x, y int32 = 0, 0
var pixelsize = 1
var width, height int
var fX, fY = 1, 1

// SetPixelSize scales all graphics by the given factor.
func SetPixelSize(size int) {
	pixelsize = size
	setGLViewPort()
}

// SetArea defines the area on the screen to be drawn to.
func SetArea(w, h int) {
	width = w
	height = h
	setGLViewPort()
}

func setGLViewPort() {
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(width/pixelsize), -float64(height/pixelsize), 0, -1, 1)
	// This is for retina stuff on macs
	// var width, height = desktop.CurrentWindow.GlfwWindow.GetFramebufferSize()
	// fX, fY := int32(width/desktop.CurrentWindow.Mode.Width), int32(height/desktop.CurrentWindow.Mode.Height)
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.Viewport(x, y, int32(width*fX), int32(height*fY))

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}
