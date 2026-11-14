# TODO

Granular queue. The 15-minute autonomous loop works top-down. Checkboxes only move on verified evidence.

## Native build repair (highest value)
- [x] Untangle spliced `meson.build`: restored clean bobgui file from git history (`ce1de57f61`, 932 lines); spliced file archived at `scripts/archive/meson.build.spliced.txt`; GTK3 `meson_options.txt` archived to `scripts/archive/meson_options.gtk3.txt`.
- [x] Truncate orphan GTK3 tails in 8 wraps (backups in `scripts/archive/wraps-spliced/`); add missing `[provide]` mappings: libffi, zlib, intl (proxy-libintl); pin pango `main` → `1.56.4` (glib 2.84 era); switch pixman centricular fork → upstream `pixman-0.46.4` (exports `idep_pixman`); add wraps: libsoup-3.0 (3.6.5), sqlite3, nghttp2, libpsl (WrapDB).
- [x] `meson setup` probe **passes** (2026-11-13): `meson setup $env:TEMP\bgtk-probe-build . -Dmedia-gstreamer=disabled -Dintrospection=disabled -Dvulkan=disabled` exits 0. GStreamer disabled due to its Windows win-pkgconfig override bug; Vulkan disabled (no SDK on this laptop).
- [ ] `meson compile` in progress (BelowNormal, log `logs/native-compile6.log`): fixed in order — (1) sysprof Unix-only headers → `-Dsysprof=disabled -Dglib:sysprof=disabled`; (2) centricular libffi multiple-definition (GCC 15) → WrapDB libffi 3.8.0; (3) cairo 1.18.2 vs MinGW `dcommon.h` → cairo 1.18.4; (4) unresolved merge conflict in `bobgui/theme/Default/Default-light.scss` → kept bobgui (HEAD) side; (5) `gnome.mkenums_simple` Windows quoting failure in `gdk/meson.build` → template-based `gnome.mkenums` with new `gdk/gdkenumtypes.{h,c}.template` (mirrors bobgui/gsk pattern). Keep fixing failures in order until compile exits 0. Also fixed: (6) nested `<gresource>` in `gdk/gen-gdk-gresources-xml.py` (spliced second document) → single bobgui gresource; (7) spliced `gdk/gdkconfig.h.meson` (unterminated generated header) → restored from `ce1de57f61`; (8) spliced C sources across gdk/gsk/bobgui (GTK3 headers like `gdkmain.h`/`gdkwindow.h` breaking `gdkversionmacros.h` includes) → restored 150 files from `ce1de57f61` (175 post-clean additions kept; spliced copies in `scripts/archive/meson-spliced/`). Note: restoration reverts the gdk template-enum fix — re-apply it after any future restore (see CHANGELOG 1.2.1).
- [ ] Re-enable media-gstreamer (fix win-pkgconfig override) and vulkan (needs Vulkan SDK) after core compile is green.
- [ ] Fix submodule state: `submodules/juce` dirty pin (25531a05-dirty vs recorded d60b7a78); `submodules/ultimatepp` present but untracked. Record intended pins, commit cleanly.

## Dashboard & UI
- [x] Consolidated single-page dashboard with tooltips/labels (1.1.0).
- [ ] Native tray icon in C++ (or Go systray lib when network allows) to replace `scripts/tray.ps1`.
- [ ] Add per-feature dashboard actions as bridges land (e.g. Media play/stop once cgo exists).

## Bridges & parity
- [x] Analyze JUCE UI widgets for intermixing → `docs/JUCE_ANALYSIS.md` (Media-first; note CHANGELOG [4.0.8] bridge files do not exist).
- [x] Research missing Qt6/Qt7 features vs C++ baseline → `docs/QT_PARITY_MATRIX.md` (top gaps: declarative/QML layer, charting, SQL model/view).
- [ ] First cgo bridge spike (Core or Media) behind a build tag, off by default.

## Housekeeping
- [ ] Decide fate of root `check-version.py` / `plan_step_complete.py` (archive copies exist in `scripts/archive/`; root copies kept for now).
- [ ] Remove stale `HANDOFF.md` claims as milestones land (XKB bug already fixed in tree; meson wrap story superseded by splice finding).
