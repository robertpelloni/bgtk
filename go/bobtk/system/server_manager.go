package system

import (
	"fmt"
	"sync"
)

// ServerStatus describes the current lifecycle state of a managed server.
type ServerStatus string

const (
	StatusStopped  ServerStatus = "STOPPED"
	StatusStarting ServerStatus = "STARTING"
	StatusRunning  ServerStatus = "RUNNING"
	StatusError    ServerStatus = "ERROR"
)

// ServerProcess represents a managed background server/daemon.
type ServerProcess struct {
	ID          string
	Name        string
	Description string
	Port        int
	Status      ServerStatus
	StartFn     func() error
	StopFn      func() error
}

// ServerManager provides centralized management, health tracking, and graceful
// shutdown for all background servers within the Bobtk architecture.
type ServerManager struct {
	mu      sync.RWMutex
	servers map[string]*ServerProcess
}

// NewServerManager instantiates a thread-safe server manager.
func NewServerManager() *ServerManager {
	return &ServerManager{
		servers: make(map[string]*ServerProcess),
	}
}

// RegisterServer registers a new server process for lifecycle monitoring.
func (sm *ServerManager) RegisterServer(id, name, desc string, port int, stopFn func() error) *ServerProcess {
	return sm.RegisterServerWithStart(id, name, desc, port, nil, stopFn)
}

// RegisterServerWithStart registers a new server process with both start and stop lifecycle handlers.
func (sm *ServerManager) RegisterServerWithStart(id, name, desc string, port int, startFn func() error, stopFn func() error) *ServerProcess {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	proc := &ServerProcess{
		ID:          id,
		Name:        name,
		Description: desc,
		Port:        port,
		Status:      StatusRunning,
		StartFn:     startFn,
		StopFn:      stopFn,
	}
	sm.servers[id] = proc
	return proc
}

// StartServer starts a stopped or error server process.
func (sm *ServerManager) StartServer(id string) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	proc, ok := sm.servers[id]
	if !ok {
		return fmt.Errorf("server %q not found", id)
	}
	if proc.Status == StatusRunning {
		return nil
	}
	proc.Status = StatusStarting
	if proc.StartFn != nil {
		if err := proc.StartFn(); err != nil {
			proc.Status = StatusError
			return err
		}
	}
	proc.Status = StatusRunning
	return nil
}

// StopServer stops a specific server by ID.
func (sm *ServerManager) StopServer(id string) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	proc, ok := sm.servers[id]
	if !ok {
		return fmt.Errorf("server %q not found", id)
	}
	if proc.Status == StatusStopped {
		return nil
	}
	if proc.StopFn != nil {
		if err := proc.StopFn(); err != nil {
			proc.Status = StatusError
			return err
		}
	}
	proc.Status = StatusStopped
	return nil
}

// RestartServer gracefully restarts an active or stopped server.
func (sm *ServerManager) RestartServer(id string) error {
	sm.mu.Lock()
	proc, ok := sm.servers[id]
	sm.mu.Unlock()

	if !ok {
		return fmt.Errorf("server %q not found", id)
	}

	if proc.Status == StatusRunning {
		if err := sm.StopServer(id); err != nil {
			return fmt.Errorf("failed to stop %s during restart: %w", id, err)
		}
	}
	return sm.StartServer(id)
}

// GetServer returns a single server process by ID.
func (sm *ServerManager) GetServer(id string) (*ServerProcess, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	proc, ok := sm.servers[id]
	return proc, ok
}

// StopAllServers stops all active registered servers.
func (sm *ServerManager) StopAllServers() []error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	var errs []error
	for _, proc := range sm.servers {
		if proc.Status == StatusRunning {
			if proc.StopFn != nil {
				if err := proc.StopFn(); err != nil {
					proc.Status = StatusError
					errs = append(errs, fmt.Errorf("failed to stop %s: %w", proc.Name, err))
					continue
				}
			}
			proc.Status = StatusStopped
		}
	}
	return errs
}

// GetServers returns a slice of all registered server processes.
func (sm *ServerManager) GetServers() []*ServerProcess {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	out := make([]*ServerProcess, 0, len(sm.servers))
	for _, proc := range sm.servers {
		out = append(out, proc)
	}
	return out
}

// RunningCount returns the number of currently running servers.
func (sm *ServerManager) RunningCount() int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	count := 0
	for _, proc := range sm.servers {
		if proc.Status == StatusRunning {
			count++
		}
	}
	return count
}
