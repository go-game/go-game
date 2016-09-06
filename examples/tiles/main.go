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
		InitFunc:    initGame,
		RenderFunc:  render,
		UpdateFunc:  logic,
		CleanupFunc: cleanupGame,
	}

	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	desktop.OpenWindow(mode)

	desktop.Run(gameState)
}

func initGame() {
	tile1 = gfx.NewImage("assets/tile1.png")
	tile2 = gfx.NewImage("assets/tile2.png")
}

func render() {
	gfx.Clear()

	for x := -64.0; x <= 1280; x += 64 {
		for y := 0.0; y <= 96; y += 32 {
			tile1.Render(x, y)
			tile2.Render(x+32, y-16)
		}
	}
}

func logic(elapsed time.Duration) {
	if keys.IsDown(keys.Esc) {
		desktop.Exit()
	}
}

func cleanupGame() {
	tile1.Delete()
	tile2.Delete()
}
