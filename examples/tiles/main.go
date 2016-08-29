package main

import (
	"time"

	"git.mbuechmann.com/go-game/desktop"
	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/gfx"
	"git.mbuechmann.com/go-game/keys"
)

var tile1 *gfx.Texture
var tile2 *gfx.Texture

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
	tile1 = gfx.NewTexture("assets/tile1.png")
	tile2 = gfx.NewTexture("assets/tile2.png")
}

func render() {
	gfx.Clear()

	for x := -64; x <= 1280; x += 64 {
		for y := 0; y <= 96; y += 32 {
			tile1.Render(float32(x), float32(y))
			tile2.Render(float32(x+32), float32(y-16))
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
