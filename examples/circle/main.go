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
	window := desktop.OpenWindow(mode)

	window.Run(&game.State{
		OnRender:  onRender,
		OnKeyDown: onKeyDown,
	})
}

func onRender() {
	gfx.Clear()
	for i := 0; i < 12; i++ {
		gfx.RenderCircle(100*float64(i+1), 100, 40, (i+2)*2)
	}

	for i := 0; i < 6; i++ {
		gfx.RenderCircle(200*(float64(i)+0.5), 300, 80, (i+2)*5)
	}
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}
