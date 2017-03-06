package controller

type Listener struct {
	// OnButtonDown gets called when a button is pressed.
	OnButtonDown func(b Button)
	// OnButtonUp gets called when a button is released.
	OnButtonUp func(b Button)
	// OnAxisMoved gets called when a axis is moved.
	OnAxisMoved func(a Axis, value float64)
	// OnConnect gets called when the controller connects.
	OnConnect func()
	// OnDisconnectConnect gets called when the controller disconnects.
	OnDisconnect func()
}
