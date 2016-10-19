// +build example

package main

import (
	"time"

	"git.mbuechmann.com/go-game/desktop"
	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/gfx"
	"git.mbuechmann.com/go-game/keys"
)

var tile1 *gfx.Image
var tile2 *gfx.Image

func main() {
	gameState := &game.State{
		OnInit:    onInit,
		OnRender:  onRender,
		OnUpdate:  onUpdate,
		OnCleanup: onCleanup,
	}

	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	window := desktop.OpenWindow(mode)

	window.Run(gameState)
}

func onInit() {
	tile1 = gfx.NewImage("assets/tile1.png")
	tile2 = gfx.NewImage("assets/tile2.png")
}

func onRender() {
	gfx.Clear()

	renderOptions1 := gfx.NewRenderOptions()
	renderOptions2 := gfx.NewRenderOptions()

	for x := -64.0; x <= 1280; x += 64 {
		for y := 0.0; y <= 96; y += 32 {
			renderOptions1.X = x
			renderOptions1.Y = y
			renderOptions2.X = x + 32
			renderOptions2.Y = y - 16

			gfx.Render(tile1, renderOptions1)
			gfx.Render(tile2, renderOptions2)
		}
	}
}

func onUpdate(elapsed time.Duration) {
	if keys.IsDown(keys.Esc) {
		desktop.Exit()
	}
}

func onCleanup() {
	tile1.Delete()
	tile2.Delete()
}
