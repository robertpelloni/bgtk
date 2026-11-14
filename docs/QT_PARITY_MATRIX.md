# Qt6 Parity Matrix — Bobtk C++ baseline vs Qt6 (2026-11-13)

Method: Bobtk column cites real files in `bobgui/cpp/` (inspected this session). Status: **Covered** = dedicated wrapper exists, **Partial** = underlying GTK/GDK capability without Bobtk wrapper, **Missing** = no counterpart found. This is a research map, not a claim of API equivalence — "Covered" means a starting wrapper, not Qt-compatible semantics.

| Qt6 module | Bobtk status | Evidence / gap |
|---|---|---|
| QtCore (objects, signals, properties) | Covered | `object_handle.hpp`, `property.hpp`, `signal.hpp`, `timer.hpp`, `settings.hpp`, `core_subsystems.hpp` (StateMachine, ScriptContext, ReflectContext) |
| QtGui (painting, images, text) | Partial | `graphics.hpp`, `canvas.hpp`, `theme.hpp`; text/layout depth lives in Pango/GDK without Bobtk wrappers |
| QtWidgets (shells, docks, commands) | Covered | `widget.hpp`, `workbench.hpp`, `app_shell.hpp`, `studio_shell.hpp`, `document_shell.hpp`, `dashboard_shell.hpp`, `dock_manager.hpp`, `action_registry.hpp`, `command_palette.hpp`, `toolbar_builder.hpp`, `tool_surface*.hpp` |
| QtNetwork (HTTP, sockets, WebSockets) | Covered | `network.hpp` (NetworkAccessManager, WebSocket, RemoteServer/Client, SyncEngine, CloudContext) |
| QtMultimedia / QtSpatialAudio | Partial | `media.hpp` (AudioDeviceManager, AudioProcessor, audio widgets); no spatial-audio wrapper; JUCE bridge planned (`docs/JUCE_ANALYSIS.md`) |
| Qt3D / QtQuick3D (scene graph) | Partial | `media.hpp` ThreeDNode/Hologram/SpatialContext + GSK GPU shaders; no scene-graph API parity |
| QtQuick / QML (declarative UI) | Missing | No declarative language or binding engine; `property.hpp` is imperative only. Largest single gap |
| QtSql | Partial | `database.hpp`, Core `SqlDatabase` (sqlite3 dep now wired in native build); no model/view SQL classes |
| QtCharts / QtDataVisualization | Partial | Core `FinChart`, `OrderBook`, `TradingTerminal` widgets; no general charting API |
| QtSvg | Partial | GDK/GSK SVG loading exists natively (`bobgui/svg`); no Bobtk `Svg` wrapper header |
| QtPrintSupport | Partial | Print headers feed `bobguitypebuiltins` (`bobgui_print_headers`); no print dialog wrapper |
| QtTest | Partial | Tools `TestRunner` wrapper + native `testsuite/`; no QtTest-style API |
| QtConcurrent | Partial | Core `QuantumScheduler`, `RealtimeKernel`, `ComputePipeline`; no futures/promises wrapper |
| QtWebEngine / WebView | Partial | Network `WebView` widget declared; engine integration unverified |
| QtDBus / platform extras | Missing | No DBus wrapper (Linux-only relevance); Windows/macOS backends are GDK-native |

## Priority gaps (feed TODO)
1. Declarative/binding layer (QtQuick/QML equivalent) — design spike, not a port.
2. General charting API beyond FinChart.
3. SQL model/view classes on top of `SqlDatabase`.
4. JUCE-backed multimedia depth (see JUCE analysis).
Qt7 modules are beta/unreleased upstream; re-run this matrix when Qt7 stabilizes rather than chasing previews.
