# HANDOFF — Bobtk session 2026-11-13 (dashboard v1.1.0 + native configure v1.2.0)

Identity of this agent: **muse-wonder-1.0 in Muse Code in C:\Users\hyper\workspace\bgtk on laptop**. Edge CDP 9222 tabs observed: Muse (muse.ai), ChatGPT, Gemini — this agent did **not** inject messages into those tabs; coordination text lives here and in the dashboard footer.

## What runs today (verified this session)
- Go dashboard http://localhost:8484 (`cd go; go run . -serve -addr :8484`): one page, init-all 6/6, graceful Quit proven. Restart after VERSION bumps (reads `VERSION` at startup).
- `go test ./...`, `go vet ./...`, `go build` green (Go 1.27.4).
- **Native configure + compile green (v1.6.0, persistent):** in-repo `_build` with `-Dmedia-gstreamer=disabled -Dintrospection=disabled -Dvulkan=disabled -Dsysprof=disabled -Dglib:sysprof=disabled -Dlibsoup-3.0:tests=false`; `meson compile -C _build` exits 0 (no-op recompile verified). Test subset: 16 run, only `gdk colorstate` ERROR (missing en_US.UTF-8 locale). Disk is tight — keep >2 GB free.

## Native build repair — how it was fixed (reproduce carefully)
1. Root `meson.build` was two interleaved files (bobgui/GTK4 932 lines + GTK3 ~813). Restored clean bobgui file via `git restore --source=ce1de57f61 --worktree -- <file>`; spliced copies archived under `scripts/archive/` (`meson.build.spliced.txt`, `meson_options.gtk3.txt`, `meson-spliced/…`, `wraps-spliced/…`). Never regex-edit `meson.build` (AGENTS.md).
2. Same restore for 43 files under gdk/bobgui/po/tools/docs/demos/examples/tests/testsuite. `gdk/quartz/meson.build` deletion is intentional (clean tree uses `gdk/macos`).
3. Wrap fixes: truncate second `directory=` orphan tails; `[provide]` for libffi (`ffi_dep`), zlib (`zlib_dep`), intl (`intl_dep`); pango pinned 1.56.4; pixman → freedesktop `pixman-0.46.4` (`idep_pixman`); new wraps libsoup-3.0 3.6.5, sqlite3, nghttp2, libpsl (last three WrapDB; new wraps need `git add -f`, `.gitignore` covers `subprojects/*.wrap`).
4. Disabled for probe: media-gstreamer (gstreamer win-pkgconfig override bug), vulkan (no SDK), introspection (speed).

## What is broken / not running
1. Native compile is **green**; remaining native work: broader `meson test` run, persistent in-repo `_build`, re-enable GStreamer (win-pkgconfig bug) and Vulkan (needs SDK). Warning: restoring from `ce1de57f61` reverts the gdk enum fix; re-apply (CHANGELOG 1.2.1/1.3.0).
2. cgo bridge: none; Go pillars flip state only (badged Planned in UI).
3. Native tray icon: none in C++; use `scripts/tray.ps1`.
4. Submodules: `submodules/juce` dirty (25531a05-dirty vs recorded d60b7a78; checked-out JUCE is 8.0.14); `submodules/ultimatepp` untracked. Needs a deliberate pin decision.
5. CHANGELOG [4.0.8] bridge headers (`juce_audio_bridge.hpp` etc.) do not exist — see `docs/JUCE_ANALYSIS.md`.

## Next queue (see TODO.md)
`meson compile` attempt → fix compile blockers → re-enable gstreamer/vulkan → JUCE bridge spike → Qt-parity matrix.

## Resume commands
```powershell
cd go; go test ./...; go vet ./...
meson setup $env:TEMP\bgtk-probe-build . -Dmedia-gstreamer=disabled -Dintrospection=disabled -Dvulkan=disabled
meson compile -C $env:TEMP\bgtk-probe-build   # BelowNormal; long
```
