package main

import (
	"log"

	//"time"
	"github.com/gin-gonic/gin"

	"golang.org/x/time/rate"
)

func Limiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		limiter := rate.NewLimiter(10, 1)
		// Check if the request is allowed by the rate limiter.
		if limiter.Allow() {
			c.Next()
		} else {
			log.Println("Rate Limit Exceed")
		}

	}

}
