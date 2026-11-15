package core_test

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"

	"github.com/bobtk/bobtk/bobtk/core"
)

func TestWorkPoolBasicExecution(t *testing.T) {
	wp := core.NewWorkPool(4, 16)
	defer wp.Shutdown()

	var counter atomic.Int32
	jobCount := 50

	for i := 0; i < jobCount; i++ {
		err := wp.Submit(func(ctx context.Context) error {
			counter.Add(1)
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected error submitting job: %v", err)
		}
	}

	wp.Wait()

	if counter.Load() != int32(jobCount) {
		t.Fatalf("expected counter %d, got %d", jobCount, counter.Load())
	}
}

func TestWorkPoolSubmitWaitable(t *testing.T) {
	wp := core.NewWorkPool(2, 8)
	defer wp.Shutdown()

	testErr := errors.New("sample task error")

	errCh, err := wp.SubmitWaitable(func(ctx context.Context) error {
		return testErr
	})
	if err != nil {
		t.Fatalf("failed to submit waitable task: %v", err)
	}

	select {
	case receivedErr := <-errCh:
		if !errors.Is(receivedErr, testErr) {
			t.Fatalf("expected error %v, got %v", testErr, receivedErr)
		}
	case <-time.After(1 * time.Second):
		t.Fatal("timed out waiting for task completion")
	}
}

func TestWorkPoolShutdown(t *testing.T) {
	wp := core.NewWorkPool(2, 4)
	wp.Shutdown()

	err := wp.Submit(func(ctx context.Context) error {
		return nil
	})
	if !errors.Is(err, core.ErrPoolStopped) {
		t.Fatalf("expected ErrPoolStopped, got %v", err)
	}
}
