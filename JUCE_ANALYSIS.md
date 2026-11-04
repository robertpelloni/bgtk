# JUCE Integration Analysis for Bobtk

JUCE is the premier C++ framework for audio applications, synthesized instruments, DSP signal chains, and high-performance native user interfaces. This analysis outlines the architecture of JUCE's core GUI and event subsystems located in `submodules/juce/modules/` and establishes their integration path within Bobtk's 6-pillar framework (C++ and Go).

---

## 1. GUI Component Hierarchy (`juce_gui_basics/components/juce_Component.h`)

### Architectural Characteristics
- **Unified Base Class (`juce::Component`)**: Every UI element (windows, buttons, sliders, custom vector canvases) inherits from `Component`.
- **Tree Hierarchy**: Components maintain parent-child relationships where child coordinates are strictly relative to their parent (`getX()`, `getY()`, `getWidth()`, `getHeight()`).
- **Paint & Resized Life Cycle**:
  - `paint(Graphics& g)`: Pure rendering callback invoked when a component is marked dirty via `repaint()`.
  - `resized()`: Invoked whenever the component's bounds change via `setBounds()` or layout recomputation.
- **Child Management**: `addAndMakeVisible(Component* child)` registers a child, sets its parent pointer, and ensures it participates in layout traversal.

### Integration with Bobtk Visual Pillar
- **Go Visual Pillar**: Implement `Component` with `AddAndMakeVisible`, `SetBounds`, `Resized`, and `Paint` hooks, allowing audio plugins, visual equalizers, and controls to nest naturally alongside Qt6-parity layouts (`BoxLayout`, `GridLayout`).
- **C++ Visual Pillar**: Map `bobtk::visual::Component` to coordinate transforms and integrate with `bobgui/cpp/widget.hpp`.

---

## 2. Event Dispatching & Timers (`juce_events/timers/juce_Timer.h`)

### Architectural Characteristics
- **Lightweight Repeating Callbacks**: `juce::Timer` enables any class to receive periodic interval callbacks via `timerCallback()`.
- **Decoupled Main Loop**: Allows audio telemetry, animation ticks, and metering meters to update smoothly at 30/60 Hz without allocating heavy OS threads per timer.

### Integration with Bobtk Core Pillar
- **Go Core Pillar**: Implement `Timer` with `StartTimer(intervalMs)`, `StopTimer()`, and concurrent tick dispatching, preventing UI blocking while syncing state changes into `ReflectContext`.

---

## 3. Audio & DSP Pipelines (`juce_audio_processors/`, `juce_dsp/`)

### Architectural Characteristics
- **AudioProcessor**: Standardized DSP unit with `processBlock(AudioBuffer<float>& buffer, MidiBuffer& midiMessages)`.
- **AudioProcessorGraph**: Modular node-based routing for synthesizers, effects plugins, and audio routing.

### Integration with Bobtk Media Pillar
- **Go Media Pillar**: Bridges `AudioBuffer` and `GainProcessor` into an extensible audio processor chain, supporting both audio DSP and MIDI 2.0 Universal MIDI Packets (UMP).

---

## Summary for Bobtk 6-Pillar Architecture

| JUCE Subsystem | Bobtk Pillar | Architectural Role in Bobtk |
|---|---|---|
| `juce_gui_basics/components/Component` | **Visual Pillar** | Hierarchical UI nesting, relative bounds layout, custom paint traversal. |
| `juce_events/timers/Timer` | **Core Pillar** | High-precision repeating callbacks for animation, audio metering, and state sync. |
| `juce_audio_processors/AudioProcessor` | **Media Pillar** | Real-time block-based audio DSP pipelines and synthesizer chains. |
| `juce_gui_extra/misc/SystemTrayIcon` | **System Pillar** | OS notification area docking and server lifecycle supervision. |
