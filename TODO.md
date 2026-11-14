# TODO

Granular queue. The 15-minute autonomous loop works top-down. Checkboxes only move on verified evidence.

## Native build repair (highest value)
- [x] Untangle spliced `meson.build`: restored clean bobgui file from git history (`ce1de57f61`, 932 lines); spliced file archived at `scripts/archive/meson.build.spliced.txt`; GTK3 `meson_options.txt` archived to `scripts/archive/meson_options.gtk3.txt`.
- [x] Truncate orphan GTK3 tails in 8 wraps (backups in `scripts/archive/wraps-spliced/`); add missing `[provide]` mappings: libffi, zlib, intl (proxy-libintl); pin pango `main` → `1.56.4` (glib 2.84 era); switch pixman centricular fork → upstream `pixman-0.46.4` (exports `idep_pixman`); add wraps: libsoup-3.0 (3.6.5), sqlite3, nghttp2, libpsl (WrapDB).
- [x] `meson setup` probe **passes** (2026-11-13): `meson setup $env:TEMP\bgtk-probe-build . -Dmedia-gstreamer=disabled -Dintrospection=disabled -Dvulkan=disabled` exits 0. GStreamer disabled due to its Windows win-pkgconfig override bug; Vulkan disabled (no SDK on this laptop).
- [x] `meson compile` **exits 0** (v1.5.0). Blockers fixed in order: sysprof off; WrapDB libffi 3.8.0; cairo 1.18.4; SCSS conflict resolved; gdk template enums; gresource generator; gdkconfig template; 271 restored sources; modules-enum exclusion (Windows backslash filter); 60 modules stubs excluded; libsoup tests off (`-Dlibsoup-3.0:tests=false`). Logs: `logs/native-compile19.log`.
- [ ] Run broader native tests (`meson test`, 7779 total). Done on `_build`: bobgui subset 16 run (1 locale ERROR); gdk+gsk compare subset stopped at 1527/6124 (1129 OK, 11 image-compare FAIL, 317 TIMEOUT, 10 ERROR); unit run (compares excluded, 1274 tests) completed: **102 OK, 1158 Fail (all exit-status-3 display-init bails: reftest/nodeparser/GL), 14 TIMEOUT, zero TAP `not ok`/assertion failures**. Conclusion: no code assertion failures anywhere; display-dependent suites need an interactive desktop session with GPU — rerun there before claiming suite health. Diagnostic (2026-11-14): a single reftest run directly (`bobgui-reftest.exe --tap -k ...ui`, devenv PATH, diagnostics on) exits 3 immediately after `Failed to set locale to en_US.UTF-8` with no test output — failure is in harness/init (Windows locale/init path), before any test executes.
- [x] Persistent `_build` in-repo configured + compiled green (v1.6.0); dashboard reports it. Disk watch: keep >2 GB free (build + links fill C:).
- [ ] Re-enable media-gstreamer: root cause identified (2026-11-14) — gstreamer's `win-pkgconfig` subproject calls `meson.override_find_program('pkg-config', ...)` unconditionally on Windows when `build-tools-source=subproject` (default), colliding with the system pkg-config found earlier. Unblock: `-Dgstreamer:build-tools-source=system` plus flex/bison/nasm on PATH (none installed system-wide today; nasm binary already at `subprojects/gstreamer/subprojects/win-nasm/nasm-2.15.04/nasm.exe`). Needs a flex/bison install decision — not done autonomously (system change).
- [ ] Re-enable vulkan: needs Vulkan SDK (`VULKAN_SDK`) installed — user/machine decision.
- [ ] Fix submodule state: `submodules/juce` dirty pin (25531a05-dirty vs recorded d60b7a78); `submodules/ultimatepp` present but untracked. Record intended pins, commit cleanly.

## Dashboard & UI
- [x] Consolidated single-page dashboard with tooltips/labels (1.1.0).
- [ ] Native tray icon in C++ (or Go systray lib when network allows) to replace `scripts/tray.ps1`.
- [ ] Add per-feature dashboard actions as bridges land (e.g. Media play/stop once cgo exists).

## Bridges & parity
- [x] Analyze JUCE UI widgets for intermixing → `docs/JUCE_ANALYSIS.md` (Media-first; note CHANGELOG [4.0.8] bridge files do not exist).
- [x] Research missing Qt6/Qt7 features vs C++ baseline → `docs/QT_PARITY_MATRIX.md` (top gaps: declarative/QML layer, charting, SQL model/view).
- [x] JUCE bridge spike: `bobgui/cpp/juce_audio_bridge.hpp` (device enumeration + sample rate + `JuceProcessorChain::process_block` DSP forwarding, `BOBGUI_WITH_JUCE`-gated, syntax-verified guard-off; enabled path uncompiled here — verify before use).
- [x] First cgo bridge spike: `go/bobtk/bridge` behind `bobtk_cgo` tag (stub tested green by default). cgo compile currently blocked on this laptop: Go 1.27.4 cgo vs GCC 15.2 pe-bigobj objects (see `bridge_cgo.go` note) — needs older GCC or newer Go.
- [ ] Link bridge against `libbobgui` from `_build` (Media `processBlock` forwarding) once cgo toolchain unblocks.

## Housekeeping
- [ ] Decide fate of root `check-version.py` / `plan_step_complete.py` (archive copies exist in `scripts/archive/`; root copies kept for now).
- [ ] Remove stale `HANDOFF.md` claims as milestones land (XKB bug already fixed in tree; meson wrap story superseded by splice finding).
