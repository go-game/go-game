package gfx

import "github.com/go-gl/gl/v2.1/gl"

func init() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
}

// Width is the width in pixels of the complete drawing area.
var Width int32

// Height is the height in pixels of the complete drawing area.
var Height int32

// Fullscreen indicates whether the graphics are in fullscreen mode or not.
var Fullscreen bool

var clearR, clearG, clearB float32
var currentCamera *Camera

var defaultFilterMode = NearestFilter

type renderer interface {
	render(*Params)
}

// SetDefaultFilterMode sets the filter mode which will be used by default when
// creating a new Image or Canvas.
func SetDefaultFilterMode(fm FilterMode) {
	defaultFilterMode = fm
}

// SetPixelSize scales all graphics by the given factor.
func SetPixelSize(size int32) {
	currentCamera.SetPixelSize(size)
}

// CurrentCamera returns the currently used Camera.
func CurrentCamera() *Camera {
	return currentCamera
}

// SetCamera sets the currently used Camera.
func SetCamera(c *Camera) {
	currentCamera = c
}

// Render uses a renderer to put pixels onto the screen directly.
func Render(r renderer, p *Params) {
	currentCamera.Render(r, p)
}

// Clear clears the whole drawing area.
func Clear() {
	gl.ClearColor(clearR, clearG, clearB, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

// SetClearColor sets the color with which the screen will be filled when calling Clear().
func SetClearColor(r, g, b float64) {
	clearR = float32(r)
	clearG = float32(g)
	clearB = float32(b)
}

func transform(p *Params) {
	gl.LoadIdentity()

	gl.Translated(p.X, -p.Y, 0)
	if activeCamera != nil {
		gl.Translated(activeCamera.posX, activeCamera.posY, 0)
	}

	gl.Translated(p.Scale.X, -p.Scale.Y, 0)
	gl.Scaled(p.Scale.Factor, p.Scale.Factor, 1)
	gl.Translated(-p.Scale.X, p.Scale.Y, 0)

	gl.Translated(p.Rot.X, -p.Rot.Y, 0)
	gl.Rotated(-p.Rot.Angle, 0, 0, 1)
	gl.Translated(-p.Rot.X, p.Rot.Y, 0)
}
