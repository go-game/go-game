package gfx

var mainCamera *Camera

// SetPixelSize scales all graphics by the given factor.
func SetPixelSize(size int) {
	mainCamera.SetPixelSize(size)
}

// SetArea defines the area on the screen to be drawn to.
func SetArea(w, h int) {
	activeCamera = NewCamera(w, h)
	mainCamera = activeCamera
}
