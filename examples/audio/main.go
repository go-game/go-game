//go:build example
// +build example

package main

import (
	"github.com/go-game/go-game/audio"
	"github.com/go-game/go-game/desktop"
	"github.com/go-game/go-game/game"
	"github.com/go-game/go-game/gfx"
	"github.com/go-game/go-game/keys"
)

var sound *audio.Sound

func main() {
	gameState := &game.State{
		OnInit:    onInit,
		OnRender:  onRender,
		OnKeyDown: onKeyDown,
	}

	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	window, err := desktop.OpenWindow(mode)
	if err != nil {
		panic(err)
	}
	gfx.SetPixelSize(4)

	window.Run(gameState)
}

func onInit() {
	var err error
	sound, err = audio.NewSound("assets/welcome.wav")
	if err != nil {
		panic(err)
	}
}

func onRender() {
	gfx.Clear()
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
