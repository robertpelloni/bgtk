package core

import "fmt"

// CorePillar manages the creation and state of all data, rendering,
// layout, machine learning, logic, and state management subsystems.
type CorePillar struct {
	initialized bool
	Scheduler   *QuantumScheduler
	Reflect     *ReflectContext
	Entities    *EntityManager
}

// NewCorePillar creates a new instance of the CorePillar with foundational subsystems.
func NewCorePillar() *CorePillar {
	return &CorePillar{
		initialized: false,
		Scheduler:   NewQuantumScheduler(16),
		Reflect:     NewReflectContext(),
		Entities:    NewEntityManager(),
	}
}

// Initialize sets up the Core subsystem and bridges to the C++ backend.
func (c *CorePillar) Initialize() {
	if c.initialized {
		return
	}

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
