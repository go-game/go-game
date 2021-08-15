package gfx

import (
	"fmt"

	"github.com/go-gl/gl/v2.1/gl"
)

// NewLineMode returns a pointer to a new LineMode with width 1.0 and not smooth.
func NewLineMode() *LineMode {
	return &LineMode{
		Width:  1.0,
		Smooth: false,
	}
}

// LineMode defines the way how (out)lines are rendered.
type LineMode struct {
	Width  float32
	Smooth bool
}

func renderPolygon(filled bool, mode *LineMode, coords ...float64) error {
	if len(coords)%2 != 0 {
		return fmt.Errorf("can only render an even number of x, y coords")
	}

	gl.Disable(gl.TEXTURE_2D)
	if filled {
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
	} else {
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	}
	renderPoints(gl.POLYGON, mode, coords...)

	gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)

	return nil
}

func renderPoints(glMode uint32, mode *LineMode, coords ...float64) {
	gl.LineWidth(mode.Width)
	if mode.Smooth {
		gl.Enable(gl.POINT_SMOOTH)
		gl.Enable(gl.LINE_SMOOTH)
		gl.Enable(gl.POLYGON_SMOOTH)
	} else {
		gl.Disable(gl.POINT_SMOOTH)
		gl.Disable(gl.LINE_SMOOTH)
		gl.Disable(gl.POLYGON_SMOOTH)
	}

	gl.Begin(glMode)
	for i := 0; i < len(coords); i += 2 {
		gl.Vertex3d(coords[i], -coords[i+1], 0)
	}
	gl.End()
}
