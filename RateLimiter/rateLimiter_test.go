package RateLimiter

import (
	"testing"
	"time"
)

func TestFixedWindowRateLimiter(t *testing.T) {
	limiter := NewFixedWindowRateLimiter(10, time.Minute)

	// Make sure that the rate limiter allows the first request.
	if !limiter.Allow() {
		t.Errorf("Expected the first request to be allowed")
	}

	// Make sure that the rate limiter denies the second request.
	if limiter.Allow() {
		t.Errorf("Expected the second request to be denied")
	}

	// Wait for the window to expire.
	time.Sleep(time.Minute)

	// Make sure that the third request is allowed.
	if !limiter.Allow() {
		t.Errorf("Expected the third request to be allowed")
	}
}

func TestLeakyBucketRateLimiter(t *testing.T) {
	// Create a new leaky bucket rate limiter with a rate of 10 requests per second and a capacity of 10 tokens.
	rateLimiter := NewLeakyBucketRateLimiter(10, 10)

	// Make 10 requests.
	for i := 0; i < 10; i++ {
		if !rateLimiter.Allow() {
			t.Errorf("Expected the request to be allowed")
		}
	}

	// Make another request. This request should be denied because the bucket is full.
	if rateLimiter.Allow() {
		t.Errorf("Expected the request to be denied")
	}

	// Tick the bucket. This should add a token to the bucket.
	time.Sleep(1 * time.Second)
	rateLimiter.Tick()

	// Make another request. This request should be allowed because the bucket now has 1 token.
	if !rateLimiter.Allow() {
		t.Errorf("Expected the request to be allowed")
	}
}

func TestFixedWindowRateLimiter_Allow(t *testing.T) {
	// Create a rate limiter with limit of 3 requests per second
	limiter := NewFixedWindowRateLimiter(3, time.Second)

	// First request should be allowed
	allowed := limiter.Allow()
	if !allowed {
		t.Errorf("First request should be allowed")
	}

	// Second request should be allowed
	allowed = limiter.Allow()
	if !allowed {
		t.Errorf("Second request should be allowed")
	}

	// Third request should be allowed
	allowed = limiter.Allow()
	if !allowed {
		t.Errorf("Third request should be allowed")
	}

	// Fourth request should be denied
	allowed = limiter.Allow()
	if allowed {
		t.Errorf("Fourth request should be denied")
	}
}

func TestFixedWindowRateLimiter_IsAllowedAt(t *testing.T) {
	// Create a rate limiter with limit of 3 requests per minute
	limiter := NewFixedWindowRateLimiter(3, time.Minute)

	// Request at the beginning of the minute should be allowed
	allowed := limiter.IsAllowedAt(time.Now())
	if !allowed {
		t.Errorf("Request at beginning of minute should be allowed")
	}
}

func TestLeakyBucketRateLimiter_Allow(t *testing.T) {
	limiter := NewLeakyBucketRateLimiter(3, 5)

	// First request should be allowed
	allowed := limiter.Allow()
	if !allowed {
		t.Errorf("First request should be allowed")
	}

	// Second request should be allowed
	allowed = limiter.Allow()
	if !allowed {
		t.Errorf("Second request should be allowed")
	}

	// Third request should be allowed
	allowed = limiter.Allow()
	if !allowed {
		t.Errorf("Third request should be allowed")
	}

	// Fourth request should be denied
	allowed = limiter.Allow()
	if allowed {
		t.Errorf("Fourth request should be denied")
	}
}

func TestLeakyBucketRateLimiter_Tick(t *testing.T) {
	limiter := NewLeakyBucketRateLimiter(3, 5)

	// Add tokens
	limiter.Tick()
	limiter.Tick()

	// Two requests should be allowed
	allowed := limiter.Allow()
	allowed = limiter.Allow()

	// Third request should be denied
	allowed = limiter.Allow()
	if allowed {
		t.Errorf("Third request should be denied")
	}
}
