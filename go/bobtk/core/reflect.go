package core

import (
	"fmt"
	"sync"
)

// Value represents a dynamically typed property inspired by Ultimate++ Value and Qt QVariant.
type Value struct {
	data interface{}
}

// NewValue wraps any native type into a dynamic Value.
func NewValue(v interface{}) Value {
	return Value{data: v}
}

// Get returns the underlying data.
func (v Value) Get() interface{} {
	return v.data
}

// String safely casts to string or returns formatted representation.
func (v Value) String() string {
	if s, ok := v.data.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", v.data)
}

// Int safely casts to int.
func (v Value) Int() int {
	if i, ok := v.data.(int); ok {
		return i
	}
	return 0
}

// Bool safely casts to bool.
func (v Value) Bool() bool {
	if b, ok := v.data.(bool); ok {
		return b
	}
	return false
}

// ReflectContext provides a centralized registry for dynamic properties, state reflection,
// and data binding between models and UI components.
type ReflectContext struct {
	mu         sync.RWMutex
	properties map[string]Value
	listeners  map[string][]func(Value)
}

// NewReflectContext instantiates a thread-safe reflection context.
func NewReflectContext() *ReflectContext {
	return &ReflectContext{
		properties: make(map[string]Value),
		listeners:  make(map[string][]func(Value)),
	}
}

// SetProperty sets a property value and dispatches change events to registered listeners.
func (rc *ReflectContext) SetProperty(key string, val interface{}) {
	v := NewValue(val)

	rc.mu.Lock()
	rc.properties[key] = v
	cbs := append([]func(Value){}, rc.listeners[key]...)
	rc.mu.Unlock()

	for _, cb := range cbs {
		cb(v)
	}
}

// GetProperty retrieves a dynamic property by key.
func (rc *ReflectContext) GetProperty(key string) (Value, bool) {
	rc.mu.RLock()
	defer rc.mu.RUnlock()

	val, ok := rc.properties[key]
	return val, ok
}

// PropertyCount returns the total number of registered properties.
func (rc *ReflectContext) PropertyCount() int {
	rc.mu.RLock()
	defer rc.mu.RUnlock()
	return len(rc.properties)
}

// OnChange registers a callback triggered whenever the specified property is mutated.
func (rc *ReflectContext) OnChange(key string, callback func(Value)) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	rc.listeners[key] = append(rc.listeners[key], callback)
}
