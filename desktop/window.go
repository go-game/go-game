package desktop

import (
  "github.com/go-gl/glfw/v3.1/glfw"
  "git.mbuechmann.com/go-game/game"
)

// Mode represents the resolution of a window
type Mode struct {
  Width int
  Height int
  Fullscreen bool
}

// Window is the os application frame where all the stuff will  happen
type Window struct {
  mode *Mode
  game *game.Game
  glfwWindow *glfw.Window
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
func OpenWindow(m *Mode, g *game.Game) *Window {
  err := glfw.Init()
  if err != nil {
    panic(err)
  }

  w := &Window{mode: m, game: g}

  var monitor *glfw.Monitor

  if m.Fullscreen {
    monitor = glfw.GetPrimaryMonitor()
  }

  w.glfwWindow, err = glfw.CreateWindow(m.Width, m.Height, g.Title, monitor, nil)
  if err != nil {
    panic(err)
  }

  w.glfwWindow.MakeContextCurrent()
  w.glfwWindow.SetInputMode(glfw.CursorMode, glfw.CursorHidden)

  return w
}

// Run starts the main game loop
func (w *Window) Run() {
  for !w.glfwWindow.ShouldClose() {
    glfw.PollEvents()
  }
  w.exit()
}

// Exit closes the window and resets the screen
func (w *Window) exit() {
  glfw.Terminate()
}
