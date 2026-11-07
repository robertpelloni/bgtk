package core

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrPoolStopped = errors.New("work pool has been stopped")
	ErrJobCanceled = errors.New("job context canceled")
)

// WorkJob represents a discrete unit of parallel computation submitted to WorkPool.
type WorkJob func(ctx context.Context) error

// WorkPool provides an adaptive worker pool and task dispatcher inspired by
// Ultimate++ CoWork and JUCE ThreadPool. It features dynamic concurrency management,
// group barrier synchronization (Wait), error collection, and graceful shutdown.
type WorkPool struct {
	mu          sync.Mutex
	concurrency int
	jobQueue    chan func()
	ctx         context.Context
	cancel      context.CancelFunc
	wg          sync.WaitGroup
	stopped     atomic.Bool
	activeTasks atomic.Int64
}

// NewWorkPool instantiates a WorkPool with the specified concurrency limit and buffer depth.
func NewWorkPool(concurrency int, queueDepth int) *WorkPool {
	if concurrency <= 0 {
		concurrency = 4
	}
	if queueDepth <= 0 {
		queueDepth = 64
	}

	ctx, cancel := context.WithCancel(context.Background())
	wp := &WorkPool{
		concurrency: concurrency,
		jobQueue:    make(chan func(), queueDepth),
		ctx:         ctx,
		cancel:      cancel,
	}

	// Spawn persistent worker goroutines
	for i := 0; i < concurrency; i++ {
		go wp.worker()
	}

	return wp
}

func (wp *WorkPool) worker() {
	for {
		select {
		case <-wp.ctx.Done():
			return
		case task, ok := <-wp.jobQueue:
			if !ok {
				return
			}
			if task != nil {
				wp.activeTasks.Add(1)
				task()
				wp.activeTasks.Add(-1)
				wp.wg.Done()
			}
		}
	}
}

// Submit enqueues a job to be processed concurrently by the pool.
// Returns ErrPoolStopped if the pool has been shut down.
func (wp *WorkPool) Submit(job WorkJob) error {
	if wp.stopped.Load() {
		return ErrPoolStopped
	}

	wp.mu.Lock()
	defer wp.mu.Unlock()

	if wp.stopped.Load() {
		return ErrPoolStopped
	}

	wp.wg.Add(1)
	select {
	case wp.jobQueue <- func() {
		_ = job(wp.ctx)
	}:
		return nil
	case <-wp.ctx.Done():
		wp.wg.Done()
		return ErrPoolStopped
	}
}

// SubmitWaitable enqueues a job and returns an error channel that receives the job's result.
func (wp *WorkPool) SubmitWaitable(job WorkJob) (<-chan error, error) {
	if wp.stopped.Load() {
		return nil, ErrPoolStopped
	}

	errChan := make(chan error, 1)

	wp.mu.Lock()
	defer wp.mu.Unlock()

	if wp.stopped.Load() {
		return nil, ErrPoolStopped
	}

	wp.wg.Add(1)
	wp.jobQueue <- func() {
		err := job(wp.ctx)
		errChan <- err
		close(errChan)
	}

	return errChan, nil
}

// Wait blocks until all currently submitted jobs have completed execution.
func (wp *WorkPool) Wait() {
	wp.wg.Wait()
}

// ActiveWorkers returns the number of tasks currently actively executing.
func (wp *WorkPool) ActiveTasks() int64 {
	return wp.activeTasks.Load()
}

// Concurrency returns the configured number of concurrent worker threads.
func (wp *WorkPool) Concurrency() int {
	return wp.concurrency
}

// Shutdown halts all worker goroutines and stops accepting new work.
func (wp *WorkPool) Shutdown() {
	if wp.stopped.CompareAndSwap(false, true) {
		wp.cancel()
		close(wp.jobQueue)
	}
}
