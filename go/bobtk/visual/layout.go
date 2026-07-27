package visual

import (
	"sync"
)

// LayoutManager defines the interface for arranging child widgets.
// It is designed to be composable and thread-safe.
// [Phase 4 Frontend Expansion]: This interface will be expanded to support WebAssembly (Wasm) DOM layout hooks.
type LayoutManager interface {
	// AddWidget adds a widget to the layout.
	AddWidget(widget Widget)

	// RemoveWidget removes a widget from the layout.
	RemoveWidget(widget Widget)

	// Calculate layouts and positions of child widgets.
	Calculate(width, height float64)

	// SetSpacing sets the spacing between items in the layout.
	SetSpacing(spacing float64)

	// SetContentsMargins sets the margins around the layout.
	SetContentsMargins(left, top, right, bottom float64)
}

// BaseLayout provides a foundational, thread-safe implementation of LayoutManager.
type BaseLayout struct {
	mu      sync.RWMutex
	widgets []Widget
	spacing float64
	margins [4]float64 // left, top, right, bottom
}

// NewBaseLayout creates a new BaseLayout.
func NewBaseLayout() *BaseLayout {
	return &BaseLayout{
		widgets: make([]Widget, 0),
	}
}

// AddWidget appends a widget to the layout.
func (l *BaseLayout) AddWidget(widget Widget) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.widgets = append(l.widgets, widget)
}

// RemoveWidget removes a widget from the layout.
func (l *BaseLayout) RemoveWidget(widget Widget) {
	l.mu.Lock()
	defer l.mu.Unlock()
	for i, w := range l.widgets {
		if w.ID() == widget.ID() {
			l.widgets = append(l.widgets[:i], l.widgets[i+1:]...)
			break
		}
	}
}

// Calculate is a placeholder for layout computation logic.
func (l *BaseLayout) Calculate(width, height float64) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	// [Phase 4 Frontend Expansion]: Placeholder for mapping coordinates to DOM node absolute positioning.
}

// SetSpacing sets the gap between widgets.
func (l *BaseLayout) SetSpacing(spacing float64) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.spacing = spacing
}

// SetContentsMargins sets the outer margins of the layout area.
func (l *BaseLayout) SetContentsMargins(left, top, right, bottom float64) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.margins = [4]float64{left, top, right, bottom}
}
