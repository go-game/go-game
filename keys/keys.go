package keys

import (
  "github.com/go-gl/glfw/v3.1/glfw"

  "git.mbuechmann.com/go-game/base"
)

var keyMap map[string]glfw.Key = map[string]glfw.Key{
  " ": glfw.KeySpace,
  "space": glfw.KeySpace,
  "up": glfw.KeyUp,
  "right": glfw.KeyRight,
  "down": glfw.KeyDown,
  "left": glfw.KeyLeft,
  "escape": glfw.KeyEscape,
  "esc": glfw.KeyEscape,
}

func Up(k string) bool {
  return base.Window.GetKey(keyMap[k]) != glfw.Press
}

func Down(k string) bool {
  return base.Window.GetKey(keyMap[k]) == glfw.Press
}
