package media

import (
	"math"
	"sync"
)

// AudioBuffer represents a multi-channel PCM audio buffer (JUCE AudioBuffer parity).
type AudioBuffer struct {
	Channels   int
	NumSamples int
	Data       [][]float32
}

// NewAudioBuffer allocates an audio buffer.
func NewAudioBuffer(channels, numSamples int) *AudioBuffer {
	data := make([][]float32, channels)
	for c := 0; c < channels; c++ {
		data[c] = make([]float32, numSamples)
	}
	return &AudioBuffer{
		Channels:   channels,
		NumSamples: numSamples,
		Data:       data,
	}
}

// AudioProcessor defines the DSP processing interface inspired by JUCE AudioProcessor.
type AudioProcessor interface {
	ProcessBlock(buffer *AudioBuffer)
	Name() string
}

// GainProcessor scales audio buffer volume by a gain factor.
type GainProcessor struct {
	Gain float32
}

func (g *GainProcessor) Name() string { return "GainProcessor" }
func (g *GainProcessor) ProcessBlock(buffer *AudioBuffer) {
	for c := 0; c < buffer.Channels; c++ {
		for i := 0; i < buffer.NumSamples; i++ {
			buffer.Data[c][i] *= g.Gain
		}
	}
}

// SineOscillatorProcessor synthesizes a pure sine waveform (JUCE ToneGeneratorAudioProcessor parity).
type SineOscillatorProcessor struct {
	Frequency  float64
	SampleRate float64
	Amplitude  float32
	phase      float64
}

func NewSineOscillatorProcessor(freq, sampleRate float64, amp float32) *SineOscillatorProcessor {
	if sampleRate <= 0 {
		sampleRate = 48000
	}
	if amp <= 0 {
		amp = 1.0
	}
	return &SineOscillatorProcessor{
		Frequency:  freq,
		SampleRate: sampleRate,
		Amplitude:  amp,
		phase:      0,
	}
}

func (s *SineOscillatorProcessor) Name() string { return "SineOscillatorProcessor" }
func (s *SineOscillatorProcessor) ProcessBlock(buffer *AudioBuffer) {
	phaseIncrement := (2.0 * math.Pi * s.Frequency) / s.SampleRate
	for i := 0; i < buffer.NumSamples; i++ {
		sample := float32(math.Sin(s.phase)) * s.Amplitude
		s.phase += phaseIncrement
		if s.phase >= 2.0*math.Pi {
			s.phase -= 2.0 * math.Pi
		}
		for c := 0; c < buffer.Channels; c++ {
			buffer.Data[c][i] += sample
		}
	}
}

// LowPassFilterProcessor applies a single-pole IIR low-pass filter (JUCE IIRFilter parity).
type LowPassFilterProcessor struct {
	Cutoff     float64
	SampleRate float64
	prevOutput []float32
}

func NewLowPassFilterProcessor(cutoff, sampleRate float64, channels int) *LowPassFilterProcessor {
	if sampleRate <= 0 {
		sampleRate = 48000
	}
	if channels <= 0 {
		channels = 2
	}
	return &LowPassFilterProcessor{
		Cutoff:     cutoff,
		SampleRate: sampleRate,
		prevOutput: make([]float32, channels),
	}
}

func (lp *LowPassFilterProcessor) Name() string { return "LowPassFilterProcessor" }
func (lp *LowPassFilterProcessor) ProcessBlock(buffer *AudioBuffer) {
	if len(lp.prevOutput) < buffer.Channels {
		lp.prevOutput = make([]float32, buffer.Channels)
	}
	// Calculate smoothing coefficient alpha = dt / (RC + dt)
	dt := 1.0 / lp.SampleRate
	rc := 1.0 / (2.0 * math.Pi * lp.Cutoff)
	alpha := float32(dt / (rc + dt))
	if alpha > 1.0 {
		alpha = 1.0
	}

	for c := 0; c < buffer.Channels; c++ {
		prev := lp.prevOutput[c]
		for i := 0; i < buffer.NumSamples; i++ {
			out := prev + alpha*(buffer.Data[c][i]-prev)
			buffer.Data[c][i] = out
			prev = out
		}
		lp.prevOutput[c] = prev
	}
}

// AudioGraph chains audio processors into an execution pipeline.
type AudioGraph struct {
	mu         sync.RWMutex
	sampleRate int
	bufferSize int
	processors []AudioProcessor
}

// NewAudioGraph instantiates an audio DSP graph.
func NewAudioGraph(sampleRate, bufferSize int) *AudioGraph {
	if sampleRate <= 0 {
		sampleRate = 48000
	}
	if bufferSize <= 0 {
		bufferSize = 256
	}
	return &AudioGraph{
		sampleRate: sampleRate,
		bufferSize: bufferSize,
		processors: make([]AudioProcessor, 0),
	}
}

// AddProcessor adds a DSP node to the chain.
func (ag *AudioGraph) AddProcessor(proc AudioProcessor) {
	ag.mu.Lock()
	defer ag.mu.Unlock()
	ag.processors = append(ag.processors, proc)
}

// Process passes an audio buffer through all chained processors.
func (ag *AudioGraph) Process(buffer *AudioBuffer) {
	ag.mu.RLock()
	defer ag.mu.RUnlock()
	for _, proc := range ag.processors {
		proc.ProcessBlock(buffer)
	}
}

// ProcessorCount returns the number of active DSP nodes.
func (ag *AudioGraph) ProcessorCount() int {
	ag.mu.RLock()
	defer ag.mu.RUnlock()
	return len(ag.processors)
}
