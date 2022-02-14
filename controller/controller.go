package controller

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

// All contains all available controllers.
var All = map[sdl.JoystickID]*Controller{}

func init() {
	err := sdl.Init(sdl.INIT_GAMECONTROLLER | sdl.INIT_HAPTIC)
	if err != nil {
		log.Fatal(err)
	}
	sdl.GameControllerEventState(sdl.ENABLE)
}

// Count returns the number of Joysticks.
func Count() int {
	return sdl.NumJoysticks()
}

// Open opens the joystick for the given id and returns it.
func Open(id sdl.JoystickID) *Controller {
	sdlCtrl := sdl.GameControllerOpen(int(id))
	haptic, _ := sdl.HapticOpen(int(id))
	if haptic != nil {
		err := haptic.RumbleInit()
		if err != nil {
			haptic = nil
		}
	}
	if All[id] == nil {
		All[id] = &Controller{}
	}
	ctrl := All[id]
	ctrl.ID = int(id)
	ctrl.ctrl = sdlCtrl
	ctrl.haptic = haptic
	ctrl.Connected = true
	ctrl.Name = sdlCtrl.Name()
	return ctrl
}

// Close closes the joystick for the given id and returns the now disconnected controller.
func Close(id sdl.JoystickID) *Controller {
	ctrl := All[id]
	ctrl.haptic.Close()
	ctrl.haptic = nil
	ctrl.ctrl.Close()
	ctrl.Connected = false
	return ctrl
}

// DispatchButtonEvent triggers a button event. This func should not be called manually.
func DispatchButtonEvent(id sdl.JoystickID, b uint8, state uint8) {
	ctrl := All[id]
	if ctrl.listener != nil {
		if state == 1 {
			if ctrl.listener.OnButtonDown != nil {
				ctrl.listener.OnButtonDown(Button(b))
			}
		} else {
			if ctrl.listener.OnButtonUp != nil {
				ctrl.listener.OnButtonUp(Button(b))
			}
		}
	}
}

// DispatchAxisEvent triggers a axis event. This func should not be called manually.
func DispatchAxisEvent(id sdl.JoystickID, a uint8, value int16) {
	ctrl := All[id]
	if ctrl.listener != nil && ctrl.listener.OnAxisMoved != nil {
		v := float64(value) / 32767.0
		if v > 1.0 {
			v = 1.0
		}
		if v < -1.0 {
			v = -1.0
		}
		ctrl.listener.OnAxisMoved(Axis(a), v)
	}
}

// Controller is a Controller, Gamepad or Joystick.
type Controller struct {
	ID        int
	Name      string
	Connected bool
	ctrl      *sdl.GameController
	haptic    *sdl.Haptic
	listener  *Listener
}

// SetListener sets the Listener that will get notified when controller events occur.
func (c *Controller) SetListener(l *Listener) {
	c.listener = l
}

// ClearListener removes the Listener for the controller.
func (c *Controller) ClearListener() {
	c.listener = nil
}

// Rumble lets the controller vibrate for a given time and strength.
func (c *Controller) Rumble(strength float32, duration uint32) error {
	if c.haptic == nil {
		return nil
	}
	return c.haptic.RumblePlay(strength, duration)
}
