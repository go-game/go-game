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
var params *gfx.Params

func main() {
	mode := &desktop.Mode{Width: 1280, Height: 1000, Fullscreen: false}
	window, err := desktop.OpenWindow(mode)
	if err != nil {
		panic(err)
	}

	window.Run(&game.State{
		OnKeyDown: onKeyDown,
		OnRender:  render,
		OnInit:    onInit,
		OnCleanup: onCleanup,
	})
}

func onInit() {
	gfx.SetPixelSize(4)
	gfx.SetClearColor(0.5, 0.5, 0.5)
	var err error
	if image, err = gfx.NewImage("assets/grey.png"); err != nil {
		panic(err)
	}
	params = gfx.NewParams()
}

func onCleanup() {
	image.Delete()
}

func render() {
	gfx.Clear()

	topLimit := 300.0
	yOff := 25.0

	params.Y = 0.0
	for x := 15.0; x < topLimit; x += 30.0 {
		params.X = x
		params.A = x / topLimit
		gfx.Render(image, params)
	}

	params.Y += yOff
	for x := 15.0; x < topLimit; x += 30.0 {
		params.X = x
		params.R = x / topLimit
		gfx.Render(image, params)
	}

	params.Y += yOff
	for x := 15.0; x < topLimit; x += 30.0 {
		params.X = x
		params.G = x / topLimit
		gfx.Render(image, params)
	}

	params.Y += yOff
	for x := 15.0; x < topLimit; x += 30.0 {
		params.X = x
		params.B = x / topLimit
		gfx.Render(image, params)
	}

	params.Y += yOff
	for x := 15.0; x < topLimit; x += 30.0 {
		params.X = x
		params.R = x / topLimit
		params.G = x / topLimit
		gfx.Render(image, params)
	}

	params.Y += yOff
	for x := 15.0; x < topLimit; x += 30.0 {
		params.X = x
		params.R = x / topLimit
		params.B = x / topLimit
		gfx.Render(image, params)
	}

	params.Y += yOff
	for x := 15.0; x < topLimit; x += 30.0 {
		params.X = x
		params.G = x / topLimit
		params.B = x / topLimit
		gfx.Render(image, params)
	}

	params.Y += yOff
	for x := 15.0; x < topLimit; x += 30.0 {
		params.X = x
		params.Rot = gfx.Rotation{Angle: x * 1.2, X: 8, Y: 8}
		gfx.Render(image, params)
	}
	params.Rot = gfx.Rotation{Angle: 0, X: 8, Y: 8}

	params.Y += yOff
	scale := 1.0
	for x := 15.0; x < topLimit; x += 30.0 {
		params.X = x
		scale += 0.1
		params.Scale.Factor = scale
		params.Scale.X = 8
		params.Scale.Y = 8
		gfx.Render(image, params)
	}
	params.Scale = gfx.Scale{Factor: 1, X: 0, Y: 0}
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}
