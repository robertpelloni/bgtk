# CHANGELOG

## [1.3.1] - 2026-04-06
### Documentation
- **Agent-Specific Instruction Harmonization**:
  - Synchronized `CLAUDE.md`, `GEMINI.md`, `GPT.md`, and `copilot-instructions.md` with current 6-pillar framework directives, single-page dashboard standards, system tray & server supervisor controls, and pre-flight diagnostics tooling.

## [1.3.0] - 2026-04-06
### Added
- **Clipboard Subsystem (Go System Pillar)**:
  - `Clipboard`: Thread-safe cross-platform clipboard abstraction (`QClipboard` and `GtkClipboard` parity) with `GetText`, `SetText`, `Clear`, and `AddListener`.
  - Exposed via `SystemPillar.Clipboard()` and integrated into `go/main.go`.
  - Comprehensive unit test `TestClipboardOperations` in `go/bobtk/system/tray_test.go`.

## [1.2.9] - 2026-04-06
### Fixed
- **Meson Build System & Sysprof Subproject Dependency**:
  - Restored `gtk=false` in `meson.build` default_options for `libsysprof_capture_dep` (repairing blind mass-renaming corruption).
  - Added missing `sysprof` feature option definition in `meson_options.txt`.

## [1.2.8] - 2026-04-06
### Added
- **MIDI 2.0 Packet Generation (Go Media Pillar)**:
  - `NewNoteOff`: Constructs MIDI 2.0 Note Off Universal MIDI Packets (UMP) with 16-bit release velocity.
  - `NewControlChange`: Constructs MIDI 2.0 Control Change UMP packets with 32-bit high-resolution parameter values.
  - `UMPPacket.Channel`: Extracts 4-bit channel index (0-15).
  - Integrated into `go/main.go` and verified in `go/bobtk/media/media_pillar_test.go`.

## [1.2.7] - 2026-04-06
### Added
- **Circular Buffer Delay DSP Processor (Go Media Pillar)**:
  - `DelayProcessor`: Multi-channel circular buffer delay line with parameterized feedback and wet/dry mix controls (`juce::dsp::DelayLine` parity).
  - Integrated into `AudioGraph` processing chain in `go/main.go` and verified in `go/bobtk/media/media_pillar_test.go`.

## [1.2.6] - 2026-04-06
### Added
- **TabWidget Navigation Container (Go Visual Pillar)**:
  - `TabWidget`: Multi-page tabbed container supporting dynamic tab addition (`AddTab`), tab index switching (`SetCurrentIndex`), and `OnTabChanged` notification (`QTabWidget` parity).
  - `TabItem`: Titled container binding page titles to active child widgets.
  - Integrated into `go/main.go` and verified in `go/bobtk/visual/layout_test.go`.

## [1.2.5] - 2026-04-06
### Added
- **Project Health & Environment Diagnostic Suite**:
  - `tools/check_health.sh`: Inspects availability and versions of Go toolchain, GCC, G++, Meson, and Ninja; evaluates submodule states; and runs Go test suites non-destructively.
  - Enhanced `tools/run_dashboard.sh` with `--check-health` (`-c`) command-line flag to perform pre-flight diagnostics before booting the operations dashboard.

## [1.2.4] - 2026-04-06
### Added
- **Server Start & Restart Operations (Go System Pillar)**:
  - `ServerManager.StartServer`: Starts an inactive or failed background daemon.
  - `ServerManager.RestartServer`: Gracefully stops and reboots an active or stopped daemon.
  - `ServerManager.GetServer`: Queries server status and configuration by ID.
  - `RegisterServerWithStart`: Allows servers to supply both initialization and shutdown lifecycle hooks.
  - Integrated into `go/main.go` via System Tray actions and verified in `go/bobtk/system/tray_test.go`.

## [1.2.3] - 2026-04-06
### Added
- **ParallelFor Concurrency Dispatcher (Go Core Pillar)**:
  - `QuantumScheduler.ParallelFor`: Concurrent loop iteration distribution across worker routines with completion barriers (`upp::CoWork::Do` / `QtConcurrent::map` parity).
  - Integrated into `go/main.go` and verified with unit test suite `TestQuantumSchedulerParallelFor` in `go/bobtk/core/core_pillar_test.go`.

## [1.2.2] - 2026-04-06
### Added
- **ComboBox & RadioButton Widget Primitives (Go Visual Pillar)**:
  - `ComboBox`: Drop-down selector supporting item registration, index tracking, text retrieval, and `OnIndexChanged` notifications (`QComboBox` parity).
  - `RadioButton`: Mutually exclusive radio option supporting group identifiers, selection states, and `OnSelected` notifications (`QRadioButton` parity).
  - Integrated into `go/main.go` and verified in `go/bobtk/visual/layout_test.go`.

## [1.2.1] - 2026-04-06
### Added
- **WebSocket Session Management & Broadcasting (Go Network Pillar)**:
  - `WebSocketSession`: Connected client session structure with per-session non-blocking outbound send queues.
  - `WebSocketHub` extensions: `RegisterClient`, `UnregisterClient`, `Broadcast`, and `ClientCount` for real-time cluster state and UI event distribution.
  - Integrated into `go/main.go` and verified in `go/bobtk/network/network_pillar_test.go`.

## [1.2.0] - 2026-04-06
### Added
- **Sine Oscillator & Low-Pass Filter DSP Processors (Go Media Pillar)**:
  - `SineOscillatorProcessor`: Real-time sine wave synthesizer node with frequency, amplitude, and sample rate parameterization (`juce::ToneGeneratorAudioProcessor` parity).
  - `LowPassFilterProcessor`: Single-pole IIR low-pass filtering node with cutoff frequency smoothing (`juce::IIRFilter` parity).
  - Extended DSP pipeline in `go/bobtk/media/audio_graph.go` and verified in `go/main.go` and `go/bobtk/media/media_pillar_test.go`.

## [1.1.9] - 2026-04-06
### Added
- **Slider & CheckBox Widget Primitives (Go Visual Pillar)**:
  - `Slider`: Stepped/continuous value selector with orientation support (`Horizontal`, `Vertical`), range limits (`Minimum`, `Maximum`), boundary clamping, and `OnValueChanged` notification (QSlider parity).
  - `CheckBox`: Toggleable option control with state tracking (`Checked`), interactive toggling (`Toggle`), and `OnToggled` notification (QCheckBox parity).
  - Integrated `Slider` and `CheckBox` directly into `go/main.go` and verified with unit test suites in `go/bobtk/visual/layout_test.go`.

## [1.1.8] - 2026-04-06
### Added
- **Subsystem Execution Span Profiling & Trace Reporting (Go Tools Pillar)**:
  - `Span`: Structure capturing span identifier, category, start timestamp, elapsed duration, and completion status.
  - `PerformanceProfiler` extensions: `StartSpan`, `EndSpan`, `GetSpans`, and `FormatTraceReport` for detailed subsystem latency analysis.
  - Integrated startup tracing into `ToolsPillar.Initialize` and `main.go`.
  - Comprehensive unit test `TestPerformanceProfilerSpansAndReport` in `go/bobtk/tools/tools_pillar_test.go`.

## [1.1.7] - 2026-04-06
### Added
- **Real-Time Telemetry & Log Streaming (Go Visual Pillar)**:
  - `ServerTelemetry`: Live CPU usage %, memory consumption MB, uptime seconds, and request counters attached to managed server processes.
  - `LogSubscriber`: Dynamic event streaming callback hook allowing external consumers to tap into dashboard log lines in real-time.
  - Formatted live telemetry strings directly within the single-page consolidated dashboard layout.
- **Interactive Global Keyboard Shortcut Dispatcher (Go Visual Pillar)**:
  - `ShortcutEntry` & `ShortcutHandler`: Accelerator registry binding keyboard sequences (e.g. `Ctrl+Q`, `Ctrl+R`, `Ctrl+S`) to executable actions with categorical and descriptive metadata.
  - `HandleKeyPress`: Key dispatcher triggering matching registered actions with automatic execution logging.
  - Added dedicated global keyboard shortcuts summary section to the single-page dashboard render.
- **Verification Suites & Runner Expansion**:
  - Added `TestDashboardTelemetryAndLogStreaming` and `TestDashboardShortcutKeyDispatcher` to `go/bobtk/visual/dashboard_test.go`.
  - Updated `go/main.go` and `tools/run_dashboard.sh` to demonstrate live telemetry updates and keyboard accelerator execution.

## [1.1.6] - 2026-04-06
### Added
- **HTTP Client & Multiplayer Sync Engine (Go Network Pillar)**:
  - `NetworkAccessManager`: HTTP request execution (`bobgui/cpp/network.hpp` and `QNetworkAccessManager` parity) supporting GET/POST methods and history tracking.
  - `SyncEngine`: Multiplayer collaborative state syncing (`bobgui/cpp/network.hpp` parity) with peer connection management, cursor tracking, and state delta broadcasting.
  - Comprehensive unit test suite in `go/bobtk/network/http_sync_test.go`.

## [1.1.5] - 2026-04-06
### Added
- **PropertyAnimation & StateMachine Subsystems (Go Core Pillar)**:
  - `PropertyAnimation`: Interpolates float values over time (`bobgui/cpp/animation.hpp` and `QPropertyAnimation` parity) supporting discrete step advancement and live updates.
  - `StateMachine` & `StateNode`: Hierarchical state transitions with event posting and enter/exit listeners (`bobguistate.h` and `QStateMachine` parity).
  - Comprehensive unit test suites in `go/bobtk/core/animation_test.go` and `go/bobtk/core/state_machine_test.go`.

## [1.1.4] - 2026-04-06
### Added
- **Canvas & 2D Vector Graphics Engine (Go Visual Pillar)**:
  - `GraphicsContext`: 2D canvas drawing pipeline (`QPainter` and `juce::Graphics` parity) supporting `DrawRect`, `FillRect`, `DrawText`, and `StrokePath`.
  - `Color`: Complete color model with predefined constants (`ColorTransparent`, `ColorWhite`, `ColorBlack`, `ColorRed`, `ColorGreen`, `ColorBlue`).
  - `Path`: Vector path builder with `StartNewSubPath`, `LineTo`, and `CloseSubPath`.
  - Integrated `GraphicsContext` directly into the `Component` `Paint` callback hierarchy.
  - Comprehensive unit test suite in `go/bobtk/visual/graphics_test.go`.

## [1.1.3] - 2026-04-06
### Added
- **FileSystem & Settings Subsystems (Go System & Core Pillars)**:
  - `FileSystemWatcher`: Monitors directories and files for changes (`QFileSystemWatcher` parity) with path registration and mutation dispatch.
  - `Settings`: Thread-safe persistent application configuration store (`QSettings` parity) with key-value access and change notifications.
  - Comprehensive unit test suites in `go/bobtk/system/fs_watcher_test.go` and `go/bobtk/core/settings_test.go`.

## [1.1.2] - 2026-04-06
### Added
- **Qt6 Signal/Slot & QAction Subsystem (Go Core & Visual Pillars)**:
  - `Signal`: Thread-safe signal emission, slot connection, and disconnection inspired by Qt6 Signals and Slots.
  - `Action`: Full user action abstraction (`QAction` parity) with text, icon, shortcut, tooltip, and reactive `Triggered`/`Toggled` signals.
  - `KeyEvent`: Key press/release event model (`QKeyEvent` parity) with keyboard modifiers (`Shift`, `Control`, `Alt`, `Meta`) and `Accept`/`Ignore` dispatching.
  - Comprehensive unit test suites in `go/bobtk/core/signal_test.go` and `go/bobtk/visual/action_test.go`.

## [1.1.1] - 2026-04-06
### Added
- **Ultimate++ Map & Serialization Integration (Go Core Pillar)**:
  - `Map`: Ordered key-value map container with index-based fast lookup inspired by U++ `upp::Map` and `Index`.
  - `StateSerializer`: JSON state persistence and configuration serialization inspired by U++ `Jsonize`.
  - Comprehensive unit test coverage in `go/bobtk/core/map_test.go`.

## [1.1.0] - 2026-04-06
### Added
- **Ultimate++ NTL Submodule Assimilation (Go Core Pillar)**:
  - `Vector`: High-performance cache-coherent dynamic array with pointer stability, fast indexed lookup, and pre-allocation inspired by U++ `upp::Vector`.
  - `HeapPool`: Thread-safe memory slab/arena pool avoiding GC allocation spikes during high-frequency UI rendering passes inspired by U++ `sheap`/`lheap`.
  - Comprehensive unit test coverage in `go/bobtk/core/upp_test.go`.

## [1.0.9] - 2026-04-06
### Added
- **Native C/C++ Runtime Bridge (Go Core Pillar)**:
  - `NativeBridge`: Manages GObject / C++ Bobtk runtime context, native handle allocation and deallocation (`AllocateNativeHandle`, `FreeNativeHandle`), and bidirectional event routing between C and Go.
  - Safe pure-Go simulation fallback alongside live CGO linking capabilities.
  - Comprehensive unit test coverage in `go/bobtk/core/bridge_test.go`.

## [1.0.8] - 2026-04-06
### Added
- **Tools Pillar Implementation (Go)**:
  - `PerformanceProfiler`: High-precision execution sampler tracking subsystem latencies and durations (`Sysprof` and `QElapsedTimer` parity).
  - `IntegratedTestRunner`: Built-in automated diagnostic runner and reporting harness.
  - Complete 6-pillar framework test suite coverage in Go (`core`, `media`, `network`, `system`, `tools`, `visual`) with 100% pass rate.

## [1.0.7] - 2026-04-06
### Added
- **JUCE Integration Analysis (`JUCE_ANALYSIS.md`)**:
  - Detailed architectural breakdown of `juce::Component` hierarchy, `juce::Timer` periodic intervals, and `juce_audio_processors` DSP pipelines mapped into Bobtk's 6 pillars.
- **JUCE-Style Component Hierarchy (Go Visual Pillar)**:
  - `Component`: Recursive UI element hierarchy with `AddAndMakeVisible`, parent-relative coordinate transforms, `SetBounds`, `Resized`, and `Paint`/`Repaint` cascades.
- **JUCE-Style Precision Timer (Go Core Pillar)**:
  - `Timer`: Periodic callback dispatcher (`StartTimer`, `StopTimer`, `timerCallback`) decoupled from heavy thread allocations.
  - Comprehensive unit test coverage in `go/bobtk/visual/component_test.go` and `go/bobtk/core/timer_test.go`.

## [1.0.6] - 2026-04-06
### Added
- **Network Pillar Implementation (Go)**:
  - `BroadwayGateway`: Remote HTML5 canvas streaming gateway and client session manager porting GDK Broadway functionality to Go.
  - `WebSocketHub`: Bidirectional JSON-RPC transport and dispatcher for distributed UI state syncing.
  - Comprehensive unit test suite in `go/bobtk/network/network_pillar_test.go`.
- **Media Pillar Implementation (Go)**:
  - `AudioGraph` & `AudioBuffer`: Low-latency DSP audio processing graph inspired by JUCE `AudioProcessor`.
  - `GainProcessor`: Standard real-time DSP gain node.
  - `MidiBus` & `UMPPacket`: MIDI 2.0 Universal MIDI Packet event parser and distribution bus.
  - Comprehensive unit test suite in `go/bobtk/media/media_pillar_test.go`.

## [1.0.5] - 2026-04-06
### Added
- **Qt6-Parity Layout & Widget System (Go Visual Pillar)**:
  - `BoxLayout`: Horizontal (`QHBoxLayout`) and vertical (`QVBoxLayout`) box layout managers supporting dynamic stretch factors, spacing, and margins.
  - `GridLayout`: 2D matrix layout manager (`QGridLayout`) with multi-row and multi-column spanning and proportional distribution.
  - `StackedLayout`: Multi-view card switcher (`QStackedLayout`) with zero-overhead active widget rendering.
  - Core Widget Primitives: `Button` (QPushButton), `Label` (QLabel), `LineEdit` (QLineEdit), `ProgressBar` (QProgressBar) with reactive event signals.
  - Comprehensive unit test suite in `go/bobtk/visual/layout_test.go`.

## [1.0.4] - 2026-04-06
### Added
- **Core Pillar Implementation (Go)**:
  - `QuantumScheduler`: High-concurrency task execution pool modeled after Ultimate++ `CoWork`.
  - `ReflectContext` & `Value`: Dynamic reflection, property binding, and change listeners modeled after U++ `Value` and Qt property systems.
  - `EntityManager`: Cache-coherent entity component storage (ECS) hierarchy in Go.
  - Comprehensive unit test suite in `go/bobtk/core/core_pillar_test.go`.

## [1.0.3] - 2026-04-06
### Added
- **System Tray Subsystem**: Cross-platform system tray notification area abstraction (`SystemTrayIcon` in C++, `SystemTray` in Go) with dynamic context menu actions and tooltip guidance.
- **Server Lifecycle Controller**: Centralized process and daemon supervision (`ServerProcessController` in C++, `ServerManager` in Go) capable of tracking active services (Broadway HTML5 gateway, local IPC daemon, realtime audio server) and providing graceful termination and forceful shutdown.
- **Consolidated Single-Page Dashboard UI**: Redesigned dashboard layout condensing all subsystem views, status metrics, and parallel concepts across the 6 pillars into a singular high-value page with zero fragmented subpages.
- **Rich Guidance & Tooltip Markers**: Integrated detailed tooltips with categories, descriptions, and keyboard shortcuts across all interactive cards and actions.
- **Unified Runner Script**: Created `tools/run_dashboard.sh` for non-destructive test execution, dashboard launching, and server monitoring.

## [1.0.2] - 2024-05-19
### Added
- Completed `UPP_ANALYSIS.md` detailing integration points for Ultimate++ `Core` utilities (Memory Allocation, CoWork concurrency, Serialization, and Callbacks) into the Bobtk 6-pillar framework.

## [1.0.1] - 2024-05-19
### Added
- Initialized the Go port architecture in the `/go` directory.
- Created the six-pillar module structure (`core`, `media`, `network`, `system`, `tools`, `visual`) within Go to prepare for CGO bindings.
- Added basic unit tests for Go module initialization.

### Fixed
- Addressed C build errors caused by renaming (`XKB_CONSUMED_MODE_BOBGUI` back to `XKB_CONSUMED_MODE_GTK` in `gdk/wayland/gdkseat-wayland.c`).
- Partially resolved `meson.build` syntax errors caused by "gtk" mass-renaming corruption.

## [1.0.0] - 2024-04-10
### Added
- Submodule: Ultimate++ (https://github.com/ultimatepp/ultimatepp)
- Submodule: JUCE (https://github.com/juce-framework/JUCE)
- Comprehensive project documentation (`VISION.md`, `ROADMAP.md`, `TODO.md`, `MEMORY.md`, `DEPLOY.md`, `CHANGELOG.md`, `IDEAS.md`, `HANDOFF.md`, etc.).
- Agent instruction files for various LLMs.

### Changed
- Mass renaming: All instances of `gtk` and `bobgui` renamed to `bobtk` to signify the new direction and future of the project.

## [4.0.6] - Unreleased
### Added
- Implemented high-level C++ wrappers for Core subsystems (`bobtk::core_subsystems::AutonomousController`, `BrainNode`, `ComputePipeline`, `SqlDatabase`, `EntityManager`, `FinChart`, `Translator`, `ModernLayout`, `MetaOrchestrator`, `QuantumScheduler`, `RealtimeKernel`, `ReflectContext`, `ScriptContext`, `StateMachine`, `Stream`).
- Implemented the `bobtk::module::CorePillar` interface to manage the creation and state of all data, rendering, layout, machine learning, logic, and state management subsystems within the core 6-pillar framework.
- Exported new C++ core classes through the master `bobgui.hpp` header.

## [4.0.7] - Unreleased
### Added
- Implemented high-level C++ wrappers for Tools subsystems (`bobtk::tools_subsystems::ForgeContext`, `ReportDesigner`, `ReportPreview`, `StudioManager`, `TestRunner`).
- Implemented the `bobtk::module::ToolsPillar` interface to manage the creation and state of IDE integrations, compilers, unit testing, and pdf reporting.
- **MILESTONE**: Completed the entire 6-pillar framework (Media, Network, System, Visual, Core, Tools) in C++, natively wrapping the underlying GObject structures and aggregating them into the master `bobgui.hpp` header.

## [4.0.8] - Unreleased
### Added
- Expanded Go architectural port to encompass all 6 pillars (`tools`, `network`, `system`, `visual`, `media`, `core`) achieving structural API parity with the C++ wrapper framework.
- Evaluated and prepared bridging components for JUCE (`bobgui/cpp/juce_audio_bridge.hpp`) allowing optional compile-time linking to JUCE's Audio Device and DSP systems.
- Evaluated and prepared bridging components for Ultimate++ (`bobgui/cpp/upp_core_bridge.hpp`) allowing optional compile-time data bindings with the Ultimate++ Core framework.

## [4.0.10] - Unreleased
### Added
- Expanded concrete `media.hpp` C++ wrapper capabilities including `AudioDeviceManager` controls (`initializeDevices`, `setSampleRate`, `play`, `stop`), `AudioProcessor` routing (`setParameter`, `processBlock`), `ThreeDNode` properties (`setTranslation`, `setRotation`, `setScale`, `loadMesh`), `MapView` functionality (`setCenter`, `setZoom`, `addMarker`), `EffectShader` uniforms (`setUniformFloat`, `setUniformVec3`), and `Timeline` controls (`play`, `pause`, `seek`).
- Instantiated backend raw C `GObject` stubs for all newly implemented media prototype definitions across `bobguiaudio.c`, `bobgui3d.c`, `bobguigis.c`, `bobguishader.c`, and `bobguitimeline.c` ensuring proper native compilation.
