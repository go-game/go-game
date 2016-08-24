package gfx

import "github.com/go-gl/gl/v2.1/gl"

var x int32 = 0
var y int32 = 0
var pixelsize = 1
var width, height int
var fX, fY = 1, 1

// SetArea defines the area to be drawn
func SetArea(w, h int) {
	width = w
	height = h

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(width/pixelsize), 0, float64(height/pixelsize), -1, 1)
  // This is for retina stuff on macs
  // var width, height = desktop.CurrentWindow.GlfwWindow.GetFramebufferSize()
	// fX, fY := int32(width/desktop.CurrentWindow.Mode.Width), int32(height/desktop.CurrentWindow.Mode.Height)
	gl.Viewport(x, y, int32(width * fX), int32(height * fY))

	gl.MatrixMode(gl.MODELVIEW)
}
