package visual

import (
	"sync"

	"github.com/bobtk/bobtk/bobtk/core"
)

// Action represents an abstract user interface action capable of binding to menus,
// toolbars, and shortcut keys (QAction parity).
type Action struct {
	mu           sync.RWMutex
	text         string
	iconName     string
	shortcut     string
	tooltip      string
	enabled      bool
	checkable    bool
	checked      bool
	Triggered    *core.Signal
	Toggled      *core.Signal
}

// NewAction creates an action with text label.
func NewAction(text string) *Action {
	return &Action{
		text:      text,
		enabled:   true,
		Triggered: core.NewSignal(),
		Toggled:   core.NewSignal(),
	}
}

// Trigger fires the action's Triggered signal.
func (a *Action) Trigger() {
	a.mu.Lock()
	if !a.enabled {
		a.mu.Unlock()
		return
	}
	if a.checkable {
		a.checked = !a.checked
		checkedState := a.checked
		a.mu.Unlock()
		a.Toggled.Emit(checkedState)
	} else {
		a.mu.Unlock()
	}
	a.Triggered.Emit(a.text)
}

func (a *Action) SetText(text string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.text = text
}

func (a *Action) Text() string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.text
}

func (a *Action) SetShortcut(sc string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.shortcut = sc
}

func (a *Action) Shortcut() string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.shortcut
}

func (a *Action) SetTooltip(tip string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.tooltip = tip
}

func (a *Action) Tooltip() string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.tooltip
}

func (a *Action) SetCheckable(checkable bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.checkable = checkable
}

func (a *Action) IsCheckable() bool {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.checkable
}

func (a *Action) SetChecked(checked bool) {
	a.mu.Lock()
	if a.checked != checked {
		a.checked = checked
		a.mu.Unlock()
		a.Toggled.Emit(checked)
		return
	}
	a.mu.Unlock()
}

func (a *Action) IsChecked() bool {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.checked
}

func (a *Action) SetEnabled(enabled bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.enabled = enabled
}

func (a *Action) IsEnabled() bool {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.enabled
}
