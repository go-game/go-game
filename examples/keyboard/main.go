package main

import (
	"time"

	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/gfx"
	"git.mbuechmann.com/go-game/examples/keyboard/sprites"
	"git.mbuechmann.com/go-game/keys"
)

var heart *sprites.Heart
var vX float32
var vY float32
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
		Title:     "Keyboard-Movement",
	}
	demo.Run()
}

func logic(elapsed time.Duration) {
	if keys.Down("esc") {
		demo.Close()
	}

	vY = 0
	vX = 0
	if keys.Down("up") {
		vY++
	}
	if keys.Down("down") {
		vY--
	}
	if keys.Down("right") {
		vX++
	}
	if keys.Down("left") {
		vX--
	}
	heart.SetDirection(vX, vY)

	heart.Update(float64(elapsed))
}

func render() {
	gfx.Clear()

	heart.Render()
}

func initGame() {
	heart = sprites.NewHeart()
}

func cleanupGame() {
	heart.Delete()
}
