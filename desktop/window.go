package desktop

import (
	"time"

	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/gfx"
	"git.mbuechmann.com/go-game/keys"
	"git.mbuechmann.com/go-game/mouse"
	"github.com/go-gl/glfw/v3.1/glfw"
)

var window *Window

func init() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
}

// Mode represents the resolution of a window and whether it is fullscreen.
type Mode struct {
	Width      int
	Height     int
	Fullscreen bool
}

// Window is the os application frame where all the stuff will  happen.
type Window struct {
	mode       *Mode
	GlfwWindow *glfw.Window
}

// FullscreenModes returns an array of all available fullscreen modes.
func FullscreenModes() []*Mode {
	monitor := glfw.GetPrimaryMonitor()
	videoModes := monitor.GetVideoModes()

	var modes = make([]*Mode, len(videoModes))
	for i, m := range videoModes {
		modes[i] = &Mode{Width: m.Width, Height: m.Height, Fullscreen: true}
	}

	return modes
}

// CurrentMode returns the mode that is currently active.
func CurrentMode() *Mode {
	if window == nil {
		monitor := glfw.GetPrimaryMonitor()
		m := monitor.GetVideoMode()
		return &Mode{Width: m.Width, Height: m.Height}
	}
	return window.mode
}

// OpenWindow creates a new window on the main monitor.
func OpenWindow(m *Mode) *Window {
	window = &Window{mode: m}

	err := window.initGlfwWindow()
	if err != nil {
		panic(err)
	}

	gfx.SetArea(m.Width, m.Height)
	keys.SetGlfwWindow(window.GlfwWindow)
	mouse.SetGlfwWindow(window.GlfwWindow)
	mouse.Hide()

	return window
}

// Run starts the main game loop for the given game state by invocing all defined callbacks in the given game state.
func Run(state *game.State) {
	if window == nil {
		panic("No open window for game state. Call OpenWindow() first")
	}

	if state.OnMouseMove != nil {
		window.GlfwWindow.SetCursorPosCallback(func(w *glfw.Window, xpos float64, ypos float64) {
			state.OnMouseMove(xpos, ypos)
		})
	}

	if state.OnMouseButtonUp != nil || state.OnMouseButtonDown != nil {
		window.GlfwWindow.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
			x, y := mouse.Position()
			if state.OnMouseButtonUp != nil && action == glfw.Release {
				state.OnMouseButtonUp(mouse.Button(button), x, y)
			}
			if state.OnMouseButtonDown != nil && action == glfw.Press {
				state.OnMouseButtonDown(mouse.Button(button), x, y)
			}
		})
	}

	if state.OnMouseWheel != nil {
		window.GlfwWindow.SetScrollCallback(func(w *glfw.Window, x, y float64) {
			state.OnMouseWheel(x, y)
		})
	}

	if state.OnKeyUp != nil || state.OnKeyDown != nil {
		window.GlfwWindow.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
			if state.OnKeyUp != nil && action == glfw.Release {
				state.OnKeyUp(keys.Key(key))
			}
			if state.OnKeyDown != nil && action == glfw.Press {
				state.OnKeyDown(keys.Key(key))
			}
		})
	}

	if state.OnInit != nil {
		state.OnInit()
	}
	last := time.Now()
	for !window.GlfwWindow.ShouldClose() {
		if state.OnUpdate != nil {
			elapsed := time.Since(last)
			last = time.Now()
			state.OnUpdate(elapsed)
		}
		if state.OnRender != nil {
			state.OnRender()
		}
		glfw.SwapInterval(1)
		window.GlfwWindow.SwapBuffers()
		glfw.PollEvents()
	}
	if state.OnCleanup != nil {
		state.OnCleanup()
	}
	glfw.Terminate()
}

// Exit closes the game and calls the cleanup callbacks.
func Exit() {
	window.GlfwWindow.SetShouldClose(true)
}

func (w *Window) initGlfwWindow() (err error) {
	var monitor *glfw.Monitor

	if w.mode.Fullscreen {
		monitor = glfw.GetPrimaryMonitor()
	}

	window.GlfwWindow, err = glfw.CreateWindow(w.mode.Width, w.mode.Height, "", monitor, nil)
	if err != nil {
		return
	}

	w.GlfwWindow.MakeContextCurrent()

	return
}
