### Core Architecture and Vision
*   **Goal**: The `bobtk` project (formerly `gtk` and `bobgui`) aims to become the ultimate multi-paradigm C++/Go UI library. It seeks to integrate features from Qt, JUCE, Ultimate++, JavaFX, WxWidgets, WinUI, and Dear ImGui.
*   **The Six-Pillar Architecture**: The system revolves around six fundamental pillars:
    1.  **Media**: Manages Audio (`AudioDeviceManager`, `AudioProcessor`), 3D spatial properties (`ThreeDNode`), GIS mapping (`MapView`), Shaders (`EffectShader`), and Timeline controls.
    2.  **Network**: Handles network operations and protocols.
    3.  **System**: Manages low-level OS integrations.
    4.  **Visual**: The core UI rendering system and widgets.
    5.  **Core**: Contains heavy-lifting sub-systems such as `AutonomousController` (AI control), `BrainNode` (LLM Inference), `ComputePipeline`, `EntityManager` (ECS), `FinChart`, `MetaOrchestrator`, `QuantumScheduler`, `RealtimeKernel`, and `StateMachine`.
    6.  **Tools**: Encompasses IDE integrations, compilers, unit testing, and reporting (e.g., `StudioManager`, `TestRunner`).
*   **C-to-C++ Abstraction**: The backend is built upon standard C-based `GObject` architecture (inherited from GTK). The recent focus has been on wrapping these raw C `GObject` functions and stubs into modern C++ abstractions.
*   **The Go Port (`/go`)**: A parallel objective is to cleanly port this architecture to Go. The C++ framework is meant to be bridged to Go via `cgo`, matching the 6-pillar structure at a high level.
*   **Submodule Integration**: The project includes `JUCE` and `Ultimate++` as git submodules. The objective is to adopt their architectural design patterns and allow developers to seamlessly intermix widgets from these frameworks.

### Codebase State & Observations
*   **Mass Renaming Corruption**: The project recently experienced a mass regex replacement, changing `gtk` and `bobgui` to `bobtk`. This broke the `meson.build` and `.wrap` files, leaving the C/C++ build system currently inoperable.
*   **Backend Stubs**: The core and media C++ wrappers exist, but the underlying C implementation files (e.g., `bobguiaudio.c`, `bobgui3d.c`, `bobguidata.c`) are currently empty or return `NULL` dummy stubs. They are waiting for actual logic implementation.

### Agent Guidelines
*   **Backwards Compatibility**: Keep the project backwards-compatible with legacy GTK software.
*   **Documentation**: Heavy emphasis is placed on deep, insightful documentation. Every UI element requires labels, tooltips, and descriptions.
*   **Model-Specific Prompts**: Specific models are expected to assume distinct roles (e.g., Claude for cross-file architecture analysis, Gemini for expansive bridging of submodules, GPT for precise TDD generation, Copilot for C++20 standard practices).

### Next Steps Forward
*   The highest priority is to safely and manually repair the `meson.build` and subproject `.wrap` dependencies to restore the C++ compile chain.
*   The `/go` module architecture has been successfully initialized (mirroring the 6 pillars) and needs to be hooked up to the C++ backend once it compiles.