package ratelimit

import (
	"sync"
)

type Limiter struct {
	mu       sync.RWMutex
	buckets  map[string]*Bucket
	capacity float64
	fillRate float64
}

func NewLimiter(capacity float64, fillRate float64) *Limiter {
	return &Limiter{
		buckets:  make(map[string]*Bucket),
		capacity: capacity,
		fillRate: fillRate,
	}
}

// Allow checks if the given key is allowed to make a request
func (l *Limiter) Allow(key string) bool {
	l.mu.RLock()
	bucket, exists := l.buckets[key]
	l.mu.RUnlock()

	if !exists {
		l.mu.Lock()

		bucket, exists = l.buckets[key]
		if !exists {
			bucket = NewBucket(l.capacity, l.fillRate)
			l.buckets[key] = bucket
		}

		l.mu.Unlock()
	}

	return bucket.Allow()
}
