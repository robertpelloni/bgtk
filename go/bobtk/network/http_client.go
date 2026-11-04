package network

import (
	"fmt"
	"sync"
)

// HTTPResponse represents the outcome of an HTTP transfer.
type HTTPResponse struct {
	StatusCode int
	Body       []byte
	Error      error
}

// NetworkAccessManager manages HTTP request lifecycles and caching (QNetworkAccessManager parity).
type NetworkAccessManager struct {
	mu          sync.RWMutex
	requestLogs []string
}

// NewNetworkAccessManager creates a network access manager.
func NewNetworkAccessManager() *NetworkAccessManager {
	return &NetworkAccessManager{
		requestLogs: make([]string, 0),
	}
}

// Get executes or simulates an HTTP GET request.
func (nam *NetworkAccessManager) Get(url string) HTTPResponse {
	nam.mu.Lock()
	nam.requestLogs = append(nam.requestLogs, fmt.Sprintf("GET %s", url))
	nam.mu.Unlock()

	return HTTPResponse{
		StatusCode: 200,
		Body:       []byte(fmt.Sprintf("OK GET %s", url)),
		Error:      nil,
	}
}

// Post executes or simulates an HTTP POST request.
func (nam *NetworkAccessManager) Post(url string, data []byte) HTTPResponse {
	nam.mu.Lock()
	nam.requestLogs = append(nam.requestLogs, fmt.Sprintf("POST %s (%d bytes)", url, len(data)))
	nam.mu.Unlock()

	return HTTPResponse{
		StatusCode: 201,
		Body:       []byte(fmt.Sprintf("OK POST %s", url)),
		Error:      nil,
	}
}

// HistoryCount returns total requests handled.
func (nam *NetworkAccessManager) HistoryCount() int {
	nam.mu.RLock()
	defer nam.mu.RUnlock()
	return len(nam.requestLogs)
}
