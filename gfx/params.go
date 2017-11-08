package gfx

// NewParams returns new Params with sensible default values.
func NewParams() *Params {
	return &Params{
		R:     1.0,
		G:     1.0,
		B:     1.0,
		A:     1.0,
		Rot:   Rotation{},
		Scale: Scale{Factor: 1},
	}
}

// Params encompasses all transformations that can be done while rendering an Image.
type Params struct {
	X     float64
	Y     float64
	R     float64
	G     float64
	B     float64
	A     float64
	Rot   Rotation
	Scale Scale
}

// Rotation describes the clockwise rotation around the center at X, Y.
type Rotation struct {
	Angle float64
	X     float64
	Y     float64
}

// Scale describes a scale with a center at x, y.
type Scale struct {
	Factor float64
	X      float64
	Y      float64
}
