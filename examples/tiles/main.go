package main

import (
	"time"

	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/keys"
	"git.mbuechmann.com/go-game/gfx"
)

var tile1 *gfx.Texture
var tile2 *gfx.Texture
var demo *game.Game

func main() {
	gameState := &game.State{
		InitFunc:    initGame,
		RenderFunc:  render,
		UpdateFunc:  logic,
		CleanupFunc: cleanupGame,
	}

	demo = &game.Game{
		State: gameState,
		PixelSize: 2,
		Title:     "Tiles",
	}
	demo.Run()
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
	if keys.Down("esc") {
		demo.Close()
	}
}

func cleanupGame() {
	tile1.Delete()
	tile2.Delete()
}
