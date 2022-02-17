//go:build example
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

const pixelSize = 4

func main() {
	gameState := &game.State{
		OnInit:      onInit,
		OnRender:    onRender,
		OnCleanup:   cleanupGame,
		OnMouseMove: onMouseMove,
		OnKeyDown:   onKeyDown,
	}

	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	window, err := desktop.OpenWindow(mode)
	if err != nil {
		panic(err)
	}

	window.Run(gameState)
}

func onMouseMove(x, y int32) {
	params.X = float64(x) / pixelSize
	params.Y = float64(y) / pixelSize
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
	gfx.SetPixelSize(pixelSize)
	image = gfx.NewImage("assets/heart.png")
}

func cleanupGame() {
	image.Delete()
}
