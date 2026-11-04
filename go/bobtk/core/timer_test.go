package core

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestTimerPeriodicExecution(t *testing.T) {
	var ticks int32

	timer := NewTimer(func() {
		atomic.AddInt32(&ticks, 1)
	})

	timer.StartTimer(10) // 10ms interval

	if !timer.IsTimerRunning() {
		t.Fatalf("expected timer to be running")
	}

	time.Sleep(55 * time.Millisecond)

	timer.StopTimer()

	if timer.IsTimerRunning() {
		t.Fatalf("expected timer to be stopped")
	}

	finalTicks := atomic.LoadInt32(&ticks)
	if finalTicks < 3 {
		t.Errorf("expected at least 3 ticks in ~55ms, got %d", finalTicks)
	}
}
