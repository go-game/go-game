package controller

import "github.com/veandco/go-sdl2/sdl"

// Button is a controller's button.
type Button int

// All possible buttons.
const (
	ButtonA             = Button(sdl.CONTROLLER_BUTTON_A)
	ButtonB             = Button(sdl.CONTROLLER_BUTTON_B)
	ButtonX             = Button(sdl.CONTROLLER_BUTTON_X)
	ButtonY             = Button(sdl.CONTROLLER_BUTTON_Y)
	ButtonBack          = Button(sdl.CONTROLLER_BUTTON_BACK)
	ButtonGuide         = Button(sdl.CONTROLLER_BUTTON_GUIDE)
	ButtonStart         = Button(sdl.CONTROLLER_BUTTON_START)
	ButtonLeftStick     = Button(sdl.CONTROLLER_BUTTON_LEFTSTICK)
	ButtonRightStick    = Button(sdl.CONTROLLER_BUTTON_RIGHTSTICK)
	ButtonLeftShoulder  = Button(sdl.CONTROLLER_BUTTON_LEFTSHOULDER)
	ButtonRightShoulder = Button(sdl.CONTROLLER_BUTTON_RIGHTSHOULDER)
	ButtonDPadUp        = Button(sdl.CONTROLLER_BUTTON_DPAD_UP)
	ButtonDPadDown      = Button(sdl.CONTROLLER_BUTTON_DPAD_DOWN)
	ButtonDPadLeft      = Button(sdl.CONTROLLER_BUTTON_DPAD_LEFT)
	ButtonDPadRight     = Button(sdl.CONTROLLER_BUTTON_DPAD_RIGHT)
)
