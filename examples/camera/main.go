// +build example

package main

import (
	"time"

	"github.com/go-game/go-game/desktop"
	"github.com/go-game/go-game/game"
	"github.com/go-game/go-game/gfx"
	"github.com/go-game/go-game/keys"
)

const width = 640
const height = 640

const tileSize = 64

var tile1 *gfx.Image
var tile2 *gfx.Image

var ball *gfx.Image
var ballParams *gfx.Params

const v = 100.0 / float64(time.Second)

var vX = 0.0
var vY = 0.0

var camera1 *gfx.Camera
var camera2 *gfx.Camera

var vcX = 0.0
var vcY = 0.0

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
	window := desktop.OpenWindow(mode)

	window.Run(state)
}

func onInit() {
	tile1 = gfx.NewImage("./assets/tile1.png")
	tile2 = gfx.NewImage("./assets/tile2.png")
	ball = gfx.NewImage("./assets/ball.png")

	camera1 = gfx.NewCamera(width, height, 0, 0, 2)
	camera2 = gfx.NewCamera(width, height, width, 0, 1)

	ballParams = gfx.NewParams()

}

func onRender() {
	gfx.Clear()
	renderCamera(camera1, 1)
	renderCamera(camera2, 2)
}

func onUpdate(delta time.Duration) {
	ballParams.X += vX * float64(delta)
	ballParams.Y += vY * float64(delta)

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
	p1 := gfx.NewParams()
	p2 := gfx.NewParams()
	for x := 0; x < 6*f; x++ {
		p1.X = float64(x * tileSize)
		p2.X = float64(x*tileSize - tileSize/2)
		for y := -1; y < 10*f; y++ {
			p1.Y = float64(y * tileSize / 2)
			p2.Y = float64(y*tileSize/2 - tileSize/4)
			c.Render(tile1, p1)
			c.Render(tile2, p2)
		}
	}

	c.Render(ball, ballParams)
}
