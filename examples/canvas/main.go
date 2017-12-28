// +build example

package main

import (
	"github.com/go-game/go-game/desktop"
	"github.com/go-game/go-game/game"
	"github.com/go-game/go-game/gfx"
	"github.com/go-game/go-game/keys"
)

var heart *gfx.Image
var grey *gfx.Image
var imageParams *gfx.Params

var canvas *gfx.Canvas
var canvasParams *gfx.Params

func main() {
	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	window := desktop.OpenWindow(mode)

	gfx.SetPixelSize(4)

	window.Run(&game.State{
		OnInit:    onInit,
		OnCleanup: onCleanup,
		OnRender:  onRender,
		OnKeyDown: onKeyDown,
	})
}

func onInit() {
	gfx.SetClearColor(0.2, 0.2, 0.2)
	heart = gfx.NewImage("assets/heart.png")
	grey = gfx.NewImage("assets/grey.png")

	var err error
	canvas, err = gfx.NewCanvas(64, 64)
	if err != nil {
		panic(err)
	}

	imageParams = gfx.NewParams()
	imageParams.X = 0
	imageParams.Y = 0
	canvas.Render(heart, imageParams)

	imageParams.X = 20
	imageParams.Y = 0
	canvas.Render(heart, imageParams)

	canvasParams = gfx.NewParams()
}

func onCleanup() {
	heart.Delete()
	grey.Delete()
	canvas.Delete()
}

func onRender() {
	gfx.Clear()

	imageParams.X = 200
	imageParams.Y = 92
	gfx.Render(grey, imageParams)

	canvasParams.X = 100
	canvasParams.Y = 100
	gfx.Render(canvas, canvasParams)

	canvasParams.X = 200
	canvasParams.Y = 100
	gfx.Render(canvas, canvasParams)
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}
