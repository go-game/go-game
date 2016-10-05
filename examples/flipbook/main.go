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

var (
	flipBook      *animation.Flipbook
	renderOptions *gfx.RenderOptions
	pages         [8]*gfx.Image

	tick       *gfx.Image
	tock       *gfx.Image
	tickTock   *gfx.Image
	roTickTock *gfx.RenderOptions

	callback   animation.OnFlip
	showSounds bool
)

func main() {
	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	desktop.OpenWindow(mode)
	gfx.SetClearColor(0.2, 0.2, 0.2, 1.0)
	gfx.SetPixelSize(4)

	desktop.Run(&game.State{
		OnInit:    onInit,
		OnCleanup: onCleanup,
		OnRender:  render,
		OnUpdate:  update,
		OnKeyDown: onKeyDown,
	})
}

func onInit() {
	pages = [8]*gfx.Image{}
	for i := 0; i < len(pages); i++ {
		pages[i] = gfx.NewImage(fmt.Sprintf("assets/pointer/%d.png", i))
	}
	tick = gfx.NewImage("assets/tick.png")
	tock = gfx.NewImage("assets/tock.png")
	tickTock = tick
	roTickTock = gfx.NewRenderOptions()
	roTickTock.X = 130
	roTickTock.Y = 100

	d := time.Second / 2
	flipBook = animation.NewFlipbook(
		true,
		&animation.Page{Duration: d, Renderer: pages[0]},
		&animation.Page{Duration: d, Renderer: pages[1]},
		&animation.Page{Duration: d, Renderer: pages[2]},
		&animation.Page{Duration: d, Renderer: pages[3]},
		&animation.Page{Duration: d, Renderer: pages[4]},
		&animation.Page{Duration: d, Renderer: pages[5]},
		&animation.Page{Duration: d, Renderer: pages[6]},
		&animation.Page{Duration: d, Renderer: pages[7]},
	)

	callback = func(page int) {
		tickTock = tock
		if page%2 == 0 {
			tickTock = tick
		}
	}

	renderOptions = gfx.NewRenderOptions()
	renderOptions.X = 100
	renderOptions.Y = 100
}

func onCleanup() {
	for i := 0; i < len(pages); i++ {
		pages[i].Delete()
	}
}

func render() {
	gfx.Clear()
	r := flipBook.CurrentRenderer()
	if r != nil {
		gfx.Render(r, renderOptions)
	}
	if showSounds {
		gfx.Render(tickTock, roTickTock)
	}
}

func update(delta time.Duration) {
	flipBook.Update(delta)
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}

	if k == keys.Space {
		if showSounds {
			flipBook.ClearPageListeners()
		} else {
			flipBook.AddPageListener(callback)
		}
		showSounds = !showSounds
	}
}
