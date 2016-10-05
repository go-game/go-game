// +build example

package main

import (
	"git.mbuechmann.com/go-game/desktop"
	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/gfx"
	"git.mbuechmann.com/go-game/keys"
)

func main() {
	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	desktop.OpenWindow(mode)
	gfx.SetClearColor(0.2, 0.2, 0.2, 1.0)

	desktop.Run(&game.State{
		OnRender:  onRender,
		OnKeyDown: onKeyDown,
	})
}

func onRender() {
	gfx.Clear()

	gfx.SetSmoothLines(false)
	for i := 0.0; i < 12; i++ {
		gfx.SetLineWidth(i + 1)
		gfx.RenderLines(10+i*100, 10, 110+i*100, 200)
	}

	for i := 0.0; i < 12; i++ {
		gfx.SetLineColor(i/11, 1, 1, 1)
		gfx.SetLineWidth(i + 1)
		gfx.RenderLines(10+i*100, 230, 110+i*100, 420)
	}

	gfx.SetSmoothLines(true)
	for i := 0.0; i < 12; i++ {
		gfx.SetLineWidth(i + 1)
		gfx.RenderLines(10+i*100, 450, 110+i*100, 640)
	}
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}
