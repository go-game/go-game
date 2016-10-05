package main

import (
	"git.mbuechmann.com/go-game/audio"
	"git.mbuechmann.com/go-game/desktop"
	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/gfx"
	"git.mbuechmann.com/go-game/keys"
)

var source *audio.Source

func main() {
	gameState := &game.State{
		OnInit:    onInit,
		OnRender:  onRender,
		OnKeyDown: onKeyDown,
	}

	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	desktop.OpenWindow(mode)
	gfx.SetPixelSize(4)

	desktop.Run(gameState)

}

func onInit() {
	data, err := audio.LoadData("assets/welcome.wav")
	if err != nil {
		panic(err)
	}

	source = data.NewSource()

	source.Play()
	return

}

func onRender() {
	gfx.Clear()
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}
