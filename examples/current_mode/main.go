//go:build example
// +build example

package main

import (
	"fmt"
	"time"

	"github.com/go-game/go-game/desktop"
)

func main() {
	mode := desktop.CurrentMode()
	printMode(mode)

	fmt.Println("Switching to fullscreen in 3 seconds for 1 second")
	time.Sleep(3 * time.Second)
	m := &desktop.Mode{Width: 1280, Height: 960, Fullscreen: true}
	_, err := desktop.OpenWindow(m)
	if err != nil {
		panic(err)
	}

	mode = desktop.CurrentMode()
	printMode(mode)
	time.Sleep(time.Second)

	desktop.Exit()
}

func printMode(mode *desktop.Mode) {
	full := "no-fullscreen"
	if mode.Fullscreen {
		full = "fullscreen"
	}
	fmt.Printf("Your current Mode: %d x %d, in %s mode\n", mode.Width, mode.Height, full)
}
