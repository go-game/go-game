package gfx

import (
	"math"

	"github.com/go-gl/gl/v2.1/gl"
)

// Circle is a geometric shape that can be rendered with gfx.Render.
type Circle struct {
	Radius   float64
	Segments int
	Filled   bool
}

func (c *Circle) render(p *Params) {
	gl.Color4d(p.R, p.G, p.B, p.A)
	coords := make([]float64, c.Segments*2)
	diff := 2 * math.Pi / float64(c.Segments)
	angle := 0.0
	for i := 0; i < c.Segments*2; i += 2 {
		coords[i] = math.Sin(angle) * c.Radius
		coords[i+1] = math.Cos(angle) * c.Radius
		angle += diff
	}

	RenderPolygon(c.Filled, coords...)
}
