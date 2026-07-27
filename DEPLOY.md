# DEPLOYMENT INSTRUCTIONS

*Deployment procedures are currently in early development as the Go port and C++ integrations evolve.*

1. **Clone the Repository:**
   ```bash
   git clone --recurse-submodules <repository_url>
   ```

2. **Update Submodules:**
   ```bash
   git submodule update --init --recursive
   ```

3. **Building the Go Architecture:**
   ```bash
   cd go
   go run main.go
   ```

4. **Building the C/C++ version:**
   - **WARNING**: The C/C++ `meson.build` system is currently experiencing failures due to a blind mass rename operation (`gtk` -> `bobtk`) which broke Subproject Wrap references. Manual intervention is required before running:
   ```bash
   # Pending wrap repairs
   meson setup _build
   ninja -C _build
   ```

## JUCE / Bobtk C++ Build Conflicts (Immediate Blockers)
- **Meson Wraps:** The `meson.build` dependency resolution is currently failing for subprojects (e.g., `libxml2`, `vpx`). The recent mass rename from `gtk` to `bobtk` requires manual auditing of all `.wrap` files in `subprojects/` to ensure they point to the correct upstream repositories or internal mirrors, as this blocks compiling the C++ bridging layer required for the Go `/media` pillar's JUCE integration.
- **Main Loop Contention:** Both JUCE (`MessageManager`) and Bobtk/GTK (`GMainLoop`) demand control of the main application thread. A strategy to nest or poll one loop inside the other must be devised before linking the UI layers.
