package media

import "fmt"

// MediaPillar manages audio graphs, DSP pipelines, MIDI 2.0 streaming, and spatial media.
type MediaPillar struct {
	initialized bool
	Audio       *AudioGraph
	Midi        *MidiBus
	Synth       *Synthesizer
}

func NewMediaPillar() *MediaPillar {
	return &MediaPillar{
		initialized: false,
		Audio:       NewAudioGraph(48000, 256),
		Midi:        NewMidiBus(),
		Synth:       NewSynthesizer(16, 48000),
	}
}

func (m *MediaPillar) Initialize() {
	if m.initialized {
		return
	}
	fmt.Println("MediaPillar initialized.")
	m.initialized = true
}

func (m *MediaPillar) AudioGraph() *AudioGraph {
	return m.Audio
}

func (m *MediaPillar) MidiBus() *MidiBus {
	return m.Midi
}

func (m *MediaPillar) Synthesizer() *Synthesizer {
	return m.Synth
}
