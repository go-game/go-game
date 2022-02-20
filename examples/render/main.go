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
}

func onCleanup() {
	image.Delete()
}

func render() {
	gfx.Clear()

	topLimit := 300.0
	yOff := 25.0

	y := 0.0
	for x := 15.0; x < topLimit; x += 30.0 {
		gfx.RenderXYColor(image, x, y, 1, 1, 1, x/topLimit)
	}

	y += yOff
	for x := 15.0; x < topLimit; x += 30.0 {
		r := x / topLimit
		gfx.RenderXYColor(image, x, y, r, 1, 1, 1)
	}

	y += yOff
	for x := 15.0; x < topLimit; x += 30.0 {
		g := x / topLimit
		gfx.RenderXYColor(image, x, y, 1, g, 1, 1)
	}

	y += yOff
	for x := 15.0; x < topLimit; x += 30.0 {
		b := x / topLimit
		gfx.RenderXYColor(image, x, y, 1, 1, b, 1)
	}

	y += yOff
	for x := 15.0; x < topLimit; x += 30.0 {
		r := x / topLimit
		g := x / topLimit
		gfx.RenderXYColor(image, x, y, r, g, 1, 1)
	}

	y += yOff
	for x := 15.0; x < topLimit; x += 30.0 {
		r := x / topLimit
		b := x / topLimit
		gfx.RenderXYColor(image, x, y, r, 1, b, 1)
	}

	y += yOff
	for x := 15.0; x < topLimit; x += 30.0 {
		g := x / topLimit
		b := x / topLimit
		gfx.RenderXYColor(image, x, y, 1, g, b, 1)
	}

	y += yOff
	for x := 15.0; x < topLimit; x += 30.0 {
		angle := x * 1.2
		gfx.RenderXYRot(image, x, y, 8, 8, angle)
	}

	y += yOff
	scale := 1.0
	for x := 15.0; x < topLimit; x += 30.0 {
		scale += 0.1
		gfx.RenderXYScale(image, x, y, 8, 8, scale)
	}
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}
