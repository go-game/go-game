// +build example

package main

import (
	"time"

	"git.mbuechmann.com/go-game/desktop"
	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/gfx"
	"git.mbuechmann.com/go-game/keys"
)

var image *gfx.Image
var tween *gfx.Tween
var elapsed time.Duration = 0

func main() {
	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	desktop.OpenWindow(mode)

	gfx.SetPixelSize(4)

	desktop.Run(&game.State{
		InitFunc:    initFunc,
		CleanupFunc: cleanup,
		UpdateFunc:  update,
		OnKeyDown:   keyDown,
		RenderFunc:  render,
	})
}

func initFunc() {
	image = gfx.NewImage("assets/heart.png")
	r1 := gfx.NewRenderOptions()
	r1.X = 10
	r1.Y = -20
	r2 := gfx.NewRenderOptions()
	r2.X = 10
	r2.Y = 220
	tween = gfx.NewTween(r1, r2, 2*time.Second, 0, true)
}

func cleanup() {
	image.Delete()
}

func keyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}

func render() {
	gfx.Clear()
	ro := tween.GetRenderOptions()
	image.Render(ro)
}

func update(delta time.Duration) {
	elapsed += delta
	tween.Update(elapsed)
}
