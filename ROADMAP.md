# ROADMAP

Single source of truth for long-term milestones. Granular tasks live in `TODO.md`; session history in `HANDOFF.md`. Version: see `VERSION` (currently 1.1.0).

## Milestone 1 — Runnable consolidated dashboard (done in 1.1.0)
- [x] Single-page Go dashboard (`go/bobtk/dashboard/`) representing all 6 pillars + native C++ catalog, with labels, descriptions and `?` tooltips; no subpages.
- [x] Live status from real pillar `Initialized()` flags; honest badges (Go today / Native C++ / Planned bridge).
- [x] Tray-equivalent controls (Quit server) + real Windows tray script (`scripts/tray.ps1`).
- [x] Unified build script (`scripts/build.ps1`), user manual (`docs/manual/MANUAL.md`).

## Milestone 2 — Native C++ build green (done; persistent `_build` green in 1.6.0)
- Spliced build files and C sources restored from the clean bobgui commit (`ce1de57f61`); wraps repaired (see CHANGELOG 1.2.0–1.4.0 for the full list). Canonical tree is bobgui (GTK4 lineage); GTK3 halves archived in `scripts/archive/`, not deleted.
- `meson setup` exits 0 (flags in `TODO.md`). `meson compile` is progressing through all libraries/demos/testsuite with blockers fixed in order (currently `logs/native-compile18.log`).
- Remaining in this milestone: compile exit 0 → `meson test` subset → re-enable GStreamer/Vulkan variants.

## Milestone 3 — Go↔C++ bridge
- cgo binding for one pillar end-to-end (suggest Core `QuantumScheduler` or Media `AudioProcessor`), then replicate the pattern.
- JUCE widget analysis (still pending from Phase 1) feeding the Media bridge; Ultimate++ mapping already in `UPP_ANALYSIS.md`.

## Milestone 4 — Parity research
- Qt6/Qt7 feature-gap matrix vs the C++ baseline; JavaFX/wxWidgets/WinUI/Dear ImGui paradigm notes in `docs/`.
- Every new backend feature must land with dashboard representation (label + description + tooltip) in the same change.

## Standing rules
- Backwards compatibility with GTK/Bobgui implementations where possible.
- `VERSION` file is the only version source; bump + commit message must name it.
- Heavy builds run at BelowNormal priority on this laptop.
