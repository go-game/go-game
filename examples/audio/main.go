// +build example

package main

import (
	"git.mbuechmann.com/go-game/audio"
	"git.mbuechmann.com/go-game/desktop"
	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/gfx"
	"git.mbuechmann.com/go-game/keys"
)

var sound *audio.Sound

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
		sound.Play()
	}
}
