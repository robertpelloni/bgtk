package system

import (
	"sync"
	"time"
)

// UrgencyLevel indicates the priority of an OS notification.
type UrgencyLevel int

const (
	UrgencyLow UrgencyLevel = iota
	UrgencyNormal
	UrgencyCritical
)

// NotificationItem encapsulates a desktop alert event.
type NotificationItem struct {
	ID        string
	Title     string
	Message   string
	Urgency   UrgencyLevel
	Timestamp time.Time
}

// NotificationCenter manages cross-platform desktop alerts (QSystemTrayIcon::showMessage parity).
type NotificationCenter struct {
	mu        sync.RWMutex
	history   []NotificationItem
	listeners []func(NotificationItem)
}

// NewNotificationCenter creates a new NotificationCenter.
func NewNotificationCenter() *NotificationCenter {
	return &NotificationCenter{
		history:   make([]NotificationItem, 0),
		listeners: make([]func(NotificationItem), 0),
	}
}

// Send dispatches a new notification to listeners and appends to history.
func (nc *NotificationCenter) Send(id, title, message string, urgency UrgencyLevel) NotificationItem {
	item := NotificationItem{
		ID:        id,
		Title:     title,
		Message:   message,
		Urgency:   urgency,
		Timestamp: time.Now(),
	}

	nc.mu.Lock()
	nc.history = append(nc.history, item)
	listenersCopy := make([]func(NotificationItem), len(nc.listeners))
	copy(listenersCopy, nc.listeners)
	nc.mu.Unlock()

	for _, listener := range listenersCopy {
		if listener != nil {
			listener(item)
		}
	}

	return item
}

// AddListener registers a notification callback.
func (nc *NotificationCenter) AddListener(listener func(NotificationItem)) {
	nc.mu.Lock()
	defer nc.mu.Unlock()
	nc.listeners = append(nc.listeners, listener)
}

// History returns a copy of all dispatched notifications.
func (nc *NotificationCenter) History() []NotificationItem {
	nc.mu.RLock()
	defer nc.mu.RUnlock()
	res := make([]NotificationItem, len(nc.history))
	copy(res, nc.history)
	return res
}

// Count returns the number of notifications sent.
func (nc *NotificationCenter) Count() int {
	nc.mu.RLock()
	defer nc.mu.RUnlock()
	return len(nc.history)
}

// Clear clears notification history.
func (nc *NotificationCenter) Clear() {
	nc.mu.Lock()
	defer nc.mu.Unlock()
	nc.history = make([]NotificationItem, 0)
}
