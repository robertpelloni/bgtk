package visual

import (
	"fmt"
	"strings"
	"time"

	"github.com/bobtk/bobtk/bobtk/system"
)

// PriorityLevel determines visual hierarchy on the dashboard.
type PriorityLevel int

const (
	PriorityNormal PriorityLevel = iota
	PriorityMedium
	PriorityHigh
	PriorityCritical
)

// TooltipMarker provides rich metadata and guidance for any UI element.
type TooltipMarker struct {
	Header      string
	Description string
	Shortcut    string
	Category    string
}

// ActionButton represents an interactive workflow action on a dashboard card.
type ActionButton struct {
	ID       string
	Label    string
	Tooltip  TooltipMarker
	Priority PriorityLevel
	Handler  func() string
}

// DashboardCard represents a consolidated visual block for a core subsystem or pillar.
type DashboardCard struct {
	ID          string
	Pillar      string
	Title       string
	Summary     string
	Tooltip     TooltipMarker
	Priority    PriorityLevel
	Metrics     map[string]string
	Actions     []ActionButton
	Subsections []string
}

// DashboardUI provides a unified, consolidated single-page dashboard for Bobtk.
// It combines all features and functions across the 6 pillars onto one singular surface,
// grouping parallel concepts and eliminating fragmented subpages.
type DashboardUI struct {
	Title         string
	Subtitle      string
	Cards         []*DashboardCard
	SystemTray    *system.SystemTray
	ServerManager *system.ServerManager
	Logs          []string
}

// NewDashboardUI instantiates a consolidated single-page dashboard.
func NewDashboardUI(sm *system.ServerManager, tray *system.SystemTray) *DashboardUI {
	if sm == nil {
		sm = system.NewServerManager()
	}
	if tray == nil {
		tray = system.NewSystemTray("utilities-system-monitor", sm)
	}

	return &DashboardUI{
		Title:         "bgtk Unified Operations & Subsystem Dashboard",
		Subtitle:      "Single-Page Consolidated Architecture Control Surface",
		Cards:         make([]*DashboardCard, 0),
		SystemTray:    tray,
		ServerManager: sm,
		Logs:          make([]string, 0),
	}
}

// AddCard registers a consolidated pillar card.
func (d *DashboardUI) AddCard(card *DashboardCard) {
	d.Cards = append(d.Cards, card)
}

// Log records an event message in the dashboard runtime log.
func (d *DashboardUI) Log(msg string) {
	entry := fmt.Sprintf("[%s] %s", time.Now().Format("15:04:05"), msg)
	d.Logs = append(d.Logs, entry)
	fmt.Println(entry)
}

// RenderSinglePage renders the unified dashboard state to a cleanly formatted text view.
func (d *DashboardUI) RenderSinglePage() string {
	var sb strings.Builder

	sb.WriteString("================================================================================\n")
	sb.WriteString(fmt.Sprintf(" %s\n", strings.ToUpper(d.Title)))
	sb.WriteString(fmt.Sprintf(" %s\n", d.Subtitle))
	sb.WriteString("================================================================================\n\n")

	// System Tray & Global Server Summary Header
	sb.WriteString("[SYSTEM TRAY & SERVER SUPERVISOR]\n")
	sb.WriteString(fmt.Sprintf("  • Status Icon: %s (Tooltip: %q, Visible: %v)\n",
		"Active", d.SystemTray.GetTooltip(), d.SystemTray.IsVisible()))
	sb.WriteString(fmt.Sprintf("  • Managed Servers: %d total, %d active\n",
		len(d.ServerManager.GetServers()), d.ServerManager.RunningCount()))

	for _, s := range d.ServerManager.GetServers() {
		sb.WriteString(fmt.Sprintf("    [%s] %s (Port: %d) - Status: %s - %s\n",
			s.ID, s.Name, s.Port, s.Status, s.Description))
	}
	sb.WriteString("\n")

	// Consolidated Pillar Cards sorted/presented by high value
	sb.WriteString("[HIGH-VALUE ARCHITECTURE CONSOLIDATION (6 PILLARS - UNIFIED SINGLE-PAGE)]\n")
	sb.WriteString("--------------------------------------------------------------------------------\n")

	for _, card := range d.Cards {
		prioStr := "NORMAL"
		if card.Priority == PriorityCritical {
			prioStr = "CRITICAL [TOP PRIORITY]"
		} else if card.Priority == PriorityHigh {
			prioStr = "HIGH VALUE"
		}

		sb.WriteString(fmt.Sprintf("✦ [%s PILLAR] %s (%s)\n", strings.ToUpper(card.Pillar), card.Title, prioStr))
		sb.WriteString(fmt.Sprintf("  Overview: %s\n", card.Summary))
		sb.WriteString(fmt.Sprintf("  ℹ Tooltip Guidance: [%s] %s (Category: %s, Shortcut: %s)\n",
			card.Tooltip.Header, card.Tooltip.Description, card.Tooltip.Category, card.Tooltip.Shortcut))

		if len(card.Metrics) > 0 {
			sb.WriteString("  Metrics / State:\n")
			for k, v := range card.Metrics {
				sb.WriteString(fmt.Sprintf("    - %s: %s\n", k, v))
			}
		}

		if len(card.Subsections) > 0 {
			sb.WriteString(fmt.Sprintf("  Consolidated Subsections: %s\n", strings.Join(card.Subsections, " | ")))
		}

		if len(card.Actions) > 0 {
			sb.WriteString("  Available Primary Actions:\n")
			for _, act := range card.Actions {
				sb.WriteString(fmt.Sprintf("    ▶ [%s] %s | Tooltip: %s (Shortcut: %s)\n",
					act.ID, act.Label, act.Tooltip.Description, act.Tooltip.Shortcut))
			}
		}
		sb.WriteString("--------------------------------------------------------------------------------\n")
	}

	return sb.String()
}
