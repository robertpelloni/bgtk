package core

import (
	"sync"
	"sync/atomic"
)

// Task represents a discrete async unit of execution.
type Task func()

// QuantumScheduler implements high-performance task-parallelism inspired by Ultimate++ CoWork.
// It executes asynchronous worker tasks across a pooled set of worker routines without blocking UI.
type QuantumScheduler struct {
	workers       int
	taskQueue     chan Task
	wg            sync.WaitGroup
	activeWorkers int32
	executedTasks int64
	closed        bool
	mu            sync.Mutex
}

// NewQuantumScheduler instantiates a scheduler with a given worker pool capacity.
func NewQuantumScheduler(workerCount int) *QuantumScheduler {
	if workerCount <= 0 {
		workerCount = 4
	}
	qs := &QuantumScheduler{
		workers:   workerCount,
		taskQueue: make(chan Task, 1024),
	}
	qs.start()
	return qs
}

func (qs *QuantumScheduler) start() {
	for i := 0; i < qs.workers; i++ {
		go func() {
			for task := range qs.taskQueue {
				atomic.AddInt32(&qs.activeWorkers, 1)
				task()
				atomic.AddInt32(&qs.activeWorkers, -1)
				atomic.AddInt64(&qs.executedTasks, 1)
				qs.wg.Done()
			}
		}()
	}
}

// Dispatch submits a task for background concurrent execution.
func (qs *QuantumScheduler) Dispatch(t Task) {
	qs.mu.Lock()
	defer qs.mu.Unlock()
	if qs.closed {
		return
	}
	qs.wg.Add(1)
	qs.taskQueue <- t
}

// Wait blocks until all currently queued tasks have finished executing.
func (qs *QuantumScheduler) Wait() {
	qs.wg.Wait()
}

// ParallelFor executes fn(i) across all indices 0..count-1 concurrently across workers (Ultimate++ CoWork::Do / QtConcurrent::map parity).
func (qs *QuantumScheduler) ParallelFor(count int, fn func(i int)) {
	if count <= 0 || fn == nil {
		return
	}
	var forWg sync.WaitGroup
	forWg.Add(count)

	for i := 0; i < count; i++ {
		idx := i
		qs.Dispatch(func() {
			defer forWg.Done()
			fn(idx)
		})
	}

	forWg.Wait()
}

// ExecutedCount returns the cumulative number of completed tasks.
func (qs *QuantumScheduler) ExecutedCount() int64 {
	return atomic.LoadInt64(&qs.executedTasks)
}

// WorkerCount returns the number of allocated concurrency workers.
func (qs *QuantumScheduler) WorkerCount() int {
	return qs.workers
}

// Close gracefully stops the task queue after in-flight tasks finish.
func (qs *QuantumScheduler) Close() {
	qs.mu.Lock()
	defer qs.mu.Unlock()
	if !qs.closed {
		qs.closed = true
		close(qs.taskQueue)
	}
}
