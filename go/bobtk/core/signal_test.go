package core

import (
	"testing"
)

func TestSignalConnectEmitDisconnect(t *testing.T) {
	sig := NewSignal()

	received := ""
	callCount := 0

	connID := sig.Connect(func(arg interface{}) {
		received = arg.(string)
		callCount++
	})

	if sig.SlotCount() != 1 {
		t.Fatalf("expected 1 slot, got %d", sig.SlotCount())
	}

	sig.Emit("Event1")

	if received != "Event1" || callCount != 1 {
		t.Errorf("expected Event1 / 1 call, got %s / %d", received, callCount)
	}

	sig.Emit("Event2")
	if received != "Event2" || callCount != 2 {
		t.Errorf("expected Event2 / 2 calls, got %s / %d", received, callCount)
	}

	disconnected := sig.Disconnect(connID)
	if !disconnected {
		t.Errorf("expected successful disconnect")
	}

	sig.Emit("Event3")
	if callCount != 2 {
		t.Errorf("expected no additional calls after disconnect, got %d", callCount)
	}
}
