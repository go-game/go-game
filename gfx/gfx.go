package gfx

import "github.com/go-gl/gl/v2.1/gl"

// Clear clear the whole drawing area
func Clear() {
	gl.LoadIdentity()
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}
