package keys

import "github.com/go-gl/glfw/v3.1/glfw"

var glfwWindow *glfw.Window

// SetGlfwWindow sets the window used for polling the keyboard.
func SetGlfwWindow(w *glfw.Window) {
	glfwWindow = w
}

// Key is a specific key.
type Key int

const (
	// Space is the space bar.
	Space = Key(glfw.KeySpace)
	// Up ist the up cursor.
	Up = Key(glfw.KeyUp)
	// Right ist the right cursor.
	Right = Key(glfw.KeyRight)
	// Down ist the down cursor.
	Down = Key(glfw.KeyDown)
	// Left ist the left cursor.
	Left = Key(glfw.KeyLeft)
	// Esc ist the escape key.
	Esc = Key(glfw.KeyEscape)
	// A key.
	A = Key(glfw.KeyA)
	// B key.
	B = Key(glfw.KeyB)
	// C key.
	C = Key(glfw.KeyC)
	// D key.
	D = Key(glfw.KeyD)
	// E key.
	E = Key(glfw.KeyE)
	// F key.
	F = Key(glfw.KeyF)
	// G key.
	G = Key(glfw.KeyG)
	// H key.
	H = Key(glfw.KeyH)
	// I key.
	I = Key(glfw.KeyI)
	// J key.
	J = Key(glfw.KeyJ)
	// K key.
	K = Key(glfw.KeyK)
	// L key.
	L = Key(glfw.KeyL)
	// M key.
	M = Key(glfw.KeyM)
	// N key.
	N = Key(glfw.KeyN)
	// O key.
	O = Key(glfw.KeyO)
	// P key.
	P = Key(glfw.KeyP)
	// Q key.
	Q = Key(glfw.KeyQ)
	// R key.
	R = Key(glfw.KeyR)
	// S key.
	S = Key(glfw.KeyS)
	// T key.
	T = Key(glfw.KeyT)
	// U key.
	U = Key(glfw.KeyU)
	// V key.
	V = Key(glfw.KeyV)
	// W key.
	W = Key(glfw.KeyW)
	// X key.
	X = Key(glfw.KeyX)
	// Y key.
	Y = Key(glfw.KeyY)
	// Z key.
	Z = Key(glfw.KeyZ)
)

// IsUp return whether the given key is not pressed.
func IsUp(k Key) bool {
	return glfwWindow.GetKey(glfw.Key(k)) != glfw.Press
}

// IsDown return whether the given key is pressed.
func IsDown(k Key) bool {
	return glfwWindow.GetKey(glfw.Key(k)) == glfw.Press
}
