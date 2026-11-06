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

func TestServerManagerStartAndRestart(t *testing.T) {
	sm := NewServerManager()
	startCalls := 0
	stopCalls := 0

	sm.RegisterServerWithStart("daemon.alpha", "Alpha Daemon", "Background worker", 8888, func() error {
		startCalls++
		return nil
	}, func() error {
		stopCalls++
		return nil
	})

	srv, ok := sm.GetServer("daemon.alpha")
	if !ok || srv.Status != StatusRunning {
		t.Fatalf("expected running daemon.alpha")
	}

	// Stop the server
	err := sm.StopServer("daemon.alpha")
	if err != nil || srv.Status != StatusStopped || stopCalls != 1 {
		t.Fatalf("expected server stopped, status=%s, stopCalls=%d", srv.Status, stopCalls)
	}

	// Start the server
	err = sm.StartServer("daemon.alpha")
	if err != nil || srv.Status != StatusRunning || startCalls != 1 {
		t.Fatalf("expected server running, status=%s, startCalls=%d", srv.Status, startCalls)
	}

	// Restart the server
	err = sm.RestartServer("daemon.alpha")
	if err != nil || srv.Status != StatusRunning || startCalls != 2 || stopCalls != 2 {
		t.Fatalf("expected restart success, status=%s, starts=%d, stops=%d", srv.Status, startCalls, stopCalls)
	}
}

func TestClipboardOperations(t *testing.T) {
	cb := NewClipboard()

	if cb.GetText() != "" {
		t.Errorf("expected empty initial clipboard, got %q", cb.GetText())
	}

	listenerCalled := false
	notifiedText := ""
	cb.AddListener(func(txt string) {
		listenerCalled = true
		notifiedText = txt
	})

	cb.SetText("Sample Clipboard Payload")
	if cb.GetText() != "Sample Clipboard Payload" || !listenerCalled || notifiedText != "Sample Clipboard Payload" {
		t.Errorf("unexpected clipboard text or listener result: %q (notified: %q)", cb.GetText(), notifiedText)
	}

	cb.Clear()
	if cb.GetText() != "" {
		t.Errorf("expected empty clipboard after clear, got %q", cb.GetText())
	}
}
