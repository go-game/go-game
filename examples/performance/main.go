// +build example

package main

import (
	"git.mbuechmann.com/go-game/desktop"
	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/gfx"
	"git.mbuechmann.com/go-game/keys"
	"git.mbuechmann.com/go-game/mouse"
	"time"
)

type heart struct {
	RenderOptions *gfx.RenderOptions
}

var hearts []*heart
var image *gfx.Image

func main() {
	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	desktop.OpenWindow(mode)

	desktop.Run(&game.State{
		InitFunc:          initGame,
		CleanupFunc:       cleanup,
		UpdateFunc:        update,
		RenderFunc:        render,
		OnKeyDown:         onKeyPress,
		OnMouseButtonDown: onMouseButtonDown,
	})
}

func initGame() {
	image = gfx.NewImage("assets/heart.png")
}

func cleanup() {
	image.Delete()
}

func render() {
	gfx.Clear()
}

func update(delta time.Duration) {

}

func onMouseButtonDown(b mouse.Button, x, y float32) {

}

func onKeyPress(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}
