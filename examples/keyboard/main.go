// +build example

package main

import (
	"time"

	"github.com/go-game/go-game/desktop"
	"github.com/go-game/go-game/game"
	"github.com/go-game/go-game/gfx"
	"github.com/go-game/go-game/keys"
)

var image *gfx.Image
var vX, vY float64 = 0, 0
var renderOptions = gfx.NewRenderOptions()

const speed float64 = 100

func main() {
	gameState := &game.State{
		OnInit:    onInit,
		OnRender:  onRender,
		OnUpdate:  onUpdate,
		OnCleanup: onCleanup,
		OnKeyUp:   onKeyUp,
		OnKeyDown: onKeyDown,
	}

	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	window := desktop.OpenWindow(mode)
	gfx.SetPixelSize(4)

	window.Run(gameState)
}

func onKeyUp(k keys.Key) {
	if k == keys.Up {
		vY += speed
	}
	if k == keys.Down {
		vY -= speed
	}
	if k == keys.Right {
		vX -= speed
	}
	if k == keys.Left {
		vX += speed
	}
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}

	if k == keys.Up {
		vY -= speed
	}
	if k == keys.Down {
		vY += speed
	}
	if k == keys.Right {
		vX += speed
	}
	if k == keys.Left {
		vX -= speed
	}
}

func onUpdate(delta time.Duration) {
	var seconds = (float64(delta) / 1000000000)
	renderOptions.X += vX * seconds
	renderOptions.Y += vY * seconds
}

func onRender() {
	gfx.Clear()
	gfx.Render(image, renderOptions)
}

func onInit() {
	image = gfx.NewImage("assets/heart.png")
}

func onCleanup() {
	image.Delete()
}
