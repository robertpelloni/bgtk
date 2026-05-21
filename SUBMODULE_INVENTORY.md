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
