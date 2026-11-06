package system

import (
	"sync"
)

// ClipboardListener receives notification when clipboard content changes.
type ClipboardListener func(newText string)

// Clipboard provides cross-platform clipboard access (QClipboard / GtkClipboard parity).
type Clipboard struct {
	mu        sync.RWMutex
	text      string
	listeners []ClipboardListener
}

// NewClipboard instantiates a thread-safe clipboard abstraction.
func NewClipboard() *Clipboard {
	return &Clipboard{
		listeners: make([]ClipboardListener, 0),
	}
}

// GetText returns the current clipboard plain text.
func (c *Clipboard) GetText() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.text
}

// SetText sets clipboard plain text and notifies registered listeners.
func (c *Clipboard) SetText(text string) {
	c.mu.Lock()
	c.text = text
	listeners := append([]ClipboardListener{}, c.listeners...)
	c.mu.Unlock()

	for _, l := range listeners {
		l(text)
	}
}

// Clear empties the clipboard content.
func (c *Clipboard) Clear() {
	c.SetText("")
}

// AddListener registers a notification callback triggered upon clipboard content changes.
func (c *Clipboard) AddListener(l ClipboardListener) {
	if l == nil {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.listeners = append(c.listeners, l)
}
