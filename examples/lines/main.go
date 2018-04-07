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

var params = gfx.NewParams()
var line *gfx.Line

func main() {
	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	window := desktop.OpenWindow(mode)
	gfx.SetClearColor(0.2, 0.2, 0.2)

	window.Run(&game.State{
		OnInit:    onInit,
		OnRender:  onRender,
		OnKeyDown: onKeyDown,
	})
}

func onInit() {
	line, _ = gfx.NewLine(gfx.NewLineMode(), 0.0, 0.0, 100.0, 150.0)
}

func onRender() {
	gfx.Clear()

	for j := 0; j < rows; j++ {
		line.Mode.Smooth = (j == 2)

		for i := 0; i < cols; i++ {
			params.X = float64(i*100) + 20.0
			params.Y = float64(j*220) + 20.0

			line.Mode.Width = float32(i+1) * 2

			params.R = 1.0
			params.G = 1.0
			if j == 1 {
				params.R = float64(i) / 11
				params.G = 1.0 - float64(i)/11
			}

			gfx.Render(line, params)
		}
	}
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}
