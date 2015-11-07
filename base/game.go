package base

import (
	"runtime"
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

// Game holds a GameState, whose callbacks will be called
type Game struct {
	Fullscreen bool
	*GameState
}

// Run starts the game
func (g *Game) Run() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	width := 1920
	height := 1080
	title := "Heart"
	var monitor *glfw.Monitor = nil
	if g.Fullscreen {
		monitor = glfw.GetPrimaryMonitor()
	}
	// modes := monitor.GetVideoModes()
	// for _, mode := range modes {
	// 	fmt.Printf("%dx%d\n", mode.Width, mode.Height)
	// }
	window, err := glfw.CreateWindow(width, height, title, monitor, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	if g.KeyHandler != nil {
		window.SetKeyCallback(g.onKey)
	}
	window.SetInputMode(glfw.CursorMode, glfw.CursorHidden)

	if err := gl.Init(); err != nil {
		panic(err)
	}

	g.initGL(width, height)

	g.GameState.InitFunc()

	last := time.Now()
	for !window.ShouldClose() {
		elapsed := time.Since(last)
		last = time.Now()
		g.GameState.UpdateFunc(elapsed)
		g.GameState.RenderFunc()
		glfw.SwapInterval(1)
		window.SwapBuffers()
		glfw.PollEvents()
	}

	g.GameState.CleanupFunc()
	glfw.Terminate()
}

func (g *Game) onKey(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	g.GameState.KeyHandler(w, key, scancode, action, mods)
}

func (g *Game) initGL(width, height int) {
	gl.Enable(gl.TEXTURE_2D)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.ClearColor(0.1, 0.1, 0.1, 0.0)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(width/4), 0, float64(height/4), -1, 1)
	gl.Viewport(0, 0, int32(width), int32(height))

	gl.MatrixMode(gl.MODELVIEW)
}
