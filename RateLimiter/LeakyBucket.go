package RateLimiter

func NewLeakyBucketRateLimiter(rate int, capacity int) *LeakyBucketRateLimiter {
	return &LeakyBucketRateLimiter{
		Rate:     rate,
		Capacity: capacity,
		Tokens:   make(chan int, capacity),
	}
}

type LeakyBucketRateLimiter struct {
	Rate     int
	Capacity int
	Tokens   chan int
}

func (r *LeakyBucketRateLimiter) Allow() bool {
	// Get the current number of tokens.
	tokenCount := len(r.Tokens)

	// If there are enough tokens, allow the request.
	if tokenCount > 0 {
		r.Tokens <- 1
		return true
	}

	// Otherwise, deny the request.
	return false
}

func (r *LeakyBucketRateLimiter) Tick() {
	// Add a token to the bucket.
	r.Tokens <- 1
}
