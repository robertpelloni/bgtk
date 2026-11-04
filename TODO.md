# TODO

- [x] Rename `gtk` to `bobtk`.
- [x] Rename `bobgui` to `bobtk`.
- [x] Add Ultimate++ submodule.
- [x] Add JUCE submodule.
- [x] Initialize primary documentation files.
- [x] Initialize `/go` directory and establish the initial module (6-Pillar framework initialized).
- [x] Analyze Ultimate++ utilities for integration points.
- [x] Implement System Tray integration in C++ and Go with background server lifecycle and quit capability.
- [x] Redesign Dashboard UI to consolidate all subsystem features across the 6 pillars into a single page with tooltip markers.
- [x] Create unified runner script `tools/run_dashboard.sh`.
- [x] Build out Core Pillar subsystems in Go (`QuantumScheduler`, `ReflectContext`, `EntityManager`).
- [x] Implement Network and Media pillar subsystems in Go (`BroadwayGateway`, `WebSocketHub`, `AudioGraph`, `MidiBus`).
- [ ] Analyze JUCE UI widgets for intermixing capabilities.
- [x] Implement Qt6 layout and widget parity in Go Visual Pillar (`BoxLayout`, `GridLayout`, `StackedLayout`, `Button`, `Label`, `LineEdit`, `ProgressBar`).
- [ ] Research missing Qt6/Qt7 features compared to the current Bobtk C++ baseline.
- [ ] Test and resolve any failing builds caused by renaming strings and paths (Meson wrap corruption).
