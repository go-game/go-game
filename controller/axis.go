package controller

import "github.com/veandco/go-sdl2/sdl"

// Axis is a controller's axis.
type Axis int

// All a possible axes.
const (
	AxisLeftX        = Axis(sdl.CONTROLLER_AXIS_LEFTX)
	AxisLeftY        = Axis(sdl.CONTROLLER_AXIS_LEFTY)
	AxisRightX       = Axis(sdl.CONTROLLER_AXIS_RIGHTX)
	AxisRightY       = Axis(sdl.CONTROLLER_AXIS_RIGHTY)
	AxisTriggerLeft  = Axis(sdl.CONTROLLER_AXIS_TRIGGERLEFT)
	AxisTriggerRight = Axis(sdl.CONTROLLER_AXIS_TRIGGERRIGHT)
)
