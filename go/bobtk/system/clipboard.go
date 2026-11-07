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
	mimeData  map[string][]byte
	listeners []ClipboardListener
}

// NewClipboard instantiates a thread-safe clipboard abstraction.
func NewClipboard() *Clipboard {
	return &Clipboard{
		mimeData:  make(map[string][]byte),
		listeners: make([]ClipboardListener, 0),
	}
}

// GetText returns the current clipboard plain text.
func (c *Clipboard) GetText() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.text
}

// HasText returns true if non-empty plain text is currently stored.
func (c *Clipboard) HasText() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.text) > 0
}

// SetText sets clipboard plain text and notifies registered listeners.
func (c *Clipboard) SetText(text string) {
	c.mu.Lock()
	c.text = text
	c.mimeData["text/plain"] = []byte(text)
	listeners := append([]ClipboardListener{}, c.listeners...)
	c.mu.Unlock()

	for _, l := range listeners {
		l(text)
	}
}

// SetData associates arbitrary binary or MIME payloads with the clipboard (QMimeData parity).
func (c *Clipboard) SetData(mimeType string, data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.mimeData[mimeType] = append([]byte{}, data...)
	if mimeType == "text/plain" {
		c.text = string(data)
	}
}

// GetData retrieves data corresponding to the requested MIME type.
func (c *Clipboard) GetData(mimeType string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.mimeData[mimeType]
	if !ok {
		return nil, false
	}
	return append([]byte{}, val...), true
}

// Clear empties the clipboard content.
func (c *Clipboard) Clear() {
	c.mu.Lock()
	c.text = ""
	c.mimeData = make(map[string][]byte)
	listeners := append([]ClipboardListener{}, c.listeners...)
	c.mu.Unlock()

	for _, l := range listeners {
		l("")
	}
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
