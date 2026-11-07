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

func TestMidi2NoteOffAndControlChange(t *testing.T) {
	// Test NoteOff
	noteOff := NewNoteOff(1, 2, 64, 0x4000)
	if noteOff.MessageType() != 0x4 {
		t.Errorf("expected message type 4, got %d", noteOff.MessageType())
	}
	if noteOff.Group() != 1 {
		t.Errorf("expected group 1, got %d", noteOff.Group())
	}
	if noteOff.Channel() != 2 {
		t.Errorf("expected channel 2, got %d", noteOff.Channel())
	}
	if noteOff.Status() != 0x82 {
		t.Errorf("expected status 0x82, got 0x%X", noteOff.Status())
	}
	if (noteOff.Word1 >> 16) != 0x4000 {
		t.Errorf("expected velocity 0x4000, got 0x%X", noteOff.Word1>>16)
	}

	// Test ControlChange
	cc := NewControlChange(0, 5, 7, 0x12345678)
	if cc.Channel() != 5 {
		t.Errorf("expected channel 5, got %d", cc.Channel())
	}
	if cc.Status() != 0xB5 {
		t.Errorf("expected status 0xB5, got 0x%X", cc.Status())
	}
	if (cc.Word0>>8)&0xFF != 7 {
		t.Errorf("expected CC index 7, got %d", (cc.Word0>>8)&0xFF)
	}
	if cc.Word1 != 0x12345678 {
		t.Errorf("expected 32-bit value 0x12345678, got 0x%X", cc.Word1)
	}
}

func TestMidi2PitchBendAndPolyPressure(t *testing.T) {
	// PitchBend test
	pb := NewPitchBend(0, 3, 0x80000000)
	if pb.MessageType() != 0x4 {
		t.Errorf("expected message type 4, got %d", pb.MessageType())
	}
	if pb.Channel() != 3 {
		t.Errorf("expected channel 3, got %d", pb.Channel())
	}
	if pb.Status() != 0xE3 {
		t.Errorf("expected status 0xE3, got 0x%X", pb.Status())
	}
	if pb.Word1 != 0x80000000 {
		t.Errorf("expected center pitch bend 0x80000000, got 0x%X", pb.Word1)
	}

	// PolyPressure test
	pp := NewPolyPressure(1, 4, 72, 0xFFFFFFFF)
	if pp.Group() != 1 {
		t.Errorf("expected group 1, got %d", pp.Group())
	}
	if pp.Channel() != 4 {
		t.Errorf("expected channel 4, got %d", pp.Channel())
	}
	if pp.Status() != 0xA4 {
		t.Errorf("expected status 0xA4, got 0x%X", pp.Status())
	}
	if (pp.Word0>>8)&0xFF != 72 {
		t.Errorf("expected note 72, got %d", (pp.Word0>>8)&0xFF)
	}
	if pp.Word1 != 0xFFFFFFFF {
		t.Errorf("expected maximum pressure 0xFFFFFFFF, got 0x%X", pp.Word1)
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

func TestDelayProcessor(t *testing.T) {
	delaySamples := 16
	delay := NewDelayProcessor(delaySamples, 0.5, 0.5, 2)

	buf1 := NewAudioBuffer(2, delaySamples)
	buf1.Data[0][0] = 1.0 // Impulse in first sample

	delay.ProcessBlock(buf1)
	if buf1.Data[0][0] != 0.5 { // 1.0 * (1 - mix) + 0.0 * mix = 0.5
		t.Errorf("expected dry sample attenuated by mix to 0.5, got %f", buf1.Data[0][0])
	}

	// Next block should receive the delayed echo at sample index 0
	buf2 := NewAudioBuffer(2, delaySamples)
	delay.ProcessBlock(buf2)

	if buf2.Data[0][0] != 0.5 { // 0.0 * (1 - mix) + 1.0 * mix = 0.5
		t.Errorf("expected delayed echo at sample 0 of buf2 to be 0.5, got %f", buf2.Data[0][0])
	}
}

func TestAudioBufferClearCopyMagnitude(t *testing.T) {
	buf1 := NewAudioBuffer(2, 64)
	buf1.Data[0][10] = 0.75
	buf1.Data[1][20] = -0.95

	if buf1.GetMagnitude(0) != 0.75 {
		t.Errorf("expected channel 0 magnitude 0.75, got %f", buf1.GetMagnitude(0))
	}
	if buf1.GetMagnitude(1) != 0.95 {
		t.Errorf("expected channel 1 magnitude 0.95, got %f", buf1.GetMagnitude(1))
	}

	buf2 := NewAudioBuffer(2, 64)
	buf2.CopyFrom(buf1)

	if buf2.GetMagnitude(0) != 0.75 || buf2.GetMagnitude(1) != 0.95 {
		t.Errorf("expected copied buffer to have matching magnitude")
	}

	buf2.Clear()
	if buf2.GetMagnitude(0) != 0.0 || buf2.GetMagnitude(1) != 0.0 {
		t.Errorf("expected 0.0 magnitude after Clear()")
	}
}

func TestHighPassFilterProcessor(t *testing.T) {
	hp := NewHighPassFilterProcessor(500.0, 48000, 2)
	if hp.Name() != "HighPassFilterProcessor" {
		t.Errorf("expected name HighPassFilterProcessor, got %s", hp.Name())
	}

	buf := NewAudioBuffer(2, 64)
	// Fill buffer with constant DC offset (1.0)
	for c := 0; c < 2; c++ {
		for i := 0; i < 64; i++ {
			buf.Data[c][i] = 1.0
		}
	}

	hp.ProcessBlock(buf)

	// High-pass filter must attenuate DC signal over time towards zero
	if buf.Data[0][63] >= buf.Data[0][0] {
		t.Errorf("expected high-pass filter to attenuate DC signal, sample 0: %f, sample 63: %f",
			buf.Data[0][0], buf.Data[0][63])
	}
}

func TestPeakingEQProcessor(t *testing.T) {
	// Center freq 1000Hz, Q 1.0, +6dB boost
	eq := NewPeakingEQProcessor(1000.0, 1.0, 6.0, 48000, 2)
	if eq.Name() != "PeakingEQProcessor" {
		t.Errorf("expected name PeakingEQProcessor, got %s", eq.Name())
	}

	buf := NewAudioBuffer(2, 64)
	// Impulse input at sample 0
	buf.Data[0][0] = 1.0
	buf.Data[1][0] = 1.0

	eq.ProcessBlock(buf)

	if buf.GetMagnitude(0) == 0.0 {
		t.Errorf("expected non-zero magnitude after EQ processing")
	}
}
