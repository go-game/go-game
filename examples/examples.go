package examples

import "github.com/go-game/go-game/gfx"

// TextImage loads the font for the examples and generates an image with the
// given text.
func TextImage(text string) *gfx.Image {
	var err error
	font, err := gfx.NewFont("assets/OpenSans-Regular.ttf", 16)
	if err != nil {
		panic(err)
	}
	font.Antialiased = true

	textImg, err := font.Render(text)
	if err != nil {
		panic(err)
	}

	return textImg
}

// RenderImage renders an image at the given coordinates.
func RenderImage(img *gfx.Image, x, y float64) {
	p := gfx.NewParams()
	p.X = x
	p.Y = y
	gfx.Render(img, p)
}
