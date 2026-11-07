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

func (w *WidgetBase) SetEnabled(enabled bool) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.enabled = enabled
}

func (w *WidgetBase) IsEnabled() bool {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.enabled
}

func (w *WidgetBase) SetVisible(visible bool) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.visible = visible
}

func (w *WidgetBase) IsVisible() bool {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.visible
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
	if !b.enabled {
		b.mu.RUnlock()
		return
	}
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

func (l *Label) SetText(text string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Text = text
}

func (l *Label) GetText() string {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.Text
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

// TabItem represents an individual titled page within a TabWidget.
type TabItem struct {
	Title   string
	Content LayoutItem
}

// TabWidget manages multiple selectable pages with tab headers (QTabWidget parity).
type TabWidget struct {
	WidgetBase
	tabs         []TabItem
	currentIndex int
	OnTabChanged func(index int, title string)
}

func NewTabWidget() *TabWidget {
	return &TabWidget{
		WidgetBase:   WidgetBase{minW: 200, minH: 150, visible: true, enabled: true},
		tabs:         make([]TabItem, 0),
		currentIndex: -1,
	}
}

func (t *TabWidget) AddTab(content LayoutItem, title string) int {
	t.mu.Lock()
	idx := len(t.tabs)
	t.tabs = append(t.tabs, TabItem{Title: title, Content: content})
	if t.currentIndex == -1 {
		t.currentIndex = 0
	}
	t.mu.Unlock()
	return idx
}

func (t *TabWidget) SetCurrentIndex(idx int) {
	t.mu.Lock()
	if idx < 0 || idx >= len(t.tabs) {
		t.mu.Unlock()
		return
	}
	t.currentIndex = idx
	title := t.tabs[idx].Title
	cb := t.OnTabChanged
	t.mu.Unlock()
	if cb != nil {
		cb(idx, title)
	}
}

func (t *TabWidget) CurrentIndex() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.currentIndex
}

func (t *TabWidget) Count() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return len(t.tabs)
}

func (t *TabWidget) Layout(x, y, w, h int) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.bounds = Rect{X: x, Y: y, W: w, H: h}

	headerH := 30
	contentY := y + headerH
	contentH := h - headerH
	if contentH < 0 {
		contentH = 0
	}

	if t.currentIndex >= 0 && t.currentIndex < len(t.tabs) {
		t.tabs[t.currentIndex].Content.SetGeometry(x, contentY, w, contentH)
	}
}

// Splitter arranges multiple widgets with resizable partitions (QSplitter parity).
type Splitter struct {
	WidgetBase
	orientation Orientation
	items       []LayoutItem
	sizes       []int
}

func NewSplitter(orientation Orientation) *Splitter {
	return &Splitter{
		WidgetBase:  WidgetBase{minW: 100, minH: 100, visible: true, enabled: true},
		orientation: orientation,
		items:       make([]LayoutItem, 0),
		sizes:       make([]int, 0),
	}
}

func (s *Splitter) AddWidget(w LayoutItem) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, w)
	s.sizes = append(s.sizes, 100)
}

func (s *Splitter) SetSizes(sizes []int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sizes = append([]int{}, sizes...)
}

func (s *Splitter) Sizes() []int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]int{}, s.sizes...)
}

func (s *Splitter) Layout(x, y, w, h int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.bounds = Rect{X: x, Y: y, W: w, H: h}

	n := len(s.items)
	if n == 0 {
		return
	}

	totalSize := 0
	for _, sz := range s.sizes {
		if sz <= 0 {
			totalSize += 1
		} else {
			totalSize += sz
		}
	}
	if totalSize <= 0 {
		totalSize = n
	}

	if s.orientation == Horizontal {
		curX := x
		for i, it := range s.items {
			weight := 1
			if i < len(s.sizes) && s.sizes[i] > 0 {
				weight = s.sizes[i]
			}
			itemW := (w * weight) / totalSize
			it.SetGeometry(curX, y, itemW, h)
			curX += itemW
		}
	} else {
		curY := y
		for i, it := range s.items {
			weight := 1
			if i < len(s.sizes) && s.sizes[i] > 0 {
				weight = s.sizes[i]
			}
			itemH := (h * weight) / totalSize
			it.SetGeometry(x, curY, w, itemH)
			curY += itemH
		}
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
