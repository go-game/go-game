package desktop

import (
	"fmt"
	"time"

	"git.mbuechmann.com/go-game/controller"
	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/gfx"
	"git.mbuechmann.com/go-game/keys"
	"git.mbuechmann.com/go-game/mouse"
	"github.com/veandco/go-sdl2/sdl"
)

var window *Window
var running = true

func init() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
}

// CurrentMode returns the mode that is currently active.
func CurrentMode() *Mode {
	mode := &sdl.DisplayMode{}
	sdl.GetCurrentDisplayMode(0, mode)
	fullscreen := false
	if window != nil && window.SDLWindow != nil {
		flags := window.SDLWindow.GetFlags()
		fullscreen = flags&sdl.WINDOW_FULLSCREEN == 1
	}

	return &Mode{Width: int(mode.W), Height: int(mode.H), Fullscreen: fullscreen}
}

// Window is the os application frame where all the stuff will  happen.
type Window struct {
	mode      *Mode
	SDLWindow *sdl.Window
	GLContext sdl.GLContext
}

// OpenWindow creates a new window on the main monitor.
func OpenWindow(m *Mode) *Window {
	options := sdl.WINDOW_OPENGL
	if m.Fullscreen {
		options |= sdl.WINDOW_FULLSCREEN
	}
	sdlWindow, err := sdl.CreateWindow("", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, m.Width, m.Height, uint32(options))
	if err != nil {
		panic(err)
	}

	context, err := sdl.GL_CreateContext(sdlWindow)
	if err != nil {
		panic(err)
	}

	window = &Window{mode: m, SDLWindow: sdlWindow, GLContext: context}

	gfx.SetArea(m.Width, m.Height)
	mouse.Hide()

	return window
}

// Run starts the main game loop for the given game state by invocing all defined callbacks in the given game state.
func (w *Window) Run(state *game.State) {
	if state.OnInit != nil {
		state.OnInit()
	}
	last := time.Now()

	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				Exit()
			case *sdl.MouseMotionEvent:
				if state.OnMouseMove != nil {
					state.OnMouseMove(t.X, t.Y)
				}
			case *sdl.MouseButtonEvent:
				if state.OnMouseButtonDown != nil && t.State == 1 {
					state.OnMouseButtonDown(mouse.Button(t.Which), float32(t.X), float32(t.Y))
				}
				if state.OnMouseButtonUp != nil && t.State == 0 {
					state.OnMouseButtonUp(mouse.Button(t.Which), float32(t.X), float32(t.Y))
				}
			case *sdl.MouseWheelEvent:
				if state.OnMouseWheel != nil {
					state.OnMouseWheel(t.X, t.Y)
				}
			case *sdl.KeyDownEvent:
				if state.OnKeyDown != nil {
					if t.Repeat == 0 {
						state.OnKeyDown(keys.Key(t.Keysym.Sym))
					}
				}
			case *sdl.KeyUpEvent:
				if state.OnKeyUp != nil {
					state.OnKeyUp(keys.Key(t.Keysym.Sym))
				}
			case *sdl.ControllerButtonEvent:
				if t.State == 1 {
					fmt.Printf("Button %d of controller %d was pressed\n", t.Button, t.Which)
				} else {
					fmt.Printf("Button %d of controller %d was released\n", t.Button, t.Which)
				}
			case *sdl.ControllerAxisEvent:
				fmt.Printf("Axis %d of controller %d moved to %d\n", t.Axis, t.Which, t.Value)
			case *sdl.ControllerDeviceEvent:
				if t.Type == sdl.CONTROLLERDEVICEADDED {
					if state.OnControllerAdded != nil {
						ctrl := controller.Open(int(t.Which))
						state.OnControllerAdded(ctrl)
					}
				}
				if t.Type == sdl.CONTROLLERDEVICEREMOVED {
					fmt.Printf("Controller %d removed\n", t.Which)
				}
				if t.Type == sdl.CONTROLLERDEVICEREMAPPED {
					fmt.Printf("Controller %d remapped\n", t.Which)
				}
			}
		}

		if state.OnUpdate != nil {
			delta := time.Since(last)
			last = time.Now()
			state.OnUpdate(delta)
		}
		if state.OnRender != nil {
			state.OnRender()
		}
		sdl.GL_SwapWindow(window.SDLWindow)
	}
	if state.OnCleanup != nil {
		state.OnCleanup()
	}
}

// Exit closes the game and calls the cleanup callbacks.
func Exit() {
	// TODO: Call cleanup of all packages

	sdl.GL_DeleteContext(window.GLContext)
	window.SDLWindow.Destroy()
	sdl.Quit()
	running = false
}
