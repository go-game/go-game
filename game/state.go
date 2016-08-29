package game

import (
	"time"

	"git.mbuechmann.com/go-game/mouse"
)

// State is used as a scaffold to generate a new game.
type State struct {
	InitFunc          func()
	RenderFunc        func()
	UpdateFunc        func(elapsed time.Duration)
	CleanupFunc       func()
	OnMouseMove       func(x, y float64)
	OnMouseButtonUp   func(b mouse.Button)
	OnMouseButtonDown func(b mouse.Button)
	OnKeyUp           func()
	OnKeyDown         func()
}
