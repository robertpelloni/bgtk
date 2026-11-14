# TODO

Granular queue. The 15-minute autonomous loop works top-down. Checkboxes only move on verified evidence.

## Native build repair (highest value)
- [ ] Untangle spliced `meson.build`: bobgui/GTK4 block (lines ~1–334, `bobgui_*` vars, `GETTEXT_PACKAGE bobgui40`) vs gtk/GTK3 block (from ~line 342, `gtk_*` vars, `gtk30`, `PACKAGE_NAME gtk+`). Keep bobgui canonical; archive the GTK3 half to `scripts/archive/` before deleting.
- [ ] Re-run probe after each splice fix: `meson setup $env:TEMP\bgtk-probe-build .` (BelowNormal). Next known error: line 343 `check_headers` list unterminated by the splice.
- [ ] Audit `subprojects/*.wrap` for rename damage (missing `revision` keys etc.), fix, then `meson setup _build` + compile on Windows.
- [ ] Fix submodule state: `submodules/juce` dirty pin (25531a05-dirty vs recorded d60b7a78); `submodules/ultimatepp` present but untracked. Record intended pins, commit cleanly.

## Dashboard & UI
- [x] Consolidated single-page dashboard with tooltips/labels (1.1.0).
- [ ] Native tray icon in C++ (or Go systray lib when network allows) to replace `scripts/tray.ps1`.
- [ ] Add per-feature dashboard actions as bridges land (e.g. Media play/stop once cgo exists).

## Bridges & parity
- [ ] Analyze JUCE UI widgets for intermixing (carried from Phase 1).
- [ ] Research missing Qt6/Qt7 features vs C++ baseline; write `docs/QT_PARITY_MATRIX.md`.
- [ ] First cgo bridge spike (Core or Media) behind a build tag, off by default.

## Housekeeping
- [ ] Decide fate of root `check-version.py` / `plan_step_complete.py` (archive copies exist in `scripts/archive/`; root copies kept for now).
- [ ] Remove stale `HANDOFF.md` claims as milestones land (XKB bug already fixed in tree; meson wrap story superseded by splice finding).
