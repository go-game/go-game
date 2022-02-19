//go:build example
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

var circle *gfx.Circle
var circleParams *gfx.Params

func main() {
	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	window, err := desktop.OpenWindow(mode)
	if err != nil {
		panic(err)
	}

	window.Run(&game.State{
		OnInit:    onInit,
		OnCleanup: onCleanup,
		OnRender:  onRender,
		OnKeyDown: onKeyDown,
	})
}

func onInit() {
	gfx.SetPixelSize(4)
	gfx.SetClearColor(0.2, 0.2, 0.2)

	var err error
	if heart, err = gfx.NewImage("assets/heart.png"); err != nil {
		panic(err)
	}
	if grey, err = gfx.NewImage("assets/grey.png"); err != nil {
		panic(err)
	}
	if canvas, err = gfx.NewCanvas(64, 64); err != nil {
		panic(err)
	}

	circleParams = gfx.NewParams()
	circleParams.X = 10
	circleParams.Y = 10

	circle, _ = gfx.NewCircle(10.0, 10, false)
	canvas.Render(circle, circleParams)

	imageParams = gfx.NewParams()
	imageParams.X = 2
	imageParams.Y = 2
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
