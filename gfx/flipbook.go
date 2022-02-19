package gfx

import (
	"time"
)

// OnFlip is a callback that gets triggered when a page is changed.
// The param page is the page number of the new page.
type OnFlip func(page int)

// Page represents one part of a FlipBook animation. May have nil as renderer for empty pages.
type Page struct {
	Duration time.Duration
	Renderer renderer
}

// FlipBook represents a timed sequence of renderers.
type FlipBook struct {
	pages     []*Page
	looping   bool
	elapsed   time.Duration
	current   int
	callbacks []OnFlip
}

// NewFlipBook returns a pointer to a new FlipBook.
func NewFlipBook(l bool, pages ...*Page) *FlipBook {
	return &FlipBook{
		pages:     pages,
		looping:   l,
		elapsed:   0,
		current:   0,
		callbacks: []OnFlip{},
	}
}

// Update updates increases the elapsed time and sets the current renderer for rendering.
func (fb *FlipBook) Update(delta time.Duration) {
	if fb.Finished() {
		return
	}

	fb.elapsed += delta

	for fb.elapsed >= fb.currentPage().Duration {
		fb.elapsed -= fb.currentPage().Duration

		fb.current++
		for _, cb := range fb.callbacks {
			cb(fb.current)
		}
		if fb.looping && fb.current >= len(fb.pages) {
			fb.current = 0
		}
	}
}

// Finished indicates if the FlipBook is at the end
func (fb *FlipBook) Finished() bool {
	if fb.looping {
		return false
	}
	return fb.current >= len(fb.pages)
}

// Reset resets the FlipBook to the first frame
func (fb *FlipBook) Reset() {
	fb.current = 0
}

// AddPageListener adds a callback that will be called when the page is changed.
func (fb *FlipBook) AddPageListener(of OnFlip) {
	fb.callbacks = append(fb.callbacks, of)
}

// ClearPageListeners removes all callback functions.
func (fb *FlipBook) ClearPageListeners() {
	fb.callbacks = []OnFlip{}
}

func (fb *FlipBook) currentPage() *Page {
	i := fb.current
	if len(fb.pages)-1 < i {
		i = len(fb.pages) - 1
	}
	return fb.pages[i]
}

func (fb *FlipBook) render() {
	if fb.currentPage().Renderer != nil {
		fb.currentPage().Renderer.render()
	}
}
