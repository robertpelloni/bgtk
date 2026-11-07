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

func TestPerformanceProfilerSpansAndReport(t *testing.T) {
	p := NewPerformanceProfiler()
	span := p.StartSpan("render_pass", "Visual")
	if span == nil || span.Completed {
		t.Fatalf("expected active in-flight span")
	}

	time.Sleep(2 * time.Millisecond)
	duration := p.EndSpan("render_pass")
	if duration <= 0 || !span.Completed {
		t.Errorf("expected completed span with positive duration")
	}

	spans := p.GetSpans()
	if len(spans) != 1 {
		t.Errorf("expected 1 registered span, got %d", len(spans))
	}

	report := p.FormatTraceReport()
	if report == "" {
		t.Errorf("expected non-empty trace report")
	}
}

func TestPerformanceProfilerAverageDurationAndFilter(t *testing.T) {
	p := NewPerformanceProfiler()

	p.StartSpan("task1", "IO")
	time.Sleep(2 * time.Millisecond)
	p.EndSpan("task1")

	p.StartSpan("task2", "IO")
	time.Sleep(4 * time.Millisecond)
	p.EndSpan("task2")

	p.StartSpan("task3", "Compute")
	time.Sleep(1 * time.Millisecond)
	p.EndSpan("task3")

	ioSpans := p.GetSpansByCategory("IO")
	if len(ioSpans) != 2 {
		t.Errorf("expected 2 IO spans, got %d", len(ioSpans))
	}

	computeSpans := p.GetSpansByCategory("Compute")
	if len(computeSpans) != 1 {
		t.Errorf("expected 1 Compute span, got %d", len(computeSpans))
	}

	avgIO := p.AverageDuration("IO")
	if avgIO <= 0 {
		t.Errorf("expected positive average duration for IO, got %v", avgIO)
	}

	p.Clear()
	if len(p.GetSpans()) != 0 || p.SampleCount() != 0 {
		t.Errorf("expected profiler cleared")
	}
}
