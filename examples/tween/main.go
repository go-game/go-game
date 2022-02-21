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

const offset time.Duration = 0

var image *gfx.Image
var tween *gfx.Tween
var running = true

func main() {
	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	window, err := desktop.OpenWindow(mode)
	if err != nil {
		panic(err)
	}

	window.Run(&game.State{
		OnInit:    onInit,
		OnCleanup: onCleanup,
		OnUpdate:  onUpdate,
		OnKeyDown: keyDown,
		OnRender:  onRender,
	})
}

func onInit() {
	gfx.SetPixelSize(4)
	var err error
	if image, err = gfx.NewImage("assets/heart.png"); err != nil {
		panic(err)
	}
	r1 := gfx.NewParams()
	r1.X = 50
	r1.Y = -16
	r1.Rot = gfx.Rotation{Angle: 0, X: 8, Y: 8}
	r1.Scale = gfx.Scale{Factor: 1, X: 8, Y: 8}
	r2 := gfx.NewParams()
	r2.X = 50
	r2.Y = 216
	r2.Rot = gfx.Rotation{Angle: 360, X: 8, Y: 8}
	r2.Scale = gfx.Scale{Factor: 5, X: 8, Y: 8}
	tween = gfx.NewTween(r1, r2, 2*time.Second, offset, true)
}

func onCleanup() {
	image.Delete()
}

func keyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
	if k == keys.Space {
		running = !running
	}
}

func onRender() {
	gfx.Clear()
	tween.Render(image)
}

func onUpdate(delta time.Duration) {
	if running {
		tween.Update(delta)
	}
}
