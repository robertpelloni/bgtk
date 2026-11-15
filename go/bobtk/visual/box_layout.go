package visual

// Orientation defines the primary direction for linear layouts (similar to Qt::Orientation).
type Orientation int

const (
	Horizontal Orientation = iota
	Vertical
)

// SizePolicy determines how a widget sizes itself within a layout (Qt::QSizePolicy parity).
type SizePolicy struct {
	HorizontalStretch int
	VerticalStretch   int
}

// LayoutItem represents an item positioned by a layout manager (QLayoutItem parity).
type LayoutItem interface {
	MinSize() (int, int)
	SetGeometry(x, y, w, h int)
}

// BoxItem holds a layout item with stretch factor.
type BoxItem struct {
	Item    LayoutItem
	Stretch int
}

// BoxLayout arranges child widgets horizontally or vertically (QBoxLayout / QHBoxLayout / QVBoxLayout parity).
type BoxLayout struct {
	orientation Orientation
	spacing     int
	margin      int
	items       []BoxItem
}

// NewHBoxLayout creates a new horizontal box layout (QHBoxLayout).
func NewHBoxLayout(spacing, margin int) *BoxLayout {
	return &BoxLayout{
		orientation: Horizontal,
		spacing:     spacing,
		margin:      margin,
		items:       make([]BoxItem, 0),
	}
}

// NewVBoxLayout creates a new vertical box layout (QVBoxLayout).
func NewVBoxLayout(spacing, margin int) *BoxLayout {
	return &BoxLayout{
		orientation: Vertical,
		spacing:     spacing,
		margin:      margin,
		items:       make([]BoxItem, 0),
	}
}

// AddItem adds an item with an optional stretch factor.
func (b *BoxLayout) AddItem(item LayoutItem, stretch int) {
	b.items = append(b.items, BoxItem{Item: item, Stretch: stretch})
}

// Count returns the number of items in the layout.
func (b *BoxLayout) Count() int {
	return len(b.items)
}

// Orientation returns the box orientation.
func (b *BoxLayout) Orientation() Orientation {
	return b.orientation
}

// Layout calculates geometries for all child items across available width and height.
func (b *BoxLayout) Layout(x, y, width, height int) {
	n := len(b.items)
	if n == 0 {
		return
	}

	totalSpacing := (n - 1) * b.spacing
	availW := width - 2*b.margin
	availH := height - 2*b.margin

	if b.orientation == Horizontal {
		availForItems := availW - totalSpacing
		if availForItems < 0 {
			availForItems = 0
		}
		totalStretch := 0
		for _, it := range b.items {
			if it.Stretch > 0 {
				totalStretch += it.Stretch
			} else {
				totalStretch += 1
			}
		}

		curX := x + b.margin
		for _, it := range b.items {
			stretch := it.Stretch
			if stretch <= 0 {
				stretch = 1
			}
			itemW := (availForItems * stretch) / totalStretch
			it.Item.SetGeometry(curX, y+b.margin, itemW, availH)
			curX += itemW + b.spacing
		}
	} else {
		availForItems := availH - totalSpacing
		if availForItems < 0 {
			availForItems = 0
		}
		totalStretch := 0
		for _, it := range b.items {
			if it.Stretch > 0 {
				totalStretch += it.Stretch
			} else {
				totalStretch += 1
			}
		}

		curY := y + b.margin
		for _, it := range b.items {
			stretch := it.Stretch
			if stretch <= 0 {
				stretch = 1
			}
			itemH := (availForItems * stretch) / totalStretch
			it.Item.SetGeometry(x+b.margin, curY, availW, itemH)
			curY += itemH + b.spacing
		}
	}
}
