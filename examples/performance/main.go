//go:build example
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
	x, y   float64
	vx, vy float64
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
	window, err := desktop.OpenWindow(mode)
	if err != nil {
		panic(err)
	}

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
	var err error
	if image, err = gfx.NewImage("assets/heart.png"); err != nil {
		panic(err)
	}
}

func onCleanup() {
	image.Delete()
}

func onRender() {
	gfx.Clear()
	for _, h := range hearts {
		gfx.RenderXY(image, h.x, h.y)
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
			h := heart{vx: vx, vy: vy}
			hearts = append(hearts, &h)
		}
	}

	for _, h := range hearts {
		h.x += h.vx * float64(delta)
		if h.vx > 0 {
			d := WIDTH - h.x - float64(image.Width())
			if d < 0 {
				h.x -= d
				h.vx *= -1
			}
		} else {
			if h.x < 0 {
				h.x *= -1
				h.vx *= -1
			}
		}

		h.vy += G * float64(delta)
		h.y += h.vy * float64(delta)
		if h.vy > 0 {
			d := HEIGHT - h.y - float64(image.Height())
			if d < 0 {
				h.y -= d
				h.vy *= -1
			}
		}
		if h.y > HEIGHT*2 {
			h.y = -float64(image.Height())
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
