// +build example

package main

import (
	"git.mbuechmann.com/go-game/desktop"
	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/gfx"
	"git.mbuechmann.com/go-game/keys"
)

var heart *gfx.Image
var grey *gfx.Image
var roImage *gfx.RenderOptions
var canvas *gfx.Canvas
var roCanvas *gfx.RenderOptions

func main() {
	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	desktop.OpenWindow(mode)

	gfx.SetPixelSize(4)

	desktop.Run(&game.State{
		OnInit:    onInit,
		OnCleanup: onCleanup,
		OnRender:  onRender,
		OnKeyDown: onKeyDown,
	})
}

func onInit() {
	heart = gfx.NewImage("assets/heart.png")
	grey = gfx.NewImage("assets/grey.png")

	var err error
	canvas, err = gfx.NewCanvas(64, 64)
	if err != nil {
		panic(err)
	}

	canvas.Clear()

	roImage = gfx.NewRenderOptions()
	roImage.X = 0
	roImage.Y = 0
	canvas.Render(heart, roImage)

	roImage.X = 20
	roImage.Y = 0
	canvas.Render(heart, roImage)

	roCanvas = gfx.NewRenderOptions()
	roCanvas.X = 100
	roCanvas.Y = 100
	gfx.Render(canvas, roCanvas)
}

func onCleanup() {
	heart.Delete()
	grey.Delete()
	canvas.Delete()
}

func onRender() {
	gfx.Clear()

	roImage.X = 200
	roImage.Y = 92
	gfx.Render(grey, roImage)

	roCanvas.X = 200
	roCanvas.Y = 100
	gfx.Render(canvas, roCanvas)
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}
