package keys

import (
  "github.com/go-gl/glfw/v3.1/glfw"

  "git.mbuechmann.com/go-game/desktop"
)

var keyMap = map[string]glfw.Key{
  " ": glfw.KeySpace,
  "space": glfw.KeySpace,
  "up": glfw.KeyUp,
  "right": glfw.KeyRight,
  "down": glfw.KeyDown,
  "left": glfw.KeyLeft,
  "escape": glfw.KeyEscape,
  "esc": glfw.KeyEscape,
}

// Up return whether the given key is not pressed
func Up(k string) bool {
  return desktop.CurrentWindow.GlfwWindow.GetKey(keyMap[k]) != glfw.Press
}

// Down return whether the given key is pressed
func Down(k string) bool {
  return desktop.CurrentWindow.GlfwWindow.GetKey(keyMap[k]) == glfw.Press
}
