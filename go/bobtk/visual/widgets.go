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

// Slider represents a draggable or stepped value selector (QSlider parity).
type Slider struct {
	WidgetBase
	Value          int
	Minimum        int
	Maximum        int
	Orientation    Orientation
	OnValueChanged func(int)
}

func NewSlider(orientation Orientation) *Slider {
	minW, minH := 100, 20
	if orientation == Vertical {
		minW, minH = 20, 100
	}
	return &Slider{
		WidgetBase:  WidgetBase{minW: minW, minH: minH, visible: true, enabled: true},
		Value:       0,
		Minimum:     0,
		Maximum:     100,
		Orientation: orientation,
	}
}

func (s *Slider) SetValue(v int) {
	s.mu.Lock()
	if v < s.Minimum {
		v = s.Minimum
	}
	if v > s.Maximum {
		v = s.Maximum
	}
	s.Value = v
	cb := s.OnValueChanged
	s.mu.Unlock()
	if cb != nil {
		cb(v)
	}
}

// CheckBox represents a toggleable option with label (QCheckBox parity).
type CheckBox struct {
	WidgetBase
	Text      string
	Checked   bool
	OnToggled func(bool)
}

func NewCheckBox(text string) *CheckBox {
	return &CheckBox{
		WidgetBase: WidgetBase{minW: 80, minH: 22, visible: true, enabled: true},
		Text:       text,
		Checked:    false,
	}
}

func (c *CheckBox) SetChecked(checked bool) {
	c.mu.Lock()
	c.Checked = checked
	cb := c.OnToggled
	c.mu.Unlock()
	if cb != nil {
		cb(checked)
	}
}

// ComboBox represents a drop-down menu / selector (QComboBox parity).
type ComboBox struct {
	WidgetBase
	Items          []string
	CurrentIndex   int
	OnIndexChanged func(int, string)
}

func NewComboBox() *ComboBox {
	return &ComboBox{
		WidgetBase:   WidgetBase{minW: 120, minH: 26, visible: true, enabled: true},
		Items:        make([]string, 0),
		CurrentIndex: -1,
	}
}

func (c *ComboBox) AddItem(item string) {
	c.mu.Lock()
	c.Items = append(c.Items, item)
	if c.CurrentIndex == -1 {
		c.CurrentIndex = 0
	}
	c.mu.Unlock()
}

func (c *ComboBox) SetCurrentIndex(idx int) {
	c.mu.Lock()
	if idx < 0 || idx >= len(c.Items) {
		c.mu.Unlock()
		return
	}
	c.CurrentIndex = idx
	item := c.Items[idx]
	cb := c.OnIndexChanged
	c.mu.Unlock()
	if cb != nil {
		cb(idx, item)
	}
}

func (c *ComboBox) CurrentText() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.CurrentIndex >= 0 && c.CurrentIndex < len(c.Items) {
		return c.Items[c.CurrentIndex]
	}
	return ""
}

// RadioButton represents an exclusive selectable radio button (QRadioButton parity).
type RadioButton struct {
	WidgetBase
	Text       string
	GroupID    string
	Checked    bool
	OnSelected func()
}

func NewRadioButton(text, groupID string) *RadioButton {
	return &RadioButton{
		WidgetBase: WidgetBase{minW: 80, minH: 22, visible: true, enabled: true},
		Text:       text,
		GroupID:    groupID,
		Checked:    false,
	}
}

func (r *RadioButton) Select() {
	r.mu.Lock()
	r.Checked = true
	cb := r.OnSelected
	r.mu.Unlock()
	if cb != nil {
		cb()
	}
}

func (c *CheckBox) Toggle() {
	c.mu.Lock()
	c.Checked = !c.Checked
	checked := c.Checked
	cb := c.OnToggled
	c.mu.Unlock()
	if cb != nil {
		cb(checked)
	}
}
