package tools

import (
	"testing"
	"time"
)

func TestToolsPillarProfilerAndRunner(t *testing.T) {
	pillar := NewToolsPillar()
	pillar.Initialize()

	profiler := pillar.PerformanceProfiler()
	profiler.Record("LayoutComputation", 12*time.Millisecond)

	if profiler.SampleCount() < 2 {
		t.Errorf("expected at least 2 profiler samples, got %d", profiler.SampleCount())
	}

	runner := pillar.IntegratedRunner()
	runner.AddResult("TestA", true, "")
	runner.AddResult("TestB", false, "mock failure")

	if runner.PassCount() < 2 {
		t.Errorf("expected at least 2 passed tests, got %d", runner.PassCount())
	}
}
