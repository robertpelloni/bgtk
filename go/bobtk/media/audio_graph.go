package media

import (
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
