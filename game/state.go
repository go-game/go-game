package game

import (
	"time"
)

// State is used as a scaffold to generate a new game.
type State struct {
	InitFunc    func()
	RenderFunc  func()
	UpdateFunc  func(elapsed time.Duration)
	CleanupFunc func()
}
