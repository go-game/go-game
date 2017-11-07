// +build example

package main

import (
	"github.com/go-game/go-game/desktop"
	"github.com/go-game/go-game/game"
	"github.com/go-game/go-game/gfx"
	"github.com/go-game/go-game/keys"
)

var font *gfx.Font
var image *gfx.Image
var ro *gfx.RenderOptions

func main() {
	gameState := &game.State{
		OnInit:    onInit,
		OnRender:  onRender,
		OnCleanup: onCleanup,
		OnKeyDown: onKeyDown,
	}

	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	window := desktop.OpenWindow(mode)

	window.Run(gameState)
}

func onRender() {
	gfx.Clear()
	gfx.Render(image, ro)
}

func onInit() {
	var err error
	font, err = gfx.NewFont("assets/OpenSans-Regular.ttf", 32)
	if err != nil {
		panic(err)
	}

	image, err = font.Render("Hello, Go-Game!")
	if err != nil {
		panic(err)
	}

	ro = gfx.NewRenderOptions()
	ro.X = 10
	ro.Y = 10
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}

func onCleanup() {
	font.Delete()
	image.Delete()
}
