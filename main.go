package main

import (
	"log"
	"net/http"

	//"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func main() {
	r := gin.Default()
	r.Use(RateLimiter())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func RateLimiter() gin.HandlerFunc {
	limiter := rate.NewLimiter(1, 4)
	return func(c *gin.Context) {

		// Check if the request is allowed by the rate limiter.
		if limiter.Allow() {
			log.Println("Allow access")
			c.Next()
		} else {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message": "Limite exceed",
			})
		}

	}

}
