package network

import "testing"

func TestNetworkPillarInitialization(t *testing.T) {
	p := NewNetworkPillar()
	if p.Initialized() {
		t.Error("new NetworkPillar should be uninitialized")
	}
	p.Initialize()
	p.Initialize()
	if !p.Initialized() {
		t.Error("NetworkPillar should be initialized after Initialize()")
	}
}
