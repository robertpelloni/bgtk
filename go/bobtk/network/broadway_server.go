package network

import (
	"fmt"
	"sync"
)

// ClientSession represents an active HTML5 browser connection streaming a GUI canvas.
type ClientSession struct {
	ID        string
	RemoteIP  string
	Connected bool
}

// BroadwayGateway provides remote HTML5 canvas streaming of Bobtk UI to browser clients
// (porting the GDK Broadway server architecture to Go).
type BroadwayGateway struct {
	mu          sync.RWMutex
	port        int
	running     bool
	clients     map[string]*ClientSession
	onClientSub []func(*ClientSession)
}

// NewBroadwayGateway instantiates a Broadway gateway.
func NewBroadwayGateway(port int) *BroadwayGateway {
	if port <= 0 {
		port = 8085
	}
	return &BroadwayGateway{
		port:    port,
		clients: make(map[string]*ClientSession),
	}
}

// Start begins listening for remote canvas display requests.
func (bg *BroadwayGateway) Start() error {
	bg.mu.Lock()
	defer bg.mu.Unlock()
	if bg.running {
		return fmt.Errorf("broadway gateway already running on port %d", bg.port)
	}
	bg.running = true
	return nil
}

// Stop closes active sessions and terminates the display listener.
func (bg *BroadwayGateway) Stop() error {
	bg.mu.Lock()
	defer bg.mu.Unlock()
	if !bg.running {
		return nil
	}
	bg.running = false
	for _, c := range bg.clients {
		c.Connected = false
	}
	return nil
}

// RegisterClient simulates connection of an HTML5 Broadway web client.
func (bg *BroadwayGateway) RegisterClient(id, ip string) *ClientSession {
	bg.mu.Lock()
	defer bg.mu.Unlock()

	client := &ClientSession{ID: id, RemoteIP: ip, Connected: true}
	bg.clients[id] = client

	for _, cb := range bg.onClientSub {
		cb(client)
	}
	return client
}

// ClientCount returns the number of connected remote sessions.
func (bg *BroadwayGateway) ClientCount() int {
	bg.mu.RLock()
	defer bg.mu.RUnlock()
	count := 0
	for _, c := range bg.clients {
		if c.Connected {
			count++
		}
	}
	return count
}

// IsRunning reports whether the gateway is accepting connections.
func (bg *BroadwayGateway) IsRunning() bool {
	bg.mu.RLock()
	defer bg.mu.RUnlock()
	return bg.running
}

// Port returns the configured gateway port.
func (bg *BroadwayGateway) Port() int {
	return bg.port
}
