package core

import "testing"

func TestCorePillarInitialization(t *testing.T) {
	cp := NewCorePillar()
	if cp.initialized {
		t.Errorf("Expected newly created CorePillar to be uninitialized")
	}

	cp.Initialize()

	if !cp.initialized {
		t.Errorf("Expected CorePillar to be initialized after calling Initialize()")
	}
}
