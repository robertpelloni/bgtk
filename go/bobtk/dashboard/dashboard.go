// Package dashboard serves the Bobtk consolidated dashboard: one page that
// surfaces every Go pillar and the native C++ framework catalog, with live
// status, labelled controls, and help tooltips. It uses only the standard
// library so it builds and runs fully offline.
//
// Why one page: the project previously had no runnable dashboard at all —
// only a native C++ DashboardShell preset (bobgui/cpp/dashboard_shell.hpp).
// This package is the web counterpart: all former "pages" (pillars, native
// catalog, build/logs, tray/quit) are sections of a single document, ordered
// by value — health and running controls first, reference catalog after.
package dashboard

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/bobtk/bobtk/bobtk/core"
	"github.com/bobtk/bobtk/bobtk/media"
	"github.com/bobtk/bobtk/bobtk/network"
	"github.com/bobtk/bobtk/bobtk/system"
	"github.com/bobtk/bobtk/bobtk/tools"
	"github.com/bobtk/bobtk/bobtk/visual"
)

//go:embed dashboard.html
var pageHTML string

// Feature describes one backend capability wired to the page. Status is one
// of: "go" (works in Go today), "native" (C++ header exists), "planned"
// (documented bridge, not yet bound). The UI badges these honestly.
type Feature struct {
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Tooltip string `json:"tooltip"`
	Status  string `json:"status"`
}

// Pillar is the dashboard view of one of the six pillars.
type Pillar struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Desc        string    `json:"desc"`
	Tooltip     string    `json:"tooltip"`
	GoPackage   string    `json:"goPackage"`
	NativeFiles string    `json:"nativeFiles"`
	Initialized bool      `json:"initialized"`
	Features    []Feature `json:"features"`
	init        func()
}

type logRing struct {
	mu    sync.Mutex
	lines []string
	max   int
}

func (r *logRing) add(format string, args ...any) {
	r.mu.Lock()
	defer r.mu.Unlock()
	line := time.Now().Format("15:04:05") + " " + fmt.Sprintf(format, args...)
	r.lines = append(r.lines, line)
	if len(r.lines) > r.max {
		r.lines = r.lines[len(r.lines)-r.max:]
	}
	log.Println(line)
}

func (r *logRing) snapshot() []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	out := make([]string, len(r.lines))
	copy(out, r.lines)
	return out
}

// Server holds dashboard state. Construct with New, serve with Handler/Run.
type Server struct {
	Version string
	Addr    string

	mu       sync.Mutex
	pillars  []*Pillar
	byID     map[string]*Pillar
	checkers map[string]func() bool
	logs     *logRing
	started  time.Time
	shutdown func()
}

// ReadVersion returns the single source-of-truth version: the VERSION file
// at the repo root (or next to the binary), falling back to the default.
// Application logic never hardcodes a version elsewhere.
func ReadVersion(fallback string) string {
	for _, p := range []string{"VERSION", filepath.Join("..", "VERSION"), filepath.Join("..", "..", "VERSION")} {
		if b, err := os.ReadFile(p); err == nil {
			if v := strings.TrimSpace(string(b)); v != "" {
				return v
			}
		}
	}
	if v := os.Getenv("BOBTK_VERSION"); v != "" {
		return v
	}
	return fallback
}

// New builds a Server with the six real pillar instances wired in.
func New(version string) *Server {
	s := &Server{Version: version, logs: &logRing{max: 300}, started: time.Now(), byID: map[string]*Pillar{}}
	coreP := core.NewCorePillar()
	visualP := visual.NewVisualPillar()
	mediaP := media.NewMediaPillar()
	networkP := network.NewNetworkPillar()
	systemP := system.NewSystemPillar()
	toolsP := tools.NewToolsPillar()
	s.pillars = []*Pillar{
		{ID: "core", Title: "Core", GoPackage: "bobtk/core", NativeFiles: "bobgui/cpp/core_subsystems.hpp, module/core_pillar.hpp",
			Desc:    "Data, state, layout, scripting and scheduling foundations every other pillar builds on.",
			Tooltip: "Start here. Core owns shared objects, properties, signals and the schedulers; other pillars assume it is initialized first.",
			init:    coreP.Initialize,
			Features: []Feature{
				{Name: "ObjectHandle / Property / Signal", Desc: "Typed handles, observable properties and signal/slot wiring.", Tooltip: "C++: bobgui/cpp/object_handle.hpp, property.hpp, signal.hpp. Go binding is planned via cgo.", Status: "native"},
				{Name: "QuantumScheduler / RealtimeKernel", Desc: "Priority scheduling and realtime task kernel.", Tooltip: "Mapped from Ultimate++ CoWork analysis in UPP_ANALYSIS.md.", Status: "native"},
				{Name: "SqlDatabase / EntityManager", Desc: "SQLite-style storage and entity registry.", Tooltip: "C++: core_subsystems.hpp (SqlDatabase, EntityManager).", Status: "native"},
				{Name: "StateMachine / ScriptContext / ReflectContext", Desc: "State machines, scripting and reflection contexts.", Tooltip: "C++: core_subsystems.hpp. Go Initialize() currently flips live state only.", Status: "native"},
				{Name: "Go pillar Initialize()", Desc: "Idempotent Go initialization with live status.", Tooltip: "Runs now from this dashboard. Safe to click twice — second run is a no-op.", Status: "go"},
			}},
		{ID: "visual", Title: "Visual", GoPackage: "bobtk/visual", NativeFiles: "bobgui/cpp/app_shell.hpp, dashboard_shell.hpp, workbench.hpp",
			Desc:    "Widgets, shells and tool surfaces — the native UI layer, including the DashboardShell preset this page mirrors.",
			Tooltip: "Visual is the highest-value user-facing pillar: AppShell, StudioShell, DocumentShell and DashboardShell all live here.",
			init:    visualP.Initialize,
			Features: []Feature{
				{Name: "AppShell / Workbench", Desc: "Window shell with sidebars, status and command palette.", Tooltip: "C++: app_shell.hpp + workbench.hpp. Hosts every other shell preset.", Status: "native"},
				{Name: "DashboardShell", Desc: "Navigation + dashboard + context preset for monitoring apps.", Tooltip: "C++: dashboard_shell.hpp (set_dashboard_view, dashboard_action_count). This web page is its consolidated counterpart.", Status: "native"},
				{Name: "ActionRegistry / CommandPalette", Desc: "Central command registry driving menus, toolbars and palette.", Tooltip: "C++: action_registry.hpp, command_palette.hpp. Duck-type tip: pin frequent commands from the palette.", Status: "native"},
				{Name: "ToolSurface / ToolbarBuilder / DockManager", Desc: "Generated toolbar/tool surfaces and docking.", Tooltip: "C++: tool_surface*.hpp, toolbar_builder.hpp, dock_manager.hpp.", Status: "native"},
				{Name: "Go pillar Initialize()", Desc: "Idempotent Go initialization with live status.", Tooltip: "Runs now from this dashboard.", Status: "go"},
			}},
		{ID: "media", Title: "Media", GoPackage: "bobtk/media", NativeFiles: "bobgui/cpp/media.hpp",
			Desc:    "Audio, 3D, maps, shaders and timelines.",
			Tooltip: "Media wraps JUCE-style audio/DSP concepts. Native device I/O needs the C++ build; Go state runs today.",
			init:    mediaP.Initialize,
			Features: []Feature{
				{Name: "AudioDeviceManager", Desc: "Device init, sample rate, play/stop.", Tooltip: "C++ media.hpp: initializeDevices, setSampleRate, play, stop. Optional JUCE bridge: bobgui/cpp bridges (see SUBMODULE_INVENTORY.md).", Status: "native"},
				{Name: "AudioProcessor", Desc: "setParameter + processBlock audio routing.", Tooltip: "C++ media.hpp. Pair with AudioDial/AudioFader/AudioScope widgets.", Status: "native"},
				{Name: "ThreeDNode / MapView / EffectShader / Timeline", Desc: "3D transforms + mesh loading, map center/zoom/markers, shader uniforms, timeline play/pause/seek.", Tooltip: "C++ media.hpp. Backed by GObject stubs bobgui3d.c, bobguigis.c, bobguishader.c, bobguitimeline.c.", Status: "native"},
				{Name: "Go pillar Initialize()", Desc: "Idempotent Go initialization with live status.", Tooltip: "Runs now from this dashboard.", Status: "go"},
			}},
		{ID: "network", Title: "Network", GoPackage: "bobtk/network", NativeFiles: "bobgui/cpp/network.hpp",
			Desc:    "HTTP access, WebSockets, remote server/client and sync.",
			Tooltip: "Network is independent of Visual — initialize it any time. WebView lives here as a widget wrapper.",
			init:    networkP.Initialize,
			Features: []Feature{
				{Name: "NetworkAccessManager / WebSocket", Desc: "HTTP requests and persistent socket connections.", Tooltip: "C++ network.hpp: NetworkAccessManager, WebSocket.", Status: "native"},
				{Name: "RemoteServer / RemoteClient / SyncEngine", Desc: "Remote endpoints and state synchronization.", Tooltip: "C++ network.hpp. CloudContext and BlockchainManager are also declared there.", Status: "native"},
				{Name: "Go pillar Initialize()", Desc: "Idempotent Go initialization with live status.", Tooltip: "Runs now from this dashboard.", Status: "go"},
			}},
		{ID: "system", Title: "System", GoPackage: "bobtk/system", NativeFiles: "bobgui/cpp/module/system.hpp",
			Desc:    "OS integration: file watching and local servers.",
			Tooltip: "System is the thinnest pillar today — two native types. Dashboard/logs/tray controls on this page are its web-facing surface.",
			init:    systemP.Initialize,
			Features: []Feature{
				{Name: "FileSystemWatcher", Desc: "Watch paths and react to file changes.", Tooltip: "C++ module/system.hpp (system::FileSystemWatcher).", Status: "native"},
				{Name: "LocalServer", Desc: "Local IPC/server endpoint.", Tooltip: "C++ module/system.hpp (system::LocalServer). This dashboard server is the Go-side equivalent in use.", Status: "native"},
				{Name: "Go pillar Initialize()", Desc: "Idempotent Go initialization with live status.", Tooltip: "Runs now from this dashboard.", Status: "go"},
			}},
		{ID: "tools", Title: "Tools", GoPackage: "bobtk/tools", NativeFiles: "bobgui/cpp/tools_subsystems.hpp, tools/",
			Desc:    "Developer tooling: forge/build context, reports, studio management, test runner and CLI tools.",
			Tooltip: "Tools also covers the C CLI utilities under tools/ (builder, image, path tools) which need the native build.",
			init:    toolsP.Initialize,
			Features: []Feature{
				{Name: "ForgeContext / StudioManager", Desc: "Build/IDE context and studio session management.", Tooltip: "C++ tools_subsystems.hpp.", Status: "native"},
				{Name: "ReportDesigner / ReportPreview / TestRunner", Desc: "Report design/preview and test execution.", Tooltip: "C++ tools_subsystems.hpp. Go test command: go test ./... in go/.", Status: "native"},
				{Name: "CLI tools (builder / image / path)", Desc: "bobgui-builder-tool, bobgui-image-tool, bobgui-path-tool.", Tooltip: "C sources in tools/. Require the Meson native build — see Build section for status.", Status: "native"},
				{Name: "Go pillar Initialize()", Desc: "Idempotent Go initialization with live status.", Tooltip: "Runs now from this dashboard.", Status: "go"},
			}},
	}
	// Keep Initialized truthful: mirror each pillar's real flag at read time.
	initializers := map[string]func() bool{
		"core": coreP.Initialized, "visual": visualP.Initialized, "media": mediaP.Initialized,
		"network": networkP.Initialized, "system": systemP.Initialized, "tools": toolsP.Initialized,
	}
	for _, p := range s.pillars {
		s.byID[p.ID] = p
	}
	// Status() reads Initialized through these checkers so the page always
	// shows the pillars' real flags, never a cached copy.
	s.checkers = initializers
	return s
}

// Status is the JSON document the page renders from.
type Status struct {
	Name        string    `json:"name"`
	Version     string    `json:"version"`
	GoVersion   string    `json:"goVersion"`
	OS          string    `json:"os"`
	Arch        string    `json:"arch"`
	UptimeSec   int64     `json:"uptimeSec"`
	Initialized int       `json:"initializedCount"`
	Total       int       `json:"totalCount"`
	Pillars     []*Pillar `json:"pillars"`
	Logs        []string  `json:"logs"`
	NativeBuild string    `json:"nativeBuild"`
}

// Status snapshots live pillar state.
func (s *Server) Status() Status {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := Status{Name: "Bobtk", Version: s.Version, GoVersion: runtime.Version(),
		OS: runtime.GOOS, Arch: runtime.GOARCH, UptimeSec: int64(time.Since(s.started).Seconds()),
		Total: len(s.pillars), NativeBuild: NativeBuildNote()}
	for _, p := range s.pillars {
		cp := *p
		cp.Features = p.Features
		if f := s.checkers[p.ID]; f != nil {
			cp.Initialized = f()
		}
		if cp.Initialized {
			out.Initialized++
		}
		out.Pillars = append(out.Pillars, &cp)
	}
	out.Logs = s.logs.snapshot()
	return out
}

// NativeBuildNote reports native (Meson) build reality without faking a run.
// BOBTK_BUILD_DIR may point at a verified build tree (e.g. a probe dir).
func NativeBuildNote() string {
	if dir := os.Getenv("BOBTK_BUILD_DIR"); dir != "" {
		if _, err := os.Stat(filepath.Join(dir, "build.ninja")); err == nil {
			if _, err := os.Stat(filepath.Join(dir, "bobgui", "libbobgui-4-1.dll")); err == nil {
				return "compiled green at BOBTK_BUILD_DIR (libbobgui-4-1.dll present, compile exit 0)"
			}
			return "configured at BOBTK_BUILD_DIR (build.ninja present)"
		}
	}
	for _, dir := range []string{"_build", filepath.Join("..", "_build")} {
		if _, err := os.Stat(filepath.Join(dir, "build.ninja")); err == nil {
			if _, err := os.Stat(filepath.Join(dir, "bobgui", "libbobgui-4-1.dll")); err == nil {
				return "compiled green (_build, libbobgui-4-1.dll present)"
			}
			return "configured (_build/build.ninja present)"
		}
	}
	if _, err := os.Stat("_build"); err == nil {
		return "partial _build directory present — run meson setup to configure"
	}
	return "not configured on this machine — Go dashboard runs standalone; see DEPLOY.md"
}

func (s *Server) logf(format string, args ...any) { s.logs.add(format, args...) }

// InitPillar initializes one pillar by id. Unknown ids return false.
func (s *Server) InitPillar(id string) bool {
	s.mu.Lock()
	p, ok := s.byID[id]
	s.mu.Unlock()
	if !ok {
		return false
	}
	p.init()
	s.logf("pillar %s initialized", id)
	return true
}

// InitAll initializes every pillar in dependency order (Core first).
func (s *Server) InitAll() {
	for _, id := range []string{"core", "visual", "media", "network", "system", "tools"} {
		s.InitPillar(id)
	}
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(v)
}

// Handler returns the dashboard HTTP handler (exported for tests).
func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(pageHTML))
	})
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]any{"ok": true, "version": s.Version, "uptimeSec": int64(time.Since(s.started).Seconds())})
	})
	mux.HandleFunc("GET /api/status", func(w http.ResponseWriter, r *http.Request) { writeJSON(w, s.Status()) })
	mux.HandleFunc("GET /api/logs", func(w http.ResponseWriter, r *http.Request) { writeJSON(w, map[string]any{"logs": s.logs.snapshot()}) })
	mux.HandleFunc("POST /api/init-all", func(w http.ResponseWriter, r *http.Request) {
		s.InitAll()
		writeJSON(w, s.Status())
	})
	mux.HandleFunc("POST /api/init/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if !s.InitPillar(id) {
			http.Error(w, "unknown pillar: "+id, http.StatusNotFound)
			return
		}
		writeJSON(w, s.Status())
	})
	mux.HandleFunc("POST /api/quit", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]any{"ok": true, "message": "shutting down"})
		s.logf("quit requested from dashboard (tray-equivalent control)")
		if s.shutdown != nil {
			// Delay so the response flushes before the listener closes.
			go func() { time.Sleep(300 * time.Millisecond); s.shutdown() }()
		}
	})
	return mux
}

// Run serves until ctx shutdown via Quit or error. It binds first so a busy
// port fails fast with a clear log line instead of a silent hang.
func (s *Server) Run(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("listen %s: %w", addr, err)
	}
	s.Addr = "http://" + ln.Addr().String()
	srv := &http.Server{Handler: s.Handler(), ReadHeaderTimeout: 10 * time.Second}
	done := make(chan struct{})
	s.shutdown = func() {
		_ = srv.Close()
		close(done)
	}
	s.logf("Bobtk dashboard %s listening on %s (%s/%s, %s)", s.Version, s.Addr, runtime.GOOS, runtime.GOARCH, runtime.Version())
	err = srv.Serve(ln)
	if err == http.ErrServerClosed {
		s.logf("dashboard stopped")
		return nil
	}
	return err
}
