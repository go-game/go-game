package audio

import (
	mix "github.com/veandco/go-sdl2/sdl_mixer"
)

type Sound struct {
	c *mix.Chunk
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
func (s *Sound) Play() {
	s.c.Play(0, 0)
}
