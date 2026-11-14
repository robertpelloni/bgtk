# scripts/ — consolidated entry points

One script per job. Originals at the repo root were **copied** (not moved) into `archive/` on 2026-11-13 for backup; the root files still work.

| Script | Job |
|---|---|
| `build.ps1` | Unified build: Go vet + test + build; `-Native` adds Meson; `-Serve` runs the dashboard; runs BelowNormal by default |
| `tray.ps1` | Windows system-tray icon for the dashboard (Open / Status / Quit Server) |
| `archive/check-version.py` | Upstream GTK version checker (backup of root `check-version.py`) |
| `archive/plan_step_complete.py` | Trivial stub kept for history (backup of root file) |

Root one-offs not yet consolidated: `make-pot` (translations, upstream), `check-version.py` (use the archive copy going forward).
