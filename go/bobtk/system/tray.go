package system

import (
	"fmt"
	"sync"
)

// TrayMenuItem represents an interactive menu option in the System Tray.
type TrayMenuItem struct {
	ID          string
	Label       string
	Tooltip     string
	IsSeparator bool
	Action      func()
}

// SystemTray provides cross-platform desktop system tray presence and menu controls.
type SystemTray struct {
	mu            sync.RWMutex
	iconName      string
	tooltip       string
	visible       bool
	menuItems     []TrayMenuItem
	serverManager *ServerManager
	quitHandler   func()
}

// NewSystemTray creates a system tray instance bound to a ServerManager.
func NewSystemTray(iconName string, sm *ServerManager) *SystemTray {
	if sm == nil {
		sm = NewServerManager()
	}
	return &SystemTray{
		iconName:      iconName,
		tooltip:       "bgtk System Tray - Autonomous Monitor",
		visible:       true,
		menuItems:     make([]TrayMenuItem, 0),
		serverManager: sm,
	}
}

// SetTooltip sets the tray tooltip text.
func (st *SystemTray) SetTooltip(tip string) {
	st.mu.Lock()
	defer st.mu.Unlock()
	st.tooltip = tip
}

// GetTooltip returns the tray tooltip text.
func (st *SystemTray) GetTooltip() string {
	st.mu.RLock()
	defer st.mu.RUnlock()
	return st.tooltip
}

// SetVisible controls tray visibility.
func (st *SystemTray) SetVisible(vis bool) {
	st.mu.Lock()
	defer st.mu.Unlock()
	st.visible = vis
}

// IsVisible returns tray visibility.
func (st *SystemTray) IsVisible() bool {
	st.mu.RLock()
	defer st.mu.RUnlock()
	return st.visible
}

// ServerManager returns the associated server manager.
func (st *SystemTray) ServerManager() *ServerManager {
	return st.serverManager
}

// AddMenuItem adds an action item to the tray menu with a descriptive tooltip.
func (st *SystemTray) AddMenuItem(id, label, tooltip string, action func()) {
	st.mu.Lock()
	defer st.mu.Unlock()

	st.menuItems = append(st.menuItems, TrayMenuItem{
		ID:          id,
		Label:       label,
		Tooltip:     tooltip,
		IsSeparator: false,
		Action:      action,
	})
}

// AddSeparator adds a visual divider to the tray menu.
func (st *SystemTray) AddSeparator() {
	st.mu.Lock()
	defer st.mu.Unlock()

	st.menuItems = append(st.menuItems, TrayMenuItem{
		IsSeparator: true,
	})
}

// MenuItems returns a copy of current tray menu items.
func (st *SystemTray) MenuItems() []TrayMenuItem {
	st.mu.RLock()
	defer st.mu.RUnlock()

	items := make([]TrayMenuItem, len(st.menuItems))
	copy(items, st.menuItems)
	return items
}

// SetQuitHandler registers an exit hook invoked upon quitting.
func (st *SystemTray) SetQuitHandler(fn func()) {
	st.mu.Lock()
	defer st.mu.Unlock()
	st.quitHandler = fn
}

// QuitApplication stops all running servers and invokes the quit handler.
func (st *SystemTray) QuitApplication() {
	fmt.Println("[SystemTray] Initiating application quit sequence...")
	if st.serverManager != nil {
		fmt.Printf("[SystemTray] Gracefully stopping %d server(s)...\n", st.serverManager.RunningCount())
		errs := st.serverManager.StopAllServers()
		for _, err := range errs {
			fmt.Printf("[SystemTray] Warning while stopping server: %v\n", err)
		}
	}
	st.mu.RLock()
	handler := st.quitHandler
	st.mu.RUnlock()

	if handler != nil {
		handler()
	}
	fmt.Println("[SystemTray] Application shutdown complete.")
}
