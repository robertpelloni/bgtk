package core

import (
	"testing"
)

func TestNativeBridgeLifecycleAndHandles(t *testing.T) {
	bridge := NewNativeBridge()
	if err := bridge.InitializeRuntime(); err != nil {
		t.Fatalf("failed to initialize native bridge: %v", err)
	}

	if !bridge.IsInitialized() {
		t.Errorf("expected bridge to be initialized")
	}

	hWindow := bridge.AllocateNativeHandle("BobguiWindow")
	hButton := bridge.AllocateNativeHandle("BobguiButton")

	if bridge.ActiveHandleCount() != 2 {
		t.Errorf("expected 2 active handles, got %d", bridge.ActiveHandleCount())
	}

	tName, ok := bridge.HandleType(hWindow)
	if !ok || tName != "BobguiWindow" {
		t.Errorf("expected BobguiWindow, got %s (ok=%v)", tName, ok)
	}

	var eventPayloadReceived string
	bridge.SubscribeNativeEvent("widget.clicked", func(h NativeHandle, payload []byte) {
		if h == hButton {
			eventPayloadReceived = string(payload)
		}
	})

	bridge.DispatchNativeEvent("widget.clicked", hButton, []byte("ok"))

	if eventPayloadReceived != "ok" {
		t.Errorf("expected payload 'ok', got %q", eventPayloadReceived)
	}

	if bridge.EventCount() != 1 {
		t.Errorf("expected 1 event dispatched, got %d", bridge.EventCount())
	}

	if err := bridge.FreeNativeHandle(hWindow); err != nil {
		t.Fatalf("failed to free native handle: %v", err)
	}

	if bridge.ActiveHandleCount() != 1 {
		t.Errorf("expected 1 active handle after free, got %d", bridge.ActiveHandleCount())
	}
}

func TestNativeBridgeTargetPlatformAndModeSwitching(t *testing.T) {
	bridge := NewNativeBridge()

	if bridge.TargetOS() == "" {
		t.Errorf("expected non-empty TargetOS")
	}
	if bridge.TargetArch() == "" {
		t.Errorf("expected non-empty TargetArch")
	}

	if bridge.Mode() != ModePureGo {
		t.Errorf("expected default ModePureGo, got %s", bridge.Mode())
	}

	bridge.SetMode(ModeLinkedC)
	if bridge.Mode() != ModeLinkedC {
		t.Errorf("expected ModeLinkedC after SetMode, got %s", bridge.Mode())
	}

	bridge.SetMode(ModePureGo)
	if bridge.Mode() != ModePureGo {
		t.Errorf("expected ModePureGo after reverting, got %s", bridge.Mode())
	}
}
