// +build example

package main

import (
	"git.mbuechmann.com/go-game/desktop"
	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/gfx"
	"git.mbuechmann.com/go-game/keys"
)

var coordsOpen []float64 = []float64{10, 10, 100, 20, 250, 25, 230, 300, 40, 280}
var coordsClosed []float64 = []float64{310, 10, 400, 20, 550, 25, 530, 300, 340, 280}

func main() {
	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	desktop.OpenWindow(mode)

	desktop.Run(&game.State{RenderFunc: render, OnKeyDown: onKeyDown})
}

func render() {
	gfx.Clear()
	gfx.RenderPolygon(false, coordsOpen...)
	gfx.RenderPolygon(true, coordsClosed...)
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}
