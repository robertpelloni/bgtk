# DEPLOY — Bobtk environment setup

## Go dashboard (works today, offline)
Requires Go 1.24+ (verified 1.27.4 on Windows).

```powershell
git clone --recurse-submodules <repository_url>
cd go
go run . -serve -addr :8484   # dashboard at http://localhost:8484
powershell -File ..\scripts\tray.ps1   # optional Windows tray icon
```

Unified alternative: `powershell -File scripts/build.ps1 -Serve`. Version comes from the repo-root `VERSION` file at process start.

## Native C++ (Meson) — currently blocked
Requires Meson + Ninja (verified present: 1.12.1 / 1.12.1 on this laptop) and a C toolchain.

```bash
meson setup _build
meson compile -C _build
meson test -C _build
```

**Known blocker (2026-11-13):** root `meson.build` splices a bobgui/GTK4 build file with a gtk/GTK3 one; configure fails at line 343 after this session's header fix. Repair queue: `TODO.md` → "Native build repair". Run probes at BelowNormal priority and into `$env:TEMP` until the tree is clean. Submodule pins (`submodules/juce` dirty, `ultimatepp` untracked) must be resolved before a reproducible native build — see `SUBMODULE_INVENTORY.md`.

## Logs & operations
- Dashboard logs: server stdout/stderr plus the in-page Live logs (last 300 lines, `GET /api/logs`).
- Health check: `GET /health`. Graceful stop: dashboard **Quit server** button, tray **Quit Server**, or `POST /api/quit`.
