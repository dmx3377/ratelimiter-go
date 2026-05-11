package ratelimit

import (
	"sync"
	"time"
)

// Bucket represents a single token bucket for a user or IP.
type Bucket struct {
	mu         sync.Mutex
	capacity   float64   // The maximum number of tokens the bucket can hold
	tokens     float64   // Current number of available tokens
	fillRate   float64   // How many tokens are added per second
	lastUpdate time.Time // The last time the bucket was accessed
}

func NewBucket(capacity float64, fillRate float64) *Bucket {
	return &Bucket{
		capacity:   capacity,
		tokens:     capacity, // Start with a full bucket
		fillRate:   fillRate,
		lastUpdate: time.Now(),
	}
}

func (b *Bucket) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()

	elapsed := now.Sub(b.lastUpdate).Seconds()

	b.tokens += elapsed * b.fillRate

	if b.tokens > b.capacity {
		b.tokens = b.capacity
	}

	b.lastUpdate = now

	if b.tokens >= 1.0 {
		b.tokens -= 1.0 // Consume one token
		return true
	}

	// Not enough tokens
	return false
}
