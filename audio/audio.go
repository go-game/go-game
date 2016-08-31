package audio

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/vova616/go-openal/openal"
)

var context *openal.Context

func init() {
	device := openal.OpenDevice("")
	context = device.CreateContext()
	context.Activate()
}

// Cleanup cleans up all audio data
func Cleanup() {
	context.Destroy()
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

func readWavFile(path string) (*soundFormat, []byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	var buff [4]byte
	f.Read(buff[:4])

	if string(buff[:4]) != "RIFF" {
		return nil, nil, fmt.Errorf("Not a WAV file")
	}

	var size int32
	binary.Read(f, binary.LittleEndian, &size)

	f.Read(buff[:4])

	if string(buff[:4]) != "WAVE" {
		return nil, nil, fmt.Errorf("Not a WAV file")
	}

	f.Read(buff[:4])

	if string(buff[:4]) != "fmt " {
		return nil, nil, fmt.Errorf("Not a WAV file")
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
		return nil, nil, fmt.Errorf("Not supported WAV file")
	}

	binary.Read(f, binary.LittleEndian, &size)

	wavData := make([]byte, size)
	n, e := f.Read(wavData)
	if e != nil {
		return nil, nil, fmt.Errorf("Cannot read WAV data")
	}
	if int32(n) != size {
		return nil, nil, fmt.Errorf("WAV data size doesnt match")
	}

	return &format, wavData, nil
}

func period(freq int, samples int) float64 {
	return float64(freq) * 2 * math.Pi * (1 / float64(samples))
}

func timeToData(t time.Duration, samples int, channels int) int {
	return int((float64(samples)/(1/t.Seconds()))+0.5) * channels
}
