# Submodule and Library Inventory

## Submodules
1. **Ultimate++** (`submodules/ultimatepp`)
   - **Purpose**: Highly optimized C++ framework for high-performance memory allocation, UI widgets, and multithreading.
   - **Integration Status**: Analyzed. Memory pool, variant arrays, and callbacks identified for wrapping.

2. **JUCE** (`submodules/juce`)
   - **Purpose**: Leading C++ framework for audio DSP, plugin wrappers (VST3/AU), and cross-platform UI.
   - **Integration Status**: Pending analysis. Aiming to bridge its `AudioProcessor` logic and UI widgets.

## Dependency Wraps (`subprojects/*.wrap`)
These dependencies are heavily used by the backend C build logic (Meson).
*Note: Due to a mass-rename corruption, many of these wrap files are currently damaged.*
- **atk, cairo, pango, fontconfig, freetype2, fribidi, harfbuzz**: Essential rendering and text shaping.
- **glib, gobject-introspection, libffi**: Core object system mapping and binding generation.
- **gdk-pixbuf, libjpeg-turbo, libpng, librsvg, libtiff, pixman**: Image rendering buffers.
- **gstreamer-full**: Multimedia processing framework.
- **wayland, wayland-protocols**: Next-gen Linux display server protocols.

## Languages and Compilers
- **C/C++**: The foundation of the framework using GObject architecture and modern C++ abstractions.
- **Go 1.24.3**: Targeted for the new high-level architectural port.

## Status note (2026-11-13, updated 2026-11-14)
- Verified on disk: submodules/juce dirty pin (25531a05-dirty vs recorded d60b7a78), submodules/ultimatepp untracked. Dashboard represents JUCE/Ultimate++ honestly as Native/Planned, never as running.
- JUCE detail (2026-11-14): HEAD is 25531a0596 ("X11: Add XInput multitouch support", JUCE 8.0.14); working tree contains **unresolved merge conflicts** (`UU examples/DemoRunner/Builds/LinuxMakefile/Makefile`, `extras/AudioPluginHost/...`, `extras/NetworkGraphicsDemo/...`, `extras/UnitTestRunner/...`) plus modifications (BREAKING_CHANGES.md, juce_opengl.cpp). Do not commit the superproject pin until the JUCE merge is resolved or the tree is reset to a clean tag (8.0.x) — owner decision.
- Ultimate++ detail: has its own git history (HEAD e3546f5bc "chore: auto checkpoint before sync") but the superproject has no gitlink for it (`.gitmodules` entry exists, index does not). Add as proper submodule at a recorded commit or remove the `.gitmodules` stanza — owner decision.
