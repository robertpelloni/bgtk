# MEMORY

- **Codebase Origins:** Started as Bobgui/GTK codebase, heavily reliant on standard C-based GObject architecture.
- **Target Architecture:** Modern, multi-paradigm C++ combined with a parallel, comprehensive Go port.
- **Design Preferences:** Clean, well-structured, easy to use, well-named, and highly documented.
- **Current Observation:** The project has undergone a massive renaming from `gtk` and `bobgui` to `bobtk`. Pathing issues or build files (e.g., meson.build, Makefiles) may require further refinement if `bobtk` replacements broke regex or internal tooling.
- [x] Evaluated JUCE submodule integration and drafted comparative architecture matrix in `docs/integration/JUCE_BOBTK_MATRIX.md`.
- Next step: Scaffold a minimal JUCE-inspired Audio Widget prototype in the Go port (`/go/visual/` or `/go/media/`) referencing the integration matrix.

- [x] Assimilated `synthetic-motion` event system into the Go `Visual` pillar (`go/bobtk/visual/motion.go`) preserving thread-safety and idiomatic concurrency for C++ interop.

- [x] Integrated `BaseMotionEmitter` into `BaseWidget` within `go/bobtk/visual/widget.go`, ensuring all widgets can handle synthetic motion parity with Qt6 and the C++ backend.

- [x] Evaluated JUCE's submodule integration, mapped its widget architecture to the 6-pillar framework, and proposed a unified interface design for the `go/bobtk/media` pillar (`MediaProcessor` and `MediaContext`).

- [x] Documented immediate conflicts between JUCE and Bobtk in `DEPLOY.md`, specifically concerning `meson.build` wraps and Main Loop contention.

- [x] Scaffolded `LayoutManager`, `BaseLayout`, and `Style` types in `go/bobtk/visual/`, implementing thread-safe property bags and composable interfaces with stubbed documentation for Phase 4 Frontend (Wasm/DOM) expansion.
