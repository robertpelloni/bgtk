package system

import "fmt"

// SystemPillar manages low-level OS integrations.
type SystemPillar struct {
	initialized bool
}

func NewSystemPillar() *SystemPillar {
	return &SystemPillar{initialized: false}
}

func (s *SystemPillar) Initialize() {
	if s.initialized { return }
	fmt.Println("SystemPillar initialized.")
	s.initialized = true
}
