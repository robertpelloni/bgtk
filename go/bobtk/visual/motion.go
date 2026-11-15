package visual

import (
	"sync"
	"time"
)

// MotionEvent represents a user interaction or synthetic motion event.
type MotionEvent struct {
	X         float64
	Y         float64
	Timestamp time.Time
	Synthetic bool // True if generated programmatically (e.g., during scrolling)
}

// MotionHandler defines an interface for receiving motion events.
type MotionHandler interface {
	HandleMotion(event MotionEvent)
}

// MotionEmitter defines an interface for components that can generate motion events.
type MotionEmitter interface {
	AddMotionHandler(handler MotionHandler)
	RemoveMotionHandler(handler MotionHandler)
	EmitMotion(event MotionEvent)
}

// BaseMotionEmitter provides a thread-safe foundation for emitting motion events.
// This is designed to interop cleanly with the C++ backend's synthetic motion logic
// (e.g. gtk_scrolled_window synthetic motion during scroll).
type BaseMotionEmitter struct {
	mu       sync.RWMutex
	handlers []MotionHandler
}

// NewBaseMotionEmitter creates a new BaseMotionEmitter.
func NewBaseMotionEmitter() *BaseMotionEmitter {
	return &BaseMotionEmitter{
		handlers: make([]MotionHandler, 0),
	}
}

// AddMotionHandler registers a new handler.
func (e *BaseMotionEmitter) AddMotionHandler(handler MotionHandler) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.handlers = append(e.handlers, handler)
}

// RemoveMotionHandler unregisters a handler.
func (e *BaseMotionEmitter) RemoveMotionHandler(handler MotionHandler) {
	e.mu.Lock()
	defer e.mu.Unlock()
	for i, h := range e.handlers {
		if h == handler {
			e.handlers = append(e.handlers[:i], e.handlers[i+1:]...)
			break
		}
	}
}

// EmitMotion dispatches the event to all registered handlers synchronously.
func (e *BaseMotionEmitter) EmitMotion(event MotionEvent) {
	e.mu.RLock()
	// Create a snapshot of handlers to avoid holding the lock while calling them
	handlersCopy := make([]MotionHandler, len(e.handlers))
	copy(handlersCopy, e.handlers)
	e.mu.RUnlock()

	for _, h := range handlersCopy {
		h.HandleMotion(event)
	}
}
