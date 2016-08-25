package desktop

import (
	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/gfx"
	"git.mbuechmann.com/go-game/keys"
	"github.com/go-gl/glfw/v3.1/glfw"
	"time"
)

var window *Window

// Mode represents the resolution of a window
type Mode struct {
	Width      int
	Height     int
	Fullscreen bool
}

// Window is the os application frame where all the stuff will  happen
type Window struct {
	mode       *Mode
	state      *game.State
	GlfwWindow *glfw.Window
}

// FullscreenModes returns an array of all available fullscreen modes
func FullscreenModes() []*Mode {
	modes := []*Mode{}

	monitor := glfw.GetPrimaryMonitor()
	videoModes := monitor.GetVideoModes()
	for _, m := range videoModes {
		modes = append(modes, &Mode{Width: m.Width, Height: m.Height, Fullscreen: true})
	}

	return modes
}

// OpenWindow creates a new window on the main monitor
func OpenWindow(m *Mode, s *game.State) *Window {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	window = &Window{mode: m, state: s}

	err = window.initGlfwWindow()
	if err != nil {
		panic(err)
	}

	gfx.SetArea(m.Width, m.Height)
	keys.SetGlfwWindow(window.GlfwWindow)

	return window
}

// Run starts the main game loop
func (w *Window) Run() {
	if w.state != nil {
		w.state.InitFunc()
	}
	last := time.Now()
	for !w.GlfwWindow.ShouldClose() {
		if w.state != nil {
			elapsed := time.Since(last)
			last = time.Now()
			w.state.UpdateFunc(elapsed)
			w.state.RenderFunc()
		}
		glfw.SwapInterval(1)
		w.GlfwWindow.SwapBuffers()
		glfw.PollEvents()
	}
	if w.state != nil {
		window.state.CleanupFunc()
	}
	glfw.Terminate()
}

// Exit closes the game and cleans up
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
	w.GlfwWindow.SetInputMode(glfw.CursorMode, glfw.CursorHidden)

	return
}
