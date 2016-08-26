package mouse

import "github.com/go-gl/glfw/v3.1/glfw"

var glfwWindow *glfw.Window

// SetGlfwWindow sets the window used for polling the keyboard.
func SetGlfwWindow(w *glfw.Window) {
	glfwWindow = w
}

// Hide makes the mouse cursor invisible
func Hide() {
	glfwWindow.SetInputMode(glfw.CursorMode, glfw.CursorHidden)
}

// Show makes the mouse cursor visible
func Show() {
	glfwWindow.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
}

func Position() (x, y float32) {
	 a, b := glfwWindow.GetCursorPos()
	//  _, height := glfwWindow.GetSize()
	 return float32(a), float32(b)
}
