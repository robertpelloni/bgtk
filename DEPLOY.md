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
