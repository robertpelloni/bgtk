package visual

import "testing"

func TestVisualPillarInitialization(t *testing.T) {
	p := NewVisualPillar()
	if p.Initialized() {
		t.Error("new VisualPillar should be uninitialized")
	}
	p.Initialize()
	p.Initialize()
	if !p.Initialized() {
		t.Error("VisualPillar should be initialized after Initialize()")
	}
}
