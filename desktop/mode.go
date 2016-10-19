package desktop

import 	"github.com/go-gl/glfw/v3.1/glfw"

// Mode represents the resolution of a window and whether it is fullscreen.
type Mode struct {
	Width      int
	Height     int
	Fullscreen bool
}

// FullscreenModes returns an array of all available fullscreen modes.
func FullscreenModes() []*Mode {
	monitor := glfw.GetPrimaryMonitor()
	videoModes := monitor.GetVideoModes()

	var modes = make([]*Mode, len(videoModes))
	for i, m := range videoModes {
		modes[i] = &Mode{Width: m.Width, Height: m.Height, Fullscreen: true}
	}

	return modes
}

// CurrentMode returns the mode that is currently active.
func CurrentMode() *Mode {
	if window == nil {
		monitor := glfw.GetPrimaryMonitor()
		m := monitor.GetVideoMode()
		return &Mode{Width: m.Width, Height: m.Height}
	}
	return window.mode
}
