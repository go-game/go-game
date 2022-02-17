//go:build example
// +build example

package main

import (
	"github.com/go-game/go-game/desktop"
	"github.com/go-game/go-game/game"
	"github.com/go-game/go-game/gfx"
	"github.com/go-game/go-game/keys"
)

var tile1 *gfx.Image
var tile2 *gfx.Image

func main() {
	gameState := &game.State{
		OnInit:    onInit,
		OnRender:  onRender,
		OnCleanup: onCleanup,
		OnKeyDown: onKeyDown,
	}

	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	window, err := desktop.OpenWindow(mode)
	if err != nil {
		panic(err)
	}

	window.Run(gameState)
}

func onInit() {
	tile1 = gfx.NewImage("assets/tile1.png")
	tile2 = gfx.NewImage("assets/tile2.png")
}

func onRender() {
	gfx.Clear()

	params1 := gfx.NewParams()
	params2 := gfx.NewParams()

	for x := -64.0; x <= 1280; x += 64 {
		for y := 0.0; y <= 96; y += 32 {
			params1.X = x
			params1.Y = y
			params2.X = x + 32
			params2.Y = y - 16

			gfx.Render(tile1, params1)
			gfx.Render(tile2, params2)
		}
	}
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}

func onCleanup() {
	tile1.Delete()
	tile2.Delete()
}
