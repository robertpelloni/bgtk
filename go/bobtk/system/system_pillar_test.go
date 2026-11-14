package system

import "testing"

func TestSystemPillarInitialization(t *testing.T) {
	p := NewSystemPillar()
	if p.Initialized() {
		t.Error("new SystemPillar should be uninitialized")
	}
	p.Initialize()
	p.Initialize()
	if !p.Initialized() {
		t.Error("SystemPillar should be initialized after Initialize()")
	}
}
