package gfx

import (
	"fmt"

	"github.com/go-gl/gl/v2.1/gl"
)

var lineWidth float32 = 1
var lineR float64 = 1
var lineG float64 = 1
var lineB float64 = 1
var lineA float64 = 1
var smoothLines bool

// SetLineWidth sets the width of the line for all primitives.
func SetLineWidth(w float64) {
	lineWidth = float32(w)
}

// SetLineColor sets the color of the line for all primitives.
func SetLineColor(r, g, b, a float64) {
	lineR, lineG, lineB, lineA = r, g, b, a
}

// SetSmoothLines sets if lines should be drawn with jagged edges or anti-aliasing.
func SetSmoothLines(b bool) {
	smoothLines = b
}

// RenderLines renders multiple lines for the given coords, coords must be an even number of floats with alternating x and y coordinates.
func RenderLines(coords ...float64) error {
	if len(coords)%4 != 0 {
		return fmt.Errorf("Can only render an even number of x, y coords")
	}
	renderPoints(gl.LINES, coords...)
	return nil
}

// RenderPolygon renders a polygon from the given coords. When filled is true, the shape will be filled with a solid color..
func RenderPolygon(filled bool, coords ...float64) error {
	gl.Disable(gl.TEXTURE_2D)
	if len(coords)%2 != 0 {
		return fmt.Errorf("Can only render an even number of x, y coords")
	}
	if filled {
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
	} else {
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	}
	var mode uint32 = gl.POLYGON
	renderPoints(mode, coords...)
	return nil
}

// RenderRectangle redners a rectangle for the given upper left and lower right corner. If filled is true the rectangle will be filled with a color.
func RenderRectangle(filled bool, x1, y1, x2, y2 float64) {
	RenderPolygon(filled, x1, y1, x1, y2, x2, y2, x2, y1)
}

func renderPoints(mode uint32, coords ...float64) {
	gl.LineWidth(lineWidth)
	if smoothLines {
		gl.Enable(gl.POINT_SMOOTH)
		gl.Enable(gl.LINE_SMOOTH)
		gl.Enable(gl.POLYGON_SMOOTH)
	} else {
		gl.Disable(gl.POINT_SMOOTH)
		gl.Disable(gl.LINE_SMOOTH)
		gl.Disable(gl.POLYGON_SMOOTH)
	}

	gl.Begin(mode)
	for i := 0; i < len(coords); i += 2 {
		gl.Vertex3d(coords[i], -coords[i+1], 0)
	}
	gl.End()
}
