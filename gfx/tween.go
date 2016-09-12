package gfx

import "time"

// NewTween returns a new Tween for two RenderOptions
func NewTween(start, end *RenderOptions, duration time.Duration, offset time.Duration, looping bool) *Tween {
	return &Tween{
		start:    start,
		end:      end,
		duration: duration,
		offset:   offset,
		looping:  looping,
		tweened:  &RenderOptions{},
	}
}

// Tween interpolates two RenderOptions over time
type Tween struct {
	start    *RenderOptions
	end      *RenderOptions
	duration time.Duration
	offset   time.Duration
	looping  bool
	tweened  *RenderOptions
}

// GetRenderOptions returns the interpolated RenderOptions
func (t *Tween) GetRenderOptions() *RenderOptions {
	return t.tweened
}

// Update updates the tweened RenderOptions
func (t *Tween) Update(elapsed time.Duration) {
	d := elapsed - t.offset

	if t.looping {
		d %= t.duration
		if d < 0 {
			d += t.duration
		}
	} else {
		if elapsed >= t.offset+t.duration {
			d = t.offset + t.duration
		}
		if elapsed < t.offset {
			d = 0
		}
	}

	f1 := 1 - float64(d)/float64(t.duration)
	f2 := 1 - f1

	t.tweened.X = t.start.X*f1 + t.end.X*f2
	t.tweened.Y = t.start.Y*f1 + t.end.Y*f2
	t.tweened.R = t.start.R*f1 + t.end.R*f2
	t.tweened.G = t.start.G*f1 + t.end.G*f2
	t.tweened.B = t.start.B*f1 + t.end.B*f2
	t.tweened.A = t.start.A*f1 + t.end.A*f2
	t.tweened.Rot.Angle = t.start.Rot.Angle*f1 + t.end.Rot.Angle*f2
	t.tweened.Rot.X = t.start.Rot.X*f1 + t.end.Rot.X*f2
	t.tweened.Rot.Y = t.start.Rot.Y*f1 + t.end.Rot.Y*f2
	t.tweened.Scale.Factor = t.start.Scale.Factor*f1 + t.end.Scale.Factor*f2
	t.tweened.Scale.X = t.start.Scale.X*f1 + t.end.Scale.X*f2
	t.tweened.Scale.Y = t.start.Scale.Y*f1 + t.end.Scale.Y*f2
}
