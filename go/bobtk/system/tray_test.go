package system

import (
	"testing"
)

func TestServerManagerAndSystemTray(t *testing.T) {
	sm := NewServerManager()
	serverStopped := false

	proc := sm.RegisterServer("test.server", "Test Server", "Mock server daemon", 1234, func() error {
		serverStopped = true
		return nil
	})

	if proc == nil {
		t.Fatalf("expected registered server process, got nil")
	}

	if sm.RunningCount() != 1 {
		t.Fatalf("expected running count 1, got %d", sm.RunningCount())
	}

	tray := NewSystemTray("test-icon", sm)
	tray.AddMenuItem("item1", "Test Item", "Test tooltip", func() {})
	tray.AddSeparator()

	if len(tray.MenuItems()) != 2 {
		t.Fatalf("expected 2 tray menu items, got %d", len(tray.MenuItems()))
	}

	tray.QuitApplication()

	if !serverStopped {
		t.Errorf("expected server to be stopped upon tray quit")
	}

	if sm.RunningCount() != 0 {
		t.Errorf("expected 0 running servers after quit, got %d", sm.RunningCount())
	}
}
