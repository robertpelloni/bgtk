# JUCE vs Bobtk Event System & Architecture Analysis

## 1. Widget Architecture Mapping
JUCE's UI architecture is built around the `juce::Component` class. This is a monolithic base class handling rendering (`paint()`), layout (`resized()`), and all event processing (mouse, keyboard, focus).
Bobtk (inheriting GTK paradigms) currently uses a more retained-mode, separated approach with CSS styling and specialized widget types (e.g., `GtkWidget`).

| Feature | JUCE (`juce::Component`) | Bobtk / Go (`visual.Widget`) | 6-Pillar Mapping |
| :--- | :--- | :--- | :--- |
| **Rendering** | Immediate-mode `paint(Graphics&)` | Retained tree, render nodes | **Visual** |
| **Layout** | Manual math in `resized()` or FlexBox | Layout Managers (`BaseLayout`) | **Visual** |
| **Styling** | Programmatic `LookAndFeel` | CSS / Property Bag (`Style`) | **Visual** |
| **Event Routing** | `MessageManager` (Main Thread) | `GLib` Event Loop / `MotionEmitter` | **System / Visual** |

## 2. Event System Comparison
### JUCE Event System
- **Core Engine:** Driven by the `juce::MessageManager`. It must run on the OS's main UI thread.
- **Delivery:** Events (mouse clicks, key presses, timers) are pulled from the OS queue and dispatched synchronously to `Component` methods (`mouseDown`, `mouseDrag`, etc.).
- **Async Messaging:** `MessageListener` and `AsyncUpdater` provide ways to defer execution to the next message loop iteration.
- **Audio Context:** JUCE strictly separates UI thread events from the high-priority Audio thread. Audio callbacks (`processBlock`) *cannot* allocate memory or lock UI mutexes.

### Bobtk Event System (Current)
- **Core Engine:** Built on GLib's `GMainContext` and `GMainLoop`.
- **Delivery:** Signal-based (e.g., `g_signal_connect`) and event structures (`GdkEvent`). Recently introduced synthetic motion events for continuous scrolling interactions.
- **Go Port (`visual.BaseMotionEmitter`):** Uses thread-safe channels or mutex-protected callback slices to dispatch events from C++ to Go.

### The Conflict
JUCE requires controlling the main thread via its `MessageManager`. GLib/GTK also requires controlling the main thread via `GMainLoop`.
**Resolution Strategy:** One loop must host the other. Typically, JUCE can be run in a mode where it integrates its timers and event dispatching into an external run loop (like GLib's), or vice versa.

## 3. Proposed Unified Interface for `/go/media`
Given JUCE's dominance in the audio domain, the Go `Media` pillar should abstract JUCE's `AudioProcessor` model.

```go
package media

// AudioBuffer represents a block of multi-channel audio data.
type AudioBuffer struct {
	Channels [][]float32
	NumSamples int
}

// MediaProcessor represents a unified interface for audio/MIDI processing,
// bridging Bobtk's generic media pillar with JUCE's low-latency requirements.
type MediaProcessor interface {
	// PrepareToPlay is called before processing starts.
	PrepareToPlay(sampleRate float64, samplesPerBlock int)

	// ProcessBlock processes the audio in a real-time, allocation-free context.
	ProcessBlock(buffer *AudioBuffer)

	// ReleaseResources cleans up after processing stops.
	ReleaseResources()
}

// MediaContext manages the backend engine (e.g., initializing JUCE AudioDeviceManager).
type MediaContext interface {
	Initialize() error
	RegisterProcessor(processor MediaProcessor)
	Start() error
	Stop() error
}
```
This design allows the Go layer to define DSP logic while a C++ bridge handles the actual low-latency JUCE `AudioIODeviceCallback` integration, shuttling data into the `AudioBuffer`.
