package gfx

import (
	_ "image/png" // needed to load png files

	"image"
	"image/draw"
	"os"

	"github.com/go-gl/gl/v2.1/gl"
)

// Image represents an image and can be rendered on the screen.
type Image struct {
	tex    *texture
	width  int
	height int
}

// Delete removes the image from memory.
func (i *Image) Delete() {
	gl.DeleteTextures(1, &i.tex.id)
}

// Width returns the width of the image in pixels.
func (i *Image) Width() int {
	return i.width
}

// Height returns the height of the image in pixels.
func (i *Image) Height() int {
	return i.height
}

func (i *Image) render(o *RenderOptions) {
	i.tex.render(o)
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
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	width := rgba.Rect.Size().X
	height := rgba.Rect.Size().Y
	tex := newTexture(width, height, rgba.Pix)

	return &Image{tex: tex, width: width, height: height}
}
