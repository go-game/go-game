// +build example

package main

import (
	"math/rand"
	"time"

	"git.mbuechmann.com/go-game/desktop"
	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/gfx"
	"git.mbuechmann.com/go-game/keys"
)

type heart struct {
	RenderOptions *gfx.RenderOptions
	Vx, Vy        float64
}

const (
	WIDTH  = 1280
	HEIGHT = 800
	G      = 600 / float64(time.Second*time.Second)
)

var hearts []*heart
var image *gfx.Image
var adding bool

func main() {
	mode := &desktop.Mode{Width: WIDTH, Height: HEIGHT, Fullscreen: false}
	desktop.OpenWindow(mode)

	desktop.Run(&game.State{
		InitFunc:    initGame,
		CleanupFunc: cleanup,
		UpdateFunc:  update,
		RenderFunc:  render,
		OnKeyDown:   onKeyDown,
		OnKeyUp:     onKeyUp,
	})
}

func initGame() {
	image = gfx.NewImage("assets/heart.png")
}

func cleanup() {
	image.Delete()
}

func render() {
	gfx.Clear()
	for _, h := range hearts {
		gfx.Render(image, h.RenderOptions)
	}
}

func update(delta time.Duration) {
	if adding {
		for i := 0; i < rand.Intn(3); i++ {
			vx := (rand.Float64()*400 + 100) / float64(time.Second)
			vy := (rand.Float64()*100 + 100) / float64(time.Second)
			h := heart{RenderOptions: gfx.NewRenderOptions(), Vx: vx, Vy: vy}
			hearts = append(hearts, &h)
		}
	}

	for _, h := range hearts {
		h.RenderOptions.X += h.Vx * float64(delta)
		if h.Vx > 0 {
			d := WIDTH - h.RenderOptions.X - float64(image.Width())
			if d < 0 {
				h.RenderOptions.X -= d
				h.Vx *= -1
			}
		} else {
			if h.RenderOptions.X < 0 {
				h.RenderOptions.X *= -1
				h.Vx *= -1
			}
		}

		h.Vy += G * float64(delta)
		h.RenderOptions.Y += h.Vy * float64(delta)
		if h.Vy > 0 {
			d := HEIGHT - h.RenderOptions.Y - float64(image.Height())
			if d < 0 {
				h.RenderOptions.Y -= d
				h.Vy *= -1
			}
		}
	}
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
	if k == keys.Space {
		adding = true
	}
}

func onKeyUp(k keys.Key) {
	if k == keys.Space {
		adding = false
	}
}
