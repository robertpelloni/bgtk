package core

import (
	"errors"
	"sync"
	"time"
)

var (
	// ErrBufferFull indicates that non-blocking TryPush failed because capacity was reached.
	ErrBufferFull = errors.New("ring buffer is full")
	// ErrBufferEmpty indicates that non-blocking TryPop failed because no elements were available.
	ErrBufferEmpty = errors.New("ring buffer is empty")
	// ErrBufferTimeout indicates that blocking Push/Pop timed out before acquiring/releasing an item.
	ErrBufferTimeout = errors.New("ring buffer operation timed out")
	// ErrBufferClosed indicates operations were attempted on a closed ring buffer.
	ErrBufferClosed = errors.New("ring buffer is closed")
)

// RingBuffer represents a thread-safe circular FIFO buffer inspired by juce::AbstractFifo
// and Ultimate++ Stream/BiVector paradigms. It is designed for low-latency producer-consumer
// event passing, audio streaming, and UI telemetry queues without dynamic memory allocation spikes.
type RingBuffer[T any] struct {
	mu       sync.Mutex
	notFull  *sync.Cond
	notEmpty *sync.Cond

	storage  []T
	capacity int
	head     int // Read index
	tail     int // Write index
	count    int // Number of active elements
	closed   bool
}

// NewRingBuffer allocates and returns a new thread-safe RingBuffer with fixed capacity.
// If capacity <= 0, a default capacity of 64 is used.
func NewRingBuffer[T any](capacity int) *RingBuffer[T] {
	if capacity <= 0 {
		capacity = 64
	}
	rb := &RingBuffer[T]{
		storage:  make([]T, capacity),
		capacity: capacity,
	}
	rb.notFull = sync.NewCond(&rb.mu)
	rb.notEmpty = sync.NewCond(&rb.mu)
	return rb
}

// Capacity returns the maximum number of items the buffer can hold.
func (rb *RingBuffer[T]) Capacity() int {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	return rb.capacity
}

// Len returns the current number of items stored in the buffer.
func (rb *RingBuffer[T]) Len() int {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	return rb.count
}

// IsFull returns true if the buffer has reached maximum capacity.
func (rb *RingBuffer[T]) IsFull() bool {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	return rb.count == rb.capacity
}

// IsEmpty returns true if no elements are currently queued in the buffer.
func (rb *RingBuffer[T]) IsEmpty() bool {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	return rb.count == 0
}

// TryPush attempts to add an item to the buffer without blocking.
// Returns ErrBufferFull if the buffer is already at capacity, or ErrBufferClosed if closed.
func (rb *RingBuffer[T]) TryPush(item T) error {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	if rb.closed {
		return ErrBufferClosed
	}
	if rb.count == rb.capacity {
		return ErrBufferFull
	}

	rb.storage[rb.tail] = item
	rb.tail = (rb.tail + 1) % rb.capacity
	rb.count++

	rb.notEmpty.Signal()
	return nil
}

// TryPop attempts to retrieve and remove the oldest item from the buffer without blocking.
// Returns ErrBufferEmpty if no elements are present, or ErrBufferClosed if closed and empty.
func (rb *RingBuffer[T]) TryPop() (T, error) {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	var zero T
	if rb.count == 0 {
		if rb.closed {
			return zero, ErrBufferClosed
		}
		return zero, ErrBufferEmpty
	}

	item := rb.storage[rb.head]
	rb.storage[rb.head] = zero // avoid memory leaks for references
	rb.head = (rb.head + 1) % rb.capacity
	rb.count--

	rb.notFull.Signal()
	return item, nil
}

// Push blocks until space is available in the buffer, then enqueues the item.
// If timeout > 0, it will return ErrBufferTimeout if space is not freed within that duration.
func (rb *RingBuffer[T]) Push(item T, timeout time.Duration) error {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	if rb.closed {
		return ErrBufferClosed
	}

	if timeout <= 0 {
		for rb.count == rb.capacity && !rb.closed {
			rb.notFull.Wait()
		}
		if rb.closed {
			return ErrBufferClosed
		}
	} else {
		deadline := time.Now().Add(timeout)
		for rb.count == rb.capacity && !rb.closed {
			remaining := time.Until(deadline)
			if remaining <= 0 {
				return ErrBufferTimeout
			}
			// Poll wait with timer
			rb.mu.Unlock()
			time.Sleep(1 * time.Millisecond)
			rb.mu.Lock()
		}
		if rb.closed {
			return ErrBufferClosed
		}
	}

	rb.storage[rb.tail] = item
	rb.tail = (rb.tail + 1) % rb.capacity
	rb.count++

	rb.notEmpty.Signal()
	return nil
}

// Pop blocks until an item is available in the buffer, then dequeues and returns it.
// If timeout > 0, it will return ErrBufferTimeout if no element is available within that duration.
func (rb *RingBuffer[T]) Pop(timeout time.Duration) (T, error) {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	var zero T
	if timeout <= 0 {
		for rb.count == 0 && !rb.closed {
			rb.notEmpty.Wait()
		}
		if rb.count == 0 && rb.closed {
			return zero, ErrBufferClosed
		}
	} else {
		deadline := time.Now().Add(timeout)
		for rb.count == 0 && !rb.closed {
			remaining := time.Until(deadline)
			if remaining <= 0 {
				return zero, ErrBufferTimeout
			}
			rb.mu.Unlock()
			time.Sleep(1 * time.Millisecond)
			rb.mu.Lock()
		}
		if rb.count == 0 && rb.closed {
			return zero, ErrBufferClosed
		}
	}

	item := rb.storage[rb.head]
	rb.storage[rb.head] = zero
	rb.head = (rb.head + 1) % rb.capacity
	rb.count--

	rb.notFull.Signal()
	return item, nil
}

// Clear removes all elements and resets read/write pointers.
func (rb *RingBuffer[T]) Clear() {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	var zero T
	for i := range rb.storage {
		rb.storage[i] = zero
	}
	rb.head = 0
	rb.tail = 0
	rb.count = 0
	rb.notFull.Broadcast()
}

// Close closes the ring buffer and wakes up all waiting producers and consumers.
func (rb *RingBuffer[T]) Close() {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	rb.closed = true
	rb.notFull.Broadcast()
	rb.notEmpty.Broadcast()
}
