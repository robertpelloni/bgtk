package network

import "fmt"

// NetworkPillar manages network operations and protocols.
type NetworkPillar struct {
	initialized bool
}

func NewNetworkPillar() *NetworkPillar {
	return &NetworkPillar{initialized: false}
}

func (n *NetworkPillar) Initialize() {
	if n.initialized { return }
	fmt.Println("NetworkPillar initialized.")
	n.initialized = true
}

// Initialized reports whether Initialize has completed.
func (n *NetworkPillar) Initialized() bool { return n.initialized }
