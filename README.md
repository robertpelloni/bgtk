# bgtk — a renamed copy of the GTK 4 tree ("Bobgtk" / "BOBGUI")

This repository is a copy of the **GTK 4** source tree (the `gtk/`, `gdk/`,
`gsk/`, `testsuite/`, Meson build, etc. are upstream GTK, copyright the GTK
project and contributors, LGPL-2.1+) that has been mechanically renamed:
"GTK" was string-replaced with "BOBGUI"/"bobgui" throughout the tree,
including a duplicate renamed copy of the library sources under `bobgui/`.

Be aware of what that means:

- The README this replaces presented the renamed tree as a GNOME-hosted
  project with official download sites, documentation and nightly flatpaks
  (`download.gnome.org/sources/bobgui`, `org.bobgui.Demo4`, …). **None of
  those exist** — they are upstream GTK links with the name swapped. The
  real project is at https://gitlab.gnome.org/GNOME/gtk.
- Nothing here has been validated as a working renamed build; treat any
  "BOBGUI" API or file as upstream GTK code with a different string in it.

## What is actually original here

- Agent-era planning docs at the root: `VISION.md`, `UPP_ANALYSIS.md`,
  `IDEA`, `IDEAS.md`, `HANDOFF.md`. `VISION.md` states ambitions (feature
  parity with Qt 6/JavaFX/wxWidgets/WinUI/Dear ImGui, JUCE and Ultimate++
  integration, a full Go port) — these are goals, not implemented work.
- `go/` — a stub Go port: `go.mod`, an 840-byte `main.go`, and a
  `go/bobtk/` directory. A starting point only.
- A rename of the upstream tree itself, per above.

## Building

Upstream GTK builds with Meson (and that build system is intact here):

```bash
meson setup _build
meson compile -C _build
meson test -C _build
```

Whether the renamed `bobgui/` copy compiles as-is is unverified.

## Status

Dormant experiment. Upstream content dates to ~October 2024; the only 2026
activity is an automated agent checkpoint commit (2026-10-19). If you want
GTK, use GTK — this repo is a renaming playground around it, kept for the
BQt/BobUI line of experiments (see also the `bqt` repo).
