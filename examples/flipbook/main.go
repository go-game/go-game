//go:build example
// +build example

package main

import (
	"fmt"
	"time"

	"github.com/go-game/go-game/desktop"
	"github.com/go-game/go-game/game"
	"github.com/go-game/go-game/gfx"
	"github.com/go-game/go-game/keys"
)

var (
	flipBook *gfx.FlipBook
	params   *gfx.Params
	pages    [8]*gfx.Image

	tick      *gfx.Image
	tock      *gfx.Image
	tickTock  *gfx.Image
	pTickTock *gfx.Params

	callback   gfx.OnFlip
	showSounds bool
)

func main() {
	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	window, err := desktop.OpenWindow(mode)
	if err != nil {
		panic(err)
	}

	window.Run(&game.State{
		OnInit:    onInit,
		OnCleanup: onCleanup,
		OnRender:  render,
		OnUpdate:  update,
		OnKeyDown: onKeyDown,
	})
}

func onInit() {
	gfx.SetClearColor(0.2, 0.2, 0.2)
	gfx.SetPixelSize(4)

	var err error
	pages = [8]*gfx.Image{}
	for i := 0; i < len(pages); i++ {
		if pages[i], err = gfx.NewImage(fmt.Sprintf("assets/pointer/%d.png", i)); err != nil {
			panic(err)
		}
	}
	if tick, err = gfx.NewImage("assets/tick.png"); err != nil {
		panic(err)
	}
	if tock, err = gfx.NewImage("assets/tock.png"); err != nil {
		panic(err)
	}
	tickTock = tick
	pTickTock = gfx.NewParams()
	pTickTock.X = 130
	pTickTock.Y = 100

	d := time.Second / 2
	flipBook = gfx.NewFlipBook(
		true,
		&gfx.Page{Duration: d, Renderer: pages[0]},
		&gfx.Page{Duration: d, Renderer: pages[1]},
		&gfx.Page{Duration: d, Renderer: pages[2]},
		&gfx.Page{Duration: d, Renderer: pages[3]},
		&gfx.Page{Duration: d, Renderer: pages[4]},
		&gfx.Page{Duration: d, Renderer: pages[5]},
		&gfx.Page{Duration: d, Renderer: pages[6]},
		&gfx.Page{Duration: d, Renderer: pages[7]},
	)

	callback = func(page int) {
		tickTock = tock
		if page%2 == 0 {
			tickTock = tick
		}
	}

	params = gfx.NewParams()
	params.X = 100
	params.Y = 100
}

func onCleanup() {
	for i := 0; i < len(pages); i++ {
		pages[i].Delete()
	}
}

func render() {
	gfx.Clear()
	gfx.Render(flipBook, params)
	if showSounds {
		gfx.Render(tickTock, pTickTock)
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
