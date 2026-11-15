package core

import (
	"runtime"
	"sync/atomic"
)

// FastMutex represents a lightweight user-space spin-mutex with progressive backoff,
// inspired by Ultimate++ SpinLock and juce::SpinLock. Designed for short critical sections
// (such as property mutations and GUI event queues) avoiding heavy OS context switching.
type FastMutex struct {
	state int32
}

// NewFastMutex creates a new unlocked FastMutex.
func NewFastMutex() *FastMutex {
	return &FastMutex{}
}

// Lock acquires the lock, spinning with exponential/progressive backoff until acquired.
func (m *FastMutex) Lock() {
	// Fast path: try acquiring directly
	if atomic.CompareAndSwapInt32(&m.state, 0, 1) {
		return
	}

	// Slow path: spin with yield
	spins := 0
	for {
		if spins < 16 {
			// CPU pause/spin
			spins++
		} else {
			// Yield processor timeslice
			runtime.Gosched()
		}

		if atomic.LoadInt32(&m.state) == 0 && atomic.CompareAndSwapInt32(&m.state, 0, 1) {
			return
		}
	}
}

// TryLock attempts to acquire the lock immediately without spinning.
// Returns true if acquired, false otherwise.
func (m *FastMutex) TryLock() bool {
	return atomic.CompareAndSwapInt32(&m.state, 0, 1)
}

// Unlock releases the lock.
func (m *FastMutex) Unlock() {
	atomic.StoreInt32(&m.state, 0)
}

// ReadWriteSpinLock represents a high-throughput multi-reader single-writer spinlock,
// inspired by juce::ReadWriteLock and Ultimate++ RWMutex.
// Bit 31: writer flag (1 = write locked)
// Bits 0-30: active reader count
type ReadWriteSpinLock struct {
	state int32
}

const (
	writerMask = 1 << 30
	readerMask = (1 << 30) - 1
)

// NewReadWriteSpinLock creates a new ReadWriteSpinLock.
func NewReadWriteSpinLock() *ReadWriteSpinLock {
	return &ReadWriteSpinLock{}
}

// RLock acquires a shared read lock. Multiple readers can hold the lock concurrently.
func (rw *ReadWriteSpinLock) RLock() {
	for {
		curr := atomic.LoadInt32(&rw.state)
		// Check if writer holds the lock
		if (curr & writerMask) == 0 {
			if atomic.CompareAndSwapInt32(&rw.state, curr, curr+1) {
				return
			}
		}
		runtime.Gosched()
	}
}

// TryRLock attempts to acquire a read lock without blocking.
func (rw *ReadWriteSpinLock) TryRLock() bool {
	curr := atomic.LoadInt32(&rw.state)
	if (curr & writerMask) != 0 {
		return false
	}
	return atomic.CompareAndSwapInt32(&rw.state, curr, curr+1)
}

// RUnlock releases a shared read lock.
func (rw *ReadWriteSpinLock) RUnlock() {
	for {
		curr := atomic.LoadInt32(&rw.state)
		if atomic.CompareAndSwapInt32(&rw.state, curr, curr-1) {
			return
		}
	}
}

// Lock acquires an exclusive write lock.
func (rw *ReadWriteSpinLock) Lock() {
	for {
		// Only acquire if no readers and no active writer
		if atomic.CompareAndSwapInt32(&rw.state, 0, writerMask) {
			return
		}
		runtime.Gosched()
	}
}

// TryLock attempts to acquire an exclusive write lock without blocking.
func (rw *ReadWriteSpinLock) TryLock() bool {
	return atomic.CompareAndSwapInt32(&rw.state, 0, writerMask)
}

// Unlock releases an exclusive write lock.
func (rw *ReadWriteSpinLock) Unlock() {
	atomic.StoreInt32(&rw.state, 0)
}
