//go:build example
// +build example

package main

import (
	"github.com/go-game/go-game/desktop"
	"github.com/go-game/go-game/game"
	"github.com/go-game/go-game/gfx"
	"github.com/go-game/go-game/keys"
)

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

	for i := 0; i < 12; i++ {
		circle, _ := gfx.NewCircle(40, (i+2)*2, false)

		x := 100.0 * float64(i+1)
		y := 100.0
		gfx.RenderXY(circle, x, y)
	}

	for i := 0; i < 6; i++ {
		circle, _ := gfx.NewCircle(80, (i+2)*5, false)

		x := 200.0 * (float64(i) + 0.5)
		y := 300.0
		gfx.RenderXY(circle, x, y)
	}

	for i := 0; i < 6; i++ {
		circle, _ := gfx.NewCircle(80, (i+2)*5, false)
		circle.Mode.Width = 2
		circle.Mode.Smooth = true

		x := 200.0 * (float64(i) + 0.5)
		y := 550.0
		gfx.RenderXY(circle, x, y)
	}
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}
