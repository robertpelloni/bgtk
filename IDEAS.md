# IDEAS FOR IMPROVEMENT

Kept deliberately aggressive; committed plans live in `ROADMAP.md`/`TODO.md`.

- **C++ Modernization:** Utilize C++20 features (concepts, modules, coroutines) for wrapping the C-based GObject layer into a more intuitive, standard-library-like API.
- **Go Bindings vs Native Port:** Evaluate whether an intermediate CGO binding layer is better for initial transition before committing to a 100% pure Go rewrite. (First spike queued in `TODO.md`.)
- **Cross-Platform Renderers:** Look into abstracting the rendering backend (GDK/GSK) to seamlessly support DearImgui's immediate-mode rendering style alongside traditional retained-mode structures.
- **Submodule Management:** Create a build script that automatically compiles only the required components of JUCE and Ultimate++ to keep build times low.
- **Dashboard as test oracle:** drive the dashboard's `/api/status` from CI so badge honesty (Go/Native/Planned) is machine-checked, not hand-waved.
- **One-binary distribution:** `go build` already yields a single `bobtk.exe` with the page embedded — add tray + auto-open-browser for a double-clickable demo.
- **Splice archaeology tool:** script that diffs the two interleaved `meson.build` halves against upstream GTK3/GTK4 tags to auto-propose the canonical file.
