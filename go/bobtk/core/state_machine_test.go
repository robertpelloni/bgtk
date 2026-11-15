package core

import (
	"testing"
)

func TestStateMachineTransitions(t *testing.T) {
	sIdle := NewStateNode("Idle")
	sActive := NewStateNode("Active")
	sPaused := NewStateNode("Paused")

	sIdle.AddTransition("START", "Active")
	sActive.AddTransition("PAUSE", "Paused")
	sPaused.AddTransition("RESUME", "Active")
	sActive.AddTransition("STOP", "Idle")

	sm := NewStateMachine(sIdle)
	sm.AddState(sActive)
	sm.AddState(sPaused)

	if sm.CurrentState() != "Idle" {
		t.Fatalf("expected initial state Idle, got %s", sm.CurrentState())
	}

	nextState, err := sm.PostEvent("START")
	if err != nil || nextState != "Active" {
		t.Errorf("expected transition to Active, got %s (err=%v)", nextState, err)
	}

	nextState, err = sm.PostEvent("PAUSE")
	if err != nil || nextState != "Paused" {
		t.Errorf("expected transition to Paused, got %s (err=%v)", nextState, err)
	}

	_, err = sm.PostEvent("INVALID_EVENT")
	if err == nil {
		t.Errorf("expected error on invalid event")
	}

	if sm.CurrentState() != "Paused" {
		t.Errorf("expected state to remain Paused after invalid event")
	}
}
