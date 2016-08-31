package audio

import (
	"github.com/vova616/go-openal/openal"
)

// Source emits sounds on the hardware.
type Source struct {
	openalSource openal.Source
}

// Play starts the playback
func (s *Source) Play() {
	s.openalSource.Play()
}

// Pause pauses the playback
func (s *Source) Pause() {
	s.openalSource.Pause()
}

// Stop stops the playback
func (s *Source) Stop() {
	s.openalSource.Stop()
}

// State in which the sound currently is
func (s *Source) State() int32 {
	return s.openalSource.State()
}
