// +build example

package main

import (
	"github.com/mbuechmann/go-game/desktop"
	"github.com/mbuechmann/go-game/game"
	"github.com/mbuechmann/go-game/gfx"
	"github.com/mbuechmann/go-game/keys"
)

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
	var i int
	for i = 0; i < 12.0; i++ {
		gfx.RenderRectangle((i%2 == 0), 10.0+float64(i)*100.0, 10.0, 100.0+float64(i)*100.0, 100.0)
	}

	for i = 0; i < 12.0; i += 2 {
		gfx.RenderRectangle((i%4 == 0), 10.0+float64(i)*100.0, 110.0, 200.0+float64(i)*100.0, 300.0)
	}

	for i = 0; i < 12.0; i += 4 {
		gfx.RenderRectangle((i%8 == 0), 10.0+float64(i)*100.0, 310.0, 400.0+float64(i)*100.0, 700.0)
	}
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}
