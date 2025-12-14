package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		duration := time.Since(startTime)
		log.Printf("%s %s - Status: %d - Duration: %v",
			c.Request.Method,
			c.Request.RequestURI,
			c.Writer.Status(),
			duration,
		)
	}
}
