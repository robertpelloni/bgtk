package visual

import (
	"sync"
)

// PaintCallback defines custom vector/canvas rendering on a Component.
type PaintCallback func(g *GraphicsContext, bounds Rect)

// Component represents the foundational building block for all user interface objects
// (porting juce::Component architecture to Go).
type Component struct {
	mu           sync.RWMutex
	name         string
	bounds       Rect
	visible      bool
	parent       *Component
	children     []*Component
	paintFn      PaintCallback
	repaintCount int
}

// NewComponent instantiates a new JUCE-style UI Component.
func NewComponent(name string) *Component {
	return &Component{
		name:     name,
		visible:  true,
		children: make([]*Component, 0),
	}
}

// Name returns the identifier of this component.
func (c *Component) Name() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.name
}

// AddAndMakeVisible adds a child component and makes it visible (juce::Component::addAndMakeVisible parity).
func (c *Component) AddAndMakeVisible(child *Component) {
	c.mu.Lock()
	defer c.mu.Unlock()

	child.parent = c
	child.visible = true
	c.children = append(c.children, child)
}

// SetBounds sets the component bounds relative to its parent (juce::Component::setBounds parity).
func (c *Component) SetBounds(x, y, width, height int) {
	c.mu.Lock()
	c.bounds = Rect{X: x, Y: y, W: width, H: height}
	c.mu.Unlock()

	c.Resized()
}

// Bounds returns the component's current bounding rectangle.
func (c *Component) Bounds() Rect {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.bounds
}

// Resized is called when the component's size changes, allowing child layout repositioning.
func (c *Component) Resized() {
	// Base implementation can be overridden or supplemented
}

// SetPaintCallback attaches a custom paint renderer.
func (c *Component) SetPaintCallback(fn PaintCallback) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.paintFn = fn
}

// Repaint marks the component dirty and executes paint callbacks down the hierarchy.
func (c *Component) Repaint() {
	c.mu.Lock()
	c.repaintCount++
	pFn := c.paintFn
	b := c.bounds
	kids := append([]*Component{}, c.children...)
	c.mu.Unlock()

	if pFn != nil {
		g := NewGraphicsContext(b)
		pFn(g, b)
	}

	for _, kid := range kids {
		if kid.IsVisible() {
			kid.Repaint()
		}
	}
}

// RepaintCount returns how many times Repaint has been triggered.
func (c *Component) RepaintCount() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.repaintCount
}

// NumChildComponents returns the count of direct child components.
func (c *Component) NumChildComponents() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.children)
}

// ChildComponent returns child at index, or nil.
func (c *Component) ChildComponent(idx int) *Component {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if idx >= 0 && idx < len(c.children) {
		return c.children[idx]
	}
	return nil
}

// Parent returns the parent component, or nil.
func (c *Component) Parent() *Component {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.parent
}

// IsVisible returns visibility state.
func (c *Component) IsVisible() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.visible
}

// SetVisible updates component visibility.
func (c *Component) SetVisible(vis bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.visible = vis
}
