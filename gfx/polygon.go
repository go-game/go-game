package gfx

import (
	"github.com/go-gl/gl/v2.1/gl"
)

// Polygon is a geometric shape that can be rendered with gfx.Render.
type Polygon struct {
	Points []float64
	Filled bool
	Mode   *LineMode
}

func (p *Polygon) render(params *Params) {
	gl.Color4d(params.R, params.G, params.B, params.A)
	_ = renderPolygon(p.Filled, p.Mode, p.Points...)
}
