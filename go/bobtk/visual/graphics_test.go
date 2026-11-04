package visual

import (
	"testing"
)

func TestGraphicsContextDrawingOperations(t *testing.T) {
	ctx := NewGraphicsContext(Rect{0, 0, 800, 600})
	ctx.SetColour(ColorRed)
	ctx.FillRect(10, 10, 100, 50)
	ctx.SetColour(ColorBlue)
	ctx.DrawText("Audio VU", 120, 10, 80, 20)

	if ctx.OperationCount() != 2 {
		t.Fatalf("expected 2 recorded draw operations, got %d", ctx.OperationCount())
	}

	lastOp, err := ctx.LastOperation()
	if err != nil {
		t.Fatalf("failed to get last operation: %v", err)
	}

	if lastOp.Type != "DRAW_TEXT" || lastOp.Text != "Audio VU" {
		t.Errorf("unexpected last operation: %+v", lastOp)
	}

	path := NewPath()
	path.StartNewSubPath(0, 0)
	path.LineTo(50, 50)
	path.CloseSubPath()

	if path.CommandCount() != 3 {
		t.Errorf("expected 3 path commands, got %d", path.CommandCount())
	}

	ctx.StrokePath(path)
	if ctx.OperationCount() != 3 {
		t.Errorf("expected 3 operations after stroking path, got %d", ctx.OperationCount())
	}
}
