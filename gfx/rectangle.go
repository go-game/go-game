package gfx

import "github.com/go-gl-legacy/gl"

// Rectangle is a geometric shape that can be rendered with gfx.Render.
type Rectangle struct {
	X1     float64
	Y1     float64
	X2     float64
	Y2     float64
	Filled bool
}

func (r *Rectangle) render(p *Params) {
	gl.Color4d(p.R, p.G, p.B, p.A)
	RenderPolygon(r.Filled, r.X1, r.Y1, r.X1, r.Y2, r.X2, r.Y2, r.X2, r.Y1)
}
