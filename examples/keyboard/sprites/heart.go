package sprites

import (
	"git.mbuechmann.com/go-game/gfx"
)

// Heart is a tiny 16x16 heart
type Heart struct {
	posX    float32
	posY    float32
	speed   float32
	vX      float32
	vY      float32
	texture *gfx.Texture
}

// Render renders the heart at its location
func (h *Heart) Render() {
	h.texture.Render(h.posX, h.posY)
}

// Delete cleans up by deleting the texture
func (h *Heart) Delete() {
	h.texture.Delete()
}

// Update upates the hearts position
func (h *Heart) Update(delta float64) {
	var seconds = float32(delta / 1000000000)
	h.posX += h.vX * seconds
	h.posY += h.vY * seconds
}

// SetDirection sets the new direction of movement
func (h *Heart) SetDirection(x, y float32) {
	h.vX = x * h.speed
	h.vY = y * h.speed
}

// NewHeart returns a pointer to a new Heart
func NewHeart() *Heart {
	t := gfx.NewTexture("assets/heart.png")
	return &Heart{posX: 100, posY: 100, speed: 100, texture: t}
}
