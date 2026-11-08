package core

import "fmt"

// CorePillar manages the creation and state of all data, rendering,
// layout, machine learning, logic, and state management subsystems.
type CorePillar struct {
	initialized  bool
	Scheduler    *QuantumScheduler
	Reflect      *ReflectContext
	Entities     *EntityManager
	Bridge       *NativeBridge
	Heap         *HeapPool
	StateMap     *Map
	Serializer   *StateSerializer
	Settings     *Settings
	StateMachine *StateMachine
	Workers      *WorkPool
}

// NewCorePillar creates a new instance of the CorePillar with foundational subsystems.
func NewCorePillar() *CorePillar {
	initialState := NewStateNode("Idle")
	initialState.AddTransition("ACTIVATE", "Running")
	runningState := NewStateNode("Running")
	runningState.AddTransition("DEACTIVATE", "Idle")

	sm := NewStateMachine(initialState)
	sm.AddState(runningState)

	return &CorePillar{
		initialized:  false,
		Scheduler:    NewQuantumScheduler(16),
		Reflect:      NewReflectContext(),
		Entities:     NewEntityManager(),
		Bridge:       NewNativeBridge(),
		Heap:         NewHeapPool(4096),
		StateMap:     NewMap(16),
		Serializer:   NewStateSerializer(),
		Settings:     NewSettings(),
		StateMachine: sm,
		Workers:      NewWorkPool(8, 32),
	}
}

// Initialize sets up the Core subsystem and bridges to the C++ backend.
func (c *CorePillar) Initialize() {
	if c.initialized {
		return
	}

	_ = c.Bridge.InitializeRuntime()
	fmt.Println("CorePillar initialized.")
	c.initialized = true
}

// QuantumScheduler returns the task execution engine.
func (c *CorePillar) QuantumScheduler() *QuantumScheduler {
	return c.Scheduler
}

// ReflectContext returns the dynamic reflection registry.
func (c *CorePillar) ReflectContext() *ReflectContext {
	return c.Reflect
}

// EntityManager returns the active ECS entity manager.
func (c *CorePillar) EntityManager() *EntityManager {
	return c.Entities
}

// NativeBridge returns the C/C++ runtime bridge.
func (c *CorePillar) NativeBridge() *NativeBridge {
	return c.Bridge
}

// HeapPool returns the Ultimate++ inspired slab memory pool.
func (c *CorePillar) HeapPool() *HeapPool {
	return c.Heap
}

// Map returns the ordered Map container.
func (c *CorePillar) Map() *Map {
	return c.StateMap
}

// StateSerializer returns the state serialization engine.
func (c *CorePillar) StateSerializer() *StateSerializer {
	return c.Serializer
}

// AppSettings returns the application configuration store.
func (c *CorePillar) AppSettings() *Settings {
	return c.Settings
}

// AppStateMachine returns the core state machine.
func (c *CorePillar) AppStateMachine() *StateMachine {
	return c.StateMachine
}

// NewRingBuffer creates a thread-safe circular FIFO buffer of specified capacity.
func (c *CorePillar) NewRingBuffer(capacity int) *RingBuffer[interface{}] {
	return NewRingBuffer[interface{}](capacity)
}

// WorkPool returns the adaptive worker pool dispatcher.
func (c *CorePillar) WorkPool() *WorkPool {
	return c.Workers
}
