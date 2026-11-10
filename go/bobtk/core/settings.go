package core

import (
	"sync"
)

// Settings provides persistent application settings and preferences (QSettings parity).
type Settings struct {
	mu        sync.RWMutex
	store     map[string]interface{}
	listeners map[string][]func(interface{})
}

// NewSettings instantiates a Settings manager.
func NewSettings() *Settings {
	return &Settings{
		store:     make(map[string]interface{}),
		listeners: make(map[string][]func(interface{})),
	}
}

// SetValue stores a key-value setting and invokes any attached listeners.
func (s *Settings) SetValue(key string, value interface{}) {
	s.mu.Lock()
	s.store[key] = value
	cbs := append([]func(interface{}){}, s.listeners[key]...)
	s.mu.Unlock()

	for _, cb := range cbs {
		cb(value)
	}
}

// Value retrieves a stored setting, or defaultVal if absent.
func (s *Settings) Value(key string, defaultVal ...interface{}) interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if val, ok := s.store[key]; ok {
		return val
	}
	if len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return nil
}

// Contains checks if a key exists in settings.
func (s *Settings) Contains(key string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.store[key]
	return ok
}

// Remove deletes a setting key.
func (s *Settings) Remove(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.store, key)
}

// AllKeys returns all configured setting keys.
func (s *Settings) AllKeys() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	keys := make([]string, 0, len(s.store))
	for k := range s.store {
		keys = append(keys, k)
	}
	return keys
}

// OnSettingChanged registers a listener for updates to a specific setting key.
func (s *Settings) OnSettingChanged(key string, cb func(interface{})) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.listeners[key] = append(s.listeners[key], cb)
}
