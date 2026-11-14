package bridge

import "testing"

func TestBridgeConsistency(t *testing.T) {
	if Available() {
		if Ping() != "pong" {
			t.Fatalf("cgo bridge Ping() = %q, want pong", Ping())
		}
	} else if Ping() != "" {
		t.Fatalf("stub bridge Ping() = %q, want empty", Ping())
	}
}
