// +build example

package main

import (
	"github.com/go-game/go-game/audio"
	"github.com/go-game/go-game/desktop"
	"github.com/go-game/go-game/examples"
	"github.com/go-game/go-game/game"
	"github.com/go-game/go-game/gfx"
	"github.com/go-game/go-game/keys"
)

var sound *audio.Sound

var textImage *gfx.Image

func main() {
	gameState := &game.State{
		OnInit:    onInit,
		OnRender:  onRender,
		OnKeyDown: onKeyDown,
	}

	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	window := desktop.OpenWindow(mode)
	gfx.SetPixelSize(4)

	window.Run(gameState)
}

func onInit() {
	textImage = examples.TextImage("Press Space to play a sound")

	var err error
	sound, err = audio.NewSound("assets/welcome.wav")
	if err != nil {
		panic(err)
	}
}

func onRender() {
	gfx.Clear()
	examples.RenderImage(textImage, 10, 10)
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
	if k == keys.Space {
		err := sound.Play()
		if err != nil {
			panic(err)
		}
	}
}
