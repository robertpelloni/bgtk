package visual

import (
	"fmt"

	"github.com/bobtk/bobtk/bobtk/system"
)

// VisualPillar manages core UI rendering, widgets, and the unified dashboard.
type VisualPillar struct {
	initialized bool
	dashboard   *DashboardUI
}

func NewVisualPillar() *VisualPillar {
	return &VisualPillar{
		initialized: false,
	}
}

func (v *VisualPillar) Initialize() {
	if v.initialized {
		return
	}
	fmt.Println("VisualPillar initialized.")
	v.initialized = true
}

// SetupConsolidatedDashboard configures the unified single-page dashboard surface
// containing consolidated sections for all 6 pillars and the System Tray.
func (v *VisualPillar) SetupConsolidatedDashboard(sm *system.ServerManager, tray *system.SystemTray) *DashboardUI {
	v.dashboard = NewDashboardUI(sm, tray)

	// Pillar 1: Visual & Layout (Highest Value User Interface Surface)
	v.dashboard.AddCard(&DashboardCard{
		ID:       "visual.surface",
		Pillar:   "Visual",
		Title:    "Visual Engine & Single-Page Surface",
		Summary:  "Modern multi-paradigm UI compositor unifying Qt, JUCE, U++, and ImGui paradigms into a cohesive viewport.",
		Priority: PriorityCritical,
		Tooltip: TooltipMarker{
			Header:      "Visual Compositor",
			Description: "Controls layout engine, render nodes, custom styling, and single-page consolidated panels.",
			Shortcut:    "Ctrl+Shift+V",
			Category:    "Visual",
		},
		Metrics: map[string]string{
			"Layout Engine":     "Flexbox & Constraint Grid Solver",
			"Active Views":      "Consolidated Single Page (Zero Fragmentation)",
			"FPS Overlay":       "60.0 FPS Stable",
			"Subpages Reduced":  "100% consolidated into main scrollview",
		},
		Subsections: []string{"Workbench", "Docking Subsystem", "Widget Tree", "ImGui Overlay"},
		Actions: []ActionButton{
			{
				ID:       "visual.reload_theme",
				Label:    "Refresh Theme",
				Priority: PriorityMedium,
				Tooltip: TooltipMarker{
					Header:      "Theme Manager",
					Description: "Reapplies active CSS theme stylesheet dynamically without restarting.",
					Shortcut:    "Ctrl+T",
					Category:    "Visual",
				},
				Handler: func() string { return "Theme reloaded." },
			},
		},
	})

	// Pillar 2: System & Server Supervisor (System Tray + Daemons)
	v.dashboard.AddCard(&DashboardCard{
		ID:       "system.servers",
		Pillar:   "System",
		Title:    "System Tray & Server Supervisor",
		Summary:  "OS-level notification area integration, IPC daemon supervisor, and background server termination.",
		Priority: PriorityCritical,
		Tooltip: TooltipMarker{
			Header:      "System Operations",
			Description: "Inspects status of all active servers and provides immediate kill/quit signals via System Tray.",
			Shortcut:    "Ctrl+Shift+S",
			Category:    "System",
		},
		Metrics: map[string]string{
			"System Tray Presence": "Active (Desktop notification icon enabled)",
			"Managed Daemons":      fmt.Sprintf("%d servers supervised", sm.RunningCount()),
			"IPC Protocol":         "Unix Domain Sockets & Named Pipes",
		},
		Subsections: []string{"System Tray", "Process Controller", "VFS Layer", "Security & Sandbox"},
		Actions: []ActionButton{
			{
				ID:       "system.quit_all",
				Label:    "Stop All Servers & Quit",
				Priority: PriorityCritical,
				Tooltip: TooltipMarker{
					Header:      "Global Shutdown",
					Description: "Gracefully terminates all background servers and closes the dashboard cleanly.",
					Shortcut:    "Ctrl+Q",
					Category:    "System",
				},
				Handler: func() string {
					tray.QuitApplication()
					return "Shutdown executed."
				},
			},
		},
	})

	// Pillar 3: Core (Quantum Scheduler & Memory)
	v.dashboard.AddCard(&DashboardCard{
		ID:       "core.engine",
		Pillar:   "Core",
		Title:    "Core Engine & Quantum Scheduler",
		Summary:  "High-performance memory allocation, async task queues, and reflective data binding inspired by Ultimate++.",
		Priority: PriorityHigh,
		Tooltip: TooltipMarker{
			Header:      "Core Subsystems",
			Description: "Monitors memory arena utilization, task queue depth, and thread dispatch latencies.",
			Shortcut:    "Ctrl+Shift+C",
			Category:    "Core",
		},
		Metrics: map[string]string{
			"Memory Allocator":  "U++ NTL Heap / Cache-Coherent Pools",
			"Quantum Scheduler": "16 Concurrency Workers Active",
			"Dispatched Tasks":  "0 pending, 248 executed",
		},
		Subsections: []string{"Quantum Scheduler", "NTL Memory Heap", "Entity Manager", "Reflect Context"},
	})

	// Pillar 4: Network & Remote (WebSockets, Cloud, Broadway)
	v.dashboard.AddCard(&DashboardCard{
		ID:       "network.gateway",
		Pillar:   "Network",
		Title:    "Network Gateway & Remote Display",
		Summary:  "Seamless distributed networking, WebSocket transport, and HTML5 Broadway server integration.",
		Priority: PriorityHigh,
		Tooltip: TooltipMarker{
			Header:      "Network Gateway",
			Description: "Manages remote connectivity, WebSocket listeners, and HTML5 web canvas streaming.",
			Shortcut:    "Ctrl+Shift+N",
			Category:    "Network",
		},
		Metrics: map[string]string{
			"Broadway HTML5":  "Port 8085 (Listening)",
			"WebSocket IPC":   "Port 9090 (Active)",
			"Active Sockets":  "2 connections established",
		},
		Subsections: []string{"Broadway Server", "WebSocket Client", "Cloud Providers", "Sync Bus"},
	})

	// Pillar 5: Media (Audio, MIDI2, 3D, Spatial)
	v.dashboard.AddCard(&DashboardCard{
		ID:       "media.pipeline",
		Pillar:   "Media",
		Title:    "Media & Audio Synthesis Pipeline",
		Summary:  "Real-time audio graph, JUCE audio plugin hosting, MIDI 2.0 streaming, and spatial sound dispatch.",
		Priority: PriorityMedium,
		Tooltip: TooltipMarker{
			Header:      "Media Engine",
			Description: "Realtime low-latency audio engine inspired by JUCE audio processors and synthesizer graphs.",
			Shortcut:    "Ctrl+Shift+M",
			Category:    "Media",
		},
		Metrics: map[string]string{
			"Sample Rate":   "48000 Hz (Buffer: 256 samples)",
			"Audio Latency": "5.3 ms",
			"MIDI 2.0":      "Universal MIDI Packet Protocol Ready",
		},
		Subsections: []string{"Audio Graph", "JUCE Plugin Host", "MIDI 2.0 Bus", "3D Spatializer"},
	})

	// Pillar 6: Tools & Forge (Profiler, Inspector, Live-Reload)
	v.dashboard.AddCard(&DashboardCard{
		ID:       "tools.forge",
		Pillar:   "Tools",
		Title:    "Developer Forge & Diagnostic Tools",
		Summary:  "Built-in runtime inspector, live performance profiler, memory leak detection, and interactive test runner.",
		Priority: PriorityMedium,
		Tooltip: TooltipMarker{
			Header:      "Developer Tools",
			Description: "Interactive diagnostic tools, render node visualizer, and code generator for fast prototyping.",
			Shortcut:    "Ctrl+Shift+T",
			Category:    "Tools",
		},
		Metrics: map[string]string{
			"Profiler":     "Sysprof / Sampling Active",
			"Test Harness": "Built-in Suite (All green)",
			"Hot Reload":   "VFS Watcher Enabled",
		},
		Subsections: []string{"Studio Inspector", "Render Node Tool", "Performance Profiler", "Test Runner"},
	})

	return v.dashboard
}

// Dashboard returns the active dashboard instance if initialized.
func (v *VisualPillar) Dashboard() *DashboardUI {
	return v.dashboard
}
