package system

import (
	"testing"
)

func TestFileSystemWatcherAddRemoveEvent(t *testing.T) {
	w := NewFileSystemWatcher()

	err := w.AddPath("/app/config")
	if err != nil {
		t.Fatalf("failed to add path: %v", err)
	}

	if len(w.WatchedPaths()) != 1 {
		t.Fatalf("expected 1 watched path, got %d", len(w.WatchedPaths()))
	}

	var receivedPath string
	var receivedOp string
	w.OnFileChanged(func(ev FileWatcherEvent) {
		receivedPath = ev.Path
		receivedOp = ev.Op
	})

	w.TriggerEvent("/app/config", "WRITE")

	if receivedPath != "/app/config" || receivedOp != "WRITE" {
		t.Errorf("expected /app/config and WRITE, got %s / %s", receivedPath, receivedOp)
	}

	err = w.RemovePath("/app/config")
	if err != nil {
		t.Fatalf("failed to remove path: %v", err)
	}

	if len(w.WatchedPaths()) != 0 {
		t.Errorf("expected 0 watched paths after removal, got %d", len(w.WatchedPaths()))
	}
}
