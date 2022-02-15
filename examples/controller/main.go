//go:build example
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

const (
	width     = int32(148)
	height    = int32(89)
	padding   = 5
	pixelSize = 2
	stickMax  = 7
)

var images map[string]*gfx.Image
var buttonsPressed = [14]bool{}

var controllerConnected bool
var leftStickX float64
var leftStickY float64
var rightStickX float64
var rightStickY float64

func onInit() {
	images = map[string]*gfx.Image{
		"bg":          gfx.NewImage("./assets/controller/background.png"),
		"left_stick":  gfx.NewImage("./assets/controller/left_stick.png"),
		"right_stick": gfx.NewImage("./assets/controller/right_stick.png"),
		"button_0":    gfx.NewImage("./assets/controller/button_0.png"),
		"button_1":    gfx.NewImage("./assets/controller/button_1.png"),
		"button_2":    gfx.NewImage("./assets/controller/button_2.png"),
		"button_3":    gfx.NewImage("./assets/controller/button_3.png"),
	}

	gfx.SetClearColor(0.8, 0.8, 0.9)
	gfx.SetPixelSize(pixelSize)
}

func main() {
	mode := &desktop.Mode{Width: (width + padding*2) * pixelSize, Height: (height + padding*2) * pixelSize, Fullscreen: false}
	window := desktop.OpenWindow(mode)

	window.Run(&game.State{
		OnInit:              onInit,
		OnRender:            onRender,
		OnKeyDown:           onKeyDown,
		OnControllerAdded:   onControllerAdded,
		OnControllerRemoved: onControllerRemoved,
	})
}

func onRender() {
	gfx.Clear()

	p := gfx.NewParams()
	p.X = padding
	p.Y = padding
	if !controllerConnected {
		p.A = 0.5
	}

	gfx.Render(images["bg"], p)

	p.X = padding + leftStickX*stickMax
	p.Y = padding + leftStickY*stickMax
	gfx.Render(images["left_stick"], p)

	p.X = padding + rightStickX*stickMax
	p.Y = padding + rightStickY*stickMax
	gfx.Render(images["right_stick"], p)

	p.X = padding
	p.Y = padding
	for i := 0; i < 4; i++ {
		if buttonsPressed[i] {
			name := fmt.Sprintf("button_%d", i)
			gfx.Render(images[name], p)
		}
	}
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}

func onControllerAdded(c *controller.Controller) {
	controllerConnected = true
	l := &controller.Listener{
		OnAxisMoved: func(a controller.Axis, value float64) {
			switch a {
			case 0:
				leftStickX = value
			case 1:
				leftStickY = value
			case 2:
				rightStickX = value
			case 3:
				rightStickY = value
			}
			fmt.Printf("Axis #%d of controller #%d has been moved by %f\n", a, c.ID, value)
		},
		OnButtonDown: func(b controller.Button) {
			fmt.Printf("Button #%d of controller #%d has been pressed\n", b, c.ID)
			buttonsPressed[b] = true
		},
		OnButtonUp: func(b controller.Button) {
			fmt.Printf("Button #%d of controller #%d has been released\n", b, c.ID)
			buttonsPressed[b] = false
		},
	}
	c.SetListener(l)
}

func onControllerRemoved(c *controller.Controller) {
	fmt.Printf("Controller #%d has been removed\n", c.ID)
}
