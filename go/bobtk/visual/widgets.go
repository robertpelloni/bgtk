package visual

import (
	"sync"
)

// Rect represents 2D bounding geometry.
type Rect struct {
	X, Y, W, H int
}

// WidgetBase provides foundational geometry and signal dispatch for visual widgets (QWidget parity).
type WidgetBase struct {
	mu       sync.RWMutex
	bounds   Rect
	minW     int
	minH     int
	visible  bool
	enabled  bool
	parent   *WidgetBase
	children []*WidgetBase
}

func (w *WidgetBase) MinSize() (int, int) {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.minW, w.minH
}

func (w *WidgetBase) SetGeometry(x, y, width, height int) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.bounds = Rect{X: x, Y: y, W: width, H: height}
}

func (w *WidgetBase) Geometry() Rect {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.bounds
}

// Button represents a clickable command button (QPushButton parity).
type Button struct {
	WidgetBase
	Text     string
	OnClick  func()
}

func NewButton(text string) *Button {
	return &Button{
		WidgetBase: WidgetBase{minW: 80, minH: 30, visible: true, enabled: true},
		Text:       text,
	}
}

func (b *Button) Click() {
	b.mu.RLock()
	cb := b.OnClick
	b.mu.RUnlock()
	if cb != nil {
		cb()
	}
}

// Label represents a text display label (QLabel parity).
type Label struct {
	WidgetBase
	Text string
}

func NewLabel(text string) *Label {
	return &Label{
		WidgetBase: WidgetBase{minW: 40, minH: 20, visible: true, enabled: true},
		Text:       text,
	}
}

// LineEdit represents a single-line text input field (QLineEdit parity).
type LineEdit struct {
	WidgetBase
	Text        string
	Placeholder string
	OnTextChange func(string)
}

func NewLineEdit(placeholder string) *LineEdit {
	return &LineEdit{
		WidgetBase:  WidgetBase{minW: 100, minH: 25, visible: true, enabled: true},
		Placeholder: placeholder,
	}
}

func (l *LineEdit) SetText(text string) {
	l.mu.Lock()
	l.Text = text
	cb := l.OnTextChange
	l.mu.Unlock()
	if cb != nil {
		cb(text)
	}
}

// ProgressBar displays operation progress (QProgressBar parity).
type ProgressBar struct {
	WidgetBase
	Value   int
	Minimum int
	Maximum int
}

func NewProgressBar() *ProgressBar {
	return &ProgressBar{
		WidgetBase: WidgetBase{minW: 100, minH: 20, visible: true, enabled: true},
		Value:      0,
		Minimum:    0,
		Maximum:    100,
	}
}

func (p *ProgressBar) SetValue(v int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if v < p.Minimum {
		v = p.Minimum
	}
	if v > p.Maximum {
		v = p.Maximum
	}
	p.Value = v
}
