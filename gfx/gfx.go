package gfx

import "github.com/go-gl/gl/v2.1/gl"

func init() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
}

var clearR, clearG, clearB, clearA float32

// Renderer imlpements render to put pixels on the screen.
type Renderer interface {
	render(o *RenderOptions)
}

// Render uses a renderer to put pixels onto the screen directly.
func Render(r Renderer, o *RenderOptions) {
	gl.LoadIdentity()

	setGLViewPort()

	gl.Translated(o.X, -o.Y, 0)

	gl.Translated(o.Scale.X, -o.Scale.Y, 0)
	gl.Scaled(o.Scale.Factor, o.Scale.Factor, 1)
	gl.Translated(-o.Scale.X, o.Scale.Y, 0)

	gl.Translated(o.Rot.X, -o.Rot.Y, 0)
	gl.Rotated(-o.Rot.Angle, 0, 0, 1)
	gl.Translated(-o.Rot.X, o.Rot.Y, 0)

	r.render(o)
}

// Clear clears the whole drawing area.
func Clear() {
	gl.ClearColor(clearR, clearG, clearB, clearA)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

// SetClearColor sets the color with which the screen will be filled when calling Clear().
func SetClearColor(r, g, b, a float64) {
	clearR = float32(r)
	clearG = float32(g)
	clearB = float32(b)
	clearA = float32(a)
}

// NewRenderOptions returns new RenderOptions with sensible default values.
func NewRenderOptions() *RenderOptions {
	return &RenderOptions{
		R:     1.0,
		G:     1.0,
		B:     1.0,
		A:     1.0,
		Rot:   Rotation{},
		Scale: Scale{Factor: 1},
	}
}

// RenderOptions encompasses all transformations that can be done while rendering an Image.
type RenderOptions struct {
	X     float64
	Y     float64
	R     float64
	G     float64
	B     float64
	A     float64
	Rot   Rotation
	Scale Scale
}

// Rotation describes the clockwise rotation around the center at X, Y.
type Rotation struct {
	Angle float64
	X     float64
	Y     float64
}

// Scale describes a scale with a center at x, y.
type Scale struct {
	Factor float64
	X      float64
	Y      float64
}
