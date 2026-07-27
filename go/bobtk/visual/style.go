package visual

import (
	"sync"
)

// Style represents visual attributes applied to a Widget.
// It acts as a thread-safe property bag for styling components.
// [Phase 4 Frontend Expansion]: This struct will be extended to serialize styling directly to CSS classes for the Wasm target.
type Style struct {
	mu         sync.RWMutex
	properties map[string]interface{}
}

// NewStyle initializes an empty, thread-safe Style object.
func NewStyle() *Style {
	return &Style{
		properties: make(map[string]interface{}),
	}
}

// Set sets a style property (e.g., "background-color").
func (s *Style) Set(key string, value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.properties[key] = value
}

// Get retrieves a style property. Returns nil if not found.
func (s *Style) Get(key string) interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if val, ok := s.properties[key]; ok {
		return val
	}
	return nil
}

// ParseQSS acts as a placeholder for parsing Qt Style Sheet strings.
func (s *Style) ParseQSS(qss string) error {
	// TODO: Implement QSS parsing to populate properties.
	return nil
}
