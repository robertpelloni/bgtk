# DEPLOY — Bobtk environment setup

## Go dashboard (works today, offline)
Requires Go 1.24+ (verified 1.27.4 on Windows).

```powershell
git clone --recurse-submodules <repository_url>
cd go
go run . -serve -addr :8484   # dashboard at http://localhost:8484
powershell -File ..\scripts\tray.ps1   # optional Windows tray icon
```

Unified alternative: `powershell -File scripts/build.ps1 -Serve`. Version comes from the repo-root `VERSION` file at process start.

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
