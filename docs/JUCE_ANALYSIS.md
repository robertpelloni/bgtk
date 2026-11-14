# JUCE Analysis — intermixing capabilities (2026-11-13)

Source inspected: `submodules/juce` (JUCE 8.0.14 per its `CMakeLists.txt`), `bobgui/cpp/media.hpp`, `bobgui/cpp/visual.hpp`. This completes the Phase-1 task "Analyze JUCE UI widgets for intermixing capabilities" carried in `TODO.md`.

## JUCE modules present
`juce_core`, `juce_events`, `juce_data_structures`, `juce_graphics`, `juce_gui_basics`, `juce_gui_extra`, `juce_audio_basics`, `juce_audio_devices`, `juce_audio_formats`, `juce_audio_processors` (+ headless), `juce_audio_utils`, `juce_dsp`, `juce_opengl`, `juce_osc`, `juce_midi_ci`, `juce_video`, `juce_javascript`, `juce_box2d`, `juce_animation`, `juce_cryptography`, `juce_analytics`, `juce_product_unlocking`.

## Mapping to Bobtk (Media pillar first — highest value)
- **AudioDeviceManager** (`media.hpp`: initializeDevices/setSampleRate/play/stop) ← `juce_audio_devices` (`AudioDeviceManager`, `AudioIODevice`). Direct conceptual match; bridge = PIMPL holding a `juce::AudioDeviceManager`, forwarding sample-rate/channel config. Risk: JUCE GPL/AGPL vs Bobtk LGPL — keep bridge optional and license-gated (see below).
- **AudioProcessor** (setParameter/processBlock) ← `juce_audio_processors` + `juce_dsp`. `processBlock(float**, float**, int)` already mirrors `juce::AudioBuffer` channel layout; bridge can wrap channels without copies.
- **AudioDial / AudioFader / AudioScope / AudioAnalyzer** (widgets) ← `juce_gui_basics` components + `juce_audio_utils` analysers. Two options: (a) native Bobgui widgets fed by JUCE analyser data (preferred — keeps one widget tree), (b) embed `juce::Component` via native window handle (heavier, platform-specific). Choose (a).
- **ThreeDNode / EffectShader** ← `juce_opengl` contexts. Defer: GSK already provides GL rendering; JUCE adds little here.
- **Timeline** ← `juce_audio_basics` position info / `juce_video`. Weak match; keep native.
- **Visual pillar generally:** do NOT intermix `juce_gui_basics` windows with Bobgui windows in one tree (two event loops, two accessibility trees). Intermix at the *processor/data* level, render natively.

## Bridge status — honest correction
`CHANGELOG.md` [4.0.8] claims bridges at `bobgui/cpp/juce_audio_bridge.hpp` and `bobgui/cpp/upp_core_bridge.hpp`. **Neither file exists** (verified by recursive search). Status is therefore: analysis done, bridge not started. First spike (queued in `TODO.md`): header-only `juce_audio_bridge.hpp` behind `BOBGUI_WITH_JUCE` compile guard, wrapping device enumeration only, with a license note (JUCE is GPLv3/AGPLv3/commercial — linking changes distribution terms; the bridge must stay opt-in).

## Recommendation order
1. Media device bridge (enumeration + sample rate) — small, testable headless via `juce_audio_processors_headless`.
2. Processor `processBlock` forwarding with JUCE DSP chain.
3. Analyser data feeding native AudioScope/AudioAnalyzer widgets.
Defer OpenGL and GUI embedding indefinitely.
