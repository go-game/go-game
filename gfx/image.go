package gfx

import (
	"image"
	"image/draw"
	_ "image/png" // needed to load png files
	"os"
)

// Image represents an image and can be rendered on the screen.
type Image struct {
	tex        *texture
	width      int
	height     int
	filterMode FilterMode
}

// Delete removes the image from memory.
func (i *Image) Delete() {
	i.tex.delete()
}

// Width returns the width of the image in pixels.
func (i *Image) Width() int {
	return i.width
}

// Height returns the height of the image in pixels.
func (i *Image) Height() int {
	return i.height
}

func (i *Image) render(p *Params) {
	i.tex.activate(i.filterMode)
	i.tex.render(p)
}

// NewImage creates a new Image from the given file name. File must be a png.
func NewImage(file string) *Image {
	imgFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		panic(err)
	}

	rgba := image.NewRGBA(img.Bounds())

	if rgba.Stride != rgba.Rect.Size().X*4 {
		panic("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)

	width := rgba.Rect.Size().X
	height := rgba.Rect.Size().Y
	tex := newTexture(width, height, rgba.Pix)

	return &Image{tex: tex, width: width, height: height, filterMode: defaultFilterMode}
}
