package media

import "testing"

func TestMediaPillarInitialization(t *testing.T) {
	p := NewMediaPillar()
	if p.Initialized() {
		t.Error("new MediaPillar should be uninitialized")
	}
	p.Initialize()
	p.Initialize() // idempotent: second call is a no-op
	if !p.Initialized() {
		t.Error("MediaPillar should be initialized after Initialize()")
	}
}
