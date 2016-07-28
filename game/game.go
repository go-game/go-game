package game

import (
	"runtime"
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

var Window *glfw.Window

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

// Game holds some configuration and a State, whose callbacks will be called
type Game struct {
	Fullscreen   bool
	WindowWidth  int
	WindowHeight int
	PixelSize    int
	Title        string
	*State
	Window *glfw.Window
}

// Run starts the game
func (g *Game) Run() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	g.initWindow()

	g.initGL()

	g.State.InitFunc()

	last := time.Now()
	for !g.Window.ShouldClose() {
		elapsed := time.Since(last)
		last = time.Now()
		g.State.UpdateFunc(elapsed)
		g.State.RenderFunc()
		glfw.SwapInterval(1)
		g.Window.SwapBuffers()
		glfw.PollEvents()
	}

	g.State.CleanupFunc()
	glfw.Terminate()
}

func(g *Game) Close() {
	Window.SetShouldClose(true)
}

func (g *Game) initWindow() {
	if g.WindowWidth == 0 {
		g.WindowWidth = 1280
	}

	if g.WindowHeight == 0 {
		g.WindowHeight = 800
	}

	if g.PixelSize == 0 {
		g.PixelSize = 1
	}

	var monitor *glfw.Monitor
	if g.Fullscreen {
		monitor = glfw.GetPrimaryMonitor()
		var mode = monitor.GetVideoMode()
		g.WindowWidth = mode.Width
		g.WindowHeight = mode.Height
	}
	window, err := glfw.CreateWindow(g.WindowWidth, g.WindowHeight, g.Title, monitor, nil)
	if err != nil {
		panic(err)
	}
	g.Window = window
	Window = window
	window.MakeContextCurrent()
	window.SetInputMode(glfw.CursorMode, glfw.CursorHidden)
}

func (g *Game) initGL() {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	gl.Enable(gl.TEXTURE_2D)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.ClearColor(0.1, 0.1, 0.1, 0.0)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(g.WindowWidth/g.PixelSize), 0, float64(g.WindowHeight/g.PixelSize), -1, 1)
	var width, height = g.Window.GetFramebufferSize()
	fX, fY := int32(width/g.WindowWidth), int32(height/g.WindowHeight)
	gl.Viewport(0, 0, fX*int32(g.WindowWidth), fY*int32(g.WindowHeight))

	gl.MatrixMode(gl.MODELVIEW)
}
