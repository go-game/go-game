// +build example

package main

import (
	"fmt"
	"time"

	"git.mbuechmann.com/go-game/desktop"
)

func main() {
	mode := desktop.CurrentMode()
	printMode(mode)

	fmt.Println("Switching to fullscreen in 3 seconds")
	time.Sleep(3 * time.Second)
	m := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: true}
	window := desktop.OpenWindow(m)

	mode = desktop.CurrentMode()
	printMode(mode)

	desktop.Exit()
}

func printMode(mode *desktop.Mode) {
	full := "no-fullscreen"
	if mode.Fullscreen {
		full = "fullscreen"
	}
	fmt.Printf("Your current Mode: %d x %d, in %s mode\n", mode.Width, mode.Height, full)
}
