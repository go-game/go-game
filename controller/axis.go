package controller

import "github.com/veandco/go-sdl2/sdl"

// Axis is a controller's axis.
type Axis int

const (
	AXIS_LEFT_X        = Axis(sdl.CONTROLLER_AXIS_LEFTX)
	AXIS_LEFT_Y        = Axis(sdl.CONTROLLER_AXIS_LEFTY)
	AXIS_RIGHT_X       = Axis(sdl.CONTROLLER_AXIS_RIGHTX)
	AXIS_RIGHT_Y       = Axis(sdl.CONTROLLER_AXIS_RIGHTY)
	AXIS_TRIGGER_LEFT  = Axis(sdl.CONTROLLER_AXIS_TRIGGERLEFT)
	AXIS_TRIGGER_RIGHT = Axis(sdl.CONTROLLER_AXIS_TRIGGERRIGHT)
)
