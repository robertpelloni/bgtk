package network

import (
	"encoding/json"
	"fmt"
	"sync"
)

// RPCMessage represents a JSON-RPC / WebSocket payload for UI state syncing.
type RPCMessage struct {
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
	ID     string          `json:"id"`
}

// RPCHandler handles an incoming RPC method call.
type RPCHandler func(params json.RawMessage) (interface{}, error)

// WebSocketSession represents a connected WebSocket/IPC client session.
type WebSocketSession struct {
	ID         string
	RemoteAddr string
	SendQueue  chan string
}

// WebSocketHub handles bidirectional JSON-RPC messaging between distributed Bobtk nodes.
type WebSocketHub struct {
	mu       sync.RWMutex
	port     int
	handlers map[string]RPCHandler
	clients  map[string]*WebSocketSession
	running  bool
}

// NewWebSocketHub instantiates a WebSocket hub.
func NewWebSocketHub(port int) *WebSocketHub {
	if port <= 0 {
		port = 9090
	}
	return &WebSocketHub{
		port:     port,
		handlers: make(map[string]RPCHandler),
		clients:  make(map[string]*WebSocketSession),
	}
}

// RegisterClient adds a client session to the hub.
func (hub *WebSocketHub) RegisterClient(id, remoteAddr string) *WebSocketSession {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	sess := &WebSocketSession{
		ID:         id,
		RemoteAddr: remoteAddr,
		SendQueue:  make(chan string, 16),
	}
	hub.clients[id] = sess
	return sess
}

// UnregisterClient removes a client session from the hub.
func (hub *WebSocketHub) UnregisterClient(id string) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	sess, ok := hub.clients[id]
	if ok {
		close(sess.SendQueue)
		delete(hub.clients, id)
	}
}

// Broadcast sends a message to all connected client sessions.
func (hub *WebSocketHub) Broadcast(msg string) {
	hub.mu.RLock()
	defer hub.mu.RUnlock()
	for _, sess := range hub.clients {
		select {
		case sess.SendQueue <- msg:
		default:
			// Non-blocking drop if client buffer is full
		}
	}
}

// ClientCount returns the number of active sessions.
func (hub *WebSocketHub) ClientCount() int {
	hub.mu.RLock()
	defer hub.mu.RUnlock()
	return len(hub.clients)
}

// RegisterHandler binds an RPC method name to a handler function.
func (hub *WebSocketHub) RegisterHandler(method string, handler RPCHandler) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	hub.handlers[method] = handler
}

// Dispatch processes an incoming RPC message and returns JSON response payload.
func (hub *WebSocketHub) Dispatch(msg RPCMessage) (string, error) {
	hub.mu.RLock()
	handler, ok := hub.handlers[msg.Method]
	hub.mu.RUnlock()

	if !ok {
		return "", fmt.Errorf("method %q not found", msg.Method)
	}

	result, err := handler(msg.Params)
	if err != nil {
		return "", err
	}

	resBytes, err := json.Marshal(map[string]interface{}{
		"id":     msg.ID,
		"result": result,
	})
	if err != nil {
		return "", err
	}

	return string(resBytes), nil
}

// Port returns the listening port.
func (hub *WebSocketHub) Port() int {
	return hub.port
}
