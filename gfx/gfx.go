package gfx

import "github.com/go-gl/gl/v2.1/gl"

func init() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
}

// Clear clears the whole drawing area.
func Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}
