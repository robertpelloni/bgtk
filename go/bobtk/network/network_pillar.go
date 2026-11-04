package network

import "fmt"

// NetworkPillar manages network operations, protocols, and remote display servers.
type NetworkPillar struct {
	initialized bool
	Broadway    *BroadwayGateway
	WSHub       *WebSocketHub
	HTTP        *NetworkAccessManager
	Sync        *SyncEngine
}

func NewNetworkPillar() *NetworkPillar {
	return &NetworkPillar{
		initialized: false,
		Broadway:    NewBroadwayGateway(8085),
		WSHub:       NewWebSocketHub(9090),
		HTTP:        NewNetworkAccessManager(),
		Sync:        NewSyncEngine(),
	}
}

func (n *NetworkPillar) Initialize() {
	if n.initialized {
		return
	}
	_ = n.Broadway.Start()
	fmt.Println("NetworkPillar initialized.")
	n.initialized = true
}

func (n *NetworkPillar) Gateway() *BroadwayGateway {
	return n.Broadway
}

func (n *NetworkPillar) Hub() *WebSocketHub {
	return n.WSHub
}

func (n *NetworkPillar) HTTPManager() *NetworkAccessManager {
	return n.HTTP
}

func (n *NetworkPillar) SyncEngine() *SyncEngine {
	return n.Sync
}
