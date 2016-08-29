package main

import (
	"time"

	"git.mbuechmann.com/go-game/desktop"
	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/gfx"
	"git.mbuechmann.com/go-game/keys"
)

var texture *gfx.Texture
var vX, vY float32 = 0, 0
var posX, posY float32 = 100, 100
var speed float32 = 100

func main() {
	gameState := &game.State{
		InitFunc:    initGame,
		RenderFunc:  render,
		UpdateFunc:  logic,
		CleanupFunc: cleanupGame,
	}

	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	desktop.OpenWindow(mode)
	gfx.SetPixelSize(4)

	desktop.Run(gameState)
}

func logic(delta time.Duration) {
	if keys.IsDown(keys.Esc) {
		desktop.Exit()
	}

	vY = 0
	vX = 0
	if keys.IsDown(keys.Up) {
		vY -= speed
	}
	if keys.IsDown(keys.Down) {
		vY += speed
	}
	if keys.IsDown(keys.Right) {
		vX += speed
	}
	if keys.IsDown(keys.Left) {
		vX -= speed
	}

	var seconds = float32(float64(delta) / 1000000000)
	posX += vX * seconds
	posY += vY * seconds
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
