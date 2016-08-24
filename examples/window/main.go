package main

import (
	"git.mbuechmann.com/go-game/desktop"
)

func main() {
  mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
  w := desktop.OpenWindow(mode, nil)

  w.Run()
}
