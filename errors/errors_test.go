package errors

import "testing"

func TestNew(t *testing.T) {
	err := New(NotFound, "user.not_found", "user not found")
	if err.Kind() != NotFound {
		t.Errorf("expected NotFound, got %v", err.Kind())
	}
	if err.Code() != "user.not_found" {
		t.Errorf("expected user.not_found, got %s", err.Code())
	}
}

func TestHTTPStatus(t *testing.T) {
	tests := []struct {
		kind Kind
		want int
	}{
		{NotFound, 404},
		{Validation, 400},
		{Conflict, 409},
		{Unauthorized, 401},
		{Forbidden, 403},
		{RateLimited, 429},
		{Unavailable, 503},
		{Internal, 500},
	}
	for _, tt := range tests {
		if got := HTTPStatus(tt.kind); got != tt.want {
			t.Errorf("HTTPStatus(%v) = %d, want %d", tt.kind, got, tt.want)
		}
	}
}
