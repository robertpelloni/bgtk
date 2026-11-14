# scripts/ — consolidated entry points

One script per job. Originals at the repo root were **copied** (not moved) into `archive/` on 2026-11-13 for backup; the root files still work.

| Script | Job |
|---|---|
| `build.ps1` | Unified build: Go vet + test + build; `-Native` adds Meson; `-Serve` runs the dashboard; runs BelowNormal by default |
| `tray.ps1` | Windows system-tray icon for the dashboard (Open / Status / Quit Server) |
| `archive/check-version.py` | Upstream GTK version checker (backup of root `check-version.py`) |
| `archive/plan_step_complete.py` | Trivial stub kept for history (backup of root file) |

Root one-offs retired 2026-11-14: `check-version.py` and `plan_step_complete.py` removed from the root (upstream leftovers; canonical copies in `archive/`, originals in git history). Remaining root script: `make-pot` (translations, upstream).
