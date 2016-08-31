package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/vova616/go-openal/openal"
)

type soundFormat struct {
	FormatTag     int16
	Channels      int16
	Samples       int32
	AvgBytes      int32
	BlockAlign    int16
	BitsPerSample int16
}

type soundFormat2 struct {
	soundFormat
	SizeOfExtension int16
}

type soundFormat3 struct {
	soundFormat2
	ValidBitsPerSample int16
	ChannelMask        int32
	SubFormat          [16]byte
}

// Source emits sounds on the hardware.
type Source struct {
	openalSource openal.Source
}

// Data is a sound sample loaded from a file.
type Data struct {
	Bytes  []byte
	Format *soundFormat
	buffer openal.Buffer
}

// NewSource returns a new Source with pitch, gain, position velocity and not looping.
func (s *Data) NewSource() *Source {
	var source = openal.NewSource()
	source.SetPitch(1)
	source.SetGain(1)
	source.SetPosition(0, 0, 0)
	source.SetVelocity(0, 0, 0)
	source.SetLooping(false)

	source.SetBuffer(s.buffer)
	return &Source{openalSource: source}
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

// LoadData creates a new sound from the given filepath
func LoadData(filepath string) (s *Data, err error) {
	var format *soundFormat
	var bytes []byte
	buffer := openal.NewBuffer()

	format, data, err := readWavFile("assets/welcome.wav")
	if err != nil {
		panic(err)
	}

	switch format.Channels {
	case 1:
		buffer.SetData(openal.FormatMono16, data[:len(data)], int32(format.Samples))
	case 2:
		buffer.SetData(openal.FormatStereo16, data[:len(data)], int32(format.Samples))
	}

	format, bytes, err = readWavFile(filepath)
	if err != nil {
		return
	}

	switch format.Channels {
	case 1:
		buffer.SetData(openal.FormatMono16, data[:len(data)], int32(format.Samples))
	case 2:
		buffer.SetData(openal.FormatStereo16, data[:len(data)], int32(format.Samples))
	}

	s = &Data{Bytes: bytes, Format: format, buffer: buffer}

	return
}

func readWavFile(path string) (*soundFormat, []byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	var buff [4]byte
	f.Read(buff[:4])

	if string(buff[:4]) != "RIFF" {
		return nil, nil, fmt.Errorf("Not a WAV file.\n")
	}

	var size int32
	binary.Read(f, binary.LittleEndian, &size)

	f.Read(buff[:4])

	if string(buff[:4]) != "WAVE" {
		return nil, nil, fmt.Errorf("Not a WAV file.\n")
	}

	f.Read(buff[:4])

	if string(buff[:4]) != "fmt " {
		return nil, nil, fmt.Errorf("Not a WAV file.\n")
	}

	binary.Read(f, binary.LittleEndian, &size)

	var format soundFormat

	switch size {
	case 16:
		binary.Read(f, binary.LittleEndian, &format)
	case 18:
		var f2 soundFormat2
		binary.Read(f, binary.LittleEndian, &f2)
		format = f2.soundFormat
	case 40:
		var f3 soundFormat3
		binary.Read(f, binary.LittleEndian, &f3)
		format = f3.soundFormat
	}

	f.Read(buff[:4])

	if string(buff[:4]) != "data" {
		return nil, nil, fmt.Errorf("Not supported WAV file.\n")
	}

	binary.Read(f, binary.LittleEndian, &size)

	wavData := make([]byte, size)
	n, e := f.Read(wavData)
	if e != nil {
		return nil, nil, fmt.Errorf("Cannot read WAV data.\n")
	}
	if int32(n) != size {
		return nil, nil, fmt.Errorf("WAV data size doesnt match.\n")
	}

	return &format, wavData, nil
}

func period(freq int, samples int) float64 {
	return float64(freq) * 2 * math.Pi * (1 / float64(samples))
}

func timeToData(t time.Duration, samples int, channels int) int {
	return int((float64(samples)/(1/t.Seconds()))+0.5) * channels
}

func main() {
	device := openal.OpenDevice("")
	context := device.CreateContext()
	defer context.Destroy()

	context.Activate()

	data, err := LoadData("assets/welcome.wav")
	if err != nil {
		panic(err)
	}

	source := data.NewSource()
	source.Play()
	for source.State() == openal.Playing {
		//loop long enough to let the wave file finish
	}
	source.Pause()
	source.Stop()
	return
}
