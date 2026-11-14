package tools

import "fmt"

// ToolsPillar manages IDE integrations, compilers, unit testing, and reporting.
type ToolsPillar struct {
	initialized bool
}

func NewToolsPillar() *ToolsPillar {
	return &ToolsPillar{initialized: false}
}

func (t *ToolsPillar) Initialize() {
	if t.initialized { return }
	fmt.Println("ToolsPillar initialized.")
	t.initialized = true
}

// Initialized reports whether Initialize has completed.
func (t *ToolsPillar) Initialized() bool { return t.initialized }
