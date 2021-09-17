package audio

import "github.com/veandco/go-sdl2/mix"

// Sound represents a loaded sound file that can be played back.
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
	return
}

// Delete deletes the sound from memory.
func (s *Sound) Delete() {
	mix.ExpireChannel(s.channel, 0)
	s.c.Free()
}

// SetVolume sets the volume between 0.0 (silent) and 1.0 (full volume).
// Any outside this range will be capped.
func (s *Sound) SetVolume(v float64) {
	s.c.Volume(int(v * mix.MAX_VOLUME))
}
