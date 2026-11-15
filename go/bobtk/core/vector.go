package core

import (
	"fmt"
	"sync"
)

// Vector provides a cache-coherent, fast dynamic array inspired by Ultimate++ NTL Vector.
// It pre-allocates contiguous memory chunks to minimize GC overhead and reallocation costs.
type Vector struct {
	mu       sync.RWMutex
	elements []interface{}
}

// NewVector creates a Vector with optional initial capacity.
func NewVector(capacity ...int) *Vector {
	capVal := 16
	if len(capacity) > 0 && capacity[0] > 0 {
		capVal = capacity[0]
	}
	return &Vector{
		elements: make([]interface{}, 0, capVal),
	}
}

// Add appends an element to the vector (upp::Vector::Add parity).
func (v *Vector) Add(elem interface{}) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.elements = append(v.elements, elem)
}

// Get returns the element at index.
func (v *Vector) Get(index int) (interface{}, error) {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if index < 0 || index >= len(v.elements) {
		return nil, fmt.Errorf("index %d out of bounds (len=%d)", index, len(v.elements))
	}
	return v.elements[index], nil
}

// Set updates the element at index.
func (v *Vector) Set(index int, elem interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	if index < 0 || index >= len(v.elements) {
		return fmt.Errorf("index %d out of bounds (len=%d)", index, len(v.elements))
	}
	v.elements[index] = elem
	return nil
}

// GetCount returns the number of elements in the vector (upp::Vector::GetCount parity).
func (v *Vector) GetCount() int {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return len(v.elements)
}

// Clear empties the vector without releasing the underlying capacity.
func (v *Vector) Clear() {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.elements = v.elements[:0]
}

// Capacity returns the pre-allocated slice capacity.
func (v *Vector) Capacity() int {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return cap(v.elements)
}
