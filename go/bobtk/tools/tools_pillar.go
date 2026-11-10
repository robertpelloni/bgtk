package tools

import (
	"fmt"
	"sync"
	"time"
)

// ProfilerSample captures CPU/latency metrics (Sysprof parity).
type ProfilerSample struct {
	Timestamp time.Time
	Label     string
	Duration  time.Duration
}

// Span represents an active or completed execution trace interval.
type Span struct {
	ID        string
	Category  string
	StartTime time.Time
	Duration  time.Duration
	Completed bool
}

// PerformanceProfiler measures subsystem timing and performance (QElapsedTimer / Sysprof parity).
type PerformanceProfiler struct {
	mu      sync.RWMutex
	samples []ProfilerSample
	spans   map[string]*Span
}

func NewPerformanceProfiler() *PerformanceProfiler {
	return &PerformanceProfiler{
		samples: make([]ProfilerSample, 0),
		spans:   make(map[string]*Span),
	}
}

func (p *PerformanceProfiler) Record(label string, duration time.Duration) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.samples = append(p.samples, ProfilerSample{
		Timestamp: time.Now(),
		Label:     label,
		Duration:  duration,
	})
}

func (p *PerformanceProfiler) SampleCount() int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return len(p.samples)
}

// StartSpan opens an execution trace span.
func (p *PerformanceProfiler) StartSpan(id, category string) *Span {
	p.mu.Lock()
	defer p.mu.Unlock()
	s := &Span{
		ID:        id,
		Category:  category,
		StartTime: time.Now(),
	}
	p.spans[id] = s
	return s
}

// EndSpan closes an active span and calculates elapsed execution duration.
func (p *PerformanceProfiler) EndSpan(id string) time.Duration {
	p.mu.Lock()
	defer p.mu.Unlock()
	s, exists := p.spans[id]
	if !exists {
		return 0
	}
	s.Duration = time.Since(s.StartTime)
	s.Completed = true
	p.samples = append(p.samples, ProfilerSample{
		Timestamp: s.StartTime,
		Label:     fmt.Sprintf("%s:%s", s.Category, s.ID),
		Duration:  s.Duration,
	})
	return s.Duration
}

// GetSpans returns all registered trace spans.
func (p *PerformanceProfiler) GetSpans() []*Span {
	p.mu.RLock()
	defer p.mu.RUnlock()
	out := make([]*Span, 0, len(p.spans))
	for _, s := range p.spans {
		out = append(out, s)
	}
	return out
}

// GetSpansByCategory returns only spans matching the specified category tag.
func (p *PerformanceProfiler) GetSpansByCategory(category string) []*Span {
	p.mu.RLock()
	defer p.mu.RUnlock()
	out := make([]*Span, 0)
	for _, s := range p.spans {
		if s.Category == category {
			out = append(out, s)
		}
	}
	return out
}

// AverageDuration computes the average duration of completed spans in a category.
func (p *PerformanceProfiler) AverageDuration(category string) time.Duration {
	p.mu.RLock()
	defer p.mu.RUnlock()
	var total time.Duration
	count := 0
	for _, s := range p.spans {
		if s.Category == category && s.Completed {
			total += s.Duration
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return total / time.Duration(count)
}

// Clear flushes all samples and spans.
func (p *PerformanceProfiler) Clear() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.samples = make([]ProfilerSample, 0)
	p.spans = make(map[string]*Span)
}

// FormatTraceReport summarizes spans into a human-readable diagnostic breakdown.
func (p *PerformanceProfiler) FormatTraceReport() string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	out := fmt.Sprintf("Performance Tracing: %d spans recorded\n", len(p.spans))
	for _, s := range p.spans {
		status := "IN_FLIGHT"
		if s.Completed {
			status = fmt.Sprintf("%0.2fms", float64(s.Duration.Microseconds())/1000.0)
		}
		out += fmt.Sprintf("  • [%s] %s -> %s\n", s.Category, s.ID, status)
	}
	return out
}

// TestResult captures test status and assertions.
type TestResult struct {
	Name   string
	Passed bool
	Error  string
}

// BenchmarkResult records throughput metrics of an automated benchmark run.
type BenchmarkResult struct {
	Name         string
	Iterations   int
	TotalElapsed time.Duration
	AverageTime  time.Duration
	OpsPerSec    float64
}

// BenchmarkRunner provides automated execution throughput analysis.
type BenchmarkRunner struct {
	mu      sync.RWMutex
	results []BenchmarkResult
}

func NewBenchmarkRunner() *BenchmarkRunner {
	return &BenchmarkRunner{
		results: make([]BenchmarkResult, 0),
	}
}

func (b *BenchmarkRunner) Run(name string, iterations int, fn func()) BenchmarkResult {
	if iterations <= 0 {
		iterations = 1000
	}

	start := time.Now()
	for i := 0; i < iterations; i++ {
		fn()
	}
	elapsed := time.Since(start)

	avg := elapsed / time.Duration(iterations)
	var opsPerSec float64
	if elapsed.Seconds() > 0 {
		opsPerSec = float64(iterations) / elapsed.Seconds()
	}

	res := BenchmarkResult{
		Name:         name,
		Iterations:   iterations,
		TotalElapsed: elapsed,
		AverageTime:  avg,
		OpsPerSec:    opsPerSec,
	}

	b.mu.Lock()
	b.results = append(b.results, res)
	b.mu.Unlock()

	return res
}

func (b *BenchmarkRunner) Results() []BenchmarkResult {
	b.mu.RLock()
	defer b.mu.RUnlock()
	out := make([]BenchmarkResult, len(b.results))
	copy(out, b.results)
	return out
}

func (b *BenchmarkRunner) Count() int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return len(b.results)
}

func (b *BenchmarkRunner) Clear() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.results = make([]BenchmarkResult, 0)
}

// IntegratedTestRunner runs and reports unit and integration checks.
type IntegratedTestRunner struct {
	mu      sync.RWMutex
	results []TestResult
}

func NewIntegratedTestRunner() *IntegratedTestRunner {
	return &IntegratedTestRunner{
		results: make([]TestResult, 0),
	}
}

func (tr *IntegratedTestRunner) AddResult(name string, passed bool, err string) {
	tr.mu.Lock()
	defer tr.mu.Unlock()
	tr.results = append(tr.results, TestResult{Name: name, Passed: passed, Error: err})
}

func (tr *IntegratedTestRunner) PassCount() int {
	tr.mu.RLock()
	defer tr.mu.RUnlock()
	count := 0
	for _, r := range tr.results {
		if r.Passed {
			count++
		}
	}
	return count
}

// ToolsPillar manages IDE integrations, compilers, unit testing, and reporting.
type ToolsPillar struct {
	initialized bool
	Profiler    *PerformanceProfiler
	Runner      *IntegratedTestRunner
	Benchmarker *BenchmarkRunner
}

func NewToolsPillar() *ToolsPillar {
	return &ToolsPillar{
		initialized: false,
		Profiler:    NewPerformanceProfiler(),
		Runner:      NewIntegratedTestRunner(),
		Benchmarker: NewBenchmarkRunner(),
	}
}

func (t *ToolsPillar) Initialize() {
	if t.initialized {
		return
	}
	t.Profiler.StartSpan("init.tools_pillar", "Subsystem")
	t.Profiler.Record("startup", 5*time.Millisecond)
	t.Runner.AddResult("PillarValidation", true, "")
	fmt.Println("ToolsPillar initialized.")
	t.Profiler.EndSpan("init.tools_pillar")
	t.initialized = true
}

func (t *ToolsPillar) PerformanceProfiler() *PerformanceProfiler {
	return t.Profiler
}

func (t *ToolsPillar) IntegratedRunner() *IntegratedTestRunner {
	return t.Runner
}

func (t *ToolsPillar) BenchmarkRunner() *BenchmarkRunner {
	return t.Benchmarker
}
