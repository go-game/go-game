package gfx

import (
	"github.com/go-gl/gl/v2.1/gl"
)

// Line is a geometric shape that can be rendered with gfx.Render.
type Line struct {
	Points []float64
	Mode   *LineMode
}

func (l *Line) render(p *Params) {
	gl.Color4d(p.R, p.G, p.B, p.A)
	renderPoints(gl.LINES, l.Mode, l.Points...)
}
