// +build example

package main

import (
	"github.com/go-game/go-game/desktop"
	"github.com/go-game/go-game/game"
	"github.com/go-game/go-game/gfx"
	"github.com/go-game/go-game/keys"
)

var params = gfx.NewParams()
var circle = &gfx.Circle{Radius: 100, Segments: 10, Filled: true}

func main() {
	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	window := desktop.OpenWindow(mode)

	window.Run(&game.State{
		OnRender:  onRender,
		OnKeyDown: onKeyDown,
	})
}

func onRender() {
	gfx.Clear()

	circle.Radius = 40
	for i := 0; i < 12; i++ {
		params.X = 100 * float64(i+1)
		params.Y = 100
		circle.Segments = (i + 2) * 2
		gfx.Render(circle, params)
	}

	circle.Radius = 80
	for i := 0; i < 6; i++ {
		params.X = 200 * (float64(i) + 0.5)
		params.Y = 300
		circle.Segments = (i + 2) * 5
		gfx.Render(circle, params)
	}
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}
