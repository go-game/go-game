package main

import (
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/mbuechmann/game/base"
	"github.com/mbuechmann/game/sprites"
)

func main() {
	game := &base.Game{
		InitFunc:    initGame,
		RenderFunc:  render,
		UpdateFunc:  logic,
		CleanupFunc: cleanupGame,
		KeyHandler:  onKey,
	}
	game.Run()
}

func logic(elapsed time.Duration) {
	heart.Update(float64(elapsed))
}

func render() {
	gl.LoadIdentity()
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	heart2.Render()
	heart.Render()
}

var heart *sprites.Heart
var heart2 *sprites.Heart

func initGame() {
	heart = sprites.NewHeart()
	heart2 = sprites.NewHeart()
}

func cleanupGame() {
	heart.Delete()
	heart2.Delete()
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
