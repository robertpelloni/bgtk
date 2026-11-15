package media_test

import (
	"testing"

	"github.com/bobtk/bobtk/bobtk/media"
)

func TestSynthesizerPolyphony(t *testing.T) {
	synth := media.NewSynthesizer(4, 48000)

	if synth.ActiveVoiceCount() != 0 {
		t.Fatalf("expected 0 active voices initially, got %d", synth.ActiveVoiceCount())
	}

	// Trigger 3 notes (C4, E4, G4 major triad)
	synth.NoteOn(60, 0.8)
	synth.NoteOn(64, 0.7)
	synth.NoteOn(67, 0.6)

	if synth.ActiveVoiceCount() != 3 {
		t.Fatalf("expected 3 active voices, got %d", synth.ActiveVoiceCount())
	}

	buf := media.NewAudioBuffer(2, 128)
	synth.RenderNextBlock(buf)

	// Verify audio waveform presence
	mag := buf.GetMagnitude(0)
	if mag <= 0.0 {
		t.Fatalf("expected synthesized audio in buffer, got magnitude %f", mag)
	}

	// Release one note
	synth.NoteOff(64)
	if synth.ActiveVoiceCount() != 2 {
		t.Fatalf("expected 2 active voices after NoteOff, got %d", synth.ActiveVoiceCount())
	}

	// All notes off
	synth.AllNotesOff()
	if synth.ActiveVoiceCount() != 0 {
		t.Fatalf("expected 0 active voices after AllNotesOff, got %d", synth.ActiveVoiceCount())
	}
}
