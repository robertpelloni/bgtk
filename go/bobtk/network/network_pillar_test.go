package network

import (
	"encoding/json"
	"testing"
)

func TestBroadwayGatewayLifecycle(t *testing.T) {
	bg := NewBroadwayGateway(8085)
	err := bg.Start()
	if err != nil {
		t.Fatalf("failed to start broadway gateway: %v", err)
	}

	if !bg.IsRunning() {
		t.Errorf("expected gateway to be running")
	}

	client := bg.RegisterClient("session-1", "127.0.0.1")
	if client == nil || !client.Connected {
		t.Fatalf("expected active client session")
	}

	if bg.ClientCount() != 1 {
		t.Errorf("expected 1 connected client, got %d", bg.ClientCount())
	}

	_ = bg.Stop()
	if bg.IsRunning() {
		t.Errorf("expected gateway to be stopped")
	}
	if bg.ClientCount() != 0 {
		t.Errorf("expected 0 connected clients after stop, got %d", bg.ClientCount())
	}
}

func TestWebSocketHubDispatch(t *testing.T) {
	hub := NewWebSocketHub(9090)

	hub.RegisterHandler("echo", func(params json.RawMessage) (interface{}, error) {
		var input string
		if err := json.Unmarshal(params, &input); err != nil {
			return nil, err
		}
		return "ECHO:" + input, nil
	})

	paramsRaw, _ := json.Marshal("HelloWorld")
	msg := RPCMessage{
		Method: "echo",
		Params: paramsRaw,
		ID:     "msg-42",
	}

	resp, err := hub.Dispatch(msg)
	if err != nil {
		t.Fatalf("failed to dispatch RPC message: %v", err)
	}

	var parsed map[string]interface{}
	_ = json.Unmarshal([]byte(resp), &parsed)
	if parsed["result"] != "ECHO:HelloWorld" {
		t.Errorf("expected ECHO:HelloWorld, got %v", parsed["result"])
	}
}
