package core

import (
	"testing"
)

func TestSettingsPersistenceAndEvents(t *testing.T) {
	s := NewSettings()

	var eventVal interface{}
	s.OnSettingChanged("editor.tab_size", func(val interface{}) {
		eventVal = val
	})

	s.SetValue("editor.tab_size", 4)
	s.SetValue("editor.font", "Fira Code")

	if s.Value("editor.tab_size") != 4 {
		t.Errorf("expected 4, got %v", s.Value("editor.tab_size"))
	}

	if eventVal != 4 {
		t.Errorf("expected eventVal 4, got %v", eventVal)
	}

	if s.Value("non.existent", "default") != "default" {
		t.Errorf("expected default fallback value")
	}

	if !s.Contains("editor.font") {
		t.Errorf("expected contains to return true for editor.font")
	}

	s.Remove("editor.font")
	if s.Contains("editor.font") {
		t.Errorf("expected font to be removed")
	}

	if len(s.AllKeys()) != 1 {
		t.Errorf("expected 1 remaining key, got %d", len(s.AllKeys()))
	}
}
