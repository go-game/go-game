// +build example

package main

import (
	"github.com/go-game/go-game/desktop"
	"github.com/go-game/go-game/game"
	"github.com/go-game/go-game/gfx"
	"github.com/go-game/go-game/keys"
)

var image *gfx.Image
var params = gfx.NewParams()

const pixelsize = 4

func main() {
	gameState := &game.State{
		OnInit:      onInit,
		OnRender:    onRender,
		OnCleanup:   cleanupGame,
		OnMouseMove: onMouseMove,
		OnKeyDown:   onKeyDown,
	}

	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	window := desktop.OpenWindow(mode)
	gfx.SetPixelSize(pixelsize)

	window.Run(gameState)
}

func onMouseMove(x, y int32) {
	params.X = float64(x) / pixelsize
	params.Y = float64(y) / pixelsize
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}

func onRender() {
	gfx.Clear()
	gfx.Render(image, params)
}

func onInit() {
	image = gfx.NewImage("assets/heart.png")
}

func cleanupGame() {
	image.Delete()
}
