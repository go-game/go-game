package gfx

import (
	"fmt"
)

// NewRectangle returns a pointer ro a new rectangle for the given coordinates.
// The given mode is optional. If no mode is given a default LineMode with width 1 is used.
func NewRectangle(x1, y1, x2, y2 float64, filled bool, mode ...*LineMode) (*Rectangle, error) {
	if len(mode) > 1 {
		return nil, fmt.Errorf("NewRectangle can be called with zero or one LineMode")
	}

	r := Rectangle{
		X1:     x1,
		Y1:     y1,
		X2:     x2,
		Y2:     y2,
		Filled: filled,
	}

	return &r, nil
}

// Rectangle is a geometric shape that can be rendered with gfx.Render.
type Rectangle struct {
	X1     float64
	Y1     float64
	X2     float64
	Y2     float64
	Filled bool
	Mode   *LineMode
}

func (r *Rectangle) render() {
	_ = renderPolygon(r.Filled, r.Mode, r.X1, r.Y1, r.X1, r.Y2, r.X2, r.Y2, r.X2, r.Y1)
}
