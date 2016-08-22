package main

import (
	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/desktop"
)

func main() {
  mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
  game := &game.Game{Title: "Demo"}
  w := desktop.OpenWindow(mode, game)

  w.Run()
}
