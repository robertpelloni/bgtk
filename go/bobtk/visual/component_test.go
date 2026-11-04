package visual

import (
	"testing"
)

func TestComponentHierarchyAndBounds(t *testing.T) {
	parent := NewComponent("MainView")
	parent.SetBounds(0, 0, 800, 600)

	child1 := NewComponent("EqualizerView")
	child2 := NewComponent("MidiMonitor")

	parent.AddAndMakeVisible(child1)
	parent.AddAndMakeVisible(child2)

	if parent.NumChildComponents() != 2 {
		t.Fatalf("expected 2 child components, got %d", parent.NumChildComponents())
	}

	if child1.Parent() != parent || child2.Parent() != parent {
		t.Errorf("child components should point to parent")
	}

	child1.SetBounds(10, 10, 300, 200)
	b1 := child1.Bounds()
	if b1.X != 10 || b1.Y != 10 || b1.W != 300 || b1.H != 200 {
		t.Errorf("unexpected bounds for child1: %+v", b1)
	}

	paintCalled := 0
	child1.SetPaintCallback(func(c *Component, b Rect) {
		paintCalled++
	})

	parent.Repaint()

	if child1.RepaintCount() == 0 {
		t.Errorf("expected child1 to have received repaint")
	}
	if paintCalled == 0 {
		t.Errorf("expected custom paint callback to be invoked")
	}
}
