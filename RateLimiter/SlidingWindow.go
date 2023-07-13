package RateLimiter

import (
	"time"
)

func NewSlidingWindowRateLimiter(windowSize time.Duration, maxRequests int) *SlidingWindowRateLimiter {
	return &SlidingWindowRateLimiter{
		WindowSize:  windowSize,
		MaxRequests: maxRequests,
		Requests:    make(map[string]int),
	}
}

type SlidingWindowRateLimiter struct {
	WindowSize  time.Duration
	MaxRequests int
	Requests    map[string]int
}

func (r *SlidingWindowRateLimiter) Allow() bool {
	// Get the current time.
	now := time.Now()

	// Get the number of requests for the current window.
	requestCount := r.Requests[now.String()]

	// If the number of requests is less than the maximum, allow the request.
	if requestCount < r.MaxRequests {
		return true
	}

	// Otherwise, deny the request.
	return false
}
