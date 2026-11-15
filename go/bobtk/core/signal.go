package core

import (
	"sync"
)

// ConnectionID uniquely identifies a slot connected to a Signal.
type ConnectionID uint64

// Signal provides a thread-safe publish-subscribe event mechanism inspired by Qt6 Signals & Slots.
type Signal struct {
	mu     sync.RWMutex
	nextID ConnectionID
	slots  map[ConnectionID]func(interface{})
}

// NewSignal creates a new Signal instance.
func NewSignal() *Signal {
	return &Signal{
		nextID: 1,
		slots:  make(map[ConnectionID]func(interface{})),
	}
}

// Connect binds a slot callback to this signal (QObject::connect parity).
func (s *Signal) Connect(slot func(interface{})) ConnectionID {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.nextID
	s.nextID++
	s.slots[id] = slot
	return id
}

// Disconnect removes a connected slot (QObject::disconnect parity).
func (s *Signal) Disconnect(id ConnectionID) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.slots[id]; exists {
		delete(s.slots, id)
		return true
	}
	return false
}

// Emit broadcasts the signal payload to all currently connected slots (emit signal parity).
func (s *Signal) Emit(arg interface{}) {
	s.mu.RLock()
	callbacks := make([]func(interface{}), 0, len(s.slots))
	for _, slot := range s.slots {
		callbacks = append(callbacks, slot)
	}
	s.mu.RUnlock()

	for _, slot := range callbacks {
		slot(arg)
	}
}

// SlotCount returns the count of active slot connections.
func (s *Signal) SlotCount() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.slots)
}
