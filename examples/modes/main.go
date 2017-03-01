// +build example

package main

import (
	"fmt"

	"git.mbuechmann.com/go-game/desktop"
)

func main() {
	fmt.Print("Available modes:\n\n")
	modes := desktop.FullscreenModes()
	for _, mode := range modes {
		fmt.Printf("%+v\n", mode)
	}
}
