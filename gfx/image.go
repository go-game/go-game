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
	id     uint32
	width  int
	height int
}

// Delete removes the image from memory.
func (t *Image) Delete() {
	gl.DeleteTextures(1, &t.id)
}

// Render renders the image on the screen at x, y.
func (t *Image) Render(x, y float64) {
	gl.BindTexture(gl.TEXTURE_2D, t.id)
	gl.Begin(gl.QUADS)
	gl.TexCoord2f(0, 0)
	gl.Vertex3d(x, -y, 0)
	gl.TexCoord2f(0, 1)
	gl.Vertex3d(x, -y-float64(t.height), 0)
	gl.TexCoord2f(1, 1)
	gl.Vertex3d(x+float64(t.width), -y-float64(t.height), 0)
	gl.TexCoord2f(1, 0)
	gl.Vertex3d(x+float64(t.width), -y, 0)
	gl.End()
}

// NewImage creates a new Image from the given file name. File must be a png.
func NewImage(file string) (t *Image) {
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

	var id uint32
	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &id)
	gl.BindTexture(gl.TEXTURE_2D, id)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	t = &Image{id: id, width: img.Bounds().Max.X, height: img.Bounds().Max.Y}

	return
}
