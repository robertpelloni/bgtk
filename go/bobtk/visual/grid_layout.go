package visual

// GridItem represents an item positioned in a 2D grid with row/column spans (QGridLayout parity).
type GridItem struct {
	Item     LayoutItem
	Row      int
	Col      int
	RowSpan  int
	ColSpan  int
}

// GridLayout arranges items in a 2D matrix layout (QGridLayout parity).
type GridLayout struct {
	horizontalSpacing int
	verticalSpacing   int
	margin            int
	items             []GridItem
}

// NewGridLayout creates a new grid layout manager.
func NewGridLayout(hSpacing, vSpacing, margin int) *GridLayout {
	return &GridLayout{
		horizontalSpacing: hSpacing,
		verticalSpacing:   vSpacing,
		margin:            margin,
		items:             make([]GridItem, 0),
	}
}

// AddItem inserts an item into the grid at (row, col) with optional spans.
func (g *GridLayout) AddItem(item LayoutItem, row, col, rowSpan, colSpan int) {
	if rowSpan <= 0 {
		rowSpan = 1
	}
	if colSpan <= 0 {
		colSpan = 1
	}
	g.items = append(g.items, GridItem{
		Item:    item,
		Row:     row,
		Col:     col,
		RowSpan: rowSpan,
		ColSpan: colSpan,
	})
}

// Count returns the number of items in the grid.
func (g *GridLayout) Count() int {
	return len(g.items)
}

// Layout calculates geometries for items across rows and columns.
func (g *GridLayout) Layout(x, y, width, height int) {
	if len(g.items) == 0 {
		return
	}

	maxRow := 0
	maxCol := 0
	for _, it := range g.items {
		endRow := it.Row + it.RowSpan
		if endRow > maxRow {
			maxRow = endRow
		}
		endCol := it.Col + it.ColSpan
		if endCol > maxCol {
			maxCol = endCol
		}
	}

	if maxRow == 0 || maxCol == 0 {
		return
	}

	availW := width - 2*g.margin - (maxCol-1)*g.horizontalSpacing
	availH := height - 2*g.margin - (maxRow-1)*g.verticalSpacing
	if availW < 0 {
		availW = 0
	}
	if availH < 0 {
		availH = 0
	}

	cellW := availW / maxCol
	cellH := availH / maxRow

	for _, it := range g.items {
		itemX := x + g.margin + it.Col*(cellW+g.horizontalSpacing)
		itemY := y + g.margin + it.Row*(cellH+g.verticalSpacing)
		itemW := it.ColSpan*cellW + (it.ColSpan-1)*g.horizontalSpacing
		itemH := it.RowSpan*cellH + (it.RowSpan-1)*g.verticalSpacing

		it.Item.SetGeometry(itemX, itemY, itemW, itemH)
	}
}
