package audio

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_mixer"
)

func init() {
	if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
		panic(err)
	}

	if err := mix.OpenAudio(mix.DEFAULT_FREQUENCY, mix.DEFAULT_FORMAT, 2, 1024); err != nil {
		log.Println(err)
		return
	}
}

// Cleanup cleans up all audio data
func Cleanup() {
	defer mix.Quit()
	defer mix.CloseAudio()
}
