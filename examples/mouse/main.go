package main

import (
	"time"

	"git.mbuechmann.com/go-game/desktop"
	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/gfx"
	"git.mbuechmann.com/go-game/keys"
)

var texture *gfx.Texture
var posX, posY float32 = 0, 0

func main() {
	gameState := &game.State{
		InitFunc:    initGame,
		RenderFunc:  render,
		UpdateFunc:  logic,
		CleanupFunc: cleanupGame,
		OnMouseMove: onMouseMove,
	}

	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	desktop.OpenWindow(mode)
	gfx.SetPixelSize(4)

	desktop.Run(gameState)
}

func onMouseMove(x, y float64) {
	posX = float32(x)
	posY = float32(y)
}

func logic(delta time.Duration) {
	if keys.Down("esc") {
		desktop.Exit()
	}
}

func render() {
	gfx.Clear()
	texture.Render(posX, posY)
}

func initGame() {
	texture = gfx.NewTexture("assets/heart.png")
}

func cleanupGame() {
	texture.Delete()
}
