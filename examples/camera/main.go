//go:build example
// +build example

package main

import (
	"time"

	"github.com/go-game/go-game/desktop"
	"github.com/go-game/go-game/game"
	"github.com/go-game/go-game/gfx"
	"github.com/go-game/go-game/keys"
)

const (
	width  = 640
	height = 640

	tileSize = 64

	v = 100.0 / float64(time.Second)
)

var (
	tile1 *gfx.Image
	tile2 *gfx.Image

	ball         *gfx.Image
	ballX, ballY float64
	vX           = 0.0
	vY           = 0.0

	camera1 *gfx.Camera
	camera2 *gfx.Camera
	vcX     = 0.0
	vcY     = 0.0
)

func main() {
	state := &game.State{
		OnInit:    onInit,
		OnRender:  onRender,
		OnUpdate:  onUpdate,
		OnKeyDown: onKeyDown,
		OnKeyUp:   onKeyUp,
		OnCleanup: onCleanup,
	}

	mode := &desktop.Mode{Width: 2 * width, Height: height, Fullscreen: false}
	window, err := desktop.OpenWindow(mode)
	if err != nil {
		panic(err)
	}

	window.Run(state)
}

func onInit() {
	var err error

	if tile1, err = gfx.NewImage("./assets/tile1.png"); err != nil {
		panic(err)
	}
	if tile2, err = gfx.NewImage("./assets/tile2.png"); err != nil {
		panic(err)
	}
	if ball, err = gfx.NewImage("./assets/ball.png"); err != nil {
		panic(err)
	}

	camera1 = gfx.NewCamera(width, height, 0, 0, 2)
	camera2 = gfx.NewCamera(width, height, width, 0, 1)
}

func onRender() {
	gfx.Clear()
	renderCamera(camera1, 1)
	renderCamera(camera2, 2)
}

func onUpdate(delta time.Duration) {
	ballX += vX * float64(delta)
	ballY += vY * float64(delta)

	camera1.Move(vcX*float64(delta), vcY*float64(delta))
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}

	switch k {
	case keys.Right:
		vX += v
	case keys.Left:
		vX -= v
	case keys.Down:
		vY += v
	case keys.Up:
		vY -= v

	case keys.W:
		vcY -= v
	case keys.A:
		vcX += v
	case keys.S:
		vcY += v
	case keys.D:
		vcX -= v
	}
}

func onKeyUp(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}

	switch k {
	case keys.Right:
		vX -= v
	case keys.Left:
		vX += v
	case keys.Down:
		vY -= v
	case keys.Up:
		vY += v

	case keys.W:
		vcY += v
	case keys.A:
		vcX -= v
	case keys.S:
		vcY -= v
	case keys.D:
		vcX += v
	}
}

func onCleanup() {
	tile1.Delete()
	tile2.Delete()
	ball.Delete()
}

func renderCamera(c *gfx.Camera, f int) {
	for x := 0; x < 6*f; x++ {
		for y := -1; y < 10*f; y++ {
			c.RenderXY(tile1, float64(x*tileSize), float64(y*tileSize/2))
			c.RenderXY(tile2, float64(x*tileSize-tileSize/2), float64(y*tileSize/2-tileSize/4))
		}
	}

	c.RenderXY(ball, ballX, ballY)
}
