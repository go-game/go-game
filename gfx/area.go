package gfx

// SetPixelSize scales all graphics by the given factor.
func SetPixelSize(size int) {
	activeCamera.SetPixelSize(size)
}

// SetArea defines the area on the screen to be drawn to.
func SetArea(w, h int) {
	activeCamera = NewCamera(w, h)
}
