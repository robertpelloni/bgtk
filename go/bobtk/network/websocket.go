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

// WebSocketHub handles bidirectional JSON-RPC messaging between distributed Bobtk nodes.
type WebSocketHub struct {
	mu       sync.RWMutex
	port     int
	handlers map[string]RPCHandler
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
	}
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
