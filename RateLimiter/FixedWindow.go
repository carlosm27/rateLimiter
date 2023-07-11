package RateLimiter

import (
	"time"
)

type FixedWindowRateLimiter struct {
	Limit      int
	WindowSize time.Duration
}

func NewFixedWindowRateLimiter(limit int, windowSize time.Duration) *FixedWindowRateLimiter {
	return &FixedWindowRateLimiter{
		Limit:      limit,
		WindowSize: windowSize,
	}
}

func (r *FixedWindowRateLimiter) Allow() bool {
	now := time.Now()
	start := time.Unix(now.Unix(), 0)

	// Count the number of requests in the current window.
	requests := 0

	for t := start; t.Before(now); t = t.Add(time.Second) {
		if t.After(t) {
			requests++
		}
	}

	// If the number of requests is less than the limit, then the request is allowed.
	return requests < r.Limit
}

func (r *FixedWindowRateLimiter) IsAllowedAt(t time.Time) bool {
	return t.Minute() < time.Now().Minute()
}
