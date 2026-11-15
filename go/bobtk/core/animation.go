package core

import (
	"sync"
	"time"
)

// PropertyAnimation interpolates a float value over time (porting bobgui/cpp/animation.hpp & QPropertyAnimation).
type PropertyAnimation struct {
	mu         sync.Mutex
	startValue float32
	endValue   float32
	durationMs int
	current    float32
	running    bool
	timer      *Timer
	onUpdate   func(float32)
	onFinished func()
}

// NewPropertyAnimation creates an animation instance with target range and duration.
func NewPropertyAnimation(start, end float32, durationMs int) *PropertyAnimation {
	if durationMs <= 0 {
		durationMs = 500
	}
	return &PropertyAnimation{
		startValue: start,
		endValue:   end,
		durationMs: durationMs,
		current:    start,
	}
}

// OnUpdate attaches a callback triggered on each animation frame.
func (pa *PropertyAnimation) OnUpdate(cb func(float32)) {
	pa.mu.Lock()
	defer pa.mu.Unlock()
	pa.onUpdate = cb
}

// OnFinished attaches a callback triggered when animation completes.
func (pa *PropertyAnimation) OnFinished(cb func()) {
	pa.mu.Lock()
	defer pa.mu.Unlock()
	pa.onFinished = cb
}

// Current returns current interpolated value.
func (pa *PropertyAnimation) Current() float32 {
	pa.mu.Lock()
	defer pa.mu.Unlock()
	return pa.current
}

// Step advances the animation progress by a discrete delta milliseconds.
func (pa *PropertyAnimation) Step(deltaMs int) (float32, bool) {
	pa.mu.Lock()
	progress := float32(deltaMs) / float32(pa.durationMs)
	if progress >= 1.0 {
		pa.current = pa.endValue
		updateCb := pa.onUpdate
		finishCb := pa.onFinished
		pa.running = false
		pa.mu.Unlock()

		if updateCb != nil {
			updateCb(pa.current)
		}
		if finishCb != nil {
			finishCb()
		}
		return pa.current, true
	}

	pa.current = pa.startValue + (pa.endValue-pa.startValue)*progress
	curr := pa.current
	updateCb := pa.onUpdate
	pa.mu.Unlock()

	if updateCb != nil {
		updateCb(curr)
	}
	return curr, false
}

// Start launches periodic frame updates.
func (pa *PropertyAnimation) Start() {
	pa.mu.Lock()
	if pa.running {
		pa.mu.Unlock()
		return
	}
	pa.running = true
	pa.current = pa.startValue
	startTime := time.Now()
	pa.mu.Unlock()

	pa.timer = NewTimer(func() {
		elapsed := int(time.Since(startTime) / time.Millisecond)
		_, finished := pa.Step(elapsed)
		if finished {
			pa.Stop()
		}
	})
	pa.timer.StartTimer(16) // ~60fps
}

// Stop halts the animation timer.
func (pa *PropertyAnimation) Stop() {
	pa.mu.Lock()
	defer pa.mu.Unlock()
	if pa.timer != nil {
		pa.timer.StopTimer()
		pa.timer = nil
	}
	pa.running = false
}

// IsRunning reports whether animation is active.
func (pa *PropertyAnimation) IsRunning() bool {
	pa.mu.Lock()
	defer pa.mu.Unlock()
	return pa.running
}
