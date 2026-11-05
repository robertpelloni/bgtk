package visual

import (
	"testing"
)

func TestHBoxLayoutCalculations(t *testing.T) {
	btn1 := NewButton("One")
	btn2 := NewButton("Two")

	hbox := NewHBoxLayout(10, 5)
	hbox.AddItem(btn1, 1)
	hbox.AddItem(btn2, 2)

	// Available width: 310, margin: 5 on each side -> 300 avail, spacing: 10 -> 290 for items
	// Stretch total: 1 + 2 = 3.
	// btn1 gets (290 * 1)/3 = 96
	// btn2 gets (290 * 2)/3 = 193
	hbox.Layout(0, 0, 310, 100)

	g1 := btn1.Geometry()
	g2 := btn2.Geometry()

	if g1.X != 5 {
		t.Errorf("expected btn1 X=5, got %d", g1.X)
	}
	if g1.W <= 0 || g2.W <= 0 {
		t.Errorf("invalid calculated widths: g1.W=%d, g2.W=%d", g1.W, g2.W)
	}
	if g2.W <= g1.W {
		t.Errorf("expected btn2 with stretch 2 to be wider than btn1 with stretch 1")
	}
}

func TestVBoxLayoutCalculations(t *testing.T) {
	lbl := NewLabel("Header")
	pbar := NewProgressBar()

	vbox := NewVBoxLayout(10, 5)
	vbox.AddItem(lbl, 1)
	vbox.AddItem(pbar, 1)

	vbox.Layout(0, 0, 200, 210)

	g1 := lbl.Geometry()
	g2 := pbar.Geometry()

	if g1.Y != 5 {
		t.Errorf("expected lbl Y=5, got %d", g1.Y)
	}
	if g2.Y <= g1.Y {
		t.Errorf("expected pbar Y to be positioned after lbl")
	}
}

func TestGridLayoutCalculations(t *testing.T) {
	b1 := NewButton("R0C0")
	b2 := NewButton("R0C1")
	b3 := NewButton("R1C0_Span2")

	grid := NewGridLayout(5, 5, 5)
	grid.AddItem(b1, 0, 0, 1, 1)
	grid.AddItem(b2, 0, 1, 1, 1)
	grid.AddItem(b3, 1, 0, 1, 2) // span 2 columns

	grid.Layout(0, 0, 215, 115)

	g1 := b1.Geometry()
	g2 := b2.Geometry()
	g3 := b3.Geometry()

	if g2.X <= g1.X {
		t.Errorf("expected col 1 to be to the right of col 0")
	}
	if g3.W <= g1.W {
		t.Errorf("expected spanning item to be wider than single cell item")
	}
}

func TestStackedLayoutSwitching(t *testing.T) {
	viewA := NewLabel("Page A")
	viewB := NewLabel("Page B")

	stack := NewStackedLayout()
	idxA := stack.AddWidget(viewA)
	idxB := stack.AddWidget(viewB)

	if idxA != 0 || idxB != 1 {
		t.Fatalf("unexpected indices: %d, %d", idxA, idxB)
	}

	if stack.CurrentIndex() != 0 {
		t.Errorf("expected initial active index 0, got %d", stack.CurrentIndex())
	}

	stack.SetCurrentIndex(1)
	if stack.CurrentWidget() != viewB {
		t.Errorf("expected active widget to be viewB")
	}

	stack.Layout(10, 10, 400, 300)
	gB := viewB.Geometry()
	if gB.W != 400 || gB.H != 300 {
		t.Errorf("expected viewB to fill allocated stacked bounds")
	}
}

func TestWidgetSignals(t *testing.T) {
	btn := NewButton("Action")
	clicked := false
	btn.OnClick = func() {
		clicked = true
	}
	btn.Click()
	if !clicked {
		t.Errorf("expected button click handler to fire")
	}

	line := NewLineEdit("Enter value")
	textMutated := ""
	line.OnTextChange = func(s string) {
		textMutated = s
	}
	line.SetText("Bobtk-UI")
	if textMutated != "Bobtk-UI" {
		t.Errorf("expected text change handler to receive mutated text")
	}

	pbar := NewProgressBar()
	pbar.SetValue(75)
	if pbar.Value != 75 {
		t.Errorf("expected progress bar value 75, got %d", pbar.Value)
	}
}

func TestSliderWidget(t *testing.T) {
	slider := NewSlider(Horizontal)
	changedVal := -1
	slider.OnValueChanged = func(v int) {
		changedVal = v
	}

	slider.SetValue(42)
	if slider.Value != 42 || changedVal != 42 {
		t.Errorf("expected slider value 42, got %d (cb: %d)", slider.Value, changedVal)
	}

	// Test clamping
	slider.SetValue(150)
	if slider.Value != 100 {
		t.Errorf("expected upper bound clamp 100, got %d", slider.Value)
	}

	slider.SetValue(-20)
	if slider.Value != 0 {
		t.Errorf("expected lower bound clamp 0, got %d", slider.Value)
	}
}

func TestCheckBoxWidget(t *testing.T) {
	cb := NewCheckBox("Enable Telemetry")
	toggled := false
	cb.OnToggled = func(checked bool) {
		toggled = checked
	}

	cb.SetChecked(true)
	if !cb.Checked || !toggled {
		t.Errorf("expected checkbox checked true")
	}

	cb.Toggle()
	if cb.Checked || toggled {
		t.Errorf("expected checkbox toggled to false")
	}
}

func TestComboBoxWidget(t *testing.T) {
	combo := NewComboBox()
	combo.AddItem("Option Alpha")
	combo.AddItem("Option Beta")
	combo.AddItem("Option Gamma")

	if combo.CurrentText() != "Option Alpha" {
		t.Errorf("expected initial text 'Option Alpha', got %q", combo.CurrentText())
	}

	selectedIdx := -1
	selectedVal := ""
	combo.OnIndexChanged = func(idx int, val string) {
		selectedIdx = idx
		selectedVal = val
	}

	combo.SetCurrentIndex(2)
	if combo.CurrentIndex != 2 || selectedIdx != 2 || selectedVal != "Option Gamma" {
		t.Errorf("expected index 2 'Option Gamma', got %d %q", selectedIdx, selectedVal)
	}
}

func TestRadioButtonWidget(t *testing.T) {
	radio := NewRadioButton("High Precision", "AudioQualityGroup")
	selected := false
	radio.OnSelected = func() {
		selected = true
	}

	radio.Select()
	if !radio.Checked || !selected {
		t.Errorf("expected radio button selected true")
	}
}

func TestTabWidget(t *testing.T) {
	tabWidget := NewTabWidget()
	page1 := NewLabel("Page 1 Content")
	page2 := NewLabel("Page 2 Content")

	idx0 := tabWidget.AddTab(page1, "Overview")
	idx1 := tabWidget.AddTab(page2, "Diagnostics")

	if idx0 != 0 || idx1 != 1 || tabWidget.Count() != 2 {
		t.Fatalf("unexpected tab indices or count: idx0=%d, idx1=%d, count=%d", idx0, idx1, tabWidget.Count())
	}

	tabWidget.Layout(0, 0, 400, 300)
	g1 := page1.Geometry()
	if g1.W != 400 || g1.H != 270 || g1.Y != 30 {
		t.Errorf("expected page1 layout W=400, H=270, Y=30, got %+v", g1)
	}

	switchedIdx := -1
	switchedTitle := ""
	tabWidget.OnTabChanged = func(idx int, title string) {
		switchedIdx = idx
		switchedTitle = title
	}

	tabWidget.SetCurrentIndex(1)
	if tabWidget.CurrentIndex() != 1 || switchedIdx != 1 || switchedTitle != "Diagnostics" {
		t.Errorf("expected tab switched to 1, got %d %q", switchedIdx, switchedTitle)
	}

	tabWidget.Layout(0, 0, 400, 300)
	g2 := page2.Geometry()
	if g2.W != 400 || g2.H != 270 || g2.Y != 30 {
		t.Errorf("expected page2 layout W=400, H=270, Y=30, got %+v", g2)
	}
}
