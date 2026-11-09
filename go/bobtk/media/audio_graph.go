package media

import (
	"math"
	"sync"
)

// AudioBuffer represents a multi-channel PCM audio buffer (JUCE AudioBuffer parity).
type AudioBuffer struct {
	Channels   int
	NumSamples int
	Data       [][]float32
}

// NewAudioBuffer allocates an audio buffer.
func NewAudioBuffer(channels, numSamples int) *AudioBuffer {
	data := make([][]float32, channels)
	for c := 0; c < channels; c++ {
		data[c] = make([]float32, numSamples)
	}
	return &AudioBuffer{
		Channels:   channels,
		NumSamples: numSamples,
		Data:       data,
	}
}

// Clear zeroes all samples in all channels of the buffer (juce::AudioBuffer::clear parity).
func (ab *AudioBuffer) Clear() {
	for c := 0; c < ab.Channels; c++ {
		for i := 0; i < ab.NumSamples; i++ {
			ab.Data[c][i] = 0.0
		}
	}
}

// CopyFrom copies samples from a source audio buffer up to matching channel and sample bounds.
func (ab *AudioBuffer) CopyFrom(src *AudioBuffer) {
	if src == nil {
		return
	}
	chans := ab.Channels
	if src.Channels < chans {
		chans = src.Channels
	}
	samples := ab.NumSamples
	if src.NumSamples < samples {
		samples = src.NumSamples
	}
	for c := 0; c < chans; c++ {
		copy(ab.Data[c][:samples], src.Data[c][:samples])
	}
}

// GetMagnitude computes the peak absolute amplitude across a channel (juce::AudioBuffer::getMagnitude parity).
func (ab *AudioBuffer) GetMagnitude(channel int) float32 {
	if channel < 0 || channel >= ab.Channels {
		return 0.0
	}
	var maxMag float32 = 0.0
	for i := 0; i < ab.NumSamples; i++ {
		v := float32(math.Abs(float64(ab.Data[channel][i])))
		if v > maxMag {
			maxMag = v
		}
	}
	return maxMag
}

// AudioProcessor defines the DSP processing interface inspired by JUCE AudioProcessor.
type AudioProcessor interface {
	ProcessBlock(buffer *AudioBuffer)
	Name() string
}

// GainProcessor scales audio buffer volume by a gain factor.
type GainProcessor struct {
	Gain float32
}

func (g *GainProcessor) Name() string { return "GainProcessor" }
func (g *GainProcessor) ProcessBlock(buffer *AudioBuffer) {
	for c := 0; c < buffer.Channels; c++ {
		for i := 0; i < buffer.NumSamples; i++ {
			buffer.Data[c][i] *= g.Gain
		}
	}
}

// SineOscillatorProcessor synthesizes a pure sine waveform (JUCE ToneGeneratorAudioProcessor parity).
type SineOscillatorProcessor struct {
	Frequency  float64
	SampleRate float64
	Amplitude  float32
	phase      float64
}

func NewSineOscillatorProcessor(freq, sampleRate float64, amp float32) *SineOscillatorProcessor {
	if sampleRate <= 0 {
		sampleRate = 48000
	}
	if amp <= 0 {
		amp = 1.0
	}
	return &SineOscillatorProcessor{
		Frequency:  freq,
		SampleRate: sampleRate,
		Amplitude:  amp,
		phase:      0,
	}
}

func (s *SineOscillatorProcessor) Name() string { return "SineOscillatorProcessor" }
func (s *SineOscillatorProcessor) ProcessBlock(buffer *AudioBuffer) {
	phaseIncrement := (2.0 * math.Pi * s.Frequency) / s.SampleRate
	for i := 0; i < buffer.NumSamples; i++ {
		sample := float32(math.Sin(s.phase)) * s.Amplitude
		s.phase += phaseIncrement
		if s.phase >= 2.0*math.Pi {
			s.phase -= 2.0 * math.Pi
		}
		for c := 0; c < buffer.Channels; c++ {
			buffer.Data[c][i] += sample
		}
	}
}

// LowPassFilterProcessor applies a single-pole IIR low-pass filter (JUCE IIRFilter parity).
type LowPassFilterProcessor struct {
	Cutoff     float64
	SampleRate float64
	prevOutput []float32
}

func NewLowPassFilterProcessor(cutoff, sampleRate float64, channels int) *LowPassFilterProcessor {
	if sampleRate <= 0 {
		sampleRate = 48000
	}
	if channels <= 0 {
		channels = 2
	}
	return &LowPassFilterProcessor{
		Cutoff:     cutoff,
		SampleRate: sampleRate,
		prevOutput: make([]float32, channels),
	}
}

func (lp *LowPassFilterProcessor) Name() string { return "LowPassFilterProcessor" }
func (lp *LowPassFilterProcessor) ProcessBlock(buffer *AudioBuffer) {
	if len(lp.prevOutput) < buffer.Channels {
		lp.prevOutput = make([]float32, buffer.Channels)
	}
	// Calculate smoothing coefficient alpha = dt / (RC + dt)
	dt := 1.0 / lp.SampleRate
	rc := 1.0 / (2.0 * math.Pi * lp.Cutoff)
	alpha := float32(dt / (rc + dt))
	if alpha > 1.0 {
		alpha = 1.0
	}

	for c := 0; c < buffer.Channels; c++ {
		prev := lp.prevOutput[c]
		for i := 0; i < buffer.NumSamples; i++ {
			out := prev + alpha*(buffer.Data[c][i]-prev)
			buffer.Data[c][i] = out
			prev = out
		}
		lp.prevOutput[c] = prev
	}
}

// HighPassFilterProcessor applies a single-pole high-pass filter (JUCE IIRFilter highpass parity).
type HighPassFilterProcessor struct {
	Cutoff     float64
	SampleRate float64
	prevInput  []float32
	prevOutput []float32
}

func NewHighPassFilterProcessor(cutoff, sampleRate float64, channels int) *HighPassFilterProcessor {
	if sampleRate <= 0 {
		sampleRate = 48000
	}
	if channels <= 0 {
		channels = 2
	}
	return &HighPassFilterProcessor{
		Cutoff:     cutoff,
		SampleRate: sampleRate,
		prevInput:  make([]float32, channels),
		prevOutput: make([]float32, channels),
	}
}

func (hp *HighPassFilterProcessor) Name() string { return "HighPassFilterProcessor" }
func (hp *HighPassFilterProcessor) ProcessBlock(buffer *AudioBuffer) {
	if len(hp.prevInput) < buffer.Channels {
		hp.prevInput = make([]float32, buffer.Channels)
		hp.prevOutput = make([]float32, buffer.Channels)
	}
	// RC filter: alpha = rc / (rc + dt)
	dt := 1.0 / hp.SampleRate
	rc := 1.0 / (2.0 * math.Pi * hp.Cutoff)
	alpha := float32(rc / (rc + dt))

	for c := 0; c < buffer.Channels; c++ {
		prevIn := hp.prevInput[c]
		prevOut := hp.prevOutput[c]
		for i := 0; i < buffer.NumSamples; i++ {
			in := buffer.Data[c][i]
			out := alpha * (prevOut + in - prevIn)
			buffer.Data[c][i] = out
			prevIn = in
			prevOut = out
		}
		hp.prevInput[c] = prevIn
		hp.prevOutput[c] = prevOut
	}
}

// PeakingEQProcessor provides parametric peaking EQ filtering (JUCE dsp::IIR::Filter parity).
type PeakingEQProcessor struct {
	CenterFreq float64
	Q          float64
	GainDB     float64
	SampleRate float64
	x1, x2     []float32
	y1, y2     []float32
}

func NewPeakingEQProcessor(centerFreq, q, gainDB, sampleRate float64, channels int) *PeakingEQProcessor {
	if sampleRate <= 0 {
		sampleRate = 48000
	}
	if channels <= 0 {
		channels = 2
	}
	if q <= 0 {
		q = 1.0
	}
	return &PeakingEQProcessor{
		CenterFreq: centerFreq,
		Q:          q,
		GainDB:     gainDB,
		SampleRate: sampleRate,
		x1:         make([]float32, channels),
		x2:         make([]float32, channels),
		y1:         make([]float32, channels),
		y2:         make([]float32, channels),
	}
}

func (eq *PeakingEQProcessor) Name() string { return "PeakingEQProcessor" }
func (eq *PeakingEQProcessor) ProcessBlock(buffer *AudioBuffer) {
	// Standard biquad peaking EQ design (RBJ Audio EQ Cookbook)
	A := math.Pow(10.0, eq.GainDB/40.0)
	omega := 2.0 * math.Pi * eq.CenterFreq / eq.SampleRate
	alpha := math.Sin(omega) / (2.0 * eq.Q)
	cosw := math.Cos(omega)

	b0 := 1.0 + alpha*A
	b1 := -2.0 * cosw
	b2 := 1.0 - alpha*A
	a0 := 1.0 + alpha/A
	a1 := -2.0 * cosw
	a2 := 1.0 - alpha/A

	// Normalize coefficients
	b0_norm := float32(b0 / a0)
	b1_norm := float32(b1 / a0)
	b2_norm := float32(b2 / a0)
	a1_norm := float32(a1 / a0)
	a2_norm := float32(a2 / a0)

	for c := 0; c < buffer.Channels; c++ {
		if c >= len(eq.x1) {
			continue
		}
		x1 := eq.x1[c]
		x2 := eq.x2[c]
		y1 := eq.y1[c]
		y2 := eq.y2[c]

		for i := 0; i < buffer.NumSamples; i++ {
			x0 := buffer.Data[c][i]
			y0 := b0_norm*x0 + b1_norm*x1 + b2_norm*x2 - a1_norm*y1 - a2_norm*y2

			x2 = x1
			x1 = x0
			y2 = y1
			y1 = y0

			buffer.Data[c][i] = y0
		}

		eq.x1[c] = x1
		eq.x2[c] = x2
		eq.y1[c] = y1
		eq.y2[c] = y2
	}
}

// DelayProcessor adds echo and reverberant delay lines (JUCE DelayLine parity).
type DelayProcessor struct {
	DelaySamples int
	Feedback     float32
	Mix          float32
	buffers      [][]float32
	writePos     int
}

func NewDelayProcessor(delaySamples int, feedback, mix float32, channels int) *DelayProcessor {
	if delaySamples <= 0 {
		delaySamples = 4800 // 100ms at 48kHz
	}
	if channels <= 0 {
		channels = 2
	}
	bufs := make([][]float32, channels)
	for c := 0; c < channels; c++ {
		bufs[c] = make([]float32, delaySamples)
	}
	return &DelayProcessor{
		DelaySamples: delaySamples,
		Feedback:     feedback,
		Mix:          mix,
		buffers:      bufs,
		writePos:     0,
	}
}

func (d *DelayProcessor) Name() string { return "DelayProcessor" }
func (d *DelayProcessor) ProcessBlock(buffer *AudioBuffer) {
	for i := 0; i < buffer.NumSamples; i++ {
		for c := 0; c < buffer.Channels; c++ {
			if c >= len(d.buffers) {
				continue
			}
			inSample := buffer.Data[c][i]
			delayedSample := d.buffers[c][d.writePos]

			// Feedback loop
			d.buffers[c][d.writePos] = inSample + (delayedSample * d.Feedback)

			// Wet / Dry blend
			buffer.Data[c][i] = inSample*(1.0-d.Mix) + delayedSample*d.Mix
		}
		d.writePos = (d.writePos + 1) % d.DelaySamples
	}
}

// AudioGraph chains audio processors into an execution pipeline.
type AudioGraph struct {
	mu         sync.RWMutex
	sampleRate int
	bufferSize int
	processors []AudioProcessor
}

// NewAudioGraph instantiates an audio DSP graph.
func NewAudioGraph(sampleRate, bufferSize int) *AudioGraph {
	if sampleRate <= 0 {
		sampleRate = 48000
	}
	if bufferSize <= 0 {
		bufferSize = 256
	}
	return &AudioGraph{
		sampleRate: sampleRate,
		bufferSize: bufferSize,
		processors: make([]AudioProcessor, 0),
	}
}

// AddProcessor adds a DSP node to the chain.
func (ag *AudioGraph) AddProcessor(proc AudioProcessor) {
	ag.mu.Lock()
	defer ag.mu.Unlock()
	ag.processors = append(ag.processors, proc)
}

// Process passes an audio buffer through all chained processors.
func (ag *AudioGraph) Process(buffer *AudioBuffer) {
	ag.mu.RLock()
	defer ag.mu.RUnlock()
	for _, proc := range ag.processors {
		proc.ProcessBlock(buffer)
	}
}

// ProcessorCount returns the number of active DSP nodes.
func (ag *AudioGraph) ProcessorCount() int {
	ag.mu.RLock()
	defer ag.mu.RUnlock()
	return len(ag.processors)
}

// MinMaxSample holds the minimum and maximum sample peak values within a waveform window.
type MinMaxSample struct {
	MinVal float32
	MaxVal float32
}

// AudioThumbnail stores and generates downsampled audio waveform peak overviews (juce::AudioThumbnail parity).
type AudioThumbnail struct {
	mu          sync.RWMutex
	samplesPerThumb int
	peaks       [][]MinMaxSample
}

// NewAudioThumbnail creates an AudioThumbnail cache.
func NewAudioThumbnail(samplesPerThumb int) *AudioThumbnail {
	if samplesPerThumb <= 0 {
		samplesPerThumb = 256
	}
	return &AudioThumbnail{
		samplesPerThumb: samplesPerThumb,
		peaks:           make([][]MinMaxSample, 0),
	}
}

// AddBlock extracts min/max peak levels from an incoming AudioBuffer and stores them.
func (at *AudioThumbnail) AddBlock(buffer *AudioBuffer) {
	if buffer == nil || buffer.Channels == 0 || buffer.NumSamples == 0 {
		return
	}

	at.mu.Lock()
	defer at.mu.Unlock()

	// Ensure channels allocated
	for len(at.peaks) < buffer.Channels {
		at.peaks = append(at.peaks, make([]MinMaxSample, 0))
	}

	for c := 0; c < buffer.Channels; c++ {
		for i := 0; i < buffer.NumSamples; i += at.samplesPerThumb {
			minVal := float32(1.0)
			maxVal := float32(-1.0)

			end := i + at.samplesPerThumb
			if end > buffer.NumSamples {
				end = buffer.NumSamples
			}

			for s := i; s < end; s++ {
				val := buffer.Data[c][s]
				if val < minVal {
					minVal = val
				}
				if val > maxVal {
					maxVal = val
				}
			}

			at.peaks[c] = append(at.peaks[c], MinMaxSample{
				MinVal: minVal,
				MaxVal: maxVal,
			})
		}
	}
}

// NumThumbnails returns the number of downsampled overview windows stored for a channel.
func (at *AudioThumbnail) NumThumbnails(channel int) int {
	at.mu.RLock()
	defer at.mu.RUnlock()
	if channel < 0 || channel >= len(at.peaks) {
		return 0
	}
	return len(at.peaks[channel])
}

// GetMinMax retrieves the min/max peak for a channel at a given overview window index.
func (at *AudioThumbnail) GetMinMax(channel, index int) (MinMaxSample, bool) {
	at.mu.RLock()
	defer at.mu.RUnlock()
	if channel < 0 || channel >= len(at.peaks) {
		return MinMaxSample{}, false
	}
	if index < 0 || index >= len(at.peaks[channel]) {
		return MinMaxSample{}, false
	}
	return at.peaks[channel][index], true
}

// Reset clears all cached peak overviews.
func (at *AudioThumbnail) Reset() {
	at.mu.Lock()
	defer at.mu.Unlock()
	at.peaks = make([][]MinMaxSample, 0)
}
