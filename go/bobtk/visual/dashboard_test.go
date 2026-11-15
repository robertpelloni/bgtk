package visual

import (
	"strings"
	"testing"

	"github.com/bobtk/bobtk/bobtk/system"
)

func TestDashboardUISinglePageRender(t *testing.T) {
	sm := system.NewServerManager()
	tray := system.NewSystemTray("test-tray", sm)
	pillar := NewVisualPillar()
	pillar.Initialize()

	dashboard := pillar.SetupConsolidatedDashboard(sm, tray)
	if dashboard == nil {
		t.Fatalf("expected non-nil dashboard")
	}

	rendered := dashboard.RenderSinglePage()
	if !strings.Contains(rendered, "Single-Page Consolidated Architecture Control Surface") {
		t.Errorf("rendered output missing expected subtitle banner")
	}

	if !strings.Contains(rendered, "SYSTEM TRAY & SERVER SUPERVISOR") {
		t.Errorf("rendered output missing System Tray section")
	}

	if !strings.Contains(rendered, "CRITICAL [TOP PRIORITY]") {
		t.Errorf("rendered output missing high-priority cards")
	}
}

func TestDashboardTelemetryAndLogStreaming(t *testing.T) {
	sm := system.NewServerManager()
	sm.RegisterServer("srv-1", "Test HTTP Server", "Handles HTTP", 8080, nil)
	tray := system.NewSystemTray("tray-icon", sm)

	d := NewDashboardUI(sm, tray)

	// Test real-time log streaming subscription
	var streamedLogs []string
	d.SubscribeLogs(func(line string) {
		streamedLogs = append(streamedLogs, line)
	})

	d.Log("Event 1: System initiated")
	d.Log("Event 2: Gateway connected")

	if len(streamedLogs) != 2 {
		t.Fatalf("expected 2 streamed logs, got %d", len(streamedLogs))
	}
	if !strings.Contains(streamedLogs[0], "Event 1: System initiated") {
		t.Errorf("unexpected log content: %s", streamedLogs[0])
	}

	// Test telemetry updates
	d.UpdateTelemetry(ServerTelemetry{
		ServerID:  "srv-1",
		CPUUsage:  14.5,
		MemoryMB:  42.8,
		UptimeSec: 3600,
		Requests:  1250,
	})

	rendered := d.RenderSinglePage()
	if !strings.Contains(rendered, "CPU: 14.5%") || !strings.Contains(rendered, "Mem: 42.8 MB") {
		t.Errorf("rendered output does not contain expected telemetry info: %s", rendered)
	}
}

func TestDashboardShortcutKeyDispatcher(t *testing.T) {
	sm := system.NewServerManager()
	sm.RegisterServer("srv-alpha", "Test Server Alpha", "Running", 5555, nil)
	tray := system.NewSystemTray("tray-icon", sm)

	d := NewDashboardUI(sm, tray)

	triggeredQuit := false
	triggeredRefresh := false

	d.RegisterShortcut("Ctrl+Q", "Quit Dashboard and Servers", "Lifecycle", func() (string, error) {
		triggeredQuit = true
		return "quitting", nil
	})

	d.RegisterShortcut("Ctrl+R", "Refresh Visual Surfaces", "Navigation", func() (string, error) {
		triggeredRefresh = true
		return "refreshed", nil
	})

	// Dispatch Ctrl+R
	res, err := d.HandleKeyPress("ctrl+r")
	if err != nil {
		t.Fatalf("unexpected error executing shortcut: %v", err)
	}
	if res != "refreshed" || !triggeredRefresh {
		t.Errorf("expected refresh shortcut to be triggered")
	}

	// Dispatch Ctrl+Q
	res, err = d.HandleKeyPress("CTRL+Q")
	if err != nil {
		t.Fatalf("unexpected error executing shortcut: %v", err)
	}
	if res != "quitting" || !triggeredQuit {
		t.Errorf("expected quit shortcut to be triggered")
	}

	// Dispatch unbound key
	_, err = d.HandleKeyPress("Ctrl+Alt+Delete")
	if err == nil {
		t.Errorf("expected error for unbound shortcut")
	}

	rendered := d.RenderSinglePage()
	if !strings.Contains(rendered, "[GLOBAL ACCELERATORS & KEYBOARD SHORTCUTS]") {
		t.Errorf("missing keyboard shortcuts section in dashboard render: %s", rendered)
	}
	if !strings.Contains(rendered, "CTRL+Q") || !strings.Contains(rendered, "CTRL+R") {
		t.Errorf("missing registered shortcuts in rendered output")
	}
}
