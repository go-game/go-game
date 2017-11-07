// +build example

package main

import (
	"fmt"

	"github.com/go-game/go-game/controller"
	"github.com/go-game/go-game/desktop"
	"github.com/go-game/go-game/game"
	"github.com/go-game/go-game/gfx"
	"github.com/go-game/go-game/keys"
)

func main() {
	mode := &desktop.Mode{Width: 100, Height: 100, Fullscreen: false}
	window := desktop.OpenWindow(mode)

	window.Run(&game.State{
		OnRender:            onRender,
		OnKeyDown:           onKeyDown,
		OnControllerAdded:   onControllerAdded,
		OnControllerRemoved: onControllerRemoved,
	})
}

func onRender() {
	gfx.Clear()
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}

func onControllerAdded(c *controller.Controller) {
	fmt.Printf("Controller #%d has connected\n", c.ID)
	l := &controller.Listener{
		OnAxisMoved: func(a controller.Axis, value float64) {
			fmt.Printf("Axis #%d of controller #%d has been moved by %f\n", a, c.ID, value)
		},
		OnButtonDown: func(b controller.Button) {
			fmt.Printf("Button #%d of controller #%d has been pressed\n", b, c.ID)
		},
		OnButtonUp: func(b controller.Button) {
			fmt.Printf("Button #%d of controller #%d has been released\n", b, c.ID)
		},
	}
	c.SetListener(l)
}

func onControllerRemoved(c *controller.Controller) {
	fmt.Printf("Controller #%d has been removed\n", c.ID)
}
