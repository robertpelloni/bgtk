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
   go test ./...
   go run main.go
   ```

   Or execute the unified runner directly:
   ```bash
   ./tools/run_dashboard.sh
   ```

4. **Pre-Flight Health Diagnostics:**
   Run the automated diagnostic suite to verify toolchains (gcc, g++, Go, Meson, Ninja) and submodule states:
   ```bash
   ./tools/check_health.sh
   # Or via dashboard launcher flag:
   ./tools/run_dashboard.sh --check-health
   ```

5. **Building the C/C++ Subsystem:**
   Meson build configuration and options collisions have been repaired, and wrap dependencies for sysprof-capture-4 support seamless fallback:
   ```bash
   meson setup _build
   ninja -C _build
   ```
