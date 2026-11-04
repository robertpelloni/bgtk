package core

import (
	"sync"
	"sync/atomic"
)

// HeapChunk represents an allocated block of memory within the heap pool.
type HeapChunk struct {
	ID   uint64
	Size int
	Data []byte
}

// HeapPool implements a fast memory slab/arena pool inspired by Ultimate++ sheap/lheap.
// It avoids repeated garbage collection spikes during high-frequency UI rendering passes.
type HeapPool struct {
	mu           sync.Mutex
	chunkSize    int
	pool         sync.Pool
	allocCount   int64
	activeChunks int64
}

// NewHeapPool instantiates a heap pool with a specified default chunk byte size.
func NewHeapPool(chunkSize int) *HeapPool {
	if chunkSize <= 0 {
		chunkSize = 4096 // 4 KB slab
	}
	hp := &HeapPool{
		chunkSize: chunkSize,
	}
	hp.pool = sync.Pool{
		New: func() interface{} {
			return make([]byte, hp.chunkSize)
		},
	}
	return hp
}

// Acquire pulls a reusable memory buffer from the slab pool.
func (hp *HeapPool) Acquire() []byte {
	atomic.AddInt64(&hp.allocCount, 1)
	atomic.AddInt64(&hp.activeChunks, 1)
	buf := hp.pool.Get().([]byte)
	for i := range buf {
		buf[i] = 0
	}
	return buf
}

// Release returns an allocated memory slab back to the pool for reuse.
func (hp *HeapPool) Release(buf []byte) {
	if len(buf) == hp.chunkSize {
		atomic.AddInt64(&hp.activeChunks, -1)
		hp.pool.Put(buf)
	}
}

// TotalAllocations returns total lifetime slabs acquired.
func (hp *HeapPool) TotalAllocations() int64 {
	return atomic.LoadInt64(&hp.allocCount)
}

// ActiveChunks returns currently leased slabs.
func (hp *HeapPool) ActiveChunks() int64 {
	return atomic.LoadInt64(&hp.activeChunks)
}

// ChunkSize returns the slab byte size.
func (hp *HeapPool) ChunkSize() int {
	return hp.chunkSize
}
