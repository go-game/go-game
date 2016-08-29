package mouse

import "github.com/go-gl/glfw/v3.1/glfw"

// Button is a specific button of the mouse.
type Button int

const (
	// Button1 is the first mouse button.
	Button1 = Button(glfw.MouseButton1)
	// Button2 is the second mouse button.
	Button2 = Button(glfw.MouseButton2)
	// Button3 is the third mouse button.
	Button3 = Button(glfw.MouseButton3)
	// Button4 is the fourth mouse button.
	Button4 = Button(glfw.MouseButton4)
	// Button5 is the fifth mouse button.
	Button5 = Button(glfw.MouseButton5)
	// Button6 is the sixth mouse button.
	Button6 = Button(glfw.MouseButton6)
	// Button7 is the seventh mouse button.
	Button7 = Button(glfw.MouseButton7)
	// Button8 is the eighth mouse button.
	Button8 = Button(glfw.MouseButton8)
)

var glfwWindow *glfw.Window

// SetGlfwWindow sets the window used for polling the keyboard.
func SetGlfwWindow(w *glfw.Window) {
	glfwWindow = w
}

// Hide makes the mouse cursor invisible.
func Hide() {
	glfwWindow.SetInputMode(glfw.CursorMode, glfw.CursorHidden)
}

// Show makes the mouse cursor visible
func Show() {
	glfwWindow.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
}

// Position returns the x and y coordinate of the mouse cursor.
func Position() (x, y float32) {
	a, b := glfwWindow.GetCursorPos()
	return float32(a), float32(b)
}

// SetPosition sets the x and y coordinate of the mouse cursor.
func SetPosition(x, y float32) {
	glfwWindow.SetCursorPos(float64(x), float64(y))
}
