# MEMORY

- **Codebase Origins:** Started as Bobgui/GTK codebase, heavily reliant on standard C-based GObject architecture.
- **Target Architecture:** Modern, multi-paradigm C++ combined with a parallel, comprehensive Go port.
- **Design Preferences:** Clean, well-structured, easy to use, well-named, and highly documented.
- **Current Observation:** The project has undergone a massive renaming from `gtk` and `bobgui` to `bobtk`. Pathing issues or build files (e.g., meson.build, Makefiles) may require further refinement if `bobtk` replacements broke regex or internal tooling.
- **System Tray and Server Process Supervisor:** Added cross-platform system tray integration (`SystemTrayIcon` in C++, `SystemTray` in Go) with dynamic context menus and background daemon lifecycle tracking (`ServerProcessController` / `ServerManager`).
- **Consolidated Dashboard Paradigm:** All features across the 6 pillars are centralized into a single-page consolidated view without fragmented subpages, featuring prominent cards, tooltips, and server shutdown controls.
- **Go Core Pillar Implementation:** `QuantumScheduler` handles concurrent pooled background execution (`CoWork` pattern), `ReflectContext` provides dynamic `Value` binding and observable properties, and `EntityManager` manages ECS component relationships.
- **Go Visual Pillar Qt6 Parity:** Implemented `BoxLayout` (`QHBoxLayout`, `QVBoxLayout`), `GridLayout` (`QGridLayout`), `StackedLayout` (`QStackedLayout`), and reactive widget controls (`Button`, `Label`, `LineEdit`, `ProgressBar`).
- **Go Network & Media Pillars:** Network pillar hosts `BroadwayGateway` (remote HTML5 canvas streaming) and `WebSocketHub` (JSON-RPC transport). Media pillar hosts `AudioGraph` (JUCE-like DSP audio processors) and `MidiBus` (MIDI 2.0 Universal MIDI Packets).
