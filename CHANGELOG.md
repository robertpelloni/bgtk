# CHANGELOG

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
