package core

import "fmt"

// CorePillar manages the creation and state of all data, rendering,
// layout, machine learning, logic, and state management subsystems.
type CorePillar struct {
	initialized bool
	Scheduler   *QuantumScheduler
	Reflect     *ReflectContext
	Entities    *EntityManager
	Bridge      *NativeBridge
	Heap        *HeapPool
	StateMap    *Map
	Serializer  *StateSerializer
}

// NewCorePillar creates a new instance of the CorePillar with foundational subsystems.
func NewCorePillar() *CorePillar {
	return &CorePillar{
		initialized: false,
		Scheduler:   NewQuantumScheduler(16),
		Reflect:     NewReflectContext(),
		Entities:    NewEntityManager(),
		Bridge:      NewNativeBridge(),
		Heap:        NewHeapPool(4096),
		StateMap:    NewMap(16),
		Serializer:  NewStateSerializer(),
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
