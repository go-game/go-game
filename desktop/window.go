package desktop

import (
	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/gfx"
	"git.mbuechmann.com/go-game/keys"
	"github.com/go-gl/glfw/v3.1/glfw"
	"time"
)

var window *Window

func init()  {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
}

// Mode represents the resolution of a window and whether it is fullscreen
type Mode struct {
	Width      int
	Height     int
	Fullscreen bool
}

// Window is the os application frame where all the stuff will  happen
type Window struct {
	mode       *Mode
	GlfwWindow *glfw.Window
}

// FullscreenModes returns an array of all available fullscreen modes
func FullscreenModes() []*Mode {
	monitor := glfw.GetPrimaryMonitor()
	videoModes := monitor.GetVideoModes()

	var modes  = make([]*Mode, len(videoModes))
	for i, m := range videoModes {
		// modes = append(modes, &Mode{Width: m.Width, Height: m.Height, Fullscreen: true})
		modes[i] = &Mode{Width: m.Width, Height: m.Height, Fullscreen: true}
	}

	return modes
}

// OpenWindow creates a new window on the main monitor
func OpenWindow(m *Mode) *Window {
	window = &Window{mode: m}

	err := window.initGlfwWindow()
	if err != nil {
		panic(err)
	}

	gfx.SetArea(m.Width, m.Height)
	keys.SetGlfwWindow(window.GlfwWindow)

	return window
}

// Run starts the main game loop
func Run(state *game.State) {
	if window == nil {
		panic("No open window for game state. Call OpenWindow() first")
	}

	if state.InitFunc != nil {
		state.InitFunc()
	}
	last := time.Now()
	for !window.GlfwWindow.ShouldClose() {
		if state.UpdateFunc != nil {
			elapsed := time.Since(last)
			last = time.Now()
			state.UpdateFunc(elapsed)
		}
		if state.RenderFunc != nil {
			state.RenderFunc()
		}
		glfw.SwapInterval(1)
		window.GlfwWindow.SwapBuffers()
		glfw.PollEvents()
	}
	if state.CleanupFunc != nil {
		state.CleanupFunc()
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
