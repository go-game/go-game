package gfx

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func init() {
	err := ttf.Init()
	if err != nil {
		panic(err)
	}
}

var fontColor = sdl.Color{R: 255, G: 255, B: 255, A: 255}

// NewFont creates a new font from a given file for a given size.
func NewFont(filename string, size int) (*Font, error) {
	sdlFont, err := ttf.OpenFont(filename, size)
	return &Font{sdlFont: sdlFont}, err
}

// Font represents a ttf font which can convert a string into an image.
type Font struct {
	Bold          bool
	Italic        bool
	Underline     bool
	Strikethrough bool
	AntiAliased   bool
	sdlFont       *ttf.Font
	sdlColor      sdl.Color
}

// Size returns the width and height for the given text.
func (f *Font) Size(text string) (x, y int, err error) {
	return f.sdlFont.SizeUTF8(text)
}

// Render converts the given Text into an Image.
func (f *Font) Render(text string) (*Image, error) {
	style := 0
	if f.Bold {
		style++
	}
	if f.Italic {
		style += 2
	}
	if f.Underline {
		style += 4
	}
	if f.Strikethrough {
		style += 8
	}
	f.sdlFont.SetStyle(style)

	var sdlSurface *sdl.Surface
	var err error

	if f.AntiAliased {
		sdlSurface, err = f.sdlFont.RenderUTF8Blended(text, fontColor)
	} else {
		sdlSurface, err = f.sdlFont.RenderUTF8Solid(text, fontColor)
	}

	if err != nil {
		return nil, err
	}

	t := newTexture(int(sdlSurface.W), int(sdlSurface.H), sdlSurface.Pixels())
	return &Image{tex: t, width: int(sdlSurface.W), height: int(sdlSurface.H)}, nil
}

// Delete deletes the font from memory.
func (f *Font) Delete() {
	f.sdlFont.Close()
}
