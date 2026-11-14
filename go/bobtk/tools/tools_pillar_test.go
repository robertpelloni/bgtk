package tools

import "testing"

func TestToolsPillarInitialization(t *testing.T) {
	p := NewToolsPillar()
	if p.Initialized() {
		t.Error("new ToolsPillar should be uninitialized")
	}
	p.Initialize()
	p.Initialize()
	if !p.Initialized() {
		t.Error("ToolsPillar should be initialized after Initialize()")
	}
}
