// +build example

package main

import (
	"github.com/go-game/go-game/desktop"
	"github.com/go-game/go-game/game"
	"github.com/go-game/go-game/gfx"
	"github.com/go-game/go-game/keys"
)

var params = gfx.NewParams()
var rect = &gfx.Rectangle{}

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

	rect.Y1 = 10.0
	rect.Y2 = 100.0
	for i := 0; i < 12.0; i++ {
		rect.Filled = (i%2 == 0)
		rect.X1 = 10.0 + float64(i)*100.0
		rect.X2 = 100.0 + float64(i)*100.0
		gfx.Render(rect, params)
	}

	rect.Y1 = 110.0
	rect.Y2 = 300.0
	for i := 0; i < 12.0; i += 2 {
		rect.Filled = (i%4 == 0)
		rect.X1 = 10.0 + float64(i)*100.0
		rect.X2 = 200.0 + float64(i)*100.0
		gfx.Render(rect, params)
	}

	rect.Y1 = 310.0
	rect.Y2 = 700.0
	for i := 0; i < 12.0; i += 4 {
		rect.Filled = (i%8 == 0)
		rect.X1 = 10.0 + float64(i)*100.0
		rect.X2 = 400.0 + float64(i)*100.0
		gfx.Render(rect, params)
	}
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}
