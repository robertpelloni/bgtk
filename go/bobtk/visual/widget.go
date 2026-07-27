package visual

import (
	"sync"
)

// Widget represents the base interface for all visual UI components.
// It is designed to be thread-safe and supports zero-initialization.
type Widget interface {
	MotionEmitter

	// ID returns the unique identifier for the widget.
	ID() string

	// Show makes the widget visible.
	Show()

	// Hide makes the widget invisible.
	Hide()

	// Update triggers a repaint or re-layout.
	Update()

	// Qt6 Parity Placeholders

	// SetFocus policy (Qt6 parity).
	SetFocus()

	// HasFocus returns true if widget has keyboard focus (Qt6 parity).
	HasFocus() bool
}

// BaseWidget provides a thread-safe foundation for custom widgets.
// It embeds BaseMotionEmitter to handle motion events automatically.
type BaseWidget struct {
	*BaseMotionEmitter // Embed the emitter for synthetic motion handling

	mu      sync.RWMutex
	id      string
	visible bool
	focused bool
}

// NewBaseWidget creates a new BaseWidget with the given ID.
func NewBaseWidget(id string) *BaseWidget {
	return &BaseWidget{
		BaseMotionEmitter: NewBaseMotionEmitter(),
		id:                id,
		visible:           true,
	}
}

// ID returns the widget's unique identifier.
func (w *BaseWidget) ID() string {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.id
}

// Show makes the widget visible.
func (w *BaseWidget) Show() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.visible = true
}

// Hide makes the widget invisible.
func (w *BaseWidget) Hide() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.visible = false
}

// Update triggers a re-render of the widget.
func (w *BaseWidget) Update() {
	// TODO: integrate with the event loop / rendering pipeline.
}

// SetFocus requests keyboard focus for the widget.
func (w *BaseWidget) SetFocus() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.focused = true
}

// HasFocus returns whether the widget currently has focus.
func (w *BaseWidget) HasFocus() bool {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.focused
}
