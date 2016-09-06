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

// NewRenderOptions returns new RenderOptions with sensible default values.
func NewRenderOptions() *RenderOptions {
	return &RenderOptions{
		R: 1.0,
		G: 1.0,
		B: 1.0,
		A: 1.0,
	}
}

// RenderOptions encompasses all transformations that can be done while rendering an Image.
type RenderOptions struct {
	X float64
	Y float64
	R float64
	G float64
	B float64
	A float64
}
