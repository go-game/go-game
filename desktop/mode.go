package desktop

import "github.com/veandco/go-sdl2/sdl"

// Mode represents the resolution of a window and whether it is fullscreen.
type Mode struct {
	Width      int32
	Height     int32
	Fullscreen bool
}

// FullscreenModes returns an array of all available fullscreen modes.
func FullscreenModes() []*Mode {
	res := []*Mode{}

	alreadyIncluded := func(m *Mode) bool {
		for _, mode := range res {
			if mode.Width == m.Width && mode.Height == m.Height {
				return true
			}
		}
		return false
	}

	count, _ := sdl.GetNumDisplayModes(0)
	for i := 0; i < count; i++ {
		sdlMode := &sdl.DisplayMode{}
		sdl.GetDisplayMode(0, i, sdlMode)
		mode := &Mode{Width: sdlMode.W, Height: sdlMode.H, Fullscreen: true}
		if !alreadyIncluded(mode) {
			res = append(res, mode)
		}
	}
	return res
}
