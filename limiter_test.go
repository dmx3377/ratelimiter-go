package ratelimit

import (
	"testing"
	"time"
)

func TestLimiter_Allow(t *testing.T) {
	l := NewLimiter(2, 1)
	key := "test-user"

	if !l.Allow(key) {
		t.Errorf("expected first request to be allowed")
	}
	if !l.Allow(key) {
		t.Errorf("expected second request to be allowed")
	}

	if l.Allow(key) {
		t.Errorf("expected third request to be blocked")
	}

	time.Sleep(1100 * time.Millisecond)

	if !l.Allow(key) {
		t.Errorf("expected request after refill to be allowed")
	}

	if l.Allow(key) {
		t.Errorf("expected subsequent request to be blocked")
	}
}