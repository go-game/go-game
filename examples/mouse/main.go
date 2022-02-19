//go:build example
// +build example

package main

import (
	"github.com/go-game/go-game/desktop"
	"github.com/go-game/go-game/game"
	"github.com/go-game/go-game/gfx"
	"github.com/go-game/go-game/keys"
)

const pixelSize = 4

var (
	image *gfx.Image
	x, y  float64
)

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

func onMouseMove(mx, my int32) {
	x = float64(mx) / pixelSize
	y = float64(my) / pixelSize
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}

func onRender() {
	gfx.Clear()
	gfx.RenderXY(image, x, y)
}

func onInit() {
	gfx.SetPixelSize(pixelSize)
	var err error
	if image, err = gfx.NewImage("assets/heart.png"); err != nil {
		panic(err)
	}
}

func cleanupGame() {
	image.Delete()
}
