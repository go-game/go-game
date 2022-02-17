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

type device struct {
	x float64
	y float64

	connected      bool
	leftStickX     float64
	leftStickY     float64
	rightStickX    float64
	rightStickY    float64
	leftTrigger    float64
	rightTrigger   float64
	buttonsPressed [15]bool
}

var (
	images     map[string]*gfx.Image
	imageNames = []string{
		"background",
		"left_stick",
		"right_stick",
		"left_trigger",
		"right_trigger",
		"button_0",
		"button_1",
		"button_2",
		"button_3",
		"button_4",
		"button_5",
		"button_6",
		"button_7",
		"button_8",
		"button_9",
		"button_10",
		"button_11",
		"button_12",
		"button_13",
		"button_14",
	}
	devices [4]device
)

func main() {
	mode := &desktop.Mode{
		Width:      (width + padding*2) * pixelSize * 2,
		Height:     (height + padding*2) * pixelSize * 2,
		Fullscreen: false,
	}
	window := desktop.OpenWindow(mode)

	window.Run(&game.State{
		OnInit:              onInit,
		OnCleanup:           cleanup,
		OnRender:            onRender,
		OnKeyDown:           onKeyDown,
		OnControllerAdded:   onControllerAdded,
		OnControllerRemoved: onControllerRemoved,
	})
}

func onInit() {
	images = map[string]*gfx.Image{}
	for _, name := range imageNames {
		fn := fmt.Sprintf("./assets/controller/%s.png", name)
		images[name] = gfx.NewImage(fn)
	}

	gfx.SetClearColor(0.8, 0.8, 0.9)
	gfx.SetPixelSize(pixelSize)

	devices[0] = device{x: padding, y: padding}
	devices[1] = device{x: padding*2 + float64(width), y: padding}
	devices[2] = device{x: padding, y: padding*2 + float64(height)}
	devices[3] = device{x: padding*2 + float64(width), y: padding*2 + float64(height)}
}

func cleanup() {
	for _, i := range images {
		i.Delete()
	}
}

func onRender() {
	gfx.Clear()

	for _, d := range devices {
		p := gfx.NewParams()
		p.X = d.x
		p.Y = d.y
		if !d.connected {
			p.A = 0.5
		}

		gfx.Render(images["background"], p)

		p.X = d.x + d.leftStickX*stickMax
		p.Y = d.y + d.leftStickY*stickMax
		gfx.Render(images["left_stick"], p)
		if d.buttonsPressed[7] {
			gfx.Render(images["button_7"], p)
		}

		p.X = d.x + d.rightStickX*stickMax
		p.Y = d.y + d.rightStickY*stickMax
		gfx.Render(images["right_stick"], p)
		if d.buttonsPressed[8] {
			gfx.Render(images["button_8"], p)
		}

		p.X = d.x
		p.Y = d.y
		for i := 0; i < 15; i++ {
			if i == 7 || i == 8 {
				continue
			}
			if d.buttonsPressed[i] {
				name := fmt.Sprintf("button_%d", i)
				gfx.Render(images[name], p)
			}
		}

		if d.leftTrigger > 0 {
			p.A = d.leftTrigger
			gfx.Render(images["left_trigger"], p)
		}
		if d.rightTrigger > 0 {
			p.A = d.rightTrigger
			gfx.Render(images["right_trigger"], p)
		}
	}
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}

func onControllerAdded(c *controller.Controller) {
	if c.ID >= len(devices) {
		return
	}

	fmt.Printf("controller #%d connected\n", c.ID)

	d := &devices[c.ID]
	d.connected = true
	l := &controller.Listener{
		OnAxisMoved: func(a controller.Axis, value float64) {
			switch a {
			case 0:
				d.leftStickX = value
			case 1:
				d.leftStickY = value
			case 2:
				d.rightStickX = value
			case 3:
				d.rightStickY = value
			case 4:
				d.leftTrigger = value
			case 5:
				d.rightTrigger = value
			}
			fmt.Printf("Axis #%d of controller #%d has been moved by %f\n", a, c.ID, value)
		},
		OnButtonDown: func(b controller.Button) {
			fmt.Printf("Button #%d of controller #%d has been pressed\n", b, c.ID)
			d.buttonsPressed[b] = true
		},
		OnButtonUp: func(b controller.Button) {
			fmt.Printf("Button #%d of controller #%d has been released\n", b, c.ID)
			d.buttonsPressed[b] = false
		},
	}
	c.SetListener(l)
}

func onControllerRemoved(c *controller.Controller) {
	if c.ID >= len(devices) {
		return
	}

	fmt.Printf("Controller #%d has been removed\n", c.ID)
	devices[c.ID].connected = false
}
