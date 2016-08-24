package gfx

import "github.com/go-gl/gl/v2.1/gl"

var x int32 = 0
var y int32 = 0
var width, height, pixelsize int

// SetArea defines the area to be drawn
func SetArea(width, height, pixelsize int) {
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(width/pixelsize), 0, float64(height/pixelsize), -1, 1)
  // This is for retina stuff on macs
  // var width, height = desktop.CurrentWindow.GlfwWindow.GetFramebufferSize()
	// fX, fY := int32(width/desktop.CurrentWindow.Mode.Width), int32(height/desktop.CurrentWindow.Mode.Height)
	gl.Viewport(x, y, int32(width), int32(height))

	gl.MatrixMode(gl.MODELVIEW)
}
