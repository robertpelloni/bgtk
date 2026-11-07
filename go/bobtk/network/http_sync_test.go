package network

import (
	"testing"
)

func TestNetworkAccessManagerRequests(t *testing.T) {
	nam := NewNetworkAccessManager()

	resGet := nam.Get("https://api.bobtk.org/v1/telemetry")
	if resGet.StatusCode != 200 || string(resGet.Body) != "OK GET https://api.bobtk.org/v1/telemetry" {
		t.Errorf("unexpected GET response: %+v", resGet)
	}

	resPost := nam.Post("https://api.bobtk.org/v1/telemetry", []byte("payload"))
	if resPost.StatusCode != 201 || string(resPost.Body) != "OK POST https://api.bobtk.org/v1/telemetry" {
		t.Errorf("unexpected POST response: %+v", resPost)
	}

	if nam.HistoryCount() != 2 {
		t.Errorf("expected 2 requests logged, got %d", nam.HistoryCount())
	}
}

func TestNetworkAccessManagerPutDeleteHeaders(t *testing.T) {
	nam := NewNetworkAccessManager()

	nam.SetHeader("Authorization", "Bearer bgtk-token-xyz")
	if nam.GetHeader("Authorization") != "Bearer bgtk-token-xyz" {
		t.Errorf("expected header value 'Bearer bgtk-token-xyz', got %q", nam.GetHeader("Authorization"))
	}

	resPut := nam.Put("https://api.bobtk.org/v1/resource/1", []byte("update payload"))
	if resPut.StatusCode != 200 || string(resPut.Body) != "OK PUT https://api.bobtk.org/v1/resource/1" {
		t.Errorf("unexpected PUT response: %+v", resPut)
	}

	resDelete := nam.Delete("https://api.bobtk.org/v1/resource/1")
	if resDelete.StatusCode != 204 || string(resDelete.Body) != "OK DELETE https://api.bobtk.org/v1/resource/1" {
		t.Errorf("unexpected DELETE response: %+v", resDelete)
	}

	if nam.HistoryCount() != 2 {
		t.Errorf("expected 2 requests logged in history, got %d", nam.HistoryCount())
	}
}

func TestSyncEnginePeersAndCursors(t *testing.T) {
	se := NewSyncEngine()

	se.ConnectPeer("192.168.1.10:9090")
	se.ConnectPeer("192.168.1.11:9090")

	if se.PeerCount() != 2 {
		t.Errorf("expected 2 connected peers, got %d", se.PeerCount())
	}

	se.TrackCursor("client-1", 150.5, 300.0)
	c, found := se.GetCursor("client-1")
	if !found || c.X != 150.5 || c.Y != 300.0 {
		t.Errorf("unexpected cursor data: %+v (found=%v)", c, found)
	}

	se.Broadcast([]byte("DELTA_PACKET_0"))
	if se.BroadcastCount() != 1 {
		t.Errorf("expected 1 broadcast, got %d", se.BroadcastCount())
	}
}
