package game

import (
	"time"

	"github.com/mbuechmann/go-game/controller"
	"github.com/mbuechmann/go-game/keys"
	"github.com/mbuechmann/go-game/mouse"
)

// State is used as a scaffold to generate a new game.
type State struct {
	// OnInit gets called once when the game state starts. Put your initialization stuff here, like loading assets.
	OnInit func()
	// OnRender gets called periodically when things have to be rendered onto the screen.
	OnRender func()
	// OnUpdate gets calls periodically when your your game objects have to be updated. elapsed indicates the duration elapsed
	OnUpdate func(delta time.Duration)
	// OnCleanup gets called once when the game state exits. Put all your cleanuo stuff here, like deleting assets.
	OnCleanup func()
	// OnMouseMove gets called whenever the user moves the mouse, where x and y are the coordinates on the screen.
	OnMouseMove func(x, y int32)
	// OnMouseButtonDown gets called when the user presses a mouse button, where b indicates which button is pressed
	// and s and y are the coordinates on the screen.
	OnMouseButtonDown func(b mouse.Button, x, y float32)
	// OnMouseButtonDown gets called when the user releases a mouse button, where b indicates which button is released
	// and s and y are the coordinates on the screen.
	OnMouseButtonUp func(b mouse.Button, x, y float32)
	// OnMouseWheel gets called when the user scrolls with the mouse wheel or the touchpad, where x and y are the amounts
	// by which the user scrolled horizontally and vertically.
	OnMouseWheel func(x, y int32)
	// OnKeyDown gets called when the user pressed a key on the keyboard, where k indicates which key is pressed.
	OnKeyDown func(k keys.Key)
	// OnKeyUp gets called when the user releases a key on the keyboard, where k indicates which key is released.
	OnKeyUp func(k keys.Key)
	// OnControllerAdded gets called when a controller is added.
	OnControllerAdded func(c *controller.Controller)
	// OnControllerRemoved gets called when a controller is removed.
	OnControllerRemoved func(c *controller.Controller)
}
