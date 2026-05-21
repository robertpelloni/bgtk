package visual

import "fmt"

// VisualPillar manages core UI rendering and widgets.
type VisualPillar struct {
	initialized bool
}

func NewVisualPillar() *VisualPillar {
	return &VisualPillar{initialized: false}
}

func (v *VisualPillar) Initialize() {
	if v.initialized { return }
	fmt.Println("VisualPillar initialized.")
	v.initialized = true
}
