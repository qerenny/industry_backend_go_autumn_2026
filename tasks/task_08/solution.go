package main

import (
	"sync"
	"time"
)

type Clock interface{ Now() time.Time }
type Limiter struct {
	mu     sync.Mutex
	clock  Clock
	rate   float64
	burst  int
	tokens float64
	last   time.Time
}

func NewLimiter(clock Clock, ratePerSec float64, burst int) *Limiter {
	l := &Limiter{
		clock: clock,
		rate:  ratePerSec,
		burst: burst,
	}

	if clock == nil || burst <= 0 {
		return l
	}
	l.tokens = float64(burst)
	l.last = clock.Now()

	return l
}
func (l *Limiter) AllowN(n int) bool {
	if n <= 0 || n > l.burst || l.clock == nil || l.burst <= 0 {
		return false
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	now := l.clock.Now()
	elapsed := now.Sub(l.last)

	if now.After(l.last) {
		if l.rate > 0 {
			l.tokens += elapsed.Seconds() * l.rate
		}
		if l.tokens > float64(l.burst) {
			l.tokens = float64(l.burst)
		}
		l.last = now
	}

	if l.tokens < float64(n) {
		return false
	}

	l.tokens -= float64(n)

	return true
}
