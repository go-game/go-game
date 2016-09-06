package main

import (
	"git.mbuechmann.com/go-game/desktop"
	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/gfx"
	"git.mbuechmann.com/go-game/keys"
)

var image *gfx.Image
var renderOptions *gfx.RenderOptions

func main() {
	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	gfx.SetPixelSize(4)
	desktop.OpenWindow(mode)

	desktop.Run(&game.State{
		OnKeyDown:   onKeyDown,
		RenderFunc:  render,
		InitFunc:    initFunc,
		CleanupFunc: cleanup,
	})
}

func initFunc() {
	image = gfx.NewImage("assets/heart.png")
	renderOptions = gfx.NewRenderOptions()
}

func cleanup() {
	image.Delete()
}

func render() {
	gfx.Clear()

	topLimit := 300.0

	renderOptions.Y = 10.0
	for x := 10.0; x < topLimit; x += 30.0 {
		renderOptions.X = x
		renderOptions.A = x / topLimit
		image.Render(renderOptions)
	}

	renderOptions.Y = 40.0
	for x := 10.0; x < topLimit; x += 30.0 {
		renderOptions.X = x
		renderOptions.R = x / topLimit
		image.Render(renderOptions)

	}

	renderOptions.Y = 70.0
	for x := 10.0; x < topLimit; x += 30.0 {
		renderOptions.X = x
		renderOptions.G = x / topLimit
		image.Render(renderOptions)

	}

	renderOptions.Y = 100.0
	for x := 10.0; x < topLimit; x += 30.0 {
		renderOptions.X = x
		renderOptions.B = x / topLimit
		image.Render(renderOptions)

	}
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}
