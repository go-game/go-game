package controller

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// All contains all available controllers.
var All = []*Controller{}

func init() {
	sdl.Init(sdl.INIT_GAMECONTROLLER | sdl.INIT_HAPTIC)
	sdl.GameControllerEventState(sdl.ENABLE)
}

// Count returns the number of Joysticks.
func Count() int {
	return sdl.NumJoysticks()
}

// Open is later private.
func Open(id int) *Controller {
	sdlCtrl := sdl.GameControllerOpen(id)
	haptic := sdl.HapticOpen(id)
	err := haptic.RumbleInit()
	if err != 0 {
		fmt.Printf("Error on RumbleInit: %d\n", err)
	}
	ctrl := &Controller{
		ID:     id,
		Name:   sdlCtrl.Name(),
		ctrl:   sdlCtrl,
		haptic: haptic,
	}
	fmt.Printf("%+v\n", haptic)
	All = append(All, ctrl)
	return ctrl
}

// Controller is a Controller, Gamepad or Joystick.
type Controller struct {
	ID     int
	Name   string
	ctrl   *sdl.GameController
	haptic *sdl.Haptic
}

func (c *Controller) Rumble(strength float32, duration uint32) {
	c.haptic.RumblePlay(strength, duration)
}
