package mouse

import "github.com/veandco/go-sdl2/sdl"

// Button is a specific button of the mouse.
type Button int

const (
	// Button1 is the first (left) mouse button.
	Button1 = Button(sdl.BUTTON_LEFT)
	// Button2 is the second (right) mouse button.
	Button2 = Button(sdl.BUTTON_RIGHT)
	// Button3 is the third (middle) mouse button.
	Button3 = Button(sdl.BUTTON_MIDDLE)
	// Button4 is the fourth mouse button.
	Button4 = Button(sdl.BUTTON_X1)
	// Button5 is the fifth mouse button.
	Button5 = Button(sdl.BUTTON_X2)
)

// Hide makes the mouse cursor invisible.
func Hide() {
	sdl.ShowCursor(sdl.DISABLE)
}

// Show makes the mouse cursor visible
func Show() {
	sdl.ShowCursor(sdl.ENABLE)
}
