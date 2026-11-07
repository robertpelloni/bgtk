package visual

// StackedLayout manages multiple widgets where only one is visible at a time (QStackedLayout parity).
type StackedLayout struct {
	currentIndex int
	items        []LayoutItem
}

// NewStackedLayout creates a new stacked layout manager.
func NewStackedLayout() *StackedLayout {
	return &StackedLayout{
		currentIndex: -1,
		items:        make([]LayoutItem, 0),
	}
}

// AddWidget appends a widget/item to the stack and returns its index.
func (s *StackedLayout) AddWidget(item LayoutItem) int {
	s.items = append(s.items, item)
	if s.currentIndex == -1 {
		s.currentIndex = 0
	}
	return len(s.items) - 1
}

// CurrentIndex returns the index of the currently active view.
func (s *StackedLayout) CurrentIndex() int {
	return s.currentIndex
}

// SetCurrentIndex switches the active view.
func (s *StackedLayout) SetCurrentIndex(idx int) {
	if idx >= 0 && idx < len(s.items) {
		s.currentIndex = idx
	}
}

// Count returns the number of stacked widgets.
func (s *StackedLayout) Count() int {
	return len(s.items)
}

// CurrentWidget returns the currently active item, or nil.
func (s *StackedLayout) CurrentWidget() LayoutItem {
	if s.currentIndex >= 0 && s.currentIndex < len(s.items) {
		return s.items[s.currentIndex]
	}
	return nil
}

// Layout sizes the currently active item to fill the allocated bounds.
func (s *StackedLayout) Layout(x, y, width, height int) {
	cur := s.CurrentWidget()
	if cur != nil {
		cur.SetGeometry(x, y, width, height)
	}
}
