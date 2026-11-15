package core

import (
	"fmt"
	"sync"
)

// Map represents an ordered key-value map container with index-based fast lookup,
// inspired by Ultimate++ upp::Map, Index, and Vector pairing.
type Map struct {
	mu     sync.RWMutex
	keys   []string
	values []interface{}
	index  map[string]int
}

// NewMap creates an empty Map container.
func NewMap(capacity ...int) *Map {
	capVal := 16
	if len(capacity) > 0 && capacity[0] > 0 {
		capVal = capacity[0]
	}
	return &Map{
		keys:   make([]string, 0, capVal),
		values: make([]interface{}, 0, capVal),
		index:  make(map[string]int),
	}
}

// Put inserts or updates a key-value entry (upp::Map::Put parity).
func (m *Map) Put(key string, val interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if idx, exists := m.index[key]; exists {
		m.values[idx] = val
		return
	}

	idx := len(m.keys)
	m.keys = append(m.keys, key)
	m.values = append(m.values, val)
	m.index[key] = idx
}

// Get returns the value associated with key.
func (m *Map) Get(key string) (interface{}, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	idx, exists := m.index[key]
	if !exists {
		return nil, false
	}
	return m.values[idx], true
}

// GetAt returns the key and value at an ordered index.
func (m *Map) GetAt(index int) (string, interface{}, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if index < 0 || index >= len(m.keys) {
		return "", nil, fmt.Errorf("index %d out of bounds (len=%d)", index, len(m.keys))
	}
	return m.keys[index], m.values[index], nil
}

// GetCount returns the number of items stored in the map.
func (m *Map) GetCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.keys)
}

// Keys returns a slice of all keys in insertion order.
func (m *Map) Keys() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]string, len(m.keys))
	copy(out, m.keys)
	return out
}

// Clear empties the map.
func (m *Map) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.keys = m.keys[:0]
	m.values = m.values[:0]
	m.index = make(map[string]int)
}
