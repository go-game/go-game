package gfx

// Polygon is a geometric shape that can be rendered with gfx.Render.
type Polygon struct {
	Points []float64
	Filled bool
	Mode   *LineMode
}

func (p *Polygon) render() {
	_ = renderPolygon(p.Filled, p.Mode, p.Points...)
}
