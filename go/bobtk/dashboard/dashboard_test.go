package dashboard

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func do(t *testing.T, h http.Handler, method, path string) *httptest.ResponseRecorder {
	t.Helper()
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(method, path, nil))
	return rec
}

func TestPageAndHealth(t *testing.T) {
	s := New("9.9.9-test")
	h := s.Handler()
	if rec := do(t, h, "GET", "/"); rec.Code != 200 || !strings.Contains(rec.Body.String(), "Bobtk Dashboard") {
		t.Fatalf("GET / = %d, want 200 dashboard page", rec.Code)
	}
	rec := do(t, h, "GET", "/health")
	var body map[string]any
	if rec.Code != 200 || json.NewDecoder(rec.Body).Decode(&body) != nil || body["ok"] != true {
		t.Fatalf("GET /health = %d %q, want ok:true", rec.Code, rec.Body.String())
	}
	if rec := do(t, h, "GET", "/nope"); rec.Code != 404 {
		t.Fatalf("GET /nope = %d, want 404", rec.Code)
	}
}

func TestStatusAndIdempotentInit(t *testing.T) {
	s := New("9.9.9-test")
	h := s.Handler()
	var st Status
	rec := do(t, h, "GET", "/api/status")
	if err := json.NewDecoder(rec.Body).Decode(&st); err != nil || st.Total != 6 || st.Initialized != 0 {
		t.Fatalf("initial status = %+v err=%v, want 0/6", st, err)
	}
	// Initialize twice: second run must stay a safe no-op (still 6/6, no error).
	for i := 0; i < 2; i++ {
		rec = do(t, h, "POST", "/api/init-all")
		if rec.Code != 200 {
			t.Fatalf("init-all #%d = %d", i, rec.Code)
		}
	}
	st = Status{}
	if err := json.NewDecoder(rec.Body).Decode(&st); err != nil || st.Initialized != 6 {
		t.Fatalf("after init-all = %+v err=%v, want 6/6", st, err)
	}
	if rec := do(t, h, "POST", "/api/init/bogus"); rec.Code != 404 {
		t.Fatalf("init unknown pillar = %d, want 404", rec.Code)
	}
}
