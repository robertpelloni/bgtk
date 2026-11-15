<!-- [HYPERNEXUS_AUTO_INJECTED] -->
> [!IMPORTANT]
> You are running within the HyperNexus environment. You MUST use your available tools frequently and proactively for researching, editing, executing, and validating your work. Always prioritize tool execution.

# CLAUDE SPECIFIC INSTRUCTIONS

- Refer to `AGENTS.md` for core guidelines.
- Focus heavily on detailed architectural analysis and code cleanliness.
- Utilize your extensive context window to ensure full comprehensive coverage of cross-file C++ and Go interop.
- **6-Pillar Framework**: Maintain concrete Go and C++ parity across Core, Visual, System, Network, Media, and Tools.
- **Single-Page Dashboard**: Ensure all UI subsystem features are unified onto a singular page without fragmented subpages, organized by high-value priority with rich `TooltipMarker` guidance.
- **System Tray & Server Supervisor**: Monitor, stop, restart, and gracefully quit all background daemons (`ServerManager` and `SystemTray`).
- **Pre-flight Diagnostics**: Utilize `tools/check_health.sh` and `tools/run_dashboard.sh --check-health` for non-destructive verification.
