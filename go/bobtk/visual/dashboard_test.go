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
