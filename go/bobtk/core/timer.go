package core

import (
	"sync"
	"time"
)

// Timer provides periodic interval callbacks (juce::Timer parity).
type Timer struct {
	mu       sync.Mutex
	ticker   *time.Ticker
	stopChan chan struct{}
	running  bool
	callback func()
	interval time.Duration
}

// NewTimer creates a new timer instance.
func NewTimer(callback func()) *Timer {
	return &Timer{
		callback: callback,
	}
}

// StartTimer starts the timer with an interval in milliseconds.
func (t *Timer) StartTimer(intervalMs int) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.running {
		t.stopLocked()
	}

	if intervalMs <= 0 {
		intervalMs = 10
	}

	t.interval = time.Duration(intervalMs) * time.Millisecond
	t.ticker = time.NewTicker(t.interval)
	t.stopChan = make(chan struct{})
	t.running = true

	go func(ticker *time.Ticker, stop <-chan struct{}) {
		for {
			select {
			case <-ticker.C:
				t.mu.Lock()
				cb := t.callback
				isRunning := t.running
				t.mu.Unlock()

				if isRunning && cb != nil {
					cb()
				}
			case <-stop:
				return
			}
		}
	}(t.ticker, t.stopChan)
}

// StopTimer stops any active timer callbacks.
func (t *Timer) StopTimer() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.stopLocked()
}

func (t *Timer) stopLocked() {
	if t.running {
		t.ticker.Stop()
		close(t.stopChan)
		t.running = false
	}
}

// IsTimerRunning reports whether the timer is currently running.
func (t *Timer) IsTimerRunning() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.running
}

// GetTimerInterval returns the configured interval in milliseconds.
func (t *Timer) GetTimerInterval() int {
	t.mu.Lock()
	defer t.mu.Unlock()
	return int(t.interval / time.Millisecond)
}
