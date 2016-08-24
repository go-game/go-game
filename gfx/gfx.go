package gfx

import "github.com/go-gl/gl/v2.1/gl"

func init()  {
	if err := gl.Init(); err != nil {
		return
	}

	gl.Enable(gl.TEXTURE_2D)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.ClearColor(0.1, 0.1, 0.1, 0.0)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
}

// Clear clears the whole drawing area
func Clear() {
	gl.LoadIdentity()
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}
