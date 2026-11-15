package core

import (
	"fmt"
	"sync"
)

// EntityID uniquely identifies an entity within the ECS / state graph.
type EntityID uint64

// Component represents modular data attached to an entity.
type Component interface{}

// Entity represents an architectural object in Bobtk's data/widget hierarchy.
type Entity struct {
	ID         EntityID
	Name       string
	Components map[string]Component
}

// EntityManager manages the lifecycle, caching, and component queries of all entities.
type EntityManager struct {
	mu         sync.RWMutex
	nextID     EntityID
	entities   map[EntityID]*Entity
	byName     map[string]EntityID
}

// NewEntityManager instantiates a thread-safe entity manager.
func NewEntityManager() *EntityManager {
	return &EntityManager{
		nextID:   1,
		entities: make(map[EntityID]*Entity),
		byName:   make(map[string]EntityID),
	}
}

// CreateEntity allocates a new entity with a distinct name.
func (em *EntityManager) CreateEntity(name string) *Entity {
	em.mu.Lock()
	defer em.mu.Unlock()

	id := em.nextID
	em.nextID++

	ent := &Entity{
		ID:         id,
		Name:       name,
		Components: make(map[string]Component),
	}

	em.entities[id] = ent
	em.byName[name] = id
	return ent
}

// AttachComponent associates a component with an entity by type key.
func (em *EntityManager) AttachComponent(id EntityID, key string, comp Component) error {
	em.mu.Lock()
	defer em.mu.Unlock()

	ent, ok := em.entities[id]
	if !ok {
		return fmt.Errorf("entity ID %d not found", id)
	}

	ent.Components[key] = comp
	return nil
}

// GetComponent retrieves a component from an entity by key.
func (em *EntityManager) GetComponent(id EntityID, key string) (Component, bool) {
	em.mu.RLock()
	defer em.mu.RUnlock()

	ent, ok := em.entities[id]
	if !ok {
		return nil, false
	}

	comp, found := ent.Components[key]
	return comp, found
}

// EntityCount returns the current count of living entities.
func (em *EntityManager) EntityCount() int {
	em.mu.RLock()
	defer em.mu.RUnlock()
	return len(em.entities)
}
