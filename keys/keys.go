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
  "a": glfw.KeyA,
  "b": glfw.KeyB,
  "c": glfw.KeyC,
  "d": glfw.KeyD,
  "e": glfw.KeyE,
  "f": glfw.KeyF,
  "g": glfw.KeyG,
  "h": glfw.KeyH,
  "i": glfw.KeyI,
  "j": glfw.KeyJ,
  "k": glfw.KeyK,
  "l": glfw.KeyL,
  "m": glfw.KeyM,
  "n": glfw.KeyN,
  "o": glfw.KeyO,
  "p": glfw.KeyP,
  "q": glfw.KeyQ,
  "r": glfw.KeyR,
  "s": glfw.KeyS,
  "t": glfw.KeyT,
  "u": glfw.KeyU,
  "v": glfw.KeyV,
  "w": glfw.KeyW,
  "x": glfw.KeyX,
  "y": glfw.KeyY,
  "z": glfw.KeyZ,
}

// Up return whether the given key is not pressed
func Up(k string) bool {
  return desktop.CurrentWindow.GlfwWindow.GetKey(keyMap[k]) != glfw.Press
}

// Down return whether the given key is pressed
func Down(k string) bool {
  return desktop.CurrentWindow.GlfwWindow.GetKey(keyMap[k]) == glfw.Press
}
