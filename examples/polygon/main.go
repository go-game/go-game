//go:build example
// +build example

package main

import (
	"github.com/go-game/go-game/desktop"
	"github.com/go-game/go-game/game"
	"github.com/go-game/go-game/gfx"
	"github.com/go-game/go-game/keys"
)

var params = gfx.NewParams()
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
	params.X = 10.0
	params.Y = 10.0
	gfx.Render(polygon, params)

	params.X = 300.0
	params.Y = 10.0
	polygon.Filled = true
	gfx.Render(polygon, params)
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}
