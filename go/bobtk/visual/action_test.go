package visual

import (
	"testing"
)

func TestActionSignalsAndToggling(t *testing.T) {
	act := NewAction("Open File")
	act.SetShortcut("Ctrl+O")
	act.SetTooltip("Opens an existing file from disk")
	act.SetCheckable(true)

	triggeredName := ""
	act.Triggered.Connect(func(arg interface{}) {
		triggeredName = arg.(string)
	})

	toggleState := false
	act.Toggled.Connect(func(arg interface{}) {
		toggleState = arg.(bool)
	})

	act.Trigger()

	if triggeredName != "Open File" {
		t.Errorf("expected Open File, got %s", triggeredName)
	}

	if !act.IsChecked() || !toggleState {
		t.Errorf("expected action to be checked after trigger")
	}

	act.Trigger()
	if act.IsChecked() || toggleState {
		t.Errorf("expected action to be unchecked after second trigger")
	}
}

func TestKeyEventModifiers(t *testing.T) {
	ev := &KeyEvent{
		Key:       65, // 'A'
		Modifiers: ControlModifier | ShiftModifier,
		Text:      "A",
	}

	if !ev.HasModifier(ControlModifier) {
		t.Errorf("expected ControlModifier to be present")
	}
	if !ev.HasModifier(ShiftModifier) {
		t.Errorf("expected ShiftModifier to be present")
	}
	if ev.HasModifier(AltModifier) {
		t.Errorf("expected AltModifier to be absent")
	}

	ev.Accept()
	if !ev.Accepted {
		t.Errorf("expected event to be accepted")
	}
}
