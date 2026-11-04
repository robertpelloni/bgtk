package core

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// NativeHandle represents a pointer to an underlying C/C++ GObject / Bobtk widget.
type NativeHandle uintptr

// BridgeMode indicates whether the native bridge is running in pure Go simulation or linked CGO mode.
type BridgeMode string

const (
	ModePureGo  BridgeMode = "PureGo-Simulation"
	ModeLinkedC BridgeMode = "CGO-NativeLinked"
)

// NativeBridge manages the boundary between the Go 6-Pillar framework and the C/C++ Bobtk backend.
type NativeBridge struct {
	mu            sync.RWMutex
	mode          BridgeMode
	initialized   bool
	nextHandleID  uintptr
	activeHandles map[NativeHandle]string
	eventCount    int64
	eventHooks    map[string][]func(NativeHandle, []byte)
}

// NewNativeBridge instantiates a native runtime bridge.
func NewNativeBridge() *NativeBridge {
	return &NativeBridge{
		mode:          ModePureGo,
		nextHandleID:  0x1000,
		activeHandles: make(map[NativeHandle]string),
		eventHooks:    make(map[string][]func(NativeHandle, []byte)),
	}
}

// InitializeRuntime sets up the GObject / Bobtk native runtime context.
func (nb *NativeBridge) InitializeRuntime() error {
	nb.mu.Lock()
	defer nb.mu.Unlock()

	if nb.initialized {
		return nil
	}

	nb.initialized = true
	return nil
}

// AllocateNativeHandle registers an underlying C/C++ widget pointer into the Go bridge.
func (nb *NativeBridge) AllocateNativeHandle(typeName string) NativeHandle {
	nb.mu.Lock()
	defer nb.mu.Unlock()

	handle := NativeHandle(nb.nextHandleID)
	nb.nextHandleID += 8

	nb.activeHandles[handle] = typeName
	return handle
}

// FreeNativeHandle unregisters a native handle and releases bound resources.
func (nb *NativeBridge) FreeNativeHandle(handle NativeHandle) error {
	nb.mu.Lock()
	defer nb.mu.Unlock()

	if _, exists := nb.activeHandles[handle]; !exists {
		return fmt.Errorf("native handle %d not found or already freed", handle)
	}

	delete(nb.activeHandles, handle)
	return nil
}

// HandleType returns the C++ type name associated with a native handle.
func (nb *NativeBridge) HandleType(handle NativeHandle) (string, bool) {
	nb.mu.RLock()
	defer nb.mu.RUnlock()

	tName, ok := nb.activeHandles[handle]
	return tName, ok
}

// ActiveHandleCount returns the count of currently registered native handles.
func (nb *NativeBridge) ActiveHandleCount() int {
	nb.mu.RLock()
	defer nb.mu.RUnlock()
	return len(nb.activeHandles)
}

// DispatchNativeEvent routes an event received from the C/C++ main loop into Go listeners.
func (nb *NativeBridge) DispatchNativeEvent(event string, handle NativeHandle, payload []byte) {
	atomic.AddInt64(&nb.eventCount, 1)

	nb.mu.RLock()
	hooks := append([]func(NativeHandle, []byte){}, nb.eventHooks[event]...)
	nb.mu.RUnlock()

	for _, hook := range hooks {
		hook(handle, payload)
	}
}

// SubscribeNativeEvent registers a handler for C/C++ native runtime events.
func (nb *NativeBridge) SubscribeNativeEvent(event string, hook func(NativeHandle, []byte)) {
	nb.mu.Lock()
	defer nb.mu.Unlock()
	nb.eventHooks[event] = append(nb.eventHooks[event], hook)
}

// EventCount returns the total number of events dispatched across the bridge.
func (nb *NativeBridge) EventCount() int64 {
	return atomic.LoadInt64(&nb.eventCount)
}

// Mode returns the active bridging mode.
func (nb *NativeBridge) Mode() BridgeMode {
	nb.mu.RLock()
	defer nb.mu.RUnlock()
	return nb.mode
}

// IsInitialized reports if the native runtime is active.
func (nb *NativeBridge) IsInitialized() bool {
	nb.mu.RLock()
	defer nb.mu.RUnlock()
	return nb.initialized
}
