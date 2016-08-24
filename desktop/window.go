package desktop

import (
  "time"
  "github.com/go-gl/glfw/v3.1/glfw"
  "github.com/go-gl/gl/v2.1/gl"
  "git.mbuechmann.com/go-game/game"
)

// CurrentWindow is the current window
var CurrentWindow *Window

// Mode represents the resolution of a window
type Mode struct {
  Width int
  Height int
  Fullscreen bool
}

// Window is the os application frame where all the stuff will  happen
type Window struct {
  mode *Mode
  state *game.State
  GlfwWindow *glfw.Window
}

// FullscreenModes returns an array of all available fullscreen modes
func FullscreenModes() []*Mode {
  modes := []*Mode{}

  monitor := glfw.GetPrimaryMonitor()
  videoModes := monitor.GetVideoModes()
  for _, m := range(videoModes) {
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

  CurrentWindow = &Window{mode: m, state: s}

  err = CurrentWindow.initGlfwWindow()
  if err != nil {
    panic(err)
  }

  err = CurrentWindow.initGL()
  if err != nil {
    panic(err)
  }

  return CurrentWindow
}

// Run starts the main game loop
func (w *Window) Run() {
  w.state.InitFunc()
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
  Exit()
}

// Exit closes the game and cleans up
func Exit() {
  CurrentWindow.state.CleanupFunc()
  CurrentWindow.GlfwWindow.SetShouldClose(true)
  glfw.Terminate()
}

func (w *Window) initGlfwWindow() (err error) {
  var monitor *glfw.Monitor

  if w.mode.Fullscreen {
    monitor = glfw.GetPrimaryMonitor()
  }

  CurrentWindow.GlfwWindow, err = glfw.CreateWindow(w.mode.Width, w.mode.Height, "", monitor, nil)
  if err != nil {
    return
  }

  w.GlfwWindow.MakeContextCurrent()
  w.GlfwWindow.SetInputMode(glfw.CursorMode, glfw.CursorHidden)

  return
}

func (w *Window) initGL() (err error) {
	if err = gl.Init(); err != nil {
		return
	}

	gl.Enable(gl.TEXTURE_2D)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.ClearColor(0.1, 0.1, 0.1, 0.0)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
  PixelSize := 1
	gl.Ortho(0, float64(w.mode.Width/PixelSize), 0, float64(w.mode.Height/PixelSize), -1, 1)
	var width, height = w.GlfwWindow.GetFramebufferSize()
	fX, fY := int32(width/w.mode.Width), int32(height/w.mode.Height)
	gl.Viewport(0, 0, fX*int32(w.mode.Width), fY*int32(w.mode.Height))

	gl.MatrixMode(gl.MODELVIEW)

  return
}
