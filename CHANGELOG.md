# CHANGELOG

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
