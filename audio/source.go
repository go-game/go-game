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

// SetPitch sets the pitch
func (s *Source) SetPitch(f float32) {
	s.openalSource.SetPitch(f)
}

// SetGain sets the gain
func (s *Source) SetGain(f float32) {
	s.openalSource.SetGain(f)
}

// SetPosition sets the position
func (s *Source) SetPosition(x, y, z float32) {
	s.openalSource.SetPosition(x, y, z)
}

// SetVelocity sets the velocity
func (s *Source) SetVelocity(x, y, z float32) {
	s.openalSource.SetVelocity(x, y, z)
}

// SetLooping sets if the source is looping
func (s *Source) SetLooping(b bool) {
	s.openalSource.SetLooping(b)
}
