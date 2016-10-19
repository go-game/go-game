// +build example

package main

import (
	"git.mbuechmann.com/go-game/desktop"
	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/gfx"
	"git.mbuechmann.com/go-game/keys"
)

const (
	Y     = 100
	X     = 100
	WIDTH = 20
)

var height float64

func main() {
	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	window := desktop.OpenWindow(mode)
	gfx.SetPixelSize(4)

	window.Run(&game.State{
		OnRender:     onRender,
		OnKeyDown:    onKeyDown,
		OnMouseWheel: onMouseWheel,
	})
}

func onMouseWheel(x, y float64) {
	height += y * 2
}

func onRender() {
	gfx.Clear()
	gfx.RenderRectangle(X, Y, X+WIDTH, Y+height)
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}
