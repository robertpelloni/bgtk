# MEMORY

Durable facts, verified against the tree (2026-11-13). Session detail: `HANDOFF.md`.

- **Origins:** GTK 4 tree mechanically renamed to Bobgui/Bobtk, plus a duplicate renamed copy under `bobgui/`. Upstream is https://gitlab.gnome.org/GNOME/gtk; renamed APIs are upstream code until proven otherwise (`README.md`).
- **Architecture:** 6 pillars (Core, Visual, Media, Network, System, Tools) in three places: C++ headers `bobgui/cpp/` (+ `module/`), Go stubs `go/bobtk/`, dashboard `go/bobtk/dashboard/` (the only runnable UI today).
- **Dashboard rule:** every backend feature gets a dashboard card with label + description + `?` tooltip and an honest badge (Go today / Native C++ / Planned bridge). No subpages — sections on one page, value-ordered.
- **Build reality:** Go builds/tests green offline (stdlib only). Native Meson build is spliced (two build files interleaved) — canonical half is bobgui/GTK4; see `TODO.md` splice map.
- **Versioning:** repo-root `VERSION` is the single source; dashboard reads it at startup. Commit messages name the version.
- **Preferences:** concise docs that mirror the UI; heavy processes at BelowNormal priority; never terminate unrelated processes; tray = `scripts/tray.ps1` until a native icon exists.
- **Submodules:** JUCE + Ultimate++ under `submodules/`; pins currently dirty/untracked — treat with care, don't commit blindly.
