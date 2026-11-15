# GPT SPECIFIC INSTRUCTIONS

- Refer to `AGENTS.md` for core guidelines.
- Focus on precise code generation, bug fixing, and algorithmic conversions from C/C++ to Go.
- Rely heavily on TDD (Test Driven Development).
- **Thread Safety & Concurrency**: Ensure all Go managers (`ServerManager`, `SystemTray`, `Clipboard`, `WebSocketHub`, `QuantumScheduler`) use idiomatic synchronization without lock contention.
- **Subproject Wraps & Build Files**: Exercise extreme caution with `meson.build` and `.wrap` files, ensuring subproject options (such as `gtk=false` in sysprof) are never blindly replaced.
- **Automated Test Validation**: Run all unit tests (`cd go && go test -v ./...`) and ensure 100% pass rates prior to pre-commit.
