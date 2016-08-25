package main

import (
	"git.mbuechmann.com/go-game/desktop"
	"git.mbuechmann.com/go-game/game"
)

func main() {
	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	desktop.OpenWindow(mode)

	desktop.Run(&game.State{})
}
