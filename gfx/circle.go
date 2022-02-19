package gfx

import (
	"fmt"
	"math"
)

// NewCircle returns a pointer to a new Circle.
// the number of segments must be at least 3.
// The given mode is optional. If no mode is given a default LineMode with width 1 is used.
func NewCircle(r float64, segments int, filled bool, mode ...*LineMode) (*Circle, error) {
	if segments < 3 {
		return nil, fmt.Errorf("a circle needs at least 3 segments, got %d", segments)
	}
	if len(mode) > 1 {
		return nil, fmt.Errorf("NewCircle can be called with zero or one LineMode")
	}

	c := Circle{
		Radius: r,
		Filled: filled,
		coords: make([]float64, segments*2),
	}
	c.setSegments(segments)

	if len(mode) == 0 {
		c.Mode = NewLineMode()
	}

	return &c, nil
}

// Circle is a geometric shape that can be rendered with gfx.Render.
type Circle struct {
	Radius   float64
	segments int
	Filled   bool
	Mode     *LineMode
	coords   []float64
}

func (c *Circle) setSegments(s int) {
	c.segments = s

	diff := 2 * math.Pi / float64(s)
	angle := 0.0
	for i := 0; i < c.segments*2; i += 2 {
		c.coords[i] = math.Sin(angle) * c.Radius
		c.coords[i+1] = math.Cos(angle) * c.Radius
		angle += diff
	}
}

func (c *Circle) render() {
	_ = renderPolygon(c.Filled, c.Mode, c.coords...)
}
