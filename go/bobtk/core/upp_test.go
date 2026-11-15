package core

import (
	"testing"
)

func TestVectorOperations(t *testing.T) {
	vec := NewVector(4)
	if vec.Capacity() < 4 {
		t.Errorf("expected capacity at least 4, got %d", vec.Capacity())
	}

	vec.Add("Item0")
	vec.Add("Item1")
	vec.Add("Item2")

	if vec.GetCount() != 3 {
		t.Fatalf("expected count 3, got %d", vec.GetCount())
	}

	val, err := vec.Get(1)
	if err != nil || val.(string) != "Item1" {
		t.Errorf("unexpected value at index 1: %v (err=%v)", val, err)
	}

	if err := vec.Set(1, "UpdatedItem1"); err != nil {
		t.Fatalf("failed to set element: %v", err)
	}

	valUpdated, _ := vec.Get(1)
	if valUpdated.(string) != "UpdatedItem1" {
		t.Errorf("expected UpdatedItem1, got %v", valUpdated)
	}

	vec.Clear()
	if vec.GetCount() != 0 {
		t.Errorf("expected count 0 after clear, got %d", vec.GetCount())
	}
	if vec.Capacity() < 4 {
		t.Errorf("expected preserved capacity after clear")
	}
}

func TestHeapPoolAllocationAndReuse(t *testing.T) {
	pool := NewHeapPool(2048)

	buf1 := pool.Acquire()
	if len(buf1) != 2048 {
		t.Fatalf("expected buffer size 2048, got %d", len(buf1))
	}

	if pool.ActiveChunks() != 1 {
		t.Errorf("expected 1 active chunk, got %d", pool.ActiveChunks())
	}

	buf2 := pool.Acquire()
	if pool.ActiveChunks() != 2 {
		t.Errorf("expected 2 active chunks, got %d", pool.ActiveChunks())
	}

	pool.Release(buf1)
	if pool.ActiveChunks() != 1 {
		t.Errorf("expected 1 active chunk after release, got %d", pool.ActiveChunks())
	}

	pool.Release(buf2)
	if pool.ActiveChunks() != 0 {
		t.Errorf("expected 0 active chunks after releasing all, got %d", pool.ActiveChunks())
	}

	if pool.TotalAllocations() != 2 {
		t.Errorf("expected 2 total allocations, got %d", pool.TotalAllocations())
	}
}
