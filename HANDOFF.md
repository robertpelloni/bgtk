# HANDOFF — Bobtk (current, v1.6.0, 2026-11-14)

Identity of this agent: **muse-wonder-1.0 in Muse Code in C:\Users\hyper\workspace\bgtk on laptop**. Edge CDP 9222 coordination completed 2026-11-14 via agent-browser: one identical coordination note (identity + dashboard v1.6.0 localhost:8484 + native build green + HANDOFF pointer) posted and verified in each of Muse (muse.ai), Gemini, and ChatGPT tabs (composers were empty; drafts cleared after send).

## What runs today (verified)
- Go dashboard http://localhost:8484 (`cd go; go run . -serve -addr :8484`): one page (Overview/Pillars/Catalog/Build-logs-tray/Docs), init-all 6/6, graceful Quit proven, browser-verified (tooltips, search, logs, 58 FPS). Restart after VERSION bumps (reads `VERSION` at startup); native status is read from `_build` on disk.
- `go test ./...`, `go vet ./...`, `go build` green (Go 1.27.4): pillars, dashboard, bridge stub all tested.
- **Native build green (persistent `_build`):** flags `-Dmedia-gstreamer=disabled -Dintrospection=disabled -Dvulkan=disabled -Dsysprof=disabled -Dglib:sysprof=disabled -Dlibsoup-3.0:tests=false`; `meson compile -C _build` exits 0; `libbobgui-4-1.dll` built; CLI tools smoke-verified under `meson devenv` PATH (`builder-tool validate` exits 0).
- Native tests: zero TAP assertion failures anywhere. Unit run (1274): 102 OK, rest display-init bails (exit 3, incl. reftest harness exiting right after the Windows locale warning). Compare run stopped at 1527/6124 (1129 OK, 11 image-pixel FAILs). Rendering/GL suites need an interactive desktop/GPU session.

## How the native build was repaired (reproduce carefully)
1. Tree interleaved bobgui/GTK4 + GTK3. Restored clean files via `git restore --source=ce1de57f61 --worktree -- <file>` (44 build files, 271 sources); spliced copies archived under `scripts/archive/`. Never regex-edit `meson.build` (AGENTS.md).
2. Wrap fixes: truncate second `directory=` orphan tails; `[provide]` for libffi/zlib/intl; pango 1.56.4; cairo 1.18.4; pixman freedesktop 0.46.4; WrapDB libffi 3.8.0/sqlite3/nghttp2/libpsl; libsoup-3.0 3.6.5 (new wraps need `git add -f`).
3. Compile blockers in order: sysprof off, libffi, cairo, SCSS conflict, gdk template enums (restoring from ce1de reverts this — re-apply, CHANGELOG 1.2.1), gresource generator, gdkconfig, spliced sources, modules-enum exclusion (Windows backslashes), 60 modules stubs excluded, libsoup tests off.

## What is not running / needs decisions
1. GStreamer: `-Dgstreamer:build-tools-source=system` + flex/bison install (machine decision). Vulkan: needs SDK.
2. cgo link: spike exists (`go/bobtk/bridge`, `bobtk_cgo` tag, stub green); compilation blocked by Go 1.27.4 vs GCC 15.2 pe-bigobj — older GCC or newer Go.
3. Native tray icon: none; use `scripts/tray.ps1` (Open/Status/Quit Server).
4. Submodules: JUCE tree has **unresolved merge conflicts** at 8.0.14 (`25531a0596`) — resolve or reset to clean tag before pinning; Ultimate++ has history but no superproject gitlink — add properly or drop stanza. Owner decisions (see `SUBMODULE_INVENTORY.md`).
5. Full native suite health: needs desktop/GPU session (see above).

## Resume commands
```powershell
cd go; go test ./...; go vet ./...
powershell -File scripts/build.ps1 -Serve        # dashboard
powershell -File scripts/build.ps1 -Native       # persistent _build (green flags encoded)
meson test -C _build --no-rebuild --suite bobgui # subset; display suites need desktop
```
