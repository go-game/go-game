// +build example

package main

import (
	"fmt"
	"time"

	"git.mbuechmann.com/go-game/controller"
	"git.mbuechmann.com/go-game/desktop"
	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/gfx"
	"git.mbuechmann.com/go-game/keys"
)

func main() {
	mode := &desktop.Mode{Width: 100, Height: 100, Fullscreen: false}
	window := desktop.OpenWindow(mode)

	window.Run(&game.State{
		OnRender:          onRender,
		OnKeyDown:         onKeyDown,
		OnUpdate:          onUpdate,
		OnControllerAdded: onControllerAdded,
	})
}

func onRender() {
	gfx.Clear()
}

func onUpdate(delta time.Duration) {
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}

func onControllerAdded(ctrl *controller.Controller) {
	fmt.Printf("Controller #%d was added and it has the name %s", ctrl.ID, ctrl.Name)
	ctrl.Rumble(1.0, 1000)
}
