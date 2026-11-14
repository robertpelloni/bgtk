# Bobtk User Manual

## Start here

Bobtk has two layers: a **Go dashboard** that runs today, and a **native C++ framework** (Meson) that is still being repaired. This manual mirrors what the dashboard shows — anything not yet built is labelled Planned in both places.

## Run the dashboard

```powershell
cd go
go run . -serve -addr :8484
```

Then open <http://localhost:8484> in a browser. One page, no subpages:

| Section | What it does |
|---|---|
| Overview | Version (from the `VERSION` file), runtime, pillars running, native-build status checked on disk |
| Pillars — run & control | All 6 pillars with Initialize buttons, live Running/Not started state, every backend feature described and badged (Go today / Native C++ / Planned bridge). Search filters this and the catalog together |
| Native C++ framework catalog | Read-only map of `bobgui/cpp/` headers and `tools/` CLI utilities, grouped by pillar |
| Build, logs & system tray | Exact build commands, live server logs (last 300 lines), tray/quit controls |
| Docs, manual & next steps | Pointers to `ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `DEPLOY.md` and the honest not-working list |

Every `?` marker is a tooltip: hover or keyboard-focus it for extra detail. Buttons say what they do; Initialize is safe to click twice (second run is a no-op).

## System tray

Native Windows tray (v1.7.0, Go — no C++ tray exists; GTK4 removed `GtkStatusIcon`):

```powershell
cd go; go run . -serve -tray -addr :8484   # server + tray in one process
# or, next to an already-running server:
bobtk.exe -tray -addr :8484
```

Fallback without the Go binary: `powershell -ExecutionPolicy Bypass -File scripts/tray.ps1`.

Menu: **Open Dashboard** (double-click works too), live **Status**, **Quit Server** (graceful — stops only the dashboard process), **Exit tray icon** (server keeps running). Without the tray, use **Quit server** in the dashboard's Build section; restart any time with the `go run` command above.

## Build everything

```powershell
powershell -File scripts/build.ps1              # Go: vet + test + build
powershell -File scripts/build.ps1 -Serve       # build, then serve dashboard
powershell -File scripts/build.ps1 -Native      # also Meson C++ (currently blocked — see Troubleshooting)
```

Heavy builds run at BelowNormal priority by default so the laptop stays responsive.

## Native CLI tools (verified 2026-11-14)

After the native build, run tools inside the build environment so DLLs resolve (`meson devenv -C _build --dump` gives the PATH; tools fail with `STATUS_DLL_NOT_FOUND` without it):

```powershell
& _build/tools/bobgui4-builder-tool.exe validate examples/builder.ui   # exits 0
& _build/tools/bobgui4-image-tool.exe --help
& _build/tools/bobgui4-path-tool.exe --help
& _build/tools/bobgui4-query-settings.exe   # prints live settings values
```

## Troubleshooting

- **Port busy:** start on another port, e.g. `go run . -serve -addr :8485`.
- **Native `meson setup` fails at `meson.build:343`:** known splice of two build files (see `TODO.md` Native build repair). The Go dashboard does not depend on it.
- **Dashboard shows old version:** it reads the repo-root `VERSION` file at process start — restart the server after a bump.
- **What is not working:** native C++ build (splice repair in progress), Go↔C++ cgo bridge (Planned), native tray icon (use `scripts/tray.ps1`). The dashboard's Docs section repeats this list live.
