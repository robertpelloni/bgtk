package system

import (
	"testing"
)

func TestNotificationCenterSendAndHistory(t *testing.T) {
	nc := NewNotificationCenter()

	if nc.Count() != 0 {
		t.Fatalf("expected initial count 0, got %d", nc.Count())
	}

	received := false
	var receivedItem NotificationItem

	nc.AddListener(func(item NotificationItem) {
		received = true
		receivedItem = item
	})

	sent := nc.Send("alert-1", "System Alert", "Broadway Server connection active", UrgencyNormal)

	if !received {
		t.Fatalf("expected listener to be triggered")
	}

	if receivedItem.ID != "alert-1" || receivedItem.Title != "System Alert" {
		t.Errorf("mismatched received item: %+v", receivedItem)
	}

	if nc.Count() != 1 {
		t.Errorf("expected count 1, got %d", nc.Count())
	}

	history := nc.History()
	if len(history) != 1 || history[0].ID != sent.ID {
		t.Errorf("unexpected history contents: %+v", history)
	}

	nc.Clear()
	if nc.Count() != 0 {
		t.Errorf("expected count 0 after Clear(), got %d", nc.Count())
	}
}
