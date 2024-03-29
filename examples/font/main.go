//go:build example
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
var p *gfx.Params

func main() {
	gameState := &game.State{
		OnInit:    onInit,
		OnRender:  onRender,
		OnCleanup: onCleanup,
		OnKeyDown: onKeyDown,
	}

	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	window, err := desktop.OpenWindow(mode)
	if err != nil {
		panic(err)
	}

	window.Run(gameState)
}

func onRender() {
	gfx.Clear()
	gfx.Render(image, p)
}

func onInit() {
	var err error
	font, err = gfx.NewFont("assets/OpenSans-Regular.ttf", 32)
	if err != nil {
		panic(err)
	}
	font.AntiAliased = true

	image, err = font.Render("Hello, Go-Game!")
	if err != nil {
		panic(err)
	}

	p = gfx.NewParams()
	p.X = 10
	p.Y = 10
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
