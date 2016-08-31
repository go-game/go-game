package main

import (
	"git.mbuechmann.com/go-game/audio"
	"github.com/vova616/go-openal/openal"
)

func main() {
	data, err := audio.LoadData("assets/welcome.wav")
	if err != nil {
		panic(err)
	}

	source := data.NewSource()

	source.Play()
	for source.State() == openal.Playing {
	}
	return
}
