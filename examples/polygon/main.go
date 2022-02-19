//go:build example
// +build example

package main

import (
	"github.com/go-game/go-game/desktop"
	"github.com/go-game/go-game/game"
	"github.com/go-game/go-game/gfx"
	"github.com/go-game/go-game/keys"
)

var polygon = &gfx.Polygon{Points: []float64{0, 0, 90, 10, 240, 15, 220, 290, 30, 270}, Mode: gfx.NewLineMode()}

func main() {
	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	window, err := desktop.OpenWindow(mode)
	if err != nil {
		panic(err)
	}

	window.Run(&game.State{
		OnRender:  onRender,
		OnKeyDown: onKeyDown,
	})
}

func onRender() {
	gfx.Clear()

	polygon.Filled = false
	x := 10.0
	y := 10.0
	gfx.RenderXY(polygon, x, y)

	x = 300.0
	y = 10.0
	polygon.Filled = true
	gfx.RenderXY(polygon, x, y)
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}
