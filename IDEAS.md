# IDEAS FOR IMPROVEMENT

- **C++ Modernization:** Utilize C++20 features (concepts, modules, coroutines) for wrapping the C-based GObject layer into a more intuitive, standard-library-like API.
- **Go Bindings vs Native Port:** Evaluate whether an intermediate CGO binding layer is better for initial transition before committing to a 100% pure Go rewrite.
- **Cross-Platform Renderers:** Look into abstracting the rendering backend (GDK/GSK) to seamlessly support DearImgui's immediate-mode rendering style alongside traditional retained-mode structures.
- **Submodule Management:** Create a build script that automatically compiles only the required components of JUCE and Ultimate++ to keep build times low.
