//go:build example
// +build example

package main

import (
	"github.com/go-game/go-game/desktop"
	"github.com/go-game/go-game/game"
	"github.com/go-game/go-game/gfx"
	"github.com/go-game/go-game/keys"
)

const rows = 3
const cols = 12

var line *gfx.Line

func main() {
	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	window, err := desktop.OpenWindow(mode)
	if err != nil {
		panic(err)
	}

	window.Run(&game.State{
		OnInit:    onInit,
		OnRender:  onRender,
		OnKeyDown: onKeyDown,
	})
}

func onInit() {
	gfx.SetClearColor(0.2, 0.2, 0.2)
	line, _ = gfx.NewLine(gfx.NewLineMode(), 0.0, 0.0, 100.0, 150.0)
}

func onRender() {
	gfx.Clear()

	for j := 0; j < rows; j++ {
		line.Mode.Smooth = j == 2

		for i := 0; i < cols; i++ {
			x := float64(i*100) + 20.0
			y := float64(j*220) + 20.0

			line.Mode.Width = float32(i+1) * 2

			r := 1.0
			g := 1.0
			if j == 1 {
				r = float64(i) / 11
				g = 1.0 - float64(i)/11
			}

			gfx.RenderXYColor(line, x, y, r, g, 1, 1)
		}
	}
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}
