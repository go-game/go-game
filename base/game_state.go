package base

import (
	"time"
)

// GameState is used as a scaffold to generate a new game
type GameState struct {
	InitFunc    func()
	RenderFunc  func()
	UpdateFunc  func(elapsed time.Duration)
	CleanupFunc func()
}
