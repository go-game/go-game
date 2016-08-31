package keys

import "github.com/go-gl/glfw/v3.1/glfw"

var glfwWindow *glfw.Window

// SetGlfwWindow sets the window used for polling the keyboard.
func SetGlfwWindow(w *glfw.Window) {
	glfwWindow = w
}

// Key is a specific key.
type Key int

// IsUp return whether the given key is not pressed.
func IsUp(k Key) bool {
	return glfwWindow.GetKey(glfw.Key(k)) != glfw.Press
}

// IsDown return whether the given key is pressed.
func IsDown(k Key) bool {
	return glfwWindow.GetKey(glfw.Key(k)) == glfw.Press
}
