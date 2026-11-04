package core

import (
	"strings"
	"testing"
)

func TestMapOperations(t *testing.T) {
	m := NewMap(4)

	m.Put("key1", "Val1")
	m.Put("key2", "Val2")
	m.Put("key3", "Val3")

	if m.GetCount() != 3 {
		t.Fatalf("expected 3 entries, got %d", m.GetCount())
	}

	val, found := m.Get("key2")
	if !found || val.(string) != "Val2" {
		t.Errorf("unexpected value for key2: %v", val)
	}

	k0, v0, err := m.GetAt(0)
	if err != nil || k0 != "key1" || v0.(string) != "Val1" {
		t.Errorf("unexpected entry at index 0: k=%s v=%v (err=%v)", k0, v0, err)
	}

	keys := m.Keys()
	if len(keys) != 3 || keys[0] != "key1" || keys[2] != "key3" {
		t.Errorf("unexpected keys order: %+v", keys)
	}

	m.Put("key2", "Val2Updated")
	valUpdated, _ := m.Get("key2")
	if valUpdated.(string) != "Val2Updated" {
		t.Errorf("expected Val2Updated, got %v", valUpdated)
	}
	if m.GetCount() != 3 {
		t.Errorf("count should remain 3 after key update")
	}
}

func TestStateSerializer(t *testing.T) {
	ser := NewStateSerializer()

	data := map[string]interface{}{
		"theme":      "DarkModern",
		"fullscreen": true,
		"bufferSize": 256,
	}

	jsonStr, err := ser.ToJSON(data)
	if err != nil {
		t.Fatalf("failed to serialize: %v", err)
	}

	if !strings.Contains(jsonStr, `"theme": "DarkModern"`) {
		t.Errorf("serialized output missing theme property: %s", jsonStr)
	}

	var parsed map[string]interface{}
	if err := ser.FromJSON(jsonStr, &parsed); err != nil {
		t.Fatalf("failed to deserialize: %v", err)
	}

	if parsed["theme"] != "DarkModern" {
		t.Errorf("expected DarkModern, got %v", parsed["theme"])
	}
}
