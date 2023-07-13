package RateLimiter

import (
	"time"
)

func NewTokenBucketRateLimiter(rate int, capacity int) *TokenBucketRateLimiter {
	return &TokenBucketRateLimiter{
		Rate:     rate,
		Capacity: capacity,
		Tokens:   0,
	}
}

type TokenBucketRateLimiter struct {
	Rate     int
	Capacity int
	LastTick time.Time
	Tokens   int
}

func (r *TokenBucketRateLimiter) Allow() bool {
	// Calculate the number of tokens that should be added to the bucket.
	tokensToAdd := int(time.Since(r.LastTick).Seconds()) * r.Rate

	// Add the tokens to the bucket.
	r.Tokens += tokensToAdd

	// If there are enough tokens, allow the request.
	if r.Tokens >= 1 {
		r.Tokens -= 1
		return true
	}

	// Otherwise, deny the request.
	return false
}

func (r *TokenBucketRateLimiter) Tick() {
	// Update the last tick time.
	r.LastTick = time.Now()
}
