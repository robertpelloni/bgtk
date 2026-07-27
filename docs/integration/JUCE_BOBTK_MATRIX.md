# JUCE and Bobtk Compatibility Matrix & Merge Strategy

## 1. Architectural Overview & Overlaps
Bobtk utilizes a 6-Pillar framework: Core, System, Network, Visual, Media, and Tools.
JUCE is traditionally monolithic but can be conceptually mapped to these pillars.

| Bobtk Pillar | JUCE Equivalent Modules | Overlapping Features | Gaps & Conflicts |
|---|---|---|---|
| **Core** | `juce_core`, `juce_events` | String handling, basic data structures, threading, timers, event loops. | JUCE's string and memory handling (e.g., `juce::String`, `juce::var`) conflicts with Bobtk's GLib/C++ standard patterns. Go interop will require a unified C-API boundary for Core structures. |
| **System** | `juce_core` (OS layer) | File I/O, process management, dynamic library loading. | File paths and system abstractions differ. JUCE relies heavily on its own `File` class which may duplicate GLib's `GFile`. |
| **Network** | `juce_core` (Network) | Sockets, URL handling. | JUCE network capabilities are minimal compared to dedicated libraries. |
| **Visual** | `juce_graphics`, `juce_gui_basics`, `juce_gui_extra` | Widget hierarchy (`juce::Component`), LookAndFeel, 2D rendering, Window management. | **Major overlap.** Bobtk's widget system (inherited from GTK) is retained-mode with CSS styling. JUCE uses a heavily code-driven `LookAndFeel` and immediate-ish drawing `paint()` methods. Integrating JUCE widgets into Bobtk's visual tree (or vice versa) is a critical blocker. |
| **Media** | `juce_audio_basics`, `juce_audio_devices`, `juce_audio_formats`, `juce_audio_processors` | Audio I/O, MIDI, Plugin hosting (VST/AU), audio graphs. | **JUCE's strongest area.** Bobtk's media pillar currently relies on GStreamer. JUCE's low-latency audio paradigm does not directly map to GStreamer pipelines. |
| **Tools** | `Projucer` | Build system generation. | Bobtk uses Meson. JUCE modules will need to be consumed via CMake/Meson wraps rather than Projucer. |

## 2. Minimal Viable Merge Strategy for `/go`
To bring JUCE's capabilities into the Go port (`/go`), we need a strategy that doesn't require rewriting JUCE in Go, but rather binding to its C++ APIs through a C compatibility layer.

### Strategy Outline:
1. **C-API Wrapper (Cgo interface):** Create a minimal C-API exposing JUCE's fundamental Audio and Windowing classes (e.g., a wrapper around `juce::AudioProcessor` and `juce::Component`).
2. **Go `Media` Pillar Integration:** Bind the C-API to `/go/media`. This will allow Go to spin up JUCE audio devices and process audio callbacks.
3. **Go `Visual` Pillar Integration:** For UI, attempting to merge the Bobtk (GTK-based) visual tree with JUCE's `Component` tree is fraught. The MVP approach is to host JUCE `Component`s inside an opaque Bobtk window handle (e.g., via a native window ID).

### C++/Go Interop Blockers:
- **Event Loop Integration:** JUCE requires its own `MessageManager` (event loop) to run on the main thread for UI tasks. Bobtk/GTK also demands the main thread. Resolving this duality (e.g., running JUCE in a headless mode or integrating the event loops via custom FD polling) is a major blocker.
- **Callback Overhead:** Audio processing in Go via Cgo incurs overhead. High-performance, low-latency audio callbacks from JUCE might suffer if bouncing frequently between C++ and Go.

## 3. Current C++ Build Blockers
- **Meson Wrap Conflicts:** The mass rename from `gtk` to `bobgui`/`bobtk` has broken subproject `.wrap` files and `meson.build` dependency declarations. The C++ build is currently failing with syntax errors and missing components, blocking immediate C++ compilation and testing of any JUCE integration.
