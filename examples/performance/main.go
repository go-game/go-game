// +build example

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-game/go-game/desktop"
	"github.com/go-game/go-game/game"
	"github.com/go-game/go-game/gfx"
	"github.com/go-game/go-game/keys"
)

type heart struct {
	Params *gfx.Params
	Vx, Vy float64
}

const (
	WIDTH  = 1280
	HEIGHT = 800
	G      = 600 / float64(time.Second*time.Second)
)

var hearts []*heart
var image *gfx.Image
var adding bool

var frames int
var elapsed time.Duration

func main() {
	mode := &desktop.Mode{Width: WIDTH, Height: HEIGHT, Fullscreen: false}
	window := desktop.OpenWindow(mode)

	window.Run(&game.State{
		OnInit:    onInit,
		OnCleanup: onCleanup,
		OnUpdate:  onUpdate,
		OnRender:  onRender,
		OnKeyDown: onKeyDown,
		OnKeyUp:   onKeyUp,
	})
}

func onInit() {
	image = gfx.NewImage("assets/heart.png")
}

func onCleanup() {
	image.Delete()
}

func onRender() {
	gfx.Clear()
	for _, h := range hearts {
		gfx.Render(image, h.Params)
	}
	frames++
	fmt.Printf("fps: %f\n", float64(frames)/float64(elapsed/time.Second))
	fmt.Printf("#hearts: %d\n", len(hearts))
}

func onUpdate(delta time.Duration) {
	elapsed += delta
	if adding {
		for i := 0; i < rand.Intn(30); i++ {
			vx := (rand.Float64()*400 + 100) / float64(time.Second)
			vy := (rand.Float64()*100 + 100) / float64(time.Second)
			h := heart{Params: gfx.NewParams(), Vx: vx, Vy: vy}
			hearts = append(hearts, &h)
		}
	}

	for _, h := range hearts {
		h.Params.X += h.Vx * float64(delta)
		if h.Vx > 0 {
			d := WIDTH - h.Params.X - float64(image.Width())
			if d < 0 {
				h.Params.X -= d
				h.Vx *= -1
			}
		} else {
			if h.Params.X < 0 {
				h.Params.X *= -1
				h.Vx *= -1
			}
		}

		h.Vy += G * float64(delta)
		h.Params.Y += h.Vy * float64(delta)
		if h.Vy > 0 {
			d := HEIGHT - h.Params.Y - float64(image.Height())
			if d < 0 {
				h.Params.Y -= d
				h.Vy *= -1
			}
		}
		if h.Params.Y > HEIGHT*2 {
			h.Params.Y = -float64(image.Height())
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
