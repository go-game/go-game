package animation

import (
	"math"
	"time"

	"github.com/mbuechmann/go-game/gfx"
)

// NewTween returns a new Tween for two RenderOptions, where start and end are the two RenderOptions to be interpolated.
// offset indicates the duration after which the interpolation starts. looping indicates if the interpolation should go
// back to the beginning when the end is reached.
func NewTween(start, end *gfx.RenderOptions, duration time.Duration, offset time.Duration, looping bool) *Tween {
	return &Tween{
		start:    start,
		end:      end,
		duration: duration,
		offset:   offset,
		looping:  looping,
		tweened:  &gfx.RenderOptions{},
		progress: -offset,
	}
}

// Tween interpolates two RenderOptions over time. Use NewTween() to create a new Tween.
type Tween struct {
	start    *gfx.RenderOptions
	end      *gfx.RenderOptions
	duration time.Duration
	offset   time.Duration
	progress time.Duration
	looping  bool
	tweened  *gfx.RenderOptions
}

// GetRenderOptions returns the interpolated RenderOptions.
func (t *Tween) GetRenderOptions() *gfx.RenderOptions {
	return t.tweened
}

// Finished indicates if the animation is at the end.
func (t *Tween) Finished() bool {
	if t.looping {
		return false
	}
	return t.progress >= t.duration
}

// Update updates the tweened RenderOptions.
func (t *Tween) Update(delta time.Duration) {
	t.progress += delta

	if t.looping {
		t.progress %= t.duration
	} else {
		if t.progress > t.duration {
			t.progress = t.duration
		}
	}

	f1 := 1 - math.Max(0.0, float64(t.progress))/float64(t.duration)
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
