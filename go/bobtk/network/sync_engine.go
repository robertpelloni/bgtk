package network

import (
	"sync"
)

// CursorPosition tracks multiplayer remote cursor coordinates.
type CursorPosition struct {
	ClientID string
	X        float32
	Y        float32
}

// SyncEngine coordinates state delta synchronization and cursor tracking across peers
// (porting bobgui/cpp/network.hpp SyncEngine to Go).
type SyncEngine struct {
	mu            sync.RWMutex
	peers         map[string]bool
	cursors       map[string]CursorPosition
	broadcastLog  [][]byte
}

// NewSyncEngine instantiates a synchronization engine.
func NewSyncEngine() *SyncEngine {
	return &SyncEngine{
		peers:        make(map[string]bool),
		cursors:      make(map[string]CursorPosition),
		broadcastLog: make([][]byte, 0),
	}
}

// ConnectPeer registers a connected peer endpoint.
func (se *SyncEngine) ConnectPeer(peerAddress string) {
	se.mu.Lock()
	defer se.mu.Unlock()
	se.peers[peerAddress] = true
}

// Broadcast dispatches a binary state delta to connected peers.
func (se *SyncEngine) Broadcast(delta []byte) {
	se.mu.Lock()
	defer se.mu.Unlock()
	se.broadcastLog = append(se.broadcastLog, delta)
}

// TrackCursor updates coordinates for a collaborator's cursor.
func (se *SyncEngine) TrackCursor(clientID string, x, y float32) {
	se.mu.Lock()
	defer se.mu.Unlock()
	se.cursors[clientID] = CursorPosition{
		ClientID: clientID,
		X:        x,
		Y:        y,
	}
}

// PeerCount returns active peer connections count.
func (se *SyncEngine) PeerCount() int {
	se.mu.RLock()
	defer se.mu.RUnlock()
	return len(se.peers)
}

// BroadcastCount returns total transmitted deltas.
func (se *SyncEngine) BroadcastCount() int {
	se.mu.RLock()
	defer se.mu.RUnlock()
	return len(se.broadcastLog)
}

// GetCursor returns tracked cursor coordinates for a client.
func (se *SyncEngine) GetCursor(clientID string) (CursorPosition, bool) {
	se.mu.RLock()
	defer se.mu.RUnlock()
	c, ok := se.cursors[clientID]
	return c, ok
}
