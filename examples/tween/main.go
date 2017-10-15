// +build example

package main

import (
	"time"

	"github.com/mbuechmann/go-game/desktop"
	"github.com/mbuechmann/go-game/game"
	"github.com/mbuechmann/go-game/gfx"
	"github.com/mbuechmann/go-game/gfx/animation"
	"github.com/mbuechmann/go-game/keys"
)

const offset time.Duration = 0

var image *gfx.Image
var tween *animation.Tween
var running = true

func main() {
	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	window := desktop.OpenWindow(mode)

	gfx.SetPixelSize(4)

	window.Run(&game.State{
		OnInit:    onInit,
		OnCleanup: onCleanup,
		OnUpdate:  onUpdate,
		OnKeyDown: keyDown,
		OnRender:  onRender,
	})
}

func onInit() {
	image = gfx.NewImage("assets/heart.png")
	r1 := gfx.NewRenderOptions()
	r1.X = 50
	r1.Y = -16
	r1.Rot = gfx.Rotation{Angle: 0, X: 8, Y: 8}
	r1.Scale = gfx.Scale{Factor: 1, X: 8, Y: 8}
	r2 := gfx.NewRenderOptions()
	r2.X = 50
	r2.Y = 216
	r2.Rot = gfx.Rotation{Angle: 360, X: 8, Y: 8}
	r2.Scale = gfx.Scale{Factor: 5, X: 8, Y: 8}
	tween = animation.NewTween(r1, r2, 2*time.Second, offset, true)
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
	ro := tween.GetRenderOptions()
	gfx.Render(image, ro)
}

func onUpdate(delta time.Duration) {
	if running {
		tween.Update(delta)
	}
}
