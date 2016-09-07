// +build example

package main

import (
	"time"

	"git.mbuechmann.com/go-game/desktop"
	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/gfx"
	"git.mbuechmann.com/go-game/keys"
)

var image *gfx.Image
var renderOptions = gfx.NewRenderOptions()

const pixelsize = 4

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
	gfx.SetPixelSize(pixelsize)

	desktop.Run(gameState)
}

func onMouseMove(x, y float64) {
	renderOptions.X = x / pixelsize
	renderOptions.Y = y / pixelsize
}

func logic(delta time.Duration) {
	if keys.IsDown(keys.Esc) {
		desktop.Exit()
	}
}

func render() {
	gfx.Clear()
	image.Render(renderOptions)
}

func initGame() {
	image = gfx.NewImage("assets/heart.png")
}

func cleanupGame() {
	image.Delete()
}
