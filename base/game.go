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

// RenderFunc renders the game
type RenderFunc func()

// UpdateFunc updated all entities
type UpdateFunc func(elapsed time.Duration)

// InitFunc is where all game objects are initialized
type InitFunc func()

// CleanupFunc is where all cleanup takes place before the game is finished
type CleanupFunc func()

// KeyHandler is where key inputs are handled
type KeyHandler func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey)

// Game is used as a scaffold to generate a new game
type Game struct {
	InitFunc
	RenderFunc
	UpdateFunc
	CleanupFunc
	KeyHandler
}

// Run starts the game
func (g *Game) Run() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	width := 1920
	height := 1080
	title := "Testing"
	var monitor = glfw.GetPrimaryMonitor()
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

	g.InitFunc()

	last := time.Now()
	for !window.ShouldClose() {
		elapsed := time.Since(last)
		last = time.Now()
		g.UpdateFunc(elapsed)
		g.RenderFunc()
		glfw.SwapInterval(1)
		window.SwapBuffers()
		glfw.PollEvents()
	}

	g.CleanupFunc()
	glfw.Terminate()
}

func (g *Game) onKey(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	g.KeyHandler(w, key, scancode, action, mods)
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
