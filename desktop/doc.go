/*
Package desktop manages the creation of windows and runs game states in thise windows.
This will be the entry point for all games running on the desktop.

To start a game you have to create a window with a mode and pass it a game state:

	func main() {
		mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
		window := desktop.OpenWindow(mode)

		window.Run(&game.State{
			// define your callbacks here...
		})
	}

*/
package desktop
