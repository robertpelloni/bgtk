package core

import (
	"sync"
)

// ValueTree represents a hierarchical data structure inspired by juce::ValueTree.
// It supports typed properties, parent-child nesting, child search, and reactive change listeners.
type ValueTree struct {
	mu         sync.RWMutex
	tag        string
	properties map[string]interface{}
	children   []*ValueTree
	parent     *ValueTree
	listeners  []func(tree *ValueTree, property string)
}

// NewValueTree creates an empty ValueTree node with a semantic tag name.
func NewValueTree(tag string) *ValueTree {
	return &ValueTree{
		tag:        tag,
		properties: make(map[string]interface{}),
		children:   make([]*ValueTree, 0),
		listeners:  make([]func(tree *ValueTree, property string), 0),
	}
}

// Type returns the node tag identifier.
func (vt *ValueTree) Type() string {
	vt.mu.RLock()
	defer vt.mu.RUnlock()
	return vt.tag
}

// SetProperty sets an attribute value and notifies registered listeners.
func (vt *ValueTree) SetProperty(name string, value interface{}) {
	vt.mu.Lock()
	vt.properties[name] = value
	listenersCopy := make([]func(tree *ValueTree, property string), len(vt.listeners))
	copy(listenersCopy, vt.listeners)
	vt.mu.Unlock()

	for _, l := range listenersCopy {
		if l != nil {
			l(vt, name)
		}
	}
}

// GetProperty retrieves an attribute value.
func (vt *ValueTree) GetProperty(name string) (interface{}, bool) {
	vt.mu.RLock()
	defer vt.mu.RUnlock()
	val, ok := vt.properties[name]
	return val, ok
}

// AddChild appends a child ValueTree node to this node.
func (vt *ValueTree) AddChild(child *ValueTree) {
	vt.mu.Lock()
	defer vt.mu.Unlock()
	child.mu.Lock()
	child.parent = vt
	child.mu.Unlock()
	vt.children = append(vt.children, child)
}

// NumChildren returns the number of immediate child nodes.
func (vt *ValueTree) NumChildren() int {
	vt.mu.RLock()
	defer vt.mu.RUnlock()
	return len(vt.children)
}

// GetChild returns the child at the specified index.
func (vt *ValueTree) GetChild(index int) *ValueTree {
	vt.mu.RLock()
	defer vt.mu.RUnlock()
	if index < 0 || index >= len(vt.children) {
		return nil
	}
	return vt.children[index]
}

// GetChildWithName returns the first child matching the tag name.
func (vt *ValueTree) GetChildWithName(tag string) *ValueTree {
	vt.mu.RLock()
	defer vt.mu.RUnlock()
	for _, child := range vt.children {
		if child.tag == tag {
			return child
		}
	}
	return nil
}

// Parent returns the parent node or nil if root.
func (vt *ValueTree) Parent() *ValueTree {
	vt.mu.RLock()
	defer vt.mu.RUnlock()
	return vt.parent
}

// AddListener registers a change callback for property mutations.
func (vt *ValueTree) AddListener(l func(tree *ValueTree, property string)) {
	vt.mu.Lock()
	defer vt.mu.Unlock()
	vt.listeners = append(vt.listeners, l)
}
