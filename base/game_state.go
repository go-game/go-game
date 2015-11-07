package base

import (
	"github.com/go-gl/glfw/v3.1/glfw"
	"time"
)

// GameState is used as a scaffold to generate a new game
type GameState struct {
	InitFunc    func()
	RenderFunc  func()
	UpdateFunc  func(elapsed time.Duration)
	CleanupFunc func()
	KeyHandler  func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey)
}
