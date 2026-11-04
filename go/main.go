package main

import (
	"fmt"

	"github.com/bobtk/bobtk/bobtk/core"
	"github.com/bobtk/bobtk/bobtk/media"
	"github.com/bobtk/bobtk/bobtk/network"
	"github.com/bobtk/bobtk/bobtk/system"
	"github.com/bobtk/bobtk/bobtk/tools"
	"github.com/bobtk/bobtk/bobtk/visual"
)

func main() {
	fmt.Println("================================================================================")
	fmt.Println("             BOBTK UNIFIED RUNTIME & OPERATIONS DASHBOARD")
	fmt.Println("================================================================================")
	fmt.Println("Initializing 6-Pillar Framework...")

	corePillar := core.NewCorePillar()
	corePillar.Initialize()

	mediaPillar := media.NewMediaPillar()
	mediaPillar.Initialize()

	networkPillar := network.NewNetworkPillar()
	networkPillar.Initialize()

	systemPillar := system.NewSystemPillar()
	systemPillar.Initialize()

	toolsPillar := tools.NewToolsPillar()
	toolsPillar.Initialize()

	visualPillar := visual.NewVisualPillar()
	visualPillar.Initialize()

	fmt.Println("Bobtk Go Port Initialization Complete.")
	fmt.Println()

	// Demonstrate Core Pillar Quantum Scheduler & Dynamic Reflection
	scheduler := corePillar.QuantumScheduler()
	reflectCtx := corePillar.ReflectContext()
	entityMgr := corePillar.EntityManager()

	reflectCtx.SetProperty("system.cluster_state", "Healthy-Online")
	reflectCtx.SetProperty("ui.active_theme", "Bgtk-Modern-Dark")

	// Create root UI entity
	rootWin := entityMgr.CreateEntity("AppWindow")
	_ = entityMgr.AttachComponent(rootWin.ID, "role", "PrimaryDashboard")

	// Dispatch worker tasks to verify scheduler throughput
	for i := 0; i < 32; i++ {
		scheduler.Dispatch(func() {
			// Simulating background telemetry acquisition
		})
	}
	scheduler.Wait()

	// Initialize the Server Supervisor & System Tray
	serverManager := system.NewServerManager()

	// Register background servers into the supervisor
	serverManager.RegisterServer("broadway.server", "Broadway Server", "HTML5 Broadway Display Gateway", 8085, func() error {
		fmt.Println("[Broadway Server] Received graceful shutdown signal. Socket closed.")
		return nil
	})
	serverManager.RegisterServer("ipc.daemon", "Local IPC Daemon", "High-throughput domain socket messaging daemon", 9090, func() error {
		fmt.Println("[IPC Daemon] Terminating local client connections and unlinking socket.")
		return nil
	})
	serverManager.RegisterServer("media.engine", "Realtime Audio Server", "Low-latency JUCE-compatible audio synthesis engine", 7777, func() error {
		fmt.Println("[Realtime Audio] Stopping audio stream and releasing soundcard hardware buffers.")
		return nil
	})

	// Configure System Tray
	tray := system.NewSystemTray("utilities-system-monitor-symbolic", serverManager)
	tray.SetTooltip("bgtk Operations Center - 3 Active Servers Monitored")

	tray.AddMenuItem("tray.status", "System Status: Online", "All 6 pillars healthy and operational", nil)
	tray.AddMenuItem("tray.stop_servers", "Stop Background Servers", "Gracefully shuts down all 3 registered background servers", func() {
		serverManager.StopAllServers()
	})
	tray.AddSeparator()
	tray.AddMenuItem("tray.quit", "Quit Dashboard and All Servers", "Halts all daemons and exits cleanly", func() {
		tray.QuitApplication()
	})

	// Demonstrate Network Gateway and Media Audio DSP Pipeline
	broadway := networkPillar.Gateway()
	_ = broadway.RegisterClient("demo-client-1", "127.0.0.1")

	audioGraph := mediaPillar.AudioGraph()
	audioGraph.AddProcessor(&media.GainProcessor{Gain: 0.85})
	audioBuf := media.NewAudioBuffer(2, 256)
	audioGraph.Process(audioBuf)

	midiBus := mediaPillar.MidiBus()
	midiBus.Broadcast(media.NewNoteOn(0, 1, 60, 0x7FFF))

	// Demonstrate JUCE Component Tree & Timer Integration
	rootComp := visual.NewComponent("RootAudioHost")
	rootComp.SetBounds(0, 0, 1024, 768)
	meterComp := visual.NewComponent("VUMeterDisplay")
	meterComp.SetBounds(20, 20, 200, 150)
	rootComp.AddAndMakeVisible(meterComp)
	rootComp.Repaint()

	// High-precision telemetry timer
	telemetryTimer := core.NewTimer(func() {
		// Telemetry tick
	})
	telemetryTimer.StartTimer(50)
	defer telemetryTimer.StopTimer()

	// Demonstrate Qt6 Layout & Widget Parity
	vbox := visual.NewVBoxLayout(8, 4)
	titleLabel := visual.NewLabel("Qt6 Parity Dashboard Header")
	actionBtn := visual.NewButton("Apply Layout")
	progressBar := visual.NewProgressBar()
	progressBar.SetValue(80)

	vbox.AddItem(titleLabel, 1)
	vbox.AddItem(actionBtn, 1)
	vbox.AddItem(progressBar, 2)
	vbox.Layout(0, 0, 400, 120)

	// Wire up the consolidated single-page dashboard
	dashboard := visualPillar.SetupConsolidatedDashboard(serverManager, tray)
	dashboard.Log("Consolidated single-page dashboard initialized successfully.")
	dashboard.Log("Qt6 Parity Layout Engine & Widgets initialized (QHBoxLayout, QVBoxLayout, QGridLayout, QStackedLayout).")
	dashboard.Log(fmt.Sprintf("JUCE Component hierarchy nested (%d children) and precision Timer active (%d ms interval).", rootComp.NumChildComponents(), telemetryTimer.GetTimerInterval()))
	dashboard.Log(fmt.Sprintf("Network Broadway Gateway active on port %d with %d connected client(s).", broadway.Port(), broadway.ClientCount()))
	dashboard.Log(fmt.Sprintf("Media Audio DSP pipeline running with %d active processor(s) and MIDI 2.0 bus ready.", audioGraph.ProcessorCount()))
	dashboard.Log(fmt.Sprintf("System tray active with %d background servers registered.", serverManager.RunningCount()))

	// Render the unified single-page layout
	fmt.Print(dashboard.RenderSinglePage())

	fmt.Println("[DEMONSTRATING SYSTEM TRAY SERVER QUIT ACTION]")
	// Test the tray's ability to cleanly stop all servers and quit
	tray.QuitApplication()
	fmt.Printf("Servers active after quit command: %d\n", serverManager.RunningCount())
	fmt.Println("================================================================================")
}
