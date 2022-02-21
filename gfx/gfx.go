package gfx

import (
	"github.com/go-gl/gl/v2.1/gl"
)

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
	render()
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

// RenderXY renders the given renderer at the position x, y.
func RenderXY(r renderer, x, y float64) {
	currentCamera.RenderXY(r, x, y)
}

// RenderXYScale renders the given renderer at the position x, y and the given scale center and amount.
func RenderXYScale(rend renderer, x, y, sx, sy, s float64) {
	currentCamera.RenderXYScale(rend, x, y, sx, sy, s)
}

// RenderXYRot renders the given renderer at the position x, y and the given rotation center and angle clockwise.
func RenderXYRot(rend renderer, x, y, rx, ry, angle float64) {
	currentCamera.RenderXYRot(rend, x, y, rx, ry, angle)
}

// RenderXYColor renders the given renderer at the position x, y and the given color components.
func RenderXYColor(rend renderer, x, y, r, g, b, a float64) {
	currentCamera.RenderXYColor(rend, x, y, r, g, b, a)
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

func transform(x, y, scaleX, scaleY, scaleFactor, rotX, rotY, angle float64) {
	gl.LoadIdentity()

	gl.Translated(x, -y, 0)
	if activeCamera != nil {
		gl.Translated(activeCamera.posX, activeCamera.posY, 0)
	}

	gl.Translated(scaleX, -scaleY, 0)
	gl.Scaled(scaleFactor, scaleFactor, 1)
	gl.Translated(-scaleX, scaleY, 0)

	gl.Translated(rotX, -rotY, 0)
	gl.Rotated(-angle, 0, 0, 1)
	gl.Translated(-rotX, rotY, 0)
}
