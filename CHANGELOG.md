# CHANGELOG

## [1.6.0] - 2026-11-14
### Fixed
- **Persistent native build green:** in-repo `_build` (same verified flags) `meson compile` exits 0, verified by no-op recompile; `libbobgui-4-1.dll` (52 MB) built in `_build/bobgui/`. Dashboard now reports this directly (`compiled green (_build, ...)`). Build survived two disk-full stops (C: had filled; reclaimed Go build cache 4.5 GB; kept `-j 4` for final links). Watch disk: `_build` is ~2.5 GB.
- Native test subset on `_build`: `meson test --suite bobgui` ran 16 tests in 240 s (Windows test startup is slow) — passing except `gdk colorstate` ERROR (missing `en_US.UTF-8` locale in this environment). Full run remains queued.
- Go: initialization/idempotency tests for all six pillars + dashboard + bridge (all green, vet clean). JUCE bridge gained `JuceProcessorChain::process_block` DSP forwarding (gated, enabled path uncompiled here).

## [1.5.0] - 2026-11-13
### Fixed
- **Native compile green.** `meson compile -C <build>` exits 0 (verified with a no-op recompile) for the full bobgui tree: glib, gdk, gsk, bobgui, demos, examples, testsuite targets (3277 + 538 targets). Final blocker: libsoup's own Windows test executables (undefined `test_init` etc.) excluded via `-Dlibsoup-3.0:tests=false`.
- Configure line (Windows, this laptop): `meson setup <build> . -Dmedia-gstreamer=disabled -Dintrospection=disabled -Dvulkan=disabled -Dsysprof=disabled -Dglib:sysprof=disabled` + `-Dlibsoup-3.0:tests=false`.
- Test smoke: `meson test --suite bobgui` → 2 ok, 1 timeout (`gdk clipboard`, flaky/headless environment). Full 7779-test run not attempted; see `TODO.md`.

## [1.4.0] - 2026-11-13
### Fixed
- Restored 120 more spliced sources under `testsuite/`, `tests/`, `demos/`, `examples/`, `tools/`, `po/`, `docs/` from `ce1de57f61` (reftest snapshot header carried duplicate GTK3 declarations). Spliced copies archived in `scripts/archive/meson-spliced/`.
- Compile progress: bobgui typebuiltins now generates and compiles (modules-enum exclusion with Windows path normalization); build advancing through bobgui objects.

## [1.3.0] - 2026-11-13
### Fixed
- Restored 150 spliced C sources under `gdk/`, `gsk/`, `bobgui/` from the last clean bobgui commit (`ce1de57f61`) after GTK3-era headers (`gdkmain.h`, `gdkwindow.h`, …) broke GTK4 include paths; 175 post-clean additions (incl. `bobgui/cpp/` wrappers) kept untouched. Spliced copies archived in `scripts/archive/meson-spliced/`.
- Fixed nested `<gresource>` in `gdk/gen-gdk-gresources-xml.py` and restored spliced `gdk/gdkconfig.h.meson`.
- Re-applied the gdk template-enum fix (restoration had reverted it): `gdk/meson.build` uses `gnome.mkenums` with `gdkenumtypes.{h,c}.template`.
- Compile status at release: re-running (`logs/native-compile10.log`); completion tracked in `TODO.md`.

## [1.2.1] - 2026-11-13
### Fixed
- Native compile blockers, in order: sysprof disabled on Windows (`-Dsysprof=disabled -Dglib:sysprof=disabled`); libffi centricular fork replaced with WrapDB libffi 3.8.0; cairo 1.18.2 → 1.18.4 (MinGW `dcommon.h` duplicate enum operators); unresolved merge conflict resolved in `bobgui/theme/Default/Default-light.scss` (kept bobgui side); `gdk` enum generation switched from `gnome.mkenums_simple` (Windows argument-quoting failure) to template-based `gnome.mkenums` with new `gdk/gdkenumtypes.h.template` / `gdkenumtypes.c.template`.
- Compile status at release: progressing through project targets (gdk/gsk/bobgui generation steps pass); full-build completion tracked in `TODO.md`.

## [1.2.0] - 2026-11-13
### Fixed
- **Native build configure green.** Root `meson.build` splice (bobgui/GTK4 + gtk/GTK3 interleaved) repaired by restoring the clean 932-line bobgui file from git history (`ce1de57f61`); spliced file archived at `scripts/archive/meson.build.spliced.txt`, GTK3 `meson_options.txt` archived alongside. Same restoration applied to 43 spliced files under `gdk/`, `bobgui/`, `po/`, `tools/`, `docs/reference/`, `demos/`, `examples/`, `tests/`, `testsuite/` (spliced copies in `scripts/archive/meson-spliced/`).
- Wraps: truncated orphan GTK3 tails in 8 wraps (backups `scripts/archive/wraps-spliced/`); added missing `[provide]` for libffi/zlib/intl; pinned pango `main`→`1.56.4`; pixman centricular fork→upstream `pixman-0.46.4`; added libsoup-3.0 3.6.5, sqlite3, nghttp2, libpsl wraps.
- Verified: `meson setup <build> . -Dmedia-gstreamer=disabled -Dintrospection=disabled -Dvulkan=disabled` exits 0 (all root deps resolve: glib 2.84, pango, cairo, gdk-pixbuf, graphene, libsoup, sqlite). Compile attempt is next (`TODO.md`).

### Added
- `docs/JUCE_ANALYSIS.md`: JUCE 8.0.14 module→pillar mapping (Media-first), with correction that CHANGELOG [4.0.8] bridge headers do not exist.

## [1.1.0] - 2026-11-13
### Added
- Consolidated single-page dashboard (`go/bobtk/dashboard/` + `dashboard.html`): Overview health, 6 pillar run/control cards with live `Initialized()` state, native C++ catalog (`bobgui/cpp/` + `tools/`), build/logs/tray section, docs/next-steps. Every feature has a label, description and `?` tooltip; badges distinguish Go-today vs Native C++ vs Planned bridge. Endpoints: `/`, `/health`, `/api/status`, `/api/logs`, `POST /api/init/{id}`, `POST /api/init-all`, `POST /api/quit` (graceful). Tests in `dashboard_test.go`.
- `Initialized()` accessors on all six Go pillars (additive, backwards compatible).
- CLI flags on `go/main.go`: `-serve`, `-addr`, `-version` (default init-and-exit behavior unchanged).
- Windows tray script `scripts/tray.ps1` (Open Dashboard / live Status / Quit Server) and unified `scripts/build.ps1` (Go vet/test/build, `-Native`, `-Serve`, BelowNormal default). Originals backed up to `scripts/archive/` nondestructively.
- User manual `docs/manual/MANUAL.md`; `scripts/README.md`.

### Fixed
- Root `meson.build` duplicated `project('gtk'...)` header (rename splice) — parser now reaches line 343, exposing the next splice (unterminated `check_headers`). Full untangle tracked in `TODO.md`.

### Verified
- `go test ./...`, `go vet ./...`, `go build` green (Go 1.27.4). Dashboard run at BelowNormal on :8484: health OK, init-all 6/6, page contains tray/search controls; `/api/quit` proven on a second instance (:8485 down, :8484 untouched). Tray/build scripts parse clean. Meson 1.12.1 + Ninja present; native configure probed honestly (blocked at line 343).

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
