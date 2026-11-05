package media

import (
	"testing"
)

func TestAudioGraphDSPProcessing(t *testing.T) {
	graph := NewAudioGraph(48000, 64)
	gain := &GainProcessor{Gain: 0.5}
	graph.AddProcessor(gain)

	if graph.ProcessorCount() != 1 {
		t.Fatalf("expected 1 processor, got %d", graph.ProcessorCount())
	}

	buf := NewAudioBuffer(2, 64)
	buf.Data[0][0] = 1.0
	buf.Data[1][0] = 0.8

	graph.Process(buf)

	if buf.Data[0][0] != 0.5 {
		t.Errorf("expected 0.5 after gain DSP, got %f", buf.Data[0][0])
	}
	if buf.Data[1][0] != 0.4 {
		t.Errorf("expected 0.4 after gain DSP, got %f", buf.Data[1][0])
	}
}

func TestMidi2UMPBus(t *testing.T) {
	bus := NewMidiBus()

	received := false
	var receivedPkt UMPPacket

	bus.Subscribe(func(pkt UMPPacket) {
		received = true
		receivedPkt = pkt
	})

	// Group 0, Channel 1, Note 60 (Middle C), Velocity 0x8000
	pkt := NewNoteOn(0, 1, 60, 0x8000)
	bus.Broadcast(pkt)

	if !received {
		t.Fatalf("expected listener to receive UMP packet")
	}

	if receivedPkt.MessageType() != 0x4 {
		t.Errorf("expected message type 4 (MIDI 2.0 Channel Voice), got %d", receivedPkt.MessageType())
	}

	if (receivedPkt.Word0 >> 8)&0xFF != 60 {
		t.Errorf("expected note 60, got %d", (receivedPkt.Word0>>8)&0xFF)
	}
}

func TestSineOscillatorProcessor(t *testing.T) {
	osc := NewSineOscillatorProcessor(440.0, 48000, 1.0)
	buf := NewAudioBuffer(2, 128)

	osc.ProcessBlock(buf)

	// After synthesis, buffer should contain non-zero audio values
	hasAudio := false
	for i := 0; i < buf.NumSamples; i++ {
		if buf.Data[0][i] != 0.0 {
			hasAudio = true
			break
		}
	}
	if !hasAudio {
		t.Errorf("expected synthesized sine wave in buffer")
	}
}

func TestLowPassFilterProcessor(t *testing.T) {
	lp := NewLowPassFilterProcessor(1000.0, 48000, 2)
	buf := NewAudioBuffer(2, 64)

	// Step input
	for c := 0; c < 2; c++ {
		for i := 0; i < 64; i++ {
			buf.Data[c][i] = 1.0
		}
	}

	lp.ProcessBlock(buf)

	// Filtered step output should start below 1.0 and smoothly rise
	if buf.Data[0][0] >= 1.0 {
		t.Errorf("expected low-pass filter to attenuate initial step, got %f", buf.Data[0][0])
	}
	if buf.Data[0][63] <= buf.Data[0][0] {
		t.Errorf("expected output to rise towards 1.0, but final was %f", buf.Data[0][63])
	}
}
