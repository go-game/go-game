package animation

import (
	"time"

	"git.mbuechmann.com/go-game/gfx"
)

// Page represents one part of a Flipbook animation. May have nil as renderer for empty pages.
type Page struct {
	Duration time.Duration
	Renderer gfx.Renderer
}

// Flipbook represents a timed sequence of renderers.
type Flipbook struct {
	pages   []*Page
	looping bool
	elapsed time.Duration
	current int
}

// NewFlipbook returns a pointer to a new Flipbook.
func NewFlipbook(l bool, pages ...*Page) *Flipbook {
	return &Flipbook{
		pages:   pages,
		looping: l,
		elapsed: 0,
		current: 0,
	}
}

// Update updates increases the elapsed time and sets the current renderer for rendering.
func (fb *Flipbook) Update(delta time.Duration) {
	if !fb.looping && fb.current >= len(fb.pages)-1 {
		return
	}

	fb.elapsed += delta

	for fb.elapsed >= fb.currentPage().Duration {
		fb.elapsed -= fb.currentPage().Duration
		fb.current++
		if fb.current >= len(fb.pages) {
			fb.current = 0
		}
	}
}

// CurrentRenderer returns the renderer to be currently rendered.
func (fb *Flipbook) CurrentRenderer() gfx.Renderer {
	return fb.currentPage().Renderer
}

func (fb *Flipbook) currentPage() *Page {
	return fb.pages[fb.current]
}
