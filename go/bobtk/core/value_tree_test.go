package core_test

import (
	"testing"

	"github.com/bobtk/bobtk/bobtk/core"
)

func TestValueTreePropertiesAndHierarchy(t *testing.T) {
	root := core.NewValueTree("ApplicationSettings")

	if root.Type() != "ApplicationSettings" {
		t.Fatalf("expected type ApplicationSettings, got %q", root.Type())
	}

	notified := false
	root.AddListener(func(tree *core.ValueTree, property string) {
		if property == "theme" {
			notified = true
		}
	})

	root.SetProperty("theme", "DarkMatrix")

	if !notified {
		t.Errorf("expected property mutation notification")
	}

	val, ok := root.GetProperty("theme")
	if !ok || val != "DarkMatrix" {
		t.Errorf("expected 'DarkMatrix', got %v", val)
	}

	child := core.NewValueTree("AudioSettings")
	child.SetProperty("sampleRate", 48000)
	child.SetProperty("bufferSize", 256)

	root.AddChild(child)

	if root.NumChildren() != 1 {
		t.Fatalf("expected 1 child, got %d", root.NumChildren())
	}

	retrieved := root.GetChildWithName("AudioSettings")
	if retrieved == nil {
		t.Fatalf("expected to find child AudioSettings")
	}

	if retrieved.Parent() != root {
		t.Errorf("expected child's parent to reference root")
	}

	sr, ok := retrieved.GetProperty("sampleRate")
	if !ok || sr != 48000 {
		t.Errorf("expected sampleRate 48000, got %v", sr)
	}
}
