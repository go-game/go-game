package main

import (
	"time"

	"git.mbuechmann.com/go-game/base"
	"git.mbuechmann.com/go-game/sprites"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

func main() {
	gameState := &base.GameState{
		InitFunc:    initGame,
		RenderFunc:  render,
		UpdateFunc:  logic,
		CleanupFunc: cleanupGame,
		KeyHandler:  onKey,
	}

	// modes := monitor.GetVideoModes()
	// for _, mode := range modes {
	// 	fmt.Printf("%dx%d\n", mode.Width, mode.Height)
	// }
	game := &base.Game{
		GameState:    gameState,
		Fullscreen:   true,
		WindowWidth:  1920,
		WindowHeihgt: 1080,
		PixelSize:    4,
		Title:        "Heart",
	}
	game.Run()
}

func logic(elapsed time.Duration) {
	heart.Update(float64(elapsed))
}

func render() {
	gl.LoadIdentity()
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	heart.Render()
}

var heart *sprites.Heart

func initGame() {
	heart = sprites.NewHeart()
}

func cleanupGame() {
	heart.Delete()
}

var vX float32
var vY float32

func onKey(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		w.SetShouldClose(true)
	}

	if key == glfw.KeyUp {
		if action == glfw.Press {
			vY++
		}
		if action == glfw.Release {
			vY--
		}
	}
	if key == glfw.KeyDown {
		if action == glfw.Press {
			vY--
		}
		if action == glfw.Release {
			vY++
		}
	}
	if key == glfw.KeyRight {
		if action == glfw.Press {
			vX++
		}
		if action == glfw.Release {
			vX--
		}
	}
	if key == glfw.KeyLeft {
		if action == glfw.Press {
			vX--
		}
		if action == glfw.Release {
			vX++
		}
	}

	heart.SetDirection(vX, vY)
}
