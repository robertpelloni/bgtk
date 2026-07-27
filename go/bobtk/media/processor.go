package media

// AudioBuffer represents a block of multi-channel audio data.
type AudioBuffer struct {
	Channels   [][]float32
	NumSamples int
}

// MediaProcessor represents a unified interface for audio/MIDI processing,
// bridging Bobtk's generic media pillar with JUCE's low-latency requirements.
type MediaProcessor interface {
	// PrepareToPlay is called before processing starts.
	PrepareToPlay(sampleRate float64, samplesPerBlock int)

	// ProcessBlock processes the audio in a real-time, allocation-free context.
	ProcessBlock(buffer *AudioBuffer)

	// ReleaseResources cleans up after processing stops.
	ReleaseResources()
}

// MediaContext manages the backend engine (e.g., initializing JUCE AudioDeviceManager).
type MediaContext interface {
	Initialize() error
	RegisterProcessor(processor MediaProcessor)
	Start() error
	Stop() error
}
