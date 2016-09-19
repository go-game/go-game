// +build example

package main

import (
	"fmt"
	"time"

	"git.mbuechmann.com/go-game/desktop"
	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/gfx"
	"git.mbuechmann.com/go-game/gfx/animation"
	"git.mbuechmann.com/go-game/keys"
)

var flipBook *animation.Flipbook
var renderOptions *gfx.RenderOptions
var images [8]*gfx.Image

func main() {
	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	desktop.OpenWindow(mode)
	gfx.SetClearColor(1.0, 1.0, 1.0, 1.0)
	gfx.SetPixelSize(4)

	desktop.Run(&game.State{
		InitFunc:    initGame,
		CleanupFunc: cleanup,
		RenderFunc:  render,
		UpdateFunc:  update,
		OnKeyDown:   onKeyDown,
	})
}

func initGame() {
	images = [8]*gfx.Image{}
	for i := 0; i < len(images); i++ {
		images[i] = gfx.NewImage(fmt.Sprintf("assets/pointer/%d.png", i))
	}

	d := time.Second / 4
	flipBook = animation.NewFlipbook(
		true,
		&animation.Page{Duration: d, Renderer: images[0]},
		&animation.Page{Duration: d, Renderer: images[1]},
		&animation.Page{Duration: d, Renderer: images[2]},
		&animation.Page{Duration: d, Renderer: images[3]},
		&animation.Page{Duration: d, Renderer: images[4]},
		&animation.Page{Duration: d, Renderer: images[5]},
		&animation.Page{Duration: d, Renderer: images[6]},
		&animation.Page{Duration: d, Renderer: images[7]},
	)
	renderOptions = gfx.NewRenderOptions()
	renderOptions.X = 100
	renderOptions.Y = 100
}

func cleanup() {
	for i := 0; i < len(images); i++ {
		images[i].Delete()
	}
}

func render() {
	gfx.Clear()
	r := flipBook.CurrentRenderer()
	if r != nil {
		gfx.Render(r, renderOptions)
	}
}

func update(delta time.Duration) {
	flipBook.Update(delta)
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}
