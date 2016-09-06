package main

import (
	"time"

	"git.mbuechmann.com/go-game/desktop"
	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/gfx"
	"git.mbuechmann.com/go-game/keys"
)

var image *gfx.Image
var vX, vY float64 = 0, 0
var posX, posY float64 = 100, 100

const speed float64 = 100

func main() {
	gameState := &game.State{
		InitFunc:    initGame,
		RenderFunc:  render,
		UpdateFunc:  logic,
		CleanupFunc: cleanupGame,
		OnKeyUp:     onKeyUp,
		OnKeyDown:   onKeyDown,
	}

	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	desktop.OpenWindow(mode)
	gfx.SetPixelSize(4)

	desktop.Run(gameState)
}

func onKeyUp(k keys.Key) {
	if k == keys.Up {
		vY += speed
	}
	if k == keys.Down {
		vY -= speed
	}
	if k == keys.Right {
		vX -= speed
	}
	if k == keys.Left {
		vX += speed
	}
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}

	if k == keys.Up {
		vY -= speed
	}
	if k == keys.Down {
		vY += speed
	}
	if k == keys.Right {
		vX += speed
	}
	if k == keys.Left {
		vX -= speed
	}
}

func logic(delta time.Duration) {
	var seconds = (float64(delta) / 1000000000)
	posX += vX * seconds
	posY += vY * seconds
}

func render() {
	gfx.Clear()
	image.Render(posX, posY)
}

func initGame() {
	image = gfx.NewImage("assets/heart.png")
}

func cleanupGame() {
	image.Delete()
}
