package gfx

import (
	"fmt"

	"github.com/go-gl/gl/v2.1/gl"
)

// Line is a geometric shape that can be rendered with gfx.Render.
type Line struct {
	Points []float64
	Mode   *LineMode
}

func (l *Line) render() {
	renderPoints(gl.LINES, l.Mode, l.Points...)
}

// NewLine returns a pointer to a new line for the given coordinates and Line
// mode.
// The number of coordinates must be at least 4.
func NewLine(mode *LineMode, coordinates ...float64) (*Line, error) {
	if len(coordinates) < 3 {
		return nil, fmt.Errorf("NewLine can only be calles with at least 4 coordinates")
	}
	return &Line{Points: coordinates, Mode: mode}, nil
}
