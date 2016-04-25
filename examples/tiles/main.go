package main

import (
	"time"

	"git.mbuechmann.com/go-game/base"
	"git.mbuechmann.com/go-game/gfx"
)

var tile1 *gfx.Texture
var tile2 *gfx.Texture

func main() {
	gameState := &base.GameState{
		InitFunc:    initGame,
		RenderFunc:  render,
		UpdateFunc:  logic,
		CleanupFunc: cleanupGame,
	}

	game := &base.Game{
		GameState: gameState,
		PixelSize: 2,
		Title:     "Tiles",
	}
	game.Run()
}

func initGame() {
	tile1 = gfx.NewTexture("assets/tile1.png")
	tile2 = gfx.NewTexture("assets/tile2.png")
}

func render() {
	gfx.Clear()
	// gl.LoadIdentity()
	// gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	for x := -64; x <= 1280; x += 64 {
		for y := 0; y <= 96; y += 32 {
			tile1.Render(float32(x), float32(y))
			tile2.Render(float32(x+32), float32(y-16))
		}
	}
}

func logic(elapsed time.Duration) {
}

func cleanupGame() {
	tile1.Delete()
	tile2.Delete()
}
