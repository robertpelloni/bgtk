package media

import (
	"math"
	"sync"
)

// SynthVoice represents a single polyphonic synthesis voice (juce::SynthesizerVoice parity).
type SynthVoice struct {
	NoteNumber int
	Velocity   float32
	IsPlaying  bool
	phase      float64
	sampleRate float64
	frequency  float64
}

// NewSynthVoice creates a new voice unit.
func NewSynthVoice(sampleRate float64) *SynthVoice {
	if sampleRate <= 0 {
		sampleRate = 48000
	}
	return &SynthVoice{
		sampleRate: sampleRate,
		IsPlaying:  false,
	}
}

// StartNote triggers voice synthesis for a MIDI note and velocity.
func (v *SynthVoice) StartNote(noteNumber int, velocity float32) {
	v.NoteNumber = noteNumber
	v.Velocity = velocity
	v.IsPlaying = true
	v.phase = 0.0
	// MIDI note to frequency: 440 * 2^((note - 69) / 12)
	v.frequency = 440.0 * math.Pow(2.0, float64(noteNumber-69)/12.0)
}

// StopNote releases the voice note.
func (v *SynthVoice) StopNote() {
	v.IsPlaying = false
	v.NoteNumber = -1
}

// RenderNextBlock synthesizes voice waveform into the buffer.
func (v *SynthVoice) RenderNextBlock(buffer *AudioBuffer, startSample, numSamples int) {
	if !v.IsPlaying || v.frequency <= 0 {
		return
	}

	phaseIncrement := (2.0 * math.Pi * v.frequency) / v.sampleRate

	for i := startSample; i < startSample+numSamples && i < buffer.NumSamples; i++ {
		sample := float32(math.Sin(v.phase)) * v.Velocity
		v.phase += phaseIncrement
		if v.phase >= 2.0*math.Pi {
			v.phase -= 2.0 * math.Pi
		}

		for c := 0; c < buffer.Channels; c++ {
			buffer.Data[c][i] += sample
		}
	}
}

// Synthesizer provides a multi-voice polyphonic synthesizer engine (juce::Synthesizer parity).
type Synthesizer struct {
	mu         sync.RWMutex
	sampleRate float64
	voices     []*SynthVoice
}

// NewSynthesizer initializes a synthesizer engine with a specified polyphony count.
func NewSynthesizer(polyphony int, sampleRate float64) *Synthesizer {
	if polyphony <= 0 {
		polyphony = 16
	}
	if sampleRate <= 0 {
		sampleRate = 48000
	}

	synth := &Synthesizer{
		sampleRate: sampleRate,
		voices:     make([]*SynthVoice, polyphony),
	}

	for i := 0; i < polyphony; i++ {
		synth.voices[i] = NewSynthVoice(sampleRate)
	}

	return synth
}

// NoteOn finds an available voice and triggers synthesis.
func (s *Synthesizer) NoteOn(noteNumber int, velocity float32) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Find free or matching voice
	var targetVoice *SynthVoice
	for _, v := range s.voices {
		if !v.IsPlaying {
			targetVoice = v
			break
		}
		if v.NoteNumber == noteNumber {
			targetVoice = v
			break
		}
	}

	// Voice stealing fallback if all voices are occupied
	if targetVoice == nil && len(s.voices) > 0 {
		targetVoice = s.voices[0]
	}

	if targetVoice != nil {
		targetVoice.StartNote(noteNumber, velocity)
	}
}

// NoteOff releases voices playing the specified note.
func (s *Synthesizer) NoteOff(noteNumber int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, v := range s.voices {
		if v.IsPlaying && v.NoteNumber == noteNumber {
			v.StopNote()
		}
	}
}

// AllNotesOff releases all active voices immediately.
func (s *Synthesizer) AllNotesOff() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, v := range s.voices {
		v.StopNote()
	}
}

// ActiveVoiceCount returns the number of currently sounding voices.
func (s *Synthesizer) ActiveVoiceCount() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	count := 0
	for _, v := range s.voices {
		if v.IsPlaying {
			count++
		}
	}
	return count
}

// RenderNextBlock synthesizes all active polyphonic voices into the output buffer.
func (s *Synthesizer) RenderNextBlock(buffer *AudioBuffer) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, v := range s.voices {
		if v.IsPlaying {
			v.RenderNextBlock(buffer, 0, buffer.NumSamples)
		}
	}
}
