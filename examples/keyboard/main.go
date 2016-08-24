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
	w := desktop.OpenWindow(mode, gameState)

	w.Run()
}

func logic(delta time.Duration) {
	if keys.Down("esc") {
		desktop.Exit()
	}

	vY = 0
	vX = 0
	if keys.Down("up") {
		vY += speed
	}
	if keys.Down("down") {
		vY -= speed
	}
	if keys.Down("right") {
		vX += speed
	}
	if keys.Down("left") {
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
