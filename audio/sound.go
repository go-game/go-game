package audio

import (
	"fmt"

	mix "github.com/veandco/go-sdl2/mix"
)

type Sound struct {
	c       *mix.Chunk
	channel int
}

// NewSound open the file for the given filename and returns a Sound.
func NewSound(fileName string) (*Sound, error) {
	chunk, err := mix.LoadWAV(fileName)
	if err != nil {
		return nil, err
	}

	return &Sound{c: chunk}, nil
}

// Play plays the sound.
func (s *Sound) Play() (err error) {
	s.channel, err = s.c.Play(-1, 0)
	fmt.Printf("channel: %d\n", s.channel)
	return
}

func (s *Sound) setVolume(v float) {
	s.c.Volume(int(v * mix.MAX_VOLUME))
}
