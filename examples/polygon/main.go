// +build example

package main

import (
	"github.com/mbuechmann/go-game/desktop"
	"github.com/mbuechmann/go-game/game"
	"github.com/mbuechmann/go-game/gfx"
	"github.com/mbuechmann/go-game/keys"
)

var coordsOpen []float64 = []float64{10, 10, 100, 20, 250, 25, 230, 300, 40, 280}
var coordsClosed []float64 = []float64{310, 10, 400, 20, 550, 25, 530, 300, 340, 280}

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
	gfx.RenderPolygon(false, coordsOpen...)
	gfx.RenderPolygon(true, coordsClosed...)
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}
