//go:build example
// +build example

package main

import (
	"github.com/go-game/go-game/desktop"
	"github.com/go-game/go-game/game"
	"github.com/go-game/go-game/gfx"
	"github.com/go-game/go-game/keys"
)

const (
	x     = 100
	y     = 100
	width = 20
)

var (
	height int32
	rect   *gfx.Rectangle
)

func main() {
	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	window, err := desktop.OpenWindow(mode)
	if err != nil {
		panic(err)
	}

	window.Run(&game.State{
		OnInit:       onInit,
		OnRender:     onRender,
		OnKeyDown:    onKeyDown,
		OnMouseWheel: onMouseWheel,
	})
}

func onInit() {
	gfx.SetPixelSize(4)
	rect = &gfx.Rectangle{X1: 0, Y1: 0, X2: width, Y2: 0, Mode: gfx.NewLineMode()}
}

func onMouseWheel(_, y int32) {
	height += y * 2
	rect.Y2 = float64(height)
}

func onRender() {
	gfx.Clear()
	gfx.RenderXY(rect, x, y)
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}
