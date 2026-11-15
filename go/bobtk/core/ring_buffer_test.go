package core

import (
	"sync"
	"testing"
	"time"
)

func TestRingBufferPushPop(t *testing.T) {
	rb := NewRingBuffer[int](4)

	if rb.Capacity() != 4 {
		t.Fatalf("expected capacity 4, got %d", rb.Capacity())
	}
	if !rb.IsEmpty() {
		t.Fatalf("expected empty buffer")
	}

	for i := 1; i <= 4; i++ {
		err := rb.Push(i, 0)
		if err != nil {
			t.Fatalf("push failed at item %d: %v", i, err)
		}
	}

	if !rb.IsFull() {
		t.Fatalf("expected buffer to be full")
	}
	if rb.Len() != 4 {
		t.Fatalf("expected length 4, got %d", rb.Len())
	}

	for i := 1; i <= 4; i++ {
		val, err := rb.Pop(0)
		if err != nil {
			t.Fatalf("pop failed at item %d: %v", i, err)
		}
		if val != i {
			t.Fatalf("expected value %d, got %d", i, val)
		}
	}

	if !rb.IsEmpty() {
		t.Fatalf("expected buffer to be empty after popping all elements")
	}
}

func TestRingBufferTryPushTryPopFullEmpty(t *testing.T) {
	rb := NewRingBuffer[string](2)

	// Test TryPop on empty
	_, err := rb.TryPop()
	if err != ErrBufferEmpty {
		t.Fatalf("expected ErrBufferEmpty, got %v", err)
	}

	// TryPush elements
	if err := rb.TryPush("alpha"); err != nil {
		t.Fatalf("failed to try push: %v", err)
	}
	if err := rb.TryPush("beta"); err != nil {
		t.Fatalf("failed to try push: %v", err)
	}

	// Test TryPush on full
	if err := rb.TryPush("gamma"); err != ErrBufferFull {
		t.Fatalf("expected ErrBufferFull, got %v", err)
	}

	// TryPop elements
	val1, err := rb.TryPop()
	if err != nil || val1 != "alpha" {
		t.Fatalf("expected alpha, got %s (err: %v)", val1, err)
	}

	val2, err := rb.TryPop()
	if err != nil || val2 != "beta" {
		t.Fatalf("expected beta, got %s (err: %v)", val2, err)
	}

	// Clear test
	_ = rb.TryPush("test-clear")
	rb.Clear()
	if rb.Len() != 0 || !rb.IsEmpty() {
		t.Fatalf("expected empty buffer after Clear()")
	}

	// Close test
	rb.Close()
	if err := rb.TryPush("closed"); err != ErrBufferClosed {
		t.Fatalf("expected ErrBufferClosed on TryPush, got %v", err)
	}
	if _, err := rb.TryPop(); err != ErrBufferClosed {
		t.Fatalf("expected ErrBufferClosed on TryPop, got %v", err)
	}
}

func TestRingBufferConcurrentProducersConsumers(t *testing.T) {
	rb := NewRingBuffer[int](16)
	numItems := 1000
	numProducers := 4
	itemsPerProducer := numItems / numProducers

	var wg sync.WaitGroup
	wg.Add(numProducers)

	for p := 0; p < numProducers; p++ {
		go func(base int) {
			defer wg.Done()
			for i := 0; i < itemsPerProducer; i++ {
				_ = rb.Push(base+i, 500*time.Millisecond)
			}
		}(p * itemsPerProducer)
	}

	consumedCount := 0
	var consumerWg sync.WaitGroup
	consumerWg.Add(1)

	go func() {
		defer consumerWg.Done()
		for i := 0; i < numItems; i++ {
			_, err := rb.Pop(1 * time.Second)
			if err == nil {
				consumedCount++
			}
		}
	}()

	wg.Wait()
	consumerWg.Wait()

	if consumedCount != numItems {
		t.Fatalf("expected %d consumed items, got %d", numItems, consumedCount)
	}
}
