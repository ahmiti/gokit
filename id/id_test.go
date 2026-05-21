package id

import "testing"

func TestNew(t *testing.T) {
	id := New("req")
	if len(id) < 10 {
		t.Errorf("id too short: %s", id)
	}
	if id[:4] != "req_" {
		t.Errorf("expected prefix req_, got %s", id[:4])
	}
}

func TestNewRaw(t *testing.T) {
	id := NewRaw()
	if len(id) != 26 {
		t.Errorf("expected 26 chars, got %d", len(id))
	}
}
