package core

import (
	"sync"
	"testing"
)

func TestFastMutexLockUnlock(t *testing.T) {
	m := NewFastMutex()

	// Direct lock/unlock test
	m.Lock()
	m.Unlock()

	// High concurrency counter increment test
	counter := 0
	numGoroutines := 10
	incrementsPerGoroutine := 1000

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for g := 0; g < numGoroutines; g++ {
		go func() {
			defer wg.Done()
			for i := 0; i < incrementsPerGoroutine; i++ {
				m.Lock()
				counter++
				m.Unlock()
			}
		}()
	}

	wg.Wait()

	expected := numGoroutines * incrementsPerGoroutine
	if counter != expected {
		t.Fatalf("expected counter %d, got %d", expected, counter)
	}
}

func TestFastMutexTryLock(t *testing.T) {
	m := NewFastMutex()

	if !m.TryLock() {
		t.Fatalf("expected TryLock to succeed on unlocked mutex")
	}

	if m.TryLock() {
		t.Fatalf("expected TryLock to fail on already locked mutex")
	}

	m.Unlock()

	if !m.TryLock() {
		t.Fatalf("expected TryLock to succeed after unlock")
	}
	m.Unlock()
}

func TestReadWriteSpinLockConcurrentReadersWriters(t *testing.T) {
	rw := NewReadWriteSpinLock()

	// TryRLock / TryLock tests
	if !rw.TryRLock() {
		t.Fatalf("expected TryRLock to succeed")
	}
	if !rw.TryRLock() {
		t.Fatalf("expected second TryRLock to succeed (multiple readers allowed)")
	}
	if rw.TryLock() {
		t.Fatalf("expected TryLock to fail while readers are active")
	}
	rw.RUnlock()
	rw.RUnlock()

	if !rw.TryLock() {
		t.Fatalf("expected TryLock to succeed after readers unlocked")
	}
	if rw.TryRLock() {
		t.Fatalf("expected TryRLock to fail while writer is active")
	}
	rw.Unlock()

	// High-contention concurrent readers/writers
	val := 0
	numReaders := 8
	numWriters := 2
	ops := 500

	var wg sync.WaitGroup
	wg.Add(numReaders + numWriters)

	for r := 0; r < numReaders; r++ {
		go func() {
			defer wg.Done()
			for i := 0; i < ops; i++ {
				rw.RLock()
				_ = val
				rw.RUnlock()
			}
		}()
	}

	for w := 0; w < numWriters; w++ {
		go func() {
			defer wg.Done()
			for i := 0; i < ops; i++ {
				rw.Lock()
				val++
				rw.Unlock()
			}
		}()
	}

	wg.Wait()

	expectedVal := numWriters * ops
	if val != expectedVal {
		t.Fatalf("expected val %d, got %d", expectedVal, val)
	}
}
